package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"
	"sort"
	"strings"
	"sync"
	"text/template"
	"time"

	"github.com/taskcluster/httpbackoff/v3"
	"github.com/taskcluster/mozilla-history/workerpool"
	"github.com/taskcluster/slugid-go/slugid"
	tcclient "github.com/taskcluster/taskcluster/v48/clients/client-go"
	"github.com/taskcluster/taskcluster/v48/clients/client-go/tcqueue"
	"github.com/taskcluster/taskcluster/v48/clients/client-go/tcworkermanager"
)

type (
	Queue tcqueue.Queue
)

type WorkerInfo struct {
	WorkerPoolID   string
	Implementation string
	Version        string
	Details        map[string]string
	hasNoArtifacts bool
	isUnknown      bool
	Imageset       string
	TotalWorkers   int
	TotalCapacity  int
}

func (w *WorkerInfo) String() string {
	revision := ""
	engine := ""

	if w.Details["engine"] != "" {
		engine = fmt.Sprintf(" %v engine", w.Details["engine"])
	}

	if w.Details["revision"] != "" {
		revision = fmt.Sprintf(" (revision %v)", w.Details["revision"])
	}

	info := fmt.Sprintf("%-15s %-9s %-20s %-54s %-10s %-10s %-9s %s",
		w.Implementation,
		w.Version,
		engine,
		revision,
		w.Details["os"],
		w.Details["arch"],
		w.Details["go"],
		w.Details["error"],
	)

	return strings.Trim(info, " ")
}

func GetImageset(wp *tcworkermanager.WorkerPoolFullDefinition) string {
	var p = wp.ProviderID
	if p == "test-provisioner" || p == "no-provisioning-nope" || p == "dummy-test-provisioner" || p == "test-dummy-provisioner" {
		return "unknown"
	}

	type LaunchConfig struct {
		// AWS
		LaunchConfig struct {
			ImageId string
		} `json:"launchConfig"`

		// GCP
		Disks []struct {
			InitializeParams struct {
				SourceImage string `json:"sourceImage"`
			} `json:"initializeParams"`
		} `json:"disks"`

		// Azure
		StorageProfile struct {
			ImageReference struct {
				Id string `json:"id"`
			} `json:"imageReference"`
		} `json:"storageProfile"`
	}
	type Config struct {
		LaunchConfigs []LaunchConfig `json:"launchConfigs"`
	}
	var cfg Config
	if err := json.Unmarshal(wp.Config, &cfg); err != nil {
		return "unknown"
	}

	if len(cfg.LaunchConfigs) > 0 {
		launchCfg := cfg.LaunchConfigs[0]
		if launchCfg.LaunchConfig.ImageId != "" {
			return launchCfg.LaunchConfig.ImageId
		}
		if len(launchCfg.Disks) > 0 && launchCfg.Disks[0].InitializeParams.SourceImage != "" {
			return launchCfg.Disks[0].InitializeParams.SourceImage
		}
		if launchCfg.StorageProfile.ImageReference.Id != "" {
			return launchCfg.StorageProfile.ImageReference.Id
		}
	}

	return "unknown"
}

var (
	// Templates for README file
	readmeTpl string = `
{{- define "row" -}}
## {{ .Title }}

Total: ` + "`" + `{{ .Count }}` + "`" + `
{{ if gt (len .Versions) 1 }}
Count by version:

| Version | Count |
| :--- | ---: |
{{ range .Versions -}}
| {{ .Key }} | {{ .Value }} |
{{ end }}
{{- end }}
{{ if gt (len .Images) 1 }}
Count by image:

| Version | Count |
| :--- | ---: |
{{ range .Images -}}
| {{ .Key }} | {{ .Value }} |
{{ end }}
{{- end }}
{{if .Count }}
| Worker Pool | Implementation | Version {{ if .FullColumns }}| Engine | Revision | OS | Arch | GO {{ end }}| Total Workers | Total Capacity |
| --- | --- | --- {{ if .FullColumns }}| --- | --- | --- | --- | --- {{ end }}| ---: | ---: |
{{ range .Filtered -}}
| **{{ .WorkerPoolID }}** | {{ .Implementation }} | {{ or .Version .Details.error }} {{ if $.FullColumns }}| {{ or .Details.engine "-" }} | {{ or (slice .Details.revision 0 10) "-" }} | {{ or .Details.os "-" }} | {{ or .Details.arch "-" }} | {{ or .Details.go "-" }} {{ end }}| {{ .TotalWorkers }} | {{ .TotalCapacity }} |
{{end}}
{{- end -}}
{{end}}

# Worker Pool Versions

{{ range . }}
{{ template "row" . }}
{{ end }}

[^1]: Those are the pools whose tasks were claimed and resolved by a worker as expected, but the worker did not publish either artifact ` + "`public/logs/live_backing.log` nor `public/logs/chain_of_trust.log`" + `, which is the source used to identify the worker implementation.

[^2]: Probing task remains pending after two hours. Those are the pools that were not able to start any worker to claim the task within two hours.
`
)

func renderTemplate(data interface{}) string {
	t := template.Must(template.New("").Parse(readmeTpl))
	var content bytes.Buffer
	t.Execute(&content, data)
	return content.String()
}

var (
	// set during build with `-ldflags "-X main.revision=$(git rev-parse HEAD)"`
	revision string = ""
)

func FilenameEscape(raw string) (escaped string) {
	return strings.Replace(strings.Replace(raw, "*", "★", -1), "/", "⁄", -1)
}

// Call with no arguments -> New task group generated
// Call with one argument (taskGroupID) -> Report generated for previously created task group
//
// Expected workflow for this tool is to:
// 1. Run without arguments to generate probing tasks and get taskGroupId
// 2. wait 2h and run report collection for the given taskGroupId
func main() {

	queue := tcqueue.NewFromEnv()

	switch len(os.Args) {
	case 1:
		taskGroupID := slugid.Nice()
		fmt.Println("Task Group ID: " + taskGroupID)
		createTasks(queue, taskGroupID)
	case 2:
		taskGroupID := os.Args[1]
		taskIDs := taskIDsForTaskGroup(queue, taskGroupID)
		if len(taskIDs) == 0 {
			log.Fatalf("No tasks with taskGroupId %q", taskGroupID)
		}
		inspect(queue, taskIDs)
	default:
		log.Fatalf("Expected zero or one program arguments, but have %v: %q", len(os.Args)-1, os.Args[1:])
	}
}

func createTasks(queue *tcqueue.Queue, taskGroupID string) {
	if revision != "" {
		log.Printf("%v built from revision %v", os.Args[0], revision)
	}
	reportPrefix := os.Getenv("REPORT_PREFIX")
	schedulerId := os.Getenv("REPORT_SCHEDULER_ID")
	if reportPrefix == "" {
		log.Print("Please export env var REPORT_PREFIX to e.g.")
		log.Print("  https://github.com/taskcluster/mozilla-history/blob/master/WorkerVersions/")
		log.Fatal("  https://github.com/taskcluster/community-history/blob/master/WorkerVersions/")
	}
	if schedulerId == "" {
		log.Println("REPORT_SCHEDULER_ID not set, using '-' as default")
		schedulerId = "-"
	}
	created := time.Now()
	for _, wt := range AllWorkerTypes() {
		fmt.Println(wt)
		x := strings.Split(wt, "/")
		provisionerID := x[0]
		workerType := x[1]

		taskID := slugid.Nice()
		taskDef := &tcqueue.TaskDefinitionRequest{
			Created: tcclient.Time(created),
			// Deadline is set to +3h to give workers enough time to spin up and execute this payload
			// After this deadline task would be resolved with Deadline Exceeded exception
			// To be able to tell worker pools that don't have functioning workers and tasks that produce no artifacts
			// this tool needs to collect data from workers before deadline expires
			// if you run the create task part at t0, be sure to collect task group info before t0+3h to have relevant data
			Deadline:     tcclient.Time(created.Add(time.Hour * 3)),
			Dependencies: []string{},
			Expires:      tcclient.Time(created.Add(time.Hour * 24 * 30)),
			Extra:        json.RawMessage("{}"),
			Metadata: struct {
				Description string `json:"description"`
				Name        string `json:"name"`
				Owner       string `json:"owner"`
				Source      string `json:"source"`
			}{
				Name: "Checking worker version on " + provisionerID + "/" + workerType,
				Description: strings.Join([]string{
					`This task is a simple probe for checking the worker implementation (and version) that runs`,
					`on this worker type. This is routinely run in order to keep track of which worker`,
					`implementations have been deployed to which worker types, and which worker types may be`,
					`broken.`,
					``,
					`Note, the payload is intentionally invalid, so it is expected that the task resolves as`,
					"`malformed-payload`, but the log file should still contain enough information for the worker",
					`implementation to be determined.`,
					``,
					`The resulting verdict will be published [here](` + reportPrefix + provisionerID + `%E2%81%84` + workerType + `).`,
				}, "\n"),
				Owner:  "pmoore@mozilla.com",
				Source: "https://github.com/taskcluster/mozilla-history/tree/master/audit-worker-versions",
			},
			Payload: json.RawMessage(`{
					"fake": "fake arbitrary payload to cause malformed-payload exception",
					"see": "https://github.com/taskcluster/mozilla-history/tree/master/audit-worker-versions"
			}`),
			Priority:      "lowest",
			ProvisionerID: provisionerID,
			Requires:      "all-completed",
			Retries:       5,
			Routes:        []string{},
			SchedulerID:   schedulerId,
			Scopes:        []string{},
			Tags:          map[string]string{},
			TaskGroupID:   taskGroupID,
			WorkerType:    workerType,
		}
		tsr, err := queue.CreateTask(taskID, taskDef)
		fatalOnError(err)

		respJSON, err := json.MarshalIndent(tsr, "", "  ")
		fatalOnError(err)

		fmt.Println(string(respJSON))
	}
	fmt.Println("Tasks created, sealing task group...")
	tg, err := queue.SealTaskGroup(taskGroupID)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Task group sealed at: %v", tg.Sealed)
}

func inspect(queue *tcqueue.Queue, taskIDs []string) {
	workermanager := tcworkermanager.NewFromEnv()
	wp := workerpool.New(50)
	workers := make([]WorkerInfo, 0)

	wp.AddWork(func(wsc *workerpool.SubmitterContext) {
		for _, taskID := range taskIDs {
			wsc.RequestChannel <- func(taskID string) workerpool.Work {
				return func(workerId int) workerpool.Result {
					statusResponse, err := queue.Status(taskID)
					if err != nil {
						fmt.Println("Could not get status for task " + taskID)
						panic(err)
					}
					workerPoolID, workerInfo := show(queue, statusResponse)
					workerPool, err := workermanager.WorkerPool(workerPoolID)
					if err != nil {
						fmt.Println("Could not fetch workerPool " + workerPoolID)
					} else {
						workerInfo.Imageset = GetImageset(workerPool)
						workerInfo.TotalWorkers = int(workerPool.RunningCount) + int(workerPool.StoppedCount) +
							int(workerPool.StoppingCount) + int(workerPool.RequestedCount)
						workerInfo.TotalCapacity = int(workerPool.RunningCapacity) + int(workerPool.StoppedCapacity) +
							int(workerPool.StoppingCapacity) + int(workerPool.RequestedCapacity)
					}
					filename := FilenameEscape(workerPoolID)
					err = ioutil.WriteFile(filename, append([]byte(workerInfo.String()), '\n'), 0644)
					if err != nil {
						log.Fatalf("Error:\n%v", err)
					}
					fmt.Printf("%-70s %s\n", workerPoolID+":", &workerInfo)
					return workerInfo
				}
			}(taskID)
		}
	})
	wp.Done()
	wp.OnComplete(func(result workerpool.Result) {
		workers = append(workers, result.(WorkerInfo))
	})

	fmt.Printf("\nWriting README.md\n")
	writeReadme(workers)
	fmt.Println("Writing workers.json")
	writeSnapshot(workers)
}

func generateReadmeSection(title string, workers []WorkerInfo, filter func(WorkerInfo) bool) map[string]interface{} {
	filtered := make([]WorkerInfo, 0)
	versions := make(map[string]int)
	imagesets := make(map[string]int)

	for _, w := range workers {
		if filter(w) {
			filtered = append(filtered, w)
			versions[w.Version]++
			imagesets[w.Imageset]++
		}
	}

	sort.Slice(filtered, func(i, j int) bool {
		return strings.Compare(filtered[i].WorkerPoolID, filtered[j].WorkerPoolID) < 0
	})

	type kv struct {
		Key   string
		Value int
	}
	var sortedVersions []kv
	for k, v := range versions {
		sortedVersions = append(sortedVersions, kv{k, v})
	}
	sort.Slice(sortedVersions, func(i, j int) bool {
		return strings.Compare(sortedVersions[i].Key, sortedVersions[j].Key) < 0
	})

	var images []kv
	for k, v := range imagesets {
		images = append(images, kv{k, v})
	}

	return map[string]interface{}{
		"FullColumns": title == "Generic Worker",
		"Filtered":    filtered,
		"Count":       len(filtered),
		"Versions":    sortedVersions,
		"Images":      images,
		"Title":       title,
	}
}

func writeReadme(workers []WorkerInfo) {
	filename := "README.md"

	sections := [5]map[string]interface{}{
		generateReadmeSection("Generic Worker", workers, func(w WorkerInfo) bool { return w.Implementation == "generic-worker" }),
		generateReadmeSection("Docker Worker", workers, func(w WorkerInfo) bool { return w.Implementation == "docker-worker" }),
		generateReadmeSection("Script Worker", workers, func(w WorkerInfo) bool { return strings.Contains(w.Implementation, "Scriptworker") }),
		generateReadmeSection("No artifacts found [^1]", workers, func(w WorkerInfo) bool { return w.hasNoArtifacts }),
		generateReadmeSection("Version not determined [^2]", workers, func(w WorkerInfo) bool { return w.isUnknown }),
	}

	contents := renderTemplate(sections)

	err := ioutil.WriteFile(filename, []byte(contents), 0644)
	if err != nil {
		log.Fatalf("Error:\n%v", err)
	}
}

func writeSnapshot(workers []WorkerInfo) {
	filename := "workers.json"

	sort.Slice(workers, func(i, j int) bool {
		return strings.Compare(workers[i].WorkerPoolID, workers[j].WorkerPoolID) <= 0
	})

	contents, err := json.MarshalIndent(workers, "", " ")
	if err != nil {
		log.Fatalf("Error:\n%v", err)
	}

	err = ioutil.WriteFile(filename, []byte(contents), 0644)
	if err != nil {
		log.Fatalf("Error:\n%v", err)
	}
}

func mustCompileToRawMessage(data interface{}) *json.RawMessage {
	bytes, err := json.Marshal(data)
	fatalOnError(err)
	var JSON json.RawMessage
	err = json.Unmarshal(bytes, &JSON)
	fatalOnError(err)
	return &JSON
}

func fatalOnError(err error) {
	if err != nil {
		log.Fatalf("Error:\n%v", err)
	}
}

func NewQueue() *Queue {
	q := tcqueue.NewFromEnv()
	Q := Queue(*q)
	return &Q
}

func AllWorkerTypes() []string {
	uniqueWorkerTypes := map[string]bool{}
	q := NewQueue()
	provisioners := q.AllProvisionerIDs()
	workerTypes := make([][]string, len(provisioners), len(provisioners))
	var wg sync.WaitGroup
	for i, p := range provisioners {
		if p == "test-provisioner" || p == "no-provisioning-nope" || p == "dummy-test-provisioner" || p == "test-dummy-provisioner" {
			continue
		}
		wg.Add(1)
		go func(p string, i int) {
			defer wg.Done()
			provWorkerTypes := q.ProvisionerWorkerTypes(p)
			workerTypes[i] = make([]string, len(provWorkerTypes), len(provWorkerTypes))
			for j, wt := range provWorkerTypes {
				workerTypes[i][j] = p + "/" + wt
			}
		}(p, i)
	}
	wg.Wait()
	for _, p := range workerTypes {
		for _, wt := range p {
			if !strings.HasPrefix(wt, "null-provisioner/test-") {
				uniqueWorkerTypes[wt] = true
			}
		}
	}

	// Now merge in known worker types according to Worker Manager
	workermanager := tcworkermanager.NewFromEnv()
	continuationToken := ""
	for {
		workerPools, err := workermanager.ListWorkerPools(continuationToken, "")
		if err != nil {
			panic(err)
		}
		for _, workerPool := range workerPools.WorkerPools {
			uniqueWorkerTypes[workerPool.WorkerPoolID] = true
		}
		continuationToken = workerPools.ContinuationToken
		if continuationToken == "" {
			break
		}
	}

	keys := make([]string, 0, len(uniqueWorkerTypes))
	for key := range uniqueWorkerTypes {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	return keys
}

func (q *Queue) AllProvisionerIDs() []string {
	Q := tcqueue.Queue(*q)
	provisioners := []string{}
	var r *tcqueue.ListProvisionersResponse
	ct := ""
	for r == nil || r.ContinuationToken != "" {
		var err error
		r, err = (&Q).ListProvisioners(ct, "")
		if err != nil {
			panic(err)
		}
		ct = r.ContinuationToken
		for _, p := range r.Provisioners {
			provisioners = append(provisioners, p.ProvisionerID)
		}
	}
	return provisioners
}

func (q *Queue) ProvisionerWorkerTypes(provisionerID string) []string {
	Q := tcqueue.Queue(*q)
	workerTypes := []string{}
	var r *tcqueue.ListWorkerTypesResponse
	ct := ""
	for r == nil || r.ContinuationToken != "" {
		var err error
		r, err = (&Q).ListWorkerTypes(provisionerID, ct, "")
		if err != nil {
			panic(err)
		}
		ct = r.ContinuationToken
		for _, p := range r.WorkerTypes {
			workerTypes = append(workerTypes, p.WorkerType)
		}
	}
	return workerTypes
}

func show(queue *tcqueue.Queue, t *tcqueue.TaskStatusResponse) (workerPoolID string, workerInfo WorkerInfo) {
	workerPoolID = t.Status.ProvisionerID + "/" + t.Status.WorkerType
	workerInfo.WorkerPoolID = workerPoolID
	if workerInfo.Details == nil {
		workerInfo.Details = map[string]string{}
	}
	if t.Status.State == "pending" {
		// We schedule task with deadline for 3h and running report generation anywhere before that time
		// In case some pools were not able to claim task within 3h we would consider this pool to have unknown workers
		// which could probably tell that something is wrong with configuration of this pool
		workerInfo.Details["error"] = "Version not determined; task not (yet) claimed"
		workerInfo.isUnknown = true
		return
	}
	var resp *http.Response
	artifactFound := ""
	for _, artifact := range []string{
		"public/logs/live_backing.log",
		"public/logs/chain_of_trust.log",
	} {
		logURL, err := queue.GetLatestArtifact_SignedURL(t.Status.TaskID, artifact, time.Hour)
		if err != nil {
			log.Fatal(2, err)
		}
		resp, _, err = httpbackoff.Get(logURL.String())
		if err == nil {
			artifactFound = artifact
			break
		}
		switch e := err.(type) {
		case httpbackoff.BadHttpResponseCode:
			if e.HttpResponseCode == 404 {
				continue
			}
		}
		log.Fatal(3, err)
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Print("*** ")
	}
	logContent := string(data)
	switch true {
	case strings.Contains(logContent, "Worker Version: "):
		re := regexp.MustCompile(`(?m)Worker Version: (.*)`)
		dw := re.FindStringSubmatch(logContent)
		val := "<UNKNOWN>"
		if len(dw) > 1 {
			val = dw[1]
		}
		workerInfo.Implementation = "docker-worker"
		workerInfo.Version = val
	case strings.Contains(logContent, "Worker Node Type:"):
		workerInfo.Implementation = "docker-worker"
		workerInfo.Version = "unknown version"
	case strings.Contains(logContent, `"generic-worker":`):
		workerInfo.Implementation = "generic-worker"
		for _, t := range []struct {
			regex string
			key   string
		}{
			{`"engine": "(.*)"`, "engine"},
			{`"https://github.com/taskcluster/.*/releases/tag/v([^"]*)"`, "version"},
			{`"revision": "([0-9a-f]{40})"`, "revision"},
			{`"go-os": "(.*)"`, "os"},
			{`"go-arch": "(.*)"`, "arch"},
			{`"go-version": "go(.*)"`, "go"},
		} {
			re := regexp.MustCompile(t.regex)
			gw := re.FindStringSubmatch(logContent)
			val := "<UNKNOWN>"
			if len(gw) > 1 {
				val = gw[1]
			}
			if t.key == "version" {
				workerInfo.Version = val
			} else {
				workerInfo.Details[t.key] = val
			}
		}
	case strings.Contains(logContent, "Task not successful due to following exception"):
		workerInfo.Implementation = "generic-worker"
		workerInfo.Version = "unknown version"
	case strings.Contains(logContent, `not allowed at task.payload.features`):
		workerInfo.Implementation = "Taskcluster Worker"
	case strings.Contains(logContent, `raise TaskVerificationError`):
		workerInfo.Implementation = "Scriptworker"
	case strings.Contains(logContent, `KeyError: 'artifacts_deps'`):
		workerInfo.Implementation = "Other Scriptworker"
	case artifactFound == "":
		workerInfo.hasNoArtifacts = true
		workerInfo.Details["error"] = "No artifacts found"
	case artifactFound == "public/logs/chain_of_trust.log":
		workerInfo.Implementation = "Scriptworker Chain of Trust"
	case strings.Contains(logContent, `os.environ.get('GITHUB_HEAD_REPO_URL', decision_json['payload']['env']['GITHUB_HEAD_REPO_URL'])`):
		workerInfo.Implementation = "Scriptworker Deepspeech" // not deepspeach?
	case strings.Contains(logContent, `balrog`):
		workerInfo.Implementation = "Scriptworker Balrog"
	case strings.Contains(logContent, `bouncerscript`):
		workerInfo.Implementation = "Scriptworker Bouncer Script"
	case strings.Contains(logContent, `beetmover`):
		workerInfo.Implementation = "Scriptworker Beetmover"
	case strings.Contains(logContent, `scriptworker`):
		workerInfo.Implementation = "Scriptworker"
	default:
		workerInfo.Details["error"] = fmt.Sprintf("Cannot determine worker implementation from log:\n%v", logContent[:1024])
		workerInfo.isUnknown = true
	}
	return
}

func taskIDsForTaskGroup(queue *tcqueue.Queue, taskGroupID string) []string {
	taskIDs := []string{}
	continuationToken := ""
	for {
		ltgr, err := queue.ListTaskGroup(taskGroupID, continuationToken, "")
		if err != nil {
			panic(err)
		}
		for _, tdas := range ltgr.Tasks {
			taskIDs = append(taskIDs, tdas.Status.TaskID)
		}
		continuationToken = ltgr.ContinuationToken
		if continuationToken == "" {
			break
		}
	}
	return taskIDs
}

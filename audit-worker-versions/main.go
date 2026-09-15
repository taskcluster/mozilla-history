package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
	"sync"
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
	WorkerPoolID             string
	Implementation           string
	Version                  string
	Details                  map[string]string
	hasNoArtifacts           bool
	isUnknown                bool
	Imageset                 string
	ImageStatus              string `json:",omitempty"`
	ProviderID               string `json:",omitempty"`
	ConfiguredMinCapacity    *int   `json:",omitempty"`
	ConfiguredMaxCapacity    *int   `json:",omitempty"`
	CapacityPerWorker        *int   `json:",omitempty"`
	ConfiguredMinWorkers     *int   `json:",omitempty"`
	ConfiguredMaxWorkers     *int   `json:",omitempty"`
	WorkerManagerLookupError string `json:",omitempty"`
	// Retained only to recognize snapshots created before configured capacity
	// fields were collected. New snapshots leave these fields unset.
	LegacyTotalWorkers  *int `json:"TotalWorkers,omitempty"`
	LegacyTotalCapacity *int `json:"TotalCapacity,omitempty"`
}

const (
	imageStatusKnown         = "known"
	imageStatusNotApplicable = "not-applicable"
	imageStatusNotDetermined = "not-determined"
	imageStatusUnavailable   = "unavailable"
	standaloneProviderID     = "standalone"
	standaloneWorkerState    = "standalone"
)

type workerManagerClient interface {
	WorkerPool(workerPoolID string) (*tcworkermanager.WorkerPoolFullDefinition, error)
	ListWorkers(provisionerID, workerType, continuationToken, limit, quarantined, workerState string) (*tcworkermanager.ListWorkersResponse, error)
}

type WorkerSnapshot struct {
	GeneratedAt    time.Time    `json:"generatedAt"`
	ProbeStartedAt time.Time    `json:"probeStartedAt"`
	TaskGroupID    string       `json:"taskGroupId"`
	Workers        []WorkerInfo `json:"workers"`
}

type taskGroupProgress struct {
	Total    int
	Terminal int
	States   map[string]int
}

func (progress taskGroupProgress) complete() bool {
	return progress.Total > 0 && progress.Terminal == progress.Total
}

func isTerminalTaskState(state string) bool {
	switch state {
	case "completed", "failed", "exception":
		return true
	default:
		return false
	}
}

func summarizeTaskGroup(tasks []tcqueue.TaskDefinitionAndStatus) taskGroupProgress {
	progress := taskGroupProgress{States: map[string]int{}}
	for _, task := range tasks {
		state := task.Status.State
		progress.Total++
		progress.States[state]++
		if isTerminalTaskState(state) {
			progress.Terminal++
		}
	}
	return progress
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

type workerPoolLaunchConfig struct {
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
	ArmDeployment struct {
		Parameters struct {
			ImageID struct {
				Value string `json:"value"`
			} `json:"imageId"`
		} `json:"parameters"`
	} `json:"armDeployment"`

	// Older Worker Manager configurations stored this directly on the
	// launch config. Newer configurations nest it under workerManager.
	CapacityPerInstance *int `json:"capacityPerInstance"`
	WorkerManager       struct {
		CapacityPerInstance *int `json:"capacityPerInstance"`
	} `json:"workerManager"`
}

type workerPoolConfig struct {
	MinCapacity   *int                     `json:"minCapacity"`
	MaxCapacity   *int                     `json:"maxCapacity"`
	LaunchConfigs []workerPoolLaunchConfig `json:"launchConfigs"`
}

func parseWorkerPoolConfig(wp *tcworkermanager.WorkerPoolFullDefinition) (workerPoolConfig, error) {
	var cfg workerPoolConfig
	if err := json.Unmarshal(wp.Config, &cfg); err != nil {
		return workerPoolConfig{}, err
	}
	return cfg, nil
}

func getImageset(providerID string, cfg workerPoolConfig) string {
	if providerID == "test-provisioner" || providerID == "no-provisioning-nope" || providerID == "dummy-test-provisioner" || providerID == "test-dummy-provisioner" {
		return "unknown"
	}

	imagesMap := map[string]struct{}{}
	for _, launchCfg := range cfg.LaunchConfigs {
		imagesMap[launchCfg.LaunchConfig.ImageId] = struct{}{}
		for _, disk := range launchCfg.Disks {
			imagesMap[disk.InitializeParams.SourceImage] = struct{}{}
		}
		imagesMap[launchCfg.StorageProfile.ImageReference.Id] = struct{}{}
		imagesMap[launchCfg.ArmDeployment.Parameters.ImageID.Value] = struct{}{}
	}
	// remove empty image name ""
	delete(imagesMap, "")
	imagesSlice := make([]string, 0, len(imagesMap))
	for image := range imagesMap {
		imagesSlice = append(imagesSlice, image)
	}
	sort.Strings(imagesSlice)
	sortedImages := strings.Join(imagesSlice, ",")
	if sortedImages == "" {
		return "unknown"
	}
	return sortedImages
}

func ceilDivide(value, divisor int) int {
	return (value + divisor - 1) / divisor
}

func enrichWorkerInfo(workerInfo *WorkerInfo, wp *tcworkermanager.WorkerPoolFullDefinition) {
	workerInfo.ProviderID = wp.ProviderID
	cfg, err := parseWorkerPoolConfig(wp)
	if err != nil {
		workerInfo.Imageset = "unknown"
		workerInfo.ImageStatus = imageStatusUnavailable
		workerInfo.WorkerManagerLookupError = "Could not parse Worker Manager configuration: " + err.Error()
		return
	}

	workerInfo.Imageset = getImageset(wp.ProviderID, cfg)
	workerInfo.ImageStatus = imageStatusKnown
	if workerInfo.Imageset == "unknown" {
		workerInfo.ImageStatus = imageStatusNotDetermined
	}
	workerInfo.ConfiguredMinCapacity = cfg.MinCapacity
	workerInfo.ConfiguredMaxCapacity = cfg.MaxCapacity

	capacities := map[int]struct{}{}
	for _, launchConfig := range cfg.LaunchConfigs {
		capacity := launchConfig.CapacityPerInstance
		if launchConfig.WorkerManager.CapacityPerInstance != nil {
			capacity = launchConfig.WorkerManager.CapacityPerInstance
		}
		if capacity != nil && *capacity > 0 {
			capacities[*capacity] = struct{}{}
		}
	}

	if len(capacities) != 1 || cfg.MinCapacity == nil || cfg.MaxCapacity == nil {
		return
	}
	for capacity := range capacities {
		capacityPerWorker := capacity
		minWorkers := ceilDivide(*cfg.MinCapacity, capacity)
		maxWorkers := ceilDivide(*cfg.MaxCapacity, capacity)
		workerInfo.CapacityPerWorker = &capacityPerWorker
		workerInfo.ConfiguredMinWorkers = &minWorkers
		workerInfo.ConfiguredMaxWorkers = &maxWorkers
	}
}

func workerPoolIsStandalone(workermanager workerManagerClient, workerPoolID string) (bool, error) {
	parts := strings.SplitN(workerPoolID, "/", 2)
	if len(parts) != 2 {
		return false, fmt.Errorf("invalid worker pool ID %q", workerPoolID)
	}

	foundStandaloneWorker := false
	continuationToken := ""
	for {
		workers, err := workermanager.ListWorkers(parts[0], parts[1], continuationToken, "", "", "")
		if err != nil {
			return false, err
		}
		for _, worker := range workers.Workers {
			if worker.State == standaloneWorkerState {
				foundStandaloneWorker = true
			}
		}
		continuationToken = workers.ContinuationToken
		if continuationToken == "" {
			break
		}
	}
	return foundStandaloneWorker, nil
}

func lookupAndEnrichWorkerInfo(workerInfo *WorkerInfo, workermanager workerManagerClient) {
	workerPool, lookupErr := workermanager.WorkerPool(workerInfo.WorkerPoolID)
	if lookupErr == nil {
		enrichWorkerInfo(workerInfo, workerPool)
		return
	}

	standalone, standaloneErr := workerPoolIsStandalone(workermanager, workerInfo.WorkerPoolID)
	if standaloneErr == nil && standalone {
		workerInfo.ProviderID = standaloneProviderID
		workerInfo.Imageset = "unknown"
		workerInfo.ImageStatus = imageStatusNotApplicable
		return
	}

	workerInfo.Imageset = "unknown"
	workerInfo.ImageStatus = imageStatusUnavailable
	workerInfo.WorkerManagerLookupError = lookupErr.Error()
	if standaloneErr != nil {
		workerInfo.WorkerManagerLookupError += "; could not check for standalone workers: " + standaloneErr.Error()
	}
}

var (
	// set during build with `-ldflags "-X main.revision=$(git rev-parse HEAD)"`
	revision string = ""
)

var (
	outputDir = "WorkerVersions"
)

func FilenameEscape(raw string) (escaped string) {
	return strings.Replace(strings.Replace(raw, "*", "★", -1), "/", "⁄", -1)
}

func EmptyDirectory(dir string) {
	err := os.RemoveAll(dir)
	if err != nil {
		panic(err)
	}
}

func WriteFile(path string, content []byte) {
	err := os.MkdirAll(filepath.Dir(path), 0755)
	if err != nil {
		panic(err)
	}

	err = os.WriteFile(path, content, 0644)
	if err != nil {
		log.Fatalf("Error:\n%v", err)
	}
}

// Call with no arguments -> New task group generated
// Call with one argument (taskGroupID) -> Report generated for previously created task group
// Call with "status taskGroupID" -> Task group progress reported
//
// Expected workflow for this tool is to:
// 1. Run without arguments to generate probing tasks and get taskGroupId
// 2. wait 2h and run report collection for the given taskGroupId
//
// Files are written to the WorkerVersions directory
func main() {
	if len(os.Args) >= 2 && os.Args[1] == "render" {
		if len(os.Args) < 3 || len(os.Args) > 4 {
			log.Fatal("Usage: audit-worker-versions render INPUT_JSON [OUTPUT_MARKDOWN]")
		}
		snapshot, err := readSnapshot(os.Args[2])
		fatalOnError(err)
		contents := renderReadme(snapshot)
		if len(os.Args) == 4 {
			WriteFile(os.Args[3], []byte(contents))
		} else {
			fmt.Print(contents)
		}
		return
	}

	queue := tcqueue.NewFromEnv()
	if len(os.Args) >= 2 && os.Args[1] == "status" {
		if len(os.Args) != 3 {
			log.Fatal("Usage: audit-worker-versions status TASK_GROUP_ID")
		}
		progress, err := taskGroupStatus(queue, os.Args[2])
		fatalOnError(err)
		printTaskGroupProgress(os.Args[2], progress)
		if !progress.complete() {
			os.Exit(3)
		}
		return
	}

	switch len(os.Args) {
	case 1:
		taskGroupID := slugid.Nice()
		fmt.Println("Task Group ID: " + taskGroupID)
		createTasks(queue, taskGroupID)
	case 2:
		taskGroupID := os.Args[1]
		EmptyDirectory(outputDir)

		taskIDs := taskIDsForTaskGroup(queue, taskGroupID)
		if len(taskIDs) == 0 {
			log.Fatalf("No tasks with taskGroupId %q", taskGroupID)
		}
		inspect(queue, taskGroupID, taskIDs)
	default:
		log.Fatalf("Expected zero or one program arguments, but have %v: %q", len(os.Args)-1, os.Args[1:])
	}
}

func taskGroupStatus(queue *tcqueue.Queue, taskGroupID string) (taskGroupProgress, error) {
	tasks := []tcqueue.TaskDefinitionAndStatus{}
	continuationToken := ""
	for {
		response, err := queue.ListTaskGroup(taskGroupID, continuationToken, "")
		if err != nil {
			return taskGroupProgress{}, err
		}
		tasks = append(tasks, response.Tasks...)
		continuationToken = response.ContinuationToken
		if continuationToken == "" {
			break
		}
	}
	if len(tasks) == 0 {
		return taskGroupProgress{}, fmt.Errorf("no tasks with taskGroupId %q", taskGroupID)
	}
	return summarizeTaskGroup(tasks), nil
}

func printTaskGroupProgress(taskGroupID string, progress taskGroupProgress) {
	stateOrder := []string{"completed", "failed", "exception", "running", "pending", "unscheduled"}
	states := make([]string, 0, len(progress.States))
	knownStates := map[string]bool{}
	for _, state := range stateOrder {
		knownStates[state] = true
		if count := progress.States[state]; count > 0 {
			states = append(states, fmt.Sprintf("%s=%d", state, count))
		}
	}
	extraStates := []string{}
	for state, count := range progress.States {
		if !knownStates[state] {
			extraStates = append(extraStates, fmt.Sprintf("%s=%d", state, count))
		}
	}
	sort.Strings(extraStates)
	states = append(states, extraStates...)
	fmt.Printf("Task group %s: %d/%d terminal (%s)\n", taskGroupID, progress.Terminal, progress.Total, strings.Join(states, ", "))
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

func inspect(queue *tcqueue.Queue, taskGroupID string, taskIDs []string) {
	EmptyDirectory(outputDir)
	probeTask, err := queue.Task(taskIDs[0])
	fatalOnError(err)
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
					lookupAndEnrichWorkerInfo(&workerInfo, workermanager)
					filename := filepath.Join(outputDir, FilenameEscape(workerPoolID))
					WriteFile(filename, append([]byte(workerInfo.String()), '\n'))
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

	snapshot := WorkerSnapshot{
		GeneratedAt:    time.Now().UTC(),
		ProbeStartedAt: time.Time(probeTask.Created).UTC(),
		TaskGroupID:    taskGroupID,
		Workers:        workers,
	}

	fmt.Printf("\nWriting README.md\n")
	writeReadme(snapshot)
	fmt.Println("Writing workers.json")
	writeSnapshot(snapshot)
}

func writeSnapshot(snapshot WorkerSnapshot) {
	filename := filepath.Join(outputDir, "workers.json")

	sort.Slice(snapshot.Workers, func(i, j int) bool {
		return strings.Compare(snapshot.Workers[i].WorkerPoolID, snapshot.Workers[j].WorkerPoolID) <= 0
	})

	contents, err := json.MarshalIndent(snapshot, "", " ")
	if err != nil {
		log.Fatalf("Error:\n%v", err)
	}

	WriteFile(filename, []byte(contents))
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
	if !taskWasClaimed(t.Status.Runs) {
		workerInfo.Details["error"] = "Version not determined; task was not claimed"
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
	data, err := io.ReadAll(resp.Body)
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

func taskWasClaimed(runs []tcqueue.RunInformation) bool {
	for _, run := range runs {
		if run.WorkerID != "" || run.WorkerGroup != "" || !time.Time(run.Started).IsZero() {
			return true
		}
	}
	return false
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

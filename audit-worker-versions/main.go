package main

import (
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
	"time"

	"github.com/taskcluster/httpbackoff/v3"
	"github.com/taskcluster/slugid-go/slugid"
	tcclient "github.com/taskcluster/taskcluster/v30/clients/client-go"
	"github.com/taskcluster/taskcluster/v30/clients/client-go/tcqueue"
	"github.com/taskcluster/taskcluster/v30/clients/client-go/tcworkermanager"
)

type (
	Queue tcqueue.Queue
)

const waitTimeMinutes = 90

var (
	// set during build with `-ldflags "-X main.revision=$(git rev-parse HEAD)"`
	revision string = ""
)

func FilenameEscape(raw string) (escaped string) {
	return strings.Replace(strings.Replace(raw, "*", "★", -1), "/", "⁄", -1)
}

func main() {
	submittedTasks := false
	if revision != "" {
		log.Printf("%v built from revision %v", os.Args[0], revision)
	}
	reportPrefix := os.Getenv("REPORT_PREFIX")
	if reportPrefix == "" {
		log.Fatal("Please export env var REPORT_PREFIX to e.g. https://github.com/taskcluster/mozilla-history/blob/master/WorkerVersions/")
	}
	tasks := map[string]tcqueue.TaskRun{}
	queue := tcqueue.NewFromEnv()
	taskGroupID := slugid.Nice()
	created := time.Now()
	poolsWithTasksWithoutLogs := []string{}
	poolsWithoutTasks := []string{}
	for _, wt := range AllWorkerTypes() {
		fmt.Println(wt)
		x := strings.Split(wt, "/")
		provisionerID := x[0]
		workerType := x[1]

		continuationToken := ""
		foundLog := false
		foundTask := false
		for {
			workers, err := queue.ListWorkers(provisionerID, workerType, continuationToken, "", "")
			if err != nil {
				panic(err)
			}
			for _, worker := range workers.Workers {
				_, err := queue.Status(worker.LatestTask.TaskID)
				if err != nil {
					continue
				}
				foundTask = true
				for _, artifact := range []string{
					"public/logs/live_backing.log",
					"public/logs/chain_of_trust.log",
				} {
					logURL, err := queue.GetLatestArtifact_SignedURL(worker.LatestTask.TaskID, artifact, time.Hour)
					if err != nil {
						panic(err)
					}
					_, _, err = httpbackoff.Get(logURL.String())
					if err != nil {
						continue
					}
				}
				tasks[wt] = tcqueue.TaskRun{
					TaskID: worker.LatestTask.TaskID,
					RunID:  worker.LatestTask.RunID,
				}
				foundLog = true
				break
			}
			continuationToken = workers.ContinuationToken
			if continuationToken == "" || foundLog {
				break
			}
		}

		if !foundTask {
			poolsWithoutTasks = append(poolsWithoutTasks, wt)
		}
		if foundTask && !foundLog {
			poolsWithTasksWithoutLogs = append(poolsWithTasksWithoutLogs, wt)
		}

		if !foundLog {
			fmt.Println("No existing task found for " + wt)
			if os.Getenv("TASKCLUSTER_CLIENT_ID") != "" {
				taskID := slugid.Nice()
				tasks[wt] = tcqueue.TaskRun{
					TaskID: taskID,
					RunID:  0,
				}
				taskDef := &tcqueue.TaskDefinitionRequest{
					Created:      tcclient.Time(created),
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
					SchedulerID:   "-",
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
				submittedTasks = true
			}
		}
	}

	if submittedTasks {
		fmt.Println("Task Group ID: " + taskGroupID)
		fmt.Printf("Now sleeping for %v mins to wait for tasks to run...\n", waitTimeMinutes)
		// Leave plenty of time for tasks to run
		time.Sleep(waitTimeMinutes * time.Minute)
	}

	for _, task := range tasks {
		statusResponse, err := queue.Status(task.TaskID)
		if err != nil {
			fmt.Println("Could not get status for task " + task.TaskID)
			panic(err)
		}
		workerPoolID, versionInfo := show(queue, statusResponse)
		filename := FilenameEscape(workerPoolID)
		err = ioutil.WriteFile(filename, append([]byte(versionInfo), '\n'), 0644)
		if err != nil {
			log.Fatalf("Error:\n%v", err)
		}
		name := workerPoolID + ":                                                                          "
		fmt.Print(name[:75])
		fmt.Println(versionInfo)
	}

	fmt.Printf("Worker Types without recent tasks: %v\n", poolsWithoutTasks)
	fmt.Printf("Worker Types with recent tasks without required logs: %v\n", poolsWithTasksWithoutLogs)
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

func show(queue *tcqueue.Queue, t *tcqueue.TaskStatusResponse) (workerPoolID, versionInfo string) {
	workerPoolID = t.Status.ProvisionerID + "/" + t.Status.WorkerType
	if t.Status.State == "pending" {
		versionInfo = fmt.Sprintf("Version not determined; probing task remains pending after %v minutes.", waitTimeMinutes)
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
		versionInfo = "docker-worker - version " + val
	case strings.Contains(logContent, "Worker Node Type:"):
		versionInfo = "docker-worker - unknown version"
	case strings.Contains(logContent, `"generic-worker":`):
		versionInfo = "generic-worker"
		for _, t := range []struct {
			regex string
			text  string
		}{
			{`"engine": "(.*)"`, `%v engine`},
			{`"https://github.com/taskcluster/.*/releases/tag/v([^"]*)"`, `%v`},
			{`"revision": "([0-9a-f]{40})"`, `(revision %v)`},
			{`"go-os": "(.*)"`, `%v`},
			{`"go-arch": "(.*)"`, `%v`},
			{`"go-version": "go(.*)"`, `Go %v`},
		} {
			re := regexp.MustCompile(t.regex)
			gw := re.FindStringSubmatch(logContent)
			val := "<UNKNOWN>"
			if len(gw) > 1 {
				val = gw[1]
			}
			versionInfo += " " + fmt.Sprintf(t.text, val)
		}
	case strings.Contains(logContent, "Task not successful due to following exception"):
		versionInfo = "generic-worker - unknown version"
	case strings.Contains(logContent, `not allowed at task.payload.features`):
		versionInfo = "Taskcluster Worker"
	case strings.Contains(logContent, `raise TaskVerificationError`):
		versionInfo = "Scriptworker"
	case strings.Contains(logContent, `KeyError: 'artifacts_deps'`):
		versionInfo = "Other Scriptworker"
	case artifactFound == "":
		versionInfo = "No artifacts found"
	case artifactFound == "public/logs/chain_of_trust.log":
		versionInfo = "Scriptworker Chain of Trust"
	case strings.Contains(logContent, `os.environ.get('GITHUB_HEAD_REPO_URL', decision_json['payload']['env']['GITHUB_HEAD_REPO_URL'])`):
		versionInfo = "Scriptworker Deepspeech"
	case strings.Contains(logContent, `balrog`):
		versionInfo = "Scriptworker Balrog"
	case strings.Contains(logContent, `bouncerscript`):
		versionInfo = "Scriptworker Bouncer Script"
	case strings.Contains(logContent, `beetmover`):
		versionInfo = "Scriptworker Beetmover"
	case strings.Contains(logContent, `scriptworker`):
		versionInfo = "Scriptworker"
	default:
		versionInfo = "Unknown"
		log.Printf("Cannot determine worker implementation from log:\n%v", logContent[:1024])
	}
	return
}

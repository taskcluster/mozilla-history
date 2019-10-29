package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"sort"
	"strings"
	"sync"
	"time"

	"github.com/taskcluster/httpbackoff"
	"github.com/taskcluster/slugid-go/slugid"
	tcclient "github.com/taskcluster/taskcluster-client-go"
	"github.com/taskcluster/taskcluster-client-go/tcawsprovisioner"
	"github.com/taskcluster/taskcluster-client-go/tcqueue"
	"github.com/taskcluster/taskcluster-client-go/tcworkermanager"
)

type (
	// This schema defines the structure of the `payload` property referred to in a
	// Taskcluster Task definition.
	GenericWorkerPayload struct {

		// Artifacts to be published.
		//
		// Since: generic-worker 1.0.0
		Artifacts []struct {

			// Explicitly set the value of the HTTP `Content-Type` response header when the artifact(s)
			// is/are served over HTTP(S). If not provided (this property is optional) the worker will
			// guess the content type of artifacts based on the filename extension of the file storing
			// the artifact content. It does this by looking at the system filename-to-mimetype mappings
			// defined in the Windows registry. Note, setting `contentType` on a directory artifact will
			// apply the same contentType to all files contained in the directory.
			//
			// See [mime.TypeByExtension](https://godoc.org/mime#TypeByExtension).
			//
			// Since: generic-worker 10.4.0
			ContentType string `json:"contentType,omitempty"`

			// Date when artifact should expire must be in the future, no earlier than task deadline, but
			// no later than task expiry. If not set, defaults to task expiry.
			//
			// Since: generic-worker 1.0.0
			Expires tcclient.Time `json:"expires,omitempty"`

			// Name of the artifact, as it will be published. If not set, `path` will be used.
			// Conventionally (although not enforced) path elements are forward slash separated. Example:
			// `public/build/a/house`. Note, no scopes are required to read artifacts beginning `public/`.
			// Artifact names not beginning `public/` are scope-protected (caller requires scopes to
			// download the artifact). See the Queue documentation for more information.
			//
			// Since: generic-worker 8.1.0
			Name string `json:"name,omitempty"`

			// Relative path of the file/directory from the task directory. Note this is not an absolute
			// path as is typically used in docker-worker, since the absolute task directory name is not
			// known when the task is submitted. Example: `dist\regedit.exe`. It doesn't matter if
			// forward slashes or backslashes are used.
			//
			// Since: generic-worker 1.0.0
			Path string `json:"path"`

			// Artifacts can be either an individual `file` or a `directory` containing
			// potentially multiple files with recursively included subdirectories.
			//
			// Since: generic-worker 1.0.0
			//
			// Possible values:
			//   * "file"
			//   * "directory"
			Type string `json:"type"`
		} `json:"artifacts,omitempty"`

		// One entry per command (consider each entry to be interpreted as a full line of
		// a Windows™ .bat file). For example:
		// ```
		// [
		//   "set",
		//   "echo hello world > hello_world.txt",
		//   "set GOPATH=C:\\Go"
		// ]
		// ```
		//
		// Since: generic-worker 0.0.1
		Command []string `json:"command"`

		// Env vars must be string to __string__ mappings (not number or boolean). For example:
		// ```
		// {
		//   "PATH": "C:\\Windows\\system32;C:\\Windows",
		//   "GOOS": "windows",
		//   "FOO_ENABLE": "true",
		//   "BAR_TOTAL": "3"
		// }
		// ```
		//
		// Since: generic-worker 0.0.1
		Env map[string]string `json:"env,omitempty"`

		// Feature flags enable additional functionality.
		//
		// Since: generic-worker 5.3.0
		Features struct {

			// An artifact named `public/chainOfTrust.json.asc` should be generated
			// which will include information for downstream tasks to build a level
			// of trust for the artifacts produced by the task and the environment
			// it ran in.
			//
			// Since: generic-worker 5.3.0
			ChainOfTrust bool `json:"chainOfTrust,omitempty"`

			// The taskcluster proxy provides an easy and safe way to make authenticated
			// taskcluster requests within the scope(s) of a particular task. See
			// [the github project](https://github.com/taskcluster/taskcluster-proxy) for more information.
			//
			// Since: generic-worker 10.6.0
			TaskclusterProxy bool `json:"taskclusterProxy,omitempty"`
		} `json:"features,omitempty"`

		// Maximum time the task container can run in seconds.
		//
		// Since: generic-worker 0.0.1
		//
		// Mininum:    1
		// Maximum:    86400
		MaxRunTime int64 `json:"maxRunTime"`

		// Directories and/or files to be mounted.
		//
		// Since: generic-worker 5.4.0
		Mounts []Mount `json:"mounts,omitempty"`

		// A list of OS Groups that the task user should be a member of. Requires
		// scope `generic-worker:os-group:<os-group>` for each group listed.
		//
		// Since: generic-worker 6.0.0
		OSGroups []string `json:"osGroups,omitempty"`

		// Specifies an artifact name for publishing RDP connection information.
		//
		// Since this is potentially sensitive data, care should be taken to publish
		// to a suitably locked down path, such as
		// `login-identity/<login-identity>/rdpinfo.json` which is only readable for
		// the given login identity (for example
		// `login-identity/mozilla-ldap/pmoore@mozilla.com/rdpInfo.txt`). See the
		// [artifact namespace guide](https://docs.taskcluster.net/manual/design/namespaces#artifacts) for more information.
		//
		// Use of this feature requires scope
		// `generic-worker:allow-rdp:<provisionerId>/<workerType>` which must be
		// declared as a task scope.
		//
		// The RDP connection data is published during task startup so that a user
		// may interact with the running task.
		//
		// The task environment will be retained for 12 hours after the task
		// completes, to enable an interactive user to perform investigative tasks.
		// After these 12 hours, the worker will delete the task's Windows user
		// account, and then continue with other tasks.
		//
		// No guarantees are given about the resolution status of the interactive
		// task, since the task is inherently non-reproducible and no automation
		// should rely on this value.
		//
		// Since: generic-worker 10.5.0
		RdpInfo string `json:"rdpInfo,omitempty"`

		// URL of a service that can indicate tasks superseding this one; the current `taskId`
		// will be appended as a query argument `taskId`. The service should return an object with
		// a `supersedes` key containing a list of `taskId`s, including the supplied `taskId`. The
		// tasks should be ordered such that each task supersedes all tasks appearing later in the
		// list.
		//
		// See [superseding](https://docs.taskcluster.net/reference/platform/taskcluster-queue/docs/superseding) for more detail.
		//
		// Since: generic-worker 10.2.2
		SupersederURL string `json:"supersederUrl,omitempty"`
	}

	Mount json.RawMessage
)

type Queue tcqueue.Queue
type Provisioner tcawsprovisioner.AwsProvisioner

const waitTimeMinutes = 90

func FilenameEscape(raw string) (escaped string) {
	return strings.Replace(strings.Replace(raw, "*", "★", -1), "/", "⁄", -1)
}

func main() {
	taskIDs := map[string]string{}
	queue := tcqueue.NewFromEnv()
	taskGroupID := slugid.Nice()
	created := time.Now()
	for _, wt := range AllWorkerTypes() {
		fmt.Println(wt)
		x := strings.Split(wt, "/")
		provisionerID := x[0]
		workerType := x[1]
		taskID := slugid.Nice()
		taskIDs[wt] = taskID
		payload := GenericWorkerPayload{
			MaxRunTime: 3600,
			Command: []string{
				`echo`,
			},
		}
		payloadJSON := mustCompileToRawMessage(payload)
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
					`broken. The probe that generated this task simply reads the task log to determine which`,
					`worker implementation it is running. If this is not a generic-worker worker, this is`,
					`likely to result in a task exception, but that is expected and not a problem. Typically`,
					`the worker implementation can still be determined from the format of the task log file.`,
					``,
					`The resulting verdict is to be published [here](https://github.com/taskcluster/mozilla-history/blob/master/WorkerVersions/` + provisionerID + `%E2%81%84` + workerType + `).`,
				}, "\n"),
				Owner:  "pmoore@mozilla.com",
				Source: "https://github.com/taskcluster/mozilla-history/tree/master/audit-worker-versions",
			},
			Payload:       *payloadJSON,
			Priority:      "highest",
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
	}

	fmt.Println("Task Group ID: " + taskGroupID)
	fmt.Printf("Now sleeping for %v mins to wait for tasks to run...\n", waitTimeMinutes)

	// Leave plenty of time for tasks to run
	time.Sleep(waitTimeMinutes * time.Minute)

	continuationToken := ""
	for {
		tasks, err := queue.ListTaskGroup(taskGroupID, continuationToken, "")
		if err != nil {
			panic(err)
		}
		for _, task := range tasks.Tasks {
			workerPoolID, versionInfo := show(queue, task)
			filename := FilenameEscape(workerPoolID)
			err := ioutil.WriteFile(filename, append([]byte(versionInfo), '\n'), 0644)
			if err != nil {
				log.Fatalf("Error:\n%v", err)
			}
			name := workerPoolID + ":                                                                          "
			fmt.Print(name[:75])
			fmt.Println(versionInfo)
		}
		continuationToken = tasks.ContinuationToken
		if continuationToken == "" {
			break
		}
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

func NewProvisioner() *Provisioner {
	p := tcawsprovisioner.NewFromEnv()
	P := Provisioner(*p)
	return &P
}

func NewQueue() *Queue {
	q := tcqueue.NewFromEnv()
	Q := Queue(*q)
	return &Q
}

func (p *Provisioner) AllWorkerTypes() []string {
	prov := tcawsprovisioner.AwsProvisioner(*p)
	wt, err := prov.ListWorkerTypes()
	if err != nil {
		panic(err)
	}
	return []string(*wt)
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
			uniqueWorkerTypes[wt] = true
		}
	}

	// Now merge in known worker types according to AWS provisioner
	p := NewProvisioner()
	provisionerWorkerTypes := p.AllWorkerTypes()
	for _, wt := range provisionerWorkerTypes {
		uniqueWorkerTypes["aws-provisioner-v1/"+wt] = true
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

func show(queue *tcqueue.Queue, t tcqueue.TaskDefinitionAndStatus) (workerPoolID, versionInfo string) {
	workerPoolID = t.Task.ProvisionerID + "/" + t.Task.WorkerType
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
	case strings.Contains(logContent, "Task not successful due to following exception"):
		versionInfo = "generic-worker - unknown version"
	case strings.Contains(logContent, "Worker Node Type:"):
		versionInfo = "docker-worker - unknown version"
	case strings.Contains(logContent, `"release": "https://github.com/taskcluster/generic-worker/releases/tag/v`):
		reVersion := regexp.MustCompile(`"https://github.com/taskcluster/generic-worker/releases/tag/v([^"]*)"`)
		gwVersion := reVersion.FindStringSubmatch(logContent)
		versionInfo = "generic-worker " + gwVersion[1]
		reRevision := regexp.MustCompile(`"revision": "([0-9a-f]{40})"`)
		gwRevision := reRevision.FindStringSubmatch(logContent)
		if len(gwRevision) > 1 {
			versionInfo += " (revision " + gwRevision[1] + ")"
		}
	case strings.Contains(logContent, `not allowed at task.payload.features`):
		versionInfo = "taskcluster-worker - unknown version"
	case strings.Contains(logContent, `raise TaskVerificationError`):
		versionInfo = "scriptworker - unknown version"
	case strings.Contains(logContent, `KeyError: 'artifacts_deps'`):
		versionInfo = "some kind of scriptworker - unknown version"
	case artifactFound == "":
		versionInfo = "No artifacts found!"
	case artifactFound == "public/logs/chain_of_trust.log":
		versionInfo = "scriptworker chain of trust - unknown version"
	case strings.Contains(logContent, `os.environ.get('GITHUB_HEAD_REPO_URL', decision_json['payload']['env']['GITHUB_HEAD_REPO_URL'])`):
		versionInfo = "scriptworker - deepspeech - unknown version"
	default:
		versionInfo = "UNKNOWN"
		log.Printf("Cannot determine worker version - log: %v", logContent)
	}
	return
}

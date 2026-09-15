package main

import (
	"encoding/json"
	"errors"
	"strings"
	"testing"

	"github.com/taskcluster/taskcluster/v48/clients/client-go/tcqueue"
	"github.com/taskcluster/taskcluster/v48/clients/client-go/tcworkermanager"
)

type fakeWorkerManager struct {
	workerPool      *tcworkermanager.WorkerPoolFullDefinition
	workerPoolError error
	workerPages     []*tcworkermanager.ListWorkersResponse
	workerError     error
	listCalls       int
}

func (manager *fakeWorkerManager) WorkerPool(string) (*tcworkermanager.WorkerPoolFullDefinition, error) {
	return manager.workerPool, manager.workerPoolError
}

func (manager *fakeWorkerManager) ListWorkers(string, string, string, string, string, string) (*tcworkermanager.ListWorkersResponse, error) {
	if manager.workerError != nil {
		return nil, manager.workerError
	}
	page := manager.workerPages[manager.listCalls]
	manager.listCalls++
	return page, nil
}

func intPointer(value int) *int {
	return &value
}

func TestEnrichWorkerInfoWithConfiguredCapacity(t *testing.T) {
	tests := []struct {
		name   string
		config string
		slots  int
	}{
		{
			name:   "current nested capacity",
			config: `{"minCapacity":1,"maxCapacity":16,"launchConfigs":[{"workerManager":{"capacityPerInstance":8}}]}`,
			slots:  8,
		},
		{
			name:   "legacy direct capacity",
			config: `{"minCapacity":1,"maxCapacity":16,"launchConfigs":[{"capacityPerInstance":8}]}`,
			slots:  8,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			worker := WorkerInfo{}
			pool := tcworkermanager.WorkerPoolFullDefinition{
				Config: json.RawMessage(test.config),
			}

			enrichWorkerInfo(&worker, &pool)

			for name, result := range map[string]struct {
				got  *int
				want int
			}{
				"minimum capacity": {worker.ConfiguredMinCapacity, 1},
				"maximum capacity": {worker.ConfiguredMaxCapacity, 16},
				"slots per worker": {worker.CapacityPerWorker, test.slots},
				"minimum workers":  {worker.ConfiguredMinWorkers, 1},
				"maximum workers":  {worker.ConfiguredMaxWorkers, 2},
			} {
				if result.got == nil || *result.got != result.want {
					t.Errorf("%s = %v, want %d", name, result.got, result.want)
				}
			}

			encoded, err := json.Marshal(worker)
			if err != nil {
				t.Fatal(err)
			}
			for _, want := range []string{
				`"ConfiguredMinCapacity":1`,
				`"ConfiguredMaxCapacity":16`,
				`"CapacityPerWorker":8`,
				`"ConfiguredMinWorkers":1`,
				`"ConfiguredMaxWorkers":2`,
			} {
				if !strings.Contains(string(encoded), want) {
					t.Errorf("serialized worker does not contain %s", want)
				}
			}
			if strings.Contains(string(encoded), "TotalWorkers") || strings.Contains(string(encoded), "TotalCapacity") {
				t.Error("new worker snapshot contains legacy totals")
			}
		})
	}
}

func TestEnrichWorkerInfoDoesNotDeriveWorkersForMixedCapacities(t *testing.T) {
	worker := WorkerInfo{}
	pool := tcworkermanager.WorkerPoolFullDefinition{
		Config: json.RawMessage(`{
			"minCapacity": 0,
			"maxCapacity": 16,
			"launchConfigs": [
				{"workerManager":{"capacityPerInstance":2}},
				{"workerManager":{"capacityPerInstance":8}}
			]
		}`),
	}

	enrichWorkerInfo(&worker, &pool)

	if worker.ConfiguredMinCapacity == nil || worker.ConfiguredMaxCapacity == nil {
		t.Fatal("configured capacity bounds were not preserved")
	}
	if worker.CapacityPerWorker != nil || worker.ConfiguredMinWorkers != nil || worker.ConfiguredMaxWorkers != nil {
		t.Fatal("worker counts were derived for heterogeneous launch capacities")
	}
}

func TestSummarizeTaskGroup(t *testing.T) {
	tasks := []tcqueue.TaskDefinitionAndStatus{
		{Status: tcqueue.TaskStatusStructure{State: "completed"}},
		{Status: tcqueue.TaskStatusStructure{State: "failed"}},
		{Status: tcqueue.TaskStatusStructure{State: "exception"}},
		{Status: tcqueue.TaskStatusStructure{State: "running"}},
		{Status: tcqueue.TaskStatusStructure{State: "pending"}},
	}

	progress := summarizeTaskGroup(tasks)

	if progress.Total != 5 || progress.Terminal != 3 {
		t.Fatalf("progress = %#v, want 3 of 5 terminal", progress)
	}
	if progress.complete() {
		t.Fatal("incomplete task group was reported complete")
	}
}

func TestTaskGroupComplete(t *testing.T) {
	progress := summarizeTaskGroup([]tcqueue.TaskDefinitionAndStatus{
		{Status: tcqueue.TaskStatusStructure{State: "completed"}},
		{Status: tcqueue.TaskStatusStructure{State: "failed"}},
		{Status: tcqueue.TaskStatusStructure{State: "exception"}},
	})

	if !progress.complete() {
		t.Fatalf("progress = %#v, want complete", progress)
	}
	if (taskGroupProgress{}).complete() {
		t.Fatal("empty task group was reported complete")
	}
}

func TestTaskWasClaimed(t *testing.T) {
	tests := []struct {
		name string
		runs []tcqueue.RunInformation
		want bool
	}{
		{name: "no runs"},
		{
			name: "pending run",
			runs: []tcqueue.RunInformation{{State: "pending"}},
		},
		{
			name: "expired without claim",
			runs: []tcqueue.RunInformation{{State: "exception", ReasonResolved: "deadline-exceeded"}},
		},
		{
			name: "claimed malformed payload",
			runs: []tcqueue.RunInformation{{State: "exception", ReasonResolved: "malformed-payload", WorkerGroup: "us-west1-b", WorkerID: "worker-1"}},
			want: true,
		},
		{
			name: "deadline exceeded after claim",
			runs: []tcqueue.RunInformation{{State: "exception", ReasonResolved: "deadline-exceeded", WorkerID: "worker-1"}},
			want: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if got := taskWasClaimed(test.runs); got != test.want {
				t.Fatalf("taskWasClaimed() = %v, want %v", got, test.want)
			}
		})
	}
}

func TestShowClassifiesDeadlineExceededWithoutClaimAsUnresponsive(t *testing.T) {
	status := &tcqueue.TaskStatusResponse{Status: tcqueue.TaskStatusStructure{
		ProvisionerID: "example",
		WorkerType:    "pool",
		State:         "exception",
		Runs: []tcqueue.RunInformation{{
			State:          "exception",
			ReasonResolved: "deadline-exceeded",
		}},
	}}

	_, worker := show(nil, status)
	if !worker.isUnknown || worker.hasNoArtifacts {
		t.Fatalf("worker classification = %#v, want unresponsive", worker)
	}
	if got := worker.Details["error"]; got != "Version not determined; task was not claimed" {
		t.Fatalf("error = %q, want task-was-not-claimed result", got)
	}
}

func TestGetImagesetReadsAzureArmDeployment(t *testing.T) {
	pool := tcworkermanager.WorkerPoolFullDefinition{
		ProviderID: "azure2",
		Config: json.RawMessage(`{
			"launchConfigs": [
				{"armDeployment":{"parameters":{"imageId":{"value":"azure/image/10"}}}},
				{"armDeployment":{"parameters":{"imageId":{"value":"azure/image/2"}}}}
			]
		}`),
	}
	worker := WorkerInfo{}

	enrichWorkerInfo(&worker, &pool)

	if worker.Imageset != "azure/image/10,azure/image/2" {
		t.Fatalf("Imageset = %q, want Azure image IDs", worker.Imageset)
	}
	if worker.ImageStatus != imageStatusKnown || worker.ProviderID != "azure2" {
		t.Fatalf("worker metadata = %#v, want known azure2 image", worker)
	}
}

func TestLookupRecognizesPaginatedStandaloneWorkers(t *testing.T) {
	manager := &fakeWorkerManager{
		workerPoolError: errors.New("worker pool not found"),
		workerPages: []*tcworkermanager.ListWorkersResponse{
			{ContinuationToken: "next", Workers: []tcworkermanager.Worker{{ProviderID: "none", State: standaloneWorkerState}}},
			{Workers: []tcworkermanager.Worker{{ProviderID: "none", State: "running"}}},
		},
	}
	worker := WorkerInfo{WorkerPoolID: "releng-hardware/example"}

	lookupAndEnrichWorkerInfo(&worker, manager)

	if worker.ProviderID != standaloneProviderID || worker.ImageStatus != imageStatusNotApplicable {
		t.Fatalf("worker metadata = %#v, want standalone/not-applicable", worker)
	}
	if worker.WorkerManagerLookupError != "" {
		t.Fatalf("standalone worker retained lookup error %q", worker.WorkerManagerLookupError)
	}
	if manager.listCalls != 2 {
		t.Fatalf("ListWorkers calls = %d, want 2", manager.listCalls)
	}
}

func TestLookupRetainsUnexpectedFailure(t *testing.T) {
	manager := &fakeWorkerManager{
		workerPoolError: errors.New("worker pool unavailable"),
		workerError:     errors.New("worker list unavailable"),
	}
	worker := WorkerInfo{WorkerPoolID: "example/pool"}

	lookupAndEnrichWorkerInfo(&worker, manager)

	if worker.ImageStatus != imageStatusUnavailable {
		t.Fatalf("ImageStatus = %q, want unavailable", worker.ImageStatus)
	}
	for _, want := range []string{"worker pool unavailable", "worker list unavailable"} {
		if !strings.Contains(worker.WorkerManagerLookupError, want) {
			t.Errorf("lookup error %q does not contain %q", worker.WorkerManagerLookupError, want)
		}
	}
}

func TestLookupWithoutStandaloneSignalRemainsUnavailable(t *testing.T) {
	manager := &fakeWorkerManager{
		workerPoolError: errors.New("worker pool not found"),
		workerPages: []*tcworkermanager.ListWorkersResponse{
			{Workers: []tcworkermanager.Worker{{ProviderID: "none", State: "running"}}},
		},
	}
	worker := WorkerInfo{WorkerPoolID: "example/pool"}

	lookupAndEnrichWorkerInfo(&worker, manager)

	if worker.ProviderID == standaloneProviderID {
		t.Fatal("worker without standalone state was classified as standalone")
	}
	if worker.ImageStatus != imageStatusUnavailable || worker.WorkerManagerLookupError == "" {
		t.Fatalf("worker metadata = %#v, want unavailable with lookup error", worker)
	}
}

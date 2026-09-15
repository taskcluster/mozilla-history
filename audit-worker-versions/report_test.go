package main

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
	"time"
)

func TestReportRevision(t *testing.T) {
	dir := t.TempDir()
	git := func(args ...string) string {
		t.Helper()
		cmd := exec.Command("git", args...)
		cmd.Dir = dir
		out, err := cmd.CombinedOutput()
		if err != nil {
			t.Fatalf("git %v: %v: %s", args, err, out)
		}
		return strings.TrimSpace(string(out))
	}
	write := func(name, content string) {
		t.Helper()
		path := filepath.Join(dir, name)
		if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
			t.Fatal(err)
		}
		if err := os.WriteFile(path, []byte(content), 0644); err != nil {
			t.Fatal(err)
		}
	}
	if revision, link := reportRevision(dir); revision != "" || link != "" {
		t.Fatal("expected no revision outside a Git repository")
	}
	git("init")
	git("config", "user.name", "Report test")
	git("config", "user.email", "report@example.com")
	git("config", "commit.gpgsign", "false")
	git("remote", "add", "origin", "git@github.com:example/reports.git")
	write("source.go", "original")
	write("WorkerVersions/README.md", "original report")
	git("add", ".")
	git("commit", "-m", "initial")
	sha := git("rev-parse", "HEAD")
	check := func(suffix string) {
		t.Helper()
		revision, link := reportRevision(dir)
		if revision != sha[:9]+suffix || link != "https://github.com/example/reports/commit/"+sha {
			t.Fatalf("reportRevision() = (%q, %q)", revision, link)
		}
	}
	check("")
	write("WorkerVersions/README.md", "regenerated")
	write("WorkerVersions/workers.json", "{}")
	write(".beads/issues.jsonl", "{}")
	write("docs/history.json", "{}")
	check("")
	write("source.go", "modified")
	check("-dirty")
	git("add", "source.go")
	check("-dirty")
	git("restore", "--staged", "source.go")
	git("restore", "source.go")
	write("new-source.go", "new")
	check("-dirty")
}

func TestRenderRevisionBesideGenerationTime(t *testing.T) {
	for _, revision := range []string{"697b57411", "697b57411-dirty"} {
		for _, timings := range [][2]string{{"probe", "collected"}, {"", "collected"}, {"", ""}} {
			got := renderTemplate(reportData{
				ProbeStartedAt: timings[0], GeneratedAt: timings[1],
				ReportGeneratedAt: "2026-09-10 22:34 UTC",
				Revision:          revision, RevisionURL: "https://github.com/example/reports/commit/697b57411",
			})
			want := "Report generated: **2026-09-10 22:34 UTC** ([" + revision + "](https://github.com/example/reports/commit/697b57411))"
			if !strings.Contains(got, want) {
				t.Fatalf("missing generation provenance %q", want)
			}
		}
	}
}

func TestWorkerPoolURL(t *testing.T) {
	t.Setenv("TASKCLUSTER_ROOT_URL", "")
	worker := WorkerInfo{WorkerPoolID: "releng-hardware/gecko-t-win7-32-hw"}
	want := "https://firefox-ci-tc.services.mozilla.com/provisioners/releng-hardware/worker-types/gecko-t-win7-32-hw?sortBy=Last%20Active&sortDirection=desc"
	if got := worker.WorkerPoolURL(); got != want {
		t.Fatalf("WorkerPoolURL() = %q, want %q", got, want)
	}
}

func TestRenderReadmeTaskclusterRootURL(t *testing.T) {
	for _, root := range []string{"", "https://stage.taskcluster.nonprod.cloudops.mozgcp.net", "https://stage.taskcluster.nonprod.cloudops.mozgcp.net/"} {
		t.Run(root, func(t *testing.T) {
			t.Setenv("TASKCLUSTER_ROOT_URL", root)
			wantRoot := strings.TrimRight(root, "/")
			if wantRoot == "" {
				wantRoot = "https://firefox-ci-tc.services.mozilla.com"
			}
			got := renderReadmeAt(WorkerSnapshot{
				TaskGroupID:    "probe/group",
				ProbeStartedAt: time.Date(2026, time.September, 9, 7, 58, 29, 0, time.UTC),
				Workers:        []WorkerInfo{{WorkerPoolID: "example/pool name", Implementation: "generic-worker", Details: map[string]string{"revision": "1234567890"}}},
			}, time.Now())
			for _, path := range []string{
				"/provisioners/example/worker-types/pool%20name?sortBy=Last%20Active&sortDirection=desc",
				"/tasks/groups/probe%2Fgroup",
			} {
				if !strings.Contains(got, "("+wantRoot+path+")") {
					t.Errorf("report missing link %q", wantRoot+path)
				}
			}
		})
	}
}

func TestSortedVersionCountsUsesDescendingNaturalOrder(t *testing.T) {
	versions := map[string]int{
		"100.0.1": 1,
		"108.0.0": 1,
		"9.10.0":  1,
		"9.2.0":   1,
		"99.2.0":  1,
	}
	want := []string{"108.0.0", "100.0.1", "99.2.0", "9.10.0", "9.2.0"}

	got := sortedVersionCounts(versions)
	for i := range want {
		if got[i].Key != want[i] {
			t.Fatalf("version %d = %q, want %q", i, got[i].Key, want[i])
		}
	}
}

func TestRenderReadmeIncludesLinksAndSubheadings(t *testing.T) {
	t.Setenv("TASKCLUSTER_ROOT_URL", "")
	workers := []WorkerInfo{{
		WorkerPoolID:          "example/pool",
		Implementation:        "generic-worker",
		Version:               "1.2.3",
		Imageset:              "image-one",
		Details:               map[string]string{"revision": "1234567890"},
		ConfiguredMinCapacity: intPointer(1),
		ConfiguredMaxCapacity: intPointer(16),
		CapacityPerWorker:     intPointer(8),
		ConfiguredMinWorkers:  intPointer(1),
		ConfiguredMaxWorkers:  intPointer(2),
	}, {
		WorkerPoolID:   "example/other-pool",
		Implementation: "generic-worker",
		Version:        "2.0.0",
		Imageset:       "image-two",
		Details:        map[string]string{"revision": "1234567890"},
	}}

	got := renderReadmeAt(WorkerSnapshot{
		TaskGroupID:    "AnhEjBL2SYuUedNBvjgsWA",
		GeneratedAt:    time.Date(2026, time.September, 9, 15, 29, 53, 0, time.UTC),
		ProbeStartedAt: time.Date(2026, time.September, 9, 7, 58, 29, 0, time.UTC),
		Workers:        workers,
	}, time.Date(2026, time.September, 10, 1, 2, 3, 0, time.UTC))
	for _, want := range []string{
		"This report shows the latest detailed inventory of Firefox CI worker pools alongside historical trends from earlier snapshots.",
		"- **Implementation and version** are inferred from the failure log produced when each pool receives an intentionally malformed probe task.",
		"- **Image and capacity metadata** come from Worker Manager.",
		"- **Summary values** are counts of worker pools, not individual workers or tasks.",
		"Probe run started: **2026-09-09 07:58 UTC** ([Taskcluster task group](https://firefox-ci-tc.services.mozilla.com/tasks/groups/AnhEjBL2SYuUedNBvjgsWA)) · Results collected: **2026-09-09 15:29 UTC** · Report generated: **2026-09-10 01:02 UTC**",
		"Total generic-worker pools: `2`",
		"Total docker-worker pools: `0`",
		"Total scriptworker pools: `0`",
		"Total pools with unknown implementation: `0`",
		"Total unresponsive worker pools: `0`",
		"### Worker pools by version",
		"### Worker pools by image",
		"intentionally malformed probe task",
		"expected to fail with a malformed-payload exception",
		"live Worker Manager launch configuration",
		"### Worker pools",
		"Configured Workers | Configured Capacity | Slots per Worker",
		"pool's autoscaling range in concurrent task slots",
		"| 1–2 | 1–16 | 8 |",
		"[**example/pool**](https://firefox-ci-tc.services.mozilla.com/provisioners/example/worker-types/pool?sortBy=Last%20Active&sortDirection=desc)",
	} {
		if !strings.Contains(got, want) {
			t.Errorf("rendered README does not contain %q", want)
		}
	}
}

func TestRenderReadmeWithoutTaskGroup(t *testing.T) {
	got := renderReadmeAt(WorkerSnapshot{
		ProbeStartedAt: time.Date(2026, time.September, 9, 7, 58, 29, 0, time.UTC),
	}, time.Date(2026, time.September, 10, 1, 2, 3, 0, time.UTC))
	want := "Probe run started: **2026-09-09 07:58 UTC** · Report generated: **2026-09-10 01:02 UTC**"
	if !strings.Contains(got, want) || strings.Contains(got, "Taskcluster task group") {
		t.Fatal("report without a task group should retain its timing line without a link")
	}
}

func TestRenderReadmeDistinguishesImageStatuses(t *testing.T) {
	details := map[string]string{"revision": "1234567890"}
	workers := []WorkerInfo{
		{Implementation: "generic-worker", Imageset: "unknown", ImageStatus: imageStatusNotApplicable, Details: details},
		{Implementation: "generic-worker", Imageset: "unknown", ImageStatus: imageStatusUnavailable, Details: details},
		{Implementation: "generic-worker", Imageset: "unknown", ImageStatus: imageStatusNotDetermined, Details: details},
		{Implementation: "generic-worker", Imageset: "azure/image", ImageStatus: imageStatusKnown, Details: details},
	}

	got := renderReadme(WorkerSnapshot{Workers: workers})

	for _, want := range []string{
		"| Image | Worker pools |",
		"| Not applicable (standalone) | 1 |",
		"| Configuration unavailable | 1 |",
		"| Image not determined | 1 |",
		"| azure/image | 1 |",
	} {
		if !strings.Contains(got, want) {
			t.Errorf("rendered README does not contain %q", want)
		}
	}
}

func TestCompactAzureImageReferences(t *testing.T) {
	tests := map[string]string{
		"gallery with redundant image name": "/subscriptions/sub/resourceGroups/rg/providers/Microsoft.Compute/galleries/win2022/images/win2022/versions/1.0.0",
		"gallery with distinct image name":  "/subscriptions/sub/resourceGroups/rg/providers/Microsoft.Compute/galleries/gallery/images/image/versions/2.0",
		"managed image":                     "/subscriptions/sub/resourceGroups/rg/providers/Microsoft.Compute/images/imageset-worker-eastus",
		"non-Azure image":                   "projects/example/global/images/worker-image",
	}
	want := map[string]string{
		"gallery with redundant image name": "Azure gallery win2022@1.0.0",
		"gallery with distinct image name":  "Azure gallery gallery/image@2.0",
		"managed image":                     "Azure image imageset-worker-eastus",
		"non-Azure image":                   "projects/example/global/images/worker-image",
	}

	for name, input := range tests {
		t.Run(name, func(t *testing.T) {
			if got := compactAzureImageReference(input); got != want[name] {
				t.Fatalf("compactAzureImageReference() = %q, want %q", got, want[name])
			}
		})
	}
}

func TestCompactImageReferencesPreservesConfiguredSet(t *testing.T) {
	input := "/subscriptions/sub/resourceGroups/rg/providers/Microsoft.Compute/images/one,/subscriptions/sub/resourceGroups/rg/providers/Microsoft.Compute/images/two"
	want := "Azure image one, Azure image two"
	if got := compactImageReferences(input); got != want {
		t.Fatalf("compactImageReferences() = %q, want %q", got, want)
	}
}

func TestCompactImageReferencesSummarizesRegionalAzureSet(t *testing.T) {
	input := strings.Join([]string{
		"/subscriptions/sub/resourceGroups/rg/providers/Microsoft.Compute/images/imageset-abcdefghijklmnopqrst-westus2-fuzzing",
		"/subscriptions/sub/resourceGroups/rg/providers/Microsoft.Compute/images/imageset-zyxwvutsrqponmlkjihg-eastus-fuzzing",
	}, ",")
	want := "Azure image set fuzzing (eastus, westus2)"
	if got := compactImageReferences(input); got != want {
		t.Fatalf("compactImageReferences() = %q, want %q", got, want)
	}
}

func TestImageHoverTitleSeparatesImageSetWithLineBreaks(t *testing.T) {
	if got := imageHoverTitle("/azure/one,/azure/two"); got != "/azure/one&#10;/azure/two" {
		t.Fatalf("imageHoverTitle() = %q, want encoded line break", got)
	}
}

func TestRenderReadmeShowsFullAzureImageOnHover(t *testing.T) {
	fullImage := "/subscriptions/sub/resourceGroups/rg/providers/Microsoft.Compute/galleries/win2022/images/win2022/versions/1.0.0"
	workers := []WorkerInfo{
		{Implementation: "generic-worker", Imageset: fullImage, ImageStatus: imageStatusKnown, Details: map[string]string{"revision": "1234567890"}},
		{Implementation: "generic-worker", Imageset: "projects/example/global/images/linux", ImageStatus: imageStatusKnown, Details: map[string]string{"revision": "1234567890"}},
	}

	got := renderReadme(WorkerSnapshot{Workers: workers})

	want := `<abbr title="` + fullImage + `">Azure gallery win2022@1.0.0</abbr>`
	if !strings.Contains(got, want) {
		t.Fatalf("rendered README does not contain Azure hover label %q", want)
	}
	if strings.Contains(got, `<abbr title="projects/example`) {
		t.Fatal("non-Azure image unexpectedly received a hover label")
	}
}

func TestRenderReadmeExplainsIncompleteProbesInline(t *testing.T) {
	workers := []WorkerInfo{
		{WorkerPoolID: "example/no-artifact", Details: map[string]string{"error": "No artifacts found"}, hasNoArtifacts: true},
		{WorkerPoolID: "example/pending", Details: map[string]string{"error": "Version not determined; task not (yet) claimed"}, isUnknown: true},
	}

	got := renderReadme(WorkerSnapshot{Workers: workers})
	for _, want := range []string{
		"## Unknown implementation\n",
		"claimed and resolved the probe task",
		"did not publish a worker log artifact",
		"## Unresponsive\n",
		"did not claim the probe task",
	} {
		if !strings.Contains(got, want) {
			t.Errorf("rendered README does not contain %q", want)
		}
	}
	if strings.Contains(got, "[^1]") || strings.Contains(got, "[^2]") {
		t.Error("rendered README contains obsolete footnote markers")
	}
}

func TestReadSnapshotRestoresRenderingState(t *testing.T) {
	filename := filepath.Join(t.TempDir(), "workers.json")
	data := `[
		{"WorkerPoolID":"one/pool","Details":{"error":"No artifacts found"},"TotalWorkers":42,"TotalCapacity":84},
		{"WorkerPoolID":"two/pool","Details":{"error":"Version not determined; task not (yet) claimed"}},
		{"WorkerPoolID":"three/pool","Details":{"error":"Version not determined; task was not claimed"}}
	]`
	if err := os.WriteFile(filename, []byte(data), 0644); err != nil {
		t.Fatal(err)
	}

	snapshot, err := readSnapshot(filename)
	if err != nil {
		t.Fatal(err)
	}
	workers := snapshot.Workers
	if !workers[0].hasNoArtifacts {
		t.Error("no-artifacts state was not restored")
	}
	if workers[0].LegacyTotalWorkers == nil || *workers[0].LegacyTotalWorkers != 42 ||
		workers[0].LegacyTotalCapacity == nil || *workers[0].LegacyTotalCapacity != 84 {
		t.Error("legacy totals were not recognized")
	}
	if !workers[1].isUnknown || !workers[2].isUnknown {
		t.Error("old and current unclaimed states were not restored")
	}
}

func TestReadSnapshotReadsMetadata(t *testing.T) {
	filename := filepath.Join(t.TempDir(), "workers.json")
	data := `{
		"generatedAt":"2026-09-09T15:29:53Z",
		"probeStartedAt":"2026-09-09T07:58:29.422Z",
		"taskGroupId":"AnhEjBL2SYuUedNBvjgsWA",
		"workers":[{"WorkerPoolID":"one/pool","Details":{}}]
	}`
	if err := os.WriteFile(filename, []byte(data), 0644); err != nil {
		t.Fatal(err)
	}

	snapshot, err := readSnapshot(filename)
	if err != nil {
		t.Fatal(err)
	}
	if got := snapshot.GeneratedAt.Format(time.RFC3339); got != "2026-09-09T15:29:53Z" {
		t.Errorf("generated time = %q", got)
	}
	if got := snapshot.ProbeStartedAt.Format(time.RFC3339Nano); got != "2026-09-09T07:58:29.422Z" {
		t.Errorf("probe start time = %q", got)
	}
	if snapshot.TaskGroupID != "AnhEjBL2SYuUedNBvjgsWA" {
		t.Errorf("task group ID = %q", snapshot.TaskGroupID)
	}
	if len(snapshot.Workers) != 1 || snapshot.Workers[0].WorkerPoolID != "one/pool" {
		t.Errorf("workers = %#v", snapshot.Workers)
	}
}

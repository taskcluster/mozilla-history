package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html"
	"net/url"
	"os"
	"os/exec"
	"path/filepath"
	"sort"
	"strings"
	"text/template"
	"time"
)

type count struct {
	Key   string
	Value int
	Hover string
}

type reportSection struct {
	Title           string
	TotalLabel      string
	Description     string
	Count           int
	Versions        []count
	Images          []count
	Filtered        []WorkerInfo
	FullColumns     bool
	HasLegacyTotals bool
}

type reportData struct {
	GeneratedAt       string
	ProbeStartedAt    string
	TaskGroupURL      string
	ReportGeneratedAt string
	Revision          string
	RevisionURL       string
	Sections          [5]reportSection
}

// Keep offline report rendering usable without Taskcluster configuration.
func reportTaskclusterRootURL() string {
	root := os.Getenv("TASKCLUSTER_ROOT_URL")
	if root == "" {
		root = "https://firefox-ci-tc.services.mozilla.com"
	}
	return strings.TrimRight(root, "/")
}

func (w WorkerInfo) WorkerPoolURL() string {
	parts := strings.SplitN(w.WorkerPoolID, "/", 2)
	if len(parts) != 2 {
		return ""
	}
	return reportTaskclusterRootURL() + "/provisioners/" +
		url.PathEscape(parts[0]) + "/worker-types/" + url.PathEscape(parts[1]) +
		"?sortBy=Last%20Active&sortDirection=desc"
}

func formatConfiguredRange(minimum, maximum *int) string {
	if minimum == nil || maximum == nil {
		return "—"
	}
	if *minimum == *maximum {
		return fmt.Sprintf("%d", *minimum)
	}
	return fmt.Sprintf("%d–%d", *minimum, *maximum)
}

func (w WorkerInfo) ConfiguredWorkerRange() string {
	return formatConfiguredRange(w.ConfiguredMinWorkers, w.ConfiguredMaxWorkers)
}

func (w WorkerInfo) ConfiguredCapacityRange() string {
	return formatConfiguredRange(w.ConfiguredMinCapacity, w.ConfiguredMaxCapacity)
}

func (w WorkerInfo) CapacityPerWorkerValue() string {
	if w.CapacityPerWorker == nil {
		return "—"
	}
	return fmt.Sprintf("%d", *w.CapacityPerWorker)
}

const readmeTpl = `
{{- define "row" -}}
## {{ .Title }}
{{ if .Description }}
{{ .Description }}
{{ end }}

Total {{ .TotalLabel }}: ` + "`" + `{{ .Count }}` + "`" + `
{{ if gt (len .Versions) 1 }}
### Worker pools by version

_Source: version information parsed from the log artifact produced when each worker claims an intentionally malformed probe task. The task is expected to fail with a malformed-payload exception; known worker implementations and versions are identified from their distinct log output._

| Version | Worker pools |
| :--- | ---: |
{{ range .Versions -}}
| {{ .Key }} | {{ .Value }} |
{{ end }}
{{- end }}
{{ if gt (len .Images) 1 }}
### Worker pools by image

_Source: image references in each pool's live Worker Manager launch configuration at report time. A value may represent multiple configured images. Standalone pools do not have a Worker Manager-managed image._

| Image | Worker pools |
| :--- | ---: |
{{ range .Images -}}
| {{ if .Hover }}<abbr title="{{ .Hover }}">{{ .Key }}</abbr>{{ else }}{{ .Key }}{{ end }} | {{ .Value }} |
{{ end }}
{{- end }}
{{if .Count }}
### Worker pools

_Configured capacity is the pool's autoscaling range in concurrent task slots. The configured worker range is derived from those bounds and the slots per worker, rounding up. An em dash means the configuration is unavailable or a worker count cannot be derived._
{{ if .HasLegacyTotals }}
_Configured values were not collected in this snapshot. Legacy totals included stopped worker records and are intentionally not displayed._
{{ end }}

| Worker Pool | Implementation | Version {{ if .FullColumns }}| Engine | Revision | OS | Arch | GO {{ end }}| Configured Workers | Configured Capacity | Slots per Worker |
| --- | --- | --- {{ if .FullColumns }}| --- | --- | --- | --- | --- {{ end }}| ---: | ---: | ---: |
{{ range .Filtered -}}
| [**{{ .WorkerPoolID }}**]({{ .WorkerPoolURL }}) | {{ .Implementation }} | {{ or .Version .Details.error }} {{ if $.FullColumns }}| {{ or .Details.engine "-" }} | {{ or (slice .Details.revision 0 10) "-" }} | {{ or .Details.os "-" }} | {{ or .Details.arch "-" }} | {{ or .Details.go "-" }} {{ end }}| {{ .ConfiguredWorkerRange }} | {{ .ConfiguredCapacityRange }} | {{ .CapacityPerWorkerValue }} |
{{end}}
{{- end -}}
{{end}}

# Worker Pool Versions

This report shows the latest detailed inventory of Firefox CI worker pools alongside historical trends from earlier snapshots.

- **Implementation and version** are inferred from the failure log produced when each pool receives an intentionally malformed probe task.
- **Image and capacity metadata** come from Worker Manager.
- **Summary values** are counts of worker pools, not individual workers or tasks.

{{ define "generation" }}Report generated: **{{ .ReportGeneratedAt }}**{{ if .Revision }} ({{ if .RevisionURL }}[{{ .Revision }}]({{ .RevisionURL }}){{ else }}{{ .Revision }}{{ end }}){{ end }}{{ end }}
{{ if .ProbeStartedAt }}Probe run started: **{{ .ProbeStartedAt }}**{{ if .TaskGroupURL }} ([Taskcluster task group]({{ .TaskGroupURL }})){{ end }}{{ if .GeneratedAt }} · Results collected: **{{ .GeneratedAt }}**{{ end }}{{ if .ReportGeneratedAt }} · {{ template "generation" . }}{{ end }}
{{ else if .GeneratedAt }}Results collected: **{{ .GeneratedAt }}**{{ if .ReportGeneratedAt }} · {{ template "generation" . }}{{ end }}
{{ else if .ReportGeneratedAt }}{{ template "generation" . }}
{{ end }}

{{ range .Sections }}
{{ template "row" . }}
{{ end }}
`

func renderTemplate(data interface{}) string {
	t := template.Must(template.New("").Parse(readmeTpl))
	var content bytes.Buffer
	if err := t.Execute(&content, data); err != nil {
		panic(err)
	}
	return strings.TrimSpace(content.String()) + "\n"
}

func sortedCounts(values map[string]int) []count {
	counts := make([]count, 0, len(values))
	for key, value := range values {
		counts = append(counts, count{Key: key, Value: value})
	}
	sort.Slice(counts, func(i, j int) bool {
		return strings.Compare(counts[i].Key, counts[j].Key) < 0
	})
	return counts
}

func naturalLess(left, right string) bool {
	for leftIndex, rightIndex := 0, 0; leftIndex < len(left) && rightIndex < len(right); {
		leftDigit := left[leftIndex] >= '0' && left[leftIndex] <= '9'
		rightDigit := right[rightIndex] >= '0' && right[rightIndex] <= '9'
		if leftDigit && rightDigit {
			leftEnd, rightEnd := leftIndex, rightIndex
			for leftEnd < len(left) && left[leftEnd] >= '0' && left[leftEnd] <= '9' {
				leftEnd++
			}
			for rightEnd < len(right) && right[rightEnd] >= '0' && right[rightEnd] <= '9' {
				rightEnd++
			}

			leftSignificant, rightSignificant := leftIndex, rightIndex
			for leftSignificant < leftEnd-1 && left[leftSignificant] == '0' {
				leftSignificant++
			}
			for rightSignificant < rightEnd-1 && right[rightSignificant] == '0' {
				rightSignificant++
			}

			leftLength := leftEnd - leftSignificant
			rightLength := rightEnd - rightSignificant
			if leftLength != rightLength {
				return leftLength < rightLength
			}
			if leftNumber, rightNumber := left[leftSignificant:leftEnd], right[rightSignificant:rightEnd]; leftNumber != rightNumber {
				return leftNumber < rightNumber
			}
			if leftRunLength, rightRunLength := leftEnd-leftIndex, rightEnd-rightIndex; leftRunLength != rightRunLength {
				return leftRunLength < rightRunLength
			}

			leftIndex, rightIndex = leftEnd, rightEnd
			continue
		}

		if left[leftIndex] != right[rightIndex] {
			return left[leftIndex] < right[rightIndex]
		}
		leftIndex++
		rightIndex++
	}

	return len(left) < len(right)
}

func sortedVersionCounts(values map[string]int) []count {
	counts := sortedCounts(values)
	sort.Slice(counts, func(i, j int) bool {
		return naturalLess(counts[j].Key, counts[i].Key)
	})
	return counts
}

func compactAzureImageReference(reference string) string {
	parts := strings.Split(strings.Trim(reference, "/"), "/")
	if len(parts) < 6 || !strings.EqualFold(parts[0], "subscriptions") {
		return reference
	}

	providerIsAzureCompute := false
	for i := 0; i+1 < len(parts); i++ {
		if strings.EqualFold(parts[i], "providers") && strings.EqualFold(parts[i+1], "Microsoft.Compute") {
			providerIsAzureCompute = true
			break
		}
	}
	if !providerIsAzureCompute {
		return reference
	}

	for i := 0; i+5 < len(parts); i++ {
		if strings.EqualFold(parts[i], "galleries") &&
			strings.EqualFold(parts[i+2], "images") &&
			strings.EqualFold(parts[i+4], "versions") {
			gallery := parts[i+1]
			image := parts[i+3]
			version := parts[i+5]
			if gallery == image {
				return fmt.Sprintf("Azure gallery %s@%s", image, version)
			}
			return fmt.Sprintf("Azure gallery %s/%s@%s", gallery, image, version)
		}
	}

	for i := 0; i+1 < len(parts); i++ {
		if strings.EqualFold(parts[i], "images") {
			return "Azure image " + parts[i+1]
		}
	}
	return reference
}

func compactAzureImageSet(references []string) (string, bool) {
	regions := make([]string, 0, len(references))
	family := ""
	for _, reference := range references {
		parts := strings.Split(strings.Trim(reference, "/"), "/")
		if len(parts) < 2 || !strings.EqualFold(parts[0], "subscriptions") ||
			!strings.EqualFold(parts[len(parts)-2], "images") {
			return "", false
		}
		nameParts := strings.Split(parts[len(parts)-1], "-")
		if len(nameParts) < 4 || nameParts[0] != "imageset" {
			return "", false
		}
		currentFamily := strings.Join(nameParts[3:], "-")
		if family != "" && currentFamily != family {
			return "", false
		}
		family = currentFamily
		regions = append(regions, nameParts[2])
	}
	if len(regions) < 2 {
		return "", false
	}
	sort.Strings(regions)
	return fmt.Sprintf("Azure image set %s (%s)", family, strings.Join(regions, ", ")), true
}

func compactImageReferences(imageset string) string {
	references := strings.Split(imageset, ",")
	for i := range references {
		references[i] = strings.TrimSpace(references[i])
	}
	if compact, ok := compactAzureImageSet(references); ok {
		return compact
	}
	for i, reference := range references {
		references[i] = compactAzureImageReference(reference)
	}
	return strings.Join(references, ", ")
}

func imageHoverTitle(imageset string) string {
	references := strings.Split(imageset, ",")
	for i, reference := range references {
		references[i] = html.EscapeString(strings.TrimSpace(reference))
	}
	return strings.Join(references, "&#10;")
}

func imageCountLabel(worker WorkerInfo) string {
	switch worker.ImageStatus {
	case imageStatusNotApplicable:
		return "Not applicable (standalone)"
	case imageStatusUnavailable:
		return "Configuration unavailable"
	case imageStatusNotDetermined:
		return "Image not determined"
	}
	if worker.WorkerManagerLookupError != "" {
		return "Configuration unavailable"
	}
	if worker.Imageset == "" || worker.Imageset == "unknown" {
		return "Image not determined"
	}
	return compactImageReferences(worker.Imageset)
}

func generateReadmeSection(title, description string, workers []WorkerInfo, filter func(WorkerInfo) bool) reportSection {
	filtered := make([]WorkerInfo, 0)
	versions := make(map[string]int)
	imagesets := make(map[string]int)
	imageHovers := make(map[string]string)
	hasLegacyTotals := false

	for _, worker := range workers {
		if filter(worker) {
			filtered = append(filtered, worker)
			versions[worker.Version]++
			imageLabel := imageCountLabel(worker)
			imagesets[imageLabel]++
			if worker.ImageStatus == imageStatusKnown && imageLabel != worker.Imageset {
				imageHovers[imageLabel] = worker.Imageset
			}
			hasLegacyTotals = hasLegacyTotals || worker.LegacyTotalWorkers != nil || worker.LegacyTotalCapacity != nil
		}
	}

	sort.Slice(filtered, func(i, j int) bool {
		return strings.Compare(filtered[i].WorkerPoolID, filtered[j].WorkerPoolID) < 0
	})

	images := sortedCounts(imagesets)
	for i := range images {
		if hover := imageHovers[images[i].Key]; hover != "" {
			images[i].Key = html.EscapeString(images[i].Key)
			images[i].Hover = imageHoverTitle(hover)
		}
	}

	return reportSection{
		Title:           title,
		TotalLabel:      sectionTotalLabel(title),
		Description:     description,
		Count:           len(filtered),
		Versions:        sortedVersionCounts(versions),
		Images:          images,
		Filtered:        filtered,
		FullColumns:     title == "Generic Worker",
		HasLegacyTotals: hasLegacyTotals,
	}
}

func sectionTotalLabel(title string) string {
	switch title {
	case "Generic Worker":
		return "generic-worker pools"
	case "Docker Worker":
		return "docker-worker pools"
	case "Script Worker":
		return "scriptworker pools"
	case "Unknown implementation":
		return "pools with unknown implementation"
	case "Unresponsive":
		return "unresponsive worker pools"
	default:
		return "worker pools"
	}
}

func writeReadme(snapshot WorkerSnapshot) {
	filename := filepath.Join(outputDir, "README.md")
	WriteFile(filename, []byte(renderReadme(snapshot)))
}

func renderReadme(snapshot WorkerSnapshot) string {
	return renderReadmeAt(snapshot, time.Now())
}

// Read provenance before writing generated artifacts. Use origin so fork-only
// commits link to the repository that actually contains them.
func reportRevision(dir string) (revision, revisionURL string) {
	git := func(args ...string) (string, error) {
		cmd := exec.Command("git", args...)
		cmd.Dir = dir
		out, err := cmd.Output()
		return strings.TrimSpace(string(out)), err
	}
	sha, err := git("rev-parse", "--verify", "HEAD")
	if err != nil || len(sha) < 9 {
		return "", ""
	}
	status, err := git("status", "--porcelain", "--untracked-files=all", "--", ":/", ":(top,exclude)WorkerVersions/**", ":(top,exclude).beads/**", ":(top,exclude)docs/history.json")
	if err != nil {
		return "", "" // Do not claim a clean revision if status is unavailable.
	}
	revision = sha[:9]
	if status != "" {
		revision += "-dirty"
	}
	remote, err := git("remote", "get-url", "origin")
	if err == nil {
		remote = strings.TrimSuffix(remote, ".git")
		remote = strings.Replace(remote, "git@github.com:", "https://github.com/", 1)
		remote = strings.Replace(remote, "ssh://git@github.com/", "https://github.com/", 1)
		if strings.HasPrefix(remote, "https://github.com/") {
			revisionURL = strings.TrimRight(remote, "/") + "/commit/" + sha
		}
	}
	return revision, revisionURL
}

func renderReadmeAt(snapshot WorkerSnapshot, reportGeneratedAt time.Time) string {
	workers := snapshot.Workers
	sections := [5]reportSection{
		generateReadmeSection("Generic Worker", "These pools use generic-worker to execute tasks. Recent versions also support docker-worker payloads and Docker container execution through d2g. [Documentation](https://docs.taskcluster.net/docs/reference/workers/generic-worker) · [Source](https://github.com/taskcluster/taskcluster/tree/main/workers/generic-worker).", workers, func(w WorkerInfo) bool { return w.Implementation == "generic-worker" }),
		generateReadmeSection("Docker Worker", "These pools use docker-worker to run tasks inside Docker containers. Docker-worker is fully deprecated, and its source repository has been removed. See the [worker migration timeline](https://taskcluster.github.io/mozilla-history/migration.html) for the transition to generic-worker.", workers, func(w WorkerInfo) bool { return w.Implementation == "docker-worker" }),
		generateReadmeSection("Script Worker", "", workers, func(w WorkerInfo) bool { return strings.Contains(w.Implementation, "Scriptworker") }),
		generateReadmeSection("Unknown implementation", "These pools claimed and resolved the probe task, but did not publish a worker log artifact. Their worker implementation and version could therefore not be identified.", workers, func(w WorkerInfo) bool { return w.hasNoArtifacts }),
		generateReadmeSection("Unresponsive", "These pools did not claim the probe task, so their worker implementation and version could not be determined.", workers, func(w WorkerInfo) bool { return w.isUnknown }),
	}

	const timestampFormat = "2006-01-02 15:04 UTC"
	data := reportData{Sections: sections}
	data.Revision, data.RevisionURL = reportRevision(".")
	if snapshot.TaskGroupID != "" {
		data.TaskGroupURL = reportTaskclusterRootURL() + "/tasks/groups/" + url.PathEscape(snapshot.TaskGroupID)
	}
	if !snapshot.GeneratedAt.IsZero() {
		data.GeneratedAt = snapshot.GeneratedAt.UTC().Format(timestampFormat)
	}
	if !snapshot.ProbeStartedAt.IsZero() {
		data.ProbeStartedAt = snapshot.ProbeStartedAt.UTC().Format(timestampFormat)
	}
	if !reportGeneratedAt.IsZero() {
		data.ReportGeneratedAt = reportGeneratedAt.UTC().Format(timestampFormat)
	}

	return renderTemplate(data)
}

func readSnapshot(filename string) (WorkerSnapshot, error) {
	contents, err := os.ReadFile(filename)
	if err != nil {
		return WorkerSnapshot{}, err
	}

	var snapshot WorkerSnapshot
	if strings.HasPrefix(strings.TrimSpace(string(contents)), "[") {
		if err := json.Unmarshal(contents, &snapshot.Workers); err != nil {
			return WorkerSnapshot{}, err
		}
	} else if err := json.Unmarshal(contents, &snapshot); err != nil {
		return WorkerSnapshot{}, err
	}

	// These flags are internal rendering state and are not serialized in the
	// snapshot. Restore them from the persisted error value for offline renders.
	for i := range snapshot.Workers {
		switch snapshot.Workers[i].Details["error"] {
		case "No artifacts found":
			snapshot.Workers[i].hasNoArtifacts = true
		case "Version not determined; task not (yet) claimed", "Version not determined; task was not claimed":
			snapshot.Workers[i].isUnknown = true
		}
	}

	return snapshot, nil
}

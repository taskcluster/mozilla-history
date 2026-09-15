# Optional Quick previews

These helpers deploy this checkout to [Mozilla Quick](https://quick.mozilla.cloud/)
for contributor previews. Visit Quick's home page for access to the platform.
They are separate from upstream's production reporting and GitHub Pages
publishing workflows.

## Scripts

- `deploy-quick` publishes a personal preview of the report and the checkout's
  saved worker data to `<username>-mozilla-history.quick.mozilla.cloud`.
- `deploy-quick-staging` publishes the same content to a separate
  `<username>-mozilla-history-staging.quick.mozilla.cloud` site for trying changes.
- `deploy-common.sh` contains the shared naming, confirmation, packaging, and
  cleanup logic. It is sourced by the two helpers, not run directly.
- `quick-head.html` supplies the Quick SDK and navigation styling injected into
  each packaged HTML page, keeping Quick integration out of upstream page sources.

Use these previews to share and review report changes behind Mozilla SSO.
Neither helper collects new data, creates probes, or publishes GitHub Pages.

## Setup

Install and configure Mozilla's Quick CLI for your account before using these
helpers. The `quick` command must be available on your `PATH`.

Site names are explicit and independent of the checkout directory or Quick's
inferred name:

- `<username>-mozilla-history`
- `<username>-mozilla-history-staging`

The username defaults to `id -un` (your local OS account). Set
`QUICK_PREVIEW_USER` to override it, for example
`QUICK_PREVIEW_USER=aerickson scripts/quick/deploy-quick`. Names are lowercased,
non-alphanumeric runs become hyphens, and names over 63 characters are rejected.

## Usage

From the repository root:

```sh
scripts/quick/deploy-quick
scripts/quick/deploy-quick-staging
```

Both helpers resolve the repository root from their own location, so they work
when invoked by absolute path from another directory. They package the static
files from `docs/` at the site root alongside `WorkerVersions/README.md` and
`WorkerVersions/workers.json` in a temporary directory. The report opens directly
at the preview URL and loads the packaged snapshot; no root redirect is needed.
Only packaged pages get the Quick SDK script and navigation styling. The temporary package
is removed when deployment finishes or fails, and the checkout is not modified.

In an interactive terminal, each helper displays the target URL and asks for
approval, defaulting to No. For explicitly approved automation, use:

```sh
scripts/quick/deploy-quick --confirm
scripts/quick/deploy-quick-staging --confirm
```

Without a terminal, the helpers exit unless `--confirm` is supplied. Use
`--help` for a usage summary. These commands deploy the files already in the
checkout; they do not refresh worker data or create Taskcluster probes.

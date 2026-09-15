#!/usr/bin/env bash

# Shared by the preview entry points; this file is not a deployment command.
deploy_preview() (
  set -euo pipefail
  suffix=$1
  confirmed=$2
  repo_dir=$(cd -- "$(dirname -- "${BASH_SOURCE[0]}")/../.." && pwd)
  username=${QUICK_PREVIEW_USER:-$(id -un)}
  username=$(printf '%s' "$username" | LC_ALL=C tr '[:upper:]' '[:lower:]' | LC_ALL=C sed -E 's/[^a-z0-9]+/-/g; s/^-+//; s/-+$//')
  if [[ -z "$username" ]]; then
    echo "Set QUICK_PREVIEW_USER to a username containing letters or digits." >&2
    exit 1
  fi
  site_name="${username}-mozilla-history${suffix}"
  if [[ ${#site_name} -gt 63 ]]; then
    echo "Preview site name exceeds 63 characters; use a shorter QUICK_PREVIEW_USER." >&2
    exit 1
  fi
  deployment_url="https://${site_name}.quick.mozilla.cloud/"
  if [[ "$confirmed" != true ]]; then
    if [[ ! -t 0 ]]; then
      echo "Deployment to ${deployment_url} requires interactive approval; pass --confirm to approve explicitly." >&2
      exit 1
    fi
    read -r -p "Deploy this checkout to ${deployment_url}? [y/N] " answer
    case "$answer" in
      y|Y|yes|YES) ;;
      *) echo "Deployment cancelled."; exit 0 ;;
    esac
  fi

  site_dir=$(mktemp -d "${TMPDIR:-/tmp}/mozilla-history-preview.XXXXXX")
  trap 'rm -rf -- "$site_dir"' EXIT
  cp "$repo_dir"/docs/*.html "$repo_dir"/docs/*.js "$repo_dir/docs/history.json" "$site_dir/"
  mkdir "$site_dir/WorkerVersions"
  cp "$repo_dir/WorkerVersions/README.md" "$repo_dir/WorkerVersions/workers.json" "$site_dir/WorkerVersions/"
  # Add Quick integration only to packaged pages, leaving source HTML unchanged.
  for page in "$repo_dir"/docs/*.html; do
    awk '
      NR == FNR { fragment = fragment $0 "\n"; next }
      /<\/head>/ { printf "%s", fragment }
      { print }
    ' "$repo_dir/scripts/quick/quick-head.html" "$page" > "$site_dir/$(basename "$page")"
  done
  echo "Deploying preview: $deployment_url"
  quick deploy "$site_dir" "$site_name"
)

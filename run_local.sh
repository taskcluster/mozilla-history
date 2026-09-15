#!/bin/sh

set -eu

port="${1:-8000}"
repo_dir=$(CDPATH= cd -- "$(dirname -- "$0")" && pwd)

case "$port" in
  *[!0-9]*|'')
    echo "Usage: $0 [port]" >&2
    exit 2
    ;;
esac

echo "Serving the local report at http://localhost:${port}/docs/index-local.html"
exec python3 -m http.server "$port" --directory "$repo_dir"

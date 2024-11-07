#!/bin/bash

export TASKCLUSTER_ROOT_URL=https://firefox-ci-tc.services.mozilla.com/
export REPORT_SCHEDULER_ID="smoketest"
export REPORT_PREFIX=https://github.com/taskcluster/mozilla-history/blob/master/WorkerVersions/

# needs to be set in the environment:
# TASKCLUSTER_CLIENT_ID
# TASKCLUSTER_ACCESS_TOKEN

set -euo pipefail

cd "$(dirname "${0}")"

# build necessary tools with
# go build -buildvcs=false -o mozilla-history
# go build -buildvcs=false -o audit-worker-versions ./audit-worker-versions
# or
# go build ./...

git pull

echo "generating audit report"
./mozilla-history

git add .
git commit -am "Audit report: $(date +'%Y-%m-%d %H:%M:%S')"

echo "running worker versions report"
./audit.sh

git add .
git commit -am "Worker versions report: $(date +'%Y-%m-%d %H:%M:%S')"

# update history
./build-docs-history.sh
git add .
git commit -v -a --no-edit --amend

git push

#!/bin/bash -e

# This script collects all revisions of the "WorkerVersions/workers.json" file
# to extract statistics by date of the commit
# Result is placed in docs/history.json file that be displayed on the github page

# node script assigned to variable that will extract stats from json files
{ nodescript=$(cat) ; } <<'HEREDOC'
	const path = require('path')
	if (process.argv.length < 3) {
		console.error('Expected some arguments')
		process.exit(1)
	}
	const runStats = (data) => {
		const implementations = {}
		const versions = {}
		const imagesets = {}
		const inc = (dict, key) => dict[key] = (dict[key] || 0) + 1
		data.forEach(worker => {
			inc(implementations, worker.Implementation)
			inc(versions, worker.Version)
			inc(imagesets, worker.Imageset)
		})
		return { implementations, versions, imagesets }
	}

	const out = {}
	for (i = 2; i < process.argv.length; i++) {
		const m = process.argv[i].match(/([0-9-]+).json/)
		const file = require(path.join(process.argv[1], process.argv[i]))
		out[m[1]] = runStats(file)
	}
	console.log(JSON.stringify(out))
HEREDOC

# build a revision history for public website

PUBLIC_DIR="./docs"
WORKERS_FILE="WorkerVersions/workers.json"
HISTORICAL_DATA="./${PUBLIC_DIR}/history.json"

mkdir -p "${PUBLIC_DIR}/data/"

# grab all versions by date
for rev in $(git rev-list master "${WORKERS_FILE}");
do
	revdate=$(git show --no-patch --no-notes --date=short --pretty='%cd' "$rev")
	echo "Fetching ${revdate} version ${rev}"
	git show "${rev}:${WORKERS_FILE}" > "${PUBLIC_DIR}/data/${revdate}.json"
done

node -e "${nodescript}" "${PWD}" ${PUBLIC_DIR}/data/*.json > $HISTORICAL_DATA
rm -rf "${PUBLIC_DIR}/data/"

echo "Done"

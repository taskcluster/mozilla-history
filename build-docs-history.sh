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
		const workers = Array.isArray(data) ? data : data.workers
		const implementations = {}
		const versions = {}
		const imagesets = {}
		// versions keyed by implementation, because the version numbers of
		// docker-worker and generic-worker share a numeric range and are
		// indistinguishable once merged into a single bucket
		const versionsByImplementation = {}
		const inc = (dict, key) => dict[key] = (dict[key] || 0) + 1
		workers.forEach(worker => {
			inc(implementations, worker.Implementation)
			inc(versions, worker.Version)
			inc(imagesets, worker.Imageset)
			if (!versionsByImplementation[worker.Implementation]) {
				versionsByImplementation[worker.Implementation] = {}
			}
			inc(versionsByImplementation[worker.Implementation], worker.Version)
		})
		return { implementations, versions, imagesets, versionsByImplementation }
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
while IFS= read -r rev;
do
	if ! git cat-file -e "${rev}:${WORKERS_FILE}" 2>/dev/null; then
		echo "Skipping ${rev}: ${WORKERS_FILE} does not exist"
		continue
	fi

	revdate=$(git show --no-patch --no-notes --date=short --pretty='%cd' "$rev")
	echo "Fetching ${revdate} version ${rev}"
	git show "${rev}:${WORKERS_FILE}" > "${PUBLIC_DIR}/data/${revdate}.json"
done < <(git rev-list --reverse HEAD -- "${WORKERS_FILE}")

node -e "${nodescript}" "${PWD}" ${PUBLIC_DIR}/data/*.json > $HISTORICAL_DATA
rm -rf "${PUBLIC_DIR}/data/"

echo "Done"

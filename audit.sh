#!/bin/bash -e

if [[ -z "$TASKCLUSTER_ROOT_URL" ]]; then
	echo "TASKCLUSTER_ROOT_URL is not set"
	exit 1
fi

echo "Generating tasks for audit"
while IFS= read -r line;
do
	[[ "$line" =~ "Task Group ID:" ]] && echo "matched" &&
	export TASK_GROUP_ID=$(echo "$line" | sed -n "s/^Task Group ID: \(.*\)$/\1/p");
	echo "$line";
done < <(audit-worker-versions);

echo "Generated tasks for TASK_GROUP_ID = $TASK_GROUP_ID"

echo "Sleep 90m"
sleep 5400
echo "Generating reports"

cd WorkerVersions
audit-worker-versions "$TASK_GROUP_ID"

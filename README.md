# mozilla-history
Taskcluster history from https://firefox-ci-tc.services.mozilla.com deployment
for:

* [Clients](/Clients)
* [Hooks](/Hooks)
* [Roles](/Roles)
* [Worker Pools](/WorkerPools)

For https://community-tc.services.mozilla.com history, please see
[community-history](https://github.com/taskcluster/community-history).

Please note, until 9 November 2019 this repository stored the history for the
https://taskcluster.net deployment. On 9 November 2019 the
https://taskcluster.net deployment was decommissioned and the
https://firefox-ci-tc.services.mozilla.com was instated.

## Entity filenames

The files are named after the entities themselves, except for the following
character conversions:

  1. `*` -> `★`

This conversion avoids illegal filenames.

  2. `/` -> `⁄`

Rather than creating nested subdirectories, this conversion avoids directory
names colliding with entity filenames.

## Installing

```
go get github.com/taskcluster/mozilla-history@v1.0.0
```

## Running

```
unset TASKCLUSTER_CLIENT_ID TASKCLUSTER_ACCESS_TOKEN TASKCLUSTER_CERTIFICATE
export TASKCLUSTER_ROOT_URL='https://firefox-ci-tc.services.mozilla.com'
mozilla-history
```

This will populate subdirectories `Clients`, `Hooks`, `Roles` and `WorkerPools`
of the current directory.

## Automating the Process

You can automate this reporting process by setting up a cron job to execute `run-report.sh` at regular intervals.

### Prerequisites
- Valid Taskcluster credentials must be set in the environment variables:
  - `TASKCLUSTER_CLIENT_ID`
  - `TASKCLUSTER_ACCESS_TOKEN`

### How it works
1. `run-report.sh` executes `audit.sh`
2. `audit.sh` schedules tasks for each worker pool to extract worker implementation details from logs
3. Results are stored in the `WorkerPools` directory
4. `mozilla-history` stores Taskcluster configurations in their respective directories:
   - Hooks definitions in `Hooks/`
   - Roles definitions in `Roles/`
   - WorkerPool definitions in `WorkerPool/`
5. `build-docs-history.sh` generates a static page containing the version history

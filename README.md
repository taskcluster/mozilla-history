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

## Entity update cadence

The `mozilla-history` command is run every 5 mins from a raspberry pi in
@petemoore's home network with the results committed to this repository and
pushed to github.

## Installing

```
go get github.com/taskcluster/mozilla-history
```

## Running

```
unset TASKCLUSTER_CLIENT_ID TASKCLUSTER_ACCESS_TOKEN TASKCLUSTER_CERTIFICATE
export TASKCLUSTER_ROOT_URL='https://firefox-ci-tc.services.mozilla.com'
mozilla-history
```

This will populate subdirectories `Clients`, `Hooks`, `Roles` and `WorkerPools`
of the current directory.

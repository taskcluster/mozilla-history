# mozilla-history
Taskcluster history from https://taskcluster.net deployment for:

* [AWS Worker Types](/AWSWorkerTypes)
* [Clients](/Clients)
* [Hooks](/Hooks)
* [Roles](/Roles)
* [Secrets (hashed)](/Secrets)

## Entity filenames

The files are named after the entities themselves, except for the following character conversions:

  1. `*` -> `★`

This conversion avoids illegal filenames.

  2. `/` -> `⁄`

Rather than creating nested subdirectories, this conversion avoids conflicting directory names and entity filenames.

## Entity update cadence

The `mozilla-history` command is run every 5 mins from a crontab entry on
@petemoore's iMac, with the results committed to this repository and pushed to
github. The update takes around 10 seconds.

## Installing

```
go get github.com/taskcluster/mozilla-history
```

## Running

_NOTE_: You need a taskcluster client that has `secrets:get:*`.

```
export TASKCLUSTER_CLIENT_ID='...'
export TASKCLUSTER_ACCESS_TOKEN='...'
export TASKCLUSTER_ROOT_URL='https://taskcluster.net'
mozilla-history
```

This will populate subdirectories `AWSWorkerTypes`, `Clients`, `Hooks`, `Roles` and `Secrets` of the current directory.

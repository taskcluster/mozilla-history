{
  "created": "2016-10-12T16:46:55.062Z",
  "description": "Custom role for Geoff Brown, created by pmoore on 12 October 2016. Initial purpose is to provide scopes that allows Geoff to help out with Windows tests troubleshooting in gecko.",
  "expandedScopes": [
    "assume:mozilla-user:gbrown@mozilla.com",
    "assume:project:releng:gecko-win-test-troubleshooter",
    "assume:project:taskcluster:tutorial",
    "generic-worker:allow-rdp:releng-hardware/gecko-t-win*",
    "queue:route:notify.email.*",
    "secrets:get:repo:github.com/mozilla-releng/OpenCloudConfig:gecko-t-win*"
  ],
  "lastModified": "2016-10-26T14:07:45.481Z",
  "roleId": "mozilla-user:gbrown@mozilla.com",
  "scopes": [
    "assume:project:releng:gecko-win-test-troubleshooter"
  ]
}

{
  "created": "2017-11-29T21:57:41.014Z",
  "description": "",
  "expandedScopes": [
    "assume:mozilla-user:garndt@mozilla.com",
    "assume:project:taskcluster:tutorial",
    "project:relops-hardware-controller:*",
    "queue:route:notify.email.*"
  ],
  "lastModified": "2017-12-07T17:17:52.065Z",
  "roleId": "mozilla-user:garndt@mozilla.com",
  "scopes": [
    "project:relops-hardware-controller:*"
  ]
}

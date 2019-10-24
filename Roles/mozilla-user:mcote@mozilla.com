{
  "created": "2016-12-19T16:19:05.593Z",
  "description": "",
  "expandedScopes": [
    "assume:mozilla-user:mcote@mozilla.com",
    "assume:project:shipit:user",
    "assume:project:taskcluster:tutorial",
    "project:shipit:analysis/use",
    "project:shipit:bugzilla",
    "project:shipit:user",
    "queue:route:notify.email.*"
  ],
  "lastModified": "2016-12-19T16:19:05.593Z",
  "roleId": "mozilla-user:mcote@mozilla.com",
  "scopes": [
    "assume:project:shipit:user"
  ]
}

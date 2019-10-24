{
  "created": "2016-12-19T16:22:28.349Z",
  "description": "",
  "expandedScopes": [
    "assume:mozilla-user:ehenry@mozilla.com",
    "assume:project:shipit:admin",
    "assume:project:shipit:user",
    "assume:project:taskcluster:tutorial",
    "project:shipit:admin",
    "project:shipit:analysis/use",
    "project:shipit:bugzilla",
    "project:shipit:user",
    "queue:route:notify.email.*"
  ],
  "lastModified": "2016-12-19T16:22:28.349Z",
  "roleId": "mozilla-user:ehenry@mozilla.com",
  "scopes": [
    "assume:project:shipit:admin"
  ]
}

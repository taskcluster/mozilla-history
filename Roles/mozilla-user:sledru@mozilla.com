{
  "created": "2016-12-19T16:21:13.415Z",
  "description": "",
  "expandedScopes": [
    "assume:mozilla-user:sledru@mozilla.com",
    "assume:project:shipit:admin",
    "assume:project:shipit:user",
    "assume:project:taskcluster:tutorial",
    "project:shipit:admin",
    "project:shipit:analysis/use",
    "project:shipit:bugzilla",
    "project:shipit:user",
    "queue:route:notify.email.*"
  ],
  "lastModified": "2016-12-19T16:21:13.415Z",
  "roleId": "mozilla-user:sledru@mozilla.com",
  "scopes": [
    "assume:project:shipit:admin"
  ]
}

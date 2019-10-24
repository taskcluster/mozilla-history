{
  "created": "2016-12-19T16:21:58.839Z",
  "description": "",
  "expandedScopes": [
    "assume:mozilla-user:rkothari@mozilla.com",
    "assume:project:shipit:admin",
    "assume:project:shipit:user",
    "assume:project:taskcluster:tutorial",
    "project:shipit:admin",
    "project:shipit:analysis/use",
    "project:shipit:bugzilla",
    "project:shipit:user",
    "queue:route:notify.email.*"
  ],
  "lastModified": "2016-12-19T16:21:58.839Z",
  "roleId": "mozilla-user:rkothari@mozilla.com",
  "scopes": [
    "assume:project:shipit:admin"
  ]
}

{
  "created": "2017-12-27T14:55:25.294Z",
  "description": "",
  "expandedScopes": [
    "assume:mozilla-user:apereira@mozilla.com",
    "assume:project:taskcluster:tc-hooks-tests",
    "assume:project:taskcluster:tutorial",
    "queue:route:notify.email.*"
  ],
  "lastModified": "2017-12-27T15:25:23.058Z",
  "roleId": "mozilla-user:apereira@mozilla.com",
  "scopes": [
    "assume:project:taskcluster:tc-hooks-tests"
  ]
}

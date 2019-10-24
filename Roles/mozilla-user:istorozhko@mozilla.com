{
  "created": "2016-12-19T21:46:28.652Z",
  "description": "",
  "expandedScopes": [
    "assume:mozilla-user:istorozhko@mozilla.com",
    "assume:project:taskcluster:github",
    "assume:project:taskcluster:tutorial",
    "queue:route:notify.email.*"
  ],
  "lastModified": "2016-12-19T21:46:28.652Z",
  "roleId": "mozilla-user:istorozhko@mozilla.com",
  "scopes": [
    "assume:project:taskcluster:github"
  ]
}

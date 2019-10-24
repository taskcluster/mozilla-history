{
  "created": "2017-10-12T23:47:07.446Z",
  "description": "",
  "expandedScopes": [
    "assume:mozilla-user:nalexander@mozilla.com",
    "assume:project:taskcluster:tutorial",
    "hooks:trigger-hook:project-releng/nightly-fennec/oak",
    "queue:route:notify.email.*"
  ],
  "lastModified": "2017-10-12T23:47:07.446Z",
  "roleId": "mozilla-user:nalexander@mozilla.com",
  "scopes": [
    "hooks:trigger-hook:project-releng/nightly-fennec/oak"
  ]
}

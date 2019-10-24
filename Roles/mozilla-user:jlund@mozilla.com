{
  "created": "2016-07-19T02:54:08.412Z",
  "description": "Authority to modify index to backfill, sometime necessary for making releases...",
  "expandedScopes": [
    "assume:mozilla-user:jlund@mozilla.com",
    "assume:project:taskcluster:tutorial",
    "index:insert-task:gecko.v2.*",
    "queue:route:notify.email.*"
  ],
  "lastModified": "2016-07-19T02:54:08.412Z",
  "roleId": "mozilla-user:jlund@mozilla.com",
  "scopes": [
    "index:insert-task:gecko.v2.*"
  ]
}

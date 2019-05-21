{
  "created": "2016-07-19T02:55:47.393Z",
  "description": "Authority to modify index to backfill, sometime necessary for making releases...",
  "expandedScopes": [
    "assume:mozilla-user:raliiev@mozilla.com",
    "assume:project:taskcluster:tutorial",
    "index:insert-task:gecko.v2.*",
    "queue:create-task:aws-provisioner-v1/tutorial",
    "queue:route:index.garbage.*",
    "queue:route:notify.email.*",
    "secrets:get:garbage/*",
    "secrets:set:garbage/*"
  ],
  "lastModified": "2016-07-19T02:55:47.393Z",
  "roleId": "mozilla-user:raliiev@mozilla.com",
  "scopes": [
    "index:insert-task:gecko.v2.*"
  ]
}

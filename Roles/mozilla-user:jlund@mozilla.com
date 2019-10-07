{
  "created": "2016-07-19T02:54:08.412Z",
  "description": "Authority to modify index to backfill, sometime necessary for making releases...",
  "expandedScopes": [
    "assume:mozilla-user:jlund@mozilla.com",
    "assume:project:taskcluster:tutorial",
    "index:insert-task:gecko.v2.*",
    "queue:create-task:aws-provisioner-v1/tutorial",
    "queue:create-task:low:built-in/*",
    "queue:route:index.garbage.*",
    "queue:route:notify.email.*",
    "queue:scheduler-id:tutorial",
    "secrets:get:garbage/*",
    "secrets:set:garbage/*"
  ],
  "lastModified": "2016-07-19T02:54:08.412Z",
  "roleId": "mozilla-user:jlund@mozilla.com",
  "scopes": [
    "index:insert-task:gecko.v2.*"
  ]
}

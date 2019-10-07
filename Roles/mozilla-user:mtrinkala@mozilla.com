{
  "created": "2016-02-01T17:57:36.780Z",
  "description": "Let trink create tasks for trink-worker",
  "expandedScopes": [
    "assume:mozilla-user:mtrinkala@mozilla.com",
    "assume:project:taskcluster:tutorial",
    "queue:cancel-task:-/*",
    "queue:create-task:aws-provisioner-v1/tutorial",
    "queue:create-task:low:built-in/*",
    "queue:define-task:trink-provisioned/trink-*",
    "queue:rerun-task:-/*",
    "queue:route:index.garbage.*",
    "queue:route:notify.email.*",
    "queue:schedule-task:trink-provisioned/trink-*",
    "queue:scheduler-id:tutorial",
    "queue:task-group-id:-/*",
    "secrets:get:garbage/*",
    "secrets:set:garbage/*"
  ],
  "lastModified": "2016-02-01T17:57:36.780Z",
  "roleId": "mozilla-user:mtrinkala@mozilla.com",
  "scopes": [
    "queue:cancel-task:-/*",
    "queue:define-task:trink-provisioned/trink-*",
    "queue:rerun-task:-/*",
    "queue:schedule-task:trink-provisioned/trink-*",
    "queue:task-group-id:-/*"
  ]
}

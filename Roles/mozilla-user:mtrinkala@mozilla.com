{
  "created": "2016-02-01T17:57:36.780Z",
  "description": "Let trink create tasks for trink-worker",
  "expandedScopes": [
    "assume:mozilla-user:mtrinkala@mozilla.com",
    "assume:project:taskcluster:tutorial",
    "queue:cancel-task:-/*",
    "queue:define-task:trink-provisioned/trink-*",
    "queue:rerun-task:-/*",
    "queue:route:notify.email.*",
    "queue:schedule-task:trink-provisioned/trink-*",
    "queue:task-group-id:-/*"
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

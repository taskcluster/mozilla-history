{
  "created": "2017-12-12T21:28:24.469Z",
  "description": "",
  "expandedScopes": [
    "assume:mozilla-user:gguthe@mozilla.com",
    "assume:project:taskcluster:tutorial",
    "queue:create-task:aws-provisioner-v1/tutorial",
    "queue:create-task:highest:test-dummy-provisioner/dummy-test-worker-type",
    "queue:create-task:low:built-in/*",
    "queue:declare-provisioner:test-dummy-provisioner#actions",
    "queue:route:index.garbage.*",
    "queue:route:notify.email.*",
    "queue:scheduler-id:-",
    "queue:scheduler-id:tutorial",
    "secrets:get:garbage/*",
    "secrets:set:garbage/*"
  ],
  "lastModified": "2017-12-12T21:53:52.884Z",
  "roleId": "mozilla-user:gguthe@mozilla.com",
  "scopes": [
    "queue:create-task:highest:test-dummy-provisioner/dummy-test-worker-type",
    "queue:declare-provisioner:test-dummy-provisioner#actions",
    "queue:scheduler-id:-"
  ]
}

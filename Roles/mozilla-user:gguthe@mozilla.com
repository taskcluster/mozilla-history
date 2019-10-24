{
  "created": "2017-12-12T21:28:24.469Z",
  "description": "",
  "expandedScopes": [
    "assume:mozilla-user:gguthe@mozilla.com",
    "assume:project:taskcluster:tutorial",
    "queue:create-task:highest:test-dummy-provisioner/dummy-test-worker-type",
    "queue:declare-provisioner:test-dummy-provisioner#actions",
    "queue:route:notify.email.*",
    "queue:scheduler-id:-"
  ],
  "lastModified": "2017-12-12T21:53:52.884Z",
  "roleId": "mozilla-user:gguthe@mozilla.com",
  "scopes": [
    "queue:create-task:highest:test-dummy-provisioner/dummy-test-worker-type",
    "queue:declare-provisioner:test-dummy-provisioner#actions",
    "queue:scheduler-id:-"
  ]
}

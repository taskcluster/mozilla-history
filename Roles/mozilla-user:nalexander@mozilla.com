{
  "created": "2017-10-12T23:47:07.446Z",
  "description": "",
  "expandedScopes": [
    "assume:mozilla-user:nalexander@mozilla.com",
    "assume:project:taskcluster:tutorial",
    "hooks:trigger-hook:project-releng/nightly-fennec/oak",
    "queue:create-task:aws-provisioner-v1/tutorial",
    "queue:create-task:low:built-in/*",
    "queue:route:index.garbage.*",
    "queue:route:notify.email.*",
    "queue:scheduler-id:tutorial",
    "secrets:get:garbage/*",
    "secrets:set:garbage/*"
  ],
  "lastModified": "2017-10-12T23:47:07.446Z",
  "roleId": "mozilla-user:nalexander@mozilla.com",
  "scopes": [
    "hooks:trigger-hook:project-releng/nightly-fennec/oak"
  ]
}

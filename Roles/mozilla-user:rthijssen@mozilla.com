{
  "created": "2016-05-12T16:45:46.109Z",
  "description": "pmoore added scopes so rthijssen can see and edit the windows passwords of the TC windows workers.",
  "expandedScopes": [
    "assume:mozilla-user:rthijssen@mozilla.com",
    "assume:project:taskcluster:tutorial",
    "auth:update-role:worker-type:aws-provisioner-v1/win*",
    "queue:create-task:aws-provisioner-v1/tutorial",
    "queue:create-task:low:built-in/*",
    "queue:route:index.garbage.*",
    "queue:route:notify.email.*",
    "queue:scheduler-id:tutorial",
    "secrets:get:garbage/*",
    "secrets:get:project/taskcluster/aws-provisioner-v1/worker-types/ssh-keys/*",
    "secrets:set:garbage/*",
    "secrets:set:project/taskcluster/aws-provisioner-v1/worker-types/ssh-keys/*"
  ],
  "lastModified": "2016-05-13T09:38:05.164Z",
  "roleId": "mozilla-user:rthijssen@mozilla.com",
  "scopes": [
    "auth:update-role:worker-type:aws-provisioner-v1/win*",
    "secrets:get:project/taskcluster/aws-provisioner-v1/worker-types/ssh-keys/*",
    "secrets:set:project/taskcluster/aws-provisioner-v1/worker-types/ssh-keys/*"
  ]
}

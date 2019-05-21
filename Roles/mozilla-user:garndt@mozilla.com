{
  "created": "2017-11-29T21:57:41.014Z",
  "description": "",
  "expandedScopes": [
    "assume:mozilla-user:garndt@mozilla.com",
    "assume:project:taskcluster:tutorial",
    "project:relops-hardware-controller:*",
    "queue:create-task:aws-provisioner-v1/tutorial",
    "queue:route:index.garbage.*",
    "queue:route:notify.email.*",
    "secrets:get:garbage/*",
    "secrets:set:garbage/*"
  ],
  "lastModified": "2017-12-07T17:17:52.065Z",
  "roleId": "mozilla-user:garndt@mozilla.com",
  "scopes": [
    "project:relops-hardware-controller:*"
  ]
}

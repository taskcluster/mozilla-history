{
  "created": "2016-08-25T10:32:00.421Z",
  "description": "For bug 1296995, a playground is required for Bryce to do his work in. Created by pmoore.",
  "expandedScopes": [
    "assume:mozilla-user:bvandyk@mozilla.com",
    "assume:project:taskcluster:tutorial",
    "queue:create-task:aws-provisioner-v1/tutorial",
    "queue:create-task:aws-provisioner-v1/win2012r2-mediatests",
    "queue:route:index.garbage.*",
    "queue:route:notify.email.*",
    "secrets:get:garbage/*",
    "secrets:set:garbage/*"
  ],
  "lastModified": "2016-08-25T10:32:00.421Z",
  "roleId": "mozilla-user:bvandyk@mozilla.com",
  "scopes": [
    "queue:create-task:aws-provisioner-v1/win2012r2-mediatests"
  ]
}

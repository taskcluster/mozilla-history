{
  "created": "2016-05-31T15:55:47.367Z",
  "description": "Tim can directly create tasks for nss-win2012r2 worker type.",
  "expandedScopes": [
    "assume:mozilla-user:ttaubert@mozilla.com",
    "assume:project:taskcluster:tutorial",
    "queue:create-task:aws-provisioner-v1/nss-win2012r2",
    "queue:create-task:aws-provisioner-v1/tutorial",
    "queue:route:index.garbage.*",
    "queue:route:notify.email.*",
    "secrets:get:garbage/*",
    "secrets:set:garbage/*"
  ],
  "lastModified": "2016-08-01T14:27:58.588Z",
  "roleId": "mozilla-user:ttaubert@mozilla.com",
  "scopes": [
    "queue:create-task:aws-provisioner-v1/nss-win2012r2"
  ]
}

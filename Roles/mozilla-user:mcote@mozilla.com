{
  "created": "2016-12-19T16:19:05.593Z",
  "description": "",
  "expandedScopes": [
    "assume:mozilla-user:mcote@mozilla.com",
    "assume:project:shipit:user",
    "assume:project:taskcluster:tutorial",
    "project:shipit:analysis/use",
    "project:shipit:bugzilla",
    "project:shipit:user",
    "queue:create-task:aws-provisioner-v1/tutorial",
    "queue:route:index.garbage.*",
    "queue:route:notify.email.*",
    "secrets:get:garbage/*",
    "secrets:set:garbage/*"
  ],
  "lastModified": "2016-12-19T16:19:05.593Z",
  "roleId": "mozilla-user:mcote@mozilla.com",
  "scopes": [
    "assume:project:shipit:user"
  ]
}

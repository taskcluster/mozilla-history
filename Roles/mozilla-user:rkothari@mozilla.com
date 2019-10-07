{
  "created": "2016-12-19T16:21:58.839Z",
  "description": "",
  "expandedScopes": [
    "assume:mozilla-user:rkothari@mozilla.com",
    "assume:project:shipit:admin",
    "assume:project:shipit:user",
    "assume:project:taskcluster:tutorial",
    "project:shipit:admin",
    "project:shipit:analysis/use",
    "project:shipit:bugzilla",
    "project:shipit:user",
    "queue:create-task:aws-provisioner-v1/tutorial",
    "queue:create-task:low:built-in/*",
    "queue:route:index.garbage.*",
    "queue:route:notify.email.*",
    "queue:scheduler-id:tutorial",
    "secrets:get:garbage/*",
    "secrets:set:garbage/*"
  ],
  "lastModified": "2016-12-19T16:21:58.839Z",
  "roleId": "mozilla-user:rkothari@mozilla.com",
  "scopes": [
    "assume:project:shipit:admin"
  ]
}

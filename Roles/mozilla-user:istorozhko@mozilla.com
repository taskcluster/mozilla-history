{
  "created": "2016-12-19T21:46:28.652Z",
  "description": "",
  "expandedScopes": [
    "assume:mozilla-user:istorozhko@mozilla.com",
    "assume:project:taskcluster:github",
    "assume:project:taskcluster:tutorial",
    "queue:create-task:aws-provisioner-v1/tutorial",
    "queue:create-task:low:built-in/*",
    "queue:route:index.garbage.*",
    "queue:route:notify.email.*",
    "queue:scheduler-id:tutorial",
    "secrets:get:garbage/*",
    "secrets:set:garbage/*"
  ],
  "lastModified": "2016-12-19T21:46:28.652Z",
  "roleId": "mozilla-user:istorozhko@mozilla.com",
  "scopes": [
    "assume:project:taskcluster:github"
  ]
}

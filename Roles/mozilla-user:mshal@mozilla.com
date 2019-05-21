{
  "created": "2016-07-15T19:50:30.403Z",
  "description": "mshal gets gecko.v2 scopes for index as he often backports fixes...",
  "expandedScopes": [
    "assume:mozilla-user:mshal@mozilla.com",
    "assume:project:taskcluster:tutorial",
    "index:insert-task:gecko.v2.*",
    "queue:create-task:aws-provisioner-v1/tutorial",
    "queue:route:index.garbage.*",
    "queue:route:notify.email.*",
    "secrets:get:garbage/*",
    "secrets:set:garbage/*"
  ],
  "lastModified": "2016-07-15T19:50:30.403Z",
  "roleId": "mozilla-user:mshal@mozilla.com",
  "scopes": [
    "index:insert-task:gecko.v2.*"
  ]
}

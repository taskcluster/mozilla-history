{
  "created": "2016-04-21T17:23:07.553Z",
  "description": "",
  "expandedScopes": [
    "assume:mozilla-user:bwong@mozilla.com",
    "assume:project:taskcluster:mozilla-github-repository",
    "assume:project:taskcluster:tutorial",
    "assume:repo:github.com/mozilla/balrog:*",
    "notify:email:*",
    "notify:irc-channel:*",
    "notify:irc-user:*",
    "queue:create-task:aws-provisioner-v1/github-worker",
    "queue:create-task:aws-provisioner-v1/github-worker-windows",
    "queue:create-task:aws-provisioner-v1/tutorial",
    "queue:create-task:aws-provisioner-v1/win2012r2",
    "queue:route:checks",
    "queue:route:garbage.*",
    "queue:route:index.garbage.*",
    "queue:route:index.project.balrog.*",
    "queue:route:notify.*",
    "queue:route:statuses",
    "queue:scheduler-id:taskcluster-github",
    "secrets:get:garbage/*",
    "secrets:get:repo:github.com/mozilla/balrog:coveralls",
    "secrets:get:repo:github.com/mozilla/balrog:dockerhub",
    "secrets:set:garbage/*"
  ],
  "lastModified": "2016-04-21T17:23:07.553Z",
  "roleId": "mozilla-user:bwong@mozilla.com",
  "scopes": [
    "assume:repo:github.com/mozilla/balrog:*"
  ]
}

{
  "created": "2016-06-29T19:55:36.713Z",
  "description": "",
  "expandedScopes": [
    "assume:mozilla-user:mbanner@mozilla.com",
    "assume:project:taskcluster:mozilla-github-repository",
    "assume:project:taskcluster:tutorial",
    "assume:repo:github.com/mozilla/loop:*",
    "notify:email:*",
    "notify:irc-channel:*",
    "notify:irc-user:*",
    "queue:cancel-task:*",
    "queue:create-task:aws-provisioner-v1/github-worker",
    "queue:create-task:aws-provisioner-v1/github-worker-windows",
    "queue:create-task:aws-provisioner-v1/win2012r2",
    "queue:route:checks",
    "queue:route:garbage.*",
    "queue:route:index.garbage.*",
    "queue:route:notify.email.*",
    "queue:route:notify.irc-channel.*",
    "queue:route:notify.irc-user.*",
    "queue:route:statuses",
    "queue:schedule-task:task-graph-scheduler/*",
    "queue:scheduler-id:taskcluster-github",
    "secrets:get:garbage/*",
    "secrets:get:repo:github.com/mozilla/example-addon-repo:*",
    "secrets:set:garbage/*",
    "secrets:set:repo:github.com/mozilla/example-addon-repo:*"
  ],
  "lastModified": "2016-09-19T13:48:36.050Z",
  "roleId": "mozilla-user:mbanner@mozilla.com",
  "scopes": [
    "assume:repo:github.com/mozilla/loop:*",
    "queue:cancel-task:*",
    "queue:schedule-task:task-graph-scheduler/*",
    "secrets:get:repo:github.com/mozilla/example-addon-repo:*",
    "secrets:set:repo:github.com/mozilla/example-addon-repo:*"
  ]
}

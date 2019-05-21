{
  "created": "2017-12-27T14:55:25.294Z",
  "description": "",
  "expandedScopes": [
    "assume:hook-id:tc-hooks-tests/tc-test-hook",
    "assume:mozilla-user:apereira@mozilla.com",
    "assume:project:taskcluster:tc-hooks-tests",
    "assume:project:taskcluster:tutorial",
    "auth:azure-table-access:jungle/*",
    "auth:azure-table:read-write:jungle/*",
    "project:taskcluster:tests:tc-hooks:scope/required/for/task/1",
    "queue:create-task:aws-provisioner-v1/tutorial",
    "queue:create-task:no-provisioner/test-worker",
    "queue:route:index.garbage.*",
    "queue:route:notify.email.*",
    "secrets:get:garbage/*",
    "secrets:set:garbage/*"
  ],
  "lastModified": "2017-12-27T15:25:23.058Z",
  "roleId": "mozilla-user:apereira@mozilla.com",
  "scopes": [
    "assume:project:taskcluster:tc-hooks-tests"
  ]
}

{
  "created": "2016-03-09T18:58:24.111Z",
  "description": "",
  "expandedScopes": [
    "assume:mozilla-user:jclaudius@mozilla.com",
    "assume:project:taskcluster:tutorial",
    "queue:create-task:aws-provisioner-v1/b2gtest",
    "queue:define-task:aws-provisioner-v1/b2gtest",
    "queue:route:notify.email.*"
  ],
  "lastModified": "2016-06-15T16:20:20.976Z",
  "roleId": "mozilla-user:jclaudius@mozilla.com",
  "scopes": [
    "queue:create-task:aws-provisioner-v1/b2gtest",
    "queue:define-task:aws-provisioner-v1/b2gtest"
  ]
}

{
  "created": "2017-08-25T17:49:13.349Z",
  "description": "https://bugzil.la/1393578",
  "expandedScopes": [
    "assume:mozilla-user:cdawson@mozilla.com",
    "assume:project:releng:treestatus/sheriff",
    "assume:project:taskcluster:tutorial",
    "project:releng:treestatus/recent_changes/revert",
    "project:releng:treestatus/trees/update",
    "queue:create-task:aws-provisioner-v1/tutorial",
    "queue:route:index.garbage.*",
    "queue:route:notify.email.*",
    "secrets:get:garbage/*",
    "secrets:set:garbage/*"
  ],
  "lastModified": "2017-08-25T17:49:13.349Z",
  "roleId": "mozilla-user:cdawson@mozilla.com",
  "scopes": [
    "assume:project:releng:treestatus/sheriff"
  ]
}

{
  "created": "2017-08-25T17:49:13.349Z",
  "description": "https://bugzil.la/1393578",
  "expandedScopes": [
    "assume:mozilla-user:cdawson@mozilla.com",
    "assume:project:releng:treestatus/sheriff",
    "assume:project:taskcluster:tutorial",
    "project:releng:treestatus/recent_changes/revert",
    "project:releng:treestatus/trees/update",
    "queue:route:notify.email.*"
  ],
  "lastModified": "2017-08-25T17:49:13.349Z",
  "roleId": "mozilla-user:cdawson@mozilla.com",
  "scopes": [
    "assume:project:releng:treestatus/sheriff"
  ]
}

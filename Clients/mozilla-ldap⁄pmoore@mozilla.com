{
  "clientId": "mozilla-ldap/pmoore@mozilla.com",
  "created": "2017-01-04T15:05:51.769Z",
  "deleteOnExpiration": false,
  "description": "Hack so I can use treeherder again",
  "disabled": false,
  "expandedScopes": [
    "assume:mozilla-user:pmoore@mozilla.com",
    "assume:project:taskcluster:tutorial",
    "queue:create-task:aws-provisioner-v1/tutorial",
    "queue:create-task:low:built-in/*",
    "queue:route:index.garbage.*",
    "queue:route:notify.email.*",
    "queue:scheduler-id:tutorial",
    "secrets:get:garbage/*",
    "secrets:set:garbage/*"
  ],
  "expires": "3017-01-31T23:00:00.000Z",
  "lastDateUsed": "1977-08-19T16:00:00.000Z",
  "lastModified": "2017-01-04T15:05:51.769Z",
  "lastRotated": "2017-01-04T15:05:51.769Z",
  "scopes": [
    "assume:mozilla-user:pmoore@mozilla.com"
  ]
}

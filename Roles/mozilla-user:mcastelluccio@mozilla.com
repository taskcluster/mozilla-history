{
  "created": "2017-09-05T14:20:03.314Z",
  "description": "Role created for bug 1396809",
  "expandedScopes": [
    "assume:mozilla-user:mcastelluccio@mozilla.com",
    "assume:project:taskcluster:tutorial",
    "github:create-status:marco-c/*",
    "queue:create-task:aws-provisioner-v1/tutorial",
    "queue:route:index.garbage.*",
    "queue:route:notify.email.*",
    "secrets:get:garbage/*",
    "secrets:get:repo:github.com/mozilla-releng/services:branch:production",
    "secrets:get:repo:github.com/mozilla-releng/services:branch:staging",
    "secrets:set:garbage/*",
    "secrets:set:repo:github.com/mozilla-releng/services:branch:production",
    "secrets:set:repo:github.com/mozilla-releng/services:branch:staging"
  ],
  "lastModified": "2017-10-17T15:17:37.970Z",
  "roleId": "mozilla-user:mcastelluccio@mozilla.com",
  "scopes": [
    "github:create-status:marco-c/*",
    "secrets:get:repo:github.com/mozilla-releng/services:branch:production",
    "secrets:get:repo:github.com/mozilla-releng/services:branch:staging",
    "secrets:set:repo:github.com/mozilla-releng/services:branch:production",
    "secrets:set:repo:github.com/mozilla-releng/services:branch:staging"
  ]
}

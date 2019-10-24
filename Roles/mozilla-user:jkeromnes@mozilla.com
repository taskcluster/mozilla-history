{
  "created": "2017-09-08T15:32:37.204Z",
  "description": "Create for bug 1396809",
  "expandedScopes": [
    "assume:mozilla-user:jkeromnes@mozilla.com",
    "assume:project:taskcluster:tutorial",
    "queue:route:notify.email.*",
    "secrets:get:repo:github.com/mozilla-releng/services:branch:production",
    "secrets:get:repo:github.com/mozilla-releng/services:branch:staging",
    "secrets:set:repo:github.com/mozilla-releng/services:branch:production",
    "secrets:set:repo:github.com/mozilla-releng/services:branch:staging"
  ],
  "lastModified": "2017-09-08T15:32:37.204Z",
  "roleId": "mozilla-user:jkeromnes@mozilla.com",
  "scopes": [
    "secrets:get:repo:github.com/mozilla-releng/services:branch:production",
    "secrets:get:repo:github.com/mozilla-releng/services:branch:staging",
    "secrets:set:repo:github.com/mozilla-releng/services:branch:production",
    "secrets:set:repo:github.com/mozilla-releng/services:branch:staging"
  ]
}

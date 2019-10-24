{
  "created": "2017-04-06T12:56:55.000Z",
  "description": "Loaner access",
  "expandedScopes": [
    "assume:mozilla-user:mikokm@gmail.com",
    "assume:project:taskcluster:tutorial",
    "queue:get-artifact:private/docker-worker/*",
    "queue:route:notify.email.*"
  ],
  "lastModified": "2017-04-06T12:57:32.353Z",
  "roleId": "mozilla-user:mikokm@gmail.com",
  "scopes": [
    "queue:get-artifact:private/docker-worker/*"
  ]
}

{
  "created": "2017-02-03T17:03:50.628Z",
  "description": "",
  "expandedScopes": [
    "assume:hook-id:project-releng/misc-third-party-library-alert-service",
    "assume:mozilla-user:tritter@mozilla.com",
    "assume:project:taskcluster:tutorial",
    "hooks:modify-hook:project-releng/misc-third-party-library-alert-service",
    "hooks:trigger-hook:project-releng/misc-third-party-library-alert-service",
    "queue:create-task:aws-provisioner-v1/gecko-t-linux-large",
    "queue:route:notify.email.*"
  ],
  "lastModified": "2017-02-03T17:19:59.017Z",
  "roleId": "mozilla-user:tritter@mozilla.com",
  "scopes": [
    "assume:hook-id:project-releng/misc-third-party-library-alert-service",
    "hooks:modify-hook:project-releng/misc-third-party-library-alert-service",
    "hooks:trigger-hook:project-releng/misc-third-party-library-alert-service"
  ]
}

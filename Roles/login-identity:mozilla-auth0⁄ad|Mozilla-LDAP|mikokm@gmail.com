{
  "created": "2018-03-09T13:50:00.329Z",
  "description": "Loaner access",
  "expandedScopes": [
    "assume:login-identity:mozilla-auth0/ad|Mozilla-LDAP|mikokm@gmail.com",
    "auth:create-client:mozilla-auth0/ad|Mozilla-LDAP|mikokm@gmail.com/*",
    "auth:delete-client:mozilla-auth0/ad|Mozilla-LDAP|mikokm@gmail.com/*",
    "auth:reset-access-token:mozilla-auth0/ad|Mozilla-LDAP|mikokm@gmail.com/*",
    "auth:update-client:mozilla-auth0/ad|Mozilla-LDAP|mikokm@gmail.com/*",
    "queue:cancel-task:taskcluster-github/*",
    "queue:create-task:built-in/fail",
    "queue:create-task:built-in/succeed",
    "queue:get-artifact:login-identity/mozilla-auth0/ad|Mozilla-LDAP|mikokm@gmail.com/*",
    "queue:get-artifact:private/docker-worker/*",
    "queue:rerun-task:taskcluster-github/*"
  ],
  "lastModified": "2018-03-09T13:50:00.329Z",
  "roleId": "login-identity:mozilla-auth0/ad|Mozilla-LDAP|mikokm@gmail.com",
  "scopes": [
    "queue:get-artifact:private/docker-worker/*"
  ]
}

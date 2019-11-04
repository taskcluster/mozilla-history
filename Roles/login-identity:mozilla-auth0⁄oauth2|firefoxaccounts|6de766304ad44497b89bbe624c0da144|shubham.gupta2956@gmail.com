{
  "created": "2019-08-09T19:30:35.884Z",
  "description": "",
  "expandedScopes": [
    "assume:login-identity:mozilla-auth0/oauth2|firefoxaccounts|6de766304ad44497b89bbe624c0da144|shubham.gupta2956@gmail.com",
    "auth:create-client:mozilla-auth0/oauth2|firefoxaccounts|6de766304ad44497b89bbe624c0da144|shubham.gupta2956@gmail.com/*",
    "auth:delete-client:mozilla-auth0/oauth2|firefoxaccounts|6de766304ad44497b89bbe624c0da144|shubham.gupta2956@gmail.com/*",
    "auth:reset-access-token:mozilla-auth0/oauth2|firefoxaccounts|6de766304ad44497b89bbe624c0da144|shubham.gupta2956@gmail.com/*",
    "auth:update-client:mozilla-auth0/oauth2|firefoxaccounts|6de766304ad44497b89bbe624c0da144|shubham.gupta2956@gmail.com/*",
    "notify:manage-denylist",
    "queue:create-task:built-in/fail",
    "queue:create-task:built-in/succeed",
    "queue:get-artifact:login-identity/mozilla-auth0/oauth2|firefoxaccounts|6de766304ad44497b89bbe624c0da144|shubham.gupta2956@gmail.com/*"
  ],
  "lastModified": "2019-08-09T19:30:57.299Z",
  "roleId": "login-identity:mozilla-auth0/oauth2|firefoxaccounts|6de766304ad44497b89bbe624c0da144|shubham.gupta2956@gmail.com",
  "scopes": [
    "notify:manage-denylist"
  ]
}

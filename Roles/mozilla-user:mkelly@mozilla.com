{
  "created": "2017-10-18T13:16:24.296Z",
  "description": "",
  "expandedScopes": [
    "assume:hook-id:project-socorro/*",
    "assume:mozilla-user:mkelly@mozilla.com",
    "assume:project-admin:socorro",
    "assume:project:socorro:*",
    "assume:project:taskcluster:tutorial",
    "assume:worker-type:proj-socorro/*",
    "auth:create-client:project/socorro/*",
    "auth:create-role:hook-id:project-socorro/*",
    "auth:create-role:project:socorro:*",
    "auth:delete-client:project/socorro/*",
    "auth:delete-role:hook-id:project-socorro/*",
    "auth:delete-role:project:socorro:*",
    "auth:disable-client:project/socorro/*",
    "auth:enable-client:project/socorro/*",
    "auth:reset-access-token:project/socorro/*",
    "auth:update-client:project/socorro/*",
    "auth:update-role:hook-id:project-socorro/*",
    "auth:update-role:project:socorro:*",
    "auth:websocktunnel-token:taskcluster-net/*",
    "hooks:modify-hook:project-socorro/*",
    "hooks:trigger-hook:project-socorro/*",
    "index:insert-task:project.socorro.*",
    "project:releng:services/tooltool/api/download/internal",
    "project:socorro:*",
    "queue:claim-work:proj-socorro/*",
    "queue:create-task:aws-provisioner-v1/symbol-upload",
    "queue:create-task:aws-provisioner-v1/tutorial",
    "queue:create-task:high:proj-socorro/*",
    "queue:create-task:highest:proj-socorro/*",
    "queue:create-task:low:proj-socorro/*",
    "queue:create-task:lowest:proj-socorro/*",
    "queue:create-task:medium:proj-socorro/*",
    "queue:create-task:very-high:proj-socorro/*",
    "queue:create-task:very-low:proj-socorro/*",
    "queue:get-artifact:project/socorro/*",
    "queue:quarantine-worker:proj-socorro/*",
    "queue:route:index.garbage.*",
    "queue:route:index.project.socorro.*",
    "queue:route:notify.email.*",
    "queue:route:notify.irc-channel.#uptime.*",
    "queue:worker-id:proj-socorro/*",
    "secrets:get:garbage/*",
    "secrets:get:project/socorro/*",
    "secrets:get:worker-type:proj-socorro/*",
    "secrets:set:garbage/*",
    "secrets:set:project/socorro/*"
  ],
  "lastModified": "2017-10-18T13:31:45.423Z",
  "roleId": "mozilla-user:mkelly@mozilla.com",
  "scopes": [
    "assume:project-admin:socorro"
  ]
}

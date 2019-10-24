{
  "created": "2016-03-11T13:01:01.817Z",
  "description": "",
  "expandedScopes": [
    "assume:hook-id:project-gaia/*",
    "assume:mozilla-user:jhylands@mozilla.com",
    "assume:project-admin:gaia",
    "assume:project:gaia:*",
    "assume:project:taskcluster:tutorial",
    "assume:worker-type:proj-gaia/*",
    "auth:create-client:project/gaia/*",
    "auth:create-role:hook-id:project-gaia/*",
    "auth:create-role:project:gaia:*",
    "auth:delete-client:project/gaia/*",
    "auth:delete-role:hook-id:project-gaia/*",
    "auth:delete-role:project:gaia:*",
    "auth:disable-client:project/gaia/*",
    "auth:enable-client:project/gaia/*",
    "auth:reset-access-token:project/gaia/*",
    "auth:update-client:project/gaia/*",
    "auth:update-role:hook-id:project-gaia/*",
    "auth:update-role:project:gaia:*",
    "auth:websocktunnel-token:taskcluster-net/*",
    "hooks:modify-hook:project-gaia/*",
    "hooks:trigger-hook:project-gaia/*",
    "index:insert-task:project.gaia.*",
    "project:gaia:*",
    "queue:claim-work:proj-gaia/*",
    "queue:create-task:high:proj-gaia/*",
    "queue:create-task:highest:proj-gaia/*",
    "queue:create-task:low:proj-gaia/*",
    "queue:create-task:lowest:proj-gaia/*",
    "queue:create-task:medium:proj-gaia/*",
    "queue:create-task:very-high:proj-gaia/*",
    "queue:create-task:very-low:proj-gaia/*",
    "queue:get-artifact:project/gaia/*",
    "queue:quarantine-worker:proj-gaia/*",
    "queue:route:index.project.gaia.*",
    "queue:route:notify.email.*",
    "queue:worker-id:proj-gaia/*",
    "secrets:get:project/gaia/*",
    "secrets:get:worker-pool:proj-gaia/*",
    "secrets:get:worker-type:proj-gaia/*",
    "secrets:set:project/gaia/*"
  ],
  "lastModified": "2016-03-28T23:49:54.625Z",
  "roleId": "mozilla-user:jhylands@mozilla.com",
  "scopes": [
    "assume:project-admin:gaia"
  ]
}

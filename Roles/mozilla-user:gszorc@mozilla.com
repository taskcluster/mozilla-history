{
  "created": "2016-08-02T20:37:53.810Z",
  "description": "* Allowing to gps to view workerTypes... We should probably grant this to everybody once secrets are removed from provisioner and in workers... But for now we just grant to gps, so he can get things done.\n* `secrets:get:repo:github.com/mozilla-releng/OpenCloudConfig:*` added per grenade 2017-01-26",
  "expandedScopes": [
    "assume:mozilla-user:gszorc@mozilla.com",
    "assume:project:taskcluster:tutorial",
    "aws-provisioner:manage-worker-type:*",
    "aws-provisioner:view-worker-type:*",
    "queue:create-task:aws-provisioner-v1/tutorial",
    "queue:create-task:low:built-in/*",
    "queue:route:index.garbage.*",
    "queue:route:notify.email.*",
    "queue:scheduler-id:tutorial",
    "secrets:get:garbage/*",
    "secrets:get:repo:github.com/mozilla-releng/OpenCloudConfig:*",
    "secrets:get:repo:github.com/taskcluster/docker-worker:*",
    "secrets:set:garbage/*"
  ],
  "lastModified": "2017-12-19T18:32:08.897Z",
  "roleId": "mozilla-user:gszorc@mozilla.com",
  "scopes": [
    "aws-provisioner:manage-worker-type:*",
    "aws-provisioner:view-worker-type:*",
    "secrets:get:repo:github.com/mozilla-releng/OpenCloudConfig:*",
    "secrets:get:repo:github.com/taskcluster/docker-worker:*"
  ]
}

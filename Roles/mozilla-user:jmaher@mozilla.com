{
  "created": "2016-10-12T16:42:45.635Z",
  "description": "Custom role for Joel Maher, created by pmoore on 12 October 2016. Initial purpose is to provide scopes that allows Joel to help out with Windows tests troubleshooting in gecko.",
  "expandedScopes": [
    "assume:mozilla-user:jmaher@mozilla.com",
    "assume:project:releng:gecko-win-test-troubleshooter",
    "assume:project:taskcluster:tutorial",
    "generic-worker:allow-rdp:releng-hardware/gecko-t-win*",
    "queue:create-task:aws-provisioner-v1/tutorial",
    "queue:route:index.garbage.*",
    "queue:route:notify.email.*",
    "secrets:get:garbage/*",
    "secrets:get:repo:github.com/mozilla-releng/OpenCloudConfig:gecko-t-win*",
    "secrets:set:garbage/*"
  ],
  "lastModified": "2016-10-26T14:07:56.925Z",
  "roleId": "mozilla-user:jmaher@mozilla.com",
  "scopes": [
    "assume:project:releng:gecko-win-test-troubleshooter"
  ]
}

{
  "created": "2016-10-10T15:45:03.733Z",
  "description": "",
  "expandedScopes": [
    "assume:hook-id:garbage/*",
    "assume:hook-id:project-deepspeech/*",
    "assume:mozilla-user:alissy@mozilla.com",
    "assume:project-admin:deepspeech",
    "assume:project:deepspeech:*",
    "assume:project:taskcluster:contributor",
    "assume:project:taskcluster:diagnostics",
    "assume:project:taskcluster:gecko:level-1-sccache-buckets",
    "assume:project:taskcluster:generic-worker-tester",
    "assume:project:taskcluster:gps-docker-worker-test-scopes",
    "assume:project:taskcluster:index-test-scopes",
    "assume:project:taskcluster:mozilla-github-repository",
    "assume:project:taskcluster:notify",
    "assume:project:taskcluster:taskcluster-lib-artifact-go-tester",
    "assume:project:taskcluster:taskcluster-proxy-tester",
    "assume:project:taskcluster:tc-client-generator-tester",
    "assume:project:taskcluster:tutorial",
    "assume:project:taskcluster:worker-test-scopes",
    "assume:repo:github.com/JohanLorenzo/fenix:cron:raptor",
    "assume:repo:github.com/mozilla/DeepSpeech:*",
    "assume:repo:github.com/mozilla/tensorflow:*",
    "assume:repo:hg.mozilla.org/try:*",
    "assume:worker-id:dummy-test-workergroup/dummy-test-worker-id",
    "assume:worker-id:my-worker-group/my-worker",
    "assume:worker-id:test-worker-group/test-worker-id",
    "assume:worker-type:dummy-test-provisioner/dummy-test-worker-type",
    "assume:worker-type:proj-deepspeech/*",
    "assume:worker-type:test-provisioner/*",
    "auth:aws-s3:read-write:taskcluster-level-1-sccache-eu-central-1/*",
    "auth:aws-s3:read-write:taskcluster-level-1-sccache-us-east-1/*",
    "auth:aws-s3:read-write:taskcluster-level-1-sccache-us-west-1/*",
    "auth:aws-s3:read-write:taskcluster-level-1-sccache-us-west-2/*",
    "auth:aws-s3:read-write:test-bucket-for-any-garbage/*",
    "auth:azure-table:read-write:fakeaccount/DuMmYtAbLe",
    "auth:create-client:garbage/*",
    "auth:create-client:project/deepspeech/*",
    "auth:create-client:project/taskcluster:generic-worker-tester/TestReclaimCancelledTask",
    "auth:create-client:project/taskcluster:generic-worker-tester/TestResolveResolvedTask",
    "auth:create-role:hook-id:project-deepspeech/*",
    "auth:create-role:project:deepspeech:*",
    "auth:create-role:repo:github.com/mozilla/DeepSpeech:*",
    "auth:create-role:repo:github.com/mozilla/tensorflow:*",
    "auth:delete-client:project/deepspeech/*",
    "auth:delete-role:hook-id:project-deepspeech/*",
    "auth:delete-role:project:deepspeech:*",
    "auth:delete-role:repo:github.com/mozilla/DeepSpeech:*",
    "auth:delete-role:repo:github.com/mozilla/tensorflow:*",
    "auth:disable-client:project/deepspeech/*",
    "auth:enable-client:project/deepspeech/*",
    "auth:gcp:access-token:sccache-3/sccache-l1*",
    "auth:gcp:access-token:sccache-3/tc-l1*",
    "auth:reset-access-token:project/deepspeech/*",
    "auth:sentry:ds-docker-worker",
    "auth:sentry:generic-worker",
    "auth:sentry:generic-worker-tests",
    "auth:sentry:playground",
    "auth:sentry:test-dummy-worker",
    "auth:statsum:ds-docker-worker",
    "auth:statsum:playground",
    "auth:statsum:test-dummy-worker",
    "auth:update-client:project/deepspeech/*",
    "auth:update-role:hook-id:project-deepspeech/*",
    "auth:update-role:project:deepspeech:*",
    "auth:update-role:repo:github.com/mozilla/DeepSpeech:*",
    "auth:update-role:repo:github.com/mozilla/tensorflow:*",
    "auth:webhooktunnel",
    "auth:websocktunnel-token:taskcluster-net/*",
    "aws-provisioner:manage-worker-type:deepspeech-*",
    "docker-worker:cache:*",
    "docker-worker:capability:device:loopbackAudio",
    "docker-worker:capability:device:loopbackVideo",
    "docker-worker:capability:privileged",
    "docker-worker:feature:allowPtrace",
    "docker-worker:image:quay.io/mozilla/builder:*",
    "docker-worker:image:quay.io/mozilla/decision:*",
    "docker-worker:image:taskcluster/builder:*",
    "docker-worker:image:taskcluster/tester:*",
    "docker-worker:relengapi-proxy:tooltool.download.internal",
    "docker-worker:relengapi-proxy:tooltool.download.public",
    "ec2-manager:manage-resources:deepspeech*",
    "generic-worker:allow-rdp:aws-provisioner-v1/deepspeech*",
    "generic-worker:allow-rdp:aws-provisioner-v1/gecko-1-b-win*",
    "generic-worker:allow-rdp:aws-provisioner-v1/gecko-t-win*",
    "generic-worker:cache:banana-cache",
    "generic-worker:cache:deepspeech-*",
    "generic-worker:cache:devtools-app",
    "generic-worker:cache:gecko-level-1-*",
    "generic-worker:cache:level-1-*",
    "generic-worker:cache:test-modifications",
    "generic-worker:cache:unknown-issuer-app-cache",
    "generic-worker:os-group:aws-provisioner-v1/deepspeech*",
    "generic-worker:os-group:aws-provisioner-v1/gecko-t-win10-64-alpha/Administrators",
    "generic-worker:os-group:aws-provisioner-v1/gecko-t-win10-64/Administrators",
    "generic-worker:os-group:aws-provisioner-v1/gecko-t-win7-32/Administrators",
    "generic-worker:os-group:test-provisioner/*",
    "generic-worker:run-as-administrator:aws-provisioner-v1/deepspeech*",
    "generic-worker:run-as-administrator:aws-provisioner-v1/gecko-t-win10-64",
    "generic-worker:run-as-administrator:aws-provisioner-v1/gecko-t-win10-64-alpha",
    "generic-worker:run-as-administrator:test-provisioner/*",
    "hooks:modify-hook:garbage/*",
    "hooks:modify-hook:project-deepspeech/*",
    "hooks:trigger-hook:garbage/*",
    "hooks:trigger-hook:project-deepspeech/*",
    "in-tree:hook-action:project-gecko/in-tree-action-1-*",
    "index:insert-task:garbage.*",
    "index:insert-task:gecko.cache.level-1.*",
    "index:insert-task:gecko.v2.try.*",
    "index:insert-task:project.deepspeech.*",
    "notify:email:*",
    "notify:irc-channel:*",
    "notify:irc-user:*",
    "project:deepspeech:*",
    "project:mobile:fenix:releng:signing:cert:dep-signing",
    "project:mobile:fenix:releng:signing:format:*",
    "project:releng:addons.mozilla.org:server:staging",
    "project:releng:balrog:action:*",
    "project:releng:balrog:channel:*",
    "project:releng:balrog:server:dep",
    "project:releng:beetmover:action:*",
    "project:releng:beetmover:bucket:dep",
    "project:releng:beetmover:bucket:dep-partner",
    "project:releng:beetmover:bucket:maven-staging",
    "project:releng:bouncer:action:*",
    "project:releng:bouncer:server:staging",
    "project:releng:bouncer:server:staging-nazgul",
    "project:releng:googleplay:dep",
    "project:releng:services/tooltool/api/download/internal",
    "project:releng:services/tooltool/api/download/public",
    "project:releng:ship-it:action:mark-as-shipped",
    "project:releng:ship-it:action:mark-as-started",
    "project:releng:ship-it:server:staging",
    "project:releng:signing:cert:dep-signing",
    "project:releng:signing:format:*",
    "project:releng:snapcraft:firefox:mock",
    "project:releng:treescript:action:*",
    "purge-cache:aws-provisioner-v1/*",
    "purge-cache:deepspeech-provisioner/*",
    "purge-cache:garbage*",
    "purge-cache:verifyprovisioner/verifyworker:verifycache",
    "queue:cancel-task:gecko-level-1/*",
    "queue:cancel-task:test-scheduler/*",
    "queue:claim-task:dummy-test-provisioner/dummy-test-worker-type",
    "queue:claim-task:no-provisioner/my-workertype",
    "queue:claim-work:deepspeech-provisioner/*",
    "queue:claim-work:dummy-test-provisioner/dummy-test-worker-type",
    "queue:claim-work:proj-deepspeech/*",
    "queue:claim-work:test-dummy-provisioner/dummy-worker-*",
    "queue:claim-work:test-provisioner/*",
    "queue:create-artifact:public/*",
    "queue:create-task:aws-provisioner-v1/ami-test*",
    "queue:create-task:aws-provisioner-v1/deepspeech-*",
    "queue:create-task:aws-provisioner-v1/github-worker",
    "queue:create-task:aws-provisioner-v1/github-worker-windows",
    "queue:create-task:aws-provisioner-v1/mobile-1-decision",
    "queue:create-task:aws-provisioner-v1/tutorial",
    "queue:create-task:aws-provisioner-v1/win2012r2",
    "queue:create-task:dummy-test-provisioner/dummy-test-type",
    "queue:create-task:dummy-test-provisioner/dummy-test-worker-type",
    "queue:create-task:gecko-t-tc-worker/*",
    "queue:create-task:high:deepspeech-provisioner/*",
    "queue:create-task:high:proj-deepspeech/*",
    "queue:create-task:highest:aws-provisioner-v1/gecko-1-decision",
    "queue:create-task:highest:aws-provisioner-v1/gecko-1-images",
    "queue:create-task:highest:aws-provisioner-v1/github-worker",
    "queue:create-task:highest:aws-provisioner-v1/mobile-1-*",
    "queue:create-task:highest:aws-provisioner-v1/mobile-3-b-ref-browser",
    "queue:create-task:highest:deepspeech-provisioner/*",
    "queue:create-task:highest:proj-deepspeech/*",
    "queue:create-task:highest:scriptworker-prov-v1/mobile-signing-dep-v1",
    "queue:create-task:highest:test-provisioner/gecko-1-b-win2012-gamma",
    "queue:create-task:highest:win-provisioner/win2008-worker",
    "queue:create-task:localprovisioner/*",
    "queue:create-task:low:aws-provisioner-v1/ami-test*",
    "queue:create-task:low:aws-provisioner-v1/gecko-1-*",
    "queue:create-task:low:aws-provisioner-v1/gecko-misc",
    "queue:create-task:low:aws-provisioner-v1/gecko-t-*",
    "queue:create-task:low:bitbar/gecko-t-*",
    "queue:create-task:low:built-in/*",
    "queue:create-task:low:deepspeech-provisioner/*",
    "queue:create-task:low:dummy-test-provisioner/dummy-test-type",
    "queue:create-task:low:gce/gecko-1-*",
    "queue:create-task:low:gce/gecko-t-*",
    "queue:create-task:low:localprovisioner/*",
    "queue:create-task:low:manual-packet/tc-worker-docker-v0",
    "queue:create-task:low:manual-packet/tc-worker-docker-v1",
    "queue:create-task:low:manual-packet/tc-worker-qemu-v1",
    "queue:create-task:low:packetnet/*",
    "queue:create-task:low:proj-deepspeech/*",
    "queue:create-task:low:releng-hardware/gecko-1-*",
    "queue:create-task:low:releng-hardware/gecko-t-*",
    "queue:create-task:low:releng-hardware/my-talos",
    "queue:create-task:low:scriptworker-k8s/gecko-1-*",
    "queue:create-task:low:scriptworker-k8s/gecko-t-*",
    "queue:create-task:low:scriptworker-prov-v1/addon-dev",
    "queue:create-task:low:scriptworker-prov-v1/balrog-dev",
    "queue:create-task:low:scriptworker-prov-v1/beetmoverworker-dev",
    "queue:create-task:low:scriptworker-prov-v1/bouncer-dev",
    "queue:create-task:low:scriptworker-prov-v1/dep-pushapk",
    "queue:create-task:low:scriptworker-prov-v1/dep-pushsnap",
    "queue:create-task:low:scriptworker-prov-v1/depsigning",
    "queue:create-task:low:scriptworker-prov-v1/depsigning-mac-v1",
    "queue:create-task:low:scriptworker-prov-v1/shipit-dev*",
    "queue:create-task:low:scriptworker-prov-v1/signing-linux-dev",
    "queue:create-task:low:scriptworker-prov-v1/signing-mac-dev",
    "queue:create-task:low:scriptworker-prov-v1/treescript-dev",
    "queue:create-task:low:tc-worker-provisioner/*",
    "queue:create-task:low:terraform-packet/*",
    "queue:create-task:low:test-dummy-provisioner/*",
    "queue:create-task:lowest:aws-provisioner-v1/deepspeech-*",
    "queue:create-task:lowest:deepspeech-provisioner/*",
    "queue:create-task:lowest:localprovisioner/deepspeech-macos",
    "queue:create-task:lowest:no-provisioner/my-workertype",
    "queue:create-task:lowest:proj-deepspeech/*",
    "queue:create-task:lowest:test-provisioner/*",
    "queue:create-task:medium:deepspeech-provisioner/*",
    "queue:create-task:medium:proj-autophone/*",
    "queue:create-task:medium:proj-awfy/*",
    "queue:create-task:medium:proj-deepspeech/*",
    "queue:create-task:medium:terraform-packet/tc-worker-docker-v1-*",
    "queue:create-task:no-provisioner/test-worker",
    "queue:create-task:packetnet/*",
    "queue:create-task:tc-worker-provisioner/*",
    "queue:create-task:test-dummy-provisioner/*",
    "queue:create-task:test-provisioner/*",
    "queue:create-task:very-high:deepspeech-provisioner/*",
    "queue:create-task:very-high:proj-deepspeech/*",
    "queue:create-task:very-low:deepspeech-provisioner/*",
    "queue:create-task:very-low:proj-deepspeech/*",
    "queue:define-task:test-dummy-provisioner/dummy-worker-*",
    "queue:get-artifact:SampleArtifacts/_/X.txt",
    "queue:get-artifact:SampleArtifacts/_/non-existent-artifact.txt",
    "queue:get-artifact:SampleArtifacts/b/c/d.jpg",
    "queue:get-artifact:private/docker-worker/*",
    "queue:get-artifact:private/openh264/*",
    "queue:get-artifact:project/deepspeech/*",
    "queue:get-artifact:project/gecko/android-*",
    "queue:get-artifact:releng/partner/*",
    "queue:get-artifact:taskcluster-proxy-test/512-random-bytes",
    "queue:poll-task-urls",
    "queue:poll-task-urls:test-dummy-provisioner/dummy-worker-*",
    "queue:quarantine-worker:deepspeech-provisioner/*",
    "queue:quarantine-worker:proj-deepspeech/*",
    "queue:report-task-completed:dummy-test-provisioner/dummy-test-worker-type",
    "queue:rerun-task:gecko-level-1/*",
    "queue:rerun-task:taskcluster-github/*",
    "queue:resolve-task",
    "queue:route:checks",
    "queue:route:coalesce.v1.try.*",
    "queue:route:garbage.*",
    "queue:route:index.garbage.*",
    "queue:route:index.gecko.cache.level-1.*",
    "queue:route:index.gecko.heavyprofile",
    "queue:route:index.gecko.v2.try.*",
    "queue:route:index.project.deepspeech.*",
    "queue:route:index.project.mobile.fenix.cache.level-1.*",
    "queue:route:index.project.mobile.fenix.v2.staging.*",
    "queue:route:index.project.mobile.fenix.v3.staging.*",
    "queue:route:notify.*",
    "queue:route:project.relman.codereview.*",
    "queue:route:statuses",
    "queue:route:tc-treeherder-stage.try.*",
    "queue:route:tc-treeherder-stage.v2.try.*",
    "queue:route:tc-treeherder.try.*",
    "queue:route:tc-treeherder.v2.try.*",
    "queue:route:test-notify.*",
    "queue:schedule-task:test-dummy-scheduler/*",
    "queue:scheduler-id:-",
    "queue:scheduler-id:garbage-level-1",
    "queue:scheduler-id:gecko-level-1",
    "queue:scheduler-id:go-test-test-scheduler",
    "queue:scheduler-id:mobile-level-1",
    "queue:scheduler-id:taskcluster-github",
    "queue:scheduler-id:test-scheduler",
    "queue:task-group-id:test-dummy-scheduler/*",
    "queue:worker-id:deepspeech-workers/*",
    "queue:worker-id:dummy-test-workergroup/dummy-test-worker-id",
    "queue:worker-id:my-worker-group/my-worker",
    "queue:worker-id:proj-deepspeech/*",
    "queue:worker-id:test-dummy-workers/dummy-worker-*",
    "queue:worker-id:test-worker-group/*",
    "secrets:get:garbage/*",
    "secrets:get:project/deepspeech/*",
    "secrets:get:project/releng/gecko/build/level-1/*",
    "secrets:get:project/relman/coverity",
    "secrets:get:project/taskcluster/gecko/hgfingerprint",
    "secrets:get:project/taskcluster/gecko/hgmointernal",
    "secrets:get:project/taskcluster/taskcluster-worker/stateless-dns",
    "secrets:get:worker-pool:dummy-test-provisioner/dummy-test-worker-type",
    "secrets:get:worker-pool:proj-deepspeech/*",
    "secrets:get:worker-pool:test-provisioner/*",
    "secrets:get:worker-type:dummy-test-provisioner/dummy-test-worker-type",
    "secrets:get:worker-type:proj-deepspeech/*",
    "secrets:get:worker-type:test-provisioner/*",
    "secrets:set:garbage/*",
    "secrets:set:project/deepspeech/*",
    "worker:cache:dummy-cache-*"
  ],
  "lastModified": "2018-03-09T01:47:55.276Z",
  "roleId": "mozilla-user:alissy@mozilla.com",
  "scopes": [
    "assume:project-admin:deepspeech",
    "assume:project:taskcluster:contributor",
    "queue:scheduler-id:taskcluster-github"
  ]
}

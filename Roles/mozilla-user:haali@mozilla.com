{
  "created": "2016-06-24T21:50:55.167Z",
  "description": "",
  "expandedScopes": [
    "assume:mozilla-group:scm_level_3",
    "assume:mozilla-user:haali@mozilla.com",
    "assume:project:taskcluster:tutorial",
    "purge-cache:aws-provisioner-v1/b2gtest-emulator:level-1-*",
    "purge-cache:aws-provisioner-v1/b2gtest-emulator:tooltool-cache",
    "purge-cache:aws-provisioner-v1/b2gtest:level-1-*",
    "purge-cache:aws-provisioner-v1/b2gtest:tooltool-cache",
    "purge-cache:aws-provisioner-v1/balrog:level-1-*",
    "purge-cache:aws-provisioner-v1/balrog:tooltool-cache",
    "purge-cache:aws-provisioner-v1/cli:level-1-*",
    "purge-cache:aws-provisioner-v1/cli:tooltool-cache",
    "purge-cache:aws-provisioner-v1/cratertest:level-1-*",
    "purge-cache:aws-provisioner-v1/cratertest:tooltool-cache",
    "purge-cache:aws-provisioner-v1/dbg-linux32:level-1-*",
    "purge-cache:aws-provisioner-v1/dbg-linux32:tooltool-cache",
    "purge-cache:aws-provisioner-v1/dbg-linux64:level-1-*",
    "purge-cache:aws-provisioner-v1/dbg-linux64:tooltool-cache",
    "purge-cache:aws-provisioner-v1/dbg-macosx64:level-1-*",
    "purge-cache:aws-provisioner-v1/dbg-macosx64:tooltool-cache",
    "purge-cache:aws-provisioner-v1/desktop-test-xlarge:level-1-*",
    "purge-cache:aws-provisioner-v1/desktop-test-xlarge:tooltool-cache",
    "purge-cache:aws-provisioner-v1/desktop-test:level-1-*",
    "purge-cache:aws-provisioner-v1/desktop-test:tooltool-cache",
    "purge-cache:aws-provisioner-v1/gecko-decision:level-1-*",
    "purge-cache:aws-provisioner-v1/gecko-decision:tooltool-cache",
    "purge-cache:aws-provisioner-v1/mulet-debug:level-1-*",
    "purge-cache:aws-provisioner-v1/mulet-debug:tooltool-cache",
    "purge-cache:aws-provisioner-v1/mulet-opt:level-1-*",
    "purge-cache:aws-provisioner-v1/mulet-opt:tooltool-cache",
    "purge-cache:aws-provisioner-v1/opt-linux32:level-1-*",
    "purge-cache:aws-provisioner-v1/opt-linux32:tooltool-cache",
    "purge-cache:aws-provisioner-v1/opt-linux64:level-1-*",
    "purge-cache:aws-provisioner-v1/opt-linux64:tooltool-cache",
    "purge-cache:aws-provisioner-v1/opt-macosx64:level-1-*",
    "purge-cache:aws-provisioner-v1/opt-macosx64:tooltool-cache",
    "purge-cache:aws-provisioner-v1/spidermonkey:level-1-*",
    "purge-cache:aws-provisioner-v1/spidermonkey:tooltool-cache",
    "purge-cache:aws-provisioner-v1/symbol-upload:level-1-*",
    "purge-cache:aws-provisioner-v1/symbol-upload:tooltool-cache",
    "queue:cancel-task:*",
    "queue:route:index.garbage.hassan.*",
    "queue:route:notify.email.*"
  ],
  "lastModified": "2017-07-12T20:53:10.938Z",
  "roleId": "mozilla-user:haali@mozilla.com",
  "scopes": [
    "assume:mozilla-group:scm_level_3",
    "purge-cache:aws-provisioner-v1/b2gtest-emulator:level-1-*",
    "purge-cache:aws-provisioner-v1/b2gtest-emulator:tooltool-cache",
    "purge-cache:aws-provisioner-v1/b2gtest:level-1-*",
    "purge-cache:aws-provisioner-v1/b2gtest:tooltool-cache",
    "purge-cache:aws-provisioner-v1/balrog:level-1-*",
    "purge-cache:aws-provisioner-v1/balrog:tooltool-cache",
    "purge-cache:aws-provisioner-v1/cli:level-1-*",
    "purge-cache:aws-provisioner-v1/cli:tooltool-cache",
    "purge-cache:aws-provisioner-v1/cratertest:level-1-*",
    "purge-cache:aws-provisioner-v1/cratertest:tooltool-cache",
    "purge-cache:aws-provisioner-v1/dbg-linux32:level-1-*",
    "purge-cache:aws-provisioner-v1/dbg-linux32:tooltool-cache",
    "purge-cache:aws-provisioner-v1/dbg-linux64:level-1-*",
    "purge-cache:aws-provisioner-v1/dbg-linux64:tooltool-cache",
    "purge-cache:aws-provisioner-v1/dbg-macosx64:level-1-*",
    "purge-cache:aws-provisioner-v1/dbg-macosx64:tooltool-cache",
    "purge-cache:aws-provisioner-v1/desktop-test-xlarge:level-1-*",
    "purge-cache:aws-provisioner-v1/desktop-test-xlarge:tooltool-cache",
    "purge-cache:aws-provisioner-v1/desktop-test:level-1-*",
    "purge-cache:aws-provisioner-v1/desktop-test:tooltool-cache",
    "purge-cache:aws-provisioner-v1/gecko-decision:level-1-*",
    "purge-cache:aws-provisioner-v1/gecko-decision:tooltool-cache",
    "purge-cache:aws-provisioner-v1/mulet-debug:level-1-*",
    "purge-cache:aws-provisioner-v1/mulet-debug:tooltool-cache",
    "purge-cache:aws-provisioner-v1/mulet-opt:level-1-*",
    "purge-cache:aws-provisioner-v1/mulet-opt:tooltool-cache",
    "purge-cache:aws-provisioner-v1/opt-linux32:level-1-*",
    "purge-cache:aws-provisioner-v1/opt-linux32:tooltool-cache",
    "purge-cache:aws-provisioner-v1/opt-linux64:level-1-*",
    "purge-cache:aws-provisioner-v1/opt-linux64:tooltool-cache",
    "purge-cache:aws-provisioner-v1/opt-macosx64:level-1-*",
    "purge-cache:aws-provisioner-v1/opt-macosx64:tooltool-cache",
    "purge-cache:aws-provisioner-v1/spidermonkey:level-1-*",
    "purge-cache:aws-provisioner-v1/spidermonkey:tooltool-cache",
    "purge-cache:aws-provisioner-v1/symbol-upload:level-1-*",
    "purge-cache:aws-provisioner-v1/symbol-upload:tooltool-cache",
    "queue:cancel-task:*",
    "queue:route:index.garbage.hassan.*"
  ]
}

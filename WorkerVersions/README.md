

# Worker Pool Versions


## Generic Worker

Total: `55`

Count by version:

| Version | Count |
| :--- | ---: |
| 14.1.2 | 1 |
| 15.1.4 | 1 |
| 16.2.0 | 4 |
| 16.5.1 | 1 |
| 30.0.2 | 3 |
| 36.0.0 | 1 |
| 40.0.3 | 3 |
| 43.2.0 | 3 |
| 44.0.0 | 1 |
| 45.0.0 | 1 |
| 54.4.1 | 26 |
| 55.1.1 | 4 |
| 55.3.4 | 5 |
| 57.0.1 | 1 |


Count by image:

| Version | Count |
| :--- | ---: |
| /subscriptions/108d46d5-fe9b-4850-9a7d-8c914aa6c1f0/resourceGroups/rg-packer-through-cib/providers/Microsoft.Compute/images/win11-64-2009-centralus-win11-22h2-avd-531c001 | 2 |
| /subscriptions/108d46d5-fe9b-4850-9a7d-8c914aa6c1f0/resourceGroups/rg-packer-through-cib/providers/Microsoft.Compute/images/win11-64-2009-eastus-win11-22h2-avd-531c001 | 2 |
| /subscriptions/108d46d5-fe9b-4850-9a7d-8c914aa6c1f0/resourceGroups/rg-packer-through-cib/providers/Microsoft.Compute/images/win11-64-2009-eastus-win11-22h2-avd-alpha | 1 |
| projects/taskcluster-imaging/global/images/generic-translations-gcp-googlecompute-2023-11-01t20-12-18z | 1 |
| /subscriptions/a30e97ab-734a-4f3b-a0e4-c51c0bff0701/resourceGroups/rg-packer-through-cib/providers/Microsoft.Compute/images/trusted-win2012r2-64-vs-py2-l3-centralus-2012-R2-Datacenter-7aa76c6 | 1 |
| ami-07db1c26bda143e33 | 1 |
| projects/taskcluster-imaging/global/images/generic-translations-gcp-googlecompute-2023-09-12t21-35-55z | 4 |
| /subscriptions/108d46d5-fe9b-4850-9a7d-8c914aa6c1f0/resourceGroups/rg-packer-through-cib/providers/Microsoft.Compute/images/win2022-64-2009-centralus-2022-datacenter-azure-edition-531c001 | 5 |
| /subscriptions/108d46d5-fe9b-4850-9a7d-8c914aa6c1f0/resourceGroups/rg-packer-through-cib/providers/Microsoft.Compute/images/win10-64-2009-eastus-win10-22h2-avd-531c001 | 1 |
| /subscriptions/a30e97ab-734a-4f3b-a0e4-c51c0bff0701/resourceGroups/rg-packer-through-cib/providers/Microsoft.Compute/images/trusted-win2012r2-64-l3-centralus-2012-R2-Datacenter-7cde253 | 1 |
| /subscriptions/108d46d5-fe9b-4850-9a7d-8c914aa6c1f0/resourceGroups/rg-packer-through-cib/providers/Microsoft.Compute/images/win11-64-2009-canadacentral-win11-22h2-avd-531c001 | 1 |
| /subscriptions/a30e97ab-734a-4f3b-a0e4-c51c0bff0701/resourceGroups/rg-packer-through-cib/providers/Microsoft.Compute/images/trusted-win2012r2-64-vs-l3-centralus-2012-R2-Datacenter-7cde253 | 1 |
|  | 26 |
| /subscriptions/108d46d5-fe9b-4850-9a7d-8c914aa6c1f0/resourceGroups/rg-packer-through-cib/providers/Microsoft.Compute/images/win11-64-2009-uksouth-win11-22h2-avd-alpha | 2 |
| /subscriptions/108d46d5-fe9b-4850-9a7d-8c914aa6c1f0/resourceGroups/rg-packer-through-cib/providers/Microsoft.Compute/images/win10-64-2009-canadacentral-win10-22h2-avd-531c001 | 2 |
| /subscriptions/a30e97ab-734a-4f3b-a0e4-c51c0bff0701/resourceGroups/rg-packer-through-cib/providers/Microsoft.Compute/images/trusted-win2022-64-2009-centralus-2022-datacenter-azure-edition-d5dee60 | 3 |
| unknown | 1 |


| Worker Pool | Implementation | Version | Engine | Revision | OS | Arch | GO | Total Workers | Total Capacity |
| --- | --- | --- | --- | --- | --- | --- | --- | ---: | ---: |
| **bitbar/gecko-t-win64-aarch64-laptop** | generic-worker | 14.1.2 | <UNKNOWN> | 13118c4c1b | windows | 386 | 1.10.8 | 0 | 0 |
| **comm-1/b-win2022** | generic-worker | 54.4.1 | multiuser | a38a78d98c | windows | amd64 | 1.20.6 | 4 | 4 |
| **comm-2/b-win2022** | generic-worker | 54.4.1 | multiuser | a38a78d98c | windows | amd64 | 1.20.6 | 4 | 4 |
| **comm-3/b-win2012-azure** | generic-worker | 43.2.0 | multiuser | 53fa7b79bc | windows | amd64 | 1.16.3 | 122 | 122 |
| **comm-3/b-win2022** | generic-worker | 54.4.1 | multiuser | a38a78d98c | windows | amd64 | 1.20.6 | 529 | 529 |
| **gecko-1/b-win2022** | generic-worker | 54.4.1 | multiuser | a38a78d98c | windows | amd64 | 1.20.6 | 829 | 829 |
| **gecko-2/b-win2022** | generic-worker | 54.4.1 | multiuser | a38a78d98c | windows | amd64 | 1.20.6 | 57 | 57 |
| **gecko-3/b-win2022** | generic-worker | 54.4.1 | multiuser | a38a78d98c | windows | amd64 | 1.20.6 | 5885 | 5885 |
| **gecko-t/win10-64-2009** | generic-worker | 54.4.1 | multiuser | a38a78d98c | windows | amd64 | 1.20.6 | 5047 | 5047 |
| **gecko-t/win10-64-2009-gpu** | generic-worker | 54.4.1 | multiuser | a38a78d98c | windows | amd64 | 1.20.6 | 618 | 618 |
| **gecko-t/win10-64-2009-source** | generic-worker | 54.4.1 | multiuser | a38a78d98c | windows | amd64 | 1.20.6 | 4 | 4 |
| **gecko-t/win11-64-2009** | generic-worker | 54.4.1 | multiuser | a38a78d98c | windows | amd64 | 1.20.6 | 20739 | 20739 |
| **gecko-t/win11-64-2009-alpha** | generic-worker | 54.4.1 | multiuser | a38a78d98c | windows | amd64 | 1.20.6 | 4 | 4 |
| **gecko-t/win11-64-2009-gpu** | generic-worker | 54.4.1 | multiuser | a38a78d98c | windows | amd64 | 1.20.6 | 17767 | 17767 |
| **gecko-t/win11-64-2009-gpu-alpha** | generic-worker | 54.4.1 | multiuser | a38a78d98c | windows | amd64 | 1.20.6 | 4 | 4 |
| **gecko-t/win11-64-2009-source** | generic-worker | 54.4.1 | multiuser | a38a78d98c | windows | amd64 | 1.20.6 | 1234 | 1234 |
| **gecko-t/win11-64-2009-source-alpha** | generic-worker | 54.4.1 | multiuser | a38a78d98c | windows | amd64 | 1.20.6 | 4 | 4 |
| **gecko-t/win11-64-2009-ssd** | generic-worker | 54.4.1 | multiuser | a38a78d98c | windows | amd64 | 1.20.6 | 2054 | 2054 |
| **gecko-t/win11-64-2009-ssd-gpu** | generic-worker | 54.4.1 | multiuser | a38a78d98c | windows | amd64 | 1.20.6 | 1762 | 1762 |
| **localprovisioner/nss-macos-10-12** | generic-worker | 15.1.4 | simple | c407e45e3f | darwin | amd64 | 1.10.8 | 0 | 0 |
| **localprovisioner/nss-macos-m1** | generic-worker | 30.0.2 | simple | 6fdba0dad3 | darwin | arm64 | 1.16.4 | 0 | 0 |
| **mozillavpn-1/b-win2022** | generic-worker | 54.4.1 | multiuser | a38a78d98c | windows | amd64 | 1.20.6 | 163 | 163 |
| **mozillavpn-3/b-win2012-azure** | generic-worker | 43.2.0 | multiuser | 53fa7b79bc | windows | amd64 | 1.16.3 | 11 | 11 |
| **mozillavpn-3/b-win2022** | generic-worker | 54.4.1 | multiuser | a38a78d98c | windows | amd64 | 1.20.6 | 72 | 72 |
| **nss-3/b-win2012-azure** | generic-worker | 43.2.0 | multiuser | 53fa7b79bc | windows | amd64 | 1.16.3 | 37 | 37 |
| **performance-hardware/gecko-t-fxrecorder** | generic-worker | 44.0.0 | multiuser | faf6f319c0 | windows | amd64 | 1.16.3 | 0 | 0 |
| **proj-autophone/gecko-t-bitbar-gw-unit-p5** | generic-worker | 36.0.0 | simple | b9cb11293f | linux | amd64 | 1.13.7 | 0 | 0 |
| **releng-hardware/applicationservices-b-1-osx1015** | generic-worker | 54.4.1 | multiuser | a38a78d98c | darwin | amd64 | 1.20.6 | 0 | 0 |
| **releng-hardware/applicationservices-b-3-osx1015** | generic-worker | 54.4.1 | multiuser | a38a78d98c | darwin | amd64 | 1.20.6 | 0 | 0 |
| **releng-hardware/gecko-1-b-osx-1015** | generic-worker | 40.0.3 | multiuser | 5f7a31ac6e | darwin | amd64 | 1.15.6 | 0 | 0 |
| **releng-hardware/gecko-1-b-osx-1015-staging** | generic-worker | 54.4.1 | multiuser | a38a78d98c | darwin | amd64 | 1.20.6 | 0 | 0 |
| **releng-hardware/gecko-3-b-osx-1015** | generic-worker | 55.3.4 | multiuser | a71f300fcb | darwin | amd64 | 1.21.1 | 0 | 0 |
| **releng-hardware/gecko-t-osx-1015-power** | generic-worker | 40.0.3 | simple | 5f7a31ac6e | darwin | amd64 | 1.15.6 | 0 | 0 |
| **releng-hardware/gecko-t-osx-1015-r8** | generic-worker | 55.3.4 | simple | a71f300fcb | darwin | amd64 | 1.21.1 | 0 | 0 |
| **releng-hardware/gecko-t-osx-1015-r8-staging** | generic-worker | 40.0.3 | simple | 5f7a31ac6e | darwin | amd64 | 1.15.6 | 0 | 0 |
| **releng-hardware/gecko-t-osx-1100-m1** | generic-worker | 55.3.4 | simple | a71f300fcb | darwin | arm64 | 1.21.1 | 0 | 0 |
| **releng-hardware/gecko-t-osx-1100-m1-loaner** | generic-worker | 30.0.2 | simple | 6fdba0dad3 | darwin | arm64 | 1.16.4 | 0 | 0 |
| **releng-hardware/gecko-t-osx-1100-m1-loaner-bug1829730** | generic-worker | 30.0.2 | simple | 6fdba0dad3 | darwin | arm64 | 1.16.4 | 0 | 0 |
| **releng-hardware/gecko-t-osx-1100-m1-staging** | generic-worker | 55.3.4 | simple | a71f300fcb | darwin | arm64 | 1.21.1 | 0 | 0 |
| **releng-hardware/gecko-t-osx-1300-m2** | generic-worker | 55.3.4 | simple | a71f300fcb | darwin | arm64 | 1.21.1 | 0 | 0 |
| **releng-hardware/gecko-t-win-talos** | generic-worker | 16.2.0 | multiuser | b321a877c8 | windows | amd64 | 1.10.8 | 0 | 0 |
| **releng-hardware/gecko-t-win10-64-1803-hw** | generic-worker | 16.2.0 | multiuser | b321a877c8 | windows | amd64 | 1.10.8 | 0 | 0 |
| **releng-hardware/gecko-t-win10-64-dev** | generic-worker | 16.2.0 | multiuser | b321a877c8 | windows | amd64 | 1.10.8 | 0 | 0 |
| **releng-hardware/gecko-t-win10-64-ref-hw** | generic-worker | 16.2.0 | multiuser | b321a877c8 | windows | amd64 | 1.10.8 | 0 | 0 |
| **releng-hardware/gecko-t-win7-32-hw** | generic-worker | 45.0.0 | multiuser | 988e8100b3 | windows | 386 | 1.19.3 | 0 | 0 |
| **releng-hardware/mozillavpn-b-1-osx** | generic-worker | 54.4.1 | multiuser | a38a78d98c | darwin | amd64 | 1.20.6 | 0 | 0 |
| **releng-hardware/mozillavpn-b-3-osx** | generic-worker | 54.4.1 | multiuser | a38a78d98c | darwin | amd64 | 1.20.6 | 0 | 0 |
| **releng-hardware/win11-64-2009-hw-ref** | generic-worker | 54.4.1 | multiuser | a38a78d98c | windows | amd64 | 1.20.6 | 0 | 0 |
| **releng-hardware/win11-64-2009-hw-ref-alpha** | generic-worker | 54.4.1 | multiuser | a38a78d98c | windows | amd64 | 1.20.6 | 0 | 0 |
| **relops-3/win2019** | generic-worker | 16.5.1 | multiuser | 02bc4c9cd7 | windows | amd64 | 1.10.8 | 4 | 4 |
| **translations-1/b-linux-aerickson-test** | generic-worker | 57.0.1 | simple | fe45895655 | linux | amd64 | 1.21.3 | 2 | 2 |
| **translations-1/b-linux-v100-gpu** | generic-worker | 55.1.1 | simple | b83fffd4bd | linux | amd64 | 1.21.1 | 181 | 181 |
| **translations-1/b-linux-v100-gpu-4** | generic-worker | 55.1.1 | simple | b83fffd4bd | linux | amd64 | 1.21.1 | 50 | 50 |
| **translations-1/b-linux-v100-gpu-4-1tb** | generic-worker | 55.1.1 | simple | b83fffd4bd | linux | amd64 | 1.21.1 | 82 | 82 |
| **translations-1/b-linux-v100-gpu-4-300gb** | generic-worker | 55.1.1 | simple | b83fffd4bd | linux | amd64 | 1.21.1 | 12 | 12 |


## Docker Worker

Total: `131`

Count by version:

| Version | Count |
| :--- | ---: |
| 38.0.5 | 121 |
| 48.3.0 | 1 |
| unknown version | 9 |


Count by image:

| Version | Count |
| :--- | ---: |
| projects/taskcluster-imaging/global/images/docker-firefoxci-gcp-lt-googlecompute-2023-04-13t21-30-28z | 1 |
| projects/taskcluster-imaging/global/images/docker-worker-gcp-u14-04-2023-06-13 | 9 |
| projects/taskcluster-imaging/global/images/docker-firefoxci-gcp-l1-googlecompute-2023-06-13t23-47-26z | 74 |
| projects/fxci-production-level3-workers/global/images/docker-firefoxci-gcp-l3-googlecompute-2023-03-16t20-10-06z | 47 |


| Worker Pool | Implementation | Version | Total Workers | Total Capacity |
| --- | --- | --- | ---: | ---: |
| **adhoc-1/b-linux-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **adhoc-1/decision-gcp** | docker-worker | 38.0.5 | 5 | 10 |
| **adhoc-1/images-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **adhoc-3/b-linux-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **adhoc-3/decision-gcp** | docker-worker | 38.0.5 | 6 | 12 |
| **adhoc-3/images-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **app-services-1/b-linux-gcp** | docker-worker | 38.0.5 | 112 | 112 |
| **app-services-1/decision-gcp** | docker-worker | 38.0.5 | 577 | 1154 |
| **app-services-1/images-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **app-services-3/b-linux-gcp** | docker-worker | 38.0.5 | 701 | 701 |
| **app-services-3/decision-gcp** | docker-worker | 38.0.5 | 34 | 68 |
| **app-services-3/images-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **ci-1/decision-gcp** | docker-worker | 38.0.5 | 22 | 44 |
| **ci-1/images-gcp** | docker-worker | 38.0.5 | 13 | 13 |
| **ci-3/decision-gcp** | docker-worker | 38.0.5 | 14 | 28 |
| **ci-3/images-gcp** | docker-worker | 38.0.5 | 4 | 4 |
| **ci-t/linux-gcp** | docker-worker | 38.0.5 | 166 | 166 |
| **code-analysis-1/linux-gcp** | docker-worker | 38.0.5 | 16 | 16 |
| **code-analysis-3/linux-gcp** | docker-worker | 38.0.5 | 14 | 14 |
| **code-coverage/bot-gcp** | docker-worker | 38.0.5 | 825 | 825 |
| **code-review/bot-gcp** | docker-worker | 38.0.5 | 3378 | 3378 |
| **comm-1/b-linux-gcp** | docker-worker | 38.0.5 | 328 | 328 |
| **comm-1/b-linux-large-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **comm-1/b-linux-xlarge-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **comm-1/decision-gcp** | docker-worker | 38.0.5 | 134 | 268 |
| **comm-1/images-gcp** | docker-worker | 38.0.5 | 9 | 9 |
| **comm-2/b-linux-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **comm-2/b-linux-large-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **comm-2/b-linux-xlarge-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **comm-2/decision-gcp** | docker-worker | 38.0.5 | 2 | 4 |
| **comm-2/images-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **comm-3/b-linux-gcp** | docker-worker | 38.0.5 | 3094 | 3094 |
| **comm-3/b-linux-large-gcp** | docker-worker | 38.0.5 | 3 | 3 |
| **comm-3/b-linux-xlarge-gcp** | docker-worker | 38.0.5 | 3 | 3 |
| **comm-3/decision-gcp** | docker-worker | 38.0.5 | 88 | 176 |
| **comm-3/images-gcp** | docker-worker | 38.0.5 | 8 | 8 |
| **comm-t/misc-gcp** | docker-worker | 38.0.5 | 2 | 16 |
| **comm-t/t-linux-large-gcp** | docker-worker | unknown version | 1404 | 1404 |
| **comm-t/t-linux-xlarge-gcp** | docker-worker | unknown version | 243 | 243 |
| **comm-t/t-linux-xlarge-source-gcp** | docker-worker | unknown version | 1012 | 1012 |
| **gecko-1/b-linux-gcp** | docker-worker | 38.0.5 | 7116 | 7116 |
| **gecko-1/b-linux-gcp-bug1797804-n2** | docker-worker | 38.0.5 | 2 | 2 |
| **gecko-1/b-linux-large-gcp** | docker-worker | 38.0.5 | 376 | 376 |
| **gecko-1/b-linux-xlarge-gcp** | docker-worker | 38.0.5 | 789 | 789 |
| **gecko-1/b-linux-xlarge-gcp-bug1797804-c2** | docker-worker | 38.0.5 | 2 | 2 |
| **gecko-1/decision-gcp** | docker-worker | 38.0.5 | 1511 | 3022 |
| **gecko-1/images-gcp** | docker-worker | 38.0.5 | 85 | 85 |
| **gecko-2/b-linux-gcp** | docker-worker | 38.0.5 | 507 | 507 |
| **gecko-2/b-linux-large-gcp** | docker-worker | 38.0.5 | 13 | 13 |
| **gecko-2/b-linux-xlarge-gcp** | docker-worker | 38.0.5 | 16 | 16 |
| **gecko-2/decision-gcp** | docker-worker | 38.0.5 | 23 | 46 |
| **gecko-2/images-gcp** | docker-worker | 38.0.5 | 8 | 8 |
| **gecko-3/b-linux-gcp** | docker-worker | 38.0.5 | 18257 | 18257 |
| **gecko-3/b-linux-large-gcp** | docker-worker | 38.0.5 | 136 | 136 |
| **gecko-3/b-linux-xlarge-gcp** | docker-worker | 38.0.5 | 1154 | 1154 |
| **gecko-3/decision-gcp** | docker-worker | 38.0.5 | 1585 | 3170 |
| **gecko-3/images-gcp** | docker-worker | 38.0.5 | 26 | 26 |
| **gecko-3/t-linux-xlarge-gcp** | docker-worker | 38.0.5 | 303 | 303 |
| **gecko-t/misc-gcp** | docker-worker | 38.0.5 | 323 | 2584 |
| **gecko-t/t-linux-kvm-gcp** | docker-worker | unknown version | 27481 | 27481 |
| **gecko-t/t-linux-large-gcp** | docker-worker | unknown version | 100746 | 100746 |
| **gecko-t/t-linux-large-hostub2204-gcp** | docker-worker | 48.3.0 | 2 | 2 |
| **gecko-t/t-linux-xlarge-gcp** | docker-worker | unknown version | 56072 | 56072 |
| **gecko-t/t-linux-xlarge-source-gcp** | docker-worker | unknown version | 10491 | 10491 |
| **glean-1/b-linux-gcp** | docker-worker | 38.0.5 | 1133 | 1133 |
| **glean-1/decision-gcp** | docker-worker | 38.0.5 | 573 | 1146 |
| **glean-1/images-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **glean-3/b-linux-gcp** | docker-worker | 38.0.5 | 1133 | 1133 |
| **glean-3/decision-gcp** | docker-worker | 38.0.5 | 3 | 6 |
| **glean-3/images-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **infra/build-decision** | docker-worker | 38.0.5 | 12 | 96 |
| **mobile-1/b-linux-gcp** | docker-worker | 38.0.5 | 1735 | 1735 |
| **mobile-1/b-linux-large-gcp** | docker-worker | 38.0.5 | 7761 | 7761 |
| **mobile-1/b-linux-xlarge-gcp** | docker-worker | 38.0.5 | 4397 | 4397 |
| **mobile-1/bitrise-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **mobile-1/decision-gcp** | docker-worker | 38.0.5 | 496 | 992 |
| **mobile-1/images-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **mobile-3/b-linux-gcp** | docker-worker | 38.0.5 | 1485 | 1485 |
| **mobile-3/b-linux-large-gcp** | docker-worker | 38.0.5 | 8117 | 8117 |
| **mobile-3/b-linux-xlarge-gcp** | docker-worker | 38.0.5 | 4169 | 4169 |
| **mobile-3/bitrise-gcp** | docker-worker | 38.0.5 | 22 | 22 |
| **mobile-3/decision-gcp** | docker-worker | 38.0.5 | 256 | 512 |
| **mobile-3/images-gcp** | docker-worker | 38.0.5 | 3 | 3 |
| **mozilla-1/b-linux-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **mozilla-1/decision-gcp** | docker-worker | 38.0.5 | 566 | 1132 |
| **mozilla-1/images-gcp** | docker-worker | 38.0.5 | 18 | 18 |
| **mozilla-t/t-linux-large-gcp** | docker-worker | unknown version | 29 | 29 |
| **mozillaonline-1/b-linux-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **mozillaonline-3/b-linux-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **mozillavpn-1/b-linux-gcp** | docker-worker | 38.0.5 | 411 | 411 |
| **mozillavpn-1/b-linux-large-gcp** | docker-worker | 38.0.5 | 287 | 287 |
| **mozillavpn-1/decision-gcp** | docker-worker | 38.0.5 | 563 | 1126 |
| **mozillavpn-1/images-gcp** | docker-worker | 38.0.5 | 24 | 24 |
| **mozillavpn-3/b-linux-gcp** | docker-worker | 38.0.5 | 145 | 145 |
| **mozillavpn-3/b-linux-large-gcp** | docker-worker | 38.0.5 | 140 | 140 |
| **mozillavpn-3/decision-gcp** | docker-worker | 38.0.5 | 22 | 44 |
| **mozillavpn-3/images-gcp** | docker-worker | 38.0.5 | 7 | 7 |
| **nss-1/decision-gcp** | docker-worker | 38.0.5 | 2 | 4 |
| **nss-1/linux-gcp** | docker-worker | 38.0.5 | 832 | 832 |
| **nss-3/linux-gcp** | docker-worker | 38.0.5 | 104 | 104 |
| **nss-t/t-linux-xlarge-gcp** | docker-worker | unknown version | 975 | 975 |
| **project-relman/relman-svc** | docker-worker | 38.0.5 | 73 | 73 |
| **releng-1/b-linux-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **releng-1/decision-gcp** | docker-worker | 38.0.5 | 26 | 52 |
| **releng-1/linux-gcp** | docker-worker | 38.0.5 | 10 | 10 |
| **releng-3/b-linux-gcp** | docker-worker | 38.0.5 | 12 | 12 |
| **releng-3/decision-gcp** | docker-worker | 38.0.5 | 18 | 36 |
| **releng-3/linux-gcp** | docker-worker | 38.0.5 | 31 | 31 |
| **releng-t/linux-gcp** | docker-worker | 38.0.5 | 126 | 126 |
| **relops-3/decision-gcp** | docker-worker | 38.0.5 | 2 | 4 |
| **scriptworker-1/b-linux-gcp** | docker-worker | 38.0.5 | 77 | 77 |
| **scriptworker-1/decision-gcp** | docker-worker | 38.0.5 | 26 | 52 |
| **scriptworker-1/images-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **scriptworker-3/b-linux-gcp** | docker-worker | 38.0.5 | 72 | 72 |
| **scriptworker-3/decision-gcp** | docker-worker | 38.0.5 | 28 | 56 |
| **scriptworker-3/images-gcp** | docker-worker | 38.0.5 | 40 | 40 |
| **taskgraph-1/decision-gcp** | docker-worker | 38.0.5 | 9 | 18 |
| **taskgraph-1/images-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **taskgraph-3/decision-gcp** | docker-worker | 38.0.5 | 2 | 4 |
| **taskgraph-3/images-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **taskgraph-t/linux-gcp** | docker-worker | 38.0.5 | 14 | 14 |
| **translations-1/b-linux-large-gcp** | docker-worker | 38.0.5 | 147 | 147 |
| **translations-1/b-linux-large-gcp-300gb** | docker-worker | 38.0.5 | 57 | 57 |
| **translations-1/decision-gcp** | docker-worker | 38.0.5 | 21 | 42 |
| **translations-1/images-gcp** | docker-worker | 38.0.5 | 12 | 12 |
| **xpi-1/b-linux-gcp** | docker-worker | 38.0.5 | 41 | 41 |
| **xpi-1/decision-gcp** | docker-worker | 38.0.5 | 13 | 26 |
| **xpi-1/images-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **xpi-3/b-linux-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **xpi-3/decision-gcp** | docker-worker | 38.0.5 | 2 | 4 |
| **xpi-3/images-gcp** | docker-worker | 38.0.5 | 2 | 2 |


## Script Worker

Total: `26`



| Worker Pool | Implementation | Version | Total Workers | Total Capacity |
| --- | --- | --- | ---: | ---: |
| **scriptworker-k8s/adhoc-t-signing** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/comm-3-balrog** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/comm-3-beetmover** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/comm-3-bouncer** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/comm-3-signing** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/comm-3-tree** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/comm-t-signing** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/gecko-3-balrog** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/gecko-3-beetmover** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/gecko-3-bouncer** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/gecko-3-signing** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/gecko-3-tree** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/gecko-t-signing** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/mobile-3-beetmover** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/mobile-3-pushapk** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/mobile-3-signing** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/mobile-t-signing** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/mozillavpn-3-signing** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-prov-v1/adhoc-3-signing-mac** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-prov-v1/adhoc-t-signing-mac** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-prov-v1/depsigning-mac-v1** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-prov-v1/mozillavpn-3-signing-mac** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-prov-v1/signing-mac-v1** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-prov-v1/tb-depsigning-mac-v1** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-prov-v1/tb-signing-mac-v1** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-prov-v1/vpn-depsigning-mac-v1** | Scriptworker Chain of Trust | <no value> | 0 | 0 |


## No artifacts found [^1]

Total: `6`


Count by image:

| Version | Count |
| :--- | ---: |
|  | 4 |
| projects/taskcluster-imaging/global/images/generic-2204-wayland-vm-gcp-googlecompute-2023-09-22t17-39-37z | 1 |
| projects/taskcluster-imaging/global/images/generic-2204-wayland-vm-gcp-aerickson-test-2023-09-20t01-30-09z | 1 |


| Worker Pool | Implementation | Version | Total Workers | Total Capacity |
| --- | --- | --- | ---: | ---: |
| **built-in/fail** |  | No artifacts found | 0 | 0 |
| **built-in/succeed** |  | No artifacts found | 0 | 0 |
| **gecko-t/t-linux-vm-2204-wayland** |  | No artifacts found | 1985 | 1985 |
| **gecko-t/t-linux-vm-2204-wayland-aerickson-test** |  | No artifacts found | 4 | 4 |
| **scriptworker-prov-v1/mac-notarization-poller** |  | No artifacts found | 0 | 0 |
| **scriptworker-prov-v1/tb-mac-notarization-poller** |  | No artifacts found | 0 | 0 |


## Version not determined [^2]

Total: `25`


Count by image:

| Version | Count |
| :--- | ---: |
| /subscriptions/108d46d5-fe9b-4850-9a7d-8c914aa6c1f0/resourceGroups/rg-packer-through-cib/providers/Microsoft.Compute/images/win2022-64-2009-centralus-2022-datacenter-azure-edition-531c001 | 1 |
| /subscriptions/108d46d5-fe9b-4850-9a7d-8c914aa6c1f0/resourceGroups/rg-packer-through-cib/providers/Microsoft.Compute/images/win2012r2-64-vs-py2-l1-centralus-2012-R2-Datacenter-7aa76c6 | 1 |
| /subscriptions/108d46d5-fe9b-4850-9a7d-8c914aa6c1f0/resourceGroups/rg-packer-through-cib/providers/Microsoft.Compute/images/win11-64-2009-centralus-win11-22h2-avd-69a7ecf | 1 |
| /subscriptions/108d46d5-fe9b-4850-9a7d-8c914aa6c1f0/resourceGroups/rg-packer-through-cib/providers/Microsoft.Compute/images/win2012r2-64-l1-centralus-2012-R2-Datacenter-21c6826 | 4 |
| /subscriptions/108d46d5-fe9b-4850-9a7d-8c914aa6c1f0/resourceGroups/rg-packer-through-cib/providers/Microsoft.Compute/images/win2012r2-64-l1-centralus-2012-R2-Datacenter-alpha | 2 |
| /subscriptions/108d46d5-fe9b-4850-9a7d-8c914aa6c1f0/resourceGroups/rg-packer-through-cib/providers/Microsoft.Compute/images/win2022-64-2009-rd-centralus-2022-datacenter-azure-edition-alpha | 1 |
| /subscriptions/a30e97ab-734a-4f3b-a0e4-c51c0bff0701/resourceGroups/rg-packer-through-cib/providers/Microsoft.Compute/images/trusted-win2022-64-2009-centralus-2022-datacenter-azure-edition-d5dee60 | 1 |
|  | 2 |
| /subscriptions/108d46d5-fe9b-4850-9a7d-8c914aa6c1f0/resourceGroups/rg-packer-through-cib/providers/Microsoft.Compute/images/win2012r2-64-vs-py2-l1-centralus-2012-R2-Datacenter-alpha | 1 |
| /subscriptions/108d46d5-fe9b-4850-9a7d-8c914aa6c1f0/resourceGroups/rg-packer-through-cib/providers/Microsoft.Compute/images/win2012r2-64-vs-l1-centralus-2012-R2-Datacenter-21c6826 | 1 |
| /subscriptions/108d46d5-fe9b-4850-9a7d-8c914aa6c1f0/resourceGroups/rg-packer-through-cib/providers/Microsoft.Compute/images/win11-64-2009-tc-centralus-win11-22h2-avd-alpha | 1 |
| /subscriptions/108d46d5-fe9b-4850-9a7d-8c914aa6c1f0/resourceGroups/rg-packer-through-cib/providers/Microsoft.Compute/images/win10-64-2009-centralus-win10-22h2-avd-alpha | 1 |
| /subscriptions/108d46d5-fe9b-4850-9a7d-8c914aa6c1f0/resourceGroups/rg-packer-through-cib/providers/Microsoft.Compute/images/masterwayz-win2012test-image-20221122010013-alpha | 1 |
| /subscriptions/108d46d5-fe9b-4850-9a7d-8c914aa6c1f0/resourceGroups/rg-packer-through-cib/providers/Microsoft.Compute/images/win10-64-2009-eastus-win10-22h2-avd-alpha | 1 |
| /subscriptions/108d46d5-fe9b-4850-9a7d-8c914aa6c1f0/resourceGroups/rg-packer-through-cib/providers/Microsoft.Compute/images/win10-64-2009-ukwest-win10-22h2-avd-alpha | 1 |
| /subscriptions/108d46d5-fe9b-4850-9a7d-8c914aa6c1f0/resourceGroups/rg-packer-through-cib/providers/Microsoft.Compute/images/win11-64-2009-canadacentral-win11-22h2-avd-531c001 | 1 |
| /subscriptions/108d46d5-fe9b-4850-9a7d-8c914aa6c1f0/resourceGroups/rg-packer-through-cib/providers/Microsoft.Compute/images/win2022-64-2009-centralus-2022-datacenter-azure-edition-alpha | 2 |
| /subscriptions/108d46d5-fe9b-4850-9a7d-8c914aa6c1f0/resourceGroups/rg-packer-through-cib/providers/Microsoft.Compute/images/win2012r2-64-vs-l1-centralus-2012-R2-Datacenter-alpha | 1 |
| /subscriptions/a30e97ab-734a-4f3b-a0e4-c51c0bff0701/resourceGroups/rg-packer-through-cib/providers/Microsoft.Compute/images/trusted-win2012r2-64-l3-centralus-2012-R2-Datacenter-7cde253 | 1 |


| Worker Pool | Implementation | Version | Total Workers | Total Capacity |
| --- | --- | --- | ---: | ---: |
| **comm-1/b-win2012-azure** |  | Version not determined; task not (yet) claimed | 9 | 9 |
| **comm-2/b-win2012-azure** |  | Version not determined; task not (yet) claimed | 9 | 9 |
| **comm-t/win11-64-2009** |  | Version not determined; task not (yet) claimed | 3757 | 3757 |
| **gecko-1/b-win2012-azure** |  | Version not determined; task not (yet) claimed | 2 | 2 |
| **gecko-1/b-win2012-azure-alpha** |  | Version not determined; task not (yet) claimed | 9 | 9 |
| **gecko-1/b-win2012-azure-beta** |  | Version not determined; task not (yet) claimed | 9 | 9 |
| **gecko-1/b-win2022-alpha** |  | Version not determined; task not (yet) claimed | 9 | 9 |
| **gecko-1/b-win2022-rd-alpha** |  | Version not determined; task not (yet) claimed | 9 | 9 |
| **gecko-2/b-win2012-azure** |  | Version not determined; task not (yet) claimed | 9 | 9 |
| **gecko-3/b-win2012-azure** |  | Version not determined; task not (yet) claimed | 19 | 19 |
| **gecko-t/win10-64-2009-alpha** |  | Version not determined; task not (yet) claimed | 9 | 9 |
| **gecko-t/win10-64-2009-gpu-alpha** |  | Version not determined; task not (yet) claimed | 9 | 9 |
| **gecko-t/win10-64-2009-source-alpha** |  | Version not determined; task not (yet) claimed | 9 | 9 |
| **gecko-t/win11-64-2009-alpha-test** |  | Version not determined; task not (yet) claimed | 9 | 9 |
| **gecko-t/win11-64-2009-tc-alpha** |  | Version not determined; task not (yet) claimed | 9 | 9 |
| **mozillavpn-1/b-win2012-azure** |  | Version not determined; task not (yet) claimed | 148 | 148 |
| **mozillavpn-1/b-win2012-azure-alpha** |  | Version not determined; task not (yet) claimed | 9 | 9 |
| **mozillavpn-1/masterwayz-win2012-azure** |  | Version not determined; task not (yet) claimed | 9 | 9 |
| **nss-1/b-win2012-azure** |  | Version not determined; task not (yet) claimed | 2278 | 2278 |
| **nss-1/b-win2012-azure-alpha** |  | Version not determined; task not (yet) claimed | 9 | 9 |
| **nss-1/b-win2022-alpha** |  | Version not determined; task not (yet) claimed | 9 | 9 |
| **proj-autophone/gecko-t-bitbar-gw-perf-a51** |  | Version not determined; task not (yet) claimed | 0 | 0 |
| **releng-hardware/gecko-t-linux-talos-1804** |  | Version not determined; task not (yet) claimed | 0 | 0 |
| **relops-1/b-win2022** |  | Version not determined; task not (yet) claimed | 17 | 17 |
| **relops-3/b-win2022** |  | Version not determined; task not (yet) claimed | 17 | 17 |



[^1]: Those are the pools whose tasks were claimed and resolved by a worker as expected, but the worker did not publish either artifact `public/logs/live_backing.log` nor `public/logs/chain_of_trust.log`, which is the source used to identify the worker implementation.

[^2]: Probing task remains pending after two hours. Those are the pools that were not able to start any worker to claim the task within two hours.

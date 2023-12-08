

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
| 30.0.2 | 2 |
| 36.0.0 | 2 |
| 40.0.3 | 2 |
| 43.2.0 | 4 |
| 44.0.0 | 1 |
| 45.0.0 | 1 |
| 54.4.1 | 25 |
| 55.1.1 | 4 |
| 55.3.4 | 6 |
| 57.0.1 | 1 |


Count by image:

| Version | Count |
| :--- | ---: |
| /subscriptions/a30e97ab-734a-4f3b-a0e4-c51c0bff0701/resourceGroups/rg-packer-through-cib/providers/Microsoft.Compute/images/trusted-win2022-64-2009-centralus-2022-datacenter-azure-edition-28c1765 | 3 |
| /subscriptions/108d46d5-fe9b-4850-9a7d-8c914aa6c1f0/resourceGroups/rg-packer-through-cib/providers/Microsoft.Compute/images/win10-64-2009-canadacentral-win10-22h2-avd-531c001 | 2 |
| /subscriptions/a30e97ab-734a-4f3b-a0e4-c51c0bff0701/resourceGroups/rg-packer-through-cib/providers/Microsoft.Compute/images/trusted-win2012r2-64-vs-py2-l3-centralus-2012-R2-Datacenter-7aa76c6 | 1 |
| /subscriptions/108d46d5-fe9b-4850-9a7d-8c914aa6c1f0/resourceGroups/rg-packer-through-cib/providers/Microsoft.Compute/images/win11-64-2009-eastus-win11-22h2-avd-531c001 | 2 |
| /subscriptions/a30e97ab-734a-4f3b-a0e4-c51c0bff0701/resourceGroups/rg-packer-through-cib/providers/Microsoft.Compute/images/trusted-win2012r2-64-vs-l3-centralus-2012-R2-Datacenter-7cde253 | 1 |
| /subscriptions/108d46d5-fe9b-4850-9a7d-8c914aa6c1f0/resourceGroups/rg-packer-through-cib/providers/Microsoft.Compute/images/win10-64-2009-eastus-win10-22h2-avd-531c001 | 1 |
| projects/taskcluster-imaging/global/images/generic-translations-gcp-googlecompute-2023-11-01t20-12-18z | 1 |
| /subscriptions/108d46d5-fe9b-4850-9a7d-8c914aa6c1f0/resourceGroups/rg-packer-through-cib/providers/Microsoft.Compute/images/win11-64-2009-canadacentral-win11-22h2-avd-531c001 | 1 |
| /subscriptions/108d46d5-fe9b-4850-9a7d-8c914aa6c1f0/resourceGroups/rg-packer-through-cib/providers/Microsoft.Compute/images/win2022-64-2009-centralus-2022-datacenter-azure-edition-531c001 | 4 |
| /subscriptions/a30e97ab-734a-4f3b-a0e4-c51c0bff0701/resourceGroups/rg-packer-through-cib/providers/Microsoft.Compute/images/trusted-win2012r2-64-l3-centralus-2012-R2-Datacenter-7cde253 | 2 |
| ami-07db1c26bda143e33 | 1 |
| /subscriptions/108d46d5-fe9b-4850-9a7d-8c914aa6c1f0/resourceGroups/rg-packer-through-cib/providers/Microsoft.Compute/images/win11-64-2009-uksouth-win11-22h2-avd-alpha | 2 |
| /subscriptions/108d46d5-fe9b-4850-9a7d-8c914aa6c1f0/resourceGroups/rg-packer-through-cib/providers/Microsoft.Compute/images/win11-64-2009-centralus-win11-22h2-avd-531c001 | 3 |
| projects/taskcluster-imaging/global/images/generic-translations-gcp-googlecompute-2023-09-12t21-35-55z | 4 |
| /subscriptions/108d46d5-fe9b-4850-9a7d-8c914aa6c1f0/resourceGroups/rg-packer-through-cib/providers/Microsoft.Compute/images/win11-64-2009-eastus-win11-22h2-avd-alpha | 1 |
|  | 25 |
| unknown | 1 |


| Worker Pool | Implementation | Version | Engine | Revision | OS | Arch | GO | Total Workers | Total Capacity |
| --- | --- | --- | --- | --- | --- | --- | --- | ---: | ---: |
| **bitbar/gecko-t-win64-aarch64-laptop** | generic-worker | 14.1.2 | <UNKNOWN> | 13118c4c1b | windows | 386 | 1.10.8 | 0 | 0 |
| **comm-1/b-win2022** | generic-worker | 54.4.1 | multiuser | a38a78d98c | windows | amd64 | 1.20.6 | 5 | 5 |
| **comm-2/b-win2022** | generic-worker | 54.4.1 | multiuser | a38a78d98c | windows | amd64 | 1.20.6 | 4 | 4 |
| **comm-3/b-win2012-azure** | generic-worker | 43.2.0 | multiuser | 53fa7b79bc | windows | amd64 | 1.16.3 | 22 | 22 |
| **comm-3/b-win2022** | generic-worker | 54.4.1 | multiuser | a38a78d98c | windows | amd64 | 1.20.6 | 628 | 628 |
| **comm-t/win11-64-2009-ssd** | generic-worker | 54.4.1 | multiuser | a38a78d98c | windows | amd64 | 1.20.6 | 513 | 513 |
| **gecko-1/b-win2022** | generic-worker | 54.4.1 | multiuser | a38a78d98c | windows | amd64 | 1.20.6 | 780 | 780 |
| **gecko-2/b-win2022** | generic-worker | 54.4.1 | multiuser | a38a78d98c | windows | amd64 | 1.20.6 | 34 | 34 |
| **gecko-3/b-win2012-azure** | generic-worker | 43.2.0 | multiuser | 53fa7b79bc | windows | amd64 | 1.16.3 | 69 | 69 |
| **gecko-3/b-win2022** | generic-worker | 54.4.1 | multiuser | a38a78d98c | windows | amd64 | 1.20.6 | 5477 | 5477 |
| **gecko-t/win10-64-2009** | generic-worker | 54.4.1 | multiuser | a38a78d98c | windows | amd64 | 1.20.6 | 6299 | 6299 |
| **gecko-t/win10-64-2009-gpu** | generic-worker | 54.4.1 | multiuser | a38a78d98c | windows | amd64 | 1.20.6 | 791 | 791 |
| **gecko-t/win10-64-2009-source** | generic-worker | 54.4.1 | multiuser | a38a78d98c | windows | amd64 | 1.20.6 | 4 | 4 |
| **gecko-t/win11-64-2009** | generic-worker | 54.4.1 | multiuser | a38a78d98c | windows | amd64 | 1.20.6 | 20334 | 20334 |
| **gecko-t/win11-64-2009-alpha** | generic-worker | 54.4.1 | multiuser | a38a78d98c | windows | amd64 | 1.20.6 | 4 | 4 |
| **gecko-t/win11-64-2009-gpu** | generic-worker | 54.4.1 | multiuser | a38a78d98c | windows | amd64 | 1.20.6 | 21077 | 21077 |
| **gecko-t/win11-64-2009-gpu-alpha** | generic-worker | 54.4.1 | multiuser | a38a78d98c | windows | amd64 | 1.20.6 | 4 | 4 |
| **gecko-t/win11-64-2009-source** | generic-worker | 54.4.1 | multiuser | a38a78d98c | windows | amd64 | 1.20.6 | 1609 | 1609 |
| **gecko-t/win11-64-2009-source-alpha** | generic-worker | 54.4.1 | multiuser | a38a78d98c | windows | amd64 | 1.20.6 | 4 | 4 |
| **gecko-t/win11-64-2009-ssd** | generic-worker | 54.4.1 | multiuser | a38a78d98c | windows | amd64 | 1.20.6 | 1759 | 1759 |
| **gecko-t/win11-64-2009-ssd-gpu** | generic-worker | 54.4.1 | multiuser | a38a78d98c | windows | amd64 | 1.20.6 | 1497 | 1497 |
| **localprovisioner/nss-macos-10-12** | generic-worker | 15.1.4 | simple | c407e45e3f | darwin | amd64 | 1.10.8 | 0 | 0 |
| **localprovisioner/nss-macos-m1** | generic-worker | 30.0.2 | simple | 6fdba0dad3 | darwin | arm64 | 1.16.4 | 0 | 0 |
| **mozillavpn-3/b-win2012-azure** | generic-worker | 43.2.0 | multiuser | 53fa7b79bc | windows | amd64 | 1.16.3 | 4 | 4 |
| **mozillavpn-3/b-win2022** | generic-worker | 54.4.1 | multiuser | a38a78d98c | windows | amd64 | 1.20.6 | 107 | 107 |
| **nss-3/b-win2012-azure** | generic-worker | 43.2.0 | multiuser | 53fa7b79bc | windows | amd64 | 1.16.3 | 30 | 30 |
| **performance-hardware/gecko-t-fxrecorder** | generic-worker | 44.0.0 | multiuser | faf6f319c0 | windows | amd64 | 1.16.3 | 0 | 0 |
| **proj-autophone/gecko-t-bitbar-gw-perf-a51** | generic-worker | 36.0.0 | simple | b9cb11293f | linux | amd64 | 1.13.7 | 0 | 0 |
| **proj-autophone/gecko-t-bitbar-gw-unit-p5** | generic-worker | 36.0.0 | simple | b9cb11293f | linux | amd64 | 1.13.7 | 0 | 0 |
| **releng-hardware/applicationservices-b-1-osx1015** | generic-worker | 55.3.4 | multiuser | a71f300fcb | darwin | amd64 | 1.21.1 | 0 | 0 |
| **releng-hardware/applicationservices-b-3-osx1015** | generic-worker | 54.4.1 | multiuser | a38a78d98c | darwin | amd64 | 1.20.6 | 0 | 0 |
| **releng-hardware/gecko-1-b-osx-1015-staging** | generic-worker | 54.4.1 | multiuser | a38a78d98c | darwin | amd64 | 1.20.6 | 0 | 0 |
| **releng-hardware/gecko-3-b-osx-1015** | generic-worker | 55.3.4 | multiuser | a71f300fcb | darwin | amd64 | 1.21.1 | 0 | 0 |
| **releng-hardware/gecko-t-osx-1015-power** | generic-worker | 40.0.3 | simple | 5f7a31ac6e | darwin | amd64 | 1.15.6 | 0 | 0 |
| **releng-hardware/gecko-t-osx-1015-r8-staging** | generic-worker | 40.0.3 | simple | 5f7a31ac6e | darwin | amd64 | 1.15.6 | 0 | 0 |
| **releng-hardware/gecko-t-osx-1100-m1** | generic-worker | 55.3.4 | simple | a71f300fcb | darwin | arm64 | 1.21.1 | 0 | 0 |
| **releng-hardware/gecko-t-osx-1100-m1-loaner** | generic-worker | 30.0.2 | simple | 6fdba0dad3 | darwin | arm64 | 1.16.4 | 0 | 0 |
| **releng-hardware/gecko-t-osx-1100-m1-staging** | generic-worker | 55.3.4 | simple | a71f300fcb | darwin | arm64 | 1.21.1 | 0 | 0 |
| **releng-hardware/gecko-t-osx-1300-m2** | generic-worker | 55.3.4 | simple | a71f300fcb | darwin | arm64 | 1.21.1 | 0 | 0 |
| **releng-hardware/gecko-t-win-talos** | generic-worker | 16.2.0 | multiuser | b321a877c8 | windows | amd64 | 1.10.8 | 0 | 0 |
| **releng-hardware/gecko-t-win10-64-1803-hw** | generic-worker | 16.2.0 | multiuser | b321a877c8 | windows | amd64 | 1.10.8 | 0 | 0 |
| **releng-hardware/gecko-t-win10-64-dev** | generic-worker | 16.2.0 | multiuser | b321a877c8 | windows | amd64 | 1.10.8 | 0 | 0 |
| **releng-hardware/gecko-t-win10-64-ref-hw** | generic-worker | 16.2.0 | multiuser | b321a877c8 | windows | amd64 | 1.10.8 | 0 | 0 |
| **releng-hardware/gecko-t-win7-32-hw** | generic-worker | 45.0.0 | multiuser | 988e8100b3 | windows | 386 | 1.19.3 | 0 | 0 |
| **releng-hardware/mozillavpn-b-1-osx** | generic-worker | 54.4.1 | multiuser | a38a78d98c | darwin | amd64 | 1.20.6 | 0 | 0 |
| **releng-hardware/mozillavpn-b-3-osx** | generic-worker | 54.4.1 | multiuser | a38a78d98c | darwin | amd64 | 1.20.6 | 0 | 0 |
| **releng-hardware/nss-1-b-osx-1015** | generic-worker | 55.3.4 | multiuser | a71f300fcb | darwin | amd64 | 1.21.1 | 0 | 0 |
| **releng-hardware/win11-64-2009-hw-ref** | generic-worker | 54.4.1 | multiuser | a38a78d98c | windows | amd64 | 1.20.6 | 0 | 0 |
| **releng-hardware/win11-64-2009-hw-ref-alpha** | generic-worker | 54.4.1 | multiuser | a38a78d98c | windows | amd64 | 1.20.6 | 0 | 0 |
| **relops-3/win2019** | generic-worker | 16.5.1 | multiuser | 02bc4c9cd7 | windows | amd64 | 1.10.8 | 19 | 19 |
| **translations-1/b-linux-aerickson-test** | generic-worker | 57.0.1 | simple | fe45895655 | linux | amd64 | 1.21.3 | 2 | 2 |
| **translations-1/b-linux-v100-gpu** | generic-worker | 55.1.1 | simple | b83fffd4bd | linux | amd64 | 1.21.1 | 94 | 94 |
| **translations-1/b-linux-v100-gpu-4** | generic-worker | 55.1.1 | simple | b83fffd4bd | linux | amd64 | 1.21.1 | 82 | 82 |
| **translations-1/b-linux-v100-gpu-4-1tb** | generic-worker | 55.1.1 | simple | b83fffd4bd | linux | amd64 | 1.21.1 | 50 | 50 |
| **translations-1/b-linux-v100-gpu-4-300gb** | generic-worker | 55.1.1 | simple | b83fffd4bd | linux | amd64 | 1.21.1 | 28 | 28 |


## Docker Worker

Total: `131`

Count by version:

| Version | Count |
| :--- | ---: |
| 38.0.5 | 122 |
| 48.3.0 | 1 |
| unknown version | 8 |


Count by image:

| Version | Count |
| :--- | ---: |
| projects/taskcluster-imaging/global/images/docker-firefoxci-gcp-l1-googlecompute-2023-06-13t23-47-26z | 74 |
| projects/fxci-production-level3-workers/global/images/docker-firefoxci-gcp-l3-googlecompute-2023-12-06t15-04-40z | 47 |
| projects/taskcluster-imaging/global/images/docker-worker-gcp-u14-04-2023-06-13 | 8 |
| projects/taskcluster-imaging/global/images/docker-firefoxci-gcp-lt-googlecompute-2023-04-13t21-30-28z | 1 |
| projects/taskcluster-imaging/global/images/docker-firefoxci-gcp-l1-googlecompute-2023-12-01t19-56-02z | 1 |


| Worker Pool | Implementation | Version | Total Workers | Total Capacity |
| --- | --- | --- | ---: | ---: |
| **adhoc-1/b-linux-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **adhoc-1/decision-gcp** | docker-worker | 38.0.5 | 24 | 48 |
| **adhoc-1/images-gcp** | docker-worker | 38.0.5 | 12 | 12 |
| **adhoc-3/b-linux-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **adhoc-3/decision-gcp** | docker-worker | 38.0.5 | 9 | 18 |
| **adhoc-3/images-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **app-services-1/b-linux-gcp** | docker-worker | 38.0.5 | 160 | 160 |
| **app-services-1/decision-gcp** | docker-worker | 38.0.5 | 583 | 1166 |
| **app-services-1/images-gcp** | docker-worker | 38.0.5 | 12 | 12 |
| **app-services-3/b-linux-gcp** | docker-worker | 38.0.5 | 543 | 543 |
| **app-services-3/decision-gcp** | docker-worker | 38.0.5 | 24 | 48 |
| **app-services-3/images-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **ci-1/decision-gcp** | docker-worker | 38.0.5 | 15 | 30 |
| **ci-1/images-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **ci-3/decision-gcp** | docker-worker | 38.0.5 | 11 | 22 |
| **ci-3/images-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **ci-t/linux-gcp** | docker-worker | 38.0.5 | 126 | 126 |
| **code-analysis-1/linux-gcp** | docker-worker | 38.0.5 | 14 | 14 |
| **code-analysis-3/linux-gcp** | docker-worker | 38.0.5 | 21 | 21 |
| **code-coverage/bot-gcp** | docker-worker | 38.0.5 | 822 | 822 |
| **code-review/bot-gcp** | docker-worker | 38.0.5 | 3095 | 3095 |
| **comm-1/b-linux-gcp** | docker-worker | 38.0.5 | 506 | 506 |
| **comm-1/b-linux-large-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **comm-1/b-linux-xlarge-gcp** | docker-worker | 38.0.5 | 6 | 6 |
| **comm-1/decision-gcp** | docker-worker | 38.0.5 | 129 | 258 |
| **comm-1/images-gcp** | docker-worker | 38.0.5 | 6 | 6 |
| **comm-2/b-linux-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **comm-2/b-linux-large-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **comm-2/b-linux-xlarge-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **comm-2/decision-gcp** | docker-worker | 38.0.5 | 2 | 4 |
| **comm-2/images-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **comm-3/b-linux-gcp** | docker-worker | 38.0.5 | 6476 | 6476 |
| **comm-3/b-linux-large-gcp** | docker-worker | 38.0.5 | 18 | 18 |
| **comm-3/b-linux-xlarge-gcp** | docker-worker | 38.0.5 | 13 | 13 |
| **comm-3/decision-gcp** | docker-worker | 38.0.5 | 84 | 168 |
| **comm-3/images-gcp** | docker-worker | 38.0.5 | 6 | 6 |
| **comm-t/misc-gcp** | docker-worker | 38.0.5 | 2 | 16 |
| **comm-t/t-linux-large-gcp** | docker-worker | unknown version | 1344 | 1344 |
| **comm-t/t-linux-xlarge-gcp** | docker-worker | unknown version | 220 | 220 |
| **comm-t/t-linux-xlarge-source-gcp** | docker-worker | unknown version | 1028 | 1028 |
| **gecko-1/b-linux-gcp** | docker-worker | 38.0.5 | 18049 | 18049 |
| **gecko-1/b-linux-gcp-bug1797804-n2** | docker-worker | 38.0.5 | 2 | 2 |
| **gecko-1/b-linux-gcp-relops757-test** | docker-worker | 38.0.5 | 86 | 86 |
| **gecko-1/b-linux-large-gcp** | docker-worker | 38.0.5 | 1195 | 1195 |
| **gecko-1/b-linux-xlarge-gcp** | docker-worker | 38.0.5 | 1296 | 1296 |
| **gecko-1/b-linux-xlarge-gcp-bug1797804-c2** | docker-worker | 38.0.5 | 3 | 3 |
| **gecko-1/decision-gcp** | docker-worker | 38.0.5 | 1322 | 2644 |
| **gecko-1/images-gcp** | docker-worker | 38.0.5 | 39 | 39 |
| **gecko-2/b-linux-gcp** | docker-worker | 38.0.5 | 2449 | 2449 |
| **gecko-2/b-linux-large-gcp** | docker-worker | 38.0.5 | 15 | 15 |
| **gecko-2/b-linux-xlarge-gcp** | docker-worker | 38.0.5 | 13 | 13 |
| **gecko-2/decision-gcp** | docker-worker | 38.0.5 | 26 | 52 |
| **gecko-2/images-gcp** | docker-worker | 38.0.5 | 7 | 7 |
| **gecko-3/b-linux-gcp** | docker-worker | 38.0.5 | 36410 | 36410 |
| **gecko-3/b-linux-large-gcp** | docker-worker | 38.0.5 | 262 | 262 |
| **gecko-3/b-linux-xlarge-gcp** | docker-worker | 38.0.5 | 1171 | 1171 |
| **gecko-3/decision-gcp** | docker-worker | 38.0.5 | 1387 | 2774 |
| **gecko-3/images-gcp** | docker-worker | 38.0.5 | 79 | 79 |
| **gecko-3/t-linux-xlarge-gcp** | docker-worker | 38.0.5 | 294 | 294 |
| **gecko-t/misc-gcp** | docker-worker | 38.0.5 | 397 | 3176 |
| **gecko-t/t-linux-large-gcp** | docker-worker | unknown version | 82975 | 82975 |
| **gecko-t/t-linux-large-hostub2204-gcp** | docker-worker | 48.3.0 | 2 | 2 |
| **gecko-t/t-linux-xlarge-gcp** | docker-worker | unknown version | 45254 | 45254 |
| **gecko-t/t-linux-xlarge-source-gcp** | docker-worker | unknown version | 10415 | 10415 |
| **glean-1/b-linux-gcp** | docker-worker | 38.0.5 | 1154 | 1154 |
| **glean-1/decision-gcp** | docker-worker | 38.0.5 | 578 | 1156 |
| **glean-1/images-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **glean-3/b-linux-gcp** | docker-worker | 38.0.5 | 1170 | 1170 |
| **glean-3/decision-gcp** | docker-worker | 38.0.5 | 7 | 14 |
| **glean-3/images-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **infra/build-decision** | docker-worker | 38.0.5 | 14 | 112 |
| **mobile-1/b-linux-gcp** | docker-worker | 38.0.5 | 2032 | 2032 |
| **mobile-1/b-linux-large-gcp** | docker-worker | 38.0.5 | 7620 | 7620 |
| **mobile-1/b-linux-xlarge-gcp** | docker-worker | 38.0.5 | 3377 | 3377 |
| **mobile-1/bitrise-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **mobile-1/decision-gcp** | docker-worker | 38.0.5 | 461 | 922 |
| **mobile-1/images-gcp** | docker-worker | 38.0.5 | 11 | 11 |
| **mobile-3/b-linux-gcp** | docker-worker | 38.0.5 | 1991 | 1991 |
| **mobile-3/b-linux-large-gcp** | docker-worker | 38.0.5 | 8284 | 8284 |
| **mobile-3/b-linux-xlarge-gcp** | docker-worker | 38.0.5 | 3713 | 3713 |
| **mobile-3/bitrise-gcp** | docker-worker | 38.0.5 | 19 | 19 |
| **mobile-3/decision-gcp** | docker-worker | 38.0.5 | 244 | 488 |
| **mobile-3/images-gcp** | docker-worker | 38.0.5 | 3 | 3 |
| **mozilla-1/b-linux-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **mozilla-1/decision-gcp** | docker-worker | 38.0.5 | 577 | 1154 |
| **mozilla-1/images-gcp** | docker-worker | 38.0.5 | 32 | 32 |
| **mozilla-t/t-linux-large-gcp** | docker-worker | unknown version | 51 | 51 |
| **mozillaonline-1/b-linux-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **mozillaonline-3/b-linux-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **mozillavpn-1/b-linux-gcp** | docker-worker | 38.0.5 | 478 | 478 |
| **mozillavpn-1/b-linux-large-gcp** | docker-worker | 38.0.5 | 342 | 342 |
| **mozillavpn-1/decision-gcp** | docker-worker | 38.0.5 | 556 | 1112 |
| **mozillavpn-1/images-gcp** | docker-worker | 38.0.5 | 19 | 19 |
| **mozillavpn-3/b-linux-gcp** | docker-worker | 38.0.5 | 214 | 214 |
| **mozillavpn-3/b-linux-large-gcp** | docker-worker | 38.0.5 | 200 | 200 |
| **mozillavpn-3/decision-gcp** | docker-worker | 38.0.5 | 31 | 62 |
| **mozillavpn-3/images-gcp** | docker-worker | 38.0.5 | 3 | 3 |
| **nss-1/decision-gcp** | docker-worker | 38.0.5 | 2 | 4 |
| **nss-1/linux-gcp** | docker-worker | 38.0.5 | 466 | 466 |
| **nss-3/linux-gcp** | docker-worker | 38.0.5 | 97 | 97 |
| **nss-t/t-linux-xlarge-gcp** | docker-worker | unknown version | 545 | 545 |
| **project-relman/relman-svc** | docker-worker | 38.0.5 | 72 | 72 |
| **releng-1/b-linux-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **releng-1/decision-gcp** | docker-worker | 38.0.5 | 29 | 58 |
| **releng-1/linux-gcp** | docker-worker | 38.0.5 | 47 | 47 |
| **releng-3/b-linux-gcp** | docker-worker | 38.0.5 | 8 | 8 |
| **releng-3/decision-gcp** | docker-worker | 38.0.5 | 21 | 42 |
| **releng-3/linux-gcp** | docker-worker | 38.0.5 | 48 | 48 |
| **releng-t/linux-gcp** | docker-worker | 38.0.5 | 209 | 209 |
| **relops-3/decision-gcp** | docker-worker | 38.0.5 | 5 | 10 |
| **scriptworker-1/b-linux-gcp** | docker-worker | 38.0.5 | 257 | 257 |
| **scriptworker-1/decision-gcp** | docker-worker | 38.0.5 | 21 | 42 |
| **scriptworker-1/images-gcp** | docker-worker | 38.0.5 | 44 | 44 |
| **scriptworker-3/b-linux-gcp** | docker-worker | 38.0.5 | 82 | 82 |
| **scriptworker-3/decision-gcp** | docker-worker | 38.0.5 | 24 | 48 |
| **scriptworker-3/images-gcp** | docker-worker | 38.0.5 | 41 | 41 |
| **taskgraph-1/decision-gcp** | docker-worker | 38.0.5 | 24 | 48 |
| **taskgraph-1/images-gcp** | docker-worker | 38.0.5 | 30 | 30 |
| **taskgraph-3/decision-gcp** | docker-worker | 38.0.5 | 15 | 30 |
| **taskgraph-3/images-gcp** | docker-worker | 38.0.5 | 11 | 11 |
| **taskgraph-t/linux-gcp** | docker-worker | 38.0.5 | 99 | 99 |
| **translations-1/b-linux-large-gcp** | docker-worker | 38.0.5 | 231 | 231 |
| **translations-1/b-linux-large-gcp-300gb** | docker-worker | 38.0.5 | 118 | 118 |
| **translations-1/decision-gcp** | docker-worker | 38.0.5 | 32 | 64 |
| **translations-1/images-gcp** | docker-worker | 38.0.5 | 21 | 21 |
| **xpi-1/b-linux-gcp** | docker-worker | 38.0.5 | 49 | 49 |
| **xpi-1/decision-gcp** | docker-worker | 38.0.5 | 19 | 38 |
| **xpi-1/images-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **xpi-3/b-linux-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **xpi-3/decision-gcp** | docker-worker | 38.0.5 | 2 | 4 |
| **xpi-3/images-gcp** | docker-worker | 38.0.5 | 2 | 2 |


## Script Worker

Total: `33`



| Worker Pool | Implementation | Version | Total Workers | Total Capacity |
| --- | --- | --- | ---: | ---: |
| **scriptworker-k8s/adhoc-t-signing-dev** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/comm-3-balrog** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/comm-3-beetmover** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/comm-3-bouncer** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/comm-3-signing** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/comm-3-tree** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/comm-t-signing** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/gecko-1-addon** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/gecko-1-balrog** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/gecko-1-beetmover** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/gecko-1-bouncer** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/gecko-1-pushflatpak** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/gecko-1-pushmsix** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/gecko-1-tree** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
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
| projects/taskcluster-imaging/global/images/generic-2204-wayland-vm-gcp-aerickson-test-2023-09-20t01-30-09z | 1 |
| projects/taskcluster-imaging/global/images/generic-2204-wayland-vm-gcp-googlecompute-2023-09-22t17-39-37z | 1 |


| Worker Pool | Implementation | Version | Total Workers | Total Capacity |
| --- | --- | --- | ---: | ---: |
| **built-in/fail** |  | No artifacts found | 0 | 0 |
| **built-in/succeed** |  | No artifacts found | 0 | 0 |
| **gecko-t/t-linux-vm-2204-wayland** |  | No artifacts found | 1597 | 1597 |
| **gecko-t/t-linux-vm-2204-wayland-aerickson-test** |  | No artifacts found | 5 | 5 |
| **scriptworker-prov-v1/mac-notarization-poller** |  | No artifacts found | 0 | 0 |
| **scriptworker-prov-v1/tb-mac-notarization-poller** |  | No artifacts found | 0 | 0 |


## Version not determined [^2]

Total: `28`


Count by image:

| Version | Count |
| :--- | ---: |
| /subscriptions/108d46d5-fe9b-4850-9a7d-8c914aa6c1f0/resourceGroups/rg-packer-through-cib/providers/Microsoft.Compute/images/win2012r2-64-l1-centralus-2012-R2-Datacenter-21c6826 | 4 |
| /subscriptions/108d46d5-fe9b-4850-9a7d-8c914aa6c1f0/resourceGroups/rg-packer-through-cib/providers/Microsoft.Compute/images/win2022-64-2009-rd-centralus-2022-datacenter-azure-edition-alpha | 1 |
| /subscriptions/108d46d5-fe9b-4850-9a7d-8c914aa6c1f0/resourceGroups/rg-packer-through-cib/providers/Microsoft.Compute/images/masterwayz-win2012test-image-20221122010013-alpha | 1 |
| /subscriptions/108d46d5-fe9b-4850-9a7d-8c914aa6c1f0/resourceGroups/rg-packer-through-cib/providers/Microsoft.Compute/images/win2012r2-64-vs-l1-centralus-2012-R2-Datacenter-21c6826 | 1 |
| /subscriptions/108d46d5-fe9b-4850-9a7d-8c914aa6c1f0/resourceGroups/rg-packer-through-cib/providers/Microsoft.Compute/images/win10-64-2009-eastus-win10-22h2-avd-alpha | 1 |
| /subscriptions/108d46d5-fe9b-4850-9a7d-8c914aa6c1f0/resourceGroups/rg-packer-through-cib/providers/Microsoft.Compute/images/win10-64-2009-ukwest-win10-22h2-avd-alpha | 1 |
| projects/taskcluster-imaging/global/images/docker-worker-gcp-u14-04-2023-06-13 | 1 |
| /subscriptions/108d46d5-fe9b-4850-9a7d-8c914aa6c1f0/resourceGroups/rg-packer-through-cib/providers/Microsoft.Compute/images/win2022-64-2009-centralus-2022-datacenter-azure-edition-alpha | 2 |
| /subscriptions/108d46d5-fe9b-4850-9a7d-8c914aa6c1f0/resourceGroups/rg-packer-through-cib/providers/Microsoft.Compute/images/win11-64-2009-tc-centralus-win11-22h2-avd-alpha | 1 |
| /subscriptions/108d46d5-fe9b-4850-9a7d-8c914aa6c1f0/resourceGroups/rg-packer-through-cib/providers/Microsoft.Compute/images/win10-64-2009-centralus-win10-22h2-avd-alpha | 1 |
| /subscriptions/108d46d5-fe9b-4850-9a7d-8c914aa6c1f0/resourceGroups/rg-packer-through-cib/providers/Microsoft.Compute/images/win11-64-2009-canadacentral-win11-22h2-avd-531c001 | 1 |
| /subscriptions/108d46d5-fe9b-4850-9a7d-8c914aa6c1f0/resourceGroups/rg-packer-through-cib/providers/Microsoft.Compute/images/win2012r2-64-vs-l1-centralus-2012-R2-Datacenter-alpha | 1 |
| /subscriptions/108d46d5-fe9b-4850-9a7d-8c914aa6c1f0/resourceGroups/rg-packer-through-cib/providers/Microsoft.Compute/images/win2022-64-2009-centralus-2022-datacenter-azure-edition-531c001 | 2 |
|  | 4 |
| /subscriptions/108d46d5-fe9b-4850-9a7d-8c914aa6c1f0/resourceGroups/rg-packer-through-cib/providers/Microsoft.Compute/images/win11-64-2009-centralus-win11-22h2-avd-69a7ecf | 1 |
| /subscriptions/108d46d5-fe9b-4850-9a7d-8c914aa6c1f0/resourceGroups/rg-packer-through-cib/providers/Microsoft.Compute/images/win2012r2-64-l1-centralus-2012-R2-Datacenter-alpha | 2 |
| /subscriptions/a30e97ab-734a-4f3b-a0e4-c51c0bff0701/resourceGroups/rg-packer-through-cib/providers/Microsoft.Compute/images/trusted-win2022-64-2009-centralus-2022-datacenter-azure-edition-28c1765 | 1 |
| /subscriptions/108d46d5-fe9b-4850-9a7d-8c914aa6c1f0/resourceGroups/rg-packer-through-cib/providers/Microsoft.Compute/images/win2012r2-64-vs-py2-l1-centralus-2012-R2-Datacenter-alpha | 1 |
| /subscriptions/108d46d5-fe9b-4850-9a7d-8c914aa6c1f0/resourceGroups/rg-packer-through-cib/providers/Microsoft.Compute/images/win2012r2-64-vs-py2-l1-centralus-2012-R2-Datacenter-7aa76c6 | 1 |


| Worker Pool | Implementation | Version | Total Workers | Total Capacity |
| --- | --- | --- | ---: | ---: |
| **comm-1/b-win2012-azure** |  | Version not determined; task not (yet) claimed | 11 | 11 |
| **comm-2/b-win2012-azure** |  | Version not determined; task not (yet) claimed | 11 | 11 |
| **comm-t/win11-64-2009** |  | Version not determined; task not (yet) claimed | 3699 | 3699 |
| **gecko-1/b-win2012-azure** |  | Version not determined; task not (yet) claimed | 2 | 2 |
| **gecko-1/b-win2012-azure-alpha** |  | Version not determined; task not (yet) claimed | 11 | 11 |
| **gecko-1/b-win2012-azure-beta** |  | Version not determined; task not (yet) claimed | 11 | 11 |
| **gecko-1/b-win2022-alpha** |  | Version not determined; task not (yet) claimed | 11 | 11 |
| **gecko-1/b-win2022-rd-alpha** |  | Version not determined; task not (yet) claimed | 11 | 11 |
| **gecko-2/b-win2012-azure** |  | Version not determined; task not (yet) claimed | 11 | 11 |
| **gecko-t/t-linux-kvm-gcp** |  | Version not determined; task not (yet) claimed | 30349 | 30349 |
| **gecko-t/win10-64-2009-alpha** |  | Version not determined; task not (yet) claimed | 11 | 11 |
| **gecko-t/win10-64-2009-gpu-alpha** |  | Version not determined; task not (yet) claimed | 11 | 11 |
| **gecko-t/win10-64-2009-source-alpha** |  | Version not determined; task not (yet) claimed | 11 | 11 |
| **gecko-t/win11-64-2009-alpha-test** |  | Version not determined; task not (yet) claimed | 11 | 11 |
| **gecko-t/win11-64-2009-tc-alpha** |  | Version not determined; task not (yet) claimed | 11 | 11 |
| **mozillavpn-1/b-win2012-azure** |  | Version not determined; task not (yet) claimed | 20 | 20 |
| **mozillavpn-1/b-win2012-azure-alpha** |  | Version not determined; task not (yet) claimed | 11 | 11 |
| **mozillavpn-1/b-win2022** |  | Version not determined; task not (yet) claimed | 239 | 239 |
| **mozillavpn-1/masterwayz-win2012-azure** |  | Version not determined; task not (yet) claimed | 11 | 11 |
| **nss-1/b-win2012-azure** |  | Version not determined; task not (yet) claimed | 1430 | 1430 |
| **nss-1/b-win2012-azure-alpha** |  | Version not determined; task not (yet) claimed | 12 | 12 |
| **nss-1/b-win2022-alpha** |  | Version not determined; task not (yet) claimed | 11 | 11 |
| **releng-hardware/gecko-1-b-osx-1015** |  | Version not determined; task not (yet) claimed | 0 | 0 |
| **releng-hardware/gecko-t-linux-talos-1804** |  | Version not determined; task not (yet) claimed | 0 | 0 |
| **releng-hardware/gecko-t-osx-1015-r8** |  | Version not determined; task not (yet) claimed | 0 | 0 |
| **releng-hardware/gecko-t-osx-1100-m1-loaner-bug1829730** |  | Version not determined; task not (yet) claimed | 0 | 0 |
| **relops-1/b-win2022** |  | Version not determined; task not (yet) claimed | 24 | 24 |
| **relops-3/b-win2022** |  | Version not determined; task not (yet) claimed | 23 | 23 |



[^1]: Those are the pools whose tasks were claimed and resolved by a worker as expected, but the worker did not publish either artifact `public/logs/live_backing.log` nor `public/logs/chain_of_trust.log`, which is the source used to identify the worker implementation.

[^2]: Probing task remains pending after two hours. Those are the pools that were not able to start any worker to claim the task within two hours.

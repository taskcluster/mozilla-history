

# Worker Pool Versions


## Generic Worker

Total: `56`

Count by version:

| Version | Count |
| :--- | ---: |
| 14.1.2 | 1 |
| 15.1.4 | 1 |
| 16.2.0 | 4 |
| 16.5.1 | 1 |
| 30.0.2 | 3 |
| 36.0.0 | 2 |
| 40.0.3 | 2 |
| 43.2.0 | 3 |
| 44.0.0 | 1 |
| 44.23.2 | 1 |
| 45.0.0 | 1 |
| 54.4.1 | 26 |
| 55.1.1 | 4 |
| 55.3.4 | 5 |
| 57.0.1 | 1 |


Count by image:

| Version | Count |
| :--- | ---: |
| /subscriptions/a30e97ab-734a-4f3b-a0e4-c51c0bff0701/resourceGroups/rg-packer-through-cib/providers/Microsoft.Compute/images/trusted-win2022-64-2009-centralus-2022-datacenter-azure-edition-d5dee60 | 3 |
| ami-07db1c26bda143e33 | 1 |
| /subscriptions/108d46d5-fe9b-4850-9a7d-8c914aa6c1f0/resourceGroups/rg-packer-through-cib/providers/Microsoft.Compute/images/win11-64-2009-eastus-win11-22h2-avd-531c001 | 2 |
| /subscriptions/a30e97ab-734a-4f3b-a0e4-c51c0bff0701/resourceGroups/rg-packer-through-cib/providers/Microsoft.Compute/images/trusted-win2012r2-64-vs-l3-centralus-2012-R2-Datacenter-7cde253 | 1 |
| /subscriptions/108d46d5-fe9b-4850-9a7d-8c914aa6c1f0/resourceGroups/rg-packer-through-cib/providers/Microsoft.Compute/images/win11-64-2009-centralus-win11-22h2-avd-531c001 | 2 |
| projects/taskcluster-imaging/global/images/generic-translations-gcp-googlecompute-2023-09-12t21-35-55z | 4 |
| /subscriptions/108d46d5-fe9b-4850-9a7d-8c914aa6c1f0/resourceGroups/rg-packer-through-cib/providers/Microsoft.Compute/images/win11-64-2009-eastus-win11-22h2-avd-alpha | 1 |
|  | 27 |
| projects/taskcluster-imaging/global/images/generic-translations-gcp-googlecompute-2023-11-01t20-12-18z | 1 |
| /subscriptions/108d46d5-fe9b-4850-9a7d-8c914aa6c1f0/resourceGroups/rg-packer-through-cib/providers/Microsoft.Compute/images/win11-64-2009-canadacentral-win11-22h2-avd-531c001 | 1 |
| /subscriptions/a30e97ab-734a-4f3b-a0e4-c51c0bff0701/resourceGroups/rg-packer-through-cib/providers/Microsoft.Compute/images/trusted-win2012r2-64-vs-py2-l3-centralus-2012-R2-Datacenter-7aa76c6 | 1 |
| /subscriptions/108d46d5-fe9b-4850-9a7d-8c914aa6c1f0/resourceGroups/rg-packer-through-cib/providers/Microsoft.Compute/images/win2022-64-2009-centralus-2022-datacenter-azure-edition-531c001 | 5 |
| /subscriptions/108d46d5-fe9b-4850-9a7d-8c914aa6c1f0/resourceGroups/rg-packer-through-cib/providers/Microsoft.Compute/images/win11-64-2009-uksouth-win11-22h2-avd-alpha | 2 |
| /subscriptions/108d46d5-fe9b-4850-9a7d-8c914aa6c1f0/resourceGroups/rg-packer-through-cib/providers/Microsoft.Compute/images/win10-64-2009-eastus-win10-22h2-avd-531c001 | 1 |
| /subscriptions/108d46d5-fe9b-4850-9a7d-8c914aa6c1f0/resourceGroups/rg-packer-through-cib/providers/Microsoft.Compute/images/win10-64-2009-canadacentral-win10-22h2-avd-531c001 | 2 |
| /subscriptions/a30e97ab-734a-4f3b-a0e4-c51c0bff0701/resourceGroups/rg-packer-through-cib/providers/Microsoft.Compute/images/trusted-win2012r2-64-l3-centralus-2012-R2-Datacenter-7cde253 | 1 |
| unknown | 1 |


| Worker Pool | Implementation | Version | Engine | Revision | OS | Arch | GO | Total Workers | Total Capacity |
| --- | --- | --- | --- | --- | --- | --- | --- | ---: | ---: |
| **bitbar/gecko-t-win64-aarch64-laptop** | generic-worker | 14.1.2 | <UNKNOWN> | 13118c4c1b | windows | 386 | 1.10.8 | 0 | 0 |
| **comm-1/b-win2022** | generic-worker | 54.4.1 | multiuser | a38a78d98c | windows | amd64 | 1.20.6 | 12 | 12 |
| **comm-2/b-win2022** | generic-worker | 54.4.1 | multiuser | a38a78d98c | windows | amd64 | 1.20.6 | 5 | 5 |
| **comm-3/b-win2012-azure** | generic-worker | 43.2.0 | multiuser | 53fa7b79bc | windows | amd64 | 1.16.3 | 62 | 62 |
| **comm-3/b-win2022** | generic-worker | 54.4.1 | multiuser | a38a78d98c | windows | amd64 | 1.20.6 | 567 | 567 |
| **gecko-1/b-win2022** | generic-worker | 54.4.1 | multiuser | a38a78d98c | windows | amd64 | 1.20.6 | 1163 | 1163 |
| **gecko-2/b-win2022** | generic-worker | 54.4.1 | multiuser | a38a78d98c | windows | amd64 | 1.20.6 | 59 | 59 |
| **gecko-3/b-win2022** | generic-worker | 54.4.1 | multiuser | a38a78d98c | windows | amd64 | 1.20.6 | 5160 | 5160 |
| **gecko-t/win10-64-2009** | generic-worker | 54.4.1 | multiuser | a38a78d98c | windows | amd64 | 1.20.6 | 4743 | 4743 |
| **gecko-t/win10-64-2009-gpu** | generic-worker | 54.4.1 | multiuser | a38a78d98c | windows | amd64 | 1.20.6 | 676 | 676 |
| **gecko-t/win10-64-2009-source** | generic-worker | 54.4.1 | multiuser | a38a78d98c | windows | amd64 | 1.20.6 | 4 | 4 |
| **gecko-t/win11-64-2009** | generic-worker | 54.4.1 | multiuser | a38a78d98c | windows | amd64 | 1.20.6 | 19289 | 19289 |
| **gecko-t/win11-64-2009-alpha** | generic-worker | 54.4.1 | multiuser | a38a78d98c | windows | amd64 | 1.20.6 | 23 | 23 |
| **gecko-t/win11-64-2009-gpu** | generic-worker | 54.4.1 | multiuser | a38a78d98c | windows | amd64 | 1.20.6 | 18450 | 18450 |
| **gecko-t/win11-64-2009-gpu-alpha** | generic-worker | 54.4.1 | multiuser | a38a78d98c | windows | amd64 | 1.20.6 | 260 | 260 |
| **gecko-t/win11-64-2009-source** | generic-worker | 54.4.1 | multiuser | a38a78d98c | windows | amd64 | 1.20.6 | 916 | 916 |
| **gecko-t/win11-64-2009-source-alpha** | generic-worker | 54.4.1 | multiuser | a38a78d98c | windows | amd64 | 1.20.6 | 12 | 12 |
| **gecko-t/win11-64-2009-ssd** | generic-worker | 54.4.1 | multiuser | a38a78d98c | windows | amd64 | 1.20.6 | 2183 | 2183 |
| **gecko-t/win11-64-2009-ssd-gpu** | generic-worker | 54.4.1 | multiuser | a38a78d98c | windows | amd64 | 1.20.6 | 1874 | 1874 |
| **localprovisioner/nss-macos-10-12** | generic-worker | 15.1.4 | simple | c407e45e3f | darwin | amd64 | 1.10.8 | 0 | 0 |
| **localprovisioner/nss-macos-m1** | generic-worker | 30.0.2 | simple | 6fdba0dad3 | darwin | arm64 | 1.16.4 | 0 | 0 |
| **mozillavpn-1/b-win2022** | generic-worker | 54.4.1 | multiuser | a38a78d98c | windows | amd64 | 1.20.6 | 119 | 119 |
| **mozillavpn-3/b-win2012-azure** | generic-worker | 43.2.0 | multiuser | 53fa7b79bc | windows | amd64 | 1.16.3 | 13 | 13 |
| **mozillavpn-3/b-win2022** | generic-worker | 54.4.1 | multiuser | a38a78d98c | windows | amd64 | 1.20.6 | 52 | 52 |
| **nss-3/b-win2012-azure** | generic-worker | 43.2.0 | multiuser | 53fa7b79bc | windows | amd64 | 1.16.3 | 139 | 139 |
| **performance-hardware/gecko-t-fxrecorder** | generic-worker | 44.0.0 | multiuser | faf6f319c0 | windows | amd64 | 1.16.3 | 0 | 0 |
| **proj-autophone/gecko-t-bitbar-gw-perf-a51** | generic-worker | 36.0.0 | simple | b9cb11293f | linux | amd64 | 1.13.7 | 0 | 0 |
| **proj-autophone/gecko-t-bitbar-gw-unit-p5** | generic-worker | 36.0.0 | simple | b9cb11293f | linux | amd64 | 1.13.7 | 0 | 0 |
| **releng-hardware/applicationservices-b-1-osx1015** | generic-worker | 54.4.1 | multiuser | a38a78d98c | darwin | amd64 | 1.20.6 | 0 | 0 |
| **releng-hardware/applicationservices-b-3-osx1015** | generic-worker | 54.4.1 | multiuser | a38a78d98c | darwin | amd64 | 1.20.6 | 0 | 0 |
| **releng-hardware/gecko-1-b-osx-1015** | generic-worker | 55.3.4 | multiuser | a71f300fcb | darwin | amd64 | 1.21.1 | 0 | 0 |
| **releng-hardware/gecko-1-b-osx-1015-staging** | generic-worker | 54.4.1 | multiuser | a38a78d98c | darwin | amd64 | 1.20.6 | 0 | 0 |
| **releng-hardware/gecko-3-b-osx-1015** | generic-worker | 55.3.4 | multiuser | a71f300fcb | darwin | amd64 | 1.21.1 | 0 | 0 |
| **releng-hardware/gecko-t-linux-talos-1804** | generic-worker | 44.23.2 | simple | da8da74563 | linux | amd64 | 1.18.5 | 0 | 0 |
| **releng-hardware/gecko-t-osx-1015-power** | generic-worker | 40.0.3 | simple | 5f7a31ac6e | darwin | amd64 | 1.15.6 | 0 | 0 |
| **releng-hardware/gecko-t-osx-1015-r8** | generic-worker | 55.3.4 | simple | a71f300fcb | darwin | amd64 | 1.21.1 | 0 | 0 |
| **releng-hardware/gecko-t-osx-1015-r8-staging** | generic-worker | 40.0.3 | simple | 5f7a31ac6e | darwin | amd64 | 1.15.6 | 0 | 0 |
| **releng-hardware/gecko-t-osx-1100-m1** | generic-worker | 55.3.4 | simple | a71f300fcb | darwin | arm64 | 1.21.1 | 0 | 0 |
| **releng-hardware/gecko-t-osx-1100-m1-loaner** | generic-worker | 30.0.2 | simple | 6fdba0dad3 | darwin | arm64 | 1.16.4 | 0 | 0 |
| **releng-hardware/gecko-t-osx-1100-m1-loaner-bug1829730** | generic-worker | 30.0.2 | simple | 6fdba0dad3 | darwin | arm64 | 1.16.4 | 0 | 0 |
| **releng-hardware/gecko-t-osx-1100-m1-staging** | generic-worker | 55.3.4 | simple | a71f300fcb | darwin | arm64 | 1.21.1 | 0 | 0 |
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
| **translations-1/b-linux-v100-gpu** | generic-worker | 55.1.1 | simple | b83fffd4bd | linux | amd64 | 1.21.1 | 142 | 142 |
| **translations-1/b-linux-v100-gpu-4** | generic-worker | 55.1.1 | simple | b83fffd4bd | linux | amd64 | 1.21.1 | 78 | 78 |
| **translations-1/b-linux-v100-gpu-4-1tb** | generic-worker | 55.1.1 | simple | b83fffd4bd | linux | amd64 | 1.21.1 | 51 | 51 |
| **translations-1/b-linux-v100-gpu-4-300gb** | generic-worker | 55.1.1 | simple | b83fffd4bd | linux | amd64 | 1.21.1 | 109 | 109 |


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
| projects/taskcluster-imaging/global/images/docker-firefoxci-gcp-l1-googlecompute-2023-06-13t23-47-26z | 74 |
| projects/fxci-production-level3-workers/global/images/docker-firefoxci-gcp-l3-googlecompute-2023-03-16t20-10-06z | 47 |
| projects/taskcluster-imaging/global/images/docker-worker-gcp-u14-04-2023-06-13 | 9 |
| projects/taskcluster-imaging/global/images/docker-firefoxci-gcp-lt-googlecompute-2023-04-13t21-30-28z | 1 |


| Worker Pool | Implementation | Version | Total Workers | Total Capacity |
| --- | --- | --- | ---: | ---: |
| **adhoc-1/b-linux-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **adhoc-1/decision-gcp** | docker-worker | 38.0.5 | 2 | 4 |
| **adhoc-1/images-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **adhoc-3/b-linux-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **adhoc-3/decision-gcp** | docker-worker | 38.0.5 | 2 | 4 |
| **adhoc-3/images-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **app-services-1/b-linux-gcp** | docker-worker | 38.0.5 | 148 | 148 |
| **app-services-1/decision-gcp** | docker-worker | 38.0.5 | 591 | 1182 |
| **app-services-1/images-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **app-services-3/b-linux-gcp** | docker-worker | 38.0.5 | 580 | 580 |
| **app-services-3/decision-gcp** | docker-worker | 38.0.5 | 29 | 58 |
| **app-services-3/images-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **ci-1/decision-gcp** | docker-worker | 38.0.5 | 18 | 36 |
| **ci-1/images-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **ci-3/decision-gcp** | docker-worker | 38.0.5 | 14 | 28 |
| **ci-3/images-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **ci-t/linux-gcp** | docker-worker | 38.0.5 | 131 | 131 |
| **code-analysis-1/linux-gcp** | docker-worker | 38.0.5 | 16 | 16 |
| **code-analysis-3/linux-gcp** | docker-worker | 38.0.5 | 16 | 16 |
| **code-coverage/bot-gcp** | docker-worker | 38.0.5 | 809 | 809 |
| **code-review/bot-gcp** | docker-worker | 38.0.5 | 2300 | 2300 |
| **comm-1/b-linux-gcp** | docker-worker | 38.0.5 | 338 | 338 |
| **comm-1/b-linux-large-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **comm-1/b-linux-xlarge-gcp** | docker-worker | 38.0.5 | 3 | 3 |
| **comm-1/decision-gcp** | docker-worker | 38.0.5 | 141 | 282 |
| **comm-1/images-gcp** | docker-worker | 38.0.5 | 8 | 8 |
| **comm-2/b-linux-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **comm-2/b-linux-large-gcp** | docker-worker | 38.0.5 | 3 | 3 |
| **comm-2/b-linux-xlarge-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **comm-2/decision-gcp** | docker-worker | 38.0.5 | 2 | 4 |
| **comm-2/images-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **comm-3/b-linux-gcp** | docker-worker | 38.0.5 | 2902 | 2902 |
| **comm-3/b-linux-large-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **comm-3/b-linux-xlarge-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **comm-3/decision-gcp** | docker-worker | 38.0.5 | 78 | 156 |
| **comm-3/images-gcp** | docker-worker | 38.0.5 | 3 | 3 |
| **comm-t/misc-gcp** | docker-worker | 38.0.5 | 3 | 24 |
| **comm-t/t-linux-large-gcp** | docker-worker | unknown version | 1638 | 1638 |
| **comm-t/t-linux-xlarge-gcp** | docker-worker | unknown version | 219 | 219 |
| **comm-t/t-linux-xlarge-source-gcp** | docker-worker | unknown version | 928 | 928 |
| **gecko-1/b-linux-gcp** | docker-worker | 38.0.5 | 6716 | 6716 |
| **gecko-1/b-linux-gcp-bug1797804-n2** | docker-worker | 38.0.5 | 2 | 2 |
| **gecko-1/b-linux-large-gcp** | docker-worker | 38.0.5 | 162 | 162 |
| **gecko-1/b-linux-xlarge-gcp** | docker-worker | 38.0.5 | 628 | 628 |
| **gecko-1/b-linux-xlarge-gcp-bug1797804-c2** | docker-worker | 38.0.5 | 2 | 2 |
| **gecko-1/decision-gcp** | docker-worker | 38.0.5 | 1436 | 2872 |
| **gecko-1/images-gcp** | docker-worker | 38.0.5 | 71 | 71 |
| **gecko-2/b-linux-gcp** | docker-worker | 38.0.5 | 771 | 771 |
| **gecko-2/b-linux-large-gcp** | docker-worker | 38.0.5 | 16 | 16 |
| **gecko-2/b-linux-xlarge-gcp** | docker-worker | 38.0.5 | 24 | 24 |
| **gecko-2/decision-gcp** | docker-worker | 38.0.5 | 20 | 40 |
| **gecko-2/images-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **gecko-3/b-linux-gcp** | docker-worker | 38.0.5 | 17518 | 17518 |
| **gecko-3/b-linux-large-gcp** | docker-worker | 38.0.5 | 135 | 135 |
| **gecko-3/b-linux-xlarge-gcp** | docker-worker | 38.0.5 | 903 | 903 |
| **gecko-3/decision-gcp** | docker-worker | 38.0.5 | 1451 | 2902 |
| **gecko-3/images-gcp** | docker-worker | 38.0.5 | 35 | 35 |
| **gecko-3/t-linux-xlarge-gcp** | docker-worker | 38.0.5 | 276 | 276 |
| **gecko-t/misc-gcp** | docker-worker | 38.0.5 | 299 | 2392 |
| **gecko-t/t-linux-kvm-gcp** | docker-worker | unknown version | 28054 | 28054 |
| **gecko-t/t-linux-large-gcp** | docker-worker | unknown version | 111190 | 111190 |
| **gecko-t/t-linux-large-hostub2204-gcp** | docker-worker | 48.3.0 | 2 | 2 |
| **gecko-t/t-linux-xlarge-gcp** | docker-worker | unknown version | 55588 | 55588 |
| **gecko-t/t-linux-xlarge-source-gcp** | docker-worker | unknown version | 9925 | 9925 |
| **glean-1/b-linux-gcp** | docker-worker | 38.0.5 | 1164 | 1164 |
| **glean-1/decision-gcp** | docker-worker | 38.0.5 | 592 | 1184 |
| **glean-1/images-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **glean-3/b-linux-gcp** | docker-worker | 38.0.5 | 1186 | 1186 |
| **glean-3/decision-gcp** | docker-worker | 38.0.5 | 6 | 12 |
| **glean-3/images-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **infra/build-decision** | docker-worker | 38.0.5 | 13 | 104 |
| **mobile-1/b-linux-gcp** | docker-worker | 38.0.5 | 1865 | 1865 |
| **mobile-1/b-linux-large-gcp** | docker-worker | 38.0.5 | 7491 | 7491 |
| **mobile-1/b-linux-xlarge-gcp** | docker-worker | 38.0.5 | 4079 | 4079 |
| **mobile-1/bitrise-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **mobile-1/decision-gcp** | docker-worker | 38.0.5 | 482 | 964 |
| **mobile-1/images-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **mobile-3/b-linux-gcp** | docker-worker | 38.0.5 | 1499 | 1499 |
| **mobile-3/b-linux-large-gcp** | docker-worker | 38.0.5 | 7673 | 7673 |
| **mobile-3/b-linux-xlarge-gcp** | docker-worker | 38.0.5 | 3944 | 3944 |
| **mobile-3/bitrise-gcp** | docker-worker | 38.0.5 | 21 | 21 |
| **mobile-3/decision-gcp** | docker-worker | 38.0.5 | 237 | 474 |
| **mobile-3/images-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **mozilla-1/b-linux-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **mozilla-1/decision-gcp** | docker-worker | 38.0.5 | 594 | 1188 |
| **mozilla-1/images-gcp** | docker-worker | 38.0.5 | 34 | 34 |
| **mozilla-t/t-linux-large-gcp** | docker-worker | unknown version | 37 | 37 |
| **mozillaonline-1/b-linux-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **mozillaonline-3/b-linux-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **mozillavpn-1/b-linux-gcp** | docker-worker | 38.0.5 | 385 | 385 |
| **mozillavpn-1/b-linux-large-gcp** | docker-worker | 38.0.5 | 263 | 263 |
| **mozillavpn-1/decision-gcp** | docker-worker | 38.0.5 | 581 | 1162 |
| **mozillavpn-1/images-gcp** | docker-worker | 38.0.5 | 7 | 7 |
| **mozillavpn-3/b-linux-gcp** | docker-worker | 38.0.5 | 186 | 186 |
| **mozillavpn-3/b-linux-large-gcp** | docker-worker | 38.0.5 | 167 | 167 |
| **mozillavpn-3/decision-gcp** | docker-worker | 38.0.5 | 29 | 58 |
| **mozillavpn-3/images-gcp** | docker-worker | 38.0.5 | 8 | 8 |
| **nss-1/decision-gcp** | docker-worker | 38.0.5 | 2 | 4 |
| **nss-1/linux-gcp** | docker-worker | 38.0.5 | 313 | 313 |
| **nss-3/linux-gcp** | docker-worker | 38.0.5 | 467 | 467 |
| **nss-t/t-linux-xlarge-gcp** | docker-worker | unknown version | 668 | 668 |
| **project-relman/relman-svc** | docker-worker | 38.0.5 | 68 | 68 |
| **releng-1/b-linux-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **releng-1/decision-gcp** | docker-worker | 38.0.5 | 16 | 32 |
| **releng-1/linux-gcp** | docker-worker | 38.0.5 | 15 | 15 |
| **releng-3/b-linux-gcp** | docker-worker | 38.0.5 | 6 | 6 |
| **releng-3/decision-gcp** | docker-worker | 38.0.5 | 13 | 26 |
| **releng-3/linux-gcp** | docker-worker | 38.0.5 | 20 | 20 |
| **releng-t/linux-gcp** | docker-worker | 38.0.5 | 108 | 108 |
| **relops-3/decision-gcp** | docker-worker | 38.0.5 | 2 | 4 |
| **scriptworker-1/b-linux-gcp** | docker-worker | 38.0.5 | 221 | 221 |
| **scriptworker-1/decision-gcp** | docker-worker | 38.0.5 | 34 | 68 |
| **scriptworker-1/images-gcp** | docker-worker | 38.0.5 | 4 | 4 |
| **scriptworker-3/b-linux-gcp** | docker-worker | 38.0.5 | 100 | 100 |
| **scriptworker-3/decision-gcp** | docker-worker | 38.0.5 | 17 | 34 |
| **scriptworker-3/images-gcp** | docker-worker | 38.0.5 | 21 | 21 |
| **taskgraph-1/decision-gcp** | docker-worker | 38.0.5 | 4 | 8 |
| **taskgraph-1/images-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **taskgraph-3/decision-gcp** | docker-worker | 38.0.5 | 2 | 4 |
| **taskgraph-3/images-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **taskgraph-t/linux-gcp** | docker-worker | 38.0.5 | 5 | 5 |
| **translations-1/b-linux-large-gcp** | docker-worker | 38.0.5 | 245 | 245 |
| **translations-1/b-linux-large-gcp-300gb** | docker-worker | 38.0.5 | 90 | 90 |
| **translations-1/decision-gcp** | docker-worker | 38.0.5 | 34 | 68 |
| **translations-1/images-gcp** | docker-worker | 38.0.5 | 5 | 5 |
| **xpi-1/b-linux-gcp** | docker-worker | 38.0.5 | 92 | 92 |
| **xpi-1/decision-gcp** | docker-worker | 38.0.5 | 26 | 52 |
| **xpi-1/images-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **xpi-3/b-linux-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **xpi-3/decision-gcp** | docker-worker | 38.0.5 | 2 | 4 |
| **xpi-3/images-gcp** | docker-worker | 38.0.5 | 2 | 2 |


## Script Worker

Total: `54`



| Worker Pool | Implementation | Version | Total Workers | Total Capacity |
| --- | --- | --- | ---: | ---: |
| **scriptworker-k8s/adhoc-1-shipit-dev** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/app-services-1-beetmover-dev** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/app-services-1-shipit** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/app-services-1-shipit-dev** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/app-services-3-beetmover** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/app-services-3-shipit** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/app-services-3-signing** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/comm-3-balrog** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/comm-3-beetmover** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/comm-3-bouncer** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/comm-3-signing** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/comm-3-tree** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/comm-t-signing** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/gecko-3-addon** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/gecko-3-balrog** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/gecko-3-beetmover** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/gecko-3-bouncer** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/gecko-3-pushflatpak** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/gecko-3-pushmsix** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/gecko-3-signing** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/gecko-3-tree** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/gecko-t-signing** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/mobile-1-tree** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/mobile-1-tree-dev** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/mobile-3-beetmover** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/mobile-3-pushapk** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/mobile-3-signing** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/mobile-3-tree** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/mobile-t-signing** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/mozillavpn-1-beetmover** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/mozillavpn-1-beetmover-dev** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/mozillavpn-1-pushapk** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/mozillavpn-1-pushapk-dev** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/mozillavpn-3-beetmover** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/mozillavpn-3-signing** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/mozillavpn-t-signing** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/xpi-1-balrog** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/xpi-1-balrog-dev** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/xpi-1-beetmover** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/xpi-1-beetmover-dev** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/xpi-1-github** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/xpi-1-github-dev** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/xpi-3-balrog** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/xpi-3-beetmover** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/xpi-3-github** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/xpi-t-signing** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
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
| **gecko-t/t-linux-vm-2204-wayland** |  | No artifacts found | 970 | 970 |
| **gecko-t/t-linux-vm-2204-wayland-aerickson-test** |  | No artifacts found | 4 | 4 |
| **scriptworker-prov-v1/mac-notarization-poller** |  | No artifacts found | 0 | 0 |
| **scriptworker-prov-v1/tb-mac-notarization-poller** |  | No artifacts found | 0 | 0 |


## Version not determined [^2]

Total: `24`


Count by image:

| Version | Count |
| :--- | ---: |
| /subscriptions/a30e97ab-734a-4f3b-a0e4-c51c0bff0701/resourceGroups/rg-packer-through-cib/providers/Microsoft.Compute/images/trusted-win2012r2-64-l3-centralus-2012-R2-Datacenter-7cde253 | 1 |
| /subscriptions/108d46d5-fe9b-4850-9a7d-8c914aa6c1f0/resourceGroups/rg-packer-through-cib/providers/Microsoft.Compute/images/win2012r2-64-l1-centralus-2012-R2-Datacenter-alpha | 2 |
|  | 1 |
| /subscriptions/a30e97ab-734a-4f3b-a0e4-c51c0bff0701/resourceGroups/rg-packer-through-cib/providers/Microsoft.Compute/images/trusted-win2022-64-2009-centralus-2022-datacenter-azure-edition-d5dee60 | 1 |
| /subscriptions/108d46d5-fe9b-4850-9a7d-8c914aa6c1f0/resourceGroups/rg-packer-through-cib/providers/Microsoft.Compute/images/masterwayz-win2012test-image-20221122010013-alpha | 1 |
| /subscriptions/108d46d5-fe9b-4850-9a7d-8c914aa6c1f0/resourceGroups/rg-packer-through-cib/providers/Microsoft.Compute/images/win2012r2-64-vs-l1-centralus-2012-R2-Datacenter-21c6826 | 1 |
| /subscriptions/108d46d5-fe9b-4850-9a7d-8c914aa6c1f0/resourceGroups/rg-packer-through-cib/providers/Microsoft.Compute/images/win10-64-2009-ukwest-win10-22h2-avd-alpha | 1 |
| /subscriptions/108d46d5-fe9b-4850-9a7d-8c914aa6c1f0/resourceGroups/rg-packer-through-cib/providers/Microsoft.Compute/images/win2022-64-2009-centralus-2022-datacenter-azure-edition-alpha | 2 |
| /subscriptions/108d46d5-fe9b-4850-9a7d-8c914aa6c1f0/resourceGroups/rg-packer-through-cib/providers/Microsoft.Compute/images/win11-64-2009-canadacentral-win11-22h2-avd-531c001 | 1 |
| /subscriptions/108d46d5-fe9b-4850-9a7d-8c914aa6c1f0/resourceGroups/rg-packer-through-cib/providers/Microsoft.Compute/images/win2022-64-2009-rd-centralus-2022-datacenter-azure-edition-alpha | 1 |
| /subscriptions/108d46d5-fe9b-4850-9a7d-8c914aa6c1f0/resourceGroups/rg-packer-through-cib/providers/Microsoft.Compute/images/win2012r2-64-vs-l1-centralus-2012-R2-Datacenter-alpha | 1 |
| /subscriptions/108d46d5-fe9b-4850-9a7d-8c914aa6c1f0/resourceGroups/rg-packer-through-cib/providers/Microsoft.Compute/images/win2012r2-64-l1-centralus-2012-R2-Datacenter-21c6826 | 4 |
| /subscriptions/108d46d5-fe9b-4850-9a7d-8c914aa6c1f0/resourceGroups/rg-packer-through-cib/providers/Microsoft.Compute/images/win2022-64-2009-centralus-2022-datacenter-azure-edition-531c001 | 1 |
| /subscriptions/108d46d5-fe9b-4850-9a7d-8c914aa6c1f0/resourceGroups/rg-packer-through-cib/providers/Microsoft.Compute/images/win2012r2-64-vs-py2-l1-centralus-2012-R2-Datacenter-7aa76c6 | 1 |
| /subscriptions/108d46d5-fe9b-4850-9a7d-8c914aa6c1f0/resourceGroups/rg-packer-through-cib/providers/Microsoft.Compute/images/win2012r2-64-vs-py2-l1-centralus-2012-R2-Datacenter-alpha | 1 |
| /subscriptions/108d46d5-fe9b-4850-9a7d-8c914aa6c1f0/resourceGroups/rg-packer-through-cib/providers/Microsoft.Compute/images/win11-64-2009-tc-centralus-win11-22h2-avd-alpha | 1 |
| /subscriptions/108d46d5-fe9b-4850-9a7d-8c914aa6c1f0/resourceGroups/rg-packer-through-cib/providers/Microsoft.Compute/images/win11-64-2009-centralus-win11-22h2-avd-69a7ecf | 1 |
| /subscriptions/108d46d5-fe9b-4850-9a7d-8c914aa6c1f0/resourceGroups/rg-packer-through-cib/providers/Microsoft.Compute/images/win10-64-2009-centralus-win10-22h2-avd-alpha | 1 |
| /subscriptions/108d46d5-fe9b-4850-9a7d-8c914aa6c1f0/resourceGroups/rg-packer-through-cib/providers/Microsoft.Compute/images/win10-64-2009-eastus-win10-22h2-avd-alpha | 1 |


| Worker Pool | Implementation | Version | Total Workers | Total Capacity |
| --- | --- | --- | ---: | ---: |
| **comm-1/b-win2012-azure** |  | Version not determined; task not (yet) claimed | 7 | 7 |
| **comm-2/b-win2012-azure** |  | Version not determined; task not (yet) claimed | 7 | 7 |
| **comm-t/win11-64-2009** |  | Version not determined; task not (yet) claimed | 4264 | 4264 |
| **gecko-1/b-win2012-azure** |  | Version not determined; task not (yet) claimed | 2 | 2 |
| **gecko-1/b-win2012-azure-alpha** |  | Version not determined; task not (yet) claimed | 7 | 7 |
| **gecko-1/b-win2012-azure-beta** |  | Version not determined; task not (yet) claimed | 7 | 7 |
| **gecko-1/b-win2022-alpha** |  | Version not determined; task not (yet) claimed | 7 | 7 |
| **gecko-1/b-win2022-rd-alpha** |  | Version not determined; task not (yet) claimed | 7 | 7 |
| **gecko-2/b-win2012-azure** |  | Version not determined; task not (yet) claimed | 7 | 7 |
| **gecko-3/b-win2012-azure** |  | Version not determined; task not (yet) claimed | 232 | 232 |
| **gecko-t/win10-64-2009-alpha** |  | Version not determined; task not (yet) claimed | 6 | 6 |
| **gecko-t/win10-64-2009-gpu-alpha** |  | Version not determined; task not (yet) claimed | 6 | 6 |
| **gecko-t/win10-64-2009-source-alpha** |  | Version not determined; task not (yet) claimed | 6 | 6 |
| **gecko-t/win11-64-2009-alpha-test** |  | Version not determined; task not (yet) claimed | 7 | 7 |
| **gecko-t/win11-64-2009-tc-alpha** |  | Version not determined; task not (yet) claimed | 8 | 8 |
| **mozillavpn-1/b-win2012-azure** |  | Version not determined; task not (yet) claimed | 122 | 122 |
| **mozillavpn-1/b-win2012-azure-alpha** |  | Version not determined; task not (yet) claimed | 8 | 8 |
| **mozillavpn-1/masterwayz-win2012-azure** |  | Version not determined; task not (yet) claimed | 8 | 8 |
| **nss-1/b-win2012-azure** |  | Version not determined; task not (yet) claimed | 1846 | 1846 |
| **nss-1/b-win2012-azure-alpha** |  | Version not determined; task not (yet) claimed | 8 | 8 |
| **nss-1/b-win2022-alpha** |  | Version not determined; task not (yet) claimed | 8 | 8 |
| **releng-hardware/gecko-t-osx-1300-m2** |  | Version not determined; task not (yet) claimed | 0 | 0 |
| **relops-1/b-win2022** |  | Version not determined; task not (yet) claimed | 17 | 17 |
| **relops-3/b-win2022** |  | Version not determined; task not (yet) claimed | 17 | 17 |



[^1]: Those are the pools whose tasks were claimed and resolved by a worker as expected, but the worker did not publish either artifact `public/logs/live_backing.log` nor `public/logs/chain_of_trust.log`, which is the source used to identify the worker implementation.

[^2]: Probing task remains pending after two hours. Those are the pools that were not able to start any worker to claim the task within two hours.

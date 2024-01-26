

# Worker Pool Versions


## Generic Worker

Total: `69`

Count by version:

| Version | Count |
| :--- | ---: |
| 14.1.2 | 1 |
| 15.1.4 | 1 |
| 16.2.0 | 3 |
| 16.5.1 | 1 |
| 30.0.2 | 1 |
| 36.0.0 | 2 |
| 40.0.3 | 3 |
| 43.2.0 | 5 |
| 44.0.0 | 1 |
| 44.23.2 | 1 |
| 45.0.0 | 1 |
| 54.4.1 | 29 |
| 55.3.4 | 10 |
| 57.0.1 | 1 |
| 59.2.0 | 3 |
| 60.1.2 | 6 |


Count by image:

| Version | Count |
| :--- | ---: |
| /subscriptions/108d46d5-fe9b-4850-9a7d-8c914aa6c1f0/resourceGroups/rg-packer-through-cib/providers/Microsoft.Compute/images/win10-64-2009-canadacentral-win10-22h2-avd-531c001 | 2 |
| /subscriptions/108d46d5-fe9b-4850-9a7d-8c914aa6c1f0/resourceGroups/rg-packer-through-cib/providers/Microsoft.Compute/images/win11-64-2009-canadacentral-win11-22h2-avd-531c001 | 2 |
| /subscriptions/a30e97ab-734a-4f3b-a0e4-c51c0bff0701/resourceGroups/rg-packer-through-cib/providers/Microsoft.Compute/images/trusted-win2022-64-2009-centralus-2022-datacenter-azure-edition-28c1765 | 3 |
| /subscriptions/108d46d5-fe9b-4850-9a7d-8c914aa6c1f0/resourceGroups/rg-packer-through-cib/providers/Microsoft.Compute/images/win2022-64-2009-centralus-2022-datacenter-azure-edition-alpha | 2 |
| /subscriptions/108d46d5-fe9b-4850-9a7d-8c914aa6c1f0/resourceGroups/rg-packer-through-cib/providers/Microsoft.Compute/images/win10-64-2009-eastus-win10-22h2-avd-531c001 | 1 |
| /subscriptions/108d46d5-fe9b-4850-9a7d-8c914aa6c1f0/resourceGroups/rg-packer-through-cib/providers/Microsoft.Compute/images/win11-64-2009-eastus-win11-22h2-avd-alpha | 1 |
| /subscriptions/108d46d5-fe9b-4850-9a7d-8c914aa6c1f0/resourceGroups/rg-packer-through-cib/providers/Microsoft.Compute/images/win11-64-2009-centralus-win11-22h2-avd-531c001 | 3 |
| /subscriptions/a30e97ab-734a-4f3b-a0e4-c51c0bff0701/resourceGroups/rg-packer-through-cib/providers/Microsoft.Compute/images/trusted-win2012r2-64-l3-centralus-2012-R2-Datacenter-7cde253 | 2 |
| /subscriptions/108d46d5-fe9b-4850-9a7d-8c914aa6c1f0/resourceGroups/rg-packer-through-cib/providers/Microsoft.Compute/images/win11-64-2009-eastus-win11-22h2-avd-531c001 | 2 |
| /subscriptions/108d46d5-fe9b-4850-9a7d-8c914aa6c1f0/resourceGroups/rg-packer-through-cib/providers/Microsoft.Compute/images/win2022-64-2009-centralus-2022-datacenter-azure-edition-531c001 | 5 |
| ami-07db1c26bda143e33 | 1 |
|  | 29 |
| projects/taskcluster-imaging/global/images/gw-fxci-gcp-l1-gui-googlecompute-2024-01-18t23-27-40z | 2 |
| /subscriptions/a30e97ab-734a-4f3b-a0e4-c51c0bff0701/resourceGroups/rg-packer-through-cib/providers/Microsoft.Compute/images/trusted-win2012r2-64-vs-l3-centralus-2012-R2-Datacenter-7cde253 | 1 |
| projects/taskcluster-imaging/global/images/generic-translations-gcp-googlecompute-2024-01-24t16-29-51z | 6 |
| /subscriptions/a30e97ab-734a-4f3b-a0e4-c51c0bff0701/resourceGroups/rg-packer-through-cib/providers/Microsoft.Compute/images/trusted-win2012r2-64-vs-py2-l3-centralus-2012-R2-Datacenter-7aa76c6 | 1 |
| /subscriptions/108d46d5-fe9b-4850-9a7d-8c914aa6c1f0/resourceGroups/rg-packer-through-cib/providers/Microsoft.Compute/images/win2012r2-64-vs-py2-l1-centralus-2012-R2-Datacenter-7aa76c6 | 1 |
| projects/fxci-production-level3-workers/global/images/gw-fxci-gcp-l3-gui-googlecompute-2024-01-22t18-50-07z | 1 |
| unknown | 1 |
| projects/taskcluster-imaging/global/images/generic-translations-gcp-googlecompute-2023-11-01t20-12-18z | 1 |
| /subscriptions/108d46d5-fe9b-4850-9a7d-8c914aa6c1f0/resourceGroups/rg-packer-through-cib/providers/Microsoft.Compute/images/win11-64-2009-uksouth-win11-22h2-avd-alpha | 2 |


| Worker Pool | Implementation | Version | Engine | Revision | OS | Arch | GO | Total Workers | Total Capacity |
| --- | --- | --- | --- | --- | --- | --- | --- | ---: | ---: |
| **bitbar/gecko-t-win64-aarch64-laptop** | generic-worker | 14.1.2 | <UNKNOWN> | 13118c4c1b | windows | 386 | 1.10.8 | 0 | 0 |
| **comm-1/b-win2022** | generic-worker | 54.4.1 | multiuser | a38a78d98c | windows | amd64 | 1.20.6 | 27 | 27 |
| **comm-2/b-win2022** | generic-worker | 54.4.1 | multiuser | a38a78d98c | windows | amd64 | 1.20.6 | 4 | 4 |
| **comm-3/b-win2012-azure** | generic-worker | 43.2.0 | multiuser | 53fa7b79bc | windows | amd64 | 1.16.3 | 65 | 65 |
| **comm-3/b-win2022** | generic-worker | 54.4.1 | multiuser | a38a78d98c | windows | amd64 | 1.20.6 | 627 | 627 |
| **comm-t/win11-64-2009** | generic-worker | 54.4.1 | multiuser | a38a78d98c | windows | amd64 | 1.20.6 | 4058 | 4058 |
| **comm-t/win11-64-2009-ssd** | generic-worker | 54.4.1 | multiuser | a38a78d98c | windows | amd64 | 1.20.6 | 580 | 580 |
| **gecko-1/b-linux-2204-kvm-gcp** | generic-worker | 59.2.0 | multiuser | 45b996ca98 | linux | amd64 | 1.21.5 | 554 | 554 |
| **gecko-1/b-win2022** | generic-worker | 54.4.1 | multiuser | a38a78d98c | windows | amd64 | 1.20.6 | 830 | 830 |
| **gecko-1/b-win2022-alpha** | generic-worker | 54.4.1 | multiuser | a38a78d98c | windows | amd64 | 1.20.6 | 14 | 14 |
| **gecko-2/b-win2022** | generic-worker | 54.4.1 | multiuser | a38a78d98c | windows | amd64 | 1.20.6 | 110 | 110 |
| **gecko-3/b-linux-2204-kvm-gcp** | generic-worker | 59.2.0 | multiuser | 45b996ca98 | linux | amd64 | 1.21.5 | 170 | 170 |
| **gecko-3/b-win2012-azure** | generic-worker | 43.2.0 | multiuser | 53fa7b79bc | windows | amd64 | 1.16.3 | 16 | 16 |
| **gecko-3/b-win2022** | generic-worker | 54.4.1 | multiuser | a38a78d98c | windows | amd64 | 1.20.6 | 5371 | 5371 |
| **gecko-t/t-linux-2204-wayland-experimental** | generic-worker | 59.2.0 | multiuser | 45b996ca98 | linux | amd64 | 1.21.5 | 73 | 73 |
| **gecko-t/win10-64-2009** | generic-worker | 54.4.1 | multiuser | a38a78d98c | windows | amd64 | 1.20.6 | 5602 | 5602 |
| **gecko-t/win10-64-2009-gpu** | generic-worker | 54.4.1 | multiuser | a38a78d98c | windows | amd64 | 1.20.6 | 639 | 639 |
| **gecko-t/win10-64-2009-source** | generic-worker | 54.4.1 | multiuser | a38a78d98c | windows | amd64 | 1.20.6 | 4 | 4 |
| **gecko-t/win11-64-2009** | generic-worker | 54.4.1 | multiuser | a38a78d98c | windows | amd64 | 1.20.6 | 17854 | 17854 |
| **gecko-t/win11-64-2009-alpha** | generic-worker | 54.4.1 | multiuser | a38a78d98c | windows | amd64 | 1.20.6 | 4 | 4 |
| **gecko-t/win11-64-2009-gpu** | generic-worker | 54.4.1 | multiuser | a38a78d98c | windows | amd64 | 1.20.6 | 19029 | 19029 |
| **gecko-t/win11-64-2009-gpu-alpha** | generic-worker | 54.4.1 | multiuser | a38a78d98c | windows | amd64 | 1.20.6 | 4 | 4 |
| **gecko-t/win11-64-2009-source** | generic-worker | 54.4.1 | multiuser | a38a78d98c | windows | amd64 | 1.20.6 | 1654 | 1654 |
| **gecko-t/win11-64-2009-source-alpha** | generic-worker | 54.4.1 | multiuser | a38a78d98c | windows | amd64 | 1.20.6 | 4 | 4 |
| **gecko-t/win11-64-2009-ssd** | generic-worker | 54.4.1 | multiuser | a38a78d98c | windows | amd64 | 1.20.6 | 1892 | 1892 |
| **gecko-t/win11-64-2009-ssd-gpu** | generic-worker | 54.4.1 | multiuser | a38a78d98c | windows | amd64 | 1.20.6 | 1592 | 1592 |
| **localprovisioner/nss-macos-10-12** | generic-worker | 15.1.4 | simple | c407e45e3f | darwin | amd64 | 1.10.8 | 0 | 0 |
| **localprovisioner/nss-macos-m1** | generic-worker | 30.0.2 | simple | 6fdba0dad3 | darwin | arm64 | 1.16.4 | 0 | 0 |
| **mozillavpn-1/b-win2022** | generic-worker | 54.4.1 | multiuser | a38a78d98c | windows | amd64 | 1.20.6 | 301 | 301 |
| **mozillavpn-3/b-win2012-azure** | generic-worker | 43.2.0 | multiuser | 53fa7b79bc | windows | amd64 | 1.16.3 | 4 | 4 |
| **mozillavpn-3/b-win2022** | generic-worker | 54.4.1 | multiuser | a38a78d98c | windows | amd64 | 1.20.6 | 119 | 119 |
| **nss-1/b-win2012-azure** | generic-worker | 43.2.0 | multiuser | 53fa7b79bc | windows | amd64 | 1.16.3 | 240 | 240 |
| **nss-1/b-win2022-alpha** | generic-worker | 54.4.1 | multiuser | a38a78d98c | windows | amd64 | 1.20.6 | 13 | 13 |
| **nss-3/b-win2012-azure** | generic-worker | 43.2.0 | multiuser | 53fa7b79bc | windows | amd64 | 1.16.3 | 111 | 111 |
| **performance-hardware/gecko-t-fxrecorder** | generic-worker | 44.0.0 | multiuser | faf6f319c0 | windows | amd64 | 1.16.3 | 0 | 0 |
| **proj-autophone/gecko-t-bitbar-gw-perf-a51** | generic-worker | 36.0.0 | simple | b9cb11293f | linux | amd64 | 1.13.7 | 0 | 0 |
| **proj-autophone/gecko-t-bitbar-gw-unit-p5** | generic-worker | 36.0.0 | simple | b9cb11293f | linux | amd64 | 1.13.7 | 0 | 0 |
| **releng-hardware/applicationservices-b-1-osx1015** | generic-worker | 54.4.1 | multiuser | a38a78d98c | darwin | amd64 | 1.20.6 | 0 | 0 |
| **releng-hardware/applicationservices-b-3-osx1015** | generic-worker | 54.4.1 | multiuser | a38a78d98c | darwin | amd64 | 1.20.6 | 0 | 0 |
| **releng-hardware/gecko-1-b-osx-1015** | generic-worker | 40.0.3 | multiuser | 5f7a31ac6e | darwin | amd64 | 1.15.6 | 0 | 0 |
| **releng-hardware/gecko-1-b-osx-1015-staging** | generic-worker | 54.4.1 | multiuser | a38a78d98c | darwin | amd64 | 1.20.6 | 0 | 0 |
| **releng-hardware/gecko-1-b-osx-arm64** | generic-worker | 55.3.4 | multiuser | a71f300fcb | darwin | arm64 | 1.21.1 | 0 | 0 |
| **releng-hardware/gecko-3-b-osx-1015** | generic-worker | 55.3.4 | multiuser | a71f300fcb | darwin | amd64 | 1.21.1 | 0 | 0 |
| **releng-hardware/gecko-3-b-osx-arm64** | generic-worker | 55.3.4 | multiuser | a71f300fcb | darwin | arm64 | 1.21.1 | 0 | 0 |
| **releng-hardware/gecko-t-linux-talos-1804** | generic-worker | 44.23.2 | simple | da8da74563 | linux | amd64 | 1.18.5 | 0 | 0 |
| **releng-hardware/gecko-t-osx-1015-power** | generic-worker | 40.0.3 | simple | 5f7a31ac6e | darwin | amd64 | 1.15.6 | 0 | 0 |
| **releng-hardware/gecko-t-osx-1015-r8** | generic-worker | 55.3.4 | simple | a71f300fcb | darwin | amd64 | 1.21.1 | 0 | 0 |
| **releng-hardware/gecko-t-osx-1015-r8-staging** | generic-worker | 40.0.3 | simple | 5f7a31ac6e | darwin | amd64 | 1.15.6 | 0 | 0 |
| **releng-hardware/gecko-t-osx-1100-m1** | generic-worker | 55.3.4 | simple | a71f300fcb | darwin | arm64 | 1.21.1 | 0 | 0 |
| **releng-hardware/gecko-t-osx-1100-m1-staging** | generic-worker | 55.3.4 | simple | a71f300fcb | darwin | arm64 | 1.21.1 | 0 | 0 |
| **releng-hardware/gecko-t-osx-1300-m2** | generic-worker | 55.3.4 | simple | a71f300fcb | darwin | arm64 | 1.21.1 | 0 | 0 |
| **releng-hardware/gecko-t-osx-1400-r8-staging** | generic-worker | 55.3.4 | simple | a71f300fcb | darwin | amd64 | 1.21.1 | 0 | 0 |
| **releng-hardware/gecko-t-win10-64-1803-hw** | generic-worker | 16.2.0 | multiuser | b321a877c8 | windows | amd64 | 1.10.8 | 0 | 0 |
| **releng-hardware/gecko-t-win10-64-dev** | generic-worker | 16.2.0 | multiuser | b321a877c8 | windows | amd64 | 1.10.8 | 0 | 0 |
| **releng-hardware/gecko-t-win10-64-ref-hw** | generic-worker | 16.2.0 | multiuser | b321a877c8 | windows | amd64 | 1.10.8 | 0 | 0 |
| **releng-hardware/gecko-t-win7-32-hw** | generic-worker | 45.0.0 | multiuser | 988e8100b3 | windows | 386 | 1.19.3 | 0 | 0 |
| **releng-hardware/mozillavpn-b-1-osx** | generic-worker | 54.4.1 | multiuser | a38a78d98c | darwin | amd64 | 1.20.6 | 0 | 0 |
| **releng-hardware/mozillavpn-b-3-osx** | generic-worker | 54.4.1 | multiuser | a38a78d98c | darwin | amd64 | 1.20.6 | 0 | 0 |
| **releng-hardware/nss-1-b-osx-1015** | generic-worker | 55.3.4 | multiuser | a71f300fcb | darwin | amd64 | 1.21.1 | 0 | 0 |
| **releng-hardware/nss-3-b-osx-1015** | generic-worker | 55.3.4 | multiuser | a71f300fcb | darwin | amd64 | 1.21.1 | 0 | 0 |
| **releng-hardware/win11-64-2009-hw-ref** | generic-worker | 54.4.1 | multiuser | a38a78d98c | windows | amd64 | 1.20.6 | 0 | 0 |
| **relops-3/win2019** | generic-worker | 16.5.1 | multiuser | 02bc4c9cd7 | windows | amd64 | 1.10.8 | 4 | 4 |
| **translations-1/b-linux-aerickson-test** | generic-worker | 57.0.1 | simple | fe45895655 | linux | amd64 | 1.21.3 | 2 | 2 |
| **translations-1/b-linux-v100-gpu** | generic-worker | 60.1.2 | simple | c9b1d20769 | linux | amd64 | 1.21.5 | 140 | 140 |
| **translations-1/b-linux-v100-gpu-4** | generic-worker | 60.1.2 | simple | c9b1d20769 | linux | amd64 | 1.21.5 | 112 | 112 |
| **translations-1/b-linux-v100-gpu-4-1tb** | generic-worker | 60.1.2 | simple | c9b1d20769 | linux | amd64 | 1.21.5 | 65 | 65 |
| **translations-1/b-linux-v100-gpu-4-1tb-standard** | generic-worker | 60.1.2 | simple | c9b1d20769 | linux | amd64 | 1.21.5 | 7 | 7 |
| **translations-1/b-linux-v100-gpu-4-300gb** | generic-worker | 60.1.2 | simple | c9b1d20769 | linux | amd64 | 1.21.5 | 260 | 260 |
| **translations-1/b-linux-v100-gpu-4-300gb-standard** | generic-worker | 60.1.2 | simple | c9b1d20769 | linux | amd64 | 1.21.5 | 4 | 4 |


## Docker Worker

Total: `133`

Count by version:

| Version | Count |
| :--- | ---: |
| 38.0.5 | 123 |
| 48.3.0 | 1 |
| unknown version | 9 |


Count by image:

| Version | Count |
| :--- | ---: |
| projects/taskcluster-imaging/global/images/docker-firefoxci-gcp-l1-googlecompute-2023-12-04t22-11-57z | 75 |
| projects/fxci-production-level3-workers/global/images/docker-firefoxci-gcp-l3-googlecompute-2023-12-06t15-04-40z | 48 |
| projects/taskcluster-imaging/global/images/docker-worker-gcp-u14-04-2023-06-13 | 9 |
| projects/taskcluster-imaging/global/images/docker-firefoxci-gcp-lt-googlecompute-2023-04-13t21-30-28z | 1 |


| Worker Pool | Implementation | Version | Total Workers | Total Capacity |
| --- | --- | --- | ---: | ---: |
| **adhoc-1/b-linux-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **adhoc-1/decision-gcp** | docker-worker | 38.0.5 | 8 | 16 |
| **adhoc-1/images-gcp** | docker-worker | 38.0.5 | 35 | 35 |
| **adhoc-3/b-linux-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **adhoc-3/decision-gcp** | docker-worker | 38.0.5 | 4 | 8 |
| **adhoc-3/images-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **app-services-1/b-linux-gcp** | docker-worker | 38.0.5 | 560 | 560 |
| **app-services-1/decision-gcp** | docker-worker | 38.0.5 | 578 | 1156 |
| **app-services-1/images-gcp** | docker-worker | 38.0.5 | 8 | 8 |
| **app-services-3/b-linux-gcp** | docker-worker | 38.0.5 | 1171 | 1171 |
| **app-services-3/decision-gcp** | docker-worker | 38.0.5 | 36 | 72 |
| **app-services-3/images-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **ci-1/decision-gcp** | docker-worker | 38.0.5 | 27 | 54 |
| **ci-1/images-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **ci-3/decision-gcp** | docker-worker | 38.0.5 | 28 | 56 |
| **ci-3/images-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **ci-t/linux-gcp** | docker-worker | 38.0.5 | 293 | 293 |
| **code-analysis-1/linux-gcp** | docker-worker | 38.0.5 | 16 | 16 |
| **code-analysis-3/linux-gcp** | docker-worker | 38.0.5 | 13 | 13 |
| **code-coverage/bot-gcp** | docker-worker | 38.0.5 | 818 | 818 |
| **code-review/bot-gcp** | docker-worker | 38.0.5 | 3074 | 3074 |
| **comm-1/b-linux-gcp** | docker-worker | 38.0.5 | 718 | 718 |
| **comm-1/b-linux-large-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **comm-1/b-linux-xlarge-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **comm-1/decision-gcp** | docker-worker | 38.0.5 | 140 | 280 |
| **comm-1/images-gcp** | docker-worker | 38.0.5 | 20 | 20 |
| **comm-2/b-linux-gcp** | docker-worker | 38.0.5 | 3 | 3 |
| **comm-2/b-linux-large-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **comm-2/b-linux-xlarge-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **comm-2/decision-gcp** | docker-worker | 38.0.5 | 2 | 4 |
| **comm-2/images-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **comm-3/b-linux-gcp** | docker-worker | 38.0.5 | 5184 | 5184 |
| **comm-3/b-linux-large-gcp** | docker-worker | 38.0.5 | 3 | 3 |
| **comm-3/b-linux-xlarge-gcp** | docker-worker | 38.0.5 | 4 | 4 |
| **comm-3/decision-gcp** | docker-worker | 38.0.5 | 85 | 170 |
| **comm-3/images-gcp** | docker-worker | 38.0.5 | 18 | 18 |
| **comm-t/misc-gcp** | docker-worker | 38.0.5 | 2 | 16 |
| **comm-t/t-linux-large-gcp** | docker-worker | unknown version | 1401 | 1401 |
| **comm-t/t-linux-xlarge-gcp** | docker-worker | unknown version | 239 | 239 |
| **comm-t/t-linux-xlarge-source-gcp** | docker-worker | unknown version | 1101 | 1101 |
| **gecko-1/b-linux-gcp** | docker-worker | 38.0.5 | 11939 | 11939 |
| **gecko-1/b-linux-gcp-bug1797804-n2** | docker-worker | 38.0.5 | 2 | 2 |
| **gecko-1/b-linux-kvm-gcp** | docker-worker | 38.0.5 | 34 | 34 |
| **gecko-1/b-linux-large-gcp** | docker-worker | 38.0.5 | 1391 | 1391 |
| **gecko-1/b-linux-xlarge-gcp** | docker-worker | 38.0.5 | 688 | 688 |
| **gecko-1/b-linux-xlarge-gcp-bug1797804-c2** | docker-worker | 38.0.5 | 5 | 5 |
| **gecko-1/decision-gcp** | docker-worker | 38.0.5 | 1345 | 2690 |
| **gecko-1/images-gcp** | docker-worker | 38.0.5 | 207 | 207 |
| **gecko-2/b-linux-gcp** | docker-worker | 38.0.5 | 2863 | 2863 |
| **gecko-2/b-linux-large-gcp** | docker-worker | 38.0.5 | 35 | 35 |
| **gecko-2/b-linux-xlarge-gcp** | docker-worker | 38.0.5 | 35 | 35 |
| **gecko-2/decision-gcp** | docker-worker | 38.0.5 | 33 | 66 |
| **gecko-2/images-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **gecko-3/b-linux-gcp** | docker-worker | 38.0.5 | 30006 | 30006 |
| **gecko-3/b-linux-kvm-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **gecko-3/b-linux-large-gcp** | docker-worker | 38.0.5 | 93 | 93 |
| **gecko-3/b-linux-xlarge-gcp** | docker-worker | 38.0.5 | 1242 | 1242 |
| **gecko-3/decision-gcp** | docker-worker | 38.0.5 | 1564 | 3128 |
| **gecko-3/images-gcp** | docker-worker | 38.0.5 | 5 | 5 |
| **gecko-3/t-linux-xlarge-gcp** | docker-worker | 38.0.5 | 226 | 226 |
| **gecko-t/misc-gcp** | docker-worker | 38.0.5 | 334 | 2672 |
| **gecko-t/t-linux-kvm-gcp** | docker-worker | unknown version | 25562 | 25562 |
| **gecko-t/t-linux-large-gcp** | docker-worker | unknown version | 79134 | 79134 |
| **gecko-t/t-linux-large-hostub2204-gcp** | docker-worker | 48.3.0 | 2 | 2 |
| **gecko-t/t-linux-xlarge-gcp** | docker-worker | unknown version | 39132 | 39132 |
| **gecko-t/t-linux-xlarge-source-gcp** | docker-worker | unknown version | 10318 | 10318 |
| **glean-1/b-linux-gcp** | docker-worker | 38.0.5 | 1199 | 1199 |
| **glean-1/decision-gcp** | docker-worker | 38.0.5 | 601 | 1202 |
| **glean-1/images-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **glean-3/b-linux-gcp** | docker-worker | 38.0.5 | 1206 | 1206 |
| **glean-3/decision-gcp** | docker-worker | 38.0.5 | 5 | 10 |
| **glean-3/images-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **infra/build-decision** | docker-worker | 38.0.5 | 13 | 104 |
| **mobile-1/b-linux-gcp** | docker-worker | 38.0.5 | 2061 | 2061 |
| **mobile-1/b-linux-large-gcp** | docker-worker | 38.0.5 | 8817 | 8817 |
| **mobile-1/b-linux-xlarge-gcp** | docker-worker | 38.0.5 | 4330 | 4330 |
| **mobile-1/bitrise-gcp** | docker-worker | 38.0.5 | 13 | 13 |
| **mobile-1/decision-gcp** | docker-worker | 38.0.5 | 465 | 930 |
| **mobile-1/images-gcp** | docker-worker | 38.0.5 | 6 | 6 |
| **mobile-3/b-linux-gcp** | docker-worker | 38.0.5 | 1555 | 1555 |
| **mobile-3/b-linux-large-gcp** | docker-worker | 38.0.5 | 7726 | 7726 |
| **mobile-3/b-linux-xlarge-gcp** | docker-worker | 38.0.5 | 3925 | 3925 |
| **mobile-3/bitrise-gcp** | docker-worker | 38.0.5 | 18 | 18 |
| **mobile-3/decision-gcp** | docker-worker | 38.0.5 | 251 | 502 |
| **mobile-3/images-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **mozilla-1/b-linux-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **mozilla-1/decision-gcp** | docker-worker | 38.0.5 | 603 | 1206 |
| **mozilla-1/images-gcp** | docker-worker | 38.0.5 | 28 | 28 |
| **mozilla-t/t-linux-large-gcp** | docker-worker | unknown version | 38 | 38 |
| **mozillaonline-1/b-linux-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **mozillaonline-3/b-linux-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **mozillavpn-1/b-linux-gcp** | docker-worker | 38.0.5 | 476 | 476 |
| **mozillavpn-1/b-linux-large-gcp** | docker-worker | 38.0.5 | 396 | 396 |
| **mozillavpn-1/decision-gcp** | docker-worker | 38.0.5 | 579 | 1158 |
| **mozillavpn-1/images-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **mozillavpn-3/b-linux-gcp** | docker-worker | 38.0.5 | 249 | 249 |
| **mozillavpn-3/b-linux-large-gcp** | docker-worker | 38.0.5 | 231 | 231 |
| **mozillavpn-3/decision-gcp** | docker-worker | 38.0.5 | 33 | 66 |
| **mozillavpn-3/images-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **nss-1/decision-gcp** | docker-worker | 38.0.5 | 2 | 4 |
| **nss-1/linux-gcp** | docker-worker | 38.0.5 | 900 | 900 |
| **nss-3/linux-gcp** | docker-worker | 38.0.5 | 509 | 509 |
| **nss-t/t-linux-xlarge-gcp** | docker-worker | unknown version | 1704 | 1704 |
| **project-relman/relman-svc** | docker-worker | 38.0.5 | 70 | 70 |
| **releng-1/b-linux-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **releng-1/decision-gcp** | docker-worker | 38.0.5 | 20 | 40 |
| **releng-1/linux-gcp** | docker-worker | 38.0.5 | 41 | 41 |
| **releng-3/b-linux-gcp** | docker-worker | 38.0.5 | 14 | 14 |
| **releng-3/decision-gcp** | docker-worker | 38.0.5 | 15 | 30 |
| **releng-3/linux-gcp** | docker-worker | 38.0.5 | 41 | 41 |
| **releng-t/linux-gcp** | docker-worker | 38.0.5 | 71 | 71 |
| **relops-3/decision-gcp** | docker-worker | 38.0.5 | 2 | 4 |
| **scriptworker-1/b-linux-gcp** | docker-worker | 38.0.5 | 44 | 44 |
| **scriptworker-1/decision-gcp** | docker-worker | 38.0.5 | 10 | 20 |
| **scriptworker-1/images-gcp** | docker-worker | 38.0.5 | 18 | 18 |
| **scriptworker-3/b-linux-gcp** | docker-worker | 38.0.5 | 37 | 37 |
| **scriptworker-3/decision-gcp** | docker-worker | 38.0.5 | 12 | 24 |
| **scriptworker-3/images-gcp** | docker-worker | 38.0.5 | 27 | 27 |
| **taskgraph-1/decision-gcp** | docker-worker | 38.0.5 | 11 | 22 |
| **taskgraph-1/images-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **taskgraph-3/decision-gcp** | docker-worker | 38.0.5 | 10 | 20 |
| **taskgraph-3/images-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **taskgraph-t/linux-gcp** | docker-worker | 38.0.5 | 29 | 29 |
| **translations-1/b-linux-large-gcp** | docker-worker | 38.0.5 | 345 | 345 |
| **translations-1/b-linux-large-gcp-300gb** | docker-worker | 38.0.5 | 136 | 136 |
| **translations-1/decision-gcp** | docker-worker | 38.0.5 | 47 | 94 |
| **translations-1/images-gcp** | docker-worker | 38.0.5 | 24 | 24 |
| **xpi-1/b-linux-gcp** | docker-worker | 38.0.5 | 16 | 16 |
| **xpi-1/decision-gcp** | docker-worker | 38.0.5 | 6 | 12 |
| **xpi-1/images-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **xpi-3/b-linux-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **xpi-3/decision-gcp** | docker-worker | 38.0.5 | 2 | 4 |
| **xpi-3/images-gcp** | docker-worker | 38.0.5 | 2 | 2 |


## Script Worker

Total: `28`



| Worker Pool | Implementation | Version | Total Workers | Total Capacity |
| --- | --- | --- | ---: | ---: |
| **scriptworker-k8s/app-services-3-beetmover** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/app-services-3-signing** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/comm-1-beetmover** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/comm-3-balrog** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/comm-3-beetmover** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/comm-3-signing** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/comm-3-tree** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/comm-t-signing** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
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

Total: `5`


Count by image:

| Version | Count |
| :--- | ---: |
|  | 4 |
| projects/taskcluster-imaging/global/images/generic-2204-wayland-vm-gcp-googlecompute-2023-09-22t17-39-37z | 1 |


| Worker Pool | Implementation | Version | Total Workers | Total Capacity |
| --- | --- | --- | ---: | ---: |
| **built-in/fail** |  | No artifacts found | 0 | 0 |
| **built-in/succeed** |  | No artifacts found | 0 | 0 |
| **gecko-t/t-linux-vm-2204-wayland** |  | No artifacts found | 1662 | 1662 |
| **scriptworker-prov-v1/mac-notarization-poller** |  | No artifacts found | 0 | 0 |
| **scriptworker-prov-v1/tb-mac-notarization-poller** |  | No artifacts found | 0 | 0 |


## Version not determined [^2]

Total: `21`


Count by image:

| Version | Count |
| :--- | ---: |
|  | 3 |
| /subscriptions/a30e97ab-734a-4f3b-a0e4-c51c0bff0701/resourceGroups/rg-packer-through-cib/providers/Microsoft.Compute/images/trusted-win2022-64-2009-centralus-2022-datacenter-azure-edition-28c1765 | 1 |
| /subscriptions/108d46d5-fe9b-4850-9a7d-8c914aa6c1f0/resourceGroups/rg-packer-through-cib/providers/Microsoft.Compute/images/win2012r2-64-vs-l1-centralus-2012-R2-Datacenter-alpha | 1 |
| /subscriptions/108d46d5-fe9b-4850-9a7d-8c914aa6c1f0/resourceGroups/rg-packer-through-cib/providers/Microsoft.Compute/images/win2012r2-64-vs-l1-centralus-2012-R2-Datacenter-21c6826 | 1 |
| /subscriptions/108d46d5-fe9b-4850-9a7d-8c914aa6c1f0/resourceGroups/rg-packer-through-cib/providers/Microsoft.Compute/images/win2012r2-64-l1-centralus-2012-R2-Datacenter-21c6826 | 4 |
| /subscriptions/108d46d5-fe9b-4850-9a7d-8c914aa6c1f0/resourceGroups/rg-packer-through-cib/providers/Microsoft.Compute/images/win2012r2-64-vs-py2-l1-centralus-2012-R2-Datacenter-alpha | 1 |
| /subscriptions/108d46d5-fe9b-4850-9a7d-8c914aa6c1f0/resourceGroups/rg-packer-through-cib/providers/Microsoft.Compute/images/masterwayz-win2012test-image-20221122010013-alpha | 1 |
| /subscriptions/108d46d5-fe9b-4850-9a7d-8c914aa6c1f0/resourceGroups/rg-packer-through-cib/providers/Microsoft.Compute/images/win10-64-2009-centralus-win10-22h2-avd-alpha | 1 |
| /subscriptions/108d46d5-fe9b-4850-9a7d-8c914aa6c1f0/resourceGroups/rg-packer-through-cib/providers/Microsoft.Compute/images/win10-64-2009-eastus-win10-22h2-avd-alpha | 1 |
| /subscriptions/108d46d5-fe9b-4850-9a7d-8c914aa6c1f0/resourceGroups/rg-packer-through-cib/providers/Microsoft.Compute/images/win10-64-2009-ukwest-win10-22h2-avd-alpha | 1 |
| /subscriptions/108d46d5-fe9b-4850-9a7d-8c914aa6c1f0/resourceGroups/rg-packer-through-cib/providers/Microsoft.Compute/images/win2022-64-2009-centralus-2022-datacenter-azure-edition-531c001 | 1 |
| /subscriptions/108d46d5-fe9b-4850-9a7d-8c914aa6c1f0/resourceGroups/rg-packer-through-cib/providers/Microsoft.Compute/images/win11-64-2009-tc-centralus-win11-22h2-avd-alpha | 1 |
| /subscriptions/108d46d5-fe9b-4850-9a7d-8c914aa6c1f0/resourceGroups/rg-packer-through-cib/providers/Microsoft.Compute/images/win11-64-2009-centralus-win11-22h2-avd-69a7ecf | 1 |
| /subscriptions/108d46d5-fe9b-4850-9a7d-8c914aa6c1f0/resourceGroups/rg-packer-through-cib/providers/Microsoft.Compute/images/win2022-64-2009-rd-centralus-2022-datacenter-azure-edition-alpha | 1 |
| /subscriptions/108d46d5-fe9b-4850-9a7d-8c914aa6c1f0/resourceGroups/rg-packer-through-cib/providers/Microsoft.Compute/images/win2012r2-64-l1-centralus-2012-R2-Datacenter-alpha | 2 |


| Worker Pool | Implementation | Version | Total Workers | Total Capacity |
| --- | --- | --- | ---: | ---: |
| **comm-1/b-win2012-azure** |  | Version not determined; task not (yet) claimed | 17 | 17 |
| **comm-2/b-win2012-azure** |  | Version not determined; task not (yet) claimed | 17 | 17 |
| **gecko-1/b-win2012-azure** |  | Version not determined; task not (yet) claimed | 19 | 19 |
| **gecko-1/b-win2012-azure-alpha** |  | Version not determined; task not (yet) claimed | 17 | 17 |
| **gecko-1/b-win2012-azure-beta** |  | Version not determined; task not (yet) claimed | 17 | 17 |
| **gecko-1/b-win2022-rd-alpha** |  | Version not determined; task not (yet) claimed | 17 | 17 |
| **gecko-2/b-win2012-azure** |  | Version not determined; task not (yet) claimed | 17 | 17 |
| **gecko-t/win10-64-2009-alpha** |  | Version not determined; task not (yet) claimed | 16 | 16 |
| **gecko-t/win10-64-2009-gpu-alpha** |  | Version not determined; task not (yet) claimed | 16 | 16 |
| **gecko-t/win10-64-2009-source-alpha** |  | Version not determined; task not (yet) claimed | 16 | 16 |
| **gecko-t/win11-64-2009-alpha-test** |  | Version not determined; task not (yet) claimed | 16 | 16 |
| **gecko-t/win11-64-2009-tc-alpha** |  | Version not determined; task not (yet) claimed | 16 | 16 |
| **mozillavpn-1/b-win2012-azure** |  | Version not determined; task not (yet) claimed | 16 | 16 |
| **mozillavpn-1/b-win2012-azure-alpha** |  | Version not determined; task not (yet) claimed | 16 | 16 |
| **mozillavpn-1/masterwayz-win2012-azure** |  | Version not determined; task not (yet) claimed | 16 | 16 |
| **nss-1/b-win2012-azure-alpha** |  | Version not determined; task not (yet) claimed | 16 | 16 |
| **releng-hardware/gecko-t-osx-1100-m1-loaner** |  | Version not determined; task not (yet) claimed | 0 | 0 |
| **releng-hardware/gecko-t-osx-1100-m1-loaner-bug1829730** |  | Version not determined; task not (yet) claimed | 0 | 0 |
| **relops-1/b-win2022** |  | Version not determined; task not (yet) claimed | 103 | 103 |
| **relops-3/b-win2022** |  | Version not determined; task not (yet) claimed | 114 | 114 |
| **translations-1/b-linux-v100-gpu-4-300gb-relops813** |  | Version not determined; task not (yet) claimed | 0 | 0 |



[^1]: Those are the pools whose tasks were claimed and resolved by a worker as expected, but the worker did not publish either artifact `public/logs/live_backing.log` nor `public/logs/chain_of_trust.log`, which is the source used to identify the worker implementation.

[^2]: Probing task remains pending after two hours. Those are the pools that were not able to start any worker to claim the task within two hours.

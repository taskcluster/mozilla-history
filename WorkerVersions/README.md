

# Worker Pool Versions


## Generic Worker

Total: `96`

Count by version:

| Version | Count |
| :--- | ---: |
| 14.1.2 | 1 |
| 15.1.4 | 1 |
| 16.2.0 | 1 |
| 16.5.1 | 1 |
| 30.0.2 | 1 |
| 36.0.0 | 5 |
| 40.0.3 | 2 |
| 43.2.0 | 9 |
| 44.0.0 | 1 |
| 45.0.0 | 1 |
| 54.4.1 | 1 |
| 55.3.4 | 1 |
| 57.0.1 | 1 |
| 60.3.4 | 16 |
| 60.4.0 | 1 |
| 61.0.0 | 1 |
| 64.2.5 | 22 |
| 64.2.6 | 9 |
| 64.2.7 | 3 |
| 64.3.0 | 16 |
| 65.1.0 | 1 |
| 67.0.0 | 1 |


Count by image:

| Version | Count |
| :--- | ---: |
|  | 40 |
| /subscriptions/108d46d5-fe9b-4850-9a7d-8c914aa6c1f0/resourceGroups/rg-packer-through-cib/providers/Microsoft.Compute/images/win2012r2-64-vs-py2-l1-centralus-2012-R2-Datacenter-7aa76c6 | 1 |
| /subscriptions/a30e97ab-734a-4f3b-a0e4-c51c0bff0701/resourceGroups/rg-packer-through-cib/providers/Microsoft.Compute/images/trusted-win2012r2-64-vs-py2-l3-centralus-2012-R2-Datacenter-7aa76c6 | 1 |
| /subscriptions/108d46d5-fe9b-4850-9a7d-8c914aa6c1f0/resourceGroups/rg-packer-through-cib/providers/Microsoft.Compute/images/win10-64-2009-eastus-win10-22h2-avd-8561b62 | 1 |
| /subscriptions/108d46d5-fe9b-4850-9a7d-8c914aa6c1f0/resourceGroups/rg-packer-through-cib/providers/Microsoft.Compute/images/win11-64-2009-northeurope-win11-22h2-avd-8561b62 | 1 |
| unknown | 1 |
| /subscriptions/108d46d5-fe9b-4850-9a7d-8c914aa6c1f0/resourceGroups/rg-packer-through-cib/providers/Microsoft.Compute/images/win2012r2-64-l1-centralus-2012-R2-Datacenter-75bd9ed | 4 |
| projects/taskcluster-imaging/global/images/gw-translations-gcp-googlecompute-2024-04-22t18-22-42z | 6 |
| /subscriptions/a30e97ab-734a-4f3b-a0e4-c51c0bff0701/resourceGroups/rg-packer-through-cib/providers/Microsoft.Compute/images/trusted-win2012r2-64-l3-centralus-2012-R2-Datacenter-7cde253 | 2 |
| projects/fxci-production-level3-workers/global/images/gw-fxci-gcp-l3-gui-googlecompute-2024-05-14t00-44-32z | 3 |
| projects/taskcluster-imaging/global/images/generic-translations-gcp-googlecompute-2023-11-01t20-12-18z | 1 |
| /subscriptions/a30e97ab-734a-4f3b-a0e4-c51c0bff0701/resourceGroups/rg-packer-through-cib/providers/Microsoft.Compute/images/trusted-win2022-64-2009-centralus-2022-datacenter-azure-edition-b1285eb | 3 |
| /subscriptions/108d46d5-fe9b-4850-9a7d-8c914aa6c1f0/resourceGroups/rg-packer-through-cib/providers/Microsoft.Compute/images/win11-64-2009-eastus-win11-22h2-avd-8561b62 | 4 |
| projects/taskcluster-imaging/global/images/gw-fxci-gcp-l1-gui-googlecompute-2024-05-31t15-18-14z | 16 |
| ami-07db1c26bda143e33 | 1 |
| projects/taskcluster-imaging/global/images/gw-fxci-gcp-l1-arm64-gui-googlecompute-2024-03-01t17-45-17z | 1 |
| /subscriptions/a30e97ab-734a-4f3b-a0e4-c51c0bff0701/resourceGroups/rg-packer-through-cib/providers/Microsoft.Compute/images/trusted-win2012r2-64-vs-l3-centralus-2012-R2-Datacenter-7cde253 | 1 |
| /subscriptions/108d46d5-fe9b-4850-9a7d-8c914aa6c1f0/resourceGroups/rg-packer-through-cib/providers/Microsoft.Compute/images/win11-64-2009-canadacentral-win11-22h2-avd-8561b62 | 2 |
| /subscriptions/108d46d5-fe9b-4850-9a7d-8c914aa6c1f0/resourceGroups/rg-packer-through-cib/providers/Microsoft.Compute/images/win2022-64-2009-centralus-2022-datacenter-azure-edition-b1285eb | 5 |
| /subscriptions/108d46d5-fe9b-4850-9a7d-8c914aa6c1f0/resourceGroups/rg-packer-through-cib/providers/Microsoft.Compute/images/win10-64-2009-canadacentral-win10-22h2-avd-8561b62 | 2 |


| Worker Pool | Implementation | Version | Engine | Revision | OS | Arch | GO | Total Workers | Total Capacity |
| --- | --- | --- | --- | --- | --- | --- | --- | ---: | ---: |
| **bitbar/gecko-t-win64-aarch64-laptop** | generic-worker | 14.1.2 | <UNKNOWN> | 13118c4c1b | windows | 386 | 1.10.8 | 0 | 0 |
| **code-analysis-1/linux-gw-gcp** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | amd64 | 1.22.2 | 87 | 87 |
| **code-analysis-3/linux-gw-gcp** | generic-worker | 64.2.6 | multiuser | edab196d7d | linux | amd64 | 1.22.2 | 73 | 73 |
| **comm-1/b-win2012-azure** | generic-worker | 43.2.0 | multiuser | 53fa7b79bc | windows | amd64 | 1.16.3 | 4 | 4 |
| **comm-1/b-win2022** | generic-worker | 64.2.5 | multiuser | 5f78921b3a | windows | amd64 | 1.22.2 | 3 | 3 |
| **comm-2/b-win2012-azure** | generic-worker | 43.2.0 | multiuser | 53fa7b79bc | windows | amd64 | 1.16.3 | 4 | 4 |
| **comm-2/b-win2022** | generic-worker | 64.2.5 | multiuser | 5f78921b3a | windows | amd64 | 1.22.2 | 3 | 3 |
| **comm-3/b-win2012-azure** | generic-worker | 43.2.0 | multiuser | 53fa7b79bc | windows | amd64 | 1.16.3 | 68 | 68 |
| **comm-3/b-win2022** | generic-worker | 64.2.5 | multiuser | 5f78921b3a | windows | amd64 | 1.22.2 | 537 | 537 |
| **comm-t/win11-64-2009** | generic-worker | 64.2.5 | multiuser | 5f78921b3a | windows | amd64 | 1.22.2 | 3678 | 3678 |
| **comm-t/win11-64-2009-ssd** | generic-worker | 64.2.5 | multiuser | 5f78921b3a | windows | amd64 | 1.22.2 | 442 | 442 |
| **gecko-1/b-linux-2204-kvm-gcp** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | amd64 | 1.22.2 | 2 | 2 |
| **gecko-1/b-linux-gcp-aarch64** | generic-worker | 60.4.0 | multiuser | f5527730e1 | linux | arm64 | 1.22.0 | 48 | 48 |
| **gecko-1/b-win2012-azure** | generic-worker | 43.2.0 | multiuser | 53fa7b79bc | windows | amd64 | 1.16.3 | 6 | 6 |
| **gecko-1/b-win2022** | generic-worker | 64.2.5 | multiuser | 5f78921b3a | windows | amd64 | 1.22.2 | 957 | 957 |
| **gecko-2/b-win2012-azure** | generic-worker | 43.2.0 | multiuser | 53fa7b79bc | windows | amd64 | 1.16.3 | 4 | 4 |
| **gecko-2/b-win2022** | generic-worker | 64.2.5 | multiuser | 5f78921b3a | windows | amd64 | 1.22.2 | 122 | 122 |
| **gecko-3/b-linux-2204-kvm-gcp** | generic-worker | 64.2.6 | multiuser | edab196d7d | linux | amd64 | 1.22.2 | 2 | 2 |
| **gecko-3/b-win2012-azure** | generic-worker | 43.2.0 | multiuser | 53fa7b79bc | windows | amd64 | 1.16.3 | 316 | 316 |
| **gecko-3/b-win2022** | generic-worker | 64.2.5 | multiuser | 5f78921b3a | windows | amd64 | 1.22.2 | 5684 | 5684 |
| **gecko-t/t-linux-2204-wayland** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | amd64 | 1.22.2 | 7659 | 7659 |
| **gecko-t/t-linux-2204-wayland-root-exp** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | amd64 | 1.22.2 | 2 | 2 |
| **gecko-t/t-linux-2204-wayland-snap** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | amd64 | 1.22.2 | 504 | 504 |
| **gecko-t/t-linux-xlarge-2204-wayland** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | amd64 | 1.22.2 | 1114 | 1114 |
| **gecko-t/win10-64-2009** | generic-worker | 64.2.5 | multiuser | 5f78921b3a | windows | amd64 | 1.22.2 | 5742 | 5742 |
| **gecko-t/win10-64-2009-gpu** | generic-worker | 64.2.5 | multiuser | 5f78921b3a | windows | amd64 | 1.22.2 | 627 | 627 |
| **gecko-t/win10-64-2009-source** | generic-worker | 64.2.5 | multiuser | 5f78921b3a | windows | amd64 | 1.22.2 | 5 | 5 |
| **gecko-t/win11-64-2009** | generic-worker | 64.2.5 | multiuser | 5f78921b3a | windows | amd64 | 1.22.2 | 42447 | 42447 |
| **gecko-t/win11-64-2009-gpu** | generic-worker | 64.2.5 | multiuser | 5f78921b3a | windows | amd64 | 1.22.2 | 21576 | 21576 |
| **gecko-t/win11-64-2009-source** | generic-worker | 64.2.5 | multiuser | 5f78921b3a | windows | amd64 | 1.22.2 | 1787 | 1787 |
| **gecko-t/win11-64-2009-ssd** | generic-worker | 64.2.5 | multiuser | 5f78921b3a | windows | amd64 | 1.22.2 | 1950 | 1950 |
| **gecko-t/win11-64-2009-ssd-gpu** | generic-worker | 64.2.5 | multiuser | 5f78921b3a | windows | amd64 | 1.22.2 | 1620 | 1620 |
| **localprovisioner/nss-macos-10-12** | generic-worker | 15.1.4 | simple | c407e45e3f | darwin | amd64 | 1.10.8 | 0 | 0 |
| **localprovisioner/nss-macos-m1** | generic-worker | 30.0.2 | simple | 6fdba0dad3 | darwin | arm64 | 1.16.4 | 0 | 0 |
| **mozilla-t/t-linux-2204-wayland** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | amd64 | 1.22.2 | 872 | 872 |
| **mozillavpn-1/b-linux-gcp-gw** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | amd64 | 1.22.2 | 40 | 40 |
| **mozillavpn-1/b-win2022** | generic-worker | 64.2.5 | multiuser | 5f78921b3a | windows | amd64 | 1.22.2 | 37 | 37 |
| **mozillavpn-3/b-linux-gcp-gw** | generic-worker | 64.2.6 | multiuser | edab196d7d | linux | amd64 | 1.22.2 | 33 | 33 |
| **mozillavpn-3/b-win2012-azure** | generic-worker | 43.2.0 | multiuser | 53fa7b79bc | windows | amd64 | 1.16.3 | 4 | 4 |
| **mozillavpn-3/b-win2022** | generic-worker | 64.2.5 | multiuser | 5f78921b3a | windows | amd64 | 1.22.2 | 32 | 32 |
| **nss-1/b-win2012-azure** | generic-worker | 43.2.0 | multiuser | 53fa7b79bc | windows | amd64 | 1.16.3 | 88 | 88 |
| **nss-3/b-win2012-azure** | generic-worker | 43.2.0 | multiuser | 53fa7b79bc | windows | amd64 | 1.16.3 | 44 | 44 |
| **performance-hardware/gecko-t-fxrecorder** | generic-worker | 44.0.0 | multiuser | faf6f319c0 | windows | amd64 | 1.16.3 | 0 | 0 |
| **proj-autophone/gecko-t-bitbar-gw-perf-a55** | generic-worker | 36.0.0 | simple | b9cb11293f | linux | amd64 | 1.13.7 | 0 | 0 |
| **proj-autophone/gecko-t-bitbar-gw-perf-p6** | generic-worker | 36.0.0 | simple | b9cb11293f | linux | amd64 | 1.13.7 | 0 | 0 |
| **proj-autophone/gecko-t-bitbar-gw-perf-s21** | generic-worker | 36.0.0 | simple | b9cb11293f | linux | amd64 | 1.13.7 | 0 | 0 |
| **proj-autophone/gecko-t-bitbar-gw-perf-s24** | generic-worker | 36.0.0 | simple | b9cb11293f | linux | amd64 | 1.13.7 | 0 | 0 |
| **proj-autophone/gecko-t-bitbar-gw-unit-p5** | generic-worker | 36.0.0 | simple | b9cb11293f | linux | amd64 | 1.13.7 | 0 | 0 |
| **releng-hardware/applicationservices-b-1-osx1015** | generic-worker | 60.3.4 | multiuser | 943a6f2b0d | darwin | amd64 | 1.22.0 | 0 | 0 |
| **releng-hardware/applicationservices-b-3-osx1015** | generic-worker | 60.3.4 | multiuser | 943a6f2b0d | darwin | amd64 | 1.22.0 | 0 | 0 |
| **releng-hardware/gecko-1-b-osx-1015** | generic-worker | 40.0.3 | multiuser | 5f7a31ac6e | darwin | amd64 | 1.15.6 | 0 | 0 |
| **releng-hardware/gecko-1-b-osx-1015-staging** | generic-worker | 54.4.1 | multiuser | a38a78d98c | darwin | amd64 | 1.20.6 | 0 | 0 |
| **releng-hardware/gecko-1-b-osx-arm64** | generic-worker | 60.3.4 | multiuser | 943a6f2b0d | darwin | arm64 | 1.22.0 | 0 | 0 |
| **releng-hardware/gecko-3-b-osx-1015** | generic-worker | 60.3.4 | multiuser | 943a6f2b0d | darwin | amd64 | 1.22.0 | 0 | 0 |
| **releng-hardware/gecko-3-b-osx-arm64** | generic-worker | 60.3.4 | multiuser | 943a6f2b0d | darwin | arm64 | 1.22.0 | 0 | 0 |
| **releng-hardware/gecko-t-linux-netperf-1804** | generic-worker | 65.1.0 | insecure | 1a085daa37 | linux | amd64 | 1.22.3 | 0 | 0 |
| **releng-hardware/gecko-t-linux-talos-1804** | generic-worker | 61.0.0 | simple | 3bd4419b4b | linux | amd64 | 1.22.1 | 0 | 0 |
| **releng-hardware/gecko-t-osx-1015-power** | generic-worker | 40.0.3 | simple | 5f7a31ac6e | darwin | amd64 | 1.15.6 | 0 | 0 |
| **releng-hardware/gecko-t-osx-1015-r8** | generic-worker | 60.3.4 | simple | 943a6f2b0d | darwin | amd64 | 1.22.0 | 0 | 0 |
| **releng-hardware/gecko-t-osx-1015-r8-staging** | generic-worker | 60.3.4 | multiuser | 943a6f2b0d | darwin | amd64 | 1.22.0 | 0 | 0 |
| **releng-hardware/gecko-t-osx-1100-m1** | generic-worker | 60.3.4 | simple | 943a6f2b0d | darwin | arm64 | 1.22.0 | 0 | 0 |
| **releng-hardware/gecko-t-osx-1100-m1-staging** | generic-worker | 60.3.4 | simple | 943a6f2b0d | darwin | arm64 | 1.22.0 | 0 | 0 |
| **releng-hardware/gecko-t-osx-1100-r8-latest** | generic-worker | 60.3.4 | simple | 943a6f2b0d | darwin | amd64 | 1.22.0 | 0 | 0 |
| **releng-hardware/gecko-t-osx-1200-r8-latest** | generic-worker | 60.3.4 | simple | 943a6f2b0d | darwin | amd64 | 1.22.0 | 0 | 0 |
| **releng-hardware/gecko-t-osx-1300-r8-latest** | generic-worker | 60.3.4 | simple | 943a6f2b0d | darwin | amd64 | 1.22.0 | 0 | 0 |
| **releng-hardware/gecko-t-osx-1400-m2** | generic-worker | 60.3.4 | simple | 943a6f2b0d | darwin | arm64 | 1.22.0 | 0 | 0 |
| **releng-hardware/gecko-t-osx-1400-r8-latest** | generic-worker | 60.3.4 | simple | 943a6f2b0d | darwin | amd64 | 1.22.0 | 0 | 0 |
| **releng-hardware/gecko-t-win10-64-1803-hw** | generic-worker | 16.2.0 | multiuser | b321a877c8 | windows | amd64 | 1.10.8 | 0 | 0 |
| **releng-hardware/gecko-t-win7-32-hw** | generic-worker | 45.0.0 | multiuser | 988e8100b3 | windows | 386 | 1.19.3 | 0 | 0 |
| **releng-hardware/mozillavpn-b-1-osx** | generic-worker | 60.3.4 | multiuser | 943a6f2b0d | darwin | amd64 | 1.22.0 | 0 | 0 |
| **releng-hardware/mozillavpn-b-3-osx** | generic-worker | 60.3.4 | multiuser | 943a6f2b0d | darwin | amd64 | 1.22.0 | 0 | 0 |
| **releng-hardware/nss-1-b-osx-1015** | generic-worker | 55.3.4 | multiuser | a71f300fcb | darwin | amd64 | 1.21.1 | 0 | 0 |
| **releng-hardware/win11-64-2009-hw** | generic-worker | 64.2.5 | multiuser | 5f78921b3a | windows | amd64 | 1.22.2 | 0 | 0 |
| **releng-hardware/win11-64-2009-hw-alpha** | generic-worker | 64.2.5 | multiuser | 5f78921b3a | windows | amd64 | 1.22.2 | 0 | 0 |
| **releng-hardware/win11-64-2009-hw-ref** | generic-worker | 64.2.5 | multiuser | 5f78921b3a | windows | amd64 | 1.22.2 | 0 | 0 |
| **releng-hardware/win11-64-2009-hw-ref-alpha** | generic-worker | 64.2.5 | multiuser | 5f78921b3a | windows | amd64 | 1.22.2 | 0 | 0 |
| **relops-3/win2019** | generic-worker | 16.5.1 | multiuser | 02bc4c9cd7 | windows | amd64 | 1.10.8 | 4 | 4 |
| **translations-1/b-linux-aerickson-test** | generic-worker | 57.0.1 | simple | fe45895655 | linux | amd64 | 1.21.3 | 2 | 2 |
| **translations-1/b-linux-large-gcp-1tb-32-256-d2g** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | amd64 | 1.22.2 | 5 | 5 |
| **translations-1/b-linux-large-gcp-1tb-32-256-std-d2g** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | amd64 | 1.22.2 | 5 | 5 |
| **translations-1/b-linux-large-gcp-1tb-64-512-d2g** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | amd64 | 1.22.2 | 138 | 138 |
| **translations-1/b-linux-large-gcp-1tb-64-512-std-d2g** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | amd64 | 1.22.2 | 46 | 46 |
| **translations-1/b-linux-large-gcp-d2g** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | amd64 | 1.22.2 | 4 | 4 |
| **translations-1/b-linux-large-gcp-d2g-1tb** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | amd64 | 1.22.2 | 4 | 4 |
| **translations-1/b-linux-large-gcp-d2g-1tb-standard** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | amd64 | 1.22.2 | 4 | 4 |
| **translations-1/b-linux-large-gcp-d2g-300gb** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | amd64 | 1.22.2 | 4 | 4 |
| **translations-1/b-linux-v100-gpu** | generic-worker | 64.2.6 | insecure | edab196d7d | linux | amd64 | 1.22.2 | 1047 | 1047 |
| **translations-1/b-linux-v100-gpu-4** | generic-worker | 64.2.6 | insecure | edab196d7d | linux | amd64 | 1.22.2 | 10292 | 10292 |
| **translations-1/b-linux-v100-gpu-4-1tb** | generic-worker | 64.2.6 | insecure | edab196d7d | linux | amd64 | 1.22.2 | 3396 | 3396 |
| **translations-1/b-linux-v100-gpu-4-1tb-standard** | generic-worker | 64.2.6 | insecure | edab196d7d | linux | amd64 | 1.22.2 | 2 | 2 |
| **translations-1/b-linux-v100-gpu-4-300gb** | generic-worker | 64.2.6 | insecure | edab196d7d | linux | amd64 | 1.22.2 | 967 | 967 |
| **translations-1/b-linux-v100-gpu-4-300gb-standard** | generic-worker | 64.2.6 | insecure | edab196d7d | linux | amd64 | 1.22.2 | 3 | 3 |
| **translations-1/bhearsum-translations-issue538** | generic-worker | 67.0.0 | multiuser | 881561bd9d | linux | amd64 | 1.22.4 | 0 | 0 |
| **translations-1/snakepit-2080ti** | generic-worker | 64.2.7 | insecure | 6f28121e31 | linux | amd64 | 1.22.2 | 0 | 0 |
| **translations-1/snakepit-quadro6k** | generic-worker | 64.2.7 | insecure | 6f28121e31 | linux | amd64 | 1.22.2 | 0 | 0 |
| **translations-1/snakepit-titanxp** | generic-worker | 64.2.7 | insecure | 6f28121e31 | linux | amd64 | 1.22.2 | 0 | 0 |


## Docker Worker

Total: `152`

Count by version:

| Version | Count |
| :--- | ---: |
| 38.0.5 | 127 |
| 48.3.0 | 1 |
| unknown version | 24 |


Count by image:

| Version | Count |
| :--- | ---: |
| projects/taskcluster-imaging/global/images/docker-firefoxci-gcp-l1-googlecompute-2024-06-14t19-52-20z | 78 |
| projects/fxci-production-level3-workers/global/images/docker-firefoxci-gcp-l3-googlecompute-2024-02-05t23-18-22z | 49 |
| projects/taskcluster-imaging/global/images/docker-worker-gcp-u14-04-2024-03-25 | 12 |
| projects/taskcluster-imaging/global/images/docker-worker-gcp-u14-04-2024-06-14 | 12 |
| projects/taskcluster-imaging/global/images/docker-firefoxci-gcp-lt-googlecompute-2023-04-13t21-30-28z | 1 |


| Worker Pool | Implementation | Version | Total Workers | Total Capacity |
| --- | --- | --- | ---: | ---: |
| **adhoc-1/b-linux-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **adhoc-1/decision-gcp** | docker-worker | 38.0.5 | 2 | 4 |
| **adhoc-1/images-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **adhoc-3/b-linux-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **adhoc-3/decision-gcp** | docker-worker | 38.0.5 | 2 | 4 |
| **adhoc-3/images-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **app-services-1/b-linux-gcp** | docker-worker | 38.0.5 | 176 | 176 |
| **app-services-1/decision-gcp** | docker-worker | 38.0.5 | 837 | 1674 |
| **app-services-1/images-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **app-services-3/b-linux-gcp** | docker-worker | 38.0.5 | 701 | 701 |
| **app-services-3/decision-gcp** | docker-worker | 38.0.5 | 27 | 54 |
| **app-services-3/images-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **ci-1/decision-gcp** | docker-worker | 38.0.5 | 64 | 128 |
| **ci-1/images-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **ci-3/decision-gcp** | docker-worker | 38.0.5 | 39 | 78 |
| **ci-3/images-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **ci-t/linux-gcp** | docker-worker | 38.0.5 | 527 | 527 |
| **code-analysis-1/linux-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **code-analysis-3/linux-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **code-coverage/bot-gcp** | docker-worker | 38.0.5 | 836 | 836 |
| **code-review/bot-gcp** | docker-worker | 38.0.5 | 521 | 612 |
| **comm-1/b-linux-gcp** | docker-worker | 38.0.5 | 710 | 710 |
| **comm-1/b-linux-large-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **comm-1/b-linux-xlarge-gcp** | docker-worker | 38.0.5 | 3 | 3 |
| **comm-1/decision-gcp** | docker-worker | 38.0.5 | 123 | 246 |
| **comm-1/images-gcp** | docker-worker | 38.0.5 | 29 | 29 |
| **comm-2/b-linux-gcp** | docker-worker | 38.0.5 | 4 | 4 |
| **comm-2/b-linux-large-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **comm-2/b-linux-xlarge-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **comm-2/decision-gcp** | docker-worker | 38.0.5 | 2 | 4 |
| **comm-2/images-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **comm-3/b-linux-gcp** | docker-worker | 38.0.5 | 5287 | 5287 |
| **comm-3/b-linux-large-gcp** | docker-worker | 38.0.5 | 4 | 4 |
| **comm-3/b-linux-xlarge-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **comm-3/decision-gcp** | docker-worker | 38.0.5 | 105 | 210 |
| **comm-3/images-gcp** | docker-worker | 38.0.5 | 23 | 23 |
| **comm-t/misc-gcp** | docker-worker | 38.0.5 | 2 | 16 |
| **comm-t/t-linux-large-gcp** | docker-worker | unknown version | 2 | 2 |
| **comm-t/t-linux-large-noscratch-gcp** | docker-worker | unknown version | 1447 | 1447 |
| **comm-t/t-linux-xlarge-gcp** | docker-worker | unknown version | 2 | 2 |
| **comm-t/t-linux-xlarge-noscratch-gcp** | docker-worker | unknown version | 195 | 195 |
| **comm-t/t-linux-xlarge-source-gcp** | docker-worker | unknown version | 986 | 986 |
| **comm-t/t-linux-xlarge-source-noscratch-gcp** | docker-worker | unknown version | 2 | 2 |
| **gecko-1/b-linux-gcp** | docker-worker | 38.0.5 | 18516 | 18516 |
| **gecko-1/b-linux-gcp-test-bug-1882320** | docker-worker | 38.0.5 | 2 | 2 |
| **gecko-1/b-linux-kvm-gcp** | docker-worker | 38.0.5 | 113 | 113 |
| **gecko-1/b-linux-large-gcp** | docker-worker | 38.0.5 | 76 | 76 |
| **gecko-1/b-linux-medium-gcp** | docker-worker | 38.0.5 | 11757 | 11757 |
| **gecko-1/b-linux-xlarge-gcp** | docker-worker | 38.0.5 | 705 | 705 |
| **gecko-1/b-linux-xlarge-gcp-bug1797804-c2** | docker-worker | 38.0.5 | 2 | 2 |
| **gecko-1/decision-gcp** | docker-worker | 38.0.5 | 1313 | 2626 |
| **gecko-1/images-gcp** | docker-worker | 38.0.5 | 131 | 131 |
| **gecko-2/b-linux-gcp** | docker-worker | 38.0.5 | 1815 | 1815 |
| **gecko-2/b-linux-kvm-gcp** | docker-worker | 38.0.5 | 7 | 7 |
| **gecko-2/b-linux-large-gcp** | docker-worker | 38.0.5 | 63 | 63 |
| **gecko-2/b-linux-xlarge-gcp** | docker-worker | 38.0.5 | 58 | 58 |
| **gecko-2/decision-gcp** | docker-worker | 38.0.5 | 31 | 62 |
| **gecko-2/images-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **gecko-3/b-linux-gcp** | docker-worker | 38.0.5 | 47253 | 47253 |
| **gecko-3/b-linux-gcp-bug1874006** | docker-worker | 38.0.5 | 2 | 2 |
| **gecko-3/b-linux-kvm-gcp** | docker-worker | 38.0.5 | 215 | 215 |
| **gecko-3/b-linux-large-gcp** | docker-worker | 38.0.5 | 180 | 180 |
| **gecko-3/b-linux-medium-gcp** | docker-worker | 38.0.5 | 11501 | 11501 |
| **gecko-3/b-linux-xlarge-gcp** | docker-worker | 38.0.5 | 1820 | 1820 |
| **gecko-3/decision-gcp** | docker-worker | 38.0.5 | 1358 | 2716 |
| **gecko-3/images-gcp** | docker-worker | 38.0.5 | 54 | 54 |
| **gecko-3/t-linux-xlarge-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **gecko-3/t-linux-xlarge-noscratch-gcp** | docker-worker | 38.0.5 | 293 | 293 |
| **gecko-t/misc-gcp** | docker-worker | 38.0.5 | 681 | 5448 |
| **gecko-t/t-linux-kvm-gcp** | docker-worker | unknown version | 2 | 2 |
| **gecko-t/t-linux-kvm-gcp-bug1862675** | docker-worker | unknown version | 2 | 2 |
| **gecko-t/t-linux-kvm-noscratch-gcp** | docker-worker | unknown version | 35575 | 35575 |
| **gecko-t/t-linux-kvm-noscratch-gcp-bug1862675** | docker-worker | unknown version | 2 | 2 |
| **gecko-t/t-linux-large-gcp** | docker-worker | unknown version | 1381 | 1381 |
| **gecko-t/t-linux-large-hostub2204-gcp** | docker-worker | 48.3.0 | 2 | 2 |
| **gecko-t/t-linux-large-noscratch-gcp** | docker-worker | unknown version | 78568 | 78568 |
| **gecko-t/t-linux-xlarge-gcp** | docker-worker | unknown version | 199 | 199 |
| **gecko-t/t-linux-xlarge-gcp-b1862675** | docker-worker | unknown version | 2 | 2 |
| **gecko-t/t-linux-xlarge-noscratch-gcp** | docker-worker | unknown version | 47536 | 47536 |
| **gecko-t/t-linux-xlarge-ns-gcp-b1862675** | docker-worker | unknown version | 2 | 2 |
| **gecko-t/t-linux-xlarge-source-gcp** | docker-worker | unknown version | 23503 | 23503 |
| **gecko-t/t-linux-xlarge-source-gcp-b1862675** | docker-worker | unknown version | 2 | 2 |
| **gecko-t/t-linux-xlarge-source-noscratch-gcp** | docker-worker | unknown version | 2 | 2 |
| **gecko-t/t-linux-xlarge-source-ns-gcp-b1862675** | docker-worker | unknown version | 2 | 2 |
| **glean-1/b-linux-gcp** | docker-worker | 38.0.5 | 37 | 37 |
| **glean-1/decision-gcp** | docker-worker | 38.0.5 | 838 | 1676 |
| **glean-1/images-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **glean-3/b-linux-gcp** | docker-worker | 38.0.5 | 38 | 38 |
| **glean-3/decision-gcp** | docker-worker | 38.0.5 | 9 | 18 |
| **glean-3/images-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **infra/build-decision** | docker-worker | 38.0.5 | 8 | 64 |
| **mobile-1/b-linux-gcp** | docker-worker | 38.0.5 | 141 | 141 |
| **mobile-1/decision-gcp** | docker-worker | 38.0.5 | 807 | 1614 |
| **mobile-1/images-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **mobile-3/b-linux-gcp** | docker-worker | 38.0.5 | 295 | 295 |
| **mobile-3/decision-gcp** | docker-worker | 38.0.5 | 110 | 220 |
| **mobile-3/images-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **mozilla-1/b-linux-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **mozilla-1/decision-gcp** | docker-worker | 38.0.5 | 818 | 1636 |
| **mozilla-1/images-gcp** | docker-worker | 38.0.5 | 11 | 11 |
| **mozilla-t/t-linux-large-gcp** | docker-worker | unknown version | 50 | 50 |
| **mozilla-t/t-linux-large-noscratch-gcp** | docker-worker | unknown version | 2 | 2 |
| **mozillaonline-1/b-linux-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **mozillaonline-3/b-linux-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **mozillavpn-1/b-linux-gcp** | docker-worker | 38.0.5 | 104 | 104 |
| **mozillavpn-1/b-linux-large-gcp** | docker-worker | 38.0.5 | 77 | 77 |
| **mozillavpn-1/decision-gcp** | docker-worker | 38.0.5 | 835 | 1670 |
| **mozillavpn-1/images-gcp** | docker-worker | 38.0.5 | 12 | 12 |
| **mozillavpn-3/b-linux-gcp** | docker-worker | 38.0.5 | 122 | 122 |
| **mozillavpn-3/b-linux-large-gcp** | docker-worker | 38.0.5 | 115 | 115 |
| **mozillavpn-3/decision-gcp** | docker-worker | 38.0.5 | 14 | 28 |
| **mozillavpn-3/images-gcp** | docker-worker | 38.0.5 | 12 | 12 |
| **nss-1/decision-gcp** | docker-worker | 38.0.5 | 2 | 4 |
| **nss-1/images-gcp** | docker-worker | 38.0.5 | 3 | 3 |
| **nss-1/linux-gcp** | docker-worker | 38.0.5 | 495 | 495 |
| **nss-3/images-gcp** | docker-worker | 38.0.5 | 3 | 3 |
| **nss-3/linux-gcp** | docker-worker | 38.0.5 | 407 | 407 |
| **nss-t/t-linux-xlarge-gcp** | docker-worker | unknown version | 955 | 955 |
| **nss-t/t-linux-xlarge-noscratch-gcp** | docker-worker | unknown version | 2 | 2 |
| **releng-1/b-linux-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **releng-1/decision-gcp** | docker-worker | 38.0.5 | 16 | 32 |
| **releng-1/linux-gcp** | docker-worker | 38.0.5 | 36 | 36 |
| **releng-3/b-linux-gcp** | docker-worker | 38.0.5 | 5 | 5 |
| **releng-3/decision-gcp** | docker-worker | 38.0.5 | 12 | 24 |
| **releng-3/linux-gcp** | docker-worker | 38.0.5 | 49 | 49 |
| **releng-t/linux-gcp** | docker-worker | 38.0.5 | 82 | 82 |
| **relops-3/decision-gcp** | docker-worker | 38.0.5 | 2 | 4 |
| **scriptworker-1/b-linux-gcp** | docker-worker | 38.0.5 | 135 | 135 |
| **scriptworker-1/decision-gcp** | docker-worker | 38.0.5 | 32 | 64 |
| **scriptworker-1/images-gcp** | docker-worker | 38.0.5 | 16 | 16 |
| **scriptworker-3/b-linux-gcp** | docker-worker | 38.0.5 | 163 | 163 |
| **scriptworker-3/decision-gcp** | docker-worker | 38.0.5 | 15 | 30 |
| **scriptworker-3/images-gcp** | docker-worker | 38.0.5 | 12 | 12 |
| **taskgraph-1/decision-gcp** | docker-worker | 38.0.5 | 7 | 14 |
| **taskgraph-1/images-gcp** | docker-worker | 38.0.5 | 13 | 13 |
| **taskgraph-3/decision-gcp** | docker-worker | 38.0.5 | 5 | 10 |
| **taskgraph-3/images-gcp** | docker-worker | 38.0.5 | 16 | 16 |
| **taskgraph-t/linux-gcp** | docker-worker | 38.0.5 | 64 | 64 |
| **translations-1/b-linux-large-gcp** | docker-worker | 38.0.5 | 505 | 505 |
| **translations-1/b-linux-large-gcp-1tb-32-256** | docker-worker | 38.0.5 | 2 | 2 |
| **translations-1/b-linux-large-gcp-1tb-32-256-standard** | docker-worker | 38.0.5 | 2 | 2 |
| **translations-1/b-linux-large-gcp-1tb-64-512** | docker-worker | 38.0.5 | 2 | 2 |
| **translations-1/b-linux-large-gcp-1tb-64-512-standard** | docker-worker | 38.0.5 | 2 | 2 |
| **translations-1/b-linux-large-gcp-300gb** | docker-worker | 38.0.5 | 570 | 570 |
| **translations-1/decision-gcp** | docker-worker | 38.0.5 | 48 | 96 |
| **translations-1/images-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **xpi-1/b-linux-gcp** | docker-worker | 38.0.5 | 229 | 229 |
| **xpi-1/decision-gcp** | docker-worker | 38.0.5 | 74 | 148 |
| **xpi-1/images-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **xpi-3/b-linux-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **xpi-3/decision-gcp** | docker-worker | 38.0.5 | 2 | 4 |
| **xpi-3/images-gcp** | docker-worker | 38.0.5 | 2 | 2 |


## Script Worker

Total: `31`



| Worker Pool | Implementation | Version | Total Workers | Total Capacity |
| --- | --- | --- | ---: | ---: |
| **scriptworker-k8s/app-services-3-signing** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/comm-3-balrog** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/comm-3-beetmover** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/comm-3-bouncer** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/comm-3-pushflatpak** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/comm-3-shipit** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/comm-3-signing** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/comm-3-tree** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/comm-t-signing** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/gecko-1-beetmover** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/gecko-3-balrog** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/gecko-3-beetmover** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/gecko-3-bouncer** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/gecko-3-pushapk** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/gecko-3-shipit** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/gecko-3-signing** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/gecko-3-tree** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/gecko-t-signing** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/mobile-3-bitrise** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/mobile-3-signing** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/mobile-t-signing** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/mozillavpn-3-signing** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
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
| projects/taskcluster-imaging/global/images/generic-2204-wayland-vm-gcp-googlecompute-2023-09-22t17-39-37z | 2 |


| Worker Pool | Implementation | Version | Total Workers | Total Capacity |
| --- | --- | --- | ---: | ---: |
| **built-in/fail** |  | No artifacts found | 0 | 0 |
| **built-in/succeed** |  | No artifacts found | 0 | 0 |
| **gecko-t/t-linux-vm-2204-wayland** |  | No artifacts found | 1121 | 1121 |
| **gecko-t/t-linux-vm-2204-wayland-snap** |  | No artifacts found | 2 | 2 |
| **scriptworker-prov-v1/mac-notarization-poller** |  | No artifacts found | 0 | 0 |
| **scriptworker-prov-v1/tb-mac-notarization-poller** |  | No artifacts found | 0 | 0 |


## Version not determined [^2]

Total: `24`


Count by image:

| Version | Count |
| :--- | ---: |
| /subscriptions/108d46d5-fe9b-4850-9a7d-8c914aa6c1f0/resourceGroups/rg-packer-through-cib/providers/Microsoft.Compute/images/win11-64-2009-tc-centralus-win11-22h2-avd-alpha | 1 |
| /subscriptions/108d46d5-fe9b-4850-9a7d-8c914aa6c1f0/resourceGroups/rg-packer-through-cib/providers/Microsoft.Compute/images/win11-64-2009-westus-win11-22h2-avd-alpha | 1 |
| /subscriptions/108d46d5-fe9b-4850-9a7d-8c914aa6c1f0/resourceGroups/rg-packer-through-cib/providers/Microsoft.Compute/images/win2012r2-64-vs-l1-centralus-2012-R2-Datacenter-alpha | 1 |
| projects/taskcluster-imaging/global/images/gw-fxci-gcp-l1-gui-googlecompute-2024-05-31t15-18-14z | 1 |
| /subscriptions/108d46d5-fe9b-4850-9a7d-8c914aa6c1f0/resourceGroups/rg-packer-through-cib/providers/Microsoft.Compute/images/win11-64-2009-uksouth-win11-22h2-avd-alpha | 2 |
| /subscriptions/108d46d5-fe9b-4850-9a7d-8c914aa6c1f0/resourceGroups/rg-packer-through-cib/providers/Microsoft.Compute/images/win11-64-2009-eastus-win11-22h2-avd-69a7ecf | 1 |
| /subscriptions/108d46d5-fe9b-4850-9a7d-8c914aa6c1f0/resourceGroups/rg-packer-through-cib/providers/Microsoft.Compute/images/win10-64-2009-uksouth-win10-22h2-avd-alpha | 1 |
| /subscriptions/108d46d5-fe9b-4850-9a7d-8c914aa6c1f0/resourceGroups/rg-packer-through-cib/providers/Microsoft.Compute/images/win2022-64-2009-rd-centralus-2022-datacenter-azure-edition-alpha | 1 |
| /subscriptions/108d46d5-fe9b-4850-9a7d-8c914aa6c1f0/resourceGroups/rg-packer-through-cib/providers/Microsoft.Compute/images/win2012r2-64-l1-centralus-2012-R2-Datacenter-alpha | 2 |
| /subscriptions/108d46d5-fe9b-4850-9a7d-8c914aa6c1f0/resourceGroups/rg-packer-through-cib/providers/Microsoft.Compute/images/win2022-64-2009-centralus-2022-datacenter-azure-edition-b1285eb | 1 |
| /subscriptions/108d46d5-fe9b-4850-9a7d-8c914aa6c1f0/resourceGroups/rg-packer-through-cib/providers/Microsoft.Compute/images/win2022-64-2009-centralus-2022-datacenter-azure-edition-alpha | 3 |
| /subscriptions/108d46d5-fe9b-4850-9a7d-8c914aa6c1f0/resourceGroups/rg-packer-through-cib/providers/Microsoft.Compute/images/win2012r2-64-vs-py2-l1-centralus-2012-R2-Datacenter-alpha | 1 |
| /subscriptions/108d46d5-fe9b-4850-9a7d-8c914aa6c1f0/resourceGroups/rg-packer-through-cib/providers/Microsoft.Compute/images/win2012r2-64-vs-l1-centralus-2012-R2-Datacenter-21c6826 | 1 |
| /subscriptions/108d46d5-fe9b-4850-9a7d-8c914aa6c1f0/resourceGroups/rg-packer-through-cib/providers/Microsoft.Compute/images/win10-64-2009-eastus-win10-22h2-avd-alpha | 1 |
| /subscriptions/a30e97ab-734a-4f3b-a0e4-c51c0bff0701/resourceGroups/rg-packer-through-cib/providers/Microsoft.Compute/images/trusted-win2022-64-2009-centralus-2022-datacenter-azure-edition-b1285eb | 1 |
| /subscriptions/108d46d5-fe9b-4850-9a7d-8c914aa6c1f0/resourceGroups/rg-packer-through-cib/providers/Microsoft.Compute/images/masterwayz-win2012test-image-20221122010013-alpha | 1 |
|  | 3 |
| /subscriptions/108d46d5-fe9b-4850-9a7d-8c914aa6c1f0/resourceGroups/rg-packer-through-cib/providers/Microsoft.Compute/images/win10-64-2009-ukwest-win10-22h2-avd-alpha | 1 |


| Worker Pool | Implementation | Version | Total Workers | Total Capacity |
| --- | --- | --- | ---: | ---: |
| **gecko-1/b-win2012-azure-alpha** |  | Version not determined; task not (yet) claimed | 45 | 45 |
| **gecko-1/b-win2012-azure-beta** |  | Version not determined; task not (yet) claimed | 45 | 45 |
| **gecko-1/b-win2022-alpha** |  | Version not determined; task not (yet) claimed | 7 | 7 |
| **gecko-1/b-win2022-rd-alpha** |  | Version not determined; task not (yet) claimed | 45 | 45 |
| **gecko-t/win10-64-2009-alpha** |  | Version not determined; task not (yet) claimed | 8 | 8 |
| **gecko-t/win10-64-2009-gpu-alpha** |  | Version not determined; task not (yet) claimed | 7 | 7 |
| **gecko-t/win10-64-2009-source-alpha** |  | Version not determined; task not (yet) claimed | 45 | 45 |
| **gecko-t/win11-64-2009-alpha** |  | Version not determined; task not (yet) claimed | 7 | 7 |
| **gecko-t/win11-64-2009-alpha-test** |  | Version not determined; task not (yet) claimed | 44 | 44 |
| **gecko-t/win11-64-2009-gpu-alpha** |  | Version not determined; task not (yet) claimed | 7 | 7 |
| **gecko-t/win11-64-2009-source-alpha** |  | Version not determined; task not (yet) claimed | 6 | 6 |
| **gecko-t/win11-64-2009-tc-alpha** |  | Version not determined; task not (yet) claimed | 46 | 46 |
| **mozilla-t/t-dagger-experimental** |  | Version not determined; task not (yet) claimed | 30 | 30 |
| **mozillavpn-1/b-win2012-azure** |  | Version not determined; task not (yet) claimed | 46 | 46 |
| **mozillavpn-1/b-win2012-azure-alpha** |  | Version not determined; task not (yet) claimed | 46 | 46 |
| **mozillavpn-1/b-win2022-alpha** |  | Version not determined; task not (yet) claimed | 7 | 7 |
| **mozillavpn-1/masterwayz-win2012-azure** |  | Version not determined; task not (yet) claimed | 46 | 46 |
| **nss-1/b-win2012-azure-alpha** |  | Version not determined; task not (yet) claimed | 43 | 43 |
| **nss-1/b-win2022-alpha** |  | Version not determined; task not (yet) claimed | 7 | 7 |
| **proj-autophone/gecko-t-bitbar-gw-perf-a51** |  | Version not determined; task not (yet) claimed | 0 | 0 |
| **releng-hardware/gecko-t-osx-1400-r8-staging** |  | Version not determined; task not (yet) claimed | 0 | 0 |
| **releng-hardware/nss-3-b-osx-1015** |  | Version not determined; task not (yet) claimed | 0 | 0 |
| **relops-1/b-win2022** |  | Version not determined; task not (yet) claimed | 232 | 232 |
| **relops-3/b-win2022** |  | Version not determined; task not (yet) claimed | 243 | 243 |



[^1]: Those are the pools whose tasks were claimed and resolved by a worker as expected, but the worker did not publish either artifact `public/logs/live_backing.log` nor `public/logs/chain_of_trust.log`, which is the source used to identify the worker implementation.

[^2]: Probing task remains pending after two hours. Those are the pools that were not able to start any worker to claim the task within two hours.



# Worker Pool Versions


## Generic Worker

Total: `53`

Count by version:

| Version | Count |
| :--- | ---: |
| 14.1.2 | 1 |
| 15.1.4 | 1 |
| 16.2.0 | 3 |
| 16.5.1 | 1 |
| 30.0.2 | 3 |
| 36.0.0 | 2 |
| 40.0.3 | 3 |
| 43.2.0 | 3 |
| 44.0.0 | 1 |
| 44.23.2 | 1 |
| 45.0.0 | 1 |
| 54.4.1 | 24 |
| 55.1.1 | 4 |
| 55.3.4 | 4 |
| 57.0.1 | 1 |


Count by image:

| Version | Count |
| :--- | ---: |
| /subscriptions/a30e97ab-734a-4f3b-a0e4-c51c0bff0701/resourceGroups/rg-packer-through-cib/providers/Microsoft.Compute/images/trusted-win2012r2-64-l3-centralus-2012-R2-Datacenter-7cde253 | 1 |
| /subscriptions/a30e97ab-734a-4f3b-a0e4-c51c0bff0701/resourceGroups/rg-packer-through-cib/providers/Microsoft.Compute/images/trusted-win2022-64-2009-centralus-2022-datacenter-azure-edition-d5dee60 | 3 |
| unknown | 1 |
| /subscriptions/108d46d5-fe9b-4850-9a7d-8c914aa6c1f0/resourceGroups/rg-packer-through-cib/providers/Microsoft.Compute/images/win11-64-2009-centralus-win11-22h2-avd-d322b59 | 1 |
|  | 26 |
| /subscriptions/108d46d5-fe9b-4850-9a7d-8c914aa6c1f0/resourceGroups/rg-packer-through-cib/providers/Microsoft.Compute/images/win10-64-2009-eastus-win10-22h2-avd-f38a3d4 | 1 |
| /subscriptions/a30e97ab-734a-4f3b-a0e4-c51c0bff0701/resourceGroups/rg-packer-through-cib/providers/Microsoft.Compute/images/trusted-win2012r2-64-vs-l3-centralus-2012-R2-Datacenter-7cde253 | 1 |
| /subscriptions/108d46d5-fe9b-4850-9a7d-8c914aa6c1f0/resourceGroups/rg-packer-through-cib/providers/Microsoft.Compute/images/win11-64-2009-eastus-win11-22h2-avd-d322b59 | 2 |
| /subscriptions/108d46d5-fe9b-4850-9a7d-8c914aa6c1f0/resourceGroups/rg-packer-through-cib/providers/Microsoft.Compute/images/win11-64-2009-canadacentral-win11-22h2-avd-d322b59 | 2 |
| /subscriptions/108d46d5-fe9b-4850-9a7d-8c914aa6c1f0/resourceGroups/rg-packer-through-cib/providers/Microsoft.Compute/images/win11-64-2009-uksouth-win11-22h2-avd-alpha | 2 |
| projects/taskcluster-imaging/global/images/generic-translations-gcp-googlecompute-2023-11-01t20-12-18z | 1 |
| /subscriptions/108d46d5-fe9b-4850-9a7d-8c914aa6c1f0/resourceGroups/rg-packer-through-cib/providers/Microsoft.Compute/images/win2022-64-2009-centralus-2022-datacenter-azure-edition-531c001 | 4 |
| /subscriptions/108d46d5-fe9b-4850-9a7d-8c914aa6c1f0/resourceGroups/rg-packer-through-cib/providers/Microsoft.Compute/images/win10-64-2009-canadacentral-win10-22h2-avd-f38a3d4 | 2 |
| ami-07db1c26bda143e33 | 1 |
| /subscriptions/a30e97ab-734a-4f3b-a0e4-c51c0bff0701/resourceGroups/rg-packer-through-cib/providers/Microsoft.Compute/images/trusted-win2012r2-64-vs-py2-l3-centralus-2012-R2-Datacenter-7aa76c6 | 1 |
| projects/taskcluster-imaging/global/images/generic-translations-gcp-googlecompute-2023-09-12t21-35-55z | 4 |


| Worker Pool | Implementation | Version | Engine | Revision | OS | Arch | GO | Total Workers | Total Capacity |
| --- | --- | --- | --- | --- | --- | --- | --- | ---: | ---: |
| **bitbar/gecko-t-win64-aarch64-laptop** | generic-worker | 14.1.2 | <UNKNOWN> | 13118c4c1b | windows | 386 | 1.10.8 | 0 | 0 |
| **comm-1/b-win2022** | generic-worker | 54.4.1 | multiuser | a38a78d98c | windows | amd64 | 1.20.6 | 4 | 4 |
| **comm-2/b-win2022** | generic-worker | 54.4.1 | multiuser | a38a78d98c | windows | amd64 | 1.20.6 | 5 | 5 |
| **comm-3/b-win2012-azure** | generic-worker | 43.2.0 | multiuser | 53fa7b79bc | windows | amd64 | 1.16.3 | 86 | 86 |
| **comm-3/b-win2022** | generic-worker | 54.4.1 | multiuser | a38a78d98c | windows | amd64 | 1.20.6 | 568 | 568 |
| **comm-t/win11-64-2009** | generic-worker | 54.4.1 | multiuser | a38a78d98c | windows | amd64 | 1.20.6 | 5286 | 5286 |
| **gecko-2/b-win2022** | generic-worker | 54.4.1 | multiuser | a38a78d98c | windows | amd64 | 1.20.6 | 73 | 73 |
| **gecko-3/b-win2022** | generic-worker | 54.4.1 | multiuser | a38a78d98c | windows | amd64 | 1.20.6 | 5696 | 5696 |
| **gecko-t/win10-64-2009** | generic-worker | 54.4.1 | multiuser | a38a78d98c | windows | amd64 | 1.20.6 | 5995 | 5995 |
| **gecko-t/win10-64-2009-gpu** | generic-worker | 54.4.1 | multiuser | a38a78d98c | windows | amd64 | 1.20.6 | 743 | 743 |
| **gecko-t/win10-64-2009-source** | generic-worker | 54.4.1 | multiuser | a38a78d98c | windows | amd64 | 1.20.6 | 4 | 4 |
| **gecko-t/win11-64-2009** | generic-worker | 54.4.1 | multiuser | a38a78d98c | windows | amd64 | 1.20.6 | 22557 | 22557 |
| **gecko-t/win11-64-2009-alpha** | generic-worker | 54.4.1 | multiuser | a38a78d98c | windows | amd64 | 1.20.6 | 4 | 4 |
| **gecko-t/win11-64-2009-gpu** | generic-worker | 54.4.1 | multiuser | a38a78d98c | windows | amd64 | 1.20.6 | 21148 | 21148 |
| **gecko-t/win11-64-2009-source-alpha** | generic-worker | 54.4.1 | multiuser | a38a78d98c | windows | amd64 | 1.20.6 | 120 | 120 |
| **gecko-t/win11-64-2009-ssd** | generic-worker | 54.4.1 | multiuser | a38a78d98c | windows | amd64 | 1.20.6 | 2254 | 2254 |
| **gecko-t/win11-64-2009-ssd-gpu** | generic-worker | 54.4.1 | multiuser | a38a78d98c | windows | amd64 | 1.20.6 | 1925 | 1925 |
| **localprovisioner/nss-macos-10-12** | generic-worker | 15.1.4 | simple | c407e45e3f | darwin | amd64 | 1.10.8 | 0 | 0 |
| **localprovisioner/nss-macos-m1** | generic-worker | 30.0.2 | simple | 6fdba0dad3 | darwin | arm64 | 1.16.4 | 0 | 0 |
| **mozillavpn-1/b-win2022** | generic-worker | 54.4.1 | multiuser | a38a78d98c | windows | amd64 | 1.20.6 | 194 | 194 |
| **mozillavpn-3/b-win2012-azure** | generic-worker | 43.2.0 | multiuser | 53fa7b79bc | windows | amd64 | 1.16.3 | 13 | 13 |
| **mozillavpn-3/b-win2022** | generic-worker | 54.4.1 | multiuser | a38a78d98c | windows | amd64 | 1.20.6 | 41 | 41 |
| **nss-3/b-win2012-azure** | generic-worker | 43.2.0 | multiuser | 53fa7b79bc | windows | amd64 | 1.16.3 | 29 | 29 |
| **performance-hardware/gecko-t-fxrecorder** | generic-worker | 44.0.0 | multiuser | faf6f319c0 | windows | amd64 | 1.16.3 | 0 | 0 |
| **proj-autophone/gecko-t-bitbar-gw-perf-a51** | generic-worker | 36.0.0 | simple | b9cb11293f | linux | amd64 | 1.13.7 | 0 | 0 |
| **proj-autophone/gecko-t-bitbar-gw-unit-p5** | generic-worker | 36.0.0 | simple | b9cb11293f | linux | amd64 | 1.13.7 | 0 | 0 |
| **releng-hardware/applicationservices-b-1-osx1015** | generic-worker | 54.4.1 | multiuser | a38a78d98c | darwin | amd64 | 1.20.6 | 0 | 0 |
| **releng-hardware/applicationservices-b-3-osx1015** | generic-worker | 54.4.1 | multiuser | a38a78d98c | darwin | amd64 | 1.20.6 | 0 | 0 |
| **releng-hardware/gecko-1-b-osx-1015** | generic-worker | 40.0.3 | multiuser | 5f7a31ac6e | darwin | amd64 | 1.15.6 | 0 | 0 |
| **releng-hardware/gecko-1-b-osx-1015-staging** | generic-worker | 54.4.1 | multiuser | a38a78d98c | darwin | amd64 | 1.20.6 | 0 | 0 |
| **releng-hardware/gecko-3-b-osx-1015** | generic-worker | 55.3.4 | multiuser | a71f300fcb | darwin | amd64 | 1.21.1 | 0 | 0 |
| **releng-hardware/gecko-t-linux-talos-1804** | generic-worker | 44.23.2 | simple | da8da74563 | linux | amd64 | 1.18.5 | 0 | 0 |
| **releng-hardware/gecko-t-osx-1015-power** | generic-worker | 40.0.3 | simple | 5f7a31ac6e | darwin | amd64 | 1.15.6 | 0 | 0 |
| **releng-hardware/gecko-t-osx-1015-r8-staging** | generic-worker | 40.0.3 | simple | 5f7a31ac6e | darwin | amd64 | 1.15.6 | 0 | 0 |
| **releng-hardware/gecko-t-osx-1100-m1** | generic-worker | 55.3.4 | simple | a71f300fcb | darwin | arm64 | 1.21.1 | 0 | 0 |
| **releng-hardware/gecko-t-osx-1100-m1-loaner** | generic-worker | 30.0.2 | simple | 6fdba0dad3 | darwin | arm64 | 1.16.4 | 0 | 0 |
| **releng-hardware/gecko-t-osx-1100-m1-loaner-bug1829730** | generic-worker | 30.0.2 | simple | 6fdba0dad3 | darwin | arm64 | 1.16.4 | 0 | 0 |
| **releng-hardware/gecko-t-osx-1100-m1-staging** | generic-worker | 55.3.4 | simple | a71f300fcb | darwin | arm64 | 1.21.1 | 0 | 0 |
| **releng-hardware/gecko-t-osx-1300-m2** | generic-worker | 55.3.4 | simple | a71f300fcb | darwin | arm64 | 1.21.1 | 0 | 0 |
| **releng-hardware/gecko-t-win-talos** | generic-worker | 16.2.0 | multiuser | b321a877c8 | windows | amd64 | 1.10.8 | 0 | 0 |
| **releng-hardware/gecko-t-win10-64-dev** | generic-worker | 16.2.0 | multiuser | b321a877c8 | windows | amd64 | 1.10.8 | 0 | 0 |
| **releng-hardware/gecko-t-win10-64-ref-hw** | generic-worker | 16.2.0 | multiuser | b321a877c8 | windows | amd64 | 1.10.8 | 0 | 0 |
| **releng-hardware/gecko-t-win7-32-hw** | generic-worker | 45.0.0 | multiuser | 988e8100b3 | windows | 386 | 1.19.3 | 0 | 0 |
| **releng-hardware/mozillavpn-b-1-osx** | generic-worker | 54.4.1 | multiuser | a38a78d98c | darwin | amd64 | 1.20.6 | 0 | 0 |
| **releng-hardware/mozillavpn-b-3-osx** | generic-worker | 54.4.1 | multiuser | a38a78d98c | darwin | amd64 | 1.20.6 | 0 | 0 |
| **releng-hardware/win11-64-2009-hw-ref** | generic-worker | 54.4.1 | multiuser | a38a78d98c | windows | amd64 | 1.20.6 | 0 | 0 |
| **releng-hardware/win11-64-2009-hw-ref-alpha** | generic-worker | 54.4.1 | multiuser | a38a78d98c | windows | amd64 | 1.20.6 | 0 | 0 |
| **relops-3/win2019** | generic-worker | 16.5.1 | multiuser | 02bc4c9cd7 | windows | amd64 | 1.10.8 | 4 | 4 |
| **translations-1/b-linux-aerickson-test** | generic-worker | 57.0.1 | simple | fe45895655 | linux | amd64 | 1.21.3 | 3 | 3 |
| **translations-1/b-linux-v100-gpu** | generic-worker | 55.1.1 | simple | b83fffd4bd | linux | amd64 | 1.21.1 | 124 | 124 |
| **translations-1/b-linux-v100-gpu-4** | generic-worker | 55.1.1 | simple | b83fffd4bd | linux | amd64 | 1.21.1 | 80 | 80 |
| **translations-1/b-linux-v100-gpu-4-1tb** | generic-worker | 55.1.1 | simple | b83fffd4bd | linux | amd64 | 1.21.1 | 82 | 82 |
| **translations-1/b-linux-v100-gpu-4-300gb** | generic-worker | 55.1.1 | simple | b83fffd4bd | linux | amd64 | 1.21.1 | 104 | 104 |


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
| **app-services-1/b-linux-gcp** | docker-worker | 38.0.5 | 176 | 176 |
| **app-services-1/decision-gcp** | docker-worker | 38.0.5 | 582 | 1164 |
| **app-services-1/images-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **app-services-3/b-linux-gcp** | docker-worker | 38.0.5 | 673 | 673 |
| **app-services-3/decision-gcp** | docker-worker | 38.0.5 | 24 | 48 |
| **app-services-3/images-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **ci-1/decision-gcp** | docker-worker | 38.0.5 | 8 | 16 |
| **ci-1/images-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **ci-3/decision-gcp** | docker-worker | 38.0.5 | 10 | 20 |
| **ci-3/images-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **ci-t/linux-gcp** | docker-worker | 38.0.5 | 78 | 78 |
| **code-analysis-1/linux-gcp** | docker-worker | 38.0.5 | 19 | 19 |
| **code-analysis-3/linux-gcp** | docker-worker | 38.0.5 | 17 | 17 |
| **code-coverage/bot-gcp** | docker-worker | 38.0.5 | 825 | 825 |
| **code-review/bot-gcp** | docker-worker | 38.0.5 | 2801 | 2801 |
| **comm-1/b-linux-gcp** | docker-worker | 38.0.5 | 445 | 445 |
| **comm-1/b-linux-large-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **comm-1/b-linux-xlarge-gcp** | docker-worker | 38.0.5 | 3 | 3 |
| **comm-1/decision-gcp** | docker-worker | 38.0.5 | 146 | 292 |
| **comm-1/images-gcp** | docker-worker | 38.0.5 | 13 | 13 |
| **comm-2/b-linux-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **comm-2/b-linux-large-gcp** | docker-worker | 38.0.5 | 3 | 3 |
| **comm-2/b-linux-xlarge-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **comm-2/decision-gcp** | docker-worker | 38.0.5 | 2 | 4 |
| **comm-2/images-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **comm-3/b-linux-gcp** | docker-worker | 38.0.5 | 3205 | 3205 |
| **comm-3/b-linux-large-gcp** | docker-worker | 38.0.5 | 5 | 5 |
| **comm-3/b-linux-xlarge-gcp** | docker-worker | 38.0.5 | 5 | 5 |
| **comm-3/decision-gcp** | docker-worker | 38.0.5 | 86 | 172 |
| **comm-3/images-gcp** | docker-worker | 38.0.5 | 14 | 14 |
| **comm-t/misc-gcp** | docker-worker | 38.0.5 | 3 | 24 |
| **comm-t/t-linux-large-gcp** | docker-worker | unknown version | 1503 | 1503 |
| **comm-t/t-linux-xlarge-gcp** | docker-worker | unknown version | 243 | 243 |
| **comm-t/t-linux-xlarge-source-gcp** | docker-worker | unknown version | 901 | 901 |
| **gecko-1/b-linux-gcp** | docker-worker | 38.0.5 | 6547 | 6547 |
| **gecko-1/b-linux-gcp-bug1797804-n2** | docker-worker | 38.0.5 | 2 | 2 |
| **gecko-1/b-linux-large-gcp** | docker-worker | 38.0.5 | 73 | 73 |
| **gecko-1/b-linux-xlarge-gcp** | docker-worker | 38.0.5 | 471 | 471 |
| **gecko-1/b-linux-xlarge-gcp-bug1797804-c2** | docker-worker | 38.0.5 | 2 | 2 |
| **gecko-1/decision-gcp** | docker-worker | 38.0.5 | 1229 | 2458 |
| **gecko-1/images-gcp** | docker-worker | 38.0.5 | 13 | 13 |
| **gecko-2/b-linux-gcp** | docker-worker | 38.0.5 | 885 | 885 |
| **gecko-2/b-linux-large-gcp** | docker-worker | 38.0.5 | 19 | 19 |
| **gecko-2/b-linux-xlarge-gcp** | docker-worker | 38.0.5 | 29 | 29 |
| **gecko-2/decision-gcp** | docker-worker | 38.0.5 | 20 | 40 |
| **gecko-2/images-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **gecko-3/b-linux-gcp** | docker-worker | 38.0.5 | 20476 | 20476 |
| **gecko-3/b-linux-large-gcp** | docker-worker | 38.0.5 | 152 | 152 |
| **gecko-3/b-linux-xlarge-gcp** | docker-worker | 38.0.5 | 1324 | 1324 |
| **gecko-3/decision-gcp** | docker-worker | 38.0.5 | 1443 | 2886 |
| **gecko-3/images-gcp** | docker-worker | 38.0.5 | 26 | 26 |
| **gecko-3/t-linux-xlarge-gcp** | docker-worker | 38.0.5 | 225 | 225 |
| **gecko-t/misc-gcp** | docker-worker | 38.0.5 | 355 | 2840 |
| **gecko-t/t-linux-kvm-gcp** | docker-worker | unknown version | 30758 | 30758 |
| **gecko-t/t-linux-large-gcp** | docker-worker | unknown version | 108296 | 108296 |
| **gecko-t/t-linux-large-hostub2204-gcp** | docker-worker | 48.3.0 | 2 | 2 |
| **gecko-t/t-linux-xlarge-gcp** | docker-worker | unknown version | 46874 | 46874 |
| **gecko-t/t-linux-xlarge-source-gcp** | docker-worker | unknown version | 10407 | 10407 |
| **glean-1/b-linux-gcp** | docker-worker | 38.0.5 | 1170 | 1170 |
| **glean-1/decision-gcp** | docker-worker | 38.0.5 | 586 | 1172 |
| **glean-1/images-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **glean-3/b-linux-gcp** | docker-worker | 38.0.5 | 1166 | 1166 |
| **glean-3/decision-gcp** | docker-worker | 38.0.5 | 6 | 12 |
| **glean-3/images-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **infra/build-decision** | docker-worker | 38.0.5 | 13 | 104 |
| **mobile-1/b-linux-gcp** | docker-worker | 38.0.5 | 1960 | 1960 |
| **mobile-1/b-linux-large-gcp** | docker-worker | 38.0.5 | 8843 | 8843 |
| **mobile-1/b-linux-xlarge-gcp** | docker-worker | 38.0.5 | 5560 | 5560 |
| **mobile-1/bitrise-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **mobile-1/decision-gcp** | docker-worker | 38.0.5 | 483 | 966 |
| **mobile-1/images-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **mobile-3/b-linux-gcp** | docker-worker | 38.0.5 | 1837 | 1837 |
| **mobile-3/b-linux-large-gcp** | docker-worker | 38.0.5 | 9134 | 9134 |
| **mobile-3/b-linux-xlarge-gcp** | docker-worker | 38.0.5 | 5089 | 5089 |
| **mobile-3/bitrise-gcp** | docker-worker | 38.0.5 | 18 | 18 |
| **mobile-3/decision-gcp** | docker-worker | 38.0.5 | 250 | 500 |
| **mobile-3/images-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **mozilla-1/b-linux-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **mozilla-1/decision-gcp** | docker-worker | 38.0.5 | 582 | 1164 |
| **mozilla-1/images-gcp** | docker-worker | 38.0.5 | 12 | 12 |
| **mozilla-t/t-linux-large-gcp** | docker-worker | unknown version | 22 | 22 |
| **mozillaonline-1/b-linux-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **mozillaonline-3/b-linux-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **mozillavpn-1/b-linux-gcp** | docker-worker | 38.0.5 | 413 | 413 |
| **mozillavpn-1/b-linux-large-gcp** | docker-worker | 38.0.5 | 303 | 303 |
| **mozillavpn-1/decision-gcp** | docker-worker | 38.0.5 | 544 | 1088 |
| **mozillavpn-1/images-gcp** | docker-worker | 38.0.5 | 8 | 8 |
| **mozillavpn-3/b-linux-gcp** | docker-worker | 38.0.5 | 188 | 188 |
| **mozillavpn-3/b-linux-large-gcp** | docker-worker | 38.0.5 | 176 | 176 |
| **mozillavpn-3/decision-gcp** | docker-worker | 38.0.5 | 28 | 56 |
| **mozillavpn-3/images-gcp** | docker-worker | 38.0.5 | 9 | 9 |
| **nss-1/decision-gcp** | docker-worker | 38.0.5 | 2 | 4 |
| **nss-1/linux-gcp** | docker-worker | 38.0.5 | 625 | 625 |
| **nss-3/linux-gcp** | docker-worker | 38.0.5 | 86 | 86 |
| **nss-t/t-linux-xlarge-gcp** | docker-worker | unknown version | 567 | 567 |
| **project-relman/relman-svc** | docker-worker | 38.0.5 | 70 | 70 |
| **releng-1/b-linux-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **releng-1/decision-gcp** | docker-worker | 38.0.5 | 24 | 48 |
| **releng-1/linux-gcp** | docker-worker | 38.0.5 | 20 | 20 |
| **releng-3/b-linux-gcp** | docker-worker | 38.0.5 | 10 | 10 |
| **releng-3/decision-gcp** | docker-worker | 38.0.5 | 12 | 24 |
| **releng-3/linux-gcp** | docker-worker | 38.0.5 | 28 | 28 |
| **releng-t/linux-gcp** | docker-worker | 38.0.5 | 116 | 116 |
| **relops-3/decision-gcp** | docker-worker | 38.0.5 | 2 | 4 |
| **scriptworker-1/b-linux-gcp** | docker-worker | 38.0.5 | 58 | 58 |
| **scriptworker-1/decision-gcp** | docker-worker | 38.0.5 | 14 | 28 |
| **scriptworker-1/images-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **scriptworker-3/b-linux-gcp** | docker-worker | 38.0.5 | 38 | 38 |
| **scriptworker-3/decision-gcp** | docker-worker | 38.0.5 | 10 | 20 |
| **scriptworker-3/images-gcp** | docker-worker | 38.0.5 | 15 | 15 |
| **taskgraph-1/decision-gcp** | docker-worker | 38.0.5 | 7 | 14 |
| **taskgraph-1/images-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **taskgraph-3/decision-gcp** | docker-worker | 38.0.5 | 2 | 4 |
| **taskgraph-3/images-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **taskgraph-t/linux-gcp** | docker-worker | 38.0.5 | 8 | 8 |
| **translations-1/b-linux-large-gcp** | docker-worker | 38.0.5 | 213 | 213 |
| **translations-1/b-linux-large-gcp-300gb** | docker-worker | 38.0.5 | 75 | 75 |
| **translations-1/decision-gcp** | docker-worker | 38.0.5 | 36 | 72 |
| **translations-1/images-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **xpi-1/b-linux-gcp** | docker-worker | 38.0.5 | 19 | 19 |
| **xpi-1/decision-gcp** | docker-worker | 38.0.5 | 9 | 18 |
| **xpi-1/images-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **xpi-3/b-linux-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **xpi-3/decision-gcp** | docker-worker | 38.0.5 | 2 | 4 |
| **xpi-3/images-gcp** | docker-worker | 38.0.5 | 2 | 2 |


## Script Worker

Total: `52`



| Worker Pool | Implementation | Version | Total Workers | Total Capacity |
| --- | --- | --- | ---: | ---: |
| **scriptworker-k8s/adhoc-1-shipit-dev** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/app-services-1-shipit** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/app-services-1-shipit-dev** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/app-services-3-beetmover** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/app-services-3-shipit** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/app-services-3-signing** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/comm-3-balrog** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/comm-3-beetmover** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/comm-3-signing** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/comm-3-tree** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/comm-t-signing** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/gecko-1-balrog** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/gecko-1-beetmover** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/gecko-1-bouncer** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/gecko-1-tree** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/gecko-3-balrog** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/gecko-3-beetmover** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/gecko-3-bouncer** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/gecko-3-signing** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/gecko-3-tree** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/gecko-t-signing** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/mobile-1-tree** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/mobile-1-tree-dev** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/mobile-3-beetmover** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/mobile-3-github** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/mobile-3-pushapk** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/mobile-3-shipit** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/mobile-3-signing** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/mobile-3-tree** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/mobile-t-signing** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/mozillavpn-1-beetmover** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/mozillavpn-1-beetmover-dev** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/mozillavpn-1-pushapk-dev** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/mozillavpn-3-beetmover** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/mozillavpn-3-signing** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/xpi-1-balrog** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/xpi-1-balrog-dev** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/xpi-1-beetmover** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/xpi-1-beetmover-dev** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/xpi-1-github** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/xpi-1-github-dev** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/xpi-3-balrog** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/xpi-3-beetmover** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/xpi-3-github** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
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
| projects/taskcluster-imaging/global/images/generic-2204-wayland-vm-gcp-aerickson-test-2023-09-20t01-30-09z | 1 |
|  | 4 |
| projects/taskcluster-imaging/global/images/generic-2204-wayland-vm-gcp-googlecompute-2023-09-22t17-39-37z | 1 |


| Worker Pool | Implementation | Version | Total Workers | Total Capacity |
| --- | --- | --- | ---: | ---: |
| **built-in/fail** |  | No artifacts found | 0 | 0 |
| **built-in/succeed** |  | No artifacts found | 0 | 0 |
| **gecko-t/t-linux-vm-2204-wayland** |  | No artifacts found | 1743 | 1743 |
| **gecko-t/t-linux-vm-2204-wayland-aerickson-test** |  | No artifacts found | 3 | 3 |
| **scriptworker-prov-v1/mac-notarization-poller** |  | No artifacts found | 0 | 0 |
| **scriptworker-prov-v1/tb-mac-notarization-poller** |  | No artifacts found | 0 | 0 |


## Version not determined [^2]

Total: `27`


Count by image:

| Version | Count |
| :--- | ---: |
| /subscriptions/108d46d5-fe9b-4850-9a7d-8c914aa6c1f0/resourceGroups/rg-packer-through-cib/providers/Microsoft.Compute/images/win2012r2-64-l1-centralus-2012-R2-Datacenter-21c6826 | 4 |
| /subscriptions/108d46d5-fe9b-4850-9a7d-8c914aa6c1f0/resourceGroups/rg-packer-through-cib/providers/Microsoft.Compute/images/win2022-64-2009-centralus-2022-datacenter-azure-edition-alpha | 2 |
| /subscriptions/108d46d5-fe9b-4850-9a7d-8c914aa6c1f0/resourceGroups/rg-packer-through-cib/providers/Microsoft.Compute/images/win11-64-2009-tc-centralus-win11-22h2-avd-alpha | 1 |
| /subscriptions/108d46d5-fe9b-4850-9a7d-8c914aa6c1f0/resourceGroups/rg-packer-through-cib/providers/Microsoft.Compute/images/win10-64-2009-centralus-win10-22h2-avd-alpha | 1 |
| /subscriptions/108d46d5-fe9b-4850-9a7d-8c914aa6c1f0/resourceGroups/rg-packer-through-cib/providers/Microsoft.Compute/images/win11-64-2009-eastus-win11-22h2-avd-alpha | 1 |
| /subscriptions/a30e97ab-734a-4f3b-a0e4-c51c0bff0701/resourceGroups/rg-packer-through-cib/providers/Microsoft.Compute/images/trusted-win2022-64-2009-centralus-2022-datacenter-azure-edition-d5dee60 | 1 |
| /subscriptions/108d46d5-fe9b-4850-9a7d-8c914aa6c1f0/resourceGroups/rg-packer-through-cib/providers/Microsoft.Compute/images/win2012r2-64-vs-l1-centralus-2012-R2-Datacenter-alpha | 1 |
|  | 2 |
| /subscriptions/108d46d5-fe9b-4850-9a7d-8c914aa6c1f0/resourceGroups/rg-packer-through-cib/providers/Microsoft.Compute/images/win2012r2-64-vs-py2-l1-centralus-2012-R2-Datacenter-7aa76c6 | 1 |
| /subscriptions/108d46d5-fe9b-4850-9a7d-8c914aa6c1f0/resourceGroups/rg-packer-through-cib/providers/Microsoft.Compute/images/win11-64-2009-centralus-win11-22h2-avd-d322b59 | 1 |
| /subscriptions/108d46d5-fe9b-4850-9a7d-8c914aa6c1f0/resourceGroups/rg-packer-through-cib/providers/Microsoft.Compute/images/win10-64-2009-ukwest-win10-22h2-avd-alpha | 1 |
| /subscriptions/108d46d5-fe9b-4850-9a7d-8c914aa6c1f0/resourceGroups/rg-packer-through-cib/providers/Microsoft.Compute/images/win2022-64-2009-rd-centralus-2022-datacenter-azure-edition-alpha | 1 |
| /subscriptions/108d46d5-fe9b-4850-9a7d-8c914aa6c1f0/resourceGroups/rg-packer-through-cib/providers/Microsoft.Compute/images/win2012r2-64-l1-centralus-2012-R2-Datacenter-alpha | 2 |
| /subscriptions/108d46d5-fe9b-4850-9a7d-8c914aa6c1f0/resourceGroups/rg-packer-through-cib/providers/Microsoft.Compute/images/win2022-64-2009-centralus-2022-datacenter-azure-edition-531c001 | 2 |
| /subscriptions/108d46d5-fe9b-4850-9a7d-8c914aa6c1f0/resourceGroups/rg-packer-through-cib/providers/Microsoft.Compute/images/masterwayz-win2012test-image-20221122010013-alpha | 1 |
| /subscriptions/108d46d5-fe9b-4850-9a7d-8c914aa6c1f0/resourceGroups/rg-packer-through-cib/providers/Microsoft.Compute/images/win11-64-2009-centralus-win11-22h2-avd-69a7ecf | 1 |
| /subscriptions/108d46d5-fe9b-4850-9a7d-8c914aa6c1f0/resourceGroups/rg-packer-through-cib/providers/Microsoft.Compute/images/win2012r2-64-vs-l1-centralus-2012-R2-Datacenter-21c6826 | 1 |
| /subscriptions/108d46d5-fe9b-4850-9a7d-8c914aa6c1f0/resourceGroups/rg-packer-through-cib/providers/Microsoft.Compute/images/win2012r2-64-vs-py2-l1-centralus-2012-R2-Datacenter-alpha | 1 |
| /subscriptions/108d46d5-fe9b-4850-9a7d-8c914aa6c1f0/resourceGroups/rg-packer-through-cib/providers/Microsoft.Compute/images/win10-64-2009-eastus-win10-22h2-avd-alpha | 1 |
| /subscriptions/a30e97ab-734a-4f3b-a0e4-c51c0bff0701/resourceGroups/rg-packer-through-cib/providers/Microsoft.Compute/images/trusted-win2012r2-64-l3-centralus-2012-R2-Datacenter-7cde253 | 1 |


| Worker Pool | Implementation | Version | Total Workers | Total Capacity |
| --- | --- | --- | ---: | ---: |
| **comm-1/b-win2012-azure** |  | Version not determined; task not (yet) claimed | 4 | 4 |
| **comm-2/b-win2012-azure** |  | Version not determined; task not (yet) claimed | 4 | 4 |
| **gecko-1/b-win2012-azure** |  | Version not determined; task not (yet) claimed | 2 | 2 |
| **gecko-1/b-win2012-azure-alpha** |  | Version not determined; task not (yet) claimed | 4 | 4 |
| **gecko-1/b-win2012-azure-beta** |  | Version not determined; task not (yet) claimed | 4 | 4 |
| **gecko-1/b-win2022** |  | Version not determined; task not (yet) claimed | 750 | 750 |
| **gecko-1/b-win2022-alpha** |  | Version not determined; task not (yet) claimed | 4 | 4 |
| **gecko-1/b-win2022-rd-alpha** |  | Version not determined; task not (yet) claimed | 4 | 4 |
| **gecko-2/b-win2012-azure** |  | Version not determined; task not (yet) claimed | 4 | 4 |
| **gecko-3/b-win2012-azure** |  | Version not determined; task not (yet) claimed | 21 | 21 |
| **gecko-t/win10-64-2009-alpha** |  | Version not determined; task not (yet) claimed | 4 | 4 |
| **gecko-t/win10-64-2009-gpu-alpha** |  | Version not determined; task not (yet) claimed | 4 | 4 |
| **gecko-t/win10-64-2009-source-alpha** |  | Version not determined; task not (yet) claimed | 6 | 6 |
| **gecko-t/win11-64-2009-alpha-test** |  | Version not determined; task not (yet) claimed | 4 | 4 |
| **gecko-t/win11-64-2009-gpu-alpha** |  | Version not determined; task not (yet) claimed | 31 | 31 |
| **gecko-t/win11-64-2009-source** |  | Version not determined; task not (yet) claimed | 927 | 927 |
| **gecko-t/win11-64-2009-tc-alpha** |  | Version not determined; task not (yet) claimed | 4 | 4 |
| **mozillavpn-1/b-win2012-azure** |  | Version not determined; task not (yet) claimed | 31 | 31 |
| **mozillavpn-1/b-win2012-azure-alpha** |  | Version not determined; task not (yet) claimed | 4 | 4 |
| **mozillavpn-1/masterwayz-win2012-azure** |  | Version not determined; task not (yet) claimed | 6 | 6 |
| **nss-1/b-win2012-azure** |  | Version not determined; task not (yet) claimed | 859 | 859 |
| **nss-1/b-win2012-azure-alpha** |  | Version not determined; task not (yet) claimed | 4 | 4 |
| **nss-1/b-win2022-alpha** |  | Version not determined; task not (yet) claimed | 4 | 4 |
| **releng-hardware/gecko-t-osx-1015-r8** |  | Version not determined; task not (yet) claimed | 0 | 0 |
| **releng-hardware/gecko-t-win10-64-1803-hw** |  | Version not determined; task not (yet) claimed | 0 | 0 |
| **relops-1/b-win2022** |  | Version not determined; task not (yet) claimed | 4 | 4 |
| **relops-3/b-win2022** |  | Version not determined; task not (yet) claimed | 4 | 4 |



[^1]: Those are the pools whose tasks were claimed and resolved by a worker as expected, but the worker did not publish either artifact `public/logs/live_backing.log` nor `public/logs/chain_of_trust.log`, which is the source used to identify the worker implementation.

[^2]: Probing task remains pending after two hours. Those are the pools that were not able to start any worker to claim the task within two hours.

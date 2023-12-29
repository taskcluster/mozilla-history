

# Worker Pool Versions


## Generic Worker

Total: `57`

Count by version:

| Version | Count |
| :--- | ---: |
| 14.1.2 | 1 |
| 15.1.4 | 1 |
| 16.2.0 | 3 |
| 16.5.1 | 1 |
| 30.0.2 | 3 |
| 36.0.0 | 2 |
| 40.0.3 | 2 |
| 43.2.0 | 4 |
| 44.0.0 | 1 |
| 44.23.2 | 1 |
| 45.0.0 | 1 |
| 54.4.1 | 24 |
| 55.1.1 | 4 |
| 55.3.4 | 7 |
| 57.0.1 | 1 |
| 59.1.3 | 1 |


Count by image:

| Version | Count |
| :--- | ---: |
| projects/taskcluster-imaging/global/images/generic-translations-gcp-googlecompute-2023-09-12t21-35-55z | 4 |
| /subscriptions/108d46d5-fe9b-4850-9a7d-8c914aa6c1f0/resourceGroups/rg-packer-through-cib/providers/Microsoft.Compute/images/win2012r2-64-vs-py2-l1-centralus-2012-R2-Datacenter-7aa76c6 | 1 |
| /subscriptions/a30e97ab-734a-4f3b-a0e4-c51c0bff0701/resourceGroups/rg-packer-through-cib/providers/Microsoft.Compute/images/trusted-win2012r2-64-vs-py2-l3-centralus-2012-R2-Datacenter-7aa76c6 | 1 |
| /subscriptions/108d46d5-fe9b-4850-9a7d-8c914aa6c1f0/resourceGroups/rg-packer-through-cib/providers/Microsoft.Compute/images/win10-64-2009-eastus-win10-22h2-avd-531c001 | 1 |
| projects/taskcluster-imaging/global/images/generic-translations-gcp-googlecompute-2023-11-01t20-12-18z | 1 |
| /subscriptions/a30e97ab-734a-4f3b-a0e4-c51c0bff0701/resourceGroups/rg-packer-through-cib/providers/Microsoft.Compute/images/trusted-win2012r2-64-vs-l3-centralus-2012-R2-Datacenter-7cde253 | 1 |
|  | 27 |
| ami-07db1c26bda143e33 | 1 |
| unknown | 1 |
| /subscriptions/108d46d5-fe9b-4850-9a7d-8c914aa6c1f0/resourceGroups/rg-packer-through-cib/providers/Microsoft.Compute/images/win11-64-2009-eastus-win11-22h2-avd-alpha | 1 |
| /subscriptions/108d46d5-fe9b-4850-9a7d-8c914aa6c1f0/resourceGroups/rg-packer-through-cib/providers/Microsoft.Compute/images/win11-64-2009-centralus-win11-22h2-avd-531c001 | 3 |
| /subscriptions/a30e97ab-734a-4f3b-a0e4-c51c0bff0701/resourceGroups/rg-packer-through-cib/providers/Microsoft.Compute/images/trusted-win2012r2-64-l3-centralus-2012-R2-Datacenter-7cde253 | 1 |
| /subscriptions/108d46d5-fe9b-4850-9a7d-8c914aa6c1f0/resourceGroups/rg-packer-through-cib/providers/Microsoft.Compute/images/win10-64-2009-canadacentral-win10-22h2-avd-531c001 | 2 |
| /subscriptions/108d46d5-fe9b-4850-9a7d-8c914aa6c1f0/resourceGroups/rg-packer-through-cib/providers/Microsoft.Compute/images/win11-64-2009-eastus-win11-22h2-avd-531c001 | 2 |
| /subscriptions/108d46d5-fe9b-4850-9a7d-8c914aa6c1f0/resourceGroups/rg-packer-through-cib/providers/Microsoft.Compute/images/win11-64-2009-uksouth-win11-22h2-avd-alpha | 2 |
| /subscriptions/a30e97ab-734a-4f3b-a0e4-c51c0bff0701/resourceGroups/rg-packer-through-cib/providers/Microsoft.Compute/images/trusted-win2022-64-2009-centralus-2022-datacenter-azure-edition-28c1765 | 3 |
| projects/taskcluster-imaging/global/images/generic-translations-gcp-googlecompute-2023-12-19t18-58-22z | 1 |
| /subscriptions/108d46d5-fe9b-4850-9a7d-8c914aa6c1f0/resourceGroups/rg-packer-through-cib/providers/Microsoft.Compute/images/win2022-64-2009-centralus-2022-datacenter-azure-edition-531c001 | 4 |


| Worker Pool | Implementation | Version | Engine | Revision | OS | Arch | GO | Total Workers | Total Capacity |
| --- | --- | --- | --- | --- | --- | --- | --- | ---: | ---: |
| **bitbar/gecko-t-win64-aarch64-laptop** | generic-worker | 14.1.2 | <UNKNOWN> | 13118c4c1b | windows | 386 | 1.10.8 | 0 | 0 |
| **comm-1/b-win2022** | generic-worker | 54.4.1 | multiuser | a38a78d98c | windows | amd64 | 1.20.6 | 4 | 4 |
| **comm-2/b-win2022** | generic-worker | 54.4.1 | multiuser | a38a78d98c | windows | amd64 | 1.20.6 | 4 | 4 |
| **comm-3/b-win2012-azure** | generic-worker | 43.2.0 | multiuser | 53fa7b79bc | windows | amd64 | 1.16.3 | 4 | 4 |
| **comm-3/b-win2022** | generic-worker | 54.4.1 | multiuser | a38a78d98c | windows | amd64 | 1.20.6 | 438 | 438 |
| **comm-t/win11-64-2009-ssd** | generic-worker | 54.4.1 | multiuser | a38a78d98c | windows | amd64 | 1.20.6 | 391 | 391 |
| **gecko-1/b-win2022** | generic-worker | 54.4.1 | multiuser | a38a78d98c | windows | amd64 | 1.20.6 | 533 | 533 |
| **gecko-3/b-win2022** | generic-worker | 54.4.1 | multiuser | a38a78d98c | windows | amd64 | 1.20.6 | 5410 | 5410 |
| **gecko-t/win10-64-2009** | generic-worker | 54.4.1 | multiuser | a38a78d98c | windows | amd64 | 1.20.6 | 4748 | 4748 |
| **gecko-t/win10-64-2009-gpu** | generic-worker | 54.4.1 | multiuser | a38a78d98c | windows | amd64 | 1.20.6 | 618 | 618 |
| **gecko-t/win10-64-2009-source** | generic-worker | 54.4.1 | multiuser | a38a78d98c | windows | amd64 | 1.20.6 | 4 | 4 |
| **gecko-t/win11-64-2009-alpha** | generic-worker | 54.4.1 | multiuser | a38a78d98c | windows | amd64 | 1.20.6 | 4 | 4 |
| **gecko-t/win11-64-2009-gpu** | generic-worker | 54.4.1 | multiuser | a38a78d98c | windows | amd64 | 1.20.6 | 14736 | 14736 |
| **gecko-t/win11-64-2009-gpu-alpha** | generic-worker | 54.4.1 | multiuser | a38a78d98c | windows | amd64 | 1.20.6 | 4 | 4 |
| **gecko-t/win11-64-2009-source** | generic-worker | 54.4.1 | multiuser | a38a78d98c | windows | amd64 | 1.20.6 | 850 | 850 |
| **gecko-t/win11-64-2009-source-alpha** | generic-worker | 54.4.1 | multiuser | a38a78d98c | windows | amd64 | 1.20.6 | 4 | 4 |
| **gecko-t/win11-64-2009-ssd** | generic-worker | 54.4.1 | multiuser | a38a78d98c | windows | amd64 | 1.20.6 | 1862 | 1862 |
| **gecko-t/win11-64-2009-ssd-gpu** | generic-worker | 54.4.1 | multiuser | a38a78d98c | windows | amd64 | 1.20.6 | 1565 | 1565 |
| **localprovisioner/nss-macos-10-12** | generic-worker | 15.1.4 | simple | c407e45e3f | darwin | amd64 | 1.10.8 | 0 | 0 |
| **localprovisioner/nss-macos-m1** | generic-worker | 30.0.2 | simple | 6fdba0dad3 | darwin | arm64 | 1.16.4 | 0 | 0 |
| **mozillavpn-1/b-win2022** | generic-worker | 54.4.1 | multiuser | a38a78d98c | windows | amd64 | 1.20.6 | 124 | 124 |
| **mozillavpn-3/b-win2012-azure** | generic-worker | 43.2.0 | multiuser | 53fa7b79bc | windows | amd64 | 1.16.3 | 4 | 4 |
| **mozillavpn-3/b-win2022** | generic-worker | 54.4.1 | multiuser | a38a78d98c | windows | amd64 | 1.20.6 | 32 | 32 |
| **nss-1/b-win2012-azure** | generic-worker | 43.2.0 | multiuser | 53fa7b79bc | windows | amd64 | 1.16.3 | 78 | 78 |
| **nss-3/b-win2012-azure** | generic-worker | 43.2.0 | multiuser | 53fa7b79bc | windows | amd64 | 1.16.3 | 29 | 29 |
| **performance-hardware/gecko-t-fxrecorder** | generic-worker | 44.0.0 | multiuser | faf6f319c0 | windows | amd64 | 1.16.3 | 0 | 0 |
| **proj-autophone/gecko-t-bitbar-gw-perf-a51** | generic-worker | 36.0.0 | simple | b9cb11293f | linux | amd64 | 1.13.7 | 0 | 0 |
| **proj-autophone/gecko-t-bitbar-gw-unit-p5** | generic-worker | 36.0.0 | simple | b9cb11293f | linux | amd64 | 1.13.7 | 0 | 0 |
| **releng-hardware/applicationservices-b-1-osx1015** | generic-worker | 55.3.4 | multiuser | a71f300fcb | darwin | amd64 | 1.21.1 | 0 | 0 |
| **releng-hardware/applicationservices-b-3-osx1015** | generic-worker | 54.4.1 | multiuser | a38a78d98c | darwin | amd64 | 1.20.6 | 0 | 0 |
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
| **releng-hardware/gecko-t-osx-1300-m2** | generic-worker | 55.3.4 | simple | a71f300fcb | darwin | arm64 | 1.21.1 | 0 | 0 |
| **releng-hardware/gecko-t-win-talos** | generic-worker | 16.2.0 | multiuser | b321a877c8 | windows | amd64 | 1.10.8 | 0 | 0 |
| **releng-hardware/gecko-t-win10-64-1803-hw** | generic-worker | 16.2.0 | multiuser | b321a877c8 | windows | amd64 | 1.10.8 | 0 | 0 |
| **releng-hardware/gecko-t-win10-64-dev** | generic-worker | 16.2.0 | multiuser | b321a877c8 | windows | amd64 | 1.10.8 | 0 | 0 |
| **releng-hardware/gecko-t-win7-32-hw** | generic-worker | 45.0.0 | multiuser | 988e8100b3 | windows | 386 | 1.19.3 | 0 | 0 |
| **releng-hardware/mozillavpn-b-1-osx** | generic-worker | 54.4.1 | multiuser | a38a78d98c | darwin | amd64 | 1.20.6 | 0 | 0 |
| **releng-hardware/mozillavpn-b-3-osx** | generic-worker | 54.4.1 | multiuser | a38a78d98c | darwin | amd64 | 1.20.6 | 0 | 0 |
| **releng-hardware/nss-1-b-osx-1015** | generic-worker | 55.3.4 | multiuser | a71f300fcb | darwin | amd64 | 1.21.1 | 0 | 0 |
| **releng-hardware/win11-64-2009-hw-ref** | generic-worker | 54.4.1 | multiuser | a38a78d98c | windows | amd64 | 1.20.6 | 0 | 0 |
| **releng-hardware/win11-64-2009-hw-ref-alpha** | generic-worker | 54.4.1 | multiuser | a38a78d98c | windows | amd64 | 1.20.6 | 0 | 0 |
| **relops-3/win2019** | generic-worker | 16.5.1 | multiuser | 02bc4c9cd7 | windows | amd64 | 1.10.8 | 4 | 4 |
| **translations-1/b-linux-aerickson-test** | generic-worker | 57.0.1 | simple | fe45895655 | linux | amd64 | 1.21.3 | 2 | 2 |
| **translations-1/b-linux-v100-gpu** | generic-worker | 55.1.1 | simple | b83fffd4bd | linux | amd64 | 1.21.1 | 188 | 188 |
| **translations-1/b-linux-v100-gpu-4** | generic-worker | 55.1.1 | simple | b83fffd4bd | linux | amd64 | 1.21.1 | 140 | 140 |
| **translations-1/b-linux-v100-gpu-4-1tb** | generic-worker | 55.1.1 | simple | b83fffd4bd | linux | amd64 | 1.21.1 | 437 | 437 |
| **translations-1/b-linux-v100-gpu-4-300gb** | generic-worker | 55.1.1 | simple | b83fffd4bd | linux | amd64 | 1.21.1 | 3 | 3 |
| **translations-1/b-linux-v100-gpu-relops782-test** | generic-worker | 59.1.3 | simple | b550fe4ca7 | linux | amd64 | 1.21.5 | 12 | 12 |


## Docker Worker

Total: `132`

Count by version:

| Version | Count |
| :--- | ---: |
| 38.0.5 | 122 |
| 48.3.0 | 1 |
| unknown version | 9 |


Count by image:

| Version | Count |
| :--- | ---: |
| projects/taskcluster-imaging/global/images/docker-firefoxci-gcp-lt-googlecompute-2023-04-13t21-30-28z | 1 |
| projects/fxci-production-level3-workers/global/images/docker-firefoxci-gcp-l3-googlecompute-2023-12-06t15-04-40z | 47 |
| projects/taskcluster-imaging/global/images/docker-firefoxci-gcp-l1-googlecompute-2023-06-13t23-47-26z | 74 |
| projects/taskcluster-imaging/global/images/docker-firefoxci-gcp-l1-googlecompute-2023-12-04t22-11-57z | 1 |
| projects/taskcluster-imaging/global/images/docker-worker-gcp-u14-04-2023-06-13 | 9 |


| Worker Pool | Implementation | Version | Total Workers | Total Capacity |
| --- | --- | --- | ---: | ---: |
| **adhoc-1/b-linux-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **adhoc-1/decision-gcp** | docker-worker | 38.0.5 | 2 | 4 |
| **adhoc-1/images-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **adhoc-3/b-linux-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **adhoc-3/decision-gcp** | docker-worker | 38.0.5 | 2 | 4 |
| **adhoc-3/images-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **app-services-1/b-linux-gcp** | docker-worker | 38.0.5 | 93 | 93 |
| **app-services-1/decision-gcp** | docker-worker | 38.0.5 | 619 | 1238 |
| **app-services-1/images-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **app-services-3/b-linux-gcp** | docker-worker | 38.0.5 | 473 | 473 |
| **app-services-3/decision-gcp** | docker-worker | 38.0.5 | 25 | 50 |
| **app-services-3/images-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **ci-1/decision-gcp** | docker-worker | 38.0.5 | 6 | 12 |
| **ci-1/images-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **ci-3/decision-gcp** | docker-worker | 38.0.5 | 5 | 10 |
| **ci-3/images-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **ci-t/linux-gcp** | docker-worker | 38.0.5 | 35 | 35 |
| **code-analysis-1/linux-gcp** | docker-worker | 38.0.5 | 19 | 19 |
| **code-analysis-3/linux-gcp** | docker-worker | 38.0.5 | 17 | 17 |
| **code-coverage/bot-gcp** | docker-worker | 38.0.5 | 818 | 818 |
| **code-review/bot-gcp** | docker-worker | 38.0.5 | 1854 | 1854 |
| **comm-1/b-linux-gcp** | docker-worker | 38.0.5 | 32 | 32 |
| **comm-1/b-linux-large-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **comm-1/b-linux-xlarge-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **comm-1/decision-gcp** | docker-worker | 38.0.5 | 27 | 54 |
| **comm-1/images-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **comm-2/b-linux-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **comm-2/b-linux-large-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **comm-2/b-linux-xlarge-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **comm-2/decision-gcp** | docker-worker | 38.0.5 | 2 | 4 |
| **comm-2/images-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **comm-3/b-linux-gcp** | docker-worker | 38.0.5 | 2006 | 2006 |
| **comm-3/b-linux-large-gcp** | docker-worker | 38.0.5 | 5 | 5 |
| **comm-3/b-linux-xlarge-gcp** | docker-worker | 38.0.5 | 5 | 5 |
| **comm-3/decision-gcp** | docker-worker | 38.0.5 | 50 | 100 |
| **comm-3/images-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **comm-t/misc-gcp** | docker-worker | 38.0.5 | 2 | 16 |
| **comm-t/t-linux-large-gcp** | docker-worker | unknown version | 620 | 620 |
| **comm-t/t-linux-xlarge-gcp** | docker-worker | unknown version | 128 | 128 |
| **comm-t/t-linux-xlarge-source-gcp** | docker-worker | unknown version | 262 | 262 |
| **gecko-1/b-linux-gcp** | docker-worker | 38.0.5 | 4478 | 4478 |
| **gecko-1/b-linux-gcp-bug1797804-n2** | docker-worker | 38.0.5 | 2 | 2 |
| **gecko-1/b-linux-gcp-relops757-test** | docker-worker | 38.0.5 | 3 | 3 |
| **gecko-1/b-linux-large-gcp** | docker-worker | 38.0.5 | 1453 | 1453 |
| **gecko-1/b-linux-xlarge-gcp** | docker-worker | 38.0.5 | 292 | 292 |
| **gecko-1/b-linux-xlarge-gcp-bug1797804-c2** | docker-worker | 38.0.5 | 4 | 4 |
| **gecko-1/decision-gcp** | docker-worker | 38.0.5 | 1633 | 3266 |
| **gecko-1/images-gcp** | docker-worker | 38.0.5 | 9 | 9 |
| **gecko-2/b-linux-gcp** | docker-worker | 38.0.5 | 679 | 679 |
| **gecko-2/b-linux-large-gcp** | docker-worker | 38.0.5 | 22 | 22 |
| **gecko-2/b-linux-xlarge-gcp** | docker-worker | 38.0.5 | 15 | 15 |
| **gecko-2/decision-gcp** | docker-worker | 38.0.5 | 15 | 30 |
| **gecko-2/images-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **gecko-3/b-linux-gcp** | docker-worker | 38.0.5 | 15624 | 15624 |
| **gecko-3/b-linux-large-gcp** | docker-worker | 38.0.5 | 117 | 117 |
| **gecko-3/b-linux-xlarge-gcp** | docker-worker | 38.0.5 | 890 | 890 |
| **gecko-3/decision-gcp** | docker-worker | 38.0.5 | 1591 | 3182 |
| **gecko-3/images-gcp** | docker-worker | 38.0.5 | 7 | 7 |
| **gecko-3/t-linux-xlarge-gcp** | docker-worker | 38.0.5 | 241 | 241 |
| **gecko-t/misc-gcp** | docker-worker | 38.0.5 | 304 | 2432 |
| **gecko-t/t-linux-kvm-gcp** | docker-worker | unknown version | 23897 | 23897 |
| **gecko-t/t-linux-large-gcp** | docker-worker | unknown version | 63020 | 63020 |
| **gecko-t/t-linux-large-hostub2204-gcp** | docker-worker | 48.3.0 | 2 | 2 |
| **gecko-t/t-linux-xlarge-gcp** | docker-worker | unknown version | 31969 | 31969 |
| **gecko-t/t-linux-xlarge-source-gcp** | docker-worker | unknown version | 8534 | 8534 |
| **glean-1/b-linux-gcp** | docker-worker | 38.0.5 | 1236 | 1236 |
| **glean-1/decision-gcp** | docker-worker | 38.0.5 | 622 | 1244 |
| **glean-1/images-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **glean-3/b-linux-gcp** | docker-worker | 38.0.5 | 1245 | 1245 |
| **glean-3/decision-gcp** | docker-worker | 38.0.5 | 2 | 4 |
| **glean-3/images-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **infra/build-decision** | docker-worker | 38.0.5 | 12 | 96 |
| **mobile-1/b-linux-gcp** | docker-worker | 38.0.5 | 1018 | 1018 |
| **mobile-1/b-linux-large-gcp** | docker-worker | 38.0.5 | 3918 | 3918 |
| **mobile-1/b-linux-xlarge-gcp** | docker-worker | 38.0.5 | 2688 | 2688 |
| **mobile-1/bitrise-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **mobile-1/decision-gcp** | docker-worker | 38.0.5 | 564 | 1128 |
| **mobile-1/images-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **mobile-3/b-linux-gcp** | docker-worker | 38.0.5 | 1150 | 1150 |
| **mobile-3/b-linux-large-gcp** | docker-worker | 38.0.5 | 6786 | 6786 |
| **mobile-3/b-linux-xlarge-gcp** | docker-worker | 38.0.5 | 2536 | 2536 |
| **mobile-3/bitrise-gcp** | docker-worker | 38.0.5 | 20 | 20 |
| **mobile-3/decision-gcp** | docker-worker | 38.0.5 | 190 | 380 |
| **mobile-3/images-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **mozilla-1/b-linux-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **mozilla-1/decision-gcp** | docker-worker | 38.0.5 | 622 | 1244 |
| **mozilla-1/images-gcp** | docker-worker | 38.0.5 | 11 | 11 |
| **mozilla-t/t-linux-large-gcp** | docker-worker | unknown version | 23 | 23 |
| **mozillaonline-1/b-linux-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **mozillaonline-3/b-linux-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **mozillavpn-1/b-linux-gcp** | docker-worker | 38.0.5 | 252 | 252 |
| **mozillavpn-1/b-linux-large-gcp** | docker-worker | 38.0.5 | 163 | 163 |
| **mozillavpn-1/decision-gcp** | docker-worker | 38.0.5 | 609 | 1218 |
| **mozillavpn-1/images-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **mozillavpn-3/b-linux-gcp** | docker-worker | 38.0.5 | 65 | 65 |
| **mozillavpn-3/b-linux-large-gcp** | docker-worker | 38.0.5 | 58 | 58 |
| **mozillavpn-3/decision-gcp** | docker-worker | 38.0.5 | 8 | 16 |
| **mozillavpn-3/images-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **nss-1/decision-gcp** | docker-worker | 38.0.5 | 2 | 4 |
| **nss-1/linux-gcp** | docker-worker | 38.0.5 | 320 | 320 |
| **nss-3/linux-gcp** | docker-worker | 38.0.5 | 88 | 88 |
| **nss-t/t-linux-xlarge-gcp** | docker-worker | unknown version | 377 | 377 |
| **project-relman/relman-svc** | docker-worker | 38.0.5 | 68 | 68 |
| **releng-1/b-linux-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **releng-1/decision-gcp** | docker-worker | 38.0.5 | 13 | 26 |
| **releng-1/linux-gcp** | docker-worker | 38.0.5 | 15 | 15 |
| **releng-3/b-linux-gcp** | docker-worker | 38.0.5 | 8 | 8 |
| **releng-3/decision-gcp** | docker-worker | 38.0.5 | 7 | 14 |
| **releng-3/linux-gcp** | docker-worker | 38.0.5 | 12 | 12 |
| **releng-t/linux-gcp** | docker-worker | 38.0.5 | 100 | 100 |
| **relops-3/decision-gcp** | docker-worker | 38.0.5 | 2 | 4 |
| **scriptworker-1/b-linux-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **scriptworker-1/decision-gcp** | docker-worker | 38.0.5 | 2 | 4 |
| **scriptworker-1/images-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **scriptworker-3/b-linux-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **scriptworker-3/decision-gcp** | docker-worker | 38.0.5 | 2 | 4 |
| **scriptworker-3/images-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **taskgraph-1/decision-gcp** | docker-worker | 38.0.5 | 7 | 14 |
| **taskgraph-1/images-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **taskgraph-3/decision-gcp** | docker-worker | 38.0.5 | 5 | 10 |
| **taskgraph-3/images-gcp** | docker-worker | 38.0.5 | 4 | 4 |
| **taskgraph-t/linux-gcp** | docker-worker | 38.0.5 | 21 | 21 |
| **translations-1/b-linux-large-gcp** | docker-worker | 38.0.5 | 69 | 69 |
| **translations-1/b-linux-large-gcp-300gb** | docker-worker | 38.0.5 | 24 | 24 |
| **translations-1/decision-gcp** | docker-worker | 38.0.5 | 18 | 36 |
| **translations-1/images-gcp** | docker-worker | 38.0.5 | 3 | 3 |
| **xpi-1/b-linux-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **xpi-1/decision-gcp** | docker-worker | 38.0.5 | 2 | 4 |
| **xpi-1/images-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **xpi-3/b-linux-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **xpi-3/decision-gcp** | docker-worker | 38.0.5 | 2 | 4 |
| **xpi-3/images-gcp** | docker-worker | 38.0.5 | 2 | 2 |


## Script Worker

Total: `25`



| Worker Pool | Implementation | Version | Total Workers | Total Capacity |
| --- | --- | --- | ---: | ---: |
| **scriptworker-k8s/app-services-3-beetmover** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/app-services-3-signing** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/comm-3-balrog** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/comm-3-beetmover** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
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
| projects/taskcluster-imaging/global/images/generic-2204-wayland-vm-gcp-googlecompute-2023-09-22t17-39-37z | 1 |
|  | 4 |


| Worker Pool | Implementation | Version | Total Workers | Total Capacity |
| --- | --- | --- | ---: | ---: |
| **built-in/fail** |  | No artifacts found | 0 | 0 |
| **built-in/succeed** |  | No artifacts found | 0 | 0 |
| **gecko-t/t-linux-vm-2204-wayland** |  | No artifacts found | 2526 | 2526 |
| **gecko-t/t-linux-vm-2204-wayland-aerickson-test** |  | No artifacts found | 4 | 4 |
| **scriptworker-prov-v1/mac-notarization-poller** |  | No artifacts found | 0 | 0 |
| **scriptworker-prov-v1/tb-mac-notarization-poller** |  | No artifacts found | 0 | 0 |


## Version not determined [^2]

Total: `26`


Count by image:

| Version | Count |
| :--- | ---: |
| /subscriptions/a30e97ab-734a-4f3b-a0e4-c51c0bff0701/resourceGroups/rg-packer-through-cib/providers/Microsoft.Compute/images/trusted-win2022-64-2009-centralus-2022-datacenter-azure-edition-28c1765 | 1 |
| /subscriptions/108d46d5-fe9b-4850-9a7d-8c914aa6c1f0/resourceGroups/rg-packer-through-cib/providers/Microsoft.Compute/images/win2012r2-64-vs-py2-l1-centralus-2012-R2-Datacenter-alpha | 1 |
| /subscriptions/108d46d5-fe9b-4850-9a7d-8c914aa6c1f0/resourceGroups/rg-packer-through-cib/providers/Microsoft.Compute/images/win10-64-2009-eastus-win10-22h2-avd-alpha | 1 |
| /subscriptions/108d46d5-fe9b-4850-9a7d-8c914aa6c1f0/resourceGroups/rg-packer-through-cib/providers/Microsoft.Compute/images/win2012r2-64-l1-centralus-2012-R2-Datacenter-21c6826 | 4 |
| /subscriptions/108d46d5-fe9b-4850-9a7d-8c914aa6c1f0/resourceGroups/rg-packer-through-cib/providers/Microsoft.Compute/images/win2012r2-64-l1-centralus-2012-R2-Datacenter-alpha | 2 |
| /subscriptions/108d46d5-fe9b-4850-9a7d-8c914aa6c1f0/resourceGroups/rg-packer-through-cib/providers/Microsoft.Compute/images/masterwayz-win2012test-image-20221122010013-alpha | 1 |
| /subscriptions/108d46d5-fe9b-4850-9a7d-8c914aa6c1f0/resourceGroups/rg-packer-through-cib/providers/Microsoft.Compute/images/win2012r2-64-vs-l1-centralus-2012-R2-Datacenter-21c6826 | 1 |
| /subscriptions/108d46d5-fe9b-4850-9a7d-8c914aa6c1f0/resourceGroups/rg-packer-through-cib/providers/Microsoft.Compute/images/win11-64-2009-centralus-win11-22h2-avd-69a7ecf | 1 |
| /subscriptions/a30e97ab-734a-4f3b-a0e4-c51c0bff0701/resourceGroups/rg-packer-through-cib/providers/Microsoft.Compute/images/trusted-win2012r2-64-l3-centralus-2012-R2-Datacenter-7cde253 | 1 |
| /subscriptions/108d46d5-fe9b-4850-9a7d-8c914aa6c1f0/resourceGroups/rg-packer-through-cib/providers/Microsoft.Compute/images/win2022-64-2009-centralus-2022-datacenter-azure-edition-alpha | 2 |
| /subscriptions/108d46d5-fe9b-4850-9a7d-8c914aa6c1f0/resourceGroups/rg-packer-through-cib/providers/Microsoft.Compute/images/win2012r2-64-vs-l1-centralus-2012-R2-Datacenter-alpha | 1 |
| /subscriptions/108d46d5-fe9b-4850-9a7d-8c914aa6c1f0/resourceGroups/rg-packer-through-cib/providers/Microsoft.Compute/images/win10-64-2009-ukwest-win10-22h2-avd-alpha | 1 |
| /subscriptions/108d46d5-fe9b-4850-9a7d-8c914aa6c1f0/resourceGroups/rg-packer-through-cib/providers/Microsoft.Compute/images/win2022-64-2009-centralus-2022-datacenter-azure-edition-531c001 | 2 |
| /subscriptions/108d46d5-fe9b-4850-9a7d-8c914aa6c1f0/resourceGroups/rg-packer-through-cib/providers/Microsoft.Compute/images/win2022-64-2009-rd-centralus-2022-datacenter-azure-edition-alpha | 1 |
|  | 2 |
| /subscriptions/108d46d5-fe9b-4850-9a7d-8c914aa6c1f0/resourceGroups/rg-packer-through-cib/providers/Microsoft.Compute/images/win11-64-2009-tc-centralus-win11-22h2-avd-alpha | 1 |
| /subscriptions/108d46d5-fe9b-4850-9a7d-8c914aa6c1f0/resourceGroups/rg-packer-through-cib/providers/Microsoft.Compute/images/win10-64-2009-centralus-win10-22h2-avd-alpha | 1 |
| /subscriptions/108d46d5-fe9b-4850-9a7d-8c914aa6c1f0/resourceGroups/rg-packer-through-cib/providers/Microsoft.Compute/images/win11-64-2009-canadacentral-win11-22h2-avd-531c001 | 2 |


| Worker Pool | Implementation | Version | Total Workers | Total Capacity |
| --- | --- | --- | ---: | ---: |
| **comm-1/b-win2012-azure** |  | Version not determined; task not (yet) claimed | 6 | 6 |
| **comm-2/b-win2012-azure** |  | Version not determined; task not (yet) claimed | 6 | 6 |
| **comm-t/win11-64-2009** |  | Version not determined; task not (yet) claimed | 1733 | 1733 |
| **gecko-1/b-win2012-azure** |  | Version not determined; task not (yet) claimed | 2 | 2 |
| **gecko-1/b-win2012-azure-alpha** |  | Version not determined; task not (yet) claimed | 6 | 6 |
| **gecko-1/b-win2012-azure-beta** |  | Version not determined; task not (yet) claimed | 6 | 6 |
| **gecko-1/b-win2022-alpha** |  | Version not determined; task not (yet) claimed | 6 | 6 |
| **gecko-1/b-win2022-rd-alpha** |  | Version not determined; task not (yet) claimed | 6 | 6 |
| **gecko-2/b-win2012-azure** |  | Version not determined; task not (yet) claimed | 6 | 6 |
| **gecko-2/b-win2022** |  | Version not determined; task not (yet) claimed | 53 | 53 |
| **gecko-3/b-win2012-azure** |  | Version not determined; task not (yet) claimed | 18 | 18 |
| **gecko-t/win10-64-2009-alpha** |  | Version not determined; task not (yet) claimed | 6 | 6 |
| **gecko-t/win10-64-2009-gpu-alpha** |  | Version not determined; task not (yet) claimed | 6 | 6 |
| **gecko-t/win10-64-2009-source-alpha** |  | Version not determined; task not (yet) claimed | 6 | 6 |
| **gecko-t/win11-64-2009** |  | Version not determined; task not (yet) claimed | 23211 | 23211 |
| **gecko-t/win11-64-2009-alpha-test** |  | Version not determined; task not (yet) claimed | 6 | 6 |
| **gecko-t/win11-64-2009-tc-alpha** |  | Version not determined; task not (yet) claimed | 7 | 7 |
| **mozillavpn-1/b-win2012-azure** |  | Version not determined; task not (yet) claimed | 7 | 7 |
| **mozillavpn-1/b-win2012-azure-alpha** |  | Version not determined; task not (yet) claimed | 7 | 7 |
| **mozillavpn-1/masterwayz-win2012-azure** |  | Version not determined; task not (yet) claimed | 7 | 7 |
| **nss-1/b-win2012-azure-alpha** |  | Version not determined; task not (yet) claimed | 7 | 7 |
| **nss-1/b-win2022-alpha** |  | Version not determined; task not (yet) claimed | 7 | 7 |
| **releng-hardware/gecko-1-b-osx-1015** |  | Version not determined; task not (yet) claimed | 0 | 0 |
| **releng-hardware/gecko-t-win10-64-ref-hw** |  | Version not determined; task not (yet) claimed | 0 | 0 |
| **relops-1/b-win2022** |  | Version not determined; task not (yet) claimed | 15 | 15 |
| **relops-3/b-win2022** |  | Version not determined; task not (yet) claimed | 15 | 15 |



[^1]: Those are the pools whose tasks were claimed and resolved by a worker as expected, but the worker did not publish either artifact `public/logs/live_backing.log` nor `public/logs/chain_of_trust.log`, which is the source used to identify the worker implementation.

[^2]: Probing task remains pending after two hours. Those are the pools that were not able to start any worker to claim the task within two hours.

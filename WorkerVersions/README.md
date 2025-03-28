

# Worker Pool Versions


## Generic Worker

Total: `148`

Count by version:

| Version | Count |
| :--- | ---: |
| 16.2.0 | 1 |
| 36.0.0 | 4 |
| 45.0.0 | 1 |
| 60.3.4 | 25 |
| 61.0.0 | 1 |
| 64.2.6 | 7 |
| 64.2.7 | 3 |
| 64.3.0 | 22 |
| 65.1.0 | 1 |
| 81.0.3 | 58 |
| 83.2.4 | 25 |


Count by image:

| Version | Count |
| :--- | ---: |
| projects/taskcluster-imaging/global/images/gw-fxci-gcp-l1-gui-googlecompute-2024-08-22t22-48-09z | 8 |
| /subscriptions/108d46d5-fe9b-4850-9a7d-8c914aa6c1f0/resourceGroups/rg-packer-worker-images/providers/Microsoft.Compute/galleries/win10_64_2009_alpha/images/win10_64_2009_alpha/versions/1.0.0 | 3 |
| /subscriptions/a30e97ab-734a-4f3b-a0e4-c51c0bff0701/resourceGroups/rg-packer-worker-images/providers/Microsoft.Compute/galleries/trusted_win2022_64_2009/images/trusted_win2022_64_2009/versions/1.0.3 | 4 |
| /subscriptions/108d46d5-fe9b-4850-9a7d-8c914aa6c1f0/resourceGroups/rg-packer-worker-images/providers/Microsoft.Compute/galleries/win11_64_2009/images/win11_64_2009/versions/1.0.3 | 7 |
| /subscriptions/108d46d5-fe9b-4850-9a7d-8c914aa6c1f0/resourceGroups/rg-packer-worker-images/providers/Microsoft.Compute/galleries/win10_64_2009/images/win10_64_2009/versions/1.0.3 | 3 |
| projects/taskcluster-imaging/global/images/gw-fxci-gcp-l1-2404-amd64-headless-googlecompute-alpha | 2 |
|  | 44 |
| projects/taskcluster-imaging/global/images/gw-translations-gcp-googlecompute-2024-04-22t18-22-42z | 7 |
| projects/taskcluster-imaging/global/images/gw-fxci-gcp-l1-2404-amd64-headless-googlecompute-2025-03-19 | 22 |
| /subscriptions/108d46d5-fe9b-4850-9a7d-8c914aa6c1f0/resourceGroups/rg-packer-worker-images/providers/Microsoft.Compute/galleries/win2022_64_2009/images/win2022_64_2009/versions/1.0.3 | 7 |
| /subscriptions/108d46d5-fe9b-4850-9a7d-8c914aa6c1f0/resourceGroups/rg-packer-worker-images/providers/Microsoft.Compute/galleries/win11_a64_24h2_builder/images/win11_a64_24h2_builder/versions/1.0.2 | 2 |
| projects/fxci-production-level3-workers/global/images/gw-fxci-gcp-l3-gui-googlecompute-2024-09-18t05-46-31z | 3 |
| projects/taskcluster-imaging/global/images/gw-fxci-gcp-l1-arm64-gui-googlecompute-2024-09-18t14-57-52z | 5 |
| /subscriptions/108d46d5-fe9b-4850-9a7d-8c914aa6c1f0/resourceGroups/rg-packer-worker-images/providers/Microsoft.Compute/galleries/win11_a64_24h2_tester_alpha/images/win11_a64_24h2_tester_alpha/versions/1.0.0 | 2 |
| /subscriptions/108d46d5-fe9b-4850-9a7d-8c914aa6c1f0/resourceGroups/rg-packer-worker-images/providers/Microsoft.Compute/galleries/win11_64_2009_alpha/images/win11_64_2009_alpha/versions/1.0.0 | 3 |
| /subscriptions/108d46d5-fe9b-4850-9a7d-8c914aa6c1f0/resourceGroups/rg-packer-worker-images/providers/Microsoft.Compute/galleries/win11_a64_24h2_tester/images/win11_a64_24h2_tester/versions/1.0.2 | 2 |
| /subscriptions/108d46d5-fe9b-4850-9a7d-8c914aa6c1f0/resourceGroups/rg-packer-worker-images/providers/Microsoft.Compute/galleries/win11_64_24h2/images/win11_64_24h2/versions/1.0.3 | 8 |
| /subscriptions/108d46d5-fe9b-4850-9a7d-8c914aa6c1f0/resourceGroups/rg-packer-worker-images/providers/Microsoft.Compute/galleries/win2022_64_2009_alpha/images/win2022_64_2009_alpha/versions/1.0.0 | 3 |
| /subscriptions/a30e97ab-734a-4f3b-a0e4-c51c0bff0701/resourceGroups/rg-packer-worker-images/providers/Microsoft.Compute/galleries/trusted_win11_a64_24h2_builder/images/trusted_win11_a64_24h2_builder/versions/1.0.2 | 1 |
| projects/taskcluster-imaging/global/images/gw-fxci-gcp-l1-2404-amd64-gui-googlecompute-alpha | 1 |
| projects/taskcluster-imaging/global/images/gw-fxci-gcp-l1-2404-arm64-headless-googlecompute-alpha | 1 |
| projects/taskcluster-imaging/global/images/gw-fxci-gcp-l1-gui-googlecompute-2025-01-13t22-33-40z | 3 |
| projects/taskcluster-imaging/global/images/gw-fxci-gcp-l1-2404-amd64-gui-googlecompute-2025-02-24 | 4 |
| projects/fxci-production-level3-workers/global/images/gw-fxci-gcp-l3-arm64-gui-googlecompute-2024-09-18t19-02-58z | 3 |


| Worker Pool | Implementation | Version | Engine | Revision | OS | Arch | GO | Total Workers | Total Capacity |
| --- | --- | --- | --- | --- | --- | --- | --- | ---: | ---: |
| **code-analysis-1/linux-gw-gcp** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | amd64 | 1.22.2 | 2 | 2 |
| **code-analysis-3/linux-gw-gcp** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | amd64 | 1.22.2 | 2 | 2 |
| **comm-1/b-win2022** | generic-worker | 81.0.3 | multiuser | da740f7e01 | windows | amd64 | 1.23.6 | 2 | 2 |
| **comm-2/b-win2022** | generic-worker | 81.0.3 | multiuser | da740f7e01 | windows | amd64 | 1.23.6 | 2 | 2 |
| **comm-3/b-win2022** | generic-worker | 81.0.3 | multiuser | da740f7e01 | windows | amd64 | 1.23.6 | 708 | 708 |
| **comm-t/t-linux-docker** | generic-worker | 83.2.4 | multiuser | 2f394bc1fe | linux | amd64 | 1.23.6 | 1 | 1 |
| **comm-t/t-linux-docker-noscratch** | generic-worker | 83.2.4 | multiuser | 2f394bc1fe | linux | amd64 | 1.23.6 | 1 | 1 |
| **comm-t/win11-64-2009** | generic-worker | 81.0.3 | multiuser | da740f7e01 | windows | amd64 | 1.23.6 | 4180 | 4180 |
| **comm-t/win11-64-2009-ssd** | generic-worker | 81.0.3 | multiuser | da740f7e01 | windows | amd64 | 1.23.6 | 337 | 337 |
| **comm-t/win11-64-24h2** | generic-worker | 81.0.3 | multiuser | da740f7e01 | windows | amd64 | 1.23.6 | 2 | 2 |
| **comm-t/win11-64-24h2-ssd** | generic-worker | 81.0.3 | multiuser | da740f7e01 | windows | amd64 | 1.23.6 | 2 | 2 |
| **comm-t/win11-a64-24h2** | generic-worker | 81.0.3 | multiuser | da740f7e01 | windows | amd64 | 1.23.6 | 2 | 2 |
| **comm-t/win11-a64-24h2-alpha** | generic-worker | 81.0.3 | multiuser | da740f7e01 | windows | amd64 | 1.23.6 | 10 | 10 |
| **gecko-1/b-linux-2204-kvm-gcp** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | amd64 | 1.22.2 | 2 | 2 |
| **gecko-1/b-linux-gcp-aarch64** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | arm64 | 1.22.2 | 14 | 14 |
| **gecko-1/b-win2022** | generic-worker | 81.0.3 | multiuser | da740f7e01 | windows | amd64 | 1.23.6 | 1605 | 1605 |
| **gecko-1/b-win2022-alpha** | generic-worker | 81.0.3 | multiuser | da740f7e01 | windows | amd64 | 1.23.6 | 19 | 19 |
| **gecko-1/images-gcp-aarch64** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | arm64 | 1.22.2 | 6 | 6 |
| **gecko-1/win11-a64-24h2-builder** | generic-worker | 81.0.3 | multiuser | da740f7e01 | windows | amd64 | 1.23.6 | 17 | 17 |
| **gecko-2/b-linux-gcp-aarch64** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | arm64 | 1.22.2 | 2 | 2 |
| **gecko-2/b-win2022** | generic-worker | 81.0.3 | multiuser | da740f7e01 | windows | amd64 | 1.23.6 | 14 | 14 |
| **gecko-2/images-gcp-aarch64** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | arm64 | 1.22.2 | 2 | 2 |
| **gecko-3/b-linux-2204-kvm-gcp** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | amd64 | 1.22.2 | 2 | 2 |
| **gecko-3/b-linux-gcp-aarch64** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | arm64 | 1.22.2 | 10 | 10 |
| **gecko-3/b-win2022** | generic-worker | 81.0.3 | multiuser | da740f7e01 | windows | amd64 | 1.23.6 | 5246 | 5246 |
| **gecko-3/images-gcp-aarch64** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | arm64 | 1.22.2 | 2 | 2 |
| **gecko-3/t-linux-2204-wayland-arm64-relsre** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | arm64 | 1.22.2 | 2 | 2 |
| **gecko-3/t-linux-2204-wayland-relsre** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | amd64 | 1.22.2 | 2 | 2 |
| **gecko-3/win11-a64-24h2-builder** | generic-worker | 81.0.3 | multiuser | da740f7e01 | windows | amd64 | 1.23.6 | 94 | 94 |
| **gecko-t/t-linux-2204-wayland** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | amd64 | 1.22.2 | 5581 | 5581 |
| **gecko-t/t-linux-2204-wayland-arm64-relsre** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | arm64 | 1.22.2 | 2 | 2 |
| **gecko-t/t-linux-2204-wayland-relsre** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | amd64 | 1.22.2 | 2 | 2 |
| **gecko-t/t-linux-2204-wayland-root-exp** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | amd64 | 1.22.2 | 2 | 2 |
| **gecko-t/t-linux-2204-wayland-snap** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | amd64 | 1.22.2 | 882 | 882 |
| **gecko-t/t-linux-2404-headless-alpha** | generic-worker | 83.2.4 | multiuser | 2f394bc1fe | linux | amd64 | 1.23.6 | 2 | 2 |
| **gecko-t/t-linux-2404-headless-arm64-alpha** | generic-worker | 83.2.4 | multiuser | 2f394bc1fe | linux | arm64 | 1.23.6 | 5 | 5 |
| **gecko-t/t-linux-2404-headless-ssd-alpha** | generic-worker | 83.2.4 | multiuser | 2f394bc1fe | linux | amd64 | 1.23.6 | 2 | 2 |
| **gecko-t/t-linux-2404-wayland** | generic-worker | 81.0.3 | multiuser | da740f7e01 | linux | amd64 | 1.23.6 | 2 | 2 |
| **gecko-t/t-linux-2404-wayland-relsre** | generic-worker | 81.0.3 | multiuser | da740f7e01 | linux | amd64 | 1.23.6 | 2 | 2 |
| **gecko-t/t-linux-2404-wayland-snap** | generic-worker | 81.0.3 | multiuser | da740f7e01 | linux | amd64 | 1.23.6 | 855 | 855 |
| **gecko-t/t-linux-docker** | generic-worker | 83.2.4 | multiuser | 2f394bc1fe | linux | amd64 | 1.23.6 | 1098 | 1098 |
| **gecko-t/t-linux-docker-noscratch** | generic-worker | 83.2.4 | multiuser | 2f394bc1fe | linux | amd64 | 1.23.6 | 7175 | 7175 |
| **gecko-t/t-linux-xlarge-2204-wayland** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | amd64 | 1.22.2 | 6638 | 6638 |
| **gecko-t/t-linux-xlarge-2404-wayland** | generic-worker | 81.0.3 | multiuser | da740f7e01 | linux | amd64 | 1.23.6 | 2 | 2 |
| **gecko-t/win10-64-2009** | generic-worker | 81.0.3 | multiuser | da740f7e01 | windows | amd64 | 1.23.6 | 3159 | 3159 |
| **gecko-t/win10-64-2009-alpha** | generic-worker | 81.0.3 | multiuser | da740f7e01 | windows | amd64 | 1.23.6 | 10 | 10 |
| **gecko-t/win10-64-2009-gpu** | generic-worker | 81.0.3 | multiuser | da740f7e01 | windows | amd64 | 1.23.6 | 429 | 429 |
| **gecko-t/win10-64-2009-gpu-alpha** | generic-worker | 81.0.3 | multiuser | da740f7e01 | windows | amd64 | 1.23.6 | 10 | 10 |
| **gecko-t/win10-64-2009-source** | generic-worker | 81.0.3 | multiuser | da740f7e01 | windows | amd64 | 1.23.6 | 2 | 2 |
| **gecko-t/win10-64-2009-source-alpha** | generic-worker | 81.0.3 | multiuser | da740f7e01 | windows | amd64 | 1.23.6 | 10 | 10 |
| **gecko-t/win11-64-2009** | generic-worker | 81.0.3 | multiuser | da740f7e01 | windows | amd64 | 1.23.6 | 3910 | 3910 |
| **gecko-t/win11-64-2009-alpha** | generic-worker | 81.0.3 | multiuser | da740f7e01 | windows | amd64 | 1.23.6 | 10 | 10 |
| **gecko-t/win11-64-2009-gpu** | generic-worker | 81.0.3 | multiuser | da740f7e01 | windows | amd64 | 1.23.6 | 2016 | 2016 |
| **gecko-t/win11-64-2009-gpu-alpha** | generic-worker | 81.0.3 | multiuser | da740f7e01 | windows | amd64 | 1.23.6 | 10 | 10 |
| **gecko-t/win11-64-2009-source** | generic-worker | 81.0.3 | multiuser | da740f7e01 | windows | amd64 | 1.23.6 | 398 | 398 |
| **gecko-t/win11-64-2009-source-alpha** | generic-worker | 81.0.3 | multiuser | da740f7e01 | windows | amd64 | 1.23.6 | 10 | 10 |
| **gecko-t/win11-64-2009-ssd** | generic-worker | 81.0.3 | multiuser | da740f7e01 | windows | amd64 | 1.23.6 | 2 | 2 |
| **gecko-t/win11-64-2009-ssd-gpu** | generic-worker | 81.0.3 | multiuser | da740f7e01 | windows | amd64 | 1.23.6 | 4 | 4 |
| **gecko-t/win11-64-24h2** | generic-worker | 81.0.3 | multiuser | da740f7e01 | windows | amd64 | 1.23.6 | 30765 | 30765 |
| **gecko-t/win11-64-24h2-gpu** | generic-worker | 81.0.3 | multiuser | da740f7e01 | windows | amd64 | 1.23.6 | 15569 | 15569 |
| **gecko-t/win11-64-24h2-large** | generic-worker | 81.0.3 | multiuser | da740f7e01 | windows | amd64 | 1.23.6 | 4160 | 4160 |
| **gecko-t/win11-64-24h2-source** | generic-worker | 81.0.3 | multiuser | da740f7e01 | windows | amd64 | 1.23.6 | 1093 | 1093 |
| **gecko-t/win11-64-24h2-ssd** | generic-worker | 81.0.3 | multiuser | da740f7e01 | windows | amd64 | 1.23.6 | 2 | 2 |
| **gecko-t/win11-64-24h2-ssd-gpu** | generic-worker | 81.0.3 | multiuser | da740f7e01 | windows | amd64 | 1.23.6 | 2 | 2 |
| **gecko-t/win11-a64-24h2** | generic-worker | 81.0.3 | multiuser | da740f7e01 | windows | amd64 | 1.23.6 | 37 | 37 |
| **gecko-t/win11-a64-24h2-alpha** | generic-worker | 81.0.3 | multiuser | da740f7e01 | windows | amd64 | 1.23.6 | 12 | 12 |
| **mozilla-1/b-win2022** | generic-worker | 81.0.3 | multiuser | da740f7e01 | windows | amd64 | 1.23.6 | 2 | 2 |
| **mozilla-t/t-linux-2204-wayland** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | amd64 | 1.22.2 | 290 | 290 |
| **mozilla-t/t-linux-2204-wayland-relsre** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | amd64 | 1.22.2 | 2 | 2 |
| **mozilla-t/t-linux-2404-wayland** | generic-worker | 81.0.3 | multiuser | da740f7e01 | linux | amd64 | 1.23.6 | 2 | 2 |
| **mozilla-t/t-linux-docker** | generic-worker | 83.2.4 | multiuser | 2f394bc1fe | linux | amd64 | 1.23.6 | 1 | 1 |
| **mozilla-t/t-linux-docker-noscratch** | generic-worker | 83.2.4 | multiuser | 2f394bc1fe | linux | amd64 | 1.23.6 | 1 | 1 |
| **mozillavpn-1/b-linux-gcp-gw** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | amd64 | 1.22.2 | 2 | 2 |
| **mozillavpn-1/b-win2022** | generic-worker | 81.0.3 | multiuser | da740f7e01 | windows | amd64 | 1.23.6 | 50 | 50 |
| **mozillavpn-1/b-win2022-alpha** | generic-worker | 81.0.3 | multiuser | da740f7e01 | windows | amd64 | 1.23.6 | 3 | 3 |
| **mozillavpn-1/win11-a64-24h2-builder** | generic-worker | 81.0.3 | multiuser | da740f7e01 | windows | amd64 | 1.23.6 | 12 | 12 |
| **mozillavpn-3/b-linux-gcp-gw** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | amd64 | 1.22.2 | 2 | 2 |
| **mozillavpn-3/b-win2022** | generic-worker | 81.0.3 | multiuser | da740f7e01 | windows | amd64 | 1.23.6 | 5 | 5 |
| **nss-1/b-win2022** | generic-worker | 81.0.3 | multiuser | da740f7e01 | windows | amd64 | 1.23.6 | 57 | 57 |
| **nss-1/b-win2022-alpha** | generic-worker | 81.0.3 | multiuser | da740f7e01 | windows | amd64 | 1.23.6 | 2 | 2 |
| **nss-3/b-win2022** | generic-worker | 81.0.3 | multiuser | da740f7e01 | windows | amd64 | 1.23.6 | 26 | 26 |
| **nss-t/t-linux-docker** | generic-worker | 83.2.4 | multiuser | 2f394bc1fe | linux | amd64 | 1.23.6 | 1 | 1 |
| **proj-autophone/gecko-t-bitbar-gw-perf-a55** | generic-worker | 36.0.0 | simple | b9cb11293f | linux | amd64 | 1.13.7 | 0 | 0 |
| **proj-autophone/gecko-t-bitbar-gw-perf-p6** | generic-worker | 36.0.0 | simple | b9cb11293f | linux | amd64 | 1.13.7 | 0 | 0 |
| **proj-autophone/gecko-t-bitbar-gw-perf-s24** | generic-worker | 36.0.0 | simple | b9cb11293f | linux | amd64 | 1.13.7 | 0 | 0 |
| **proj-autophone/gecko-t-bitbar-gw-unit-p5** | generic-worker | 36.0.0 | simple | b9cb11293f | linux | amd64 | 1.13.7 | 0 | 0 |
| **releng-hardware/applicationservices-b-1-osx1015** | generic-worker | 60.3.4 | multiuser | 943a6f2b0d | darwin | amd64 | 1.22.0 | 0 | 0 |
| **releng-hardware/applicationservices-b-3-osx1015** | generic-worker | 60.3.4 | multiuser | 943a6f2b0d | darwin | amd64 | 1.22.0 | 0 | 0 |
| **releng-hardware/gecko-1-b-osx-1015** | generic-worker | 60.3.4 | multiuser | 943a6f2b0d | darwin | amd64 | 1.22.0 | 0 | 0 |
| **releng-hardware/gecko-1-b-osx-1015-staging** | generic-worker | 60.3.4 | multiuser | 943a6f2b0d | darwin | amd64 | 1.22.0 | 0 | 0 |
| **releng-hardware/gecko-1-b-osx-arm64** | generic-worker | 60.3.4 | multiuser | 943a6f2b0d | darwin | arm64 | 1.22.0 | 0 | 0 |
| **releng-hardware/gecko-3-b-osx-1015** | generic-worker | 60.3.4 | multiuser | 943a6f2b0d | darwin | amd64 | 1.22.0 | 0 | 0 |
| **releng-hardware/gecko-t-linux-netperf-1804** | generic-worker | 65.1.0 | insecure | 1a085daa37 | linux | amd64 | 1.22.3 | 0 | 0 |
| **releng-hardware/gecko-t-linux-talos-1804** | generic-worker | 61.0.0 | simple | 3bd4419b4b | linux | amd64 | 1.22.1 | 0 | 0 |
| **releng-hardware/gecko-t-osx-1015-r8** | generic-worker | 60.3.4 | simple | 943a6f2b0d | darwin | amd64 | 1.22.0 | 0 | 0 |
| **releng-hardware/gecko-t-osx-1015-r8-staging** | generic-worker | 60.3.4 | simple | 943a6f2b0d | darwin | amd64 | 1.22.0 | 0 | 0 |
| **releng-hardware/gecko-t-osx-1100-m1** | generic-worker | 60.3.4 | simple | 943a6f2b0d | darwin | arm64 | 1.22.0 | 0 | 0 |
| **releng-hardware/gecko-t-osx-1100-m1-staging** | generic-worker | 60.3.4 | simple | 943a6f2b0d | darwin | arm64 | 1.22.0 | 0 | 0 |
| **releng-hardware/gecko-t-osx-1100-r8-latest** | generic-worker | 60.3.4 | simple | 943a6f2b0d | darwin | amd64 | 1.22.0 | 0 | 0 |
| **releng-hardware/gecko-t-osx-1200-r8-latest** | generic-worker | 60.3.4 | simple | 943a6f2b0d | darwin | amd64 | 1.22.0 | 0 | 0 |
| **releng-hardware/gecko-t-osx-1300-r8-latest** | generic-worker | 60.3.4 | simple | 943a6f2b0d | darwin | amd64 | 1.22.0 | 0 | 0 |
| **releng-hardware/gecko-t-osx-1400-m2** | generic-worker | 60.3.4 | simple | 943a6f2b0d | darwin | arm64 | 1.22.0 | 0 | 0 |
| **releng-hardware/gecko-t-osx-1400-m2-staging** | generic-worker | 60.3.4 | simple | 943a6f2b0d | darwin | arm64 | 1.22.0 | 0 | 0 |
| **releng-hardware/gecko-t-osx-1400-r8** | generic-worker | 60.3.4 | simple | 943a6f2b0d | darwin | amd64 | 1.22.0 | 0 | 0 |
| **releng-hardware/gecko-t-osx-1400-r8-latest** | generic-worker | 60.3.4 | simple | 943a6f2b0d | darwin | amd64 | 1.22.0 | 0 | 0 |
| **releng-hardware/gecko-t-osx-1400-r8-staging** | generic-worker | 60.3.4 | simple | 943a6f2b0d | darwin | amd64 | 1.22.0 | 0 | 0 |
| **releng-hardware/gecko-t-osx-1500-m-vms** | generic-worker | 60.3.4 | simple | 943a6f2b0d | darwin | arm64 | 1.22.0 | 0 | 0 |
| **releng-hardware/gecko-t-win10-64-1803-hw** | generic-worker | 16.2.0 | multiuser | b321a877c8 | windows | amd64 | 1.10.8 | 0 | 0 |
| **releng-hardware/gecko-t-win7-32-hw** | generic-worker | 45.0.0 | multiuser | 988e8100b3 | windows | 386 | 1.19.3 | 0 | 0 |
| **releng-hardware/mozilla-b-1-osx** | generic-worker | 60.3.4 | multiuser | 943a6f2b0d | darwin | arm64 | 1.22.0 | 0 | 0 |
| **releng-hardware/mozilla-b-3-osx** | generic-worker | 60.3.4 | multiuser | 943a6f2b0d | darwin | arm64 | 1.22.0 | 0 | 0 |
| **releng-hardware/mozillavpn-b-1-osx** | generic-worker | 60.3.4 | multiuser | 943a6f2b0d | darwin | amd64 | 1.22.0 | 0 | 0 |
| **releng-hardware/mozillavpn-b-3-osx** | generic-worker | 60.3.4 | multiuser | 943a6f2b0d | darwin | amd64 | 1.22.0 | 0 | 0 |
| **releng-hardware/nss-1-b-osx-1015** | generic-worker | 60.3.4 | multiuser | 943a6f2b0d | darwin | amd64 | 1.22.0 | 0 | 0 |
| **releng-hardware/nss-3-b-osx-1015** | generic-worker | 60.3.4 | multiuser | 943a6f2b0d | darwin | amd64 | 1.22.0 | 0 | 0 |
| **releng-hardware/win10-64-2009-hw** | generic-worker | 81.0.3 | multiuser | da740f7e01 | windows | amd64 | 1.23.6 | 0 | 0 |
| **releng-hardware/win10-64-2009-hw-alpha** | generic-worker | 81.0.3 | multiuser | da740f7e01 | windows | amd64 | 1.23.6 | 0 | 0 |
| **releng-hardware/win11-64-2009-hw-ref** | generic-worker | 81.0.3 | multiuser | da740f7e01 | windows | amd64 | 1.23.6 | 0 | 0 |
| **releng-hardware/win11-64-24h2-hw** | generic-worker | 81.0.3 | multiuser | da740f7e01 | windows | amd64 | 1.23.6 | 0 | 0 |
| **releng-hardware/win11-64-24h2-hw-alpha** | generic-worker | 81.0.3 | multiuser | da740f7e01 | windows | amd64 | 1.23.6 | 0 | 0 |
| **releng-hardware/win11-64-24h2-hw-perf-sheriff** | generic-worker | 81.0.3 | multiuser | da740f7e01 | windows | amd64 | 1.23.6 | 0 | 0 |
| **releng-hardware/win11-64-24h2-hw-ref** | generic-worker | 81.0.3 | multiuser | da740f7e01 | windows | amd64 | 1.23.6 | 0 | 0 |
| **releng-hardware/win11-64-24h2-hw-ref-alpha** | generic-worker | 81.0.3 | multiuser | da740f7e01 | windows | amd64 | 1.23.6 | 0 | 0 |
| **translations-1/b-linux-large-gcp-1tb-32-256-d2g** | generic-worker | 83.2.4 | multiuser | 2f394bc1fe | linux | amd64 | 1.23.6 | 20 | 20 |
| **translations-1/b-linux-large-gcp-1tb-32-256-std-d2g** | generic-worker | 83.2.4 | multiuser | 2f394bc1fe | linux | amd64 | 1.23.6 | 7 | 7 |
| **translations-1/b-linux-large-gcp-1tb-64-512-d2g** | generic-worker | 83.2.4 | multiuser | 2f394bc1fe | linux | amd64 | 1.23.6 | 3 | 3 |
| **translations-1/b-linux-large-gcp-1tb-64-512-std-d2g** | generic-worker | 83.2.4 | multiuser | 2f394bc1fe | linux | amd64 | 1.23.6 | 2 | 2 |
| **translations-1/b-linux-large-gcp-d2g** | generic-worker | 83.2.4 | multiuser | 2f394bc1fe | linux | amd64 | 1.23.6 | 361 | 361 |
| **translations-1/b-linux-large-gcp-d2g-1tb** | generic-worker | 83.2.4 | multiuser | 2f394bc1fe | linux | amd64 | 1.23.6 | 2 | 2 |
| **translations-1/b-linux-large-gcp-d2g-1tb-standard** | generic-worker | 83.2.4 | multiuser | 2f394bc1fe | linux | amd64 | 1.23.6 | 2 | 2 |
| **translations-1/b-linux-large-gcp-d2g-300gb** | generic-worker | 83.2.4 | multiuser | 2f394bc1fe | linux | amd64 | 1.23.6 | 85 | 85 |
| **translations-1/b-linux-v100-gpu** | generic-worker | 64.2.6 | insecure | edab196d7d | linux | amd64 | 1.22.2 | 151 | 151 |
| **translations-1/b-linux-v100-gpu-4** | generic-worker | 64.2.6 | insecure | edab196d7d | linux | amd64 | 1.22.2 | 45 | 45 |
| **translations-1/b-linux-v100-gpu-4-1tb** | generic-worker | 64.2.6 | insecure | edab196d7d | linux | amd64 | 1.22.2 | 11 | 11 |
| **translations-1/b-linux-v100-gpu-4-1tb-standard** | generic-worker | 64.2.6 | insecure | edab196d7d | linux | amd64 | 1.22.2 | 4 | 4 |
| **translations-1/b-linux-v100-gpu-4-2tb** | generic-worker | 64.2.6 | insecure | edab196d7d | linux | amd64 | 1.22.2 | 27 | 27 |
| **translations-1/b-linux-v100-gpu-4-300gb** | generic-worker | 64.2.6 | insecure | edab196d7d | linux | amd64 | 1.22.2 | 96 | 96 |
| **translations-1/b-linux-v100-gpu-4-300gb-standard** | generic-worker | 64.2.6 | insecure | edab196d7d | linux | amd64 | 1.22.2 | 3 | 3 |
| **translations-1/b-linux-v100-gpu-d2g** | generic-worker | 83.2.4 | multiuser | 2f394bc1fe | linux | amd64 | 1.23.6 | 5 | 5 |
| **translations-1/b-linux-v100-gpu-d2g-4** | generic-worker | 83.2.4 | multiuser | 2f394bc1fe | linux | amd64 | 1.23.6 | 2 | 2 |
| **translations-1/b-linux-v100-gpu-d2g-4-1tb** | generic-worker | 83.2.4 | multiuser | 2f394bc1fe | linux | amd64 | 1.23.6 | 2 | 2 |
| **translations-1/b-linux-v100-gpu-d2g-4-1tb-standard** | generic-worker | 83.2.4 | multiuser | 2f394bc1fe | linux | amd64 | 1.23.6 | 2 | 2 |
| **translations-1/b-linux-v100-gpu-d2g-4-2tb** | generic-worker | 83.2.4 | multiuser | 2f394bc1fe | linux | amd64 | 1.23.6 | 2 | 2 |
| **translations-1/b-linux-v100-gpu-d2g-4-300gb** | generic-worker | 83.2.4 | multiuser | 2f394bc1fe | linux | amd64 | 1.23.6 | 6 | 6 |
| **translations-1/b-linux-v100-gpu-d2g-4-300gb-standard** | generic-worker | 83.2.4 | multiuser | 2f394bc1fe | linux | amd64 | 1.23.6 | 2 | 2 |
| **translations-1/snakepit-2080ti** | generic-worker | 64.2.7 | insecure | 6f28121e31 | linux | amd64 | 1.22.2 | 0 | 0 |
| **translations-1/snakepit-quadro6k** | generic-worker | 64.2.7 | insecure | 6f28121e31 | linux | amd64 | 1.22.2 | 0 | 0 |
| **translations-1/snakepit-titanxp** | generic-worker | 64.2.7 | insecure | 6f28121e31 | linux | amd64 | 1.22.2 | 0 | 0 |


## Docker Worker

Total: `153`

Count by version:

| Version | Count |
| :--- | ---: |
| 38.0.5 | 128 |
| 48.3.0 | 1 |
| unknown version | 24 |


Count by image:

| Version | Count |
| :--- | ---: |
| projects/taskcluster-imaging/global/images/docker-worker-gcp-u14-04-2024-06-14 | 24 |
| projects/taskcluster-imaging/global/images/docker-firefoxci-gcp-lt-googlecompute-2023-04-13t21-30-28z | 1 |
| projects/taskcluster-imaging/global/images/docker-firefoxci-gcp-l1-googlecompute-2024-06-14t19-52-20z | 78 |
| projects/fxci-production-level3-workers/global/images/docker-firefoxci-gcp-l3-googlecompute-2024-02-05t23-18-22z | 50 |


| Worker Pool | Implementation | Version | Total Workers | Total Capacity |
| --- | --- | --- | ---: | ---: |
| **adhoc-1/b-linux-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **adhoc-1/decision-gcp** | docker-worker | 38.0.5 | 2 | 4 |
| **adhoc-1/images-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **adhoc-3/b-linux-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **adhoc-3/decision-gcp** | docker-worker | 38.0.5 | 2 | 4 |
| **adhoc-3/images-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **app-services-1/b-linux-gcp** | docker-worker | 38.0.5 | 1017 | 1017 |
| **app-services-1/decision-gcp** | docker-worker | 38.0.5 | 761 | 1522 |
| **app-services-1/images-gcp** | docker-worker | 38.0.5 | 3 | 3 |
| **app-services-3/b-linux-gcp** | docker-worker | 38.0.5 | 1192 | 1192 |
| **app-services-3/decision-gcp** | docker-worker | 38.0.5 | 31 | 62 |
| **app-services-3/images-gcp** | docker-worker | 38.0.5 | 3 | 3 |
| **ci-1/decision-gcp** | docker-worker | 38.0.5 | 22 | 44 |
| **ci-1/images-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **ci-3/decision-gcp** | docker-worker | 38.0.5 | 12 | 24 |
| **ci-3/images-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **ci-t/linux-gcp** | docker-worker | 38.0.5 | 30 | 30 |
| **code-analysis-1/linux-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **code-analysis-3/linux-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **code-coverage/bot-gcp** | docker-worker | 38.0.5 | 867 | 867 |
| **code-review/bot-gcp** | docker-worker | 38.0.5 | 466 | 932 |
| **comm-1/b-linux-gcp** | docker-worker | 38.0.5 | 1279 | 1279 |
| **comm-1/b-linux-large-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **comm-1/b-linux-xlarge-gcp** | docker-worker | 38.0.5 | 7 | 7 |
| **comm-1/decision-gcp** | docker-worker | 38.0.5 | 231 | 462 |
| **comm-1/images-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **comm-2/b-linux-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **comm-2/b-linux-large-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **comm-2/b-linux-xlarge-gcp** | docker-worker | 38.0.5 | 3 | 3 |
| **comm-2/decision-gcp** | docker-worker | 38.0.5 | 2 | 4 |
| **comm-2/images-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **comm-3/b-linux-gcp** | docker-worker | 38.0.5 | 6353 | 6353 |
| **comm-3/b-linux-large-gcp** | docker-worker | 38.0.5 | 4 | 4 |
| **comm-3/b-linux-xlarge-gcp** | docker-worker | 38.0.5 | 7 | 7 |
| **comm-3/decision-gcp** | docker-worker | 38.0.5 | 121 | 242 |
| **comm-3/images-gcp** | docker-worker | 38.0.5 | 22 | 22 |
| **comm-t/misc-gcp** | docker-worker | 38.0.5 | 2 | 16 |
| **comm-t/t-linux-large-gcp** | docker-worker | unknown version | 20 | 20 |
| **comm-t/t-linux-large-noscratch-gcp** | docker-worker | unknown version | 3097 | 3097 |
| **comm-t/t-linux-xlarge-gcp** | docker-worker | unknown version | 2 | 2 |
| **comm-t/t-linux-xlarge-noscratch-gcp** | docker-worker | unknown version | 387 | 387 |
| **comm-t/t-linux-xlarge-source-gcp** | docker-worker | unknown version | 1894 | 1894 |
| **comm-t/t-linux-xlarge-source-noscratch-gcp** | docker-worker | unknown version | 2 | 2 |
| **gecko-1/b-linux-gcp** | docker-worker | 38.0.5 | 23656 | 23656 |
| **gecko-1/b-linux-gcp-test-bug-1882320** | docker-worker | 38.0.5 | 2 | 2 |
| **gecko-1/b-linux-kvm-gcp** | docker-worker | 38.0.5 | 426 | 426 |
| **gecko-1/b-linux-large-gcp** | docker-worker | 38.0.5 | 100 | 100 |
| **gecko-1/b-linux-medium-gcp** | docker-worker | 38.0.5 | 18023 | 18023 |
| **gecko-1/b-linux-xlarge-gcp** | docker-worker | 38.0.5 | 2196 | 2196 |
| **gecko-1/b-linux-xlarge-gcp-bug1797804-c2** | docker-worker | 38.0.5 | 7 | 7 |
| **gecko-1/decision-gcp** | docker-worker | 38.0.5 | 958 | 1916 |
| **gecko-1/images-gcp** | docker-worker | 38.0.5 | 118 | 118 |
| **gecko-2/b-linux-gcp** | docker-worker | 38.0.5 | 611 | 611 |
| **gecko-2/b-linux-kvm-gcp** | docker-worker | 38.0.5 | 28 | 28 |
| **gecko-2/b-linux-large-gcp** | docker-worker | 38.0.5 | 16 | 16 |
| **gecko-2/b-linux-xlarge-gcp** | docker-worker | 38.0.5 | 88 | 88 |
| **gecko-2/decision-gcp** | docker-worker | 38.0.5 | 26 | 52 |
| **gecko-2/images-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **gecko-3/b-linux-gcp** | docker-worker | 38.0.5 | 60544 | 60544 |
| **gecko-3/b-linux-kvm-gcp** | docker-worker | 38.0.5 | 723 | 723 |
| **gecko-3/b-linux-large-gcp** | docker-worker | 38.0.5 | 222 | 222 |
| **gecko-3/b-linux-medium-gcp** | docker-worker | 38.0.5 | 14052 | 14052 |
| **gecko-3/b-linux-xlarge-gcp** | docker-worker | 38.0.5 | 4692 | 4692 |
| **gecko-3/decision-gcp** | docker-worker | 38.0.5 | 1524 | 3048 |
| **gecko-3/images-gcp** | docker-worker | 38.0.5 | 15 | 15 |
| **gecko-t/misc-gcp** | docker-worker | 38.0.5 | 748 | 5984 |
| **gecko-t/t-linux-kvm-gcp** | docker-worker | unknown version | 2 | 2 |
| **gecko-t/t-linux-kvm-gcp-bug1862675** | docker-worker | unknown version | 2 | 2 |
| **gecko-t/t-linux-kvm-noscratch-gcp** | docker-worker | unknown version | 72255 | 72255 |
| **gecko-t/t-linux-kvm-noscratch-gcp-bug1862675** | docker-worker | unknown version | 3 | 3 |
| **gecko-t/t-linux-large-gcp** | docker-worker | unknown version | 1472 | 1472 |
| **gecko-t/t-linux-large-hostub2204-gcp** | docker-worker | 48.3.0 | 2 | 2 |
| **gecko-t/t-linux-large-noscratch-gcp** | docker-worker | unknown version | 119518 | 119518 |
| **gecko-t/t-linux-xlarge-gcp** | docker-worker | unknown version | 175 | 175 |
| **gecko-t/t-linux-xlarge-gcp-b1862675** | docker-worker | unknown version | 2 | 2 |
| **gecko-t/t-linux-xlarge-noscratch-gcp** | docker-worker | unknown version | 65225 | 65225 |
| **gecko-t/t-linux-xlarge-ns-gcp-b1862675** | docker-worker | unknown version | 2 | 2 |
| **gecko-t/t-linux-xlarge-source-gcp** | docker-worker | unknown version | 33723 | 33723 |
| **gecko-t/t-linux-xlarge-source-gcp-b1862675** | docker-worker | unknown version | 2 | 2 |
| **gecko-t/t-linux-xlarge-source-noscratch-gcp** | docker-worker | unknown version | 2 | 2 |
| **gecko-t/t-linux-xlarge-source-ns-gcp-b1862675** | docker-worker | unknown version | 2 | 2 |
| **glean-1/b-linux-gcp** | docker-worker | 38.0.5 | 16 | 16 |
| **glean-1/decision-gcp** | docker-worker | 38.0.5 | 768 | 1536 |
| **glean-1/images-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **glean-3/b-linux-gcp** | docker-worker | 38.0.5 | 10 | 10 |
| **glean-3/decision-gcp** | docker-worker | 38.0.5 | 4 | 8 |
| **glean-3/images-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **infra/build-decision** | docker-worker | 38.0.5 | 1 | 8 |
| **mobile-1/b-linux-gcp** | docker-worker | 38.0.5 | 204 | 204 |
| **mobile-1/decision-gcp** | docker-worker | 38.0.5 | 705 | 1410 |
| **mobile-1/images-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **mobile-3/b-linux-gcp** | docker-worker | 38.0.5 | 312 | 312 |
| **mobile-3/decision-gcp** | docker-worker | 38.0.5 | 127 | 254 |
| **mobile-3/images-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **mozilla-1/b-linux-gcp** | docker-worker | 38.0.5 | 4 | 4 |
| **mozilla-1/decision-gcp** | docker-worker | 38.0.5 | 753 | 1506 |
| **mozilla-1/images-gcp** | docker-worker | 38.0.5 | 10 | 10 |
| **mozilla-3/b-linux-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **mozilla-3/decision-gcp** | docker-worker | 38.0.5 | 217 | 434 |
| **mozilla-3/images-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **mozilla-t/t-linux-large-gcp** | docker-worker | unknown version | 266 | 266 |
| **mozilla-t/t-linux-large-noscratch-gcp** | docker-worker | unknown version | 2 | 2 |
| **mozillaonline-1/b-linux-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **mozillaonline-3/b-linux-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **mozillavpn-1/b-linux-gcp** | docker-worker | 38.0.5 | 247 | 247 |
| **mozillavpn-1/b-linux-large-gcp** | docker-worker | 38.0.5 | 166 | 166 |
| **mozillavpn-1/decision-gcp** | docker-worker | 38.0.5 | 763 | 1526 |
| **mozillavpn-1/images-gcp** | docker-worker | 38.0.5 | 12 | 12 |
| **mozillavpn-3/b-linux-gcp** | docker-worker | 38.0.5 | 11 | 11 |
| **mozillavpn-3/b-linux-large-gcp** | docker-worker | 38.0.5 | 10 | 10 |
| **mozillavpn-3/decision-gcp** | docker-worker | 38.0.5 | 4 | 8 |
| **mozillavpn-3/images-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **nss-1/decision-gcp** | docker-worker | 38.0.5 | 11 | 22 |
| **nss-1/images-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **nss-1/linux-gcp** | docker-worker | 38.0.5 | 179 | 179 |
| **nss-3/decision-gcp** | docker-worker | 38.0.5 | 4 | 8 |
| **nss-3/images-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **nss-3/linux-gcp** | docker-worker | 38.0.5 | 75 | 75 |
| **nss-t/t-linux-xlarge-gcp** | docker-worker | unknown version | 308 | 308 |
| **nss-t/t-linux-xlarge-noscratch-gcp** | docker-worker | unknown version | 2 | 2 |
| **releng-1/b-linux-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **releng-1/decision-gcp** | docker-worker | 38.0.5 | 15 | 30 |
| **releng-1/linux-gcp** | docker-worker | 38.0.5 | 22 | 22 |
| **releng-3/b-linux-gcp** | docker-worker | 38.0.5 | 5 | 5 |
| **releng-3/decision-gcp** | docker-worker | 38.0.5 | 11 | 22 |
| **releng-3/linux-gcp** | docker-worker | 38.0.5 | 26 | 26 |
| **releng-t/linux-gcp** | docker-worker | 38.0.5 | 99 | 99 |
| **relops-3/decision-gcp** | docker-worker | 38.0.5 | 2 | 4 |
| **scriptworker-1/b-linux-gcp** | docker-worker | 38.0.5 | 78 | 78 |
| **scriptworker-1/decision-gcp** | docker-worker | 38.0.5 | 13 | 26 |
| **scriptworker-1/images-gcp** | docker-worker | 38.0.5 | 38 | 38 |
| **scriptworker-3/b-linux-gcp** | docker-worker | 38.0.5 | 41 | 41 |
| **scriptworker-3/decision-gcp** | docker-worker | 38.0.5 | 7 | 14 |
| **scriptworker-3/images-gcp** | docker-worker | 38.0.5 | 15 | 15 |
| **taskgraph-1/decision-gcp** | docker-worker | 38.0.5 | 5 | 10 |
| **taskgraph-1/images-gcp** | docker-worker | 38.0.5 | 5 | 5 |
| **taskgraph-3/decision-gcp** | docker-worker | 38.0.5 | 4 | 8 |
| **taskgraph-3/images-gcp** | docker-worker | 38.0.5 | 3 | 3 |
| **taskgraph-t/linux-gcp** | docker-worker | 38.0.5 | 42 | 42 |
| **translations-1/b-linux-large-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **translations-1/b-linux-large-gcp-1tb-32-256** | docker-worker | 38.0.5 | 2 | 2 |
| **translations-1/b-linux-large-gcp-1tb-32-256-standard** | docker-worker | 38.0.5 | 2 | 2 |
| **translations-1/b-linux-large-gcp-1tb-64-512** | docker-worker | 38.0.5 | 2 | 2 |
| **translations-1/b-linux-large-gcp-1tb-64-512-standard** | docker-worker | 38.0.5 | 2 | 2 |
| **translations-1/b-linux-large-gcp-300gb** | docker-worker | 38.0.5 | 2 | 2 |
| **translations-1/decision-gcp** | docker-worker | 38.0.5 | 1713 | 3426 |
| **translations-1/images-gcp** | docker-worker | 38.0.5 | 12 | 12 |
| **xpi-1/b-linux-gcp** | docker-worker | 38.0.5 | 89 | 89 |
| **xpi-1/decision-gcp** | docker-worker | 38.0.5 | 28 | 56 |
| **xpi-1/images-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **xpi-3/b-linux-gcp** | docker-worker | 38.0.5 | 3 | 3 |
| **xpi-3/decision-gcp** | docker-worker | 38.0.5 | 5 | 10 |
| **xpi-3/images-gcp** | docker-worker | 38.0.5 | 2 | 2 |


## Script Worker

Total: `35`



| Worker Pool | Implementation | Version | Total Workers | Total Capacity |
| --- | --- | --- | ---: | ---: |
| **scriptworker-k8s/app-services-3-beetmover** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/app-services-3-signing** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/comm-1-pushmsix** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/comm-3-balrog** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/comm-3-beetmover** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/comm-3-bouncer** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/comm-3-pushmsix** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/comm-3-signing** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/comm-3-tree** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/comm-t-signing** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/gecko-1-balrog** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/gecko-3-addon** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/gecko-3-balrog** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/gecko-3-beetmover** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/gecko-3-bouncer** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/gecko-3-pushapk** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/gecko-3-pushflatpak** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/gecko-3-pushmsix** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/gecko-3-shipit** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/gecko-3-signing** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/gecko-3-tree** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/gecko-t-signing** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/mobile-3-bitrise** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/mobile-3-pushapk** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/mobile-3-signing** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/mobile-t-signing** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
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
| **gecko-t/t-linux-vm-2204-wayland** |  | No artifacts found | 1008 | 1008 |
| **gecko-t/t-linux-vm-2204-wayland-snap** |  | No artifacts found | 2 | 2 |
| **scriptworker-prov-v1/mac-notarization-poller** |  | No artifacts found | 0 | 0 |
| **scriptworker-prov-v1/tb-mac-notarization-poller** |  | No artifacts found | 0 | 0 |


## Version not determined [^2]

Total: `20`


Count by image:

| Version | Count |
| :--- | ---: |
|  | 10 |
| /subscriptions/a30e97ab-734a-4f3b-a0e4-c51c0bff0701/resourceGroups/rg-packer-worker-images/providers/Microsoft.Compute/galleries/trusted_win2022_64_2009/images/trusted_win2022_64_2009/versions/1.0.3 | 2 |
| /subscriptions/108d46d5-fe9b-4850-9a7d-8c914aa6c1f0/resourceGroups/rg-packer-worker-images/providers/Microsoft.Compute/galleries/win2022_64_2009/images/win2022_64_2009/versions/1.0.3 | 1 |
| /subscriptions/108d46d5-fe9b-4850-9a7d-8c914aa6c1f0/resourceGroups/rg-packer-worker-images/providers/Microsoft.Compute/galleries/win11_64_24h2_alpha/images/win11_64_24h2_alpha/versions/1.0.0 | 4 |
| projects/taskcluster-imaging/global/images/gw-fxci-gcp-l1-2404-amd64-googlecompute-alpha | 1 |
| /subscriptions/108d46d5-fe9b-4850-9a7d-8c914aa6c1f0/resourceGroups/rg-packer-worker-images/providers/Microsoft.Compute/galleries/win11_a64_24h2_builder_alpha/images/win11_a64_24h2_builder_alpha/versions/1.0.0 | 1 |
| /subscriptions/108d46d5-fe9b-4850-9a7d-8c914aa6c1f0/resourceGroups/rg-packer-through-cib/providers/Microsoft.Compute/images/win2012r2-64-l1-centralus-2012-R2-Datacenter-75bd9ed,/subscriptions/108d46d5-fe9b-4850-9a7d-8c914aa6c1f0/resourceGroups/rg-packer-through-cib/providers/Microsoft.Compute/images/win2012r2-64-l1-eastus2-2012-R2-Datacenter-75bd9ed,/subscriptions/108d46d5-fe9b-4850-9a7d-8c914aa6c1f0/resourceGroups/rg-packer-through-cib/providers/Microsoft.Compute/images/win2012r2-64-l1-northcentralus-2012-R2-Datacenter-75bd9ed,/subscriptions/108d46d5-fe9b-4850-9a7d-8c914aa6c1f0/resourceGroups/rg-packer-through-cib/providers/Microsoft.Compute/images/win2012r2-64-l1-westus2-2012-R2-Datacenter-75bd9ed | 1 |


| Worker Pool | Implementation | Version | Total Workers | Total Capacity |
| --- | --- | --- | ---: | ---: |
| **gecko-1/b-win2012-azure** |  | Version not determined; task not (yet) claimed | 2 | 2 |
| **gecko-1/win11-a64-24h2-builder-alpha** |  | Version not determined; task not (yet) claimed | 11 | 11 |
| **gecko-t/t-linux-2404-relsre** |  | Version not determined; task not (yet) claimed | 5 | 5 |
| **gecko-t/win11-64-24h2-alpha** |  | Version not determined; task not (yet) claimed | 947 | 947 |
| **gecko-t/win11-64-24h2-gpu-alpha** |  | Version not determined; task not (yet) claimed | 346 | 346 |
| **gecko-t/win11-64-24h2-large-alpha** |  | Version not determined; task not (yet) claimed | 14 | 14 |
| **gecko-t/win11-64-24h2-source-alpha** |  | Version not determined; task not (yet) claimed | 346 | 346 |
| **mozilla-3/b-win2022** |  | Version not determined; task not (yet) claimed | 215 | 215 |
| **proj-autophone/gecko-t-lambda-alpha-a55** |  | Version not determined; task not (yet) claimed | 0 | 0 |
| **proj-autophone/gecko-t-lambda-perf-a55** |  | Version not determined; task not (yet) claimed | 0 | 0 |
| **releng-hardware/gecko-3-b-osx-arm64** |  | Version not determined; task not (yet) claimed | 0 | 0 |
| **releng-hardware/gecko-t-osx-1500-m4-staging** |  | Version not determined; task not (yet) claimed | 0 | 0 |
| **releng-hardware/win11-64-2009-hw** |  | Version not determined; task not (yet) claimed | 0 | 0 |
| **relops-1/b-win2022** |  | Version not determined; task not (yet) claimed | 122 | 122 |
| **relops-3/b-win2022** |  | Version not determined; task not (yet) claimed | 127 | 127 |
| **scriptworker-prov-v1/dep-adhoc-signing-mac14m2** |  | Version not determined; task not (yet) claimed | 0 | 0 |
| **scriptworker-prov-v1/dep-comm-signing-mac14m2** |  | Version not determined; task not (yet) claimed | 0 | 0 |
| **scriptworker-prov-v1/dep-gecko-signing-mac14m2** |  | Version not determined; task not (yet) claimed | 0 | 0 |
| **scriptworker-prov-v1/dep-mozillavpn-signing-mac14m2** |  | Version not determined; task not (yet) claimed | 0 | 0 |
| **scriptworker-prov-v1/gecko-signing-mac14m2** |  | Version not determined; task not (yet) claimed | 0 | 0 |



[^1]: Those are the pools whose tasks were claimed and resolved by a worker as expected, but the worker did not publish either artifact `public/logs/live_backing.log` nor `public/logs/chain_of_trust.log`, which is the source used to identify the worker implementation.

[^2]: Probing task remains pending after two hours. Those are the pools that were not able to start any worker to claim the task within two hours.

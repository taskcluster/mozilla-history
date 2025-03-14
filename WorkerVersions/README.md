

# Worker Pool Versions


## Generic Worker

Total: `139`

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
| 72.1.0 | 3 |
| 77.3.1 | 3 |
| 81.0.3 | 66 |
| 83.2.4 | 2 |


Count by image:

| Version | Count |
| :--- | ---: |
| projects/taskcluster-imaging/global/images/gw-fxci-gcp-l1-gui-googlecompute-2024-08-22t22-48-09z | 8 |
| projects/taskcluster-imaging/global/images/gw-fxci-gcp-l1-2404-amd64-headless-googlecompute-alpha | 2 |
| /subscriptions/108d46d5-fe9b-4850-9a7d-8c914aa6c1f0/resourceGroups/rg-packer-worker-images/providers/Microsoft.Compute/galleries/win11_64_24h2_alpha/images/win11_64_24h2_alpha/versions/1.0.0 | 4 |
| projects/taskcluster-imaging/global/images/gw-translations-gcp-googlecompute-2024-04-22t18-22-42z | 7 |
| projects/taskcluster-imaging/global/images/gw-fxci-gcp-l1-2404-amd64-headless-googlecompute-2025-02-21 | 15 |
| /subscriptions/a30e97ab-734a-4f3b-a0e4-c51c0bff0701/resourceGroups/rg-packer-worker-images/providers/Microsoft.Compute/galleries/trusted_win2022_64_2009/images/trusted_win2022_64_2009/versions/1.0.2 | 4 |
| projects/fxci-production-level3-workers/global/images/gw-fxci-gcp-l3-gui-googlecompute-2024-09-18t05-46-31z | 3 |
| projects/taskcluster-imaging/global/images/gw-fxci-gcp-l1-2404-amd64-gui-googlecompute-2025-02-24 | 4 |
| /subscriptions/108d46d5-fe9b-4850-9a7d-8c914aa6c1f0/resourceGroups/rg-packer-worker-images/providers/Microsoft.Compute/galleries/win10_64_2009/images/win10_64_2009/versions/1.0.2 | 2 |
| /subscriptions/108d46d5-fe9b-4850-9a7d-8c914aa6c1f0/resourceGroups/rg-packer-worker-images/providers/Microsoft.Compute/galleries/win10_64_2009_alpha/images/win10_64_2009_alpha/versions/1.0.0 | 3 |
| /subscriptions/108d46d5-fe9b-4850-9a7d-8c914aa6c1f0/resourceGroups/rg-packer-worker-images/providers/Microsoft.Compute/galleries/win11_64_2009_alpha/images/win11_64_2009_alpha/versions/1.0.0 | 3 |
| /subscriptions/a30e97ab-734a-4f3b-a0e4-c51c0bff0701/resourceGroups/rg-packer-worker-images/providers/Microsoft.Compute/galleries/trusted_win11_a64_24h2_builder/images/trusted_win11_a64_24h2_builder/versions/1.0.2 | 1 |
| /subscriptions/108d46d5-fe9b-4850-9a7d-8c914aa6c1f0/resourceGroups/rg-packer-worker-images/providers/Microsoft.Compute/galleries/win2022_64_2009/images/win2022_64_2009/versions/1.0.2 | 7 |
|  | 44 |
| projects/taskcluster-imaging/global/images/gw-fxci-gcp-l1-arm64-gui-googlecompute-2024-09-18t14-57-52z | 5 |
| projects/fxci-production-level3-workers/global/images/gw-fxci-gcp-l3-arm64-gui-googlecompute-2024-09-18t19-02-58z | 3 |
| projects/taskcluster-imaging/global/images/gw-fxci-gcp-l1-2404-amd64-gui-googlecompute-alpha | 1 |
| projects/taskcluster-imaging/global/images/gw-fxci-gcp-l1-gui-googlecompute-2025-01-13t22-33-40z | 3 |
| /subscriptions/108d46d5-fe9b-4850-9a7d-8c914aa6c1f0/resourceGroups/rg-packer-worker-images/providers/Microsoft.Compute/galleries/win11_a64_24h2_tester/images/win11_a64_24h2_tester/versions/1.0.2 | 2 |
| /subscriptions/108d46d5-fe9b-4850-9a7d-8c914aa6c1f0/resourceGroups/rg-packer-worker-images/providers/Microsoft.Compute/galleries/win11_a64_24h2_builder_alpha/images/win11_a64_24h2_builder_alpha/versions/1.0.0 | 1 |
| /subscriptions/108d46d5-fe9b-4850-9a7d-8c914aa6c1f0/resourceGroups/rg-packer-worker-images/providers/Microsoft.Compute/galleries/win11_64_24h2/images/win11_64_24h2/versions/1.0.2 | 8 |
| /subscriptions/108d46d5-fe9b-4850-9a7d-8c914aa6c1f0/resourceGroups/rg-packer-worker-images/providers/Microsoft.Compute/galleries/win11_64_2009/images/win11_64_2009/versions/1.0.2 | 7 |
| /subscriptions/108d46d5-fe9b-4850-9a7d-8c914aa6c1f0/resourceGroups/rg-packer-worker-images/providers/Microsoft.Compute/galleries/win11_a64_24h2_builder/images/win11_a64_24h2_builder/versions/1.0.2 | 2 |


| Worker Pool | Implementation | Version | Engine | Revision | OS | Arch | GO | Total Workers | Total Capacity |
| --- | --- | --- | --- | --- | --- | --- | --- | ---: | ---: |
| **code-analysis-1/linux-gw-gcp** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | amd64 | 1.22.2 | 19 | 19 |
| **code-analysis-3/linux-gw-gcp** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | amd64 | 1.22.2 | 2 | 2 |
| **comm-1/b-win2022** | generic-worker | 81.0.3 | multiuser | da740f7e01 | windows | amd64 | 1.23.6 | 2 | 2 |
| **comm-2/b-win2022** | generic-worker | 81.0.3 | multiuser | da740f7e01 | windows | amd64 | 1.23.6 | 2 | 2 |
| **comm-3/b-win2022** | generic-worker | 81.0.3 | multiuser | da740f7e01 | windows | amd64 | 1.23.6 | 616 | 616 |
| **comm-t/win11-64-2009** | generic-worker | 81.0.3 | multiuser | da740f7e01 | windows | amd64 | 1.23.6 | 3305 | 3305 |
| **comm-t/win11-64-2009-ssd** | generic-worker | 81.0.3 | multiuser | da740f7e01 | windows | amd64 | 1.23.6 | 376 | 376 |
| **comm-t/win11-64-24h2** | generic-worker | 81.0.3 | multiuser | da740f7e01 | windows | amd64 | 1.23.6 | 116 | 116 |
| **comm-t/win11-64-24h2-ssd** | generic-worker | 81.0.3 | multiuser | da740f7e01 | windows | amd64 | 1.23.6 | 2 | 2 |
| **comm-t/win11-a64-24h2** | generic-worker | 81.0.3 | multiuser | da740f7e01 | windows | amd64 | 1.23.6 | 2 | 2 |
| **gecko-1/b-linux-2204-kvm-gcp** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | amd64 | 1.22.2 | 2 | 2 |
| **gecko-1/b-linux-gcp-aarch64** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | arm64 | 1.22.2 | 10 | 10 |
| **gecko-1/b-win2022** | generic-worker | 81.0.3 | multiuser | da740f7e01 | windows | amd64 | 1.23.6 | 543 | 543 |
| **gecko-1/images-gcp-aarch64** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | arm64 | 1.22.2 | 4 | 4 |
| **gecko-1/win11-a64-24h2-builder** | generic-worker | 81.0.3 | multiuser | da740f7e01 | windows | amd64 | 1.23.6 | 21 | 21 |
| **gecko-1/win11-a64-24h2-builder-alpha** | generic-worker | 81.0.3 | multiuser | da740f7e01 | windows | amd64 | 1.23.6 | 2 | 2 |
| **gecko-2/b-linux-gcp-aarch64** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | arm64 | 1.22.2 | 2 | 2 |
| **gecko-2/b-win2022** | generic-worker | 81.0.3 | multiuser | da740f7e01 | windows | amd64 | 1.23.6 | 16 | 16 |
| **gecko-2/images-gcp-aarch64** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | arm64 | 1.22.2 | 2 | 2 |
| **gecko-3/b-linux-2204-kvm-gcp** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | amd64 | 1.22.2 | 2 | 2 |
| **gecko-3/b-linux-gcp-aarch64** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | arm64 | 1.22.2 | 2 | 2 |
| **gecko-3/b-win2022** | generic-worker | 81.0.3 | multiuser | da740f7e01 | windows | amd64 | 1.23.6 | 4690 | 4690 |
| **gecko-3/images-gcp-aarch64** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | arm64 | 1.22.2 | 2 | 2 |
| **gecko-3/t-linux-2204-wayland-arm64-relsre** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | arm64 | 1.22.2 | 2 | 2 |
| **gecko-3/t-linux-2204-wayland-relsre** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | amd64 | 1.22.2 | 2 | 2 |
| **gecko-3/win11-a64-24h2-builder** | generic-worker | 81.0.3 | multiuser | da740f7e01 | windows | amd64 | 1.23.6 | 149 | 149 |
| **gecko-t/t-linux-2204-wayland** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | amd64 | 1.22.2 | 5233 | 5233 |
| **gecko-t/t-linux-2204-wayland-arm64-relsre** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | arm64 | 1.22.2 | 2 | 2 |
| **gecko-t/t-linux-2204-wayland-relsre** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | amd64 | 1.22.2 | 2 | 2 |
| **gecko-t/t-linux-2204-wayland-root-exp** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | amd64 | 1.22.2 | 2 | 2 |
| **gecko-t/t-linux-2204-wayland-snap** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | amd64 | 1.22.2 | 297 | 297 |
| **gecko-t/t-linux-2404-headless-alpha** | generic-worker | 83.2.4 | multiuser | 2f394bc1fe | linux | amd64 | 1.23.6 | 263 | 263 |
| **gecko-t/t-linux-2404-headless-ssd-alpha** | generic-worker | 83.2.4 | multiuser | 2f394bc1fe | linux | amd64 | 1.23.6 | 109 | 109 |
| **gecko-t/t-linux-2404-wayland** | generic-worker | 81.0.3 | multiuser | da740f7e01 | linux | amd64 | 1.23.6 | 2 | 2 |
| **gecko-t/t-linux-2404-wayland-relsre** | generic-worker | 81.0.3 | multiuser | da740f7e01 | linux | amd64 | 1.23.6 | 2 | 2 |
| **gecko-t/t-linux-2404-wayland-snap** | generic-worker | 81.0.3 | multiuser | da740f7e01 | linux | amd64 | 1.23.6 | 270 | 270 |
| **gecko-t/t-linux-xlarge-2204-wayland** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | amd64 | 1.22.2 | 6120 | 6120 |
| **gecko-t/t-linux-xlarge-2404-wayland** | generic-worker | 81.0.3 | multiuser | da740f7e01 | linux | amd64 | 1.23.6 | 2 | 2 |
| **gecko-t/win10-64-2009-alpha** | generic-worker | 81.0.3 | multiuser | da740f7e01 | windows | amd64 | 1.23.6 | 2 | 2 |
| **gecko-t/win10-64-2009-gpu** | generic-worker | 81.0.3 | multiuser | da740f7e01 | windows | amd64 | 1.23.6 | 449 | 449 |
| **gecko-t/win10-64-2009-gpu-alpha** | generic-worker | 81.0.3 | multiuser | da740f7e01 | windows | amd64 | 1.23.6 | 2 | 2 |
| **gecko-t/win10-64-2009-source** | generic-worker | 81.0.3 | multiuser | da740f7e01 | windows | amd64 | 1.23.6 | 2 | 2 |
| **gecko-t/win10-64-2009-source-alpha** | generic-worker | 81.0.3 | multiuser | da740f7e01 | windows | amd64 | 1.23.6 | 2 | 2 |
| **gecko-t/win11-64-2009** | generic-worker | 81.0.3 | multiuser | da740f7e01 | windows | amd64 | 1.23.6 | 3037 | 3037 |
| **gecko-t/win11-64-2009-alpha** | generic-worker | 81.0.3 | multiuser | da740f7e01 | windows | amd64 | 1.23.6 | 2 | 2 |
| **gecko-t/win11-64-2009-gpu** | generic-worker | 81.0.3 | multiuser | da740f7e01 | windows | amd64 | 1.23.6 | 3138 | 3138 |
| **gecko-t/win11-64-2009-gpu-alpha** | generic-worker | 81.0.3 | multiuser | da740f7e01 | windows | amd64 | 1.23.6 | 3 | 3 |
| **gecko-t/win11-64-2009-source** | generic-worker | 81.0.3 | multiuser | da740f7e01 | windows | amd64 | 1.23.6 | 462 | 462 |
| **gecko-t/win11-64-2009-source-alpha** | generic-worker | 81.0.3 | multiuser | da740f7e01 | windows | amd64 | 1.23.6 | 2 | 2 |
| **gecko-t/win11-64-2009-ssd** | generic-worker | 81.0.3 | multiuser | da740f7e01 | windows | amd64 | 1.23.6 | 2 | 2 |
| **gecko-t/win11-64-2009-ssd-gpu** | generic-worker | 81.0.3 | multiuser | da740f7e01 | windows | amd64 | 1.23.6 | 4 | 4 |
| **gecko-t/win11-64-24h2** | generic-worker | 81.0.3 | multiuser | da740f7e01 | windows | amd64 | 1.23.6 | 34728 | 34728 |
| **gecko-t/win11-64-24h2-alpha** | generic-worker | 81.0.3 | multiuser | da740f7e01 | windows | amd64 | 1.23.6 | 372 | 372 |
| **gecko-t/win11-64-24h2-gpu** | generic-worker | 81.0.3 | multiuser | da740f7e01 | windows | amd64 | 1.23.6 | 16974 | 16974 |
| **gecko-t/win11-64-24h2-gpu-alpha** | generic-worker | 81.0.3 | multiuser | da740f7e01 | windows | amd64 | 1.23.6 | 87 | 87 |
| **gecko-t/win11-64-24h2-large** | generic-worker | 81.0.3 | multiuser | da740f7e01 | windows | amd64 | 1.23.6 | 4469 | 4469 |
| **gecko-t/win11-64-24h2-large-alpha** | generic-worker | 81.0.3 | multiuser | da740f7e01 | windows | amd64 | 1.23.6 | 121 | 121 |
| **gecko-t/win11-64-24h2-source** | generic-worker | 81.0.3 | multiuser | da740f7e01 | windows | amd64 | 1.23.6 | 1095 | 1095 |
| **gecko-t/win11-64-24h2-source-alpha** | generic-worker | 81.0.3 | multiuser | da740f7e01 | windows | amd64 | 1.23.6 | 184 | 184 |
| **gecko-t/win11-64-24h2-ssd** | generic-worker | 81.0.3 | multiuser | da740f7e01 | windows | amd64 | 1.23.6 | 2 | 2 |
| **gecko-t/win11-64-24h2-ssd-gpu** | generic-worker | 81.0.3 | multiuser | da740f7e01 | windows | amd64 | 1.23.6 | 2 | 2 |
| **gecko-t/win11-a64-24h2** | generic-worker | 81.0.3 | multiuser | da740f7e01 | windows | amd64 | 1.23.6 | 26 | 26 |
| **mozilla-1/b-win2022** | generic-worker | 81.0.3 | multiuser | da740f7e01 | windows | amd64 | 1.23.6 | 2 | 2 |
| **mozilla-t/t-linux-2204-wayland** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | amd64 | 1.22.2 | 285 | 285 |
| **mozilla-t/t-linux-2204-wayland-relsre** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | amd64 | 1.22.2 | 2 | 2 |
| **mozilla-t/t-linux-2404-wayland** | generic-worker | 81.0.3 | multiuser | da740f7e01 | linux | amd64 | 1.23.6 | 2 | 2 |
| **mozillavpn-1/b-linux-gcp-gw** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | amd64 | 1.22.2 | 2 | 2 |
| **mozillavpn-1/b-win2022** | generic-worker | 81.0.3 | multiuser | da740f7e01 | windows | amd64 | 1.23.6 | 43 | 43 |
| **mozillavpn-1/win11-a64-24h2-builder** | generic-worker | 81.0.3 | multiuser | da740f7e01 | windows | amd64 | 1.23.6 | 3 | 3 |
| **mozillavpn-3/b-linux-gcp-gw** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | amd64 | 1.22.2 | 2 | 2 |
| **mozillavpn-3/b-win2022** | generic-worker | 81.0.3 | multiuser | da740f7e01 | windows | amd64 | 1.23.6 | 27 | 27 |
| **nss-1/b-win2022** | generic-worker | 81.0.3 | multiuser | da740f7e01 | windows | amd64 | 1.23.6 | 119 | 119 |
| **nss-3/b-win2022** | generic-worker | 81.0.3 | multiuser | da740f7e01 | windows | amd64 | 1.23.6 | 103 | 103 |
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
| **releng-hardware/gecko-3-b-osx-arm64** | generic-worker | 60.3.4 | multiuser | 943a6f2b0d | darwin | arm64 | 1.22.0 | 0 | 0 |
| **releng-hardware/gecko-t-linux-netperf-1804** | generic-worker | 65.1.0 | insecure | 1a085daa37 | linux | amd64 | 1.22.3 | 0 | 0 |
| **releng-hardware/gecko-t-linux-talos-1804** | generic-worker | 61.0.0 | simple | 3bd4419b4b | linux | amd64 | 1.22.1 | 0 | 0 |
| **releng-hardware/gecko-t-osx-1015-r8-staging** | generic-worker | 60.3.4 | simple | 943a6f2b0d | darwin | amd64 | 1.22.0 | 0 | 0 |
| **releng-hardware/gecko-t-osx-1100-m1** | generic-worker | 60.3.4 | simple | 943a6f2b0d | darwin | arm64 | 1.22.0 | 0 | 0 |
| **releng-hardware/gecko-t-osx-1100-m1-staging** | generic-worker | 60.3.4 | simple | 943a6f2b0d | darwin | arm64 | 1.22.0 | 0 | 0 |
| **releng-hardware/gecko-t-osx-1100-r8-latest** | generic-worker | 60.3.4 | simple | 943a6f2b0d | darwin | amd64 | 1.22.0 | 0 | 0 |
| **releng-hardware/gecko-t-osx-1200-r8-latest** | generic-worker | 60.3.4 | simple | 943a6f2b0d | darwin | amd64 | 1.22.0 | 0 | 0 |
| **releng-hardware/gecko-t-osx-1300-r8-latest** | generic-worker | 60.3.4 | simple | 943a6f2b0d | darwin | amd64 | 1.22.0 | 0 | 0 |
| **releng-hardware/gecko-t-osx-1400-m2** | generic-worker | 60.3.4 | simple | 943a6f2b0d | darwin | arm64 | 1.22.0 | 0 | 0 |
| **releng-hardware/gecko-t-osx-1400-m2-staging** | generic-worker | 60.3.4 | simple | 943a6f2b0d | darwin | arm64 | 1.22.0 | 0 | 0 |
| **releng-hardware/gecko-t-osx-1400-r8-latest** | generic-worker | 60.3.4 | simple | 943a6f2b0d | darwin | amd64 | 1.22.0 | 0 | 0 |
| **releng-hardware/gecko-t-osx-1400-r8-staging** | generic-worker | 60.3.4 | simple | 943a6f2b0d | darwin | amd64 | 1.22.0 | 0 | 0 |
| **releng-hardware/gecko-t-osx-1500-m-vms** | generic-worker | 60.3.4 | simple | 943a6f2b0d | darwin | arm64 | 1.22.0 | 0 | 0 |
| **releng-hardware/gecko-t-osx-1500-m4-staging** | generic-worker | 60.3.4 | simple | 943a6f2b0d | darwin | arm64 | 1.22.0 | 0 | 0 |
| **releng-hardware/gecko-t-win10-64-1803-hw** | generic-worker | 16.2.0 | multiuser | b321a877c8 | windows | amd64 | 1.10.8 | 0 | 0 |
| **releng-hardware/gecko-t-win7-32-hw** | generic-worker | 45.0.0 | multiuser | 988e8100b3 | windows | 386 | 1.19.3 | 0 | 0 |
| **releng-hardware/mozilla-b-1-osx** | generic-worker | 60.3.4 | multiuser | 943a6f2b0d | darwin | arm64 | 1.22.0 | 0 | 0 |
| **releng-hardware/mozilla-b-3-osx** | generic-worker | 60.3.4 | multiuser | 943a6f2b0d | darwin | arm64 | 1.22.0 | 0 | 0 |
| **releng-hardware/mozillavpn-b-1-osx** | generic-worker | 60.3.4 | multiuser | 943a6f2b0d | darwin | amd64 | 1.22.0 | 0 | 0 |
| **releng-hardware/mozillavpn-b-3-osx** | generic-worker | 60.3.4 | multiuser | 943a6f2b0d | darwin | amd64 | 1.22.0 | 0 | 0 |
| **releng-hardware/nss-1-b-osx-1015** | generic-worker | 60.3.4 | multiuser | 943a6f2b0d | darwin | amd64 | 1.22.0 | 0 | 0 |
| **releng-hardware/nss-3-b-osx-1015** | generic-worker | 60.3.4 | multiuser | 943a6f2b0d | darwin | amd64 | 1.22.0 | 0 | 0 |
| **releng-hardware/win10-64-2009-hw** | generic-worker | 72.1.0 | multiuser | 114df53f48 | windows | amd64 | 1.23.2 | 0 | 0 |
| **releng-hardware/win11-64-2009-hw** | generic-worker | 72.1.0 | multiuser | 114df53f48 | windows | amd64 | 1.23.2 | 0 | 0 |
| **releng-hardware/win11-64-2009-hw-ref** | generic-worker | 72.1.0 | multiuser | 114df53f48 | windows | amd64 | 1.23.2 | 0 | 0 |
| **releng-hardware/win11-64-24h2-hw** | generic-worker | 77.3.1 | multiuser | 959a204190 | windows | amd64 | 1.23.4 | 0 | 0 |
| **releng-hardware/win11-64-24h2-hw-alpha** | generic-worker | 81.0.3 | multiuser | da740f7e01 | windows | amd64 | 1.23.6 | 0 | 0 |
| **releng-hardware/win11-64-24h2-hw-perf-sheriff** | generic-worker | 77.3.1 | multiuser | 959a204190 | windows | amd64 | 1.23.4 | 0 | 0 |
| **releng-hardware/win11-64-24h2-hw-ref** | generic-worker | 77.3.1 | multiuser | 959a204190 | windows | amd64 | 1.23.4 | 0 | 0 |
| **releng-hardware/win11-64-24h2-hw-ref-alpha** | generic-worker | 81.0.3 | multiuser | da740f7e01 | windows | amd64 | 1.23.6 | 0 | 0 |
| **translations-1/b-linux-large-gcp-1tb-32-256-d2g** | generic-worker | 81.0.3 | multiuser | da740f7e01 | linux | amd64 | 1.23.6 | 16 | 16 |
| **translations-1/b-linux-large-gcp-1tb-32-256-std-d2g** | generic-worker | 81.0.3 | multiuser | da740f7e01 | linux | amd64 | 1.23.6 | 5 | 5 |
| **translations-1/b-linux-large-gcp-1tb-64-512-d2g** | generic-worker | 81.0.3 | multiuser | da740f7e01 | linux | amd64 | 1.23.6 | 2 | 2 |
| **translations-1/b-linux-large-gcp-1tb-64-512-std-d2g** | generic-worker | 81.0.3 | multiuser | da740f7e01 | linux | amd64 | 1.23.6 | 2 | 2 |
| **translations-1/b-linux-large-gcp-d2g** | generic-worker | 81.0.3 | multiuser | da740f7e01 | linux | amd64 | 1.23.6 | 150 | 150 |
| **translations-1/b-linux-large-gcp-d2g-1tb** | generic-worker | 81.0.3 | multiuser | da740f7e01 | linux | amd64 | 1.23.6 | 2 | 2 |
| **translations-1/b-linux-large-gcp-d2g-1tb-standard** | generic-worker | 81.0.3 | multiuser | da740f7e01 | linux | amd64 | 1.23.6 | 2 | 2 |
| **translations-1/b-linux-large-gcp-d2g-300gb** | generic-worker | 81.0.3 | multiuser | da740f7e01 | linux | amd64 | 1.23.6 | 136 | 136 |
| **translations-1/b-linux-v100-gpu** | generic-worker | 64.2.6 | insecure | edab196d7d | linux | amd64 | 1.22.2 | 67 | 67 |
| **translations-1/b-linux-v100-gpu-4** | generic-worker | 64.2.6 | insecure | edab196d7d | linux | amd64 | 1.22.2 | 30 | 30 |
| **translations-1/b-linux-v100-gpu-4-1tb** | generic-worker | 64.2.6 | insecure | edab196d7d | linux | amd64 | 1.22.2 | 6 | 6 |
| **translations-1/b-linux-v100-gpu-4-1tb-standard** | generic-worker | 64.2.6 | insecure | edab196d7d | linux | amd64 | 1.22.2 | 2 | 2 |
| **translations-1/b-linux-v100-gpu-4-2tb** | generic-worker | 64.2.6 | insecure | edab196d7d | linux | amd64 | 1.22.2 | 6 | 6 |
| **translations-1/b-linux-v100-gpu-4-300gb** | generic-worker | 64.2.6 | insecure | edab196d7d | linux | amd64 | 1.22.2 | 68 | 68 |
| **translations-1/b-linux-v100-gpu-4-300gb-standard** | generic-worker | 64.2.6 | insecure | edab196d7d | linux | amd64 | 1.22.2 | 2 | 2 |
| **translations-1/b-linux-v100-gpu-d2g** | generic-worker | 81.0.3 | multiuser | da740f7e01 | linux | amd64 | 1.23.6 | 11 | 11 |
| **translations-1/b-linux-v100-gpu-d2g-4** | generic-worker | 81.0.3 | multiuser | da740f7e01 | linux | amd64 | 1.23.6 | 48 | 48 |
| **translations-1/b-linux-v100-gpu-d2g-4-1tb** | generic-worker | 81.0.3 | multiuser | da740f7e01 | linux | amd64 | 1.23.6 | 28 | 28 |
| **translations-1/b-linux-v100-gpu-d2g-4-1tb-standard** | generic-worker | 81.0.3 | multiuser | da740f7e01 | linux | amd64 | 1.23.6 | 2 | 2 |
| **translations-1/b-linux-v100-gpu-d2g-4-2tb** | generic-worker | 81.0.3 | multiuser | da740f7e01 | linux | amd64 | 1.23.6 | 3 | 3 |
| **translations-1/b-linux-v100-gpu-d2g-4-300gb** | generic-worker | 81.0.3 | multiuser | da740f7e01 | linux | amd64 | 1.23.6 | 15 | 15 |
| **translations-1/b-linux-v100-gpu-d2g-4-300gb-standard** | generic-worker | 81.0.3 | multiuser | da740f7e01 | linux | amd64 | 1.23.6 | 2 | 2 |
| **translations-1/snakepit-2080ti** | generic-worker | 64.2.7 | insecure | 6f28121e31 | linux | amd64 | 1.22.2 | 0 | 0 |
| **translations-1/snakepit-quadro6k** | generic-worker | 64.2.7 | insecure | 6f28121e31 | linux | amd64 | 1.22.2 | 0 | 0 |
| **translations-1/snakepit-titanxp** | generic-worker | 64.2.7 | insecure | 6f28121e31 | linux | amd64 | 1.22.2 | 0 | 0 |


## Docker Worker

Total: `155`

Count by version:

| Version | Count |
| :--- | ---: |
| 38.0.5 | 130 |
| 48.3.0 | 1 |
| unknown version | 24 |


Count by image:

| Version | Count |
| :--- | ---: |
| projects/taskcluster-imaging/global/images/docker-firefoxci-gcp-lt-googlecompute-2023-04-13t21-30-28z | 1 |
| projects/fxci-production-level3-workers/global/images/docker-firefoxci-gcp-l3-googlecompute-2024-02-05t23-18-22z | 52 |
| projects/taskcluster-imaging/global/images/docker-firefoxci-gcp-l1-googlecompute-2024-06-14t19-52-20z | 78 |
| projects/taskcluster-imaging/global/images/docker-worker-gcp-u14-04-2024-06-14 | 24 |


| Worker Pool | Implementation | Version | Total Workers | Total Capacity |
| --- | --- | --- | ---: | ---: |
| **adhoc-1/b-linux-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **adhoc-1/decision-gcp** | docker-worker | 38.0.5 | 2 | 4 |
| **adhoc-1/images-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **adhoc-3/b-linux-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **adhoc-3/decision-gcp** | docker-worker | 38.0.5 | 2 | 4 |
| **adhoc-3/images-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **app-services-1/b-linux-gcp** | docker-worker | 38.0.5 | 424 | 424 |
| **app-services-1/decision-gcp** | docker-worker | 38.0.5 | 1513 | 3026 |
| **app-services-1/images-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **app-services-3/b-linux-gcp** | docker-worker | 38.0.5 | 513 | 513 |
| **app-services-3/decision-gcp** | docker-worker | 38.0.5 | 47 | 94 |
| **app-services-3/images-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **ci-1/decision-gcp** | docker-worker | 38.0.5 | 99 | 198 |
| **ci-1/images-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **ci-3/decision-gcp** | docker-worker | 38.0.5 | 5 | 10 |
| **ci-3/images-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **ci-t/linux-gcp** | docker-worker | 38.0.5 | 15 | 15 |
| **code-analysis-1/linux-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **code-analysis-3/linux-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **code-coverage/bot-gcp** | docker-worker | 38.0.5 | 845 | 845 |
| **code-review/bot-gcp** | docker-worker | 38.0.5 | 1123 | 2246 |
| **comm-1/b-linux-gcp** | docker-worker | 38.0.5 | 792 | 792 |
| **comm-1/b-linux-large-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **comm-1/b-linux-xlarge-gcp** | docker-worker | 38.0.5 | 3 | 3 |
| **comm-1/decision-gcp** | docker-worker | 38.0.5 | 265 | 530 |
| **comm-1/images-gcp** | docker-worker | 38.0.5 | 24 | 24 |
| **comm-2/b-linux-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **comm-2/b-linux-large-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **comm-2/b-linux-xlarge-gcp** | docker-worker | 38.0.5 | 3 | 3 |
| **comm-2/decision-gcp** | docker-worker | 38.0.5 | 2 | 4 |
| **comm-2/images-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **comm-3/b-linux-gcp** | docker-worker | 38.0.5 | 4915 | 4915 |
| **comm-3/b-linux-large-gcp** | docker-worker | 38.0.5 | 8 | 8 |
| **comm-3/b-linux-xlarge-gcp** | docker-worker | 38.0.5 | 12 | 12 |
| **comm-3/decision-gcp** | docker-worker | 38.0.5 | 132 | 264 |
| **comm-3/images-gcp** | docker-worker | 38.0.5 | 8 | 8 |
| **comm-t/misc-gcp** | docker-worker | 38.0.5 | 2 | 16 |
| **comm-t/t-linux-large-gcp** | docker-worker | unknown version | 21 | 21 |
| **comm-t/t-linux-large-noscratch-gcp** | docker-worker | unknown version | 2141 | 2141 |
| **comm-t/t-linux-xlarge-gcp** | docker-worker | unknown version | 2 | 2 |
| **comm-t/t-linux-xlarge-noscratch-gcp** | docker-worker | unknown version | 284 | 284 |
| **comm-t/t-linux-xlarge-source-gcp** | docker-worker | unknown version | 1380 | 1380 |
| **comm-t/t-linux-xlarge-source-noscratch-gcp** | docker-worker | unknown version | 2 | 2 |
| **gecko-1/b-linux-gcp** | docker-worker | 38.0.5 | 15364 | 15364 |
| **gecko-1/b-linux-gcp-test-bug-1882320** | docker-worker | 38.0.5 | 2 | 2 |
| **gecko-1/b-linux-kvm-gcp** | docker-worker | 38.0.5 | 316 | 316 |
| **gecko-1/b-linux-large-gcp** | docker-worker | 38.0.5 | 70 | 70 |
| **gecko-1/b-linux-medium-gcp** | docker-worker | 38.0.5 | 13906 | 13906 |
| **gecko-1/b-linux-xlarge-gcp** | docker-worker | 38.0.5 | 1667 | 1667 |
| **gecko-1/b-linux-xlarge-gcp-bug1797804-c2** | docker-worker | 38.0.5 | 3 | 3 |
| **gecko-1/decision-gcp** | docker-worker | 38.0.5 | 1851 | 3702 |
| **gecko-1/images-gcp** | docker-worker | 38.0.5 | 15 | 15 |
| **gecko-2/b-linux-gcp** | docker-worker | 38.0.5 | 926 | 926 |
| **gecko-2/b-linux-kvm-gcp** | docker-worker | 38.0.5 | 41 | 41 |
| **gecko-2/b-linux-large-gcp** | docker-worker | 38.0.5 | 24 | 24 |
| **gecko-2/b-linux-xlarge-gcp** | docker-worker | 38.0.5 | 159 | 159 |
| **gecko-2/decision-gcp** | docker-worker | 38.0.5 | 38 | 76 |
| **gecko-2/images-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **gecko-3/b-linux-gcp** | docker-worker | 38.0.5 | 48917 | 48917 |
| **gecko-3/b-linux-kvm-gcp** | docker-worker | 38.0.5 | 652 | 652 |
| **gecko-3/b-linux-large-gcp** | docker-worker | 38.0.5 | 300 | 300 |
| **gecko-3/b-linux-medium-gcp** | docker-worker | 38.0.5 | 12678 | 12678 |
| **gecko-3/b-linux-xlarge-gcp** | docker-worker | 38.0.5 | 4305 | 4305 |
| **gecko-3/decision-gcp** | docker-worker | 38.0.5 | 2033 | 4066 |
| **gecko-3/images-gcp** | docker-worker | 38.0.5 | 38 | 38 |
| **gecko-3/t-linux-xlarge-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **gecko-3/t-linux-xlarge-noscratch-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **gecko-t/misc-gcp** | docker-worker | 38.0.5 | 1258 | 10064 |
| **gecko-t/t-linux-kvm-gcp** | docker-worker | unknown version | 2 | 2 |
| **gecko-t/t-linux-kvm-gcp-bug1862675** | docker-worker | unknown version | 2 | 2 |
| **gecko-t/t-linux-kvm-noscratch-gcp** | docker-worker | unknown version | 67360 | 67360 |
| **gecko-t/t-linux-kvm-noscratch-gcp-bug1862675** | docker-worker | unknown version | 2 | 2 |
| **gecko-t/t-linux-large-gcp** | docker-worker | unknown version | 1656 | 1656 |
| **gecko-t/t-linux-large-hostub2204-gcp** | docker-worker | 48.3.0 | 2 | 2 |
| **gecko-t/t-linux-large-noscratch-gcp** | docker-worker | unknown version | 116833 | 116833 |
| **gecko-t/t-linux-xlarge-gcp** | docker-worker | unknown version | 168 | 168 |
| **gecko-t/t-linux-xlarge-gcp-b1862675** | docker-worker | unknown version | 2 | 2 |
| **gecko-t/t-linux-xlarge-noscratch-gcp** | docker-worker | unknown version | 69442 | 69442 |
| **gecko-t/t-linux-xlarge-ns-gcp-b1862675** | docker-worker | unknown version | 2 | 2 |
| **gecko-t/t-linux-xlarge-source-gcp** | docker-worker | unknown version | 30407 | 30407 |
| **gecko-t/t-linux-xlarge-source-gcp-b1862675** | docker-worker | unknown version | 2 | 2 |
| **gecko-t/t-linux-xlarge-source-noscratch-gcp** | docker-worker | unknown version | 2 | 2 |
| **gecko-t/t-linux-xlarge-source-ns-gcp-b1862675** | docker-worker | unknown version | 2 | 2 |
| **glean-1/b-linux-gcp** | docker-worker | 38.0.5 | 20 | 20 |
| **glean-1/decision-gcp** | docker-worker | 38.0.5 | 1419 | 2838 |
| **glean-1/images-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **glean-3/b-linux-gcp** | docker-worker | 38.0.5 | 18 | 18 |
| **glean-3/decision-gcp** | docker-worker | 38.0.5 | 5 | 10 |
| **glean-3/images-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **infra/build-decision** | docker-worker | 38.0.5 | 3 | 24 |
| **mobile-1/b-linux-gcp** | docker-worker | 38.0.5 | 179 | 179 |
| **mobile-1/decision-gcp** | docker-worker | 38.0.5 | 1266 | 2532 |
| **mobile-1/images-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **mobile-3/b-linux-gcp** | docker-worker | 38.0.5 | 295 | 295 |
| **mobile-3/decision-gcp** | docker-worker | 38.0.5 | 207 | 414 |
| **mobile-3/images-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **mozilla-1/b-linux-gcp** | docker-worker | 38.0.5 | 17 | 17 |
| **mozilla-1/decision-gcp** | docker-worker | 38.0.5 | 1383 | 2766 |
| **mozilla-1/images-gcp** | docker-worker | 38.0.5 | 14 | 14 |
| **mozilla-3/b-linux-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **mozilla-3/decision-gcp** | docker-worker | 38.0.5 | 450 | 900 |
| **mozilla-3/images-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **mozilla-t/t-linux-large-gcp** | docker-worker | unknown version | 46 | 46 |
| **mozilla-t/t-linux-large-noscratch-gcp** | docker-worker | unknown version | 2 | 2 |
| **mozillaonline-1/b-linux-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **mozillaonline-3/b-linux-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **mozillavpn-1/b-linux-gcp** | docker-worker | 38.0.5 | 208 | 208 |
| **mozillavpn-1/b-linux-large-gcp** | docker-worker | 38.0.5 | 144 | 144 |
| **mozillavpn-1/decision-gcp** | docker-worker | 38.0.5 | 1440 | 2880 |
| **mozillavpn-1/images-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **mozillavpn-3/b-linux-gcp** | docker-worker | 38.0.5 | 89 | 89 |
| **mozillavpn-3/b-linux-large-gcp** | docker-worker | 38.0.5 | 87 | 87 |
| **mozillavpn-3/decision-gcp** | docker-worker | 38.0.5 | 15 | 30 |
| **mozillavpn-3/images-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **nss-1/decision-gcp** | docker-worker | 38.0.5 | 13 | 26 |
| **nss-1/images-gcp** | docker-worker | 38.0.5 | 7 | 7 |
| **nss-1/linux-gcp** | docker-worker | 38.0.5 | 508 | 508 |
| **nss-3/decision-gcp** | docker-worker | 38.0.5 | 7 | 14 |
| **nss-3/images-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **nss-3/linux-gcp** | docker-worker | 38.0.5 | 421 | 421 |
| **nss-t/t-linux-xlarge-gcp** | docker-worker | unknown version | 1196 | 1196 |
| **nss-t/t-linux-xlarge-noscratch-gcp** | docker-worker | unknown version | 2 | 2 |
| **releng-1/b-linux-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **releng-1/decision-gcp** | docker-worker | 38.0.5 | 12 | 24 |
| **releng-1/linux-gcp** | docker-worker | 38.0.5 | 17 | 17 |
| **releng-3/b-linux-gcp** | docker-worker | 38.0.5 | 11 | 11 |
| **releng-3/decision-gcp** | docker-worker | 38.0.5 | 9 | 18 |
| **releng-3/linux-gcp** | docker-worker | 38.0.5 | 22 | 22 |
| **releng-t/linux-gcp** | docker-worker | 38.0.5 | 66 | 66 |
| **relops-3/decision-gcp** | docker-worker | 38.0.5 | 2 | 4 |
| **scriptworker-1/b-linux-gcp** | docker-worker | 38.0.5 | 45 | 45 |
| **scriptworker-1/decision-gcp** | docker-worker | 38.0.5 | 64 | 128 |
| **scriptworker-1/images-gcp** | docker-worker | 38.0.5 | 23 | 23 |
| **scriptworker-3/b-linux-gcp** | docker-worker | 38.0.5 | 55 | 55 |
| **scriptworker-3/decision-gcp** | docker-worker | 38.0.5 | 5 | 10 |
| **scriptworker-3/images-gcp** | docker-worker | 38.0.5 | 32 | 32 |
| **taskgraph-1/decision-gcp** | docker-worker | 38.0.5 | 9 | 18 |
| **taskgraph-1/images-gcp** | docker-worker | 38.0.5 | 16 | 16 |
| **taskgraph-3/decision-gcp** | docker-worker | 38.0.5 | 4 | 8 |
| **taskgraph-3/images-gcp** | docker-worker | 38.0.5 | 6 | 6 |
| **taskgraph-t/linux-gcp** | docker-worker | 38.0.5 | 72 | 72 |
| **translations-1/b-linux-large-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **translations-1/b-linux-large-gcp-1tb-32-256** | docker-worker | 38.0.5 | 2 | 2 |
| **translations-1/b-linux-large-gcp-1tb-32-256-standard** | docker-worker | 38.0.5 | 2 | 2 |
| **translations-1/b-linux-large-gcp-1tb-64-512** | docker-worker | 38.0.5 | 2 | 2 |
| **translations-1/b-linux-large-gcp-1tb-64-512-standard** | docker-worker | 38.0.5 | 2 | 2 |
| **translations-1/b-linux-large-gcp-300gb** | docker-worker | 38.0.5 | 2 | 2 |
| **translations-1/decision-gcp** | docker-worker | 38.0.5 | 3077 | 6154 |
| **translations-1/images-gcp** | docker-worker | 38.0.5 | 5 | 5 |
| **xpi-1/b-linux-gcp** | docker-worker | 38.0.5 | 47 | 47 |
| **xpi-1/decision-gcp** | docker-worker | 38.0.5 | 18 | 36 |
| **xpi-1/images-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **xpi-3/b-linux-gcp** | docker-worker | 38.0.5 | 3 | 3 |
| **xpi-3/decision-gcp** | docker-worker | 38.0.5 | 5 | 10 |
| **xpi-3/images-gcp** | docker-worker | 38.0.5 | 2 | 2 |


## Script Worker

Total: `37`



| Worker Pool | Implementation | Version | Total Workers | Total Capacity |
| --- | --- | --- | ---: | ---: |
| **scriptworker-k8s/app-services-3-beetmover** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/app-services-3-signing** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/comm-1-pushmsix** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/comm-3-balrog** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/comm-3-beetmover** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/comm-3-bouncer** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/comm-3-pushmsix** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/comm-3-shipit** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/comm-3-signing** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/comm-3-tree** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/comm-t-signing** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/gecko-1-pushapk** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/gecko-3-balrog** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/gecko-3-beetmover** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/gecko-3-bouncer** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/gecko-3-pushapk** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/gecko-3-shipit** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/gecko-3-signing** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/gecko-3-tree** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/gecko-t-signing** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/mobile-3-bitrise** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/mobile-3-pushapk** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/mobile-3-signing** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/mobile-t-signing** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/mozillavpn-3-signing** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/xpi-t-signing** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-prov-v1/adhoc-3-signing-mac** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-prov-v1/adhoc-depsigning-mac14m2** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-prov-v1/adhoc-t-signing-mac** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-prov-v1/depsigning-mac-v1** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-prov-v1/gecko-depsigning-mac14m2** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-prov-v1/mozillavpn-3-signing-mac** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-prov-v1/signing-mac-v1** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-prov-v1/tb-depsigning-mac-v1** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-prov-v1/tb-depsigning-mac14m2** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
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
| **gecko-t/t-linux-vm-2204-wayland** |  | No artifacts found | 1506 | 1506 |
| **gecko-t/t-linux-vm-2204-wayland-snap** |  | No artifacts found | 2 | 2 |
| **scriptworker-prov-v1/mac-notarization-poller** |  | No artifacts found | 0 | 0 |
| **scriptworker-prov-v1/tb-mac-notarization-poller** |  | No artifacts found | 0 | 0 |


## Version not determined [^2]

Total: `14`


Count by image:

| Version | Count |
| :--- | ---: |
| /subscriptions/108d46d5-fe9b-4850-9a7d-8c914aa6c1f0/resourceGroups/rg-packer-worker-images/providers/Microsoft.Compute/galleries/win10_64_2009/images/win10_64_2009/versions/1.0.2 | 1 |
| projects/taskcluster-imaging/global/images/gw-fxci-gcp-l1-2404-amd64-googlecompute-alpha | 1 |
| /subscriptions/108d46d5-fe9b-4850-9a7d-8c914aa6c1f0/resourceGroups/rg-packer-worker-images/providers/Microsoft.Compute/galleries/win2022_64_2009_alpha/images/win2022_64_2009_alpha/versions/1.0.0 | 3 |
| /subscriptions/108d46d5-fe9b-4850-9a7d-8c914aa6c1f0/resourceGroups/rg-packer-through-cib/providers/Microsoft.Compute/images/win2012r2-64-l1-centralus-2012-R2-Datacenter-75bd9ed,/subscriptions/108d46d5-fe9b-4850-9a7d-8c914aa6c1f0/resourceGroups/rg-packer-through-cib/providers/Microsoft.Compute/images/win2012r2-64-l1-eastus2-2012-R2-Datacenter-75bd9ed,/subscriptions/108d46d5-fe9b-4850-9a7d-8c914aa6c1f0/resourceGroups/rg-packer-through-cib/providers/Microsoft.Compute/images/win2012r2-64-l1-northcentralus-2012-R2-Datacenter-75bd9ed,/subscriptions/108d46d5-fe9b-4850-9a7d-8c914aa6c1f0/resourceGroups/rg-packer-through-cib/providers/Microsoft.Compute/images/win2012r2-64-l1-westus2-2012-R2-Datacenter-75bd9ed | 1 |
| /subscriptions/108d46d5-fe9b-4850-9a7d-8c914aa6c1f0/resourceGroups/rg-packer-worker-images/providers/Microsoft.Compute/galleries/win11_a64_24h2_tester_alpha/images/win11_a64_24h2_tester_alpha/versions/1.0.0 | 2 |
|  | 3 |
| /subscriptions/108d46d5-fe9b-4850-9a7d-8c914aa6c1f0/resourceGroups/rg-packer-worker-images/providers/Microsoft.Compute/galleries/win2022_64_2009/images/win2022_64_2009/versions/1.0.2 | 1 |
| /subscriptions/a30e97ab-734a-4f3b-a0e4-c51c0bff0701/resourceGroups/rg-packer-worker-images/providers/Microsoft.Compute/galleries/trusted_win2022_64_2009/images/trusted_win2022_64_2009/versions/1.0.2 | 2 |


| Worker Pool | Implementation | Version | Total Workers | Total Capacity |
| --- | --- | --- | ---: | ---: |
| **comm-t/win11-a64-24h2-alpha** |  | Version not determined; task not (yet) claimed | 7 | 7 |
| **gecko-1/b-win2012-azure** |  | Version not determined; task not (yet) claimed | 2 | 2 |
| **gecko-1/b-win2022-alpha** |  | Version not determined; task not (yet) claimed | 312 | 312 |
| **gecko-t/t-linux-2404-relsre** |  | Version not determined; task not (yet) claimed | 3 | 3 |
| **gecko-t/win10-64-2009** |  | Version not determined; task not (yet) claimed | 4019 | 4019 |
| **gecko-t/win11-a64-24h2-alpha** |  | Version not determined; task not (yet) claimed | 8 | 8 |
| **mozilla-3/b-win2022** |  | Version not determined; task not (yet) claimed | 382 | 382 |
| **mozillavpn-1/b-win2022-alpha** |  | Version not determined; task not (yet) claimed | 7 | 7 |
| **nss-1/b-win2022-alpha** |  | Version not determined; task not (yet) claimed | 7 | 7 |
| **proj-autophone/gecko-t-lambda-alpha-a55** |  | Version not determined; task not (yet) claimed | 0 | 0 |
| **releng-hardware/gecko-t-osx-1015-r8** |  | Version not determined; task not (yet) claimed | 0 | 0 |
| **releng-hardware/gecko-t-osx-1400-r8** |  | Version not determined; task not (yet) claimed | 0 | 0 |
| **relops-1/b-win2022** |  | Version not determined; task not (yet) claimed | 222 | 222 |
| **relops-3/b-win2022** |  | Version not determined; task not (yet) claimed | 228 | 228 |



[^1]: Those are the pools whose tasks were claimed and resolved by a worker as expected, but the worker did not publish either artifact `public/logs/live_backing.log` nor `public/logs/chain_of_trust.log`, which is the source used to identify the worker implementation.

[^2]: Probing task remains pending after two hours. Those are the pools that were not able to start any worker to claim the task within two hours.

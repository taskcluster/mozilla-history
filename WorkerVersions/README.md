

# Worker Pool Versions


## Generic Worker

Total: `350`

Count by version:

| Version | Count |
| :--- | ---: |
| 36.0.0 | 4 |
| 45.0.0 | 1 |
| 60.3.4 | 19 |
| 61.0.0 | 1 |
| 64.3.0 | 26 |
| 65.1.0 | 1 |
| 84.1.2 | 7 |
| 86.0.2 | 3 |
| 87.0.0 | 1 |
| 88.0.2 | 2 |
| 88.1.3 | 48 |
| 91.0.2 | 2 |
| 91.1.0 | 4 |
| 91.1.1 | 2 |
| 93.1.2 | 208 |
| 93.1.4 | 16 |
| 94.0.1 | 5 |


Count by image:

| Version | Count |
| :--- | ---: |
| projects/taskcluster-imaging/global/images/gw-fxci-gcp-l1-arm64-gui-googlecompute-2024-09-18t14-57-52z | 9 |
| projects/taskcluster-imaging/global/images/gw-fxci-gcp-l1-2404-amd64-gui-googlecompute-alpha | 1 |
| projects/taskcluster-imaging/global/images/gw-fxci-gcp-l1-2404-amd64-headless-googlecompute-2025-11-14 | 117 |
| projects/taskcluster-imaging/global/images/gw-fxci-gcp-l1-gui-googlecompute-2024-08-22t22-48-09z | 8 |
|  | 39 |
| projects/taskcluster-imaging/global/images/gw-fxci-gcp-l1-gui-googlecompute-2025-01-13t22-33-40z | 2 |
| projects/taskcluster-imaging/global/images/gw-fxci-gcp-l1-2404-amd64-headless-googlecompute-alpha | 6 |
| projects/taskcluster-imaging/global/images/gw-fxci-gcp-l1-2404-arm64-headless-googlecompute-alpha | 1 |
| unknown | 54 |
| /subscriptions/108d46d5-fe9b-4850-9a7d-8c914aa6c1f0/resourceGroups/rg-packer-worker-images/providers/Microsoft.Compute/galleries/win11_a64_24h2_tester_alpha/images/win11_a64_24h2_tester_alpha/versions/1.0.0 | 4 |
| projects/fxci-production-level3-workers/global/images/gw-fxci-gcp-l3-2404-arm64-headless-googlecompute-2025-11-14 | 6 |
| projects/fxci-production-level3-workers/global/images/gw-fxci-gcp-l3-2404-amd64-headless-googlecompute-alpha | 1 |
| projects/taskcluster-imaging/global/images/gw-fxci-gcp-l1-2404-amd64-gui-googlecompute-2025-11-14 | 4 |
| projects/taskcluster-imaging/global/images/gw-fxci-gcp-l1-2404-amd64-headless-googlecompute-alpha-tc | 7 |
| projects/fxci-production-level3-workers/global/images/gw-fxci-gcp-l3-2404-amd64-headless-googlecompute-2025-11-14 | 58 |
| /subscriptions/108d46d5-fe9b-4850-9a7d-8c914aa6c1f0/resourceGroups/rg-packer-worker-images/providers/Microsoft.Compute/galleries/win11_a64_24h2_builder_alpha/images/win11_a64_24h2_builder_alpha/versions/1.0.0 | 2 |
| /subscriptions/108d46d5-fe9b-4850-9a7d-8c914aa6c1f0/resourceGroups/rg-packer-worker-images/providers/Microsoft.Compute/galleries/win11_a64_24h2_builder/images/win11_a64_24h2_builder/versions/1.0.8 | 3 |
| projects/taskcluster-imaging/global/images/gw-fxci-gcp-l1-2404-arm64-headless-googlecompute-2025-11-14 | 14 |
| projects/fxci-production-level3-workers/global/images/gw-fxci-gcp-l3-gui-googlecompute-2024-09-18t05-46-31z | 3 |
| projects/fxci-production-level3-workers/global/images/gw-fxci-gcp-l3-arm64-gui-googlecompute-2024-09-18t19-02-58z | 4 |
| /subscriptions/108d46d5-fe9b-4850-9a7d-8c914aa6c1f0/resourceGroups/rg-packer-worker-images/providers/Microsoft.Compute/galleries/win11_a64_24h2_tester/images/win11_a64_24h2_tester/versions/1.0.8 | 4 |
| /subscriptions/a30e97ab-734a-4f3b-a0e4-c51c0bff0701/resourceGroups/rg-packer-worker-images/providers/Microsoft.Compute/galleries/trusted_win11_a64_24h2_builder/images/trusted_win11_a64_24h2_builder/versions/1.0.8 | 3 |


| Worker Pool | Implementation | Version | Engine | Revision | OS | Arch | GO | Total Workers | Total Capacity |
| --- | --- | --- | --- | --- | --- | --- | --- | ---: | ---: |
| **adhoc-1/decision** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 3 | 3 |
| **adhoc-1/images** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 5 | 5 |
| **adhoc-3/decision** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 2 | 2 |
| **adhoc-3/images** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 2 | 2 |
| **app-services-1/b-linux** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 6 | 6 |
| **app-services-1/b-linux-docker-amd** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 2 | 2 |
| **app-services-1/decision** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 1172 | 1172 |
| **app-services-1/images** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 2 | 2 |
| **app-services-3/b-linux** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 76 | 76 |
| **app-services-3/b-linux-docker-amd** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 2 | 2 |
| **app-services-3/decision** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 5 | 5 |
| **app-services-3/images** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 2 | 2 |
| **code-analysis-1/linux-docker** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 2 | 2 |
| **code-analysis-1/linux-gw-gcp** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | amd64 | 1.22.2 | 23 | 23 |
| **code-analysis-3/linux-docker** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 2 | 2 |
| **code-analysis-3/linux-gw-gcp** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | amd64 | 1.22.2 | 2 | 2 |
| **code-coverage/bot** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 230 | 230 |
| **code-review/bot** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 7631 | 7631 |
| **comm-1/b-linux** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 3 | 3 |
| **comm-1/b-linux-aarch64** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | arm64 | 1.25.4 | 2 | 2 |
| **comm-1/b-linux-docker-amd** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 15 | 15 |
| **comm-1/b-linux-docker-large-amd** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 2 | 2 |
| **comm-1/b-linux-docker-xlarge-amd** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 3 | 3 |
| **comm-1/b-linux-gcp-aarch64** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | arm64 | 1.22.2 | 2 | 2 |
| **comm-1/b-linux-large** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 2 | 2 |
| **comm-1/b-linux-xlarge** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 2 | 2 |
| **comm-1/b-win2022** | generic-worker | 88.1.3 | multiuser | ffc99d77c0 | windows | amd64 | 1.25.1 | 2 | 2 |
| **comm-1/decision** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 4 | 4 |
| **comm-1/images** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 2 | 2 |
| **comm-1/images-aarch64** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | arm64 | 1.25.4 | 2 | 2 |
| **comm-1/images-gcp-aarch64** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | arm64 | 1.22.2 | 2 | 2 |
| **comm-2/b-linux** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 2 | 2 |
| **comm-2/b-linux-aarch64** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | arm64 | 1.25.4 | 2 | 2 |
| **comm-2/b-linux-docker-amd** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 2 | 2 |
| **comm-2/b-linux-docker-large-amd** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 2 | 2 |
| **comm-2/b-linux-docker-xlarge-amd** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 2 | 2 |
| **comm-2/b-linux-gcp-aarch64** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | arm64 | 1.22.2 | 2 | 2 |
| **comm-2/b-linux-large** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 2 | 2 |
| **comm-2/b-linux-xlarge** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 2 | 2 |
| **comm-2/decision** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 2 | 2 |
| **comm-2/images** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 2 | 2 |
| **comm-2/images-aarch64** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | arm64 | 1.25.4 | 2 | 2 |
| **comm-2/images-gcp-aarch64** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | arm64 | 1.22.2 | 2 | 2 |
| **comm-3/b-linux** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 2 | 2 |
| **comm-3/b-linux-aarch64** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | arm64 | 1.25.4 | 2 | 2 |
| **comm-3/b-linux-docker-amd** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 1286 | 1286 |
| **comm-3/b-linux-docker-large-amd** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 8 | 8 |
| **comm-3/b-linux-docker-xlarge-amd** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 10 | 10 |
| **comm-3/b-linux-gcp-aarch64** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | arm64 | 1.22.2 | 2 | 2 |
| **comm-3/b-linux-large** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 2 | 2 |
| **comm-3/b-linux-xlarge** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 2 | 2 |
| **comm-3/b-win2022** | generic-worker | 88.1.3 | multiuser | ffc99d77c0 | windows | amd64 | 1.25.1 | 433 | 433 |
| **comm-3/decision** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 2 | 2 |
| **comm-3/images** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 2 | 2 |
| **comm-3/images-aarch64** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | arm64 | 1.25.4 | 2 | 2 |
| **comm-3/images-gcp-aarch64** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | arm64 | 1.22.2 | 2 | 2 |
| **comm-t/misc** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 2 | 2 |
| **comm-t/t-linux-arm64-docker** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | arm64 | 1.25.4 | 2 | 2 |
| **comm-t/t-linux-docker** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 2 | 2 |
| **comm-t/t-linux-docker-amd** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 670 | 670 |
| **comm-t/t-linux-docker-noscratch** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 2 | 2 |
| **comm-t/t-linux-docker-noscratch-amd** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 2747 | 2747 |
| **comm-t/win11-64-2009** | generic-worker | 88.1.3 | multiuser | ffc99d77c0 | windows | amd64 | 1.25.1 | 2 | 2 |
| **comm-t/win11-64-2009-ssd** | generic-worker | 88.1.3 | multiuser | ffc99d77c0 | windows | amd64 | 1.25.1 | 2 | 2 |
| **comm-t/win11-64-24h2** | generic-worker | 88.1.3 | multiuser | ffc99d77c0 | windows | amd64 | 1.25.1 | 2801 | 2801 |
| **comm-t/win11-64-24h2-ssd** | generic-worker | 88.1.3 | multiuser | ffc99d77c0 | windows | amd64 | 1.25.1 | 340 | 340 |
| **comm-t/win11-a64-24h2** | generic-worker | 88.1.3 | multiuser | ffc99d77c0 | windows | amd64 | 1.25.1 | 2 | 2 |
| **comm-t/win11-a64-24h2-alpha** | generic-worker | 93.1.4 | multiuser | 1088d85c3b | windows | arm64 | 1.25.4 | 2 | 2 |
| **enterprise-1/b-linux** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 394 | 394 |
| **enterprise-1/b-linux-aarch64** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | arm64 | 1.25.4 | 2 | 2 |
| **enterprise-1/b-linux-docker-amd** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 236 | 236 |
| **enterprise-1/b-linux-docker-large-amd** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 4 | 4 |
| **enterprise-1/b-linux-docker-xlarge-amd** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 222 | 222 |
| **enterprise-1/b-win2022** | generic-worker | 88.1.3 | multiuser | ffc99d77c0 | windows | amd64 | 1.25.1 | 119 | 119 |
| **enterprise-1/decision** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 1160 | 1160 |
| **enterprise-1/images** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 3 | 3 |
| **enterprise-1/images-aarch64** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | arm64 | 1.25.4 | 2 | 2 |
| **enterprise-3/b-linux** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 217 | 217 |
| **enterprise-3/b-linux-aarch64** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | arm64 | 1.25.4 | 2 | 2 |
| **enterprise-3/b-linux-docker-amd** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 158 | 158 |
| **enterprise-3/b-linux-docker-large-amd** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 5 | 5 |
| **enterprise-3/b-linux-docker-xlarge-amd** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 120 | 120 |
| **enterprise-3/b-win2022** | generic-worker | 88.1.3 | multiuser | ffc99d77c0 | windows | amd64 | 1.25.1 | 59 | 59 |
| **enterprise-3/decision** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 13 | 13 |
| **enterprise-3/images** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 3 | 3 |
| **enterprise-3/images-aarch64** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | arm64 | 1.25.4 | 2 | 2 |
| **enterprise-t/misc** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 182 | 182 |
| **enterprise-t/t-linux-docker-amd** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 1244 | 1244 |
| **enterprise-t/win11-64-24h2-source** | generic-worker | 88.1.3 | multiuser | ffc99d77c0 | windows | amd64 | 1.25.1 | 146 | 146 |
| **gecko-1/b-linux** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 4288 | 4288 |
| **gecko-1/b-linux-2204-kvm-gcp** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | amd64 | 1.22.2 | 2 | 2 |
| **gecko-1/b-linux-aarch64** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | arm64 | 1.25.4 | 8 | 8 |
| **gecko-1/b-linux-docker-alpha** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 2 | 2 |
| **gecko-1/b-linux-docker-amd** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 6204 | 6204 |
| **gecko-1/b-linux-docker-large-amd** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 80 | 80 |
| **gecko-1/b-linux-docker-xlarge-amd** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 944 | 944 |
| **gecko-1/b-linux-gcp-aarch64** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | arm64 | 1.22.2 | 2 | 2 |
| **gecko-1/b-linux-kvm** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 226 | 226 |
| **gecko-1/b-linux-large** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 2 | 2 |
| **gecko-1/b-linux-medium** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 10795 | 10795 |
| **gecko-1/b-linux-xlarge** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 5 | 5 |
| **gecko-1/b-win2022** | generic-worker | 88.1.3 | multiuser | ffc99d77c0 | windows | amd64 | 1.25.1 | 686 | 686 |
| **gecko-1/b-win2022-alpha** | generic-worker | 93.1.4 | multiuser | 1088d85c3b | windows | amd64 | 1.25.4 | 15 | 15 |
| **gecko-1/b-win2022-xxlarge** | generic-worker | 88.1.3 | multiuser | ffc99d77c0 | windows | amd64 | 1.25.1 | 80 | 80 |
| **gecko-1/decision** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 2265 | 2265 |
| **gecko-1/images** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 57 | 57 |
| **gecko-1/images-aarch64** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | arm64 | 1.25.4 | 7 | 7 |
| **gecko-1/images-gcp-aarch64** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | arm64 | 1.22.2 | 2 | 2 |
| **gecko-1/win11-a64-24h2-builder** | generic-worker | 88.1.3 | multiuser | ffc99d77c0 | windows | amd64 | 1.25.1 | 43 | 43 |
| **gecko-1/win11-a64-24h2-builder-alpha** | generic-worker | 93.1.4 | multiuser | 1088d85c3b | windows | arm64 | 1.25.4 | 2 | 2 |
| **gecko-2/b-linux** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 30 | 30 |
| **gecko-2/b-linux-aarch64** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | arm64 | 1.25.4 | 2 | 2 |
| **gecko-2/b-linux-docker-amd** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 107 | 107 |
| **gecko-2/b-linux-docker-large-amd** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 6 | 6 |
| **gecko-2/b-linux-docker-xlarge-amd** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 21 | 21 |
| **gecko-2/b-linux-gcp-aarch64** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | arm64 | 1.22.2 | 2 | 2 |
| **gecko-2/b-linux-kvm** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 7 | 7 |
| **gecko-2/b-linux-large** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 2 | 2 |
| **gecko-2/b-linux-medium** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 134 | 134 |
| **gecko-2/b-linux-xlarge** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 2 | 2 |
| **gecko-2/b-win2022** | generic-worker | 88.1.3 | multiuser | ffc99d77c0 | windows | amd64 | 1.25.1 | 2 | 2 |
| **gecko-2/decision** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 10 | 10 |
| **gecko-2/images** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 2 | 2 |
| **gecko-2/images-aarch64** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | arm64 | 1.25.4 | 2 | 2 |
| **gecko-2/images-gcp-aarch64** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | arm64 | 1.22.2 | 2 | 2 |
| **gecko-3/b-linux** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 12674 | 12674 |
| **gecko-3/b-linux-2204-kvm-gcp** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | amd64 | 1.22.2 | 2 | 2 |
| **gecko-3/b-linux-aarch64** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | arm64 | 1.25.4 | 4 | 4 |
| **gecko-3/b-linux-docker-alpha** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 3 | 3 |
| **gecko-3/b-linux-docker-amd** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 14439 | 14439 |
| **gecko-3/b-linux-docker-large-amd** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 152 | 152 |
| **gecko-3/b-linux-docker-xlarge-amd** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 2987 | 2987 |
| **gecko-3/b-linux-gcp-aarch64** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | arm64 | 1.22.2 | 2 | 2 |
| **gecko-3/b-linux-kvm** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 583 | 583 |
| **gecko-3/b-linux-large** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 2 | 2 |
| **gecko-3/b-linux-medium** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 16717 | 16717 |
| **gecko-3/b-linux-xlarge** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 20 | 20 |
| **gecko-3/b-win2022** | generic-worker | 88.1.3 | multiuser | ffc99d77c0 | windows | amd64 | 1.25.1 | 5164 | 5164 |
| **gecko-3/b-win2022-xxlarge** | generic-worker | 88.1.3 | multiuser | ffc99d77c0 | windows | amd64 | 1.25.1 | 12 | 12 |
| **gecko-3/decision** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 4273 | 4273 |
| **gecko-3/images** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 2 | 2 |
| **gecko-3/images-aarch64** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | arm64 | 1.25.4 | 2 | 2 |
| **gecko-3/images-gcp-aarch64** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | arm64 | 1.22.2 | 2 | 2 |
| **gecko-3/win11-a64-24h2-builder** | generic-worker | 88.1.3 | multiuser | ffc99d77c0 | windows | amd64 | 1.25.1 | 53 | 53 |
| **gecko-t/misc** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 525 | 525 |
| **gecko-t/t-linux-2204-wayland** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | amd64 | 1.22.2 | 7862 | 7862 |
| **gecko-t/t-linux-2204-wayland-arm64-relsre** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | arm64 | 1.22.2 | 2 | 2 |
| **gecko-t/t-linux-2204-wayland-relsre** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | amd64 | 1.22.2 | 2 | 2 |
| **gecko-t/t-linux-2204-wayland-root-exp** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | amd64 | 1.22.2 | 2 | 2 |
| **gecko-t/t-linux-2204-wayland-snap** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | amd64 | 1.22.2 | 139 | 139 |
| **gecko-t/t-linux-2404-headless-alpha** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 2 | 2 |
| **gecko-t/t-linux-2404-headless-arm64-alpha** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | arm64 | 1.25.4 | 2 | 2 |
| **gecko-t/t-linux-2404-headless-ssd-alpha** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 2 | 2 |
| **gecko-t/t-linux-2404-wayland** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 2 | 2 |
| **gecko-t/t-linux-2404-wayland-relsre** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 2 | 2 |
| **gecko-t/t-linux-2404-wayland-snap** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 139 | 139 |
| **gecko-t/t-linux-arm64-docker** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | arm64 | 1.25.4 | 21 | 21 |
| **gecko-t/t-linux-docker** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 1337 | 1337 |
| **gecko-t/t-linux-docker-16c32gb-amd** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 2859 | 2859 |
| **gecko-t/t-linux-docker-amd** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 20916 | 20916 |
| **gecko-t/t-linux-docker-amd-alpha** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 2 | 2 |
| **gecko-t/t-linux-docker-kvm** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 62652 | 62652 |
| **gecko-t/t-linux-docker-kvm-alpha** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 2 | 2 |
| **gecko-t/t-linux-docker-noscratch** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 5462 | 5462 |
| **gecko-t/t-linux-docker-noscratch-amd** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 82843 | 82843 |
| **gecko-t/t-linux-docker-noscratch-amd-alpha** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 2 | 2 |
| **gecko-t/t-linux-xlarge-2204-wayland** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | amd64 | 1.22.2 | 9754 | 9754 |
| **gecko-t/t-linux-xlarge-2404-wayland** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 2 | 2 |
| **gecko-t/win10-64-2009** | generic-worker | 88.1.3 | multiuser | ffc99d77c0 | windows | amd64 | 1.25.1 | 3978 | 3978 |
| **gecko-t/win10-64-2009-alpha** | generic-worker | 91.1.0 | multiuser | b8f34332c3 | windows | amd64 | 1.25.3 | 2 | 2 |
| **gecko-t/win10-64-2009-gpu** | generic-worker | 88.1.3 | multiuser | ffc99d77c0 | windows | amd64 | 1.25.1 | 445 | 445 |
| **gecko-t/win10-64-2009-gpu-alpha** | generic-worker | 91.1.0 | multiuser | b8f34332c3 | windows | amd64 | 1.25.3 | 2 | 2 |
| **gecko-t/win10-64-2009-source** | generic-worker | 88.1.3 | multiuser | ffc99d77c0 | windows | amd64 | 1.25.1 | 34 | 34 |
| **gecko-t/win10-64-2009-source-alpha** | generic-worker | 91.1.0 | multiuser | b8f34332c3 | windows | amd64 | 1.25.3 | 2 | 2 |
| **gecko-t/win10-64-2009-webgpu** | generic-worker | 88.1.3 | multiuser | ffc99d77c0 | windows | amd64 | 1.25.1 | 2 | 2 |
| **gecko-t/win10-64-2009-webgpu-alpha** | generic-worker | 91.1.0 | multiuser | b8f34332c3 | windows | amd64 | 1.25.3 | 2 | 2 |
| **gecko-t/win11-64-2009** | generic-worker | 88.1.3 | multiuser | ffc99d77c0 | windows | amd64 | 1.25.1 | 677 | 677 |
| **gecko-t/win11-64-2009-alpha** | generic-worker | 93.1.4 | multiuser | 1088d85c3b | windows | amd64 | 1.25.4 | 2 | 2 |
| **gecko-t/win11-64-2009-gpu** | generic-worker | 88.1.3 | multiuser | ffc99d77c0 | windows | amd64 | 1.25.1 | 404 | 404 |
| **gecko-t/win11-64-2009-gpu-alpha** | generic-worker | 93.1.4 | multiuser | 1088d85c3b | windows | amd64 | 1.25.4 | 2 | 2 |
| **gecko-t/win11-64-2009-source** | generic-worker | 88.1.3 | multiuser | ffc99d77c0 | windows | amd64 | 1.25.1 | 2 | 2 |
| **gecko-t/win11-64-2009-source-alpha** | generic-worker | 93.1.4 | multiuser | 1088d85c3b | windows | amd64 | 1.25.4 | 2 | 2 |
| **gecko-t/win11-64-2009-ssd** | generic-worker | 88.1.3 | multiuser | ffc99d77c0 | windows | amd64 | 1.25.1 | 2 | 2 |
| **gecko-t/win11-64-2009-ssd-gpu** | generic-worker | 88.1.3 | multiuser | ffc99d77c0 | windows | amd64 | 1.25.1 | 2 | 2 |
| **gecko-t/win11-64-2009-webgpu** | generic-worker | 88.1.3 | multiuser | ffc99d77c0 | windows | amd64 | 1.25.1 | 2 | 2 |
| **gecko-t/win11-64-2009-webgpu-alpha** | generic-worker | 93.1.4 | multiuser | 1088d85c3b | windows | amd64 | 1.25.4 | 3 | 3 |
| **gecko-t/win11-64-24h2** | generic-worker | 88.1.3 | multiuser | ffc99d77c0 | windows | amd64 | 1.25.1 | 43651 | 43651 |
| **gecko-t/win11-64-24h2-alpha** | generic-worker | 94.0.1 | multiuser | e38b5ecf97 | windows | amd64 | 1.25.4 | 2 | 2 |
| **gecko-t/win11-64-24h2-gpu** | generic-worker | 88.1.3 | multiuser | ffc99d77c0 | windows | amd64 | 1.25.1 | 17619 | 17619 |
| **gecko-t/win11-64-24h2-gpu-alpha** | generic-worker | 94.0.1 | multiuser | e38b5ecf97 | windows | amd64 | 1.25.4 | 2 | 2 |
| **gecko-t/win11-64-24h2-large** | generic-worker | 88.1.3 | multiuser | ffc99d77c0 | windows | amd64 | 1.25.1 | 1656 | 1656 |
| **gecko-t/win11-64-24h2-large-alpha** | generic-worker | 94.0.1 | multiuser | e38b5ecf97 | windows | amd64 | 1.25.4 | 2 | 2 |
| **gecko-t/win11-64-24h2-source** | generic-worker | 88.1.3 | multiuser | ffc99d77c0 | windows | amd64 | 1.25.1 | 1520 | 1520 |
| **gecko-t/win11-64-24h2-source-alpha** | generic-worker | 94.0.1 | multiuser | e38b5ecf97 | windows | amd64 | 1.25.4 | 2 | 2 |
| **gecko-t/win11-64-24h2-ssd** | generic-worker | 88.1.3 | multiuser | ffc99d77c0 | windows | amd64 | 1.25.1 | 2 | 2 |
| **gecko-t/win11-64-24h2-ssd-gpu** | generic-worker | 88.1.3 | multiuser | ffc99d77c0 | windows | amd64 | 1.25.1 | 2 | 2 |
| **gecko-t/win11-64-24h2-webgpu** | generic-worker | 88.1.3 | multiuser | ffc99d77c0 | windows | amd64 | 1.25.1 | 3108 | 3108 |
| **gecko-t/win11-64-24h2-webgpu-alpha** | generic-worker | 94.0.1 | multiuser | e38b5ecf97 | windows | amd64 | 1.25.4 | 2 | 2 |
| **gecko-t/win11-a64-24h2** | generic-worker | 88.1.3 | multiuser | ffc99d77c0 | windows | amd64 | 1.25.1 | 168 | 168 |
| **gecko-t/win11-a64-24h2-alpha** | generic-worker | 93.1.4 | multiuser | 1088d85c3b | windows | arm64 | 1.25.4 | 2 | 2 |
| **glean-1/b-linux** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 4 | 4 |
| **glean-1/decision** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 1156 | 1156 |
| **glean-1/images** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 2 | 2 |
| **glean-3/b-linux** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 6 | 6 |
| **glean-3/decision** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 3 | 3 |
| **glean-3/images** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 2 | 2 |
| **mobile-1/b-linux** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 213 | 213 |
| **mobile-1/decision** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 1148 | 1148 |
| **mobile-1/images** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 2 | 2 |
| **mobile-3/b-linux** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 232 | 232 |
| **mobile-3/decision** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 64 | 64 |
| **mobile-3/images** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 2 | 2 |
| **mozilla-1/b-linux** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 2 | 2 |
| **mozilla-1/b-win2022** | generic-worker | 88.1.3 | multiuser | ffc99d77c0 | windows | amd64 | 1.25.1 | 2 | 2 |
| **mozilla-1/decision** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 1159 | 1159 |
| **mozilla-1/images** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 22 | 22 |
| **mozilla-3/b-linux** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 2 | 2 |
| **mozilla-3/b-win2022** | generic-worker | 88.1.3 | multiuser | ffc99d77c0 | windows | amd64 | 1.25.1 | 3 | 3 |
| **mozilla-3/decision** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 2 | 2 |
| **mozilla-3/images** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 2 | 2 |
| **mozilla-t/t-linux-2204-wayland** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | amd64 | 1.22.2 | 500 | 500 |
| **mozilla-t/t-linux-2204-wayland-relsre** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | amd64 | 1.22.2 | 2 | 2 |
| **mozilla-t/t-linux-2404-wayland** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 2 | 2 |
| **mozilla-t/t-linux-arm64-docker** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | arm64 | 1.25.4 | 2 | 2 |
| **mozilla-t/t-linux-docker** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 34 | 34 |
| **mozilla-t/t-linux-docker-amd** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 3 | 3 |
| **mozilla-t/t-linux-docker-noscratch** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 23 | 23 |
| **mozilla-t/t-linux-docker-noscratch-amd** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 2 | 2 |
| **mozillavpn-1/b-linux** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 173 | 173 |
| **mozillavpn-1/b-linux-gcp-gw** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | amd64 | 1.22.2 | 2 | 2 |
| **mozillavpn-1/b-linux-large** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 118 | 118 |
| **mozillavpn-1/b-win2022** | generic-worker | 88.1.3 | multiuser | ffc99d77c0 | windows | amd64 | 1.25.1 | 51 | 51 |
| **mozillavpn-1/b-win2022-alpha** | generic-worker | 93.1.4 | multiuser | 1088d85c3b | windows | amd64 | 1.25.4 | 2 | 2 |
| **mozillavpn-1/decision** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 1147 | 1147 |
| **mozillavpn-1/images** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 3 | 3 |
| **mozillavpn-1/win11-a64-24h2-builder** | generic-worker | 88.1.3 | multiuser | ffc99d77c0 | windows | amd64 | 1.25.1 | 2 | 2 |
| **mozillavpn-3/b-linux** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 64 | 64 |
| **mozillavpn-3/b-linux-gcp-gw** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | amd64 | 1.22.2 | 2 | 2 |
| **mozillavpn-3/b-linux-large** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 57 | 57 |
| **mozillavpn-3/b-win2022** | generic-worker | 88.1.3 | multiuser | ffc99d77c0 | windows | amd64 | 1.25.1 | 14 | 14 |
| **mozillavpn-3/decision** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 8 | 8 |
| **mozillavpn-3/images** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 2 | 2 |
| **mozillavpn-3/win11-a64-24h2-builder** | generic-worker | 88.1.3 | multiuser | ffc99d77c0 | windows | amd64 | 1.25.1 | 2 | 2 |
| **nss-1/b-win2022** | generic-worker | 88.1.3 | multiuser | ffc99d77c0 | windows | amd64 | 1.25.1 | 30 | 30 |
| **nss-1/b-win2022-alpha** | generic-worker | 93.1.4 | multiuser | 1088d85c3b | windows | amd64 | 1.25.4 | 2 | 2 |
| **nss-1/decision** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 2 | 2 |
| **nss-1/images** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 2 | 2 |
| **nss-1/linux-docker** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 2 | 2 |
| **nss-1/win11-a64-24h2** | generic-worker | 88.1.3 | multiuser | ffc99d77c0 | windows | amd64 | 1.25.1 | 2 | 2 |
| **nss-1/win11-a64-24h2-alpha** | generic-worker | 93.1.4 | multiuser | 1088d85c3b | windows | arm64 | 1.25.4 | 3 | 3 |
| **nss-1/win11-a64-24h2-builder** | generic-worker | 88.1.3 | multiuser | ffc99d77c0 | windows | amd64 | 1.25.1 | 2 | 2 |
| **nss-1/win11-a64-24h2-builder-alpha** | generic-worker | 93.1.4 | multiuser | 1088d85c3b | windows | arm64 | 1.25.4 | 2 | 2 |
| **nss-3/b-win2022** | generic-worker | 88.1.3 | multiuser | ffc99d77c0 | windows | amd64 | 1.25.1 | 37 | 37 |
| **nss-3/decision** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 3 | 3 |
| **nss-3/images** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 2 | 2 |
| **nss-3/linux-docker** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 75 | 75 |
| **nss-3/win11-a64-24h2-builder** | generic-worker | 88.1.3 | multiuser | ffc99d77c0 | windows | amd64 | 1.25.1 | 2 | 2 |
| **nss-t/t-linux-arm64-docker** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | arm64 | 1.25.4 | 2 | 2 |
| **nss-t/t-linux-docker** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 166 | 166 |
| **nss-t/t-linux-docker-amd** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 2 | 2 |
| **nss-t/win11-a64-24h2** | generic-worker | 88.1.3 | multiuser | ffc99d77c0 | windows | amd64 | 1.25.1 | 3 | 3 |
| **nss-t/win11-a64-24h2-alpha** | generic-worker | 93.1.4 | multiuser | 1088d85c3b | windows | arm64 | 1.25.4 | 2 | 2 |
| **proj-autophone/gecko-t-bitbar-gw-perf-a55** | generic-worker | 36.0.0 | simple | b9cb11293f | linux | amd64 | 1.13.7 | 0 | 0 |
| **proj-autophone/gecko-t-bitbar-gw-perf-p6** | generic-worker | 36.0.0 | simple | b9cb11293f | linux | amd64 | 1.13.7 | 0 | 0 |
| **proj-autophone/gecko-t-bitbar-gw-perf-s24** | generic-worker | 36.0.0 | simple | b9cb11293f | linux | amd64 | 1.13.7 | 0 | 0 |
| **proj-autophone/gecko-t-bitbar-gw-unit-p5** | generic-worker | 36.0.0 | simple | b9cb11293f | linux | amd64 | 1.13.7 | 0 | 0 |
| **proj-autophone/gecko-t-lambda-perf-a55** | generic-worker | 87.0.0 | insecure | 99a1fcafbb | linux | amd64 | 1.24.5 | 0 | 0 |
| **releng-1/decision** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 14 | 14 |
| **releng-1/linux-docker** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 9 | 9 |
| **releng-3/decision** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 10 | 10 |
| **releng-3/linux-docker** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 16 | 16 |
| **releng-hardware/applicationservices-b-1-osx1015** | generic-worker | 60.3.4 | multiuser | 943a6f2b0d | darwin | amd64 | 1.22.0 | 0 | 0 |
| **releng-hardware/applicationservices-b-3-osx1015** | generic-worker | 60.3.4 | multiuser | 943a6f2b0d | darwin | amd64 | 1.22.0 | 0 | 0 |
| **releng-hardware/enterprise-1-b-osx-arm64** | generic-worker | 60.3.4 | multiuser | 943a6f2b0d | darwin | arm64 | 1.22.0 | 0 | 0 |
| **releng-hardware/enterprise-3-b-osx-arm64** | generic-worker | 60.3.4 | multiuser | 943a6f2b0d | darwin | arm64 | 1.22.0 | 0 | 0 |
| **releng-hardware/gecko-1-b-osx-1015** | generic-worker | 60.3.4 | multiuser | 943a6f2b0d | darwin | amd64 | 1.22.0 | 0 | 0 |
| **releng-hardware/gecko-1-b-osx-1015-staging** | generic-worker | 60.3.4 | multiuser | 943a6f2b0d | darwin | amd64 | 1.22.0 | 0 | 0 |
| **releng-hardware/gecko-1-b-osx-arm64** | generic-worker | 60.3.4 | multiuser | 943a6f2b0d | darwin | arm64 | 1.22.0 | 0 | 0 |
| **releng-hardware/gecko-3-b-osx-1015** | generic-worker | 60.3.4 | multiuser | 943a6f2b0d | darwin | amd64 | 1.22.0 | 0 | 0 |
| **releng-hardware/gecko-3-b-osx-arm64** | generic-worker | 60.3.4 | multiuser | 943a6f2b0d | darwin | arm64 | 1.22.0 | 0 | 0 |
| **releng-hardware/gecko-t-linux-netperf-1804** | generic-worker | 65.1.0 | insecure | 1a085daa37 | linux | amd64 | 1.22.3 | 0 | 0 |
| **releng-hardware/gecko-t-linux-netperf-2404** | generic-worker | 88.0.2 | insecure | fcaf4a25fc | linux | amd64 | 1.24.5 | 0 | 0 |
| **releng-hardware/gecko-t-linux-talos-1804** | generic-worker | 61.0.0 | simple | 3bd4419b4b | linux | amd64 | 1.22.1 | 0 | 0 |
| **releng-hardware/gecko-t-linux-talos-2404** | generic-worker | 88.0.2 | insecure | fcaf4a25fc | linux | amd64 | 1.24.5 | 0 | 0 |
| **releng-hardware/gecko-t-osx-1015-r8** | generic-worker | 60.3.4 | simple | 943a6f2b0d | darwin | amd64 | 1.22.0 | 0 | 0 |
| **releng-hardware/gecko-t-osx-1015-r8-staging** | generic-worker | 60.3.4 | simple | 943a6f2b0d | darwin | amd64 | 1.22.0 | 0 | 0 |
| **releng-hardware/gecko-t-osx-1400-r8** | generic-worker | 60.3.4 | simple | 943a6f2b0d | darwin | amd64 | 1.22.0 | 0 | 0 |
| **releng-hardware/gecko-t-osx-1400-r8-staging** | generic-worker | 60.3.4 | simple | 943a6f2b0d | darwin | amd64 | 1.22.0 | 0 | 0 |
| **releng-hardware/gecko-t-osx-1500-m-vms** | generic-worker | 60.3.4 | simple | 943a6f2b0d | darwin | arm64 | 1.22.0 | 0 | 0 |
| **releng-hardware/gecko-t-osx-1500-m4** | generic-worker | 91.0.2 | multiuser | 06628b3721 | darwin | arm64 | 1.25.3 | 0 | 0 |
| **releng-hardware/gecko-t-osx-1500-m4-ipv6** | generic-worker | 60.3.4 | simple | 943a6f2b0d | darwin | arm64 | 1.22.0 | 0 | 0 |
| **releng-hardware/gecko-t-osx-1500-m4-staging** | generic-worker | 91.0.2 | multiuser | 06628b3721 | darwin | arm64 | 1.25.3 | 0 | 0 |
| **releng-hardware/gecko-t-win7-32-hw** | generic-worker | 45.0.0 | multiuser | 988e8100b3 | windows | 386 | 1.19.3 | 0 | 0 |
| **releng-hardware/mozillavpn-b-1-osx** | generic-worker | 60.3.4 | multiuser | 943a6f2b0d | darwin | amd64 | 1.22.0 | 0 | 0 |
| **releng-hardware/mozillavpn-b-3-osx** | generic-worker | 60.3.4 | multiuser | 943a6f2b0d | darwin | amd64 | 1.22.0 | 0 | 0 |
| **releng-hardware/nss-1-b-osx-1015** | generic-worker | 60.3.4 | multiuser | 943a6f2b0d | darwin | amd64 | 1.22.0 | 0 | 0 |
| **releng-hardware/nss-3-b-osx-1015** | generic-worker | 60.3.4 | multiuser | 943a6f2b0d | darwin | amd64 | 1.22.0 | 0 | 0 |
| **releng-hardware/win10-64-2009-hw** | generic-worker | 86.0.2 | multiuser | 9ce326432e | windows | amd64 | 1.24.4 | 0 | 0 |
| **releng-hardware/win10-64-2009-hw-alpha** | generic-worker | 86.0.2 | multiuser | 9ce326432e | windows | amd64 | 1.24.4 | 0 | 0 |
| **releng-hardware/win11-64-24h2-hw** | generic-worker | 91.1.1 | multiuser | 9ee80dc4d0 | windows | amd64 | 1.25.3 | 0 | 0 |
| **releng-hardware/win11-64-24h2-hw-alpha** | generic-worker | 93.1.4 | multiuser | 1088d85c3b | windows | amd64 | 1.25.4 | 0 | 0 |
| **releng-hardware/win11-64-24h2-hw-perf-sheriff** | generic-worker | 86.0.2 | multiuser | 9ce326432e | windows | amd64 | 1.24.4 | 0 | 0 |
| **releng-hardware/win11-64-24h2-hw-ref** | generic-worker | 91.1.1 | multiuser | 9ee80dc4d0 | windows | amd64 | 1.25.3 | 0 | 0 |
| **releng-hardware/win11-64-24h2-hw-ref-alpha** | generic-worker | 93.1.4 | multiuser | 1088d85c3b | windows | amd64 | 1.25.4 | 0 | 0 |
| **releng-hardware/win11-64-24h2-hw-relops1213** | generic-worker | 93.1.4 | multiuser | 1088d85c3b | windows | amd64 | 1.25.4 | 0 | 0 |
| **releng-t/linux-docker** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 68 | 68 |
| **relops-1/decision** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 2 | 2 |
| **relops-1/images** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 2 | 2 |
| **relops-3/decision** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 2 | 2 |
| **relops-3/images** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 2 | 2 |
| **scriptworker-1/decision** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 2 | 2 |
| **scriptworker-1/images** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 2 | 2 |
| **scriptworker-3/decision** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 2 | 2 |
| **scriptworker-3/images** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 2 | 2 |
| **taskgraph-1/decision** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 15 | 15 |
| **taskgraph-1/images** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 10 | 10 |
| **taskgraph-3/decision** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 7 | 7 |
| **taskgraph-3/images** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 6 | 6 |
| **taskgraph-t/linux-docker** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 105 | 105 |
| **taskgraph-t/win11-64-24h2** | generic-worker | 88.1.3 | multiuser | ffc99d77c0 | windows | amd64 | 1.25.1 | 9 | 9 |
| **translations-1/b-linux-large-gcp-1tb-32-256-d2g** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 8 | 8 |
| **translations-1/b-linux-large-gcp-1tb-32-256-std-d2g** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 7 | 7 |
| **translations-1/b-linux-large-gcp-1tb-64-512-d2g** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 2 | 2 |
| **translations-1/b-linux-large-gcp-1tb-64-512-std-d2g** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 3 | 3 |
| **translations-1/b-linux-large-gcp-d2g** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 185 | 185 |
| **translations-1/b-linux-large-gcp-d2g-1tb** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 2 | 2 |
| **translations-1/b-linux-large-gcp-d2g-1tb-standard** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 2 | 2 |
| **translations-1/b-linux-large-gcp-d2g-300gb** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 109 | 109 |
| **translations-1/b-linux-v100-gpu-d2g** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 63 | 63 |
| **translations-1/b-linux-v100-gpu-d2g-4** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 46 | 46 |
| **translations-1/b-linux-v100-gpu-d2g-4-1tb** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 4 | 4 |
| **translations-1/b-linux-v100-gpu-d2g-4-1tb-alpha** | generic-worker | 84.1.2 | multiuser | 4a139f395a | linux | amd64 | 1.24.4 | 2 | 2 |
| **translations-1/b-linux-v100-gpu-d2g-4-1tb-standard** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 2 | 2 |
| **translations-1/b-linux-v100-gpu-d2g-4-1tb-std-alpha** | generic-worker | 84.1.2 | multiuser | 4a139f395a | linux | amd64 | 1.24.4 | 2 | 2 |
| **translations-1/b-linux-v100-gpu-d2g-4-2tb** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 7 | 7 |
| **translations-1/b-linux-v100-gpu-d2g-4-2tb-alpha** | generic-worker | 84.1.2 | multiuser | 4a139f395a | linux | amd64 | 1.24.4 | 2 | 2 |
| **translations-1/b-linux-v100-gpu-d2g-4-300gb** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 21 | 21 |
| **translations-1/b-linux-v100-gpu-d2g-4-300gb-alpha** | generic-worker | 84.1.2 | multiuser | 4a139f395a | linux | amd64 | 1.24.4 | 2 | 2 |
| **translations-1/b-linux-v100-gpu-d2g-4-300gb-standard** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 2 | 2 |
| **translations-1/b-linux-v100-gpu-d2g-4-300gb-std-alpha** | generic-worker | 84.1.2 | multiuser | 4a139f395a | linux | amd64 | 1.24.4 | 2 | 2 |
| **translations-1/b-linux-v100-gpu-d2g-4-alpha** | generic-worker | 84.1.2 | multiuser | 4a139f395a | linux | amd64 | 1.24.4 | 2 | 2 |
| **translations-1/b-linux-v100-gpu-d2g-alpha** | generic-worker | 84.1.2 | multiuser | 4a139f395a | linux | amd64 | 1.24.4 | 2 | 2 |
| **translations-1/decision** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 1148 | 1148 |
| **translations-1/images** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 2 | 2 |
| **xpi-1/b-linux** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 48 | 48 |
| **xpi-1/decision** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 19 | 19 |
| **xpi-1/images** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 2 | 2 |
| **xpi-3/b-linux** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 2 | 2 |
| **xpi-3/decision** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 2 | 2 |
| **xpi-3/images** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 2 | 2 |


## Docker Worker

Total: `157`

Count by version:

| Version | Count |
| :--- | ---: |
| 38.0.5 | 137 |
| 48.3.0 | 1 |
| unknown version | 19 |


Count by image:

| Version | Count |
| :--- | ---: |
| projects/fxci-production-level3-workers/global/images/docker-firefoxci-gcp-l3-googlecompute-2024-02-05t23-18-22z | 52 |
| projects/taskcluster-imaging/global/images/docker-worker-gcp-u14-04-2025-06-16 | 19 |
| projects/taskcluster-imaging/global/images/docker-firefoxci-gcp-lt-googlecompute-2023-04-13t21-30-28z | 1 |
| projects/taskcluster-imaging/global/images/docker-firefoxci-gcp-l1-googlecompute-2025-06-13t18-31-38z | 85 |


| Worker Pool | Implementation | Version | Total Workers | Total Capacity |
| --- | --- | --- | ---: | ---: |
| **app-services-1/b-linux-amd** | docker-worker | 38.0.5 | 3 | 3 |
| **app-services-1/b-linux-gcp** | docker-worker | 38.0.5 | 387 | 387 |
| **app-services-1/decision-gcp** | docker-worker | 38.0.5 | 822 | 1644 |
| **app-services-1/images-gcp** | docker-worker | 38.0.5 | 5 | 5 |
| **app-services-3/b-linux-amd** | docker-worker | 38.0.5 | 2 | 2 |
| **app-services-3/b-linux-gcp** | docker-worker | 38.0.5 | 159 | 159 |
| **app-services-3/decision-gcp** | docker-worker | 38.0.5 | 11 | 22 |
| **app-services-3/images-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **code-analysis-1/linux-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **code-analysis-3/linux-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **comm-1/b-linux-amd** | docker-worker | 38.0.5 | 2 | 2 |
| **comm-1/b-linux-gcp** | docker-worker | 38.0.5 | 196 | 196 |
| **comm-1/b-linux-large-amd** | docker-worker | 38.0.5 | 2 | 2 |
| **comm-1/b-linux-large-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **comm-1/b-linux-xlarge-amd** | docker-worker | 38.0.5 | 2 | 2 |
| **comm-1/b-linux-xlarge-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **comm-1/decision-gcp** | docker-worker | 38.0.5 | 65 | 65 |
| **comm-1/images-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **comm-2/b-linux-amd** | docker-worker | 38.0.5 | 3 | 3 |
| **comm-2/b-linux-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **comm-2/b-linux-large-amd** | docker-worker | 38.0.5 | 2 | 2 |
| **comm-2/b-linux-large-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **comm-2/b-linux-xlarge-amd** | docker-worker | 38.0.5 | 2 | 2 |
| **comm-2/b-linux-xlarge-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **comm-2/decision-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **comm-2/images-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **comm-3/b-linux-amd** | docker-worker | 38.0.5 | 3 | 3 |
| **comm-3/b-linux-gcp** | docker-worker | 38.0.5 | 2363 | 2363 |
| **comm-3/b-linux-large-amd** | docker-worker | 38.0.5 | 3 | 3 |
| **comm-3/b-linux-large-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **comm-3/b-linux-xlarge-amd** | docker-worker | 38.0.5 | 2 | 2 |
| **comm-3/b-linux-xlarge-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **comm-3/decision-gcp** | docker-worker | 38.0.5 | 107 | 107 |
| **comm-3/images-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **comm-t/misc-gcp** | docker-worker | 38.0.5 | 2 | 16 |
| **comm-t/t-linux-large-gcp** | docker-worker | unknown version | 2 | 2 |
| **comm-t/t-linux-large-noscratch-gcp** | docker-worker | unknown version | 2 | 2 |
| **comm-t/t-linux-xlarge-gcp** | docker-worker | unknown version | 2 | 2 |
| **comm-t/t-linux-xlarge-noscratch-gcp** | docker-worker | unknown version | 2 | 2 |
| **comm-t/t-linux-xlarge-source-gcp** | docker-worker | unknown version | 2 | 2 |
| **comm-t/t-linux-xlarge-source-noscratch-gcp** | docker-worker | unknown version | 2 | 2 |
| **gecko-1/b-linux-amd** | docker-worker | 38.0.5 | 539 | 539 |
| **gecko-1/b-linux-gcp** | docker-worker | 38.0.5 | 105 | 105 |
| **gecko-1/b-linux-gcp-relops1411** | docker-worker | 38.0.5 | 2 | 2 |
| **gecko-1/b-linux-gcp-test-bug-1882320** | docker-worker | 38.0.5 | 2 | 2 |
| **gecko-1/b-linux-kvm-gcp** | docker-worker | 38.0.5 | 49 | 49 |
| **gecko-1/b-linux-large-amd** | docker-worker | 38.0.5 | 2 | 2 |
| **gecko-1/b-linux-large-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **gecko-1/b-linux-medium-gcp** | docker-worker | 38.0.5 | 54 | 54 |
| **gecko-1/b-linux-xlarge-amd** | docker-worker | 38.0.5 | 92 | 92 |
| **gecko-1/b-linux-xlarge-gcp** | docker-worker | 38.0.5 | 79 | 79 |
| **gecko-1/b-linux-xlarge-gcp-bug1797804-c2** | docker-worker | 38.0.5 | 2 | 2 |
| **gecko-1/decision-gcp** | docker-worker | 38.0.5 | 865 | 865 |
| **gecko-1/decision-gcp-c3d4-amd-bug1990935** | docker-worker | 38.0.5 | 2 | 2 |
| **gecko-1/decision-gcp-c3d8-amd-bug1990935** | docker-worker | 38.0.5 | 2 | 2 |
| **gecko-1/decision-gcp-c4d4-amd-bug1990935** | docker-worker | 38.0.5 | 2 | 2 |
| **gecko-1/decision-gcp-c4d4-hcpu-amd-bug1990935** | docker-worker | 38.0.5 | 2 | 2 |
| **gecko-1/decision-gcp-c4d8-amd-bug1990935** | docker-worker | 38.0.5 | 2 | 2 |
| **gecko-1/decision-gcp-c4d8-hcpu-amd-bug1990935** | docker-worker | 38.0.5 | 2 | 2 |
| **gecko-1/decision-gcp-c4d8-lssd-amd-bug1990935** | docker-worker | 38.0.5 | 2 | 2 |
| **gecko-1/images-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **gecko-2/b-linux-amd** | docker-worker | 38.0.5 | 24 | 24 |
| **gecko-2/b-linux-gcp** | docker-worker | 38.0.5 | 4 | 4 |
| **gecko-2/b-linux-kvm-gcp** | docker-worker | 38.0.5 | 3 | 3 |
| **gecko-2/b-linux-large-amd** | docker-worker | 38.0.5 | 3 | 3 |
| **gecko-2/b-linux-large-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **gecko-2/b-linux-medium-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **gecko-2/b-linux-xlarge-amd** | docker-worker | 38.0.5 | 2 | 2 |
| **gecko-2/b-linux-xlarge-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **gecko-2/decision-gcp** | docker-worker | 38.0.5 | 10 | 10 |
| **gecko-2/images-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **gecko-3/b-linux-amd** | docker-worker | 38.0.5 | 1084 | 1084 |
| **gecko-3/b-linux-gcp** | docker-worker | 38.0.5 | 881 | 881 |
| **gecko-3/b-linux-kvm-gcp** | docker-worker | 38.0.5 | 5 | 5 |
| **gecko-3/b-linux-large-amd** | docker-worker | 38.0.5 | 3 | 3 |
| **gecko-3/b-linux-large-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **gecko-3/b-linux-medium-gcp** | docker-worker | 38.0.5 | 178 | 178 |
| **gecko-3/b-linux-xlarge-amd** | docker-worker | 38.0.5 | 18 | 18 |
| **gecko-3/b-linux-xlarge-gcp** | docker-worker | 38.0.5 | 37 | 37 |
| **gecko-3/decision-gcp** | docker-worker | 38.0.5 | 889 | 889 |
| **gecko-3/images-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **gecko-t/misc-gcp** | docker-worker | 38.0.5 | 13 | 104 |
| **gecko-t/t-linux-kvm-gcp** | docker-worker | unknown version | 2 | 2 |
| **gecko-t/t-linux-kvm-gcp-relops1411** | docker-worker | unknown version | 2 | 2 |
| **gecko-t/t-linux-kvm-noscratch-gcp** | docker-worker | unknown version | 5 | 5 |
| **gecko-t/t-linux-large-gcp** | docker-worker | unknown version | 2 | 2 |
| **gecko-t/t-linux-large-hostub2204-gcp** | docker-worker | 48.3.0 | 2 | 2 |
| **gecko-t/t-linux-large-noscratch-gcp** | docker-worker | unknown version | 1182 | 1182 |
| **gecko-t/t-linux-xlarge-gcp** | docker-worker | unknown version | 19 | 19 |
| **gecko-t/t-linux-xlarge-noscratch-gcp** | docker-worker | unknown version | 738 | 738 |
| **gecko-t/t-linux-xlarge-source-gcp** | docker-worker | unknown version | 88 | 88 |
| **gecko-t/t-linux-xlarge-source-noscratch-gcp** | docker-worker | unknown version | 2 | 2 |
| **glean-1/b-linux-gcp** | docker-worker | 38.0.5 | 11 | 11 |
| **glean-1/decision-gcp** | docker-worker | 38.0.5 | 826 | 1652 |
| **glean-1/images-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **glean-3/b-linux-gcp** | docker-worker | 38.0.5 | 18 | 18 |
| **glean-3/decision-gcp** | docker-worker | 38.0.5 | 6 | 12 |
| **glean-3/images-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **infra/build-decision** | docker-worker | 38.0.5 | 2 | 16 |
| **mobile-1/b-linux-gcp** | docker-worker | 38.0.5 | 60 | 60 |
| **mobile-1/decision-gcp** | docker-worker | 38.0.5 | 776 | 1552 |
| **mobile-1/images-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **mobile-3/b-linux-gcp** | docker-worker | 38.0.5 | 41 | 41 |
| **mobile-3/decision-gcp** | docker-worker | 38.0.5 | 45 | 90 |
| **mobile-3/images-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **mozilla-1/b-linux-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **mozilla-1/decision-gcp** | docker-worker | 38.0.5 | 814 | 1628 |
| **mozilla-1/images-gcp** | docker-worker | 38.0.5 | 14 | 14 |
| **mozilla-3/b-linux-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **mozilla-3/decision-gcp** | docker-worker | 38.0.5 | 401 | 802 |
| **mozilla-3/images-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **mozilla-t/t-linux-large-gcp** | docker-worker | unknown version | 25 | 25 |
| **mozilla-t/t-linux-large-noscratch-gcp** | docker-worker | unknown version | 2 | 2 |
| **mozillavpn-1/b-linux-gcp** | docker-worker | 38.0.5 | 78 | 78 |
| **mozillavpn-1/b-linux-large-gcp** | docker-worker | 38.0.5 | 50 | 50 |
| **mozillavpn-1/decision-gcp** | docker-worker | 38.0.5 | 823 | 1646 |
| **mozillavpn-1/images-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **mozillavpn-3/b-linux-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **mozillavpn-3/b-linux-large-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **mozillavpn-3/decision-gcp** | docker-worker | 38.0.5 | 2 | 4 |
| **mozillavpn-3/images-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **nss-1/decision-gcp** | docker-worker | 38.0.5 | 3 | 6 |
| **nss-1/images-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **nss-1/linux-gcp** | docker-worker | 38.0.5 | 83 | 83 |
| **nss-3/decision-gcp** | docker-worker | 38.0.5 | 2 | 4 |
| **nss-3/images-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **nss-3/linux-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **nss-t/t-linux-xlarge-gcp** | docker-worker | unknown version | 2 | 2 |
| **nss-t/t-linux-xlarge-noscratch-gcp** | docker-worker | unknown version | 2 | 2 |
| **releng-1/b-linux-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **releng-1/decision-gcp** | docker-worker | 38.0.5 | 2 | 4 |
| **releng-1/linux-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **releng-3/b-linux-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **releng-3/decision-gcp** | docker-worker | 38.0.5 | 2 | 4 |
| **releng-3/linux-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **relops-1/decision-gcp** | docker-worker | 38.0.5 | 2 | 4 |
| **relops-1/images-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **relops-3/decision-gcp** | docker-worker | 38.0.5 | 2 | 4 |
| **relops-3/images-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **scriptworker-1/b-linux-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **scriptworker-1/decision-gcp** | docker-worker | 38.0.5 | 2 | 4 |
| **scriptworker-1/images-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **scriptworker-3/b-linux-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **scriptworker-3/decision-gcp** | docker-worker | 38.0.5 | 2 | 4 |
| **scriptworker-3/images-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **taskgraph-1/decision-gcp** | docker-worker | 38.0.5 | 4 | 8 |
| **taskgraph-1/images-gcp** | docker-worker | 38.0.5 | 3 | 3 |
| **taskgraph-3/decision-gcp** | docker-worker | 38.0.5 | 2 | 4 |
| **taskgraph-3/images-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **translations-1/decision-gcp** | docker-worker | 38.0.5 | 1926 | 3852 |
| **translations-1/images-gcp** | docker-worker | 38.0.5 | 7 | 7 |
| **xpi-1/b-linux-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **xpi-1/decision-gcp** | docker-worker | 38.0.5 | 8 | 16 |
| **xpi-1/images-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **xpi-3/b-linux-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **xpi-3/decision-gcp** | docker-worker | 38.0.5 | 2 | 4 |
| **xpi-3/images-gcp** | docker-worker | 38.0.5 | 2 | 2 |


## Script Worker

Total: `32`



| Worker Pool | Implementation | Version | Total Workers | Total Capacity |
| --- | --- | --- | ---: | ---: |
| **scriptworker-k8s/comm-3-balrog** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/comm-3-beetmover** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/comm-3-bouncer** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/comm-3-shipit** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/comm-3-signing** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/comm-3-tree** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/comm-t-signing** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/gecko-1-beetmover** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/gecko-3-balrog** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/gecko-3-beetmover** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/gecko-3-bouncer** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/gecko-3-lando** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/gecko-3-pushapk** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/gecko-3-shipit** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/gecko-3-signing** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/gecko-t-signing** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/mobile-1-pushapk** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/mobile-3-bitrise** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/mobile-3-pushapk** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/mobile-3-signing** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/mobile-t-signing** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/mozillavpn-3-signing** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/translations-1-beetmover** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-prov-v1/comm-signing-mac14m2** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-prov-v1/dep-adhoc-signing-mac14m2** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-prov-v1/dep-comm-signing-mac14m2** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-prov-v1/dep-enterprise-signing-mac14m2** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-prov-v1/dep-gecko-signing-mac14m2** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-prov-v1/dep-mozillavpn-signing-mac14m2** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-prov-v1/enterprise-signing-mac14m2** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-prov-v1/gecko-signing-mac14m2** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-prov-v1/mozillavpn-signing-mac14m2** | Scriptworker Chain of Trust | <no value> | 0 | 0 |


## No artifacts found [^1]

Total: `4`


Count by image:

| Version | Count |
| :--- | ---: |
| projects/taskcluster-imaging/global/images/generic-2204-wayland-vm-gcp-googlecompute-2023-09-22t17-39-37z | 2 |
|  | 2 |


| Worker Pool | Implementation | Version | Total Workers | Total Capacity |
| --- | --- | --- | ---: | ---: |
| **built-in/fail** |  | No artifacts found | 0 | 0 |
| **built-in/succeed** |  | No artifacts found | 0 | 0 |
| **gecko-t/t-linux-vm-2204-wayland** |  | No artifacts found | 54 | 54 |
| **gecko-t/t-linux-vm-2204-wayland-snap** |  | No artifacts found | 2 | 2 |


## Version not determined [^2]

Total: `3`


Count by image:

| Version | Count |
| :--- | ---: |
|  | 1 |
| unknown | 1 |
| projects/taskcluster-imaging/global/images/gw-fxci-gcp-l1-2404-amd64-googlecompute-alpha | 1 |


| Worker Pool | Implementation | Version | Total Workers | Total Capacity |
| --- | --- | --- | ---: | ---: |
| **gecko-t/t-linux-2404-relsre** |  | Version not determined; task not (yet) claimed | 3 | 3 |
| **gecko-t/win11-64-24h2-amd** |  | Version not determined; task not (yet) claimed | 190 | 190 |
| **scriptworker-prov-v1/adhoc-signing-mac14m2** |  | Version not determined; task not (yet) claimed | 0 | 0 |



[^1]: Those are the pools whose tasks were claimed and resolved by a worker as expected, but the worker did not publish either artifact `public/logs/live_backing.log` nor `public/logs/chain_of_trust.log`, which is the source used to identify the worker implementation.

[^2]: Probing task remains pending after two hours. Those are the pools that were not able to start any worker to claim the task within two hours.



# Worker Pool Versions


## Generic Worker

Total: `353`

Count by version:

| Version | Count |
| :--- | ---: |
| 36.0.0 | 3 |
| 45.0.0 | 1 |
| 60.3.4 | 16 |
| 61.0.0 | 1 |
| 64.3.0 | 26 |
| 72.0.1 | 1 |
| 84.1.2 | 7 |
| 86.0.2 | 3 |
| 87.0.0 | 1 |
| 88.0.2 | 2 |
| 88.1.3 | 50 |
| 91.0.2 | 2 |
| 91.1.0 | 4 |
| 91.1.1 | 1 |
| 93.1.2 | 3 |
| 93.1.4 | 13 |
| 94.1.0 | 1 |
| 95.1.1 | 7 |
| 95.1.3 | 211 |


Count by image:

| Version | Count |
| :--- | ---: |
| projects/taskcluster-imaging/global/images/gw-fxci-gcp-l1-2404-amd64-headless-googlecompute-2026-01-14 | 120 |
|  | 35 |
| projects/taskcluster-imaging/global/images/gw-fxci-gcp-l1-gui-googlecompute-2025-01-13t22-33-40z | 2 |
| projects/taskcluster-imaging/global/images/gw-fxci-gcp-l1-2404-arm64-headless-googlecompute-2026-01-14 | 14 |
| unknown | 58 |
| /subscriptions/108d46d5-fe9b-4850-9a7d-8c914aa6c1f0/resourceGroups/rg-packer-worker-images/providers/Microsoft.Compute/galleries/win11_a64_24h2_tester_alpha/images/win11_a64_24h2_tester_alpha/versions/1.0.0 | 4 |
| projects/taskcluster-imaging/global/images/gw-fxci-gcp-l1-gui-googlecompute-2024-08-22t22-48-09z | 8 |
| /subscriptions/108d46d5-fe9b-4850-9a7d-8c914aa6c1f0/resourceGroups/rg-packer-worker-images/providers/Microsoft.Compute/galleries/win11_a64_24h2_builder_alpha/images/win11_a64_24h2_builder_alpha/versions/1.0.0 | 2 |
| projects/fxci-production-level3-workers/global/images/gw-fxci-gcp-l3-2404-arm64-headless-googlecompute-2026-01-14 | 6 |
| projects/fxci-production-level3-workers/global/images/gw-fxci-gcp-l3-2404-amd64-headless-googlecompute-alpha | 1 |
| projects/taskcluster-imaging/global/images/gw-fxci-gcp-l1-2404-amd64-gui-googlecompute-alpha | 1 |
| projects/taskcluster-imaging/global/images/gw-fxci-gcp-l1-2404-amd64-gui-googlecompute-2026-01-14 | 4 |
| projects/taskcluster-imaging/global/images/gw-fxci-gcp-l1-arm64-gui-googlecompute-2024-09-18t14-57-52z | 9 |
| projects/fxci-production-level3-workers/global/images/gw-fxci-gcp-l3-gui-googlecompute-2024-09-18t05-46-31z | 3 |
| /subscriptions/a30e97ab-734a-4f3b-a0e4-c51c0bff0701/resourceGroups/rg-packer-worker-images/providers/Microsoft.Compute/galleries/trusted_win11_a64_24h2_builder/images/trusted_win11_a64_24h2_builder/versions/1.0.8 | 3 |
| projects/fxci-production-level3-workers/global/images/gw-fxci-gcp-l3-2404-amd64-headless-googlecompute-2026-01-14 | 58 |
| projects/taskcluster-imaging/global/images/gw-fxci-gcp-l1-2404-arm64-headless-googlecompute-alpha | 1 |
| projects/taskcluster-imaging/global/images/gw-fxci-gcp-l1-2404-amd64-headless-googlecompute-alpha | 6 |
| projects/taskcluster-imaging/global/images/gw-fxci-gcp-l1-2404-amd64-headless-googlecompute-alpha-tc | 7 |
| /subscriptions/108d46d5-fe9b-4850-9a7d-8c914aa6c1f0/resourceGroups/rg-packer-worker-images/providers/Microsoft.Compute/galleries/win11_a64_24h2_tester/images/win11_a64_24h2_tester/versions/1.0.8 | 4 |
| /subscriptions/108d46d5-fe9b-4850-9a7d-8c914aa6c1f0/resourceGroups/rg-packer-worker-images/providers/Microsoft.Compute/galleries/win11_a64_24h2_builder/images/win11_a64_24h2_builder/versions/1.0.8 | 3 |
| projects/fxci-production-level3-workers/global/images/gw-fxci-gcp-l3-arm64-gui-googlecompute-2024-09-18t19-02-58z | 4 |


| Worker Pool | Implementation | Version | Engine | Revision | OS | Arch | GO | Total Workers | Total Capacity |
| --- | --- | --- | --- | --- | --- | --- | --- | ---: | ---: |
| **adhoc-1/decision** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 2 | 2 |
| **adhoc-1/images** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 3 | 3 |
| **adhoc-3/decision** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 2 | 2 |
| **adhoc-3/images** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 2 | 2 |
| **app-services-1/b-linux** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 519 | 519 |
| **app-services-1/b-linux-docker-amd** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 2 | 2 |
| **app-services-1/decision** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 1083 | 1083 |
| **app-services-1/images** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 4 | 4 |
| **app-services-3/b-linux** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 956 | 956 |
| **app-services-3/b-linux-docker-amd** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 2 | 2 |
| **app-services-3/decision** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 32 | 32 |
| **app-services-3/images** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 3 | 3 |
| **code-analysis-1/linux-docker** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 2 | 2 |
| **code-analysis-1/linux-gw-gcp** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | amd64 | 1.22.2 | 26 | 26 |
| **code-analysis-3/linux-docker** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 2 | 2 |
| **code-analysis-3/linux-gw-gcp** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | amd64 | 1.22.2 | 2 | 2 |
| **code-coverage/bot** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 267 | 267 |
| **code-review/bot** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 7344 | 7344 |
| **comm-1/b-linux** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 4 | 4 |
| **comm-1/b-linux-aarch64** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | arm64 | 1.25.5 | 2 | 2 |
| **comm-1/b-linux-docker-amd** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 32 | 32 |
| **comm-1/b-linux-docker-large-amd** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 6 | 6 |
| **comm-1/b-linux-docker-xlarge-amd** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 2 | 2 |
| **comm-1/b-linux-gcp-aarch64** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | arm64 | 1.22.2 | 2 | 2 |
| **comm-1/b-linux-large** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 2 | 2 |
| **comm-1/b-linux-xlarge** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 2 | 2 |
| **comm-1/b-win2022** | generic-worker | 88.1.3 | multiuser | ffc99d77c0 | windows | amd64 | 1.25.1 | 5 | 5 |
| **comm-1/decision** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 2 | 2 |
| **comm-1/images** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 2 | 2 |
| **comm-1/images-aarch64** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | arm64 | 1.25.5 | 2 | 2 |
| **comm-1/images-gcp-aarch64** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | arm64 | 1.22.2 | 2 | 2 |
| **comm-2/b-linux** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 2 | 2 |
| **comm-2/b-linux-aarch64** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | arm64 | 1.25.5 | 2 | 2 |
| **comm-2/b-linux-docker-amd** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 2 | 2 |
| **comm-2/b-linux-docker-large-amd** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 2 | 2 |
| **comm-2/b-linux-docker-xlarge-amd** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 2 | 2 |
| **comm-2/b-linux-gcp-aarch64** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | arm64 | 1.22.2 | 2 | 2 |
| **comm-2/b-linux-large** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 2 | 2 |
| **comm-2/b-linux-xlarge** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 2 | 2 |
| **comm-2/decision** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 2 | 2 |
| **comm-2/images** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 2 | 2 |
| **comm-2/images-aarch64** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | arm64 | 1.25.5 | 2 | 2 |
| **comm-2/images-gcp-aarch64** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | arm64 | 1.22.2 | 2 | 2 |
| **comm-3/b-linux** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 5 | 5 |
| **comm-3/b-linux-aarch64** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | arm64 | 1.25.5 | 2 | 2 |
| **comm-3/b-linux-docker-amd** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 1849 | 1849 |
| **comm-3/b-linux-docker-large-amd** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 11 | 11 |
| **comm-3/b-linux-docker-xlarge-amd** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 7 | 7 |
| **comm-3/b-linux-gcp-aarch64** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | arm64 | 1.22.2 | 2 | 2 |
| **comm-3/b-linux-large** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 2 | 2 |
| **comm-3/b-linux-xlarge** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 2 | 2 |
| **comm-3/b-win2022** | generic-worker | 88.1.3 | multiuser | ffc99d77c0 | windows | amd64 | 1.25.1 | 693 | 693 |
| **comm-3/decision** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 2 | 2 |
| **comm-3/images** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 4 | 4 |
| **comm-3/images-aarch64** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | arm64 | 1.25.5 | 2 | 2 |
| **comm-3/images-gcp-aarch64** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | arm64 | 1.22.2 | 2 | 2 |
| **comm-t/misc** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 2 | 2 |
| **comm-t/t-linux-arm64-docker** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | arm64 | 1.25.5 | 2 | 2 |
| **comm-t/t-linux-docker** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 34 | 34 |
| **comm-t/t-linux-docker-amd** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 1781 | 1781 |
| **comm-t/t-linux-docker-noscratch** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 117 | 117 |
| **comm-t/t-linux-docker-noscratch-amd** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 2745 | 2745 |
| **comm-t/win11-64-2009** | generic-worker | 88.1.3 | multiuser | ffc99d77c0 | windows | amd64 | 1.25.1 | 2 | 2 |
| **comm-t/win11-64-2009-ssd** | generic-worker | 88.1.3 | multiuser | ffc99d77c0 | windows | amd64 | 1.25.1 | 2 | 2 |
| **comm-t/win11-64-24h2** | generic-worker | 88.1.3 | multiuser | ffc99d77c0 | windows | amd64 | 1.25.1 | 4101 | 4101 |
| **comm-t/win11-64-24h2-ssd** | generic-worker | 88.1.3 | multiuser | ffc99d77c0 | windows | amd64 | 1.25.1 | 414 | 414 |
| **comm-t/win11-a64-24h2** | generic-worker | 88.1.3 | multiuser | ffc99d77c0 | windows | amd64 | 1.25.1 | 2 | 2 |
| **comm-t/win11-a64-24h2-alpha** | generic-worker | 93.1.4 | multiuser | 1088d85c3b | windows | arm64 | 1.25.4 | 2 | 2 |
| **enterprise-1/b-linux** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 927 | 927 |
| **enterprise-1/b-linux-aarch64** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | arm64 | 1.25.5 | 7 | 7 |
| **enterprise-1/b-linux-docker-amd** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 835 | 835 |
| **enterprise-1/b-linux-docker-large-amd** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 30 | 30 |
| **enterprise-1/b-linux-docker-xlarge-amd** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 587 | 587 |
| **enterprise-1/b-win2022** | generic-worker | 88.1.3 | multiuser | ffc99d77c0 | windows | amd64 | 1.25.1 | 262 | 262 |
| **enterprise-1/decision** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 1058 | 1058 |
| **enterprise-1/images** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 35 | 35 |
| **enterprise-1/images-aarch64** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | arm64 | 1.25.5 | 5 | 5 |
| **enterprise-3/b-linux** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 466 | 466 |
| **enterprise-3/b-linux-aarch64** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | arm64 | 1.25.5 | 5 | 5 |
| **enterprise-3/b-linux-docker-amd** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 559 | 559 |
| **enterprise-3/b-linux-docker-large-amd** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 24 | 24 |
| **enterprise-3/b-linux-docker-xlarge-amd** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 252 | 252 |
| **enterprise-3/b-win2022** | generic-worker | 88.1.3 | multiuser | ffc99d77c0 | windows | amd64 | 1.25.1 | 145 | 145 |
| **enterprise-3/decision** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 29 | 29 |
| **enterprise-3/images** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 38 | 38 |
| **enterprise-3/images-aarch64** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | arm64 | 1.25.5 | 5 | 5 |
| **enterprise-t/misc** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 278 | 278 |
| **enterprise-t/t-linux-docker-16c32gb-amd** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 227 | 227 |
| **enterprise-t/t-linux-docker-amd** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 4282 | 4282 |
| **enterprise-t/t-linux-docker-noscratch-amd** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 956 | 956 |
| **enterprise-t/win11-64-24h2** | generic-worker | 88.1.3 | multiuser | ffc99d77c0 | windows | amd64 | 1.25.1 | 718 | 718 |
| **enterprise-t/win11-64-24h2-alpha** | generic-worker | 95.1.1 | multiuser | cbb4697b6c | windows | amd64 | 1.25.5 | 49 | 49 |
| **enterprise-t/win11-64-24h2-source** | generic-worker | 88.1.3 | multiuser | ffc99d77c0 | windows | amd64 | 1.25.1 | 2501 | 2501 |
| **enterprise-t/win11-64-24h2-source-alpha** | generic-worker | 95.1.1 | multiuser | cbb4697b6c | windows | amd64 | 1.25.5 | 526 | 526 |
| **gecko-1/b-linux** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 6105 | 6105 |
| **gecko-1/b-linux-2204-kvm-gcp** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | amd64 | 1.22.2 | 2 | 2 |
| **gecko-1/b-linux-aarch64** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | arm64 | 1.25.5 | 36 | 36 |
| **gecko-1/b-linux-docker-alpha** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 2 | 2 |
| **gecko-1/b-linux-docker-amd** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 8846 | 8846 |
| **gecko-1/b-linux-docker-large-amd** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 163 | 163 |
| **gecko-1/b-linux-docker-xlarge-amd** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 1183 | 1183 |
| **gecko-1/b-linux-gcp-aarch64** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | arm64 | 1.22.2 | 2 | 2 |
| **gecko-1/b-linux-kvm** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 339 | 339 |
| **gecko-1/b-linux-large** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 2 | 2 |
| **gecko-1/b-linux-medium** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 11037 | 11037 |
| **gecko-1/b-linux-xlarge** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 13 | 13 |
| **gecko-1/b-win2022** | generic-worker | 88.1.3 | multiuser | ffc99d77c0 | windows | amd64 | 1.25.1 | 757 | 757 |
| **gecko-1/b-win2022-alpha** | generic-worker | 93.1.2 | multiuser | e954a7555f | windows | amd64 | 1.25.4 | 2 | 2 |
| **gecko-1/b-win2022-xxlarge** | generic-worker | 88.1.3 | multiuser | ffc99d77c0 | windows | amd64 | 1.25.1 | 9 | 9 |
| **gecko-1/decision** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 2595 | 2595 |
| **gecko-1/images** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 77 | 77 |
| **gecko-1/images-aarch64** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | arm64 | 1.25.5 | 9 | 9 |
| **gecko-1/images-gcp-aarch64** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | arm64 | 1.22.2 | 2 | 2 |
| **gecko-1/win11-a64-24h2-builder** | generic-worker | 88.1.3 | multiuser | ffc99d77c0 | windows | amd64 | 1.25.1 | 39 | 39 |
| **gecko-1/win11-a64-24h2-builder-alpha** | generic-worker | 93.1.4 | multiuser | 1088d85c3b | windows | arm64 | 1.25.4 | 2 | 2 |
| **gecko-2/b-linux** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 50 | 50 |
| **gecko-2/b-linux-aarch64** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | arm64 | 1.25.5 | 2 | 2 |
| **gecko-2/b-linux-docker-amd** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 78 | 78 |
| **gecko-2/b-linux-docker-large-amd** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 4 | 4 |
| **gecko-2/b-linux-docker-xlarge-amd** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 12 | 12 |
| **gecko-2/b-linux-gcp-aarch64** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | arm64 | 1.22.2 | 2 | 2 |
| **gecko-2/b-linux-kvm** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 4 | 4 |
| **gecko-2/b-linux-large** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 2 | 2 |
| **gecko-2/b-linux-medium** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 59 | 59 |
| **gecko-2/b-linux-xlarge** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 2 | 2 |
| **gecko-2/b-win2022** | generic-worker | 88.1.3 | multiuser | ffc99d77c0 | windows | amd64 | 1.25.1 | 5 | 5 |
| **gecko-2/decision** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 17 | 17 |
| **gecko-2/images** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 2 | 2 |
| **gecko-2/images-aarch64** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | arm64 | 1.25.5 | 2 | 2 |
| **gecko-2/images-gcp-aarch64** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | arm64 | 1.22.2 | 2 | 2 |
| **gecko-3/b-linux** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 14822 | 14822 |
| **gecko-3/b-linux-2204-kvm-gcp** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | amd64 | 1.22.2 | 2 | 2 |
| **gecko-3/b-linux-aarch64** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | arm64 | 1.25.5 | 12 | 12 |
| **gecko-3/b-linux-docker-alpha** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 2 | 2 |
| **gecko-3/b-linux-docker-amd** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 20433 | 20433 |
| **gecko-3/b-linux-docker-large-amd** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 273 | 273 |
| **gecko-3/b-linux-docker-xlarge-amd** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 2627 | 2627 |
| **gecko-3/b-linux-gcp-aarch64** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | arm64 | 1.22.2 | 2 | 2 |
| **gecko-3/b-linux-kvm** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 497 | 497 |
| **gecko-3/b-linux-large** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 2 | 2 |
| **gecko-3/b-linux-medium** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 20659 | 20659 |
| **gecko-3/b-linux-xlarge** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 31 | 31 |
| **gecko-3/b-win2022** | generic-worker | 88.1.3 | multiuser | ffc99d77c0 | windows | amd64 | 1.25.1 | 5131 | 5131 |
| **gecko-3/b-win2022-xxlarge** | generic-worker | 88.1.3 | multiuser | ffc99d77c0 | windows | amd64 | 1.25.1 | 14 | 14 |
| **gecko-3/decision** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 4484 | 4484 |
| **gecko-3/images** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 32 | 32 |
| **gecko-3/images-aarch64** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | arm64 | 1.25.5 | 5 | 5 |
| **gecko-3/images-gcp-aarch64** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | arm64 | 1.22.2 | 2 | 2 |
| **gecko-3/win11-a64-24h2-builder** | generic-worker | 88.1.3 | multiuser | ffc99d77c0 | windows | amd64 | 1.25.1 | 52 | 52 |
| **gecko-t/misc** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 885 | 885 |
| **gecko-t/t-linux-2204-wayland** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | amd64 | 1.22.2 | 8617 | 8617 |
| **gecko-t/t-linux-2204-wayland-arm64-relsre** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | arm64 | 1.22.2 | 2 | 2 |
| **gecko-t/t-linux-2204-wayland-relsre** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | amd64 | 1.22.2 | 2 | 2 |
| **gecko-t/t-linux-2204-wayland-root-exp** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | amd64 | 1.22.2 | 2 | 2 |
| **gecko-t/t-linux-2204-wayland-snap** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | amd64 | 1.22.2 | 119 | 119 |
| **gecko-t/t-linux-2404-headless-alpha** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 2 | 2 |
| **gecko-t/t-linux-2404-headless-arm64-alpha** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | arm64 | 1.25.5 | 2 | 2 |
| **gecko-t/t-linux-2404-headless-ssd-alpha** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 2 | 2 |
| **gecko-t/t-linux-2404-wayland** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 2 | 2 |
| **gecko-t/t-linux-2404-wayland-relsre** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 2 | 2 |
| **gecko-t/t-linux-2404-wayland-snap** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 128 | 128 |
| **gecko-t/t-linux-arm64-docker** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | arm64 | 1.25.5 | 32 | 32 |
| **gecko-t/t-linux-docker** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 1739 | 1739 |
| **gecko-t/t-linux-docker-16c32gb-amd** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 3922 | 3922 |
| **gecko-t/t-linux-docker-amd** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 31324 | 31324 |
| **gecko-t/t-linux-docker-amd-alpha** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 2 | 2 |
| **gecko-t/t-linux-docker-kvm** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 76453 | 76453 |
| **gecko-t/t-linux-docker-kvm-alpha** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 2 | 2 |
| **gecko-t/t-linux-docker-noscratch** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 1678 | 1678 |
| **gecko-t/t-linux-docker-noscratch-amd** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 94775 | 94775 |
| **gecko-t/t-linux-docker-noscratch-amd-alpha** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 12 | 12 |
| **gecko-t/t-linux-xlarge-2204-wayland** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | amd64 | 1.22.2 | 10981 | 10981 |
| **gecko-t/t-linux-xlarge-2404-wayland** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 2 | 2 |
| **gecko-t/win10-64-2009** | generic-worker | 88.1.3 | multiuser | ffc99d77c0 | windows | amd64 | 1.25.1 | 4487 | 4487 |
| **gecko-t/win10-64-2009-alpha** | generic-worker | 91.1.0 | multiuser | b8f34332c3 | windows | amd64 | 1.25.3 | 2 | 2 |
| **gecko-t/win10-64-2009-gpu** | generic-worker | 88.1.3 | multiuser | ffc99d77c0 | windows | amd64 | 1.25.1 | 468 | 468 |
| **gecko-t/win10-64-2009-gpu-alpha** | generic-worker | 91.1.0 | multiuser | b8f34332c3 | windows | amd64 | 1.25.3 | 2 | 2 |
| **gecko-t/win10-64-2009-source** | generic-worker | 88.1.3 | multiuser | ffc99d77c0 | windows | amd64 | 1.25.1 | 47 | 47 |
| **gecko-t/win10-64-2009-source-alpha** | generic-worker | 91.1.0 | multiuser | b8f34332c3 | windows | amd64 | 1.25.3 | 3 | 3 |
| **gecko-t/win10-64-2009-webgpu** | generic-worker | 88.1.3 | multiuser | ffc99d77c0 | windows | amd64 | 1.25.1 | 2 | 2 |
| **gecko-t/win10-64-2009-webgpu-alpha** | generic-worker | 91.1.0 | multiuser | b8f34332c3 | windows | amd64 | 1.25.3 | 3 | 3 |
| **gecko-t/win11-64-2009** | generic-worker | 88.1.3 | multiuser | ffc99d77c0 | windows | amd64 | 1.25.1 | 451 | 451 |
| **gecko-t/win11-64-2009-alpha** | generic-worker | 93.1.4 | multiuser | 1088d85c3b | windows | amd64 | 1.25.4 | 2 | 2 |
| **gecko-t/win11-64-2009-gpu** | generic-worker | 88.1.3 | multiuser | ffc99d77c0 | windows | amd64 | 1.25.1 | 293 | 293 |
| **gecko-t/win11-64-2009-gpu-alpha** | generic-worker | 93.1.4 | multiuser | 1088d85c3b | windows | amd64 | 1.25.4 | 2 | 2 |
| **gecko-t/win11-64-2009-source** | generic-worker | 88.1.3 | multiuser | ffc99d77c0 | windows | amd64 | 1.25.1 | 2 | 2 |
| **gecko-t/win11-64-2009-source-alpha** | generic-worker | 93.1.4 | multiuser | 1088d85c3b | windows | amd64 | 1.25.4 | 2 | 2 |
| **gecko-t/win11-64-2009-ssd** | generic-worker | 88.1.3 | multiuser | ffc99d77c0 | windows | amd64 | 1.25.1 | 2 | 2 |
| **gecko-t/win11-64-2009-ssd-gpu** | generic-worker | 88.1.3 | multiuser | ffc99d77c0 | windows | amd64 | 1.25.1 | 2 | 2 |
| **gecko-t/win11-64-2009-webgpu** | generic-worker | 88.1.3 | multiuser | ffc99d77c0 | windows | amd64 | 1.25.1 | 2 | 2 |
| **gecko-t/win11-64-2009-webgpu-alpha** | generic-worker | 93.1.4 | multiuser | 1088d85c3b | windows | amd64 | 1.25.4 | 4 | 4 |
| **gecko-t/win11-64-24h2** | generic-worker | 88.1.3 | multiuser | ffc99d77c0 | windows | amd64 | 1.25.1 | 48121 | 48121 |
| **gecko-t/win11-64-24h2-alpha** | generic-worker | 95.1.1 | multiuser | cbb4697b6c | windows | amd64 | 1.25.5 | 25 | 25 |
| **gecko-t/win11-64-24h2-amd** | generic-worker | 94.1.0 | multiuser | 44d8d7264b | windows | amd64 | 1.25.5 | 3681 | 3681 |
| **gecko-t/win11-64-24h2-gpu** | generic-worker | 88.1.3 | multiuser | ffc99d77c0 | windows | amd64 | 1.25.1 | 17168 | 17168 |
| **gecko-t/win11-64-24h2-gpu-alpha** | generic-worker | 95.1.1 | multiuser | cbb4697b6c | windows | amd64 | 1.25.5 | 2 | 2 |
| **gecko-t/win11-64-24h2-large** | generic-worker | 88.1.3 | multiuser | ffc99d77c0 | windows | amd64 | 1.25.1 | 1855 | 1855 |
| **gecko-t/win11-64-24h2-large-alpha** | generic-worker | 95.1.1 | multiuser | cbb4697b6c | windows | amd64 | 1.25.5 | 2 | 2 |
| **gecko-t/win11-64-24h2-source** | generic-worker | 88.1.3 | multiuser | ffc99d77c0 | windows | amd64 | 1.25.1 | 2955 | 2955 |
| **gecko-t/win11-64-24h2-source-alpha** | generic-worker | 95.1.1 | multiuser | cbb4697b6c | windows | amd64 | 1.25.5 | 2 | 2 |
| **gecko-t/win11-64-24h2-ssd** | generic-worker | 88.1.3 | multiuser | ffc99d77c0 | windows | amd64 | 1.25.1 | 2 | 2 |
| **gecko-t/win11-64-24h2-ssd-gpu** | generic-worker | 88.1.3 | multiuser | ffc99d77c0 | windows | amd64 | 1.25.1 | 2 | 2 |
| **gecko-t/win11-64-24h2-webgpu** | generic-worker | 88.1.3 | multiuser | ffc99d77c0 | windows | amd64 | 1.25.1 | 3662 | 3662 |
| **gecko-t/win11-64-24h2-webgpu-alpha** | generic-worker | 95.1.1 | multiuser | cbb4697b6c | windows | amd64 | 1.25.5 | 2 | 2 |
| **gecko-t/win11-a64-24h2** | generic-worker | 88.1.3 | multiuser | ffc99d77c0 | windows | amd64 | 1.25.1 | 186 | 186 |
| **gecko-t/win11-a64-24h2-alpha** | generic-worker | 93.1.4 | multiuser | 1088d85c3b | windows | arm64 | 1.25.4 | 2 | 2 |
| **glean-1/b-linux** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 41 | 41 |
| **glean-1/decision** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 1073 | 1073 |
| **glean-1/images** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 2 | 2 |
| **glean-3/b-linux** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 35 | 35 |
| **glean-3/decision** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 10 | 10 |
| **glean-3/images** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 2 | 2 |
| **mobile-1/b-linux** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 162 | 162 |
| **mobile-1/decision** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 1048 | 1048 |
| **mobile-1/images** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 2 | 2 |
| **mobile-3/b-linux** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 212 | 212 |
| **mobile-3/decision** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 116 | 116 |
| **mobile-3/images** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 2 | 2 |
| **mozilla-1/b-linux** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 2 | 2 |
| **mozilla-1/b-win2022** | generic-worker | 88.1.3 | multiuser | ffc99d77c0 | windows | amd64 | 1.25.1 | 2 | 2 |
| **mozilla-1/decision** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 1063 | 1063 |
| **mozilla-1/images** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 6 | 6 |
| **mozilla-3/b-linux** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 2 | 2 |
| **mozilla-3/b-win2022** | generic-worker | 88.1.3 | multiuser | ffc99d77c0 | windows | amd64 | 1.25.1 | 2 | 2 |
| **mozilla-3/decision** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 428 | 428 |
| **mozilla-3/images** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 2 | 2 |
| **mozilla-t/t-linux-2204-wayland** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | amd64 | 1.22.2 | 558 | 558 |
| **mozilla-t/t-linux-2204-wayland-relsre** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | amd64 | 1.22.2 | 2 | 2 |
| **mozilla-t/t-linux-2404-wayland** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 2 | 2 |
| **mozilla-t/t-linux-arm64-docker** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | arm64 | 1.25.5 | 2 | 2 |
| **mozilla-t/t-linux-docker** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 12 | 12 |
| **mozilla-t/t-linux-docker-amd** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 2 | 2 |
| **mozilla-t/t-linux-docker-noscratch** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 4 | 4 |
| **mozilla-t/t-linux-docker-noscratch-amd** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 2 | 2 |
| **mozillavpn-1/b-linux** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 450 | 450 |
| **mozillavpn-1/b-linux-gcp-gw** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | amd64 | 1.22.2 | 2 | 2 |
| **mozillavpn-1/b-linux-large** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 331 | 331 |
| **mozillavpn-1/b-win2022** | generic-worker | 88.1.3 | multiuser | ffc99d77c0 | windows | amd64 | 1.25.1 | 137 | 137 |
| **mozillavpn-1/b-win2022-alpha** | generic-worker | 93.1.2 | multiuser | e954a7555f | windows | amd64 | 1.25.4 | 2 | 2 |
| **mozillavpn-1/decision** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 1058 | 1058 |
| **mozillavpn-1/images** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 16 | 16 |
| **mozillavpn-1/win11-a64-24h2-builder** | generic-worker | 88.1.3 | multiuser | ffc99d77c0 | windows | amd64 | 1.25.1 | 10 | 10 |
| **mozillavpn-3/b-linux** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 205 | 205 |
| **mozillavpn-3/b-linux-gcp-gw** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | amd64 | 1.22.2 | 2 | 2 |
| **mozillavpn-3/b-linux-large** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 205 | 205 |
| **mozillavpn-3/b-win2022** | generic-worker | 88.1.3 | multiuser | ffc99d77c0 | windows | amd64 | 1.25.1 | 66 | 66 |
| **mozillavpn-3/decision** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 28 | 28 |
| **mozillavpn-3/images** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 5 | 5 |
| **mozillavpn-3/win11-a64-24h2-builder** | generic-worker | 88.1.3 | multiuser | ffc99d77c0 | windows | amd64 | 1.25.1 | 4 | 4 |
| **nss-1/b-win2022** | generic-worker | 88.1.3 | multiuser | ffc99d77c0 | windows | amd64 | 1.25.1 | 142 | 142 |
| **nss-1/b-win2022-alpha** | generic-worker | 93.1.2 | multiuser | e954a7555f | windows | amd64 | 1.25.4 | 2 | 2 |
| **nss-1/decision** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 7 | 7 |
| **nss-1/images** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 2 | 2 |
| **nss-1/linux-docker** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 337 | 337 |
| **nss-1/win11-a64-24h2** | generic-worker | 88.1.3 | multiuser | ffc99d77c0 | windows | amd64 | 1.25.1 | 2 | 2 |
| **nss-1/win11-a64-24h2-alpha** | generic-worker | 93.1.4 | multiuser | 1088d85c3b | windows | arm64 | 1.25.4 | 2 | 2 |
| **nss-1/win11-a64-24h2-builder** | generic-worker | 88.1.3 | multiuser | ffc99d77c0 | windows | amd64 | 1.25.1 | 2 | 2 |
| **nss-1/win11-a64-24h2-builder-alpha** | generic-worker | 93.1.4 | multiuser | 1088d85c3b | windows | arm64 | 1.25.4 | 2 | 2 |
| **nss-3/b-win2022** | generic-worker | 88.1.3 | multiuser | ffc99d77c0 | windows | amd64 | 1.25.1 | 272 | 272 |
| **nss-3/decision** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 11 | 11 |
| **nss-3/images** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 2 | 2 |
| **nss-3/linux-docker** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 580 | 580 |
| **nss-3/win11-a64-24h2-builder** | generic-worker | 88.1.3 | multiuser | ffc99d77c0 | windows | amd64 | 1.25.1 | 2 | 2 |
| **nss-t/t-linux-arm64-docker** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | arm64 | 1.25.5 | 2 | 2 |
| **nss-t/t-linux-docker** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 936 | 936 |
| **nss-t/t-linux-docker-amd** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 2 | 2 |
| **nss-t/win11-a64-24h2** | generic-worker | 88.1.3 | multiuser | ffc99d77c0 | windows | amd64 | 1.25.1 | 2 | 2 |
| **nss-t/win11-a64-24h2-alpha** | generic-worker | 93.1.4 | multiuser | 1088d85c3b | windows | arm64 | 1.25.4 | 3 | 3 |
| **proj-autophone/gecko-t-bitbar-gw-perf-a55** | generic-worker | 36.0.0 | simple | b9cb11293f | linux | amd64 | 1.13.7 | 0 | 0 |
| **proj-autophone/gecko-t-bitbar-gw-test-1** | generic-worker | 36.0.0 | simple | b9cb11293f | linux | amd64 | 1.13.7 | 0 | 0 |
| **proj-autophone/gecko-t-bitbar-gw-unit-p5** | generic-worker | 36.0.0 | simple | b9cb11293f | linux | amd64 | 1.13.7 | 0 | 0 |
| **proj-autophone/gecko-t-lambda-perf-a55** | generic-worker | 87.0.0 | insecure | 99a1fcafbb | linux | amd64 | 1.24.5 | 0 | 0 |
| **releng-1/decision** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 44 | 44 |
| **releng-1/linux-docker** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 33 | 33 |
| **releng-3/decision** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 30 | 30 |
| **releng-3/linux-docker** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 39 | 39 |
| **releng-hardware/applicationservices-b-1-osx1015** | generic-worker | 60.3.4 | multiuser | 943a6f2b0d | darwin | amd64 | 1.22.0 | 0 | 0 |
| **releng-hardware/applicationservices-b-3-osx1015** | generic-worker | 60.3.4 | multiuser | 943a6f2b0d | darwin | amd64 | 1.22.0 | 0 | 0 |
| **releng-hardware/enterprise-1-b-osx-arm64** | generic-worker | 60.3.4 | multiuser | 943a6f2b0d | darwin | arm64 | 1.22.0 | 0 | 0 |
| **releng-hardware/enterprise-3-b-osx-arm64** | generic-worker | 60.3.4 | multiuser | 943a6f2b0d | darwin | arm64 | 1.22.0 | 0 | 0 |
| **releng-hardware/gecko-1-b-osx-1015-staging** | generic-worker | 60.3.4 | multiuser | 943a6f2b0d | darwin | amd64 | 1.22.0 | 0 | 0 |
| **releng-hardware/gecko-1-b-osx-arm64** | generic-worker | 60.3.4 | multiuser | 943a6f2b0d | darwin | arm64 | 1.22.0 | 0 | 0 |
| **releng-hardware/gecko-3-b-osx-1015** | generic-worker | 60.3.4 | multiuser | 943a6f2b0d | darwin | amd64 | 1.22.0 | 0 | 0 |
| **releng-hardware/gecko-3-b-osx-arm64** | generic-worker | 60.3.4 | multiuser | 943a6f2b0d | darwin | arm64 | 1.22.0 | 0 | 0 |
| **releng-hardware/gecko-t-linux-netperf-2404** | generic-worker | 88.0.2 | insecure | fcaf4a25fc | linux | amd64 | 1.24.5 | 0 | 0 |
| **releng-hardware/gecko-t-linux-talos-1804** | generic-worker | 61.0.0 | simple | 3bd4419b4b | linux | amd64 | 1.22.1 | 0 | 0 |
| **releng-hardware/gecko-t-linux-talos-2404** | generic-worker | 88.0.2 | insecure | fcaf4a25fc | linux | amd64 | 1.24.5 | 0 | 0 |
| **releng-hardware/gecko-t-osx-1015-r8** | generic-worker | 60.3.4 | simple | 943a6f2b0d | darwin | amd64 | 1.22.0 | 0 | 0 |
| **releng-hardware/gecko-t-osx-1015-r8-staging** | generic-worker | 72.0.1 | insecure | fa5416dc69 | darwin | amd64 | 1.23.1 | 0 | 0 |
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
| **releng-hardware/win11-64-24h2-hw-ref** | generic-worker | 88.1.3 | multiuser | ffc99d77c0 | windows | amd64 | 1.25.1 | 0 | 0 |
| **releng-hardware/win11-64-24h2-hw-ref-alpha** | generic-worker | 93.1.4 | multiuser | 1088d85c3b | windows | amd64 | 1.25.4 | 0 | 0 |
| **releng-hardware/win11-64-24h2-hw-relops1213** | generic-worker | 93.1.4 | multiuser | 1088d85c3b | windows | amd64 | 1.25.4 | 0 | 0 |
| **releng-t/linux-docker** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 217 | 217 |
| **relops-1/decision** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 2 | 2 |
| **relops-1/images** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 2 | 2 |
| **relops-3/decision** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 8 | 8 |
| **relops-3/images** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 2 | 2 |
| **scriptworker-1/decision** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 4 | 4 |
| **scriptworker-1/images** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 17 | 17 |
| **scriptworker-3/decision** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 5 | 5 |
| **scriptworker-3/images** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 22 | 22 |
| **taskgraph-1/decision** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 24 | 24 |
| **taskgraph-1/images** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 19 | 19 |
| **taskgraph-3/decision** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 9 | 9 |
| **taskgraph-3/images** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 10 | 10 |
| **taskgraph-t/linux-docker** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 152 | 152 |
| **taskgraph-t/win11-64-24h2** | generic-worker | 88.1.3 | multiuser | ffc99d77c0 | windows | amd64 | 1.25.1 | 8 | 8 |
| **translations-1/b-linux-large-gcp-1tb-32-256-d2g** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 20 | 20 |
| **translations-1/b-linux-large-gcp-1tb-32-256-std-d2g** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 2 | 2 |
| **translations-1/b-linux-large-gcp-1tb-64-512-d2g** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 2 | 2 |
| **translations-1/b-linux-large-gcp-1tb-64-512-std-d2g** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 2 | 2 |
| **translations-1/b-linux-large-gcp-d2g** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 231 | 231 |
| **translations-1/b-linux-large-gcp-d2g-1tb** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 2 | 2 |
| **translations-1/b-linux-large-gcp-d2g-1tb-standard** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 2 | 2 |
| **translations-1/b-linux-large-gcp-d2g-300gb** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 104 | 104 |
| **translations-1/b-linux-v100-gpu-d2g** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 89 | 89 |
| **translations-1/b-linux-v100-gpu-d2g-4** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 17 | 17 |
| **translations-1/b-linux-v100-gpu-d2g-4-1tb** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 10 | 10 |
| **translations-1/b-linux-v100-gpu-d2g-4-1tb-alpha** | generic-worker | 84.1.2 | multiuser | 4a139f395a | linux | amd64 | 1.24.4 | 2 | 2 |
| **translations-1/b-linux-v100-gpu-d2g-4-1tb-standard** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 3 | 3 |
| **translations-1/b-linux-v100-gpu-d2g-4-1tb-std-alpha** | generic-worker | 84.1.2 | multiuser | 4a139f395a | linux | amd64 | 1.24.4 | 2 | 2 |
| **translations-1/b-linux-v100-gpu-d2g-4-2tb** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 10 | 10 |
| **translations-1/b-linux-v100-gpu-d2g-4-2tb-alpha** | generic-worker | 84.1.2 | multiuser | 4a139f395a | linux | amd64 | 1.24.4 | 2 | 2 |
| **translations-1/b-linux-v100-gpu-d2g-4-300gb** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 25 | 25 |
| **translations-1/b-linux-v100-gpu-d2g-4-300gb-alpha** | generic-worker | 84.1.2 | multiuser | 4a139f395a | linux | amd64 | 1.24.4 | 2 | 2 |
| **translations-1/b-linux-v100-gpu-d2g-4-300gb-standard** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 2 | 2 |
| **translations-1/b-linux-v100-gpu-d2g-4-300gb-std-alpha** | generic-worker | 84.1.2 | multiuser | 4a139f395a | linux | amd64 | 1.24.4 | 2 | 2 |
| **translations-1/b-linux-v100-gpu-d2g-4-alpha** | generic-worker | 84.1.2 | multiuser | 4a139f395a | linux | amd64 | 1.24.4 | 2 | 2 |
| **translations-1/b-linux-v100-gpu-d2g-alpha** | generic-worker | 84.1.2 | multiuser | 4a139f395a | linux | amd64 | 1.24.4 | 2 | 2 |
| **translations-1/decision** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 1073 | 1073 |
| **translations-1/images** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 6 | 6 |
| **translations-t/t-linux-gw-noscratch-amd** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 1 | 1 |
| **xpi-1/b-linux** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 111 | 111 |
| **xpi-1/decision** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 45 | 45 |
| **xpi-1/images** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 2 | 2 |
| **xpi-3/b-linux** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 3 | 3 |
| **xpi-3/decision** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 4 | 4 |
| **xpi-3/images** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 2 | 2 |


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
| projects/taskcluster-imaging/global/images/docker-firefoxci-gcp-lt-googlecompute-2023-04-13t21-30-28z | 1 |
| projects/fxci-production-level3-workers/global/images/docker-firefoxci-gcp-l3-googlecompute-2024-02-05t23-18-22z | 52 |
| projects/taskcluster-imaging/global/images/docker-firefoxci-gcp-l1-googlecompute-2025-06-13t18-31-38z | 85 |
| projects/taskcluster-imaging/global/images/docker-worker-gcp-u14-04-2025-06-16 | 19 |


| Worker Pool | Implementation | Version | Total Workers | Total Capacity |
| --- | --- | --- | ---: | ---: |
| **app-services-1/b-linux-amd** | docker-worker | 38.0.5 | 2 | 2 |
| **app-services-1/b-linux-gcp** | docker-worker | 38.0.5 | 127 | 127 |
| **app-services-1/decision-gcp** | docker-worker | 38.0.5 | 808 | 1616 |
| **app-services-1/images-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **app-services-3/b-linux-amd** | docker-worker | 38.0.5 | 2 | 2 |
| **app-services-3/b-linux-gcp** | docker-worker | 38.0.5 | 3 | 3 |
| **app-services-3/decision-gcp** | docker-worker | 38.0.5 | 2 | 4 |
| **app-services-3/images-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **code-analysis-1/linux-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **code-analysis-3/linux-gcp** | docker-worker | 38.0.5 | 3 | 3 |
| **comm-1/b-linux-amd** | docker-worker | 38.0.5 | 2 | 2 |
| **comm-1/b-linux-gcp** | docker-worker | 38.0.5 | 518 | 518 |
| **comm-1/b-linux-large-amd** | docker-worker | 38.0.5 | 2 | 2 |
| **comm-1/b-linux-large-gcp** | docker-worker | 38.0.5 | 5 | 5 |
| **comm-1/b-linux-xlarge-amd** | docker-worker | 38.0.5 | 2 | 2 |
| **comm-1/b-linux-xlarge-gcp** | docker-worker | 38.0.5 | 5 | 5 |
| **comm-1/decision-gcp** | docker-worker | 38.0.5 | 210 | 210 |
| **comm-1/images-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **comm-2/b-linux-amd** | docker-worker | 38.0.5 | 2 | 2 |
| **comm-2/b-linux-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **comm-2/b-linux-large-amd** | docker-worker | 38.0.5 | 2 | 2 |
| **comm-2/b-linux-large-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **comm-2/b-linux-xlarge-amd** | docker-worker | 38.0.5 | 2 | 2 |
| **comm-2/b-linux-xlarge-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **comm-2/decision-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **comm-2/images-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **comm-3/b-linux-amd** | docker-worker | 38.0.5 | 2 | 2 |
| **comm-3/b-linux-gcp** | docker-worker | 38.0.5 | 3888 | 3888 |
| **comm-3/b-linux-large-amd** | docker-worker | 38.0.5 | 2 | 2 |
| **comm-3/b-linux-large-gcp** | docker-worker | 38.0.5 | 5 | 5 |
| **comm-3/b-linux-xlarge-amd** | docker-worker | 38.0.5 | 3 | 3 |
| **comm-3/b-linux-xlarge-gcp** | docker-worker | 38.0.5 | 6 | 6 |
| **comm-3/decision-gcp** | docker-worker | 38.0.5 | 113 | 113 |
| **comm-3/images-gcp** | docker-worker | 38.0.5 | 9 | 9 |
| **comm-t/misc-gcp** | docker-worker | 38.0.5 | 2 | 16 |
| **comm-t/t-linux-large-gcp** | docker-worker | unknown version | 2 | 2 |
| **comm-t/t-linux-large-noscratch-gcp** | docker-worker | unknown version | 2 | 2 |
| **comm-t/t-linux-xlarge-gcp** | docker-worker | unknown version | 2 | 2 |
| **comm-t/t-linux-xlarge-noscratch-gcp** | docker-worker | unknown version | 2 | 2 |
| **comm-t/t-linux-xlarge-source-gcp** | docker-worker | unknown version | 2 | 2 |
| **comm-t/t-linux-xlarge-source-noscratch-gcp** | docker-worker | unknown version | 2 | 2 |
| **gecko-1/b-linux-amd** | docker-worker | 38.0.5 | 20 | 20 |
| **gecko-1/b-linux-gcp** | docker-worker | 38.0.5 | 6 | 6 |
| **gecko-1/b-linux-gcp-relops1411** | docker-worker | 38.0.5 | 2 | 2 |
| **gecko-1/b-linux-gcp-test-bug-1882320** | docker-worker | 38.0.5 | 2 | 2 |
| **gecko-1/b-linux-kvm-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **gecko-1/b-linux-large-amd** | docker-worker | 38.0.5 | 2 | 2 |
| **gecko-1/b-linux-large-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **gecko-1/b-linux-medium-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **gecko-1/b-linux-xlarge-amd** | docker-worker | 38.0.5 | 2 | 2 |
| **gecko-1/b-linux-xlarge-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **gecko-1/b-linux-xlarge-gcp-bug1797804-c2** | docker-worker | 38.0.5 | 3 | 3 |
| **gecko-1/decision-gcp** | docker-worker | 38.0.5 | 731 | 731 |
| **gecko-1/decision-gcp-c3d4-amd-bug1990935** | docker-worker | 38.0.5 | 2 | 2 |
| **gecko-1/decision-gcp-c3d8-amd-bug1990935** | docker-worker | 38.0.5 | 2 | 2 |
| **gecko-1/decision-gcp-c4d4-amd-bug1990935** | docker-worker | 38.0.5 | 2 | 2 |
| **gecko-1/decision-gcp-c4d4-hcpu-amd-bug1990935** | docker-worker | 38.0.5 | 2 | 2 |
| **gecko-1/decision-gcp-c4d8-amd-bug1990935** | docker-worker | 38.0.5 | 2 | 2 |
| **gecko-1/decision-gcp-c4d8-hcpu-amd-bug1990935** | docker-worker | 38.0.5 | 2 | 2 |
| **gecko-1/decision-gcp-c4d8-lssd-amd-bug1990935** | docker-worker | 38.0.5 | 2 | 2 |
| **gecko-1/images-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **gecko-2/b-linux-amd** | docker-worker | 38.0.5 | 2 | 2 |
| **gecko-2/b-linux-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **gecko-2/b-linux-kvm-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **gecko-2/b-linux-large-amd** | docker-worker | 38.0.5 | 2 | 2 |
| **gecko-2/b-linux-large-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **gecko-2/b-linux-medium-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **gecko-2/b-linux-xlarge-amd** | docker-worker | 38.0.5 | 2 | 2 |
| **gecko-2/b-linux-xlarge-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **gecko-2/decision-gcp** | docker-worker | 38.0.5 | 10 | 10 |
| **gecko-2/images-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **gecko-3/b-linux-amd** | docker-worker | 38.0.5 | 221 | 221 |
| **gecko-3/b-linux-gcp** | docker-worker | 38.0.5 | 207 | 207 |
| **gecko-3/b-linux-kvm-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **gecko-3/b-linux-large-amd** | docker-worker | 38.0.5 | 2 | 2 |
| **gecko-3/b-linux-large-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **gecko-3/b-linux-medium-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **gecko-3/b-linux-xlarge-amd** | docker-worker | 38.0.5 | 2 | 2 |
| **gecko-3/b-linux-xlarge-gcp** | docker-worker | 38.0.5 | 8 | 8 |
| **gecko-3/decision-gcp** | docker-worker | 38.0.5 | 795 | 795 |
| **gecko-3/images-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **gecko-t/misc-gcp** | docker-worker | 38.0.5 | 8 | 64 |
| **gecko-t/t-linux-kvm-gcp** | docker-worker | unknown version | 2 | 2 |
| **gecko-t/t-linux-kvm-gcp-relops1411** | docker-worker | unknown version | 2 | 2 |
| **gecko-t/t-linux-kvm-noscratch-gcp** | docker-worker | unknown version | 4 | 4 |
| **gecko-t/t-linux-large-gcp** | docker-worker | unknown version | 2 | 2 |
| **gecko-t/t-linux-large-hostub2204-gcp** | docker-worker | 48.3.0 | 2 | 2 |
| **gecko-t/t-linux-large-noscratch-gcp** | docker-worker | unknown version | 707 | 707 |
| **gecko-t/t-linux-xlarge-gcp** | docker-worker | unknown version | 13 | 13 |
| **gecko-t/t-linux-xlarge-noscratch-gcp** | docker-worker | unknown version | 492 | 492 |
| **gecko-t/t-linux-xlarge-source-gcp** | docker-worker | unknown version | 70 | 70 |
| **gecko-t/t-linux-xlarge-source-noscratch-gcp** | docker-worker | unknown version | 2 | 2 |
| **glean-1/b-linux-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **glean-1/decision-gcp** | docker-worker | 38.0.5 | 793 | 1586 |
| **glean-1/images-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **glean-3/b-linux-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **glean-3/decision-gcp** | docker-worker | 38.0.5 | 2 | 4 |
| **glean-3/images-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **infra/build-decision** | docker-worker | 38.0.5 | 1 | 8 |
| **mobile-1/b-linux-gcp** | docker-worker | 38.0.5 | 62 | 62 |
| **mobile-1/decision-gcp** | docker-worker | 38.0.5 | 765 | 1530 |
| **mobile-1/images-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **mobile-3/b-linux-gcp** | docker-worker | 38.0.5 | 8 | 8 |
| **mobile-3/decision-gcp** | docker-worker | 38.0.5 | 9 | 18 |
| **mobile-3/images-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **mozilla-1/b-linux-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **mozilla-1/decision-gcp** | docker-worker | 38.0.5 | 789 | 1578 |
| **mozilla-1/images-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **mozilla-3/b-linux-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **mozilla-3/decision-gcp** | docker-worker | 38.0.5 | 2 | 4 |
| **mozilla-3/images-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **mozilla-t/t-linux-large-gcp** | docker-worker | unknown version | 2 | 2 |
| **mozilla-t/t-linux-large-noscratch-gcp** | docker-worker | unknown version | 3 | 3 |
| **mozillavpn-1/b-linux-gcp** | docker-worker | 38.0.5 | 48 | 48 |
| **mozillavpn-1/b-linux-large-gcp** | docker-worker | 38.0.5 | 32 | 32 |
| **mozillavpn-1/decision-gcp** | docker-worker | 38.0.5 | 790 | 1580 |
| **mozillavpn-1/images-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **mozillavpn-3/b-linux-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **mozillavpn-3/b-linux-large-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **mozillavpn-3/decision-gcp** | docker-worker | 38.0.5 | 2 | 4 |
| **mozillavpn-3/images-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **nss-1/decision-gcp** | docker-worker | 38.0.5 | 2 | 4 |
| **nss-1/images-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **nss-1/linux-gcp** | docker-worker | 38.0.5 | 2 | 2 |
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
| **taskgraph-1/decision-gcp** | docker-worker | 38.0.5 | 2 | 4 |
| **taskgraph-1/images-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **taskgraph-3/decision-gcp** | docker-worker | 38.0.5 | 2 | 4 |
| **taskgraph-3/images-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **translations-1/decision-gcp** | docker-worker | 38.0.5 | 1789 | 3578 |
| **translations-1/images-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **xpi-1/b-linux-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **xpi-1/decision-gcp** | docker-worker | 38.0.5 | 19 | 38 |
| **xpi-1/images-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **xpi-3/b-linux-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **xpi-3/decision-gcp** | docker-worker | 38.0.5 | 2 | 4 |
| **xpi-3/images-gcp** | docker-worker | 38.0.5 | 2 | 2 |


## Script Worker

Total: `41`



| Worker Pool | Implementation | Version | Total Workers | Total Capacity |
| --- | --- | --- | ---: | ---: |
| **scriptworker-k8s/app-services-3-beetmover** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/app-services-3-signing** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/comm-3-balrog** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/comm-3-beetmover** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/comm-3-bouncer** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/comm-3-shipit** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/comm-3-signing** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/comm-3-tree** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/comm-t-signing** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/enterprise-3-signing** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/enterprise-t-signing** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/gecko-1-addon** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/gecko-1-balrog** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/gecko-1-beetmover** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/gecko-1-bouncer** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/gecko-1-lando** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
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
| **scriptworker-k8s/xpi-t-signing** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
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
| **gecko-t/t-linux-vm-2204-wayland** |  | No artifacts found | 34 | 34 |
| **gecko-t/t-linux-vm-2204-wayland-snap** |  | No artifacts found | 3 | 3 |


## Version not determined [^2]

Total: `7`


Count by image:

| Version | Count |
| :--- | ---: |
| projects/taskcluster-imaging/global/images/gw-fxci-gcp-l1-2404-amd64-googlecompute-alpha | 1 |
|  | 6 |


| Worker Pool | Implementation | Version | Total Workers | Total Capacity |
| --- | --- | --- | ---: | ---: |
| **gecko-t/t-linux-2404-relsre** |  | Version not determined; task not (yet) claimed | 4 | 4 |
| **proj-autophone/gecko-t-bitbar-gw-perf-p6** |  | Version not determined; task not (yet) claimed | 0 | 0 |
| **proj-autophone/gecko-t-bitbar-gw-perf-s24** |  | Version not determined; task not (yet) claimed | 0 | 0 |
| **releng-hardware/gecko-1-b-osx-1015** |  | Version not determined; task not (yet) claimed | 0 | 0 |
| **releng-hardware/gecko-t-linux-netperf-1804** |  | Version not determined; task not (yet) claimed | 0 | 0 |
| **releng-hardware/gecko-t-osx-1400-r8** |  | Version not determined; task not (yet) claimed | 0 | 0 |
| **scriptworker-prov-v1/adhoc-signing-mac14m2** |  | Version not determined; task not (yet) claimed | 0 | 0 |



[^1]: Those are the pools whose tasks were claimed and resolved by a worker as expected, but the worker did not publish either artifact `public/logs/live_backing.log` nor `public/logs/chain_of_trust.log`, which is the source used to identify the worker implementation.

[^2]: Probing task remains pending after two hours. Those are the pools that were not able to start any worker to claim the task within two hours.

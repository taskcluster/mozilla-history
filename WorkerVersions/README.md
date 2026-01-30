

# Worker Pool Versions


## Generic Worker

Total: `349`

Count by version:

| Version | Count |
| :--- | ---: |
| 36.0.0 | 2 |
| 45.0.0 | 1 |
| 60.3.4 | 18 |
| 61.0.0 | 1 |
| 64.3.0 | 26 |
| 65.1.0 | 1 |
| 84.1.2 | 7 |
| 86.0.2 | 3 |
| 87.0.0 | 1 |
| 88.0.2 | 2 |
| 88.1.3 | 12 |
| 91.0.2 | 1 |
| 91.1.0 | 42 |
| 91.1.1 | 2 |
| 93.1.4 | 10 |
| 94.1.0 | 1 |
| 95.0.3 | 1 |
| 95.1.3 | 211 |
| 96.1.0 | 7 |


Count by image:

| Version | Count |
| :--- | ---: |
| /subscriptions/108d46d5-fe9b-4850-9a7d-8c914aa6c1f0/resourceGroups/rg-packer-worker-images/providers/Microsoft.Compute/galleries/win11_a64_24h2_tester_alpha/images/win11_a64_24h2_tester_alpha/versions/1.0.0 | 4 |
| /subscriptions/a30e97ab-734a-4f3b-a0e4-c51c0bff0701/resourceGroups/rg-packer-worker-images/providers/Microsoft.Compute/galleries/trusted_win11_a64_24h2_builder/images/trusted_win11_a64_24h2_builder/versions/1.0.9 | 3 |
| projects/taskcluster-imaging/global/images/gw-fxci-gcp-l1-arm64-gui-googlecompute-2024-09-18t14-57-52z | 9 |
| projects/taskcluster-imaging/global/images/gw-fxci-gcp-l1-2404-arm64-headless-googlecompute-alpha | 1 |
| projects/taskcluster-imaging/global/images/gw-fxci-gcp-l1-gui-googlecompute-2025-01-13t22-33-40z | 2 |
| /subscriptions/108d46d5-fe9b-4850-9a7d-8c914aa6c1f0/resourceGroups/rg-packer-worker-images/providers/Microsoft.Compute/galleries/win11_a64_24h2_tester/images/win11_a64_24h2_tester/versions/1.0.9 | 4 |
| /subscriptions/108d46d5-fe9b-4850-9a7d-8c914aa6c1f0/resourceGroups/rg-packer-worker-images/providers/Microsoft.Compute/galleries/win11_a64_24h2_builder_alpha/images/win11_a64_24h2_builder_alpha/versions/1.0.0 | 2 |
| /subscriptions/108d46d5-fe9b-4850-9a7d-8c914aa6c1f0/resourceGroups/rg-packer-worker-images/providers/Microsoft.Compute/galleries/win11_a64_24h2_builder/images/win11_a64_24h2_builder/versions/1.0.9 | 3 |
| projects/taskcluster-imaging/global/images/gw-fxci-gcp-l1-2404-amd64-headless-googlecompute-alpha | 6 |
|  | 36 |
| projects/taskcluster-imaging/global/images/gw-fxci-gcp-l1-2404-amd64-gui-googlecompute-2026-01-14 | 4 |
| projects/taskcluster-imaging/global/images/gw-fxci-gcp-l1-2404-amd64-headless-googlecompute-2026-01-14 | 120 |
| projects/fxci-production-level3-workers/global/images/gw-fxci-gcp-l3-2404-amd64-headless-googlecompute-2026-01-14 | 58 |
| projects/taskcluster-imaging/global/images/gw-fxci-gcp-l1-2404-arm64-headless-googlecompute-2026-01-14 | 14 |
| projects/fxci-production-level3-workers/global/images/gw-fxci-gcp-l3-gui-googlecompute-2024-09-18t05-46-31z | 3 |
| projects/fxci-production-level3-workers/global/images/gw-fxci-gcp-l3-2404-arm64-headless-googlecompute-2026-01-14 | 6 |
| projects/taskcluster-imaging/global/images/gw-fxci-gcp-l1-2404-amd64-gui-googlecompute-alpha | 1 |
| projects/fxci-production-level3-workers/global/images/gw-fxci-gcp-l3-arm64-gui-googlecompute-2024-09-18t19-02-58z | 4 |
| projects/fxci-production-level3-workers/global/images/gw-fxci-gcp-l3-2404-amd64-headless-googlecompute-alpha | 1 |
| projects/taskcluster-imaging/global/images/gw-fxci-gcp-l1-gui-googlecompute-2024-08-22t22-48-09z | 8 |
| unknown | 53 |
| projects/taskcluster-imaging/global/images/gw-fxci-gcp-l1-2404-amd64-headless-googlecompute-alpha-tc | 7 |


| Worker Pool | Implementation | Version | Engine | Revision | OS | Arch | GO | Total Workers | Total Capacity |
| --- | --- | --- | --- | --- | --- | --- | --- | ---: | ---: |
| **adhoc-1/decision** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 4 | 4 |
| **adhoc-1/images** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 6 | 6 |
| **adhoc-3/decision** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 3 | 3 |
| **adhoc-3/images** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 2 | 2 |
| **app-services-1/b-linux** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 369 | 369 |
| **app-services-1/b-linux-docker-amd** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 2 | 2 |
| **app-services-1/decision** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 1240 | 1240 |
| **app-services-1/images** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 2 | 2 |
| **app-services-3/b-linux** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 832 | 832 |
| **app-services-3/b-linux-docker-amd** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 2 | 2 |
| **app-services-3/decision** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 29 | 29 |
| **app-services-3/images** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 2 | 2 |
| **code-analysis-1/linux-docker** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 2 | 2 |
| **code-analysis-1/linux-gw-gcp** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | amd64 | 1.22.2 | 43 | 43 |
| **code-analysis-3/linux-docker** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 2 | 2 |
| **code-analysis-3/linux-gw-gcp** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | amd64 | 1.22.2 | 2 | 2 |
| **code-coverage/bot** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 254 | 254 |
| **code-review/bot** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 6342 | 6342 |
| **comm-1/b-linux** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 105 | 105 |
| **comm-1/b-linux-aarch64** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | arm64 | 1.25.5 | 21 | 21 |
| **comm-1/b-linux-docker-amd** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 63 | 63 |
| **comm-1/b-linux-docker-large-amd** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 7 | 7 |
| **comm-1/b-linux-docker-xlarge-amd** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 5 | 5 |
| **comm-1/b-linux-gcp-aarch64** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | arm64 | 1.22.2 | 2 | 2 |
| **comm-1/b-linux-large** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 2 | 2 |
| **comm-1/b-linux-xlarge** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 2 | 2 |
| **comm-1/b-win2022** | generic-worker | 91.1.0 | multiuser | b8f34332c3 | windows | amd64 | 1.25.3 | 2 | 2 |
| **comm-1/decision** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 27 | 27 |
| **comm-1/images** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 29 | 29 |
| **comm-1/images-aarch64** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | arm64 | 1.25.5 | 2 | 2 |
| **comm-1/images-gcp-aarch64** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | arm64 | 1.22.2 | 2 | 2 |
| **comm-2/b-linux** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 2 | 2 |
| **comm-2/b-linux-aarch64** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | arm64 | 1.25.5 | 2 | 2 |
| **comm-2/b-linux-docker-amd** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 3 | 3 |
| **comm-2/b-linux-docker-large-amd** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 2 | 2 |
| **comm-2/b-linux-docker-xlarge-amd** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 2 | 2 |
| **comm-2/b-linux-gcp-aarch64** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | arm64 | 1.22.2 | 2 | 2 |
| **comm-2/b-linux-large** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 2 | 2 |
| **comm-2/b-linux-xlarge** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 2 | 2 |
| **comm-2/decision** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 2 | 2 |
| **comm-2/images** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 2 | 2 |
| **comm-2/images-aarch64** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | arm64 | 1.25.5 | 2 | 2 |
| **comm-2/images-gcp-aarch64** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | arm64 | 1.22.2 | 2 | 2 |
| **comm-3/b-linux** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 401 | 401 |
| **comm-3/b-linux-aarch64** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | arm64 | 1.25.5 | 5 | 5 |
| **comm-3/b-linux-docker-amd** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 1597 | 1597 |
| **comm-3/b-linux-docker-large-amd** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 7 | 7 |
| **comm-3/b-linux-docker-xlarge-amd** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 5 | 5 |
| **comm-3/b-linux-gcp-aarch64** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | arm64 | 1.22.2 | 2 | 2 |
| **comm-3/b-linux-large** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 2 | 2 |
| **comm-3/b-linux-xlarge** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 2 | 2 |
| **comm-3/b-win2022** | generic-worker | 91.1.0 | multiuser | b8f34332c3 | windows | amd64 | 1.25.3 | 470 | 470 |
| **comm-3/decision** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 14 | 14 |
| **comm-3/images** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 14 | 14 |
| **comm-3/images-aarch64** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | arm64 | 1.25.5 | 2 | 2 |
| **comm-3/images-gcp-aarch64** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | arm64 | 1.22.2 | 2 | 2 |
| **comm-t/misc** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 2 | 2 |
| **comm-t/t-linux-arm64-docker** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | arm64 | 1.25.5 | 2 | 2 |
| **comm-t/t-linux-docker** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 13 | 13 |
| **comm-t/t-linux-docker-amd** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 2172 | 2172 |
| **comm-t/t-linux-docker-noscratch** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 48 | 48 |
| **comm-t/t-linux-docker-noscratch-amd** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 2293 | 2293 |
| **comm-t/win11-64-2009** | generic-worker | 91.1.0 | multiuser | b8f34332c3 | windows | amd64 | 1.25.3 | 2 | 2 |
| **comm-t/win11-64-2009-ssd** | generic-worker | 91.1.0 | multiuser | b8f34332c3 | windows | amd64 | 1.25.3 | 2 | 2 |
| **comm-t/win11-64-24h2** | generic-worker | 88.1.3 | multiuser | ffc99d77c0 | windows | amd64 | 1.25.1 | 1882 | 1882 |
| **comm-t/win11-64-24h2-ssd** | generic-worker | 88.1.3 | multiuser | ffc99d77c0 | windows | amd64 | 1.25.1 | 194 | 194 |
| **comm-t/win11-a64-24h2** | generic-worker | 91.1.0 | multiuser | b8f34332c3 | windows | arm64 | 1.25.3 | 2 | 2 |
| **comm-t/win11-a64-24h2-alpha** | generic-worker | 93.1.4 | multiuser | 1088d85c3b | windows | arm64 | 1.25.4 | 7 | 7 |
| **enterprise-1/b-linux** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 1836 | 1836 |
| **enterprise-1/b-linux-aarch64** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | arm64 | 1.25.5 | 2 | 2 |
| **enterprise-1/b-linux-docker-amd** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 455 | 455 |
| **enterprise-1/b-linux-docker-large-amd** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 7 | 7 |
| **enterprise-1/b-linux-docker-xlarge-amd** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 613 | 613 |
| **enterprise-1/b-win2022** | generic-worker | 91.1.0 | multiuser | b8f34332c3 | windows | amd64 | 1.25.3 | 334 | 334 |
| **enterprise-1/decision** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 1215 | 1215 |
| **enterprise-1/images** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 3 | 3 |
| **enterprise-1/images-aarch64** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | arm64 | 1.25.5 | 2 | 2 |
| **enterprise-3/b-linux** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 505 | 505 |
| **enterprise-3/b-linux-aarch64** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | arm64 | 1.25.5 | 2 | 2 |
| **enterprise-3/b-linux-docker-amd** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 304 | 304 |
| **enterprise-3/b-linux-docker-large-amd** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 7 | 7 |
| **enterprise-3/b-linux-docker-xlarge-amd** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 288 | 288 |
| **enterprise-3/b-win2022** | generic-worker | 91.1.0 | multiuser | b8f34332c3 | windows | amd64 | 1.25.3 | 195 | 195 |
| **enterprise-3/decision** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 36 | 36 |
| **enterprise-3/images** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 4 | 4 |
| **enterprise-3/images-aarch64** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | arm64 | 1.25.5 | 2 | 2 |
| **enterprise-t/misc** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 53 | 53 |
| **enterprise-t/t-linux-docker-16c32gb-amd** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 585 | 585 |
| **enterprise-t/t-linux-docker-amd** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 5750 | 5750 |
| **enterprise-t/t-linux-docker-noscratch-amd** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 2637 | 2637 |
| **enterprise-t/win11-64-24h2** | generic-worker | 88.1.3 | multiuser | ffc99d77c0 | windows | amd64 | 1.25.1 | 1799 | 1799 |
| **enterprise-t/win11-64-24h2-alpha** | generic-worker | 96.1.0 | multiuser | e25f647384 | windows | amd64 | 1.25.6 | 2 | 2 |
| **enterprise-t/win11-64-24h2-source** | generic-worker | 88.1.3 | multiuser | ffc99d77c0 | windows | amd64 | 1.25.1 | 921 | 921 |
| **enterprise-t/win11-64-24h2-source-alpha** | generic-worker | 96.1.0 | multiuser | e25f647384 | windows | amd64 | 1.25.6 | 3459 | 3459 |
| **gecko-1/b-linux** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 5701 | 5701 |
| **gecko-1/b-linux-2204-kvm-gcp** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | amd64 | 1.22.2 | 2 | 2 |
| **gecko-1/b-linux-aarch64** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | arm64 | 1.25.5 | 10 | 10 |
| **gecko-1/b-linux-docker-alpha** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 2 | 2 |
| **gecko-1/b-linux-docker-amd** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 9225 | 9225 |
| **gecko-1/b-linux-docker-large-amd** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 69 | 69 |
| **gecko-1/b-linux-docker-xlarge-amd** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 1055 | 1055 |
| **gecko-1/b-linux-gcp-aarch64** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | arm64 | 1.22.2 | 6 | 6 |
| **gecko-1/b-linux-kvm** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 359 | 359 |
| **gecko-1/b-linux-large** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 4 | 4 |
| **gecko-1/b-linux-medium** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 14913 | 14913 |
| **gecko-1/b-linux-xlarge** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 28 | 28 |
| **gecko-1/b-win2022** | generic-worker | 91.1.0 | multiuser | b8f34332c3 | windows | amd64 | 1.25.3 | 733 | 733 |
| **gecko-1/b-win2022-xxlarge** | generic-worker | 91.1.0 | multiuser | b8f34332c3 | windows | amd64 | 1.25.3 | 17 | 17 |
| **gecko-1/decision** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 4986 | 4986 |
| **gecko-1/images** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 154 | 154 |
| **gecko-1/images-aarch64** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | arm64 | 1.25.5 | 16 | 16 |
| **gecko-1/images-gcp-aarch64** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | arm64 | 1.22.2 | 3 | 3 |
| **gecko-1/win11-a64-24h2-builder** | generic-worker | 91.1.0 | multiuser | b8f34332c3 | windows | arm64 | 1.25.3 | 29 | 29 |
| **gecko-1/win11-a64-24h2-builder-alpha** | generic-worker | 93.1.4 | multiuser | 1088d85c3b | windows | arm64 | 1.25.4 | 3 | 3 |
| **gecko-2/b-linux** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 205 | 205 |
| **gecko-2/b-linux-aarch64** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | arm64 | 1.25.5 | 2 | 2 |
| **gecko-2/b-linux-docker-amd** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 226 | 226 |
| **gecko-2/b-linux-docker-large-amd** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 8 | 8 |
| **gecko-2/b-linux-docker-xlarge-amd** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 41 | 41 |
| **gecko-2/b-linux-gcp-aarch64** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | arm64 | 1.22.2 | 2 | 2 |
| **gecko-2/b-linux-kvm** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 14 | 14 |
| **gecko-2/b-linux-large** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 2 | 2 |
| **gecko-2/b-linux-medium** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 228 | 228 |
| **gecko-2/b-linux-xlarge** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 2 | 2 |
| **gecko-2/b-win2022** | generic-worker | 91.1.0 | multiuser | b8f34332c3 | windows | amd64 | 1.25.3 | 17 | 17 |
| **gecko-2/decision** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 20 | 20 |
| **gecko-2/images** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 2 | 2 |
| **gecko-2/images-aarch64** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | arm64 | 1.25.5 | 2 | 2 |
| **gecko-2/images-gcp-aarch64** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | arm64 | 1.22.2 | 2 | 2 |
| **gecko-3/b-linux** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 20103 | 20103 |
| **gecko-3/b-linux-2204-kvm-gcp** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | amd64 | 1.22.2 | 2 | 2 |
| **gecko-3/b-linux-aarch64** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | arm64 | 1.25.5 | 48 | 48 |
| **gecko-3/b-linux-docker-alpha** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 2 | 2 |
| **gecko-3/b-linux-docker-amd** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 22529 | 22529 |
| **gecko-3/b-linux-docker-large-amd** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 187 | 187 |
| **gecko-3/b-linux-docker-xlarge-amd** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 2244 | 2244 |
| **gecko-3/b-linux-gcp-aarch64** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | arm64 | 1.22.2 | 2 | 2 |
| **gecko-3/b-linux-kvm** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 383 | 383 |
| **gecko-3/b-linux-large** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 43 | 43 |
| **gecko-3/b-linux-medium** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 21435 | 21435 |
| **gecko-3/b-linux-xlarge** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 97 | 97 |
| **gecko-3/b-win2022** | generic-worker | 91.1.0 | multiuser | b8f34332c3 | windows | amd64 | 1.25.3 | 7074 | 7074 |
| **gecko-3/b-win2022-xxlarge** | generic-worker | 91.1.0 | multiuser | b8f34332c3 | windows | amd64 | 1.25.3 | 11 | 11 |
| **gecko-3/decision** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 4988 | 4988 |
| **gecko-3/images** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 128 | 128 |
| **gecko-3/images-aarch64** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | arm64 | 1.25.5 | 18 | 18 |
| **gecko-3/images-gcp-aarch64** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | arm64 | 1.22.2 | 2 | 2 |
| **gecko-3/win11-a64-24h2-builder** | generic-worker | 91.1.0 | multiuser | b8f34332c3 | windows | arm64 | 1.25.3 | 71 | 71 |
| **gecko-t/misc** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 769 | 769 |
| **gecko-t/t-linux-2204-wayland** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | amd64 | 1.22.2 | 9128 | 9128 |
| **gecko-t/t-linux-2204-wayland-arm64-relsre** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | arm64 | 1.22.2 | 2 | 2 |
| **gecko-t/t-linux-2204-wayland-relsre** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | amd64 | 1.22.2 | 2 | 2 |
| **gecko-t/t-linux-2204-wayland-root-exp** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | amd64 | 1.22.2 | 2 | 2 |
| **gecko-t/t-linux-2204-wayland-snap** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | amd64 | 1.22.2 | 147 | 147 |
| **gecko-t/t-linux-2404-headless-alpha** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 2 | 2 |
| **gecko-t/t-linux-2404-headless-arm64-alpha** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | arm64 | 1.25.5 | 2 | 2 |
| **gecko-t/t-linux-2404-headless-ssd-alpha** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 2 | 2 |
| **gecko-t/t-linux-2404-wayland** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 2 | 2 |
| **gecko-t/t-linux-2404-wayland-relsre** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 2 | 2 |
| **gecko-t/t-linux-2404-wayland-snap** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 140 | 140 |
| **gecko-t/t-linux-arm64-docker** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | arm64 | 1.25.5 | 33 | 33 |
| **gecko-t/t-linux-docker** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 1887 | 1887 |
| **gecko-t/t-linux-docker-16c32gb-amd** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 4801 | 4801 |
| **gecko-t/t-linux-docker-amd** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 26724 | 26724 |
| **gecko-t/t-linux-docker-amd-alpha** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 2 | 2 |
| **gecko-t/t-linux-docker-kvm** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 86348 | 86348 |
| **gecko-t/t-linux-docker-kvm-alpha** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 2 | 2 |
| **gecko-t/t-linux-docker-noscratch** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 6240 | 6240 |
| **gecko-t/t-linux-docker-noscratch-amd** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 107171 | 107171 |
| **gecko-t/t-linux-docker-noscratch-amd-alpha** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 2 | 2 |
| **gecko-t/t-linux-xlarge-2204-wayland** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | amd64 | 1.22.2 | 11027 | 11027 |
| **gecko-t/t-linux-xlarge-2404-wayland** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 2 | 2 |
| **gecko-t/win10-64-2009** | generic-worker | 91.1.0 | multiuser | b8f34332c3 | windows | amd64 | 1.25.3 | 4607 | 4607 |
| **gecko-t/win10-64-2009-gpu** | generic-worker | 91.1.0 | multiuser | b8f34332c3 | windows | amd64 | 1.25.3 | 519 | 519 |
| **gecko-t/win10-64-2009-source** | generic-worker | 91.1.0 | multiuser | b8f34332c3 | windows | amd64 | 1.25.3 | 111 | 111 |
| **gecko-t/win10-64-2009-webgpu** | generic-worker | 91.1.0 | multiuser | b8f34332c3 | windows | amd64 | 1.25.3 | 2 | 2 |
| **gecko-t/win11-64-2009** | generic-worker | 91.1.0 | multiuser | b8f34332c3 | windows | amd64 | 1.25.3 | 666 | 666 |
| **gecko-t/win11-64-2009-alpha** | generic-worker | 93.1.4 | multiuser | 1088d85c3b | windows | amd64 | 1.25.4 | 2 | 2 |
| **gecko-t/win11-64-2009-gpu** | generic-worker | 91.1.0 | multiuser | b8f34332c3 | windows | amd64 | 1.25.3 | 402 | 402 |
| **gecko-t/win11-64-2009-gpu-alpha** | generic-worker | 93.1.4 | multiuser | 1088d85c3b | windows | amd64 | 1.25.4 | 2 | 2 |
| **gecko-t/win11-64-2009-source** | generic-worker | 91.1.0 | multiuser | b8f34332c3 | windows | amd64 | 1.25.3 | 2 | 2 |
| **gecko-t/win11-64-2009-source-alpha** | generic-worker | 93.1.4 | multiuser | 1088d85c3b | windows | amd64 | 1.25.4 | 2 | 2 |
| **gecko-t/win11-64-2009-ssd** | generic-worker | 91.1.0 | multiuser | b8f34332c3 | windows | amd64 | 1.25.3 | 2 | 2 |
| **gecko-t/win11-64-2009-ssd-gpu** | generic-worker | 91.1.0 | multiuser | b8f34332c3 | windows | amd64 | 1.25.3 | 2 | 2 |
| **gecko-t/win11-64-2009-webgpu** | generic-worker | 91.1.0 | multiuser | b8f34332c3 | windows | amd64 | 1.25.3 | 2 | 2 |
| **gecko-t/win11-64-2009-webgpu-alpha** | generic-worker | 93.1.4 | multiuser | 1088d85c3b | windows | amd64 | 1.25.4 | 2 | 2 |
| **gecko-t/win11-64-24h2** | generic-worker | 88.1.3 | multiuser | ffc99d77c0 | windows | amd64 | 1.25.1 | 47403 | 47403 |
| **gecko-t/win11-64-24h2-alpha** | generic-worker | 96.1.0 | multiuser | e25f647384 | windows | amd64 | 1.25.6 | 2619 | 2619 |
| **gecko-t/win11-64-24h2-amd** | generic-worker | 94.1.0 | multiuser | 44d8d7264b | windows | amd64 | 1.25.5 | 2 | 2 |
| **gecko-t/win11-64-24h2-gpu** | generic-worker | 88.1.3 | multiuser | ffc99d77c0 | windows | amd64 | 1.25.1 | 17914 | 17914 |
| **gecko-t/win11-64-24h2-gpu-alpha** | generic-worker | 96.1.0 | multiuser | e25f647384 | windows | amd64 | 1.25.6 | 481 | 481 |
| **gecko-t/win11-64-24h2-large** | generic-worker | 88.1.3 | multiuser | ffc99d77c0 | windows | amd64 | 1.25.1 | 1826 | 1826 |
| **gecko-t/win11-64-24h2-large-alpha** | generic-worker | 96.1.0 | multiuser | e25f647384 | windows | amd64 | 1.25.6 | 192 | 192 |
| **gecko-t/win11-64-24h2-source** | generic-worker | 88.1.3 | multiuser | ffc99d77c0 | windows | amd64 | 1.25.1 | 2378 | 2378 |
| **gecko-t/win11-64-24h2-source-alpha** | generic-worker | 96.1.0 | multiuser | e25f647384 | windows | amd64 | 1.25.6 | 2 | 2 |
| **gecko-t/win11-64-24h2-ssd** | generic-worker | 88.1.3 | multiuser | ffc99d77c0 | windows | amd64 | 1.25.1 | 2 | 2 |
| **gecko-t/win11-64-24h2-ssd-gpu** | generic-worker | 88.1.3 | multiuser | ffc99d77c0 | windows | amd64 | 1.25.1 | 2 | 2 |
| **gecko-t/win11-64-24h2-webgpu** | generic-worker | 88.1.3 | multiuser | ffc99d77c0 | windows | amd64 | 1.25.1 | 3764 | 3764 |
| **gecko-t/win11-64-24h2-webgpu-alpha** | generic-worker | 96.1.0 | multiuser | e25f647384 | windows | amd64 | 1.25.6 | 2 | 2 |
| **gecko-t/win11-a64-24h2** | generic-worker | 91.1.0 | multiuser | b8f34332c3 | windows | arm64 | 1.25.3 | 193 | 193 |
| **gecko-t/win11-a64-24h2-alpha** | generic-worker | 93.1.4 | multiuser | 1088d85c3b | windows | arm64 | 1.25.4 | 3 | 3 |
| **glean-1/b-linux** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 12 | 12 |
| **glean-1/decision** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 1223 | 1223 |
| **glean-1/images** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 2 | 2 |
| **glean-3/b-linux** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 11 | 11 |
| **glean-3/decision** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 4 | 4 |
| **glean-3/images** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 2 | 2 |
| **mobile-1/b-linux** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 161 | 161 |
| **mobile-1/decision** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 1189 | 1189 |
| **mobile-1/images** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 2 | 2 |
| **mobile-3/b-linux** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 248 | 248 |
| **mobile-3/decision** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 133 | 133 |
| **mobile-3/images** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 2 | 2 |
| **mozilla-1/b-linux** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 2 | 2 |
| **mozilla-1/b-win2022** | generic-worker | 91.1.0 | multiuser | b8f34332c3 | windows | amd64 | 1.25.3 | 2 | 2 |
| **mozilla-1/decision** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 1210 | 1210 |
| **mozilla-1/images** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 2 | 2 |
| **mozilla-3/b-linux** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 2 | 2 |
| **mozilla-3/b-win2022** | generic-worker | 91.1.0 | multiuser | b8f34332c3 | windows | amd64 | 1.25.3 | 3 | 3 |
| **mozilla-3/decision** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 424 | 424 |
| **mozilla-3/images** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 2 | 2 |
| **mozilla-t/t-linux-2204-wayland** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | amd64 | 1.22.2 | 553 | 553 |
| **mozilla-t/t-linux-2204-wayland-relsre** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | amd64 | 1.22.2 | 2 | 2 |
| **mozilla-t/t-linux-2404-wayland** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 2 | 2 |
| **mozilla-t/t-linux-arm64-docker** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | arm64 | 1.25.5 | 2 | 2 |
| **mozilla-t/t-linux-docker** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 35 | 35 |
| **mozilla-t/t-linux-docker-amd** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 2 | 2 |
| **mozilla-t/t-linux-docker-noscratch** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 6 | 6 |
| **mozilla-t/t-linux-docker-noscratch-amd** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 3 | 3 |
| **mozillavpn-1/b-linux** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 408 | 408 |
| **mozillavpn-1/b-linux-gcp-gw** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | amd64 | 1.22.2 | 2 | 2 |
| **mozillavpn-1/b-linux-large** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 292 | 292 |
| **mozillavpn-1/b-win2022** | generic-worker | 91.1.0 | multiuser | b8f34332c3 | windows | amd64 | 1.25.3 | 96 | 96 |
| **mozillavpn-1/b-win2022-alpha** | generic-worker | 91.1.0 | multiuser | b8f34332c3 | windows | amd64 | 1.25.3 | 2 | 2 |
| **mozillavpn-1/decision** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 1214 | 1214 |
| **mozillavpn-1/images** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 33 | 33 |
| **mozillavpn-1/win11-a64-24h2-builder** | generic-worker | 91.1.0 | multiuser | b8f34332c3 | windows | arm64 | 1.25.3 | 2 | 2 |
| **mozillavpn-3/b-linux** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 201 | 201 |
| **mozillavpn-3/b-linux-gcp-gw** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | amd64 | 1.22.2 | 2 | 2 |
| **mozillavpn-3/b-linux-large** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 179 | 179 |
| **mozillavpn-3/b-win2022** | generic-worker | 91.1.0 | multiuser | b8f34332c3 | windows | amd64 | 1.25.3 | 96 | 96 |
| **mozillavpn-3/decision** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 29 | 29 |
| **mozillavpn-3/images** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 3 | 3 |
| **mozillavpn-3/win11-a64-24h2-builder** | generic-worker | 91.1.0 | multiuser | b8f34332c3 | windows | arm64 | 1.25.3 | 2 | 2 |
| **nss-1/b-win2022** | generic-worker | 91.1.0 | multiuser | b8f34332c3 | windows | amd64 | 1.25.3 | 54 | 54 |
| **nss-1/b-win2022-alpha** | generic-worker | 91.1.0 | multiuser | b8f34332c3 | windows | amd64 | 1.25.3 | 2 | 2 |
| **nss-1/decision** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 4 | 4 |
| **nss-1/images** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 2 | 2 |
| **nss-1/linux-docker** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 108 | 108 |
| **nss-1/win11-a64-24h2** | generic-worker | 91.1.0 | multiuser | b8f34332c3 | windows | arm64 | 1.25.3 | 3 | 3 |
| **nss-1/win11-a64-24h2-alpha** | generic-worker | 93.1.4 | multiuser | 1088d85c3b | windows | arm64 | 1.25.4 | 3 | 3 |
| **nss-1/win11-a64-24h2-builder** | generic-worker | 91.1.0 | multiuser | b8f34332c3 | windows | arm64 | 1.25.3 | 2 | 2 |
| **nss-1/win11-a64-24h2-builder-alpha** | generic-worker | 93.1.4 | multiuser | 1088d85c3b | windows | arm64 | 1.25.4 | 3 | 3 |
| **nss-3/b-win2022** | generic-worker | 91.1.0 | multiuser | b8f34332c3 | windows | amd64 | 1.25.3 | 28 | 28 |
| **nss-3/decision** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 3 | 3 |
| **nss-3/images** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 2 | 2 |
| **nss-3/linux-docker** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 75 | 75 |
| **nss-3/win11-a64-24h2-builder** | generic-worker | 91.1.0 | multiuser | b8f34332c3 | windows | arm64 | 1.25.3 | 2 | 2 |
| **nss-t/t-linux-arm64-docker** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | arm64 | 1.25.5 | 2 | 2 |
| **nss-t/t-linux-docker** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 203 | 203 |
| **nss-t/t-linux-docker-amd** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 3 | 3 |
| **nss-t/win11-a64-24h2** | generic-worker | 91.1.0 | multiuser | b8f34332c3 | windows | arm64 | 1.25.3 | 2 | 2 |
| **nss-t/win11-a64-24h2-alpha** | generic-worker | 93.1.4 | multiuser | 1088d85c3b | windows | arm64 | 1.25.4 | 3 | 3 |
| **proj-autophone/gecko-t-bitbar-gw-perf-a55** | generic-worker | 36.0.0 | simple | b9cb11293f | linux | amd64 | 1.13.7 | 0 | 0 |
| **proj-autophone/gecko-t-bitbar-gw-unit-p5** | generic-worker | 36.0.0 | simple | b9cb11293f | linux | amd64 | 1.13.7 | 0 | 0 |
| **proj-autophone/gecko-t-lambda-perf-a55** | generic-worker | 87.0.0 | insecure | 99a1fcafbb | linux | amd64 | 1.24.5 | 0 | 0 |
| **releng-1/decision** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 88 | 88 |
| **releng-1/linux-docker** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 128 | 128 |
| **releng-3/decision** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 57 | 57 |
| **releng-3/linux-docker** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 109 | 109 |
| **releng-hardware/applicationservices-b-1-osx1015** | generic-worker | 60.3.4 | multiuser | 943a6f2b0d | darwin | amd64 | 1.22.0 | 0 | 0 |
| **releng-hardware/applicationservices-b-3-osx1015** | generic-worker | 60.3.4 | multiuser | 943a6f2b0d | darwin | amd64 | 1.22.0 | 0 | 0 |
| **releng-hardware/enterprise-1-b-osx-arm64** | generic-worker | 60.3.4 | multiuser | 943a6f2b0d | darwin | arm64 | 1.22.0 | 0 | 0 |
| **releng-hardware/enterprise-3-b-osx-arm64** | generic-worker | 60.3.4 | multiuser | 943a6f2b0d | darwin | arm64 | 1.22.0 | 0 | 0 |
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
| **releng-hardware/gecko-t-osx-1500-m4** | generic-worker | 95.0.3 | multiuser | a996b8269b | darwin | arm64 | 1.25.5 | 0 | 0 |
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
| **releng-hardware/win11-64-24h2-hw-alpha** | generic-worker | 91.1.0 | multiuser | b8f34332c3 | windows | amd64 | 1.25.3 | 0 | 0 |
| **releng-hardware/win11-64-24h2-hw-perf-sheriff** | generic-worker | 86.0.2 | multiuser | 9ce326432e | windows | amd64 | 1.24.4 | 0 | 0 |
| **releng-hardware/win11-64-24h2-hw-ref** | generic-worker | 91.1.1 | multiuser | 9ee80dc4d0 | windows | amd64 | 1.25.3 | 0 | 0 |
| **releng-hardware/win11-64-24h2-hw-ref-alpha** | generic-worker | 91.1.0 | multiuser | b8f34332c3 | windows | amd64 | 1.25.3 | 0 | 0 |
| **releng-hardware/win11-64-24h2-hw-relops1213** | generic-worker | 91.1.0 | multiuser | b8f34332c3 | windows | amd64 | 1.25.3 | 0 | 0 |
| **releng-t/linux-docker** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 405 | 405 |
| **relops-1/decision** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 2 | 2 |
| **relops-1/images** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 2 | 2 |
| **relops-3/decision** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 3 | 3 |
| **relops-3/images** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 2 | 2 |
| **scriptworker-1/decision** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 10 | 10 |
| **scriptworker-1/images** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 37 | 37 |
| **scriptworker-3/decision** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 11 | 11 |
| **scriptworker-3/images** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 58 | 58 |
| **taskgraph-1/decision** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 11 | 11 |
| **taskgraph-1/images** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 3 | 3 |
| **taskgraph-3/decision** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 8 | 8 |
| **taskgraph-3/images** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 3 | 3 |
| **taskgraph-t/linux-docker** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 70 | 70 |
| **taskgraph-t/win11-64-24h2** | generic-worker | 88.1.3 | multiuser | ffc99d77c0 | windows | amd64 | 1.25.1 | 2 | 2 |
| **translations-1/b-linux-large-gcp-1tb-32-256-d2g** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 29 | 29 |
| **translations-1/b-linux-large-gcp-1tb-32-256-std-d2g** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 2 | 2 |
| **translations-1/b-linux-large-gcp-1tb-64-512-d2g** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 2 | 2 |
| **translations-1/b-linux-large-gcp-1tb-64-512-std-d2g** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 2 | 2 |
| **translations-1/b-linux-large-gcp-d2g** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 184 | 184 |
| **translations-1/b-linux-large-gcp-d2g-1tb** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 2 | 2 |
| **translations-1/b-linux-large-gcp-d2g-1tb-standard** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 2 | 2 |
| **translations-1/b-linux-large-gcp-d2g-300gb** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 77 | 77 |
| **translations-1/b-linux-v100-gpu-d2g** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 72 | 72 |
| **translations-1/b-linux-v100-gpu-d2g-4** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 22 | 22 |
| **translations-1/b-linux-v100-gpu-d2g-4-1tb** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 17 | 17 |
| **translations-1/b-linux-v100-gpu-d2g-4-1tb-alpha** | generic-worker | 84.1.2 | multiuser | 4a139f395a | linux | amd64 | 1.24.4 | 2 | 2 |
| **translations-1/b-linux-v100-gpu-d2g-4-1tb-standard** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 2 | 2 |
| **translations-1/b-linux-v100-gpu-d2g-4-1tb-std-alpha** | generic-worker | 84.1.2 | multiuser | 4a139f395a | linux | amd64 | 1.24.4 | 2 | 2 |
| **translations-1/b-linux-v100-gpu-d2g-4-2tb** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 10 | 10 |
| **translations-1/b-linux-v100-gpu-d2g-4-2tb-alpha** | generic-worker | 84.1.2 | multiuser | 4a139f395a | linux | amd64 | 1.24.4 | 2 | 2 |
| **translations-1/b-linux-v100-gpu-d2g-4-300gb** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 27 | 27 |
| **translations-1/b-linux-v100-gpu-d2g-4-300gb-alpha** | generic-worker | 84.1.2 | multiuser | 4a139f395a | linux | amd64 | 1.24.4 | 2 | 2 |
| **translations-1/b-linux-v100-gpu-d2g-4-300gb-standard** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 4 | 4 |
| **translations-1/b-linux-v100-gpu-d2g-4-300gb-std-alpha** | generic-worker | 84.1.2 | multiuser | 4a139f395a | linux | amd64 | 1.24.4 | 2 | 2 |
| **translations-1/b-linux-v100-gpu-d2g-4-alpha** | generic-worker | 84.1.2 | multiuser | 4a139f395a | linux | amd64 | 1.24.4 | 2 | 2 |
| **translations-1/b-linux-v100-gpu-d2g-alpha** | generic-worker | 84.1.2 | multiuser | 4a139f395a | linux | amd64 | 1.24.4 | 2 | 2 |
| **translations-1/decision** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 1232 | 1232 |
| **translations-1/images** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 2 | 2 |
| **translations-t/t-linux-gw-noscratch-amd** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 3 | 3 |
| **xpi-1/b-linux** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 45 | 45 |
| **xpi-1/decision** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 25 | 25 |
| **xpi-1/images** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 2 | 2 |
| **xpi-3/b-linux** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 2 | 2 |
| **xpi-3/decision** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 2 | 2 |
| **xpi-3/images** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 2 | 2 |


## Docker Worker

Total: `151`

Count by version:

| Version | Count |
| :--- | ---: |
| 38.0.5 | 131 |
| 48.3.0 | 1 |
| unknown version | 19 |


Count by image:

| Version | Count |
| :--- | ---: |
| projects/taskcluster-imaging/global/images/docker-firefoxci-gcp-l1-googlecompute-2025-06-13t18-31-38z | 82 |
| projects/fxci-production-level3-workers/global/images/docker-firefoxci-gcp-l3-googlecompute-2024-02-05t23-18-22z | 49 |
| projects/taskcluster-imaging/global/images/docker-worker-gcp-u14-04-2025-06-16 | 19 |
| projects/taskcluster-imaging/global/images/docker-firefoxci-gcp-lt-googlecompute-2023-04-13t21-30-28z | 1 |


| Worker Pool | Implementation | Version | Total Workers | Total Capacity |
| --- | --- | --- | ---: | ---: |
| **app-services-1/b-linux-amd** | docker-worker | 38.0.5 | 3 | 3 |
| **app-services-1/b-linux-gcp** | docker-worker | 38.0.5 | 14 | 14 |
| **app-services-1/decision-gcp** | docker-worker | 38.0.5 | 798 | 1596 |
| **app-services-1/images-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **app-services-3/b-linux-amd** | docker-worker | 38.0.5 | 3 | 3 |
| **app-services-3/b-linux-gcp** | docker-worker | 38.0.5 | 39 | 39 |
| **app-services-3/decision-gcp** | docker-worker | 38.0.5 | 3 | 6 |
| **app-services-3/images-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **code-analysis-1/linux-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **code-analysis-3/linux-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **comm-1/b-linux-amd** | docker-worker | 38.0.5 | 2 | 2 |
| **comm-1/b-linux-gcp** | docker-worker | 38.0.5 | 403 | 403 |
| **comm-1/b-linux-large-amd** | docker-worker | 38.0.5 | 2 | 2 |
| **comm-1/b-linux-large-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **comm-1/b-linux-xlarge-amd** | docker-worker | 38.0.5 | 2 | 2 |
| **comm-1/b-linux-xlarge-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **comm-1/decision-gcp** | docker-worker | 38.0.5 | 191 | 191 |
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
| **comm-3/b-linux-gcp** | docker-worker | 38.0.5 | 3615 | 3615 |
| **comm-3/b-linux-large-amd** | docker-worker | 38.0.5 | 2 | 2 |
| **comm-3/b-linux-large-gcp** | docker-worker | 38.0.5 | 4 | 4 |
| **comm-3/b-linux-xlarge-amd** | docker-worker | 38.0.5 | 2 | 2 |
| **comm-3/b-linux-xlarge-gcp** | docker-worker | 38.0.5 | 7 | 7 |
| **comm-3/decision-gcp** | docker-worker | 38.0.5 | 101 | 101 |
| **comm-3/images-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **comm-t/misc-gcp** | docker-worker | 38.0.5 | 2 | 16 |
| **comm-t/t-linux-large-gcp** | docker-worker | unknown version | 2 | 2 |
| **comm-t/t-linux-large-noscratch-gcp** | docker-worker | unknown version | 2 | 2 |
| **comm-t/t-linux-xlarge-gcp** | docker-worker | unknown version | 2 | 2 |
| **comm-t/t-linux-xlarge-noscratch-gcp** | docker-worker | unknown version | 2 | 2 |
| **comm-t/t-linux-xlarge-source-gcp** | docker-worker | unknown version | 2 | 2 |
| **comm-t/t-linux-xlarge-source-noscratch-gcp** | docker-worker | unknown version | 2 | 2 |
| **gecko-1/b-linux-amd** | docker-worker | 38.0.5 | 88 | 88 |
| **gecko-1/b-linux-gcp** | docker-worker | 38.0.5 | 30 | 30 |
| **gecko-1/b-linux-gcp-relops1411** | docker-worker | 38.0.5 | 2 | 2 |
| **gecko-1/b-linux-gcp-test-bug-1882320** | docker-worker | 38.0.5 | 2 | 2 |
| **gecko-1/b-linux-kvm-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **gecko-1/b-linux-large-amd** | docker-worker | 38.0.5 | 3 | 3 |
| **gecko-1/b-linux-large-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **gecko-1/b-linux-medium-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **gecko-1/b-linux-xlarge-amd** | docker-worker | 38.0.5 | 3 | 3 |
| **gecko-1/b-linux-xlarge-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **gecko-1/b-linux-xlarge-gcp-bug1797804-c2** | docker-worker | 38.0.5 | 2 | 2 |
| **gecko-1/decision-gcp** | docker-worker | 38.0.5 | 9 | 9 |
| **gecko-1/decision-gcp-c3d4-amd-bug1990935** | docker-worker | 38.0.5 | 2 | 2 |
| **gecko-1/decision-gcp-c3d8-amd-bug1990935** | docker-worker | 38.0.5 | 2 | 2 |
| **gecko-1/decision-gcp-c4d4-amd-bug1990935** | docker-worker | 38.0.5 | 2 | 2 |
| **gecko-1/decision-gcp-c4d4-hcpu-amd-bug1990935** | docker-worker | 38.0.5 | 2 | 2 |
| **gecko-1/decision-gcp-c4d8-amd-bug1990935** | docker-worker | 38.0.5 | 2 | 2 |
| **gecko-1/decision-gcp-c4d8-hcpu-amd-bug1990935** | docker-worker | 38.0.5 | 2 | 2 |
| **gecko-1/decision-gcp-c4d8-lssd-amd-bug1990935** | docker-worker | 38.0.5 | 2 | 2 |
| **gecko-1/images-gcp** | docker-worker | 38.0.5 | 31 | 31 |
| **gecko-2/b-linux-amd** | docker-worker | 38.0.5 | 3 | 3 |
| **gecko-2/b-linux-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **gecko-2/b-linux-kvm-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **gecko-2/b-linux-large-amd** | docker-worker | 38.0.5 | 3 | 3 |
| **gecko-2/b-linux-large-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **gecko-2/b-linux-medium-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **gecko-2/b-linux-xlarge-amd** | docker-worker | 38.0.5 | 2 | 2 |
| **gecko-2/b-linux-xlarge-gcp** | docker-worker | 38.0.5 | 3 | 3 |
| **gecko-2/decision-gcp** | docker-worker | 38.0.5 | 10 | 10 |
| **gecko-2/images-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **gecko-3/b-linux-amd** | docker-worker | 38.0.5 | 172 | 172 |
| **gecko-3/b-linux-gcp** | docker-worker | 38.0.5 | 304 | 304 |
| **gecko-3/b-linux-kvm-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **gecko-3/b-linux-large-amd** | docker-worker | 38.0.5 | 2 | 2 |
| **gecko-3/b-linux-large-gcp** | docker-worker | 38.0.5 | 3 | 3 |
| **gecko-3/b-linux-medium-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **gecko-3/b-linux-xlarge-amd** | docker-worker | 38.0.5 | 2 | 2 |
| **gecko-3/b-linux-xlarge-gcp** | docker-worker | 38.0.5 | 9 | 9 |
| **gecko-3/decision-gcp** | docker-worker | 38.0.5 | 819 | 819 |
| **gecko-3/images-gcp** | docker-worker | 38.0.5 | 3 | 3 |
| **gecko-t/misc-gcp** | docker-worker | 38.0.5 | 22 | 176 |
| **gecko-t/t-linux-kvm-gcp** | docker-worker | unknown version | 2 | 2 |
| **gecko-t/t-linux-kvm-gcp-relops1411** | docker-worker | unknown version | 2 | 2 |
| **gecko-t/t-linux-kvm-noscratch-gcp** | docker-worker | unknown version | 5 | 5 |
| **gecko-t/t-linux-large-gcp** | docker-worker | unknown version | 2 | 2 |
| **gecko-t/t-linux-large-hostub2204-gcp** | docker-worker | 48.3.0 | 2 | 2 |
| **gecko-t/t-linux-large-noscratch-gcp** | docker-worker | unknown version | 1245 | 1245 |
| **gecko-t/t-linux-xlarge-gcp** | docker-worker | unknown version | 15 | 15 |
| **gecko-t/t-linux-xlarge-noscratch-gcp** | docker-worker | unknown version | 751 | 751 |
| **gecko-t/t-linux-xlarge-source-gcp** | docker-worker | unknown version | 73 | 73 |
| **gecko-t/t-linux-xlarge-source-noscratch-gcp** | docker-worker | unknown version | 2 | 2 |
| **glean-1/b-linux-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **glean-1/decision-gcp** | docker-worker | 38.0.5 | 785 | 1570 |
| **glean-1/images-gcp** | docker-worker | 38.0.5 | 3 | 3 |
| **glean-3/b-linux-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **glean-3/decision-gcp** | docker-worker | 38.0.5 | 2 | 4 |
| **glean-3/images-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **infra/build-decision** | docker-worker | 38.0.5 | 1 | 8 |
| **mobile-1/b-linux-gcp** | docker-worker | 38.0.5 | 15 | 15 |
| **mobile-1/decision-gcp** | docker-worker | 38.0.5 | 784 | 1568 |
| **mobile-1/images-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **mobile-3/b-linux-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **mobile-3/decision-gcp** | docker-worker | 38.0.5 | 2 | 4 |
| **mobile-3/images-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **mozilla-1/b-linux-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **mozilla-1/decision-gcp** | docker-worker | 38.0.5 | 786 | 1572 |
| **mozilla-1/images-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **mozilla-3/b-linux-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **mozilla-3/decision-gcp** | docker-worker | 38.0.5 | 2 | 4 |
| **mozilla-3/images-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **mozilla-t/t-linux-large-gcp** | docker-worker | unknown version | 2 | 2 |
| **mozilla-t/t-linux-large-noscratch-gcp** | docker-worker | unknown version | 2 | 2 |
| **mozillavpn-1/b-linux-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **mozillavpn-1/b-linux-large-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **mozillavpn-1/decision-gcp** | docker-worker | 38.0.5 | 790 | 1580 |
| **mozillavpn-1/images-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **mozillavpn-3/b-linux-gcp** | docker-worker | 38.0.5 | 3 | 3 |
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
| **translations-1/decision-gcp** | docker-worker | 38.0.5 | 1762 | 3524 |
| **translations-1/images-gcp** | docker-worker | 38.0.5 | 2 | 2 |


## Script Worker

Total: `46`



| Worker Pool | Implementation | Version | Total Workers | Total Capacity |
| --- | --- | --- | ---: | ---: |
| **scriptworker-k8s/adhoc-t-signing** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
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
| **scriptworker-k8s/gecko-1-pushflatpak** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/gecko-1-pushmsix** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/gecko-3-balrog** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/gecko-3-beetmover** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/gecko-3-bouncer** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/gecko-3-lando** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/gecko-3-pushapk** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/gecko-3-shipit** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/gecko-3-signing** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/gecko-t-signing** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/mobile-3-bitrise** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/mobile-3-pushapk** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/mobile-3-signing** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/mobile-t-signing** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/mozillavpn-1-beetmover** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/mozillavpn-3-beetmover** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/mozillavpn-3-signing** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/mozillavpn-t-signing** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/translations-1-beetmover** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/xpi-t-signing** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-prov-v1/adhoc-signing-mac14m2** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
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
| **gecko-t/t-linux-vm-2204-wayland** |  | No artifacts found | 60 | 60 |
| **gecko-t/t-linux-vm-2204-wayland-snap** |  | No artifacts found | 2 | 2 |


## Version not determined [^2]

Total: `10`


Count by image:

| Version | Count |
| :--- | ---: |
| unknown | 5 |
| projects/taskcluster-imaging/global/images/gw-fxci-gcp-l1-2404-amd64-googlecompute-alpha | 1 |
|  | 4 |


| Worker Pool | Implementation | Version | Total Workers | Total Capacity |
| --- | --- | --- | ---: | ---: |
| **gecko-1/b-win2022-alpha** |  | Version not determined; task not (yet) claimed | 143 | 143 |
| **gecko-t/t-linux-2404-relsre** |  | Version not determined; task not (yet) claimed | 4 | 4 |
| **gecko-t/win10-64-2009-alpha** |  | Version not determined; task not (yet) claimed | 451 | 451 |
| **gecko-t/win10-64-2009-gpu-alpha** |  | Version not determined; task not (yet) claimed | 451 | 451 |
| **gecko-t/win10-64-2009-source-alpha** |  | Version not determined; task not (yet) claimed | 451 | 451 |
| **gecko-t/win10-64-2009-webgpu-alpha** |  | Version not determined; task not (yet) claimed | 451 | 451 |
| **proj-autophone/gecko-t-bitbar-gw-perf-p6** |  | Version not determined; task not (yet) claimed | 0 | 0 |
| **proj-autophone/gecko-t-bitbar-gw-perf-s24** |  | Version not determined; task not (yet) claimed | 0 | 0 |
| **proj-autophone/gecko-t-bitbar-gw-test-1** |  | Version not determined; task not (yet) claimed | 0 | 0 |
| **releng-hardware/gecko-1-b-osx-1015** |  | Version not determined; task not (yet) claimed | 0 | 0 |



[^1]: Those are the pools whose tasks were claimed and resolved by a worker as expected, but the worker did not publish either artifact `public/logs/live_backing.log` nor `public/logs/chain_of_trust.log`, which is the source used to identify the worker implementation.

[^2]: Probing task remains pending after two hours. Those are the pools that were not able to start any worker to claim the task within two hours.

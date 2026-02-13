

# Worker Pool Versions


## Generic Worker

Total: `342`

Count by version:

| Version | Count |
| :--- | ---: |
| 36.0.0 | 3 |
| 45.0.0 | 1 |
| 60.3.4 | 16 |
| 61.0.0 | 1 |
| 64.3.0 | 26 |
| 84.1.2 | 7 |
| 86.0.2 | 2 |
| 87.0.0 | 1 |
| 88.0.2 | 2 |
| 91.0.2 | 2 |
| 91.1.0 | 43 |
| 94.1.0 | 1 |
| 95.1.3 | 211 |
| 96.1.0 | 24 |
| 96.2.2 | 2 |


Count by image:

| Version | Count |
| :--- | ---: |
| projects/taskcluster-imaging/global/images/gw-fxci-gcp-l1-2404-amd64-headless-googlecompute-alpha | 6 |
| projects/fxci-production-level3-workers/global/images/gw-fxci-gcp-l3-2404-amd64-headless-googlecompute-alpha | 1 |
| unknown | 47 |
| /subscriptions/108d46d5-fe9b-4850-9a7d-8c914aa6c1f0/resourceGroups/rg-packer-worker-images/providers/Microsoft.Compute/galleries/win11_a64_24h2_builder_alpha/images/win11_a64_24h2_builder_alpha/versions/1.0.0 | 2 |
| /subscriptions/a30e97ab-734a-4f3b-a0e4-c51c0bff0701/resourceGroups/rg-packer-worker-images/providers/Microsoft.Compute/galleries/trusted_win11_a64_24h2_builder/images/trusted_win11_a64_24h2_builder/versions/1.0.9 | 4 |
| projects/taskcluster-imaging/global/images/gw-fxci-gcp-l1-gui-googlecompute-2025-01-13t22-33-40z | 2 |
| projects/fxci-production-level3-workers/global/images/gw-fxci-gcp-l3-gui-googlecompute-2024-09-18t05-46-31z | 3 |
| projects/taskcluster-imaging/global/images/gw-fxci-gcp-l1-2404-amd64-headless-googlecompute-2026-01-14 | 120 |
| projects/taskcluster-imaging/global/images/gw-fxci-gcp-l1-arm64-gui-googlecompute-2024-09-18t14-57-52z | 9 |
| projects/fxci-production-level3-workers/global/images/gw-fxci-gcp-l3-2404-arm64-headless-googlecompute-2026-01-14 | 6 |
| projects/taskcluster-imaging/global/images/gw-fxci-gcp-l1-2404-amd64-gui-googlecompute-alpha | 1 |
|  | 34 |
| projects/taskcluster-imaging/global/images/gw-fxci-gcp-l1-2404-amd64-gui-googlecompute-2026-01-14 | 4 |
| /subscriptions/108d46d5-fe9b-4850-9a7d-8c914aa6c1f0/resourceGroups/rg-packer-worker-images/providers/Microsoft.Compute/galleries/win11_a64_24h2_builder/images/win11_a64_24h2_builder/versions/1.0.9 | 4 |
| projects/fxci-production-level3-workers/global/images/gw-fxci-gcp-l3-2404-amd64-headless-googlecompute-2026-01-14 | 58 |
| projects/fxci-production-level3-workers/global/images/gw-fxci-gcp-l3-arm64-gui-googlecompute-2024-09-18t19-02-58z | 4 |
| projects/taskcluster-imaging/global/images/gw-fxci-gcp-l1-2404-arm64-headless-googlecompute-alpha | 1 |
| projects/taskcluster-imaging/global/images/gw-fxci-gcp-l1-2404-amd64-headless-googlecompute-alpha-tc | 7 |
| projects/taskcluster-imaging/global/images/gw-fxci-gcp-l1-gui-googlecompute-2024-08-22t22-48-09z | 8 |
| /subscriptions/108d46d5-fe9b-4850-9a7d-8c914aa6c1f0/resourceGroups/rg-packer-worker-images/providers/Microsoft.Compute/galleries/win11_a64_24h2_tester_alpha/images/win11_a64_24h2_tester_alpha/versions/1.0.0 | 3 |
| /subscriptions/108d46d5-fe9b-4850-9a7d-8c914aa6c1f0/resourceGroups/rg-packer-worker-images/providers/Microsoft.Compute/galleries/win11_a64_24h2_tester/images/win11_a64_24h2_tester/versions/1.0.9 | 4 |
| projects/taskcluster-imaging/global/images/gw-fxci-gcp-l1-2404-arm64-headless-googlecompute-2026-01-14 | 14 |


| Worker Pool | Implementation | Version | Engine | Revision | OS | Arch | GO | Total Workers | Total Capacity |
| --- | --- | --- | --- | --- | --- | --- | --- | ---: | ---: |
| **adhoc-1/decision** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 22 | 22 |
| **adhoc-1/images** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 71 | 71 |
| **adhoc-3/decision** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 8 | 8 |
| **adhoc-3/images** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 4 | 4 |
| **app-services-1/b-linux** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 196 | 196 |
| **app-services-1/b-linux-docker-amd** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 3 | 3 |
| **app-services-1/decision** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 1257 | 1257 |
| **app-services-1/images** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 2 | 2 |
| **app-services-3/b-linux** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 845 | 845 |
| **app-services-3/b-linux-docker-amd** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 2 | 2 |
| **app-services-3/decision** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 28 | 28 |
| **app-services-3/images** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 2 | 2 |
| **code-analysis-1/linux-docker** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 2 | 2 |
| **code-analysis-1/linux-gw-gcp** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | amd64 | 1.22.2 | 74 | 74 |
| **code-analysis-3/linux-docker** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 2 | 2 |
| **code-analysis-3/linux-gw-gcp** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | amd64 | 1.22.2 | 44 | 44 |
| **code-coverage/bot** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 226 | 226 |
| **code-review/bot** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 7100 | 7100 |
| **comm-1/b-linux** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 438 | 438 |
| **comm-1/b-linux-aarch64** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | arm64 | 1.25.5 | 2 | 2 |
| **comm-1/b-linux-docker-amd** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 11 | 11 |
| **comm-1/b-linux-docker-large-amd** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 2 | 2 |
| **comm-1/b-linux-docker-xlarge-amd** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 3 | 3 |
| **comm-1/b-linux-gcp-aarch64** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | arm64 | 1.22.2 | 2 | 2 |
| **comm-1/b-linux-large** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 2 | 2 |
| **comm-1/b-linux-xlarge** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 2 | 2 |
| **comm-1/b-win2022** | generic-worker | 91.1.0 | multiuser | b8f34332c3 | windows | amd64 | 1.25.3 | 2 | 2 |
| **comm-1/decision** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 152 | 152 |
| **comm-1/images** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 8 | 8 |
| **comm-1/images-aarch64** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | arm64 | 1.25.5 | 2 | 2 |
| **comm-1/images-gcp-aarch64** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | arm64 | 1.22.2 | 3 | 3 |
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
| **comm-3/b-linux** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 2404 | 2404 |
| **comm-3/b-linux-aarch64** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | arm64 | 1.25.5 | 3 | 3 |
| **comm-3/b-linux-docker-amd** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 1676 | 1676 |
| **comm-3/b-linux-docker-large-amd** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 3 | 3 |
| **comm-3/b-linux-docker-xlarge-amd** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 2 | 2 |
| **comm-3/b-linux-gcp-aarch64** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | arm64 | 1.22.2 | 2 | 2 |
| **comm-3/b-linux-large** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 2 | 2 |
| **comm-3/b-linux-xlarge** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 3 | 3 |
| **comm-3/b-win2022** | generic-worker | 91.1.0 | multiuser | b8f34332c3 | windows | amd64 | 1.25.3 | 672 | 672 |
| **comm-3/decision** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 87 | 87 |
| **comm-3/images** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 3 | 3 |
| **comm-3/images-aarch64** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | arm64 | 1.25.5 | 2 | 2 |
| **comm-3/images-gcp-aarch64** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | arm64 | 1.22.2 | 2 | 2 |
| **comm-t/misc** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 2 | 2 |
| **comm-t/t-linux-arm64-docker** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | arm64 | 1.25.5 | 2 | 2 |
| **comm-t/t-linux-docker** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 2 | 2 |
| **comm-t/t-linux-docker-amd** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 1537 | 1537 |
| **comm-t/t-linux-docker-noscratch** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 2 | 2 |
| **comm-t/t-linux-docker-noscratch-amd** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 2676 | 2676 |
| **comm-t/win11-64-2009** | generic-worker | 91.1.0 | multiuser | b8f34332c3 | windows | amd64 | 1.25.3 | 2 | 2 |
| **comm-t/win11-64-2009-ssd** | generic-worker | 91.1.0 | multiuser | b8f34332c3 | windows | amd64 | 1.25.3 | 2 | 2 |
| **comm-t/win11-64-24h2** | generic-worker | 96.1.0 | multiuser | e25f647384 | windows | amd64 | 1.25.6 | 4346 | 4346 |
| **comm-t/win11-64-24h2-ssd** | generic-worker | 96.1.0 | multiuser | e25f647384 | windows | amd64 | 1.25.6 | 423 | 423 |
| **comm-t/win11-a64-24h2** | generic-worker | 91.1.0 | multiuser | b8f34332c3 | windows | arm64 | 1.25.3 | 2 | 2 |
| **enterprise-1/b-linux** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 1117 | 1117 |
| **enterprise-1/b-linux-aarch64** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | arm64 | 1.25.5 | 6 | 6 |
| **enterprise-1/b-linux-docker-amd** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 533 | 533 |
| **enterprise-1/b-linux-docker-large-amd** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 15 | 15 |
| **enterprise-1/b-linux-docker-xlarge-amd** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 607 | 607 |
| **enterprise-1/b-win2022** | generic-worker | 91.1.0 | multiuser | b8f34332c3 | windows | amd64 | 1.25.3 | 273 | 273 |
| **enterprise-1/decision** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 1229 | 1229 |
| **enterprise-1/images** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 8 | 8 |
| **enterprise-1/images-aarch64** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | arm64 | 1.25.5 | 2 | 2 |
| **enterprise-1/win11-a64-24h2-builder** | generic-worker | 91.1.0 | multiuser | b8f34332c3 | windows | arm64 | 1.25.3 | 1 | 1 |
| **enterprise-3/b-linux** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 769 | 769 |
| **enterprise-3/b-linux-aarch64** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | arm64 | 1.25.5 | 6 | 6 |
| **enterprise-3/b-linux-docker-amd** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 553 | 553 |
| **enterprise-3/b-linux-docker-large-amd** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 19 | 19 |
| **enterprise-3/b-linux-docker-xlarge-amd** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 404 | 404 |
| **enterprise-3/b-win2022** | generic-worker | 91.1.0 | multiuser | b8f34332c3 | windows | amd64 | 1.25.3 | 248 | 248 |
| **enterprise-3/decision** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 46 | 46 |
| **enterprise-3/images** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 8 | 8 |
| **enterprise-3/images-aarch64** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | arm64 | 1.25.5 | 2 | 2 |
| **enterprise-3/win11-a64-24h2-builder** | generic-worker | 91.1.0 | multiuser | b8f34332c3 | windows | arm64 | 1.25.3 | 12 | 12 |
| **enterprise-t/misc** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 124 | 124 |
| **enterprise-t/t-linux-docker-16c32gb-amd** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 625 | 625 |
| **enterprise-t/t-linux-docker-amd** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 4057 | 4057 |
| **enterprise-t/t-linux-docker-noscratch-amd** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 3096 | 3096 |
| **enterprise-t/win11-64-24h2** | generic-worker | 96.1.0 | multiuser | e25f647384 | windows | amd64 | 1.25.6 | 1902 | 1902 |
| **enterprise-t/win11-64-24h2-source** | generic-worker | 96.1.0 | multiuser | e25f647384 | windows | amd64 | 1.25.6 | 100 | 100 |
| **gecko-1/b-linux** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 6478 | 6478 |
| **gecko-1/b-linux-2204-kvm-gcp** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | amd64 | 1.22.2 | 2 | 2 |
| **gecko-1/b-linux-aarch64** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | arm64 | 1.25.5 | 2 | 2 |
| **gecko-1/b-linux-docker-alpha** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 2 | 2 |
| **gecko-1/b-linux-docker-amd** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 8600 | 8600 |
| **gecko-1/b-linux-docker-large-amd** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 129 | 129 |
| **gecko-1/b-linux-docker-xlarge-amd** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 1071 | 1071 |
| **gecko-1/b-linux-gcp-aarch64** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | arm64 | 1.22.2 | 2 | 2 |
| **gecko-1/b-linux-kvm** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 328 | 328 |
| **gecko-1/b-linux-large** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 3 | 3 |
| **gecko-1/b-linux-medium** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 21385 | 21385 |
| **gecko-1/b-linux-xlarge** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 30 | 30 |
| **gecko-1/b-win2022** | generic-worker | 91.1.0 | multiuser | b8f34332c3 | windows | amd64 | 1.25.3 | 870 | 870 |
| **gecko-1/b-win2022-alpha** | generic-worker | 96.1.0 | multiuser | e25f647384 | windows | amd64 | 1.25.6 | 403 | 403 |
| **gecko-1/b-win2022-xxlarge** | generic-worker | 91.1.0 | multiuser | b8f34332c3 | windows | amd64 | 1.25.3 | 11 | 11 |
| **gecko-1/decision** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 4792 | 4792 |
| **gecko-1/images** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 60 | 60 |
| **gecko-1/images-aarch64** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | arm64 | 1.25.5 | 2 | 2 |
| **gecko-1/images-gcp-aarch64** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | arm64 | 1.22.2 | 2 | 2 |
| **gecko-1/win11-a64-24h2-builder** | generic-worker | 91.1.0 | multiuser | b8f34332c3 | windows | arm64 | 1.25.3 | 26 | 26 |
| **gecko-1/win11-a64-24h2-builder-alpha** | generic-worker | 96.1.0 | multiuser | e25f647384 | windows | arm64 | 1.25.6 | 5 | 5 |
| **gecko-2/b-linux** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 110 | 110 |
| **gecko-2/b-linux-aarch64** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | arm64 | 1.25.5 | 3 | 3 |
| **gecko-2/b-linux-docker-amd** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 90 | 90 |
| **gecko-2/b-linux-docker-large-amd** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 4 | 4 |
| **gecko-2/b-linux-docker-xlarge-amd** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 15 | 15 |
| **gecko-2/b-linux-gcp-aarch64** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | arm64 | 1.22.2 | 3 | 3 |
| **gecko-2/b-linux-kvm** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 7 | 7 |
| **gecko-2/b-linux-large** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 2 | 2 |
| **gecko-2/b-linux-medium** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 104 | 104 |
| **gecko-2/b-linux-xlarge** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 2 | 2 |
| **gecko-2/b-win2022** | generic-worker | 91.1.0 | multiuser | b8f34332c3 | windows | amd64 | 1.25.3 | 8 | 8 |
| **gecko-2/decision** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 12 | 12 |
| **gecko-2/images** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 2 | 2 |
| **gecko-2/images-aarch64** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | arm64 | 1.25.5 | 2 | 2 |
| **gecko-2/images-gcp-aarch64** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | arm64 | 1.22.2 | 2 | 2 |
| **gecko-3/b-linux** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 18334 | 18334 |
| **gecko-3/b-linux-2204-kvm-gcp** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | amd64 | 1.22.2 | 2 | 2 |
| **gecko-3/b-linux-aarch64** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | arm64 | 1.25.5 | 2 | 2 |
| **gecko-3/b-linux-docker-alpha** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 2 | 2 |
| **gecko-3/b-linux-docker-amd** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 23565 | 23565 |
| **gecko-3/b-linux-docker-large-amd** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 205 | 205 |
| **gecko-3/b-linux-docker-xlarge-amd** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 4007 | 4007 |
| **gecko-3/b-linux-gcp-aarch64** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | arm64 | 1.22.2 | 2 | 2 |
| **gecko-3/b-linux-kvm** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 717 | 717 |
| **gecko-3/b-linux-large** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 2 | 2 |
| **gecko-3/b-linux-medium** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 21756 | 21756 |
| **gecko-3/b-linux-xlarge** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 83 | 83 |
| **gecko-3/b-win2022** | generic-worker | 91.1.0 | multiuser | b8f34332c3 | windows | amd64 | 1.25.3 | 6194 | 6194 |
| **gecko-3/b-win2022-xxlarge** | generic-worker | 91.1.0 | multiuser | b8f34332c3 | windows | amd64 | 1.25.3 | 13 | 13 |
| **gecko-3/decision** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 5538 | 5538 |
| **gecko-3/images** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 14 | 14 |
| **gecko-3/images-aarch64** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | arm64 | 1.25.5 | 2 | 2 |
| **gecko-3/images-gcp-aarch64** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | arm64 | 1.22.2 | 2 | 2 |
| **gecko-3/win11-a64-24h2-builder** | generic-worker | 91.1.0 | multiuser | b8f34332c3 | windows | arm64 | 1.25.3 | 61 | 61 |
| **gecko-t/misc** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 1102 | 1102 |
| **gecko-t/t-linux-2204-wayland** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | amd64 | 1.22.2 | 8732 | 8732 |
| **gecko-t/t-linux-2204-wayland-arm64-relsre** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | arm64 | 1.22.2 | 3 | 3 |
| **gecko-t/t-linux-2204-wayland-relsre** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | amd64 | 1.22.2 | 3 | 3 |
| **gecko-t/t-linux-2204-wayland-root-exp** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | amd64 | 1.22.2 | 2 | 2 |
| **gecko-t/t-linux-2204-wayland-snap** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | amd64 | 1.22.2 | 113 | 113 |
| **gecko-t/t-linux-2404-headless-alpha** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 2 | 2 |
| **gecko-t/t-linux-2404-headless-arm64-alpha** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | arm64 | 1.25.5 | 3 | 3 |
| **gecko-t/t-linux-2404-headless-ssd-alpha** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 2 | 2 |
| **gecko-t/t-linux-2404-wayland** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 2 | 2 |
| **gecko-t/t-linux-2404-wayland-relsre** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 2 | 2 |
| **gecko-t/t-linux-2404-wayland-snap** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 114 | 114 |
| **gecko-t/t-linux-arm64-docker** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | arm64 | 1.25.5 | 35 | 35 |
| **gecko-t/t-linux-docker** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 1961 | 1961 |
| **gecko-t/t-linux-docker-16c32gb-amd** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 4041 | 4041 |
| **gecko-t/t-linux-docker-amd** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 30619 | 30619 |
| **gecko-t/t-linux-docker-amd-alpha** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 2 | 2 |
| **gecko-t/t-linux-docker-kvm** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 80126 | 80126 |
| **gecko-t/t-linux-docker-kvm-alpha** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 2 | 2 |
| **gecko-t/t-linux-docker-noscratch** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 7175 | 7175 |
| **gecko-t/t-linux-docker-noscratch-amd** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 109951 | 109951 |
| **gecko-t/t-linux-docker-noscratch-amd-alpha** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 2 | 2 |
| **gecko-t/t-linux-xlarge-2204-wayland** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | amd64 | 1.22.2 | 11366 | 11366 |
| **gecko-t/t-linux-xlarge-2404-wayland** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 3 | 3 |
| **gecko-t/win10-64-2009** | generic-worker | 91.1.0 | multiuser | b8f34332c3 | windows | amd64 | 1.25.3 | 4984 | 4984 |
| **gecko-t/win10-64-2009-gpu** | generic-worker | 91.1.0 | multiuser | b8f34332c3 | windows | amd64 | 1.25.3 | 531 | 531 |
| **gecko-t/win10-64-2009-source** | generic-worker | 91.1.0 | multiuser | b8f34332c3 | windows | amd64 | 1.25.3 | 84 | 84 |
| **gecko-t/win10-64-2009-webgpu** | generic-worker | 91.1.0 | multiuser | b8f34332c3 | windows | amd64 | 1.25.3 | 2 | 2 |
| **gecko-t/win11-64-2009** | generic-worker | 91.1.0 | multiuser | b8f34332c3 | windows | amd64 | 1.25.3 | 2474 | 2474 |
| **gecko-t/win11-64-2009-alpha** | generic-worker | 96.1.0 | multiuser | e25f647384 | windows | amd64 | 1.25.6 | 2 | 2 |
| **gecko-t/win11-64-2009-gpu** | generic-worker | 91.1.0 | multiuser | b8f34332c3 | windows | amd64 | 1.25.3 | 1584 | 1584 |
| **gecko-t/win11-64-2009-gpu-alpha** | generic-worker | 96.1.0 | multiuser | e25f647384 | windows | amd64 | 1.25.6 | 2 | 2 |
| **gecko-t/win11-64-2009-source** | generic-worker | 91.1.0 | multiuser | b8f34332c3 | windows | amd64 | 1.25.3 | 3 | 3 |
| **gecko-t/win11-64-2009-source-alpha** | generic-worker | 96.1.0 | multiuser | e25f647384 | windows | amd64 | 1.25.6 | 2 | 2 |
| **gecko-t/win11-64-2009-ssd** | generic-worker | 91.1.0 | multiuser | b8f34332c3 | windows | amd64 | 1.25.3 | 2 | 2 |
| **gecko-t/win11-64-2009-ssd-gpu** | generic-worker | 91.1.0 | multiuser | b8f34332c3 | windows | amd64 | 1.25.3 | 2 | 2 |
| **gecko-t/win11-64-2009-webgpu** | generic-worker | 91.1.0 | multiuser | b8f34332c3 | windows | amd64 | 1.25.3 | 2 | 2 |
| **gecko-t/win11-64-2009-webgpu-alpha** | generic-worker | 96.1.0 | multiuser | e25f647384 | windows | amd64 | 1.25.6 | 2 | 2 |
| **gecko-t/win11-64-24h2** | generic-worker | 96.1.0 | multiuser | e25f647384 | windows | amd64 | 1.25.6 | 48639 | 48639 |
| **gecko-t/win11-64-24h2-amd** | generic-worker | 94.1.0 | multiuser | 44d8d7264b | windows | amd64 | 1.25.5 | 2 | 2 |
| **gecko-t/win11-64-24h2-gpu** | generic-worker | 96.1.0 | multiuser | e25f647384 | windows | amd64 | 1.25.6 | 18058 | 18058 |
| **gecko-t/win11-64-24h2-large** | generic-worker | 96.1.0 | multiuser | e25f647384 | windows | amd64 | 1.25.6 | 2234 | 2234 |
| **gecko-t/win11-64-24h2-source** | generic-worker | 96.1.0 | multiuser | e25f647384 | windows | amd64 | 1.25.6 | 2066 | 2066 |
| **gecko-t/win11-64-24h2-ssd** | generic-worker | 96.1.0 | multiuser | e25f647384 | windows | amd64 | 1.25.6 | 2 | 2 |
| **gecko-t/win11-64-24h2-ssd-gpu** | generic-worker | 96.1.0 | multiuser | e25f647384 | windows | amd64 | 1.25.6 | 2 | 2 |
| **gecko-t/win11-64-24h2-webgpu** | generic-worker | 96.1.0 | multiuser | e25f647384 | windows | amd64 | 1.25.6 | 3649 | 3649 |
| **gecko-t/win11-a64-24h2** | generic-worker | 91.1.0 | multiuser | b8f34332c3 | windows | arm64 | 1.25.3 | 177 | 177 |
| **gecko-t/win11-a64-24h2-alpha** | generic-worker | 96.1.0 | multiuser | e25f647384 | windows | arm64 | 1.25.6 | 11 | 11 |
| **glean-1/b-linux** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 21 | 21 |
| **glean-1/decision** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 1231 | 1231 |
| **glean-1/images** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 2 | 2 |
| **glean-3/b-linux** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 14 | 14 |
| **glean-3/decision** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 5 | 5 |
| **glean-3/images** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 2 | 2 |
| **mobile-1/b-linux** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 76 | 76 |
| **mobile-1/decision** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 1193 | 1193 |
| **mobile-1/images** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 2 | 2 |
| **mobile-3/b-linux** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 105 | 105 |
| **mobile-3/decision** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 113 | 113 |
| **mobile-3/images** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 2 | 2 |
| **mozilla-1/b-linux** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 2 | 2 |
| **mozilla-1/b-win2022** | generic-worker | 91.1.0 | multiuser | b8f34332c3 | windows | amd64 | 1.25.3 | 2 | 2 |
| **mozilla-1/decision** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 1214 | 1214 |
| **mozilla-1/images** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 11 | 11 |
| **mozilla-3/b-linux** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 2 | 2 |
| **mozilla-3/b-win2022** | generic-worker | 91.1.0 | multiuser | b8f34332c3 | windows | amd64 | 1.25.3 | 4 | 4 |
| **mozilla-3/decision** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 431 | 431 |
| **mozilla-3/images** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 2 | 2 |
| **mozilla-t/t-linux-2204-wayland** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | amd64 | 1.22.2 | 568 | 568 |
| **mozilla-t/t-linux-2204-wayland-relsre** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | amd64 | 1.22.2 | 2 | 2 |
| **mozilla-t/t-linux-2404-wayland** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 3 | 3 |
| **mozilla-t/t-linux-arm64-docker** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | arm64 | 1.25.5 | 2 | 2 |
| **mozilla-t/t-linux-docker** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 70 | 70 |
| **mozilla-t/t-linux-docker-amd** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 2 | 2 |
| **mozilla-t/t-linux-docker-noscratch** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 39 | 39 |
| **mozilla-t/t-linux-docker-noscratch-amd** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 2 | 2 |
| **mozillavpn-1/b-linux** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 531 | 531 |
| **mozillavpn-1/b-linux-gcp-gw** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | amd64 | 1.22.2 | 2 | 2 |
| **mozillavpn-1/b-linux-large** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 394 | 394 |
| **mozillavpn-1/b-win2022** | generic-worker | 91.1.0 | multiuser | b8f34332c3 | windows | amd64 | 1.25.3 | 136 | 136 |
| **mozillavpn-1/b-win2022-alpha** | generic-worker | 96.1.0 | multiuser | e25f647384 | windows | amd64 | 1.25.6 | 2 | 2 |
| **mozillavpn-1/decision** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 1216 | 1216 |
| **mozillavpn-1/images** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 78 | 78 |
| **mozillavpn-1/win11-a64-24h2-builder** | generic-worker | 91.1.0 | multiuser | b8f34332c3 | windows | arm64 | 1.25.3 | 3 | 3 |
| **mozillavpn-3/b-linux** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 151 | 151 |
| **mozillavpn-3/b-linux-gcp-gw** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | amd64 | 1.22.2 | 2 | 2 |
| **mozillavpn-3/b-linux-large** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 140 | 140 |
| **mozillavpn-3/b-win2022** | generic-worker | 91.1.0 | multiuser | b8f34332c3 | windows | amd64 | 1.25.3 | 59 | 59 |
| **mozillavpn-3/decision** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 28 | 28 |
| **mozillavpn-3/images** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 2 | 2 |
| **mozillavpn-3/win11-a64-24h2-builder** | generic-worker | 91.1.0 | multiuser | b8f34332c3 | windows | arm64 | 1.25.3 | 2 | 2 |
| **nss-1/b-win2022** | generic-worker | 91.1.0 | multiuser | b8f34332c3 | windows | amd64 | 1.25.3 | 34 | 34 |
| **nss-1/b-win2022-alpha** | generic-worker | 96.1.0 | multiuser | e25f647384 | windows | amd64 | 1.25.6 | 2 | 2 |
| **nss-1/decision** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 3 | 3 |
| **nss-1/images** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 2 | 2 |
| **nss-1/linux-docker** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 75 | 75 |
| **nss-1/win11-a64-24h2** | generic-worker | 91.1.0 | multiuser | b8f34332c3 | windows | arm64 | 1.25.3 | 2 | 2 |
| **nss-1/win11-a64-24h2-alpha** | generic-worker | 96.1.0 | multiuser | e25f647384 | windows | arm64 | 1.25.6 | 12 | 12 |
| **nss-1/win11-a64-24h2-builder** | generic-worker | 91.1.0 | multiuser | b8f34332c3 | windows | arm64 | 1.25.3 | 2 | 2 |
| **nss-1/win11-a64-24h2-builder-alpha** | generic-worker | 96.1.0 | multiuser | e25f647384 | windows | arm64 | 1.25.6 | 5 | 5 |
| **nss-3/b-win2022** | generic-worker | 91.1.0 | multiuser | b8f34332c3 | windows | amd64 | 1.25.3 | 71 | 71 |
| **nss-3/decision** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 6 | 6 |
| **nss-3/images** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 2 | 2 |
| **nss-3/linux-docker** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 275 | 275 |
| **nss-3/win11-a64-24h2-builder** | generic-worker | 91.1.0 | multiuser | b8f34332c3 | windows | arm64 | 1.25.3 | 2 | 2 |
| **nss-t/t-linux-arm64-docker** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | arm64 | 1.25.5 | 2 | 2 |
| **nss-t/t-linux-docker** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 372 | 372 |
| **nss-t/t-linux-docker-amd** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 2 | 2 |
| **nss-t/win11-a64-24h2** | generic-worker | 91.1.0 | multiuser | b8f34332c3 | windows | arm64 | 1.25.3 | 2 | 2 |
| **nss-t/win11-a64-24h2-alpha** | generic-worker | 96.1.0 | multiuser | e25f647384 | windows | arm64 | 1.25.6 | 4 | 4 |
| **proj-autophone/gecko-t-bitbar-gw-perf-a55** | generic-worker | 36.0.0 | simple | b9cb11293f | linux | amd64 | 1.13.7 | 0 | 0 |
| **proj-autophone/gecko-t-bitbar-gw-test-1** | generic-worker | 36.0.0 | simple | b9cb11293f | linux | amd64 | 1.13.7 | 0 | 0 |
| **proj-autophone/gecko-t-bitbar-gw-unit-p5** | generic-worker | 36.0.0 | simple | b9cb11293f | linux | amd64 | 1.13.7 | 0 | 0 |
| **proj-autophone/gecko-t-lambda-perf-a55** | generic-worker | 87.0.0 | insecure | 99a1fcafbb | linux | amd64 | 1.24.5 | 0 | 0 |
| **releng-1/decision** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 101 | 101 |
| **releng-1/linux-docker** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 107 | 107 |
| **releng-3/decision** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 45 | 45 |
| **releng-3/linux-docker** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 69 | 69 |
| **releng-hardware/applicationservices-b-1-osx1015** | generic-worker | 60.3.4 | multiuser | 943a6f2b0d | darwin | amd64 | 1.22.0 | 0 | 0 |
| **releng-hardware/applicationservices-b-3-osx1015** | generic-worker | 60.3.4 | multiuser | 943a6f2b0d | darwin | amd64 | 1.22.0 | 0 | 0 |
| **releng-hardware/enterprise-1-b-osx-arm64** | generic-worker | 60.3.4 | multiuser | 943a6f2b0d | darwin | arm64 | 1.22.0 | 0 | 0 |
| **releng-hardware/enterprise-3-b-osx-arm64** | generic-worker | 60.3.4 | multiuser | 943a6f2b0d | darwin | arm64 | 1.22.0 | 0 | 0 |
| **releng-hardware/gecko-1-b-osx-1015** | generic-worker | 60.3.4 | multiuser | 943a6f2b0d | darwin | amd64 | 1.22.0 | 0 | 0 |
| **releng-hardware/gecko-1-b-osx-1015-staging** | generic-worker | 60.3.4 | multiuser | 943a6f2b0d | darwin | amd64 | 1.22.0 | 0 | 0 |
| **releng-hardware/gecko-1-b-osx-arm64** | generic-worker | 60.3.4 | multiuser | 943a6f2b0d | darwin | arm64 | 1.22.0 | 0 | 0 |
| **releng-hardware/gecko-3-b-osx-arm64** | generic-worker | 60.3.4 | multiuser | 943a6f2b0d | darwin | arm64 | 1.22.0 | 0 | 0 |
| **releng-hardware/gecko-t-linux-netperf-2404** | generic-worker | 88.0.2 | insecure | fcaf4a25fc | linux | amd64 | 1.24.5 | 0 | 0 |
| **releng-hardware/gecko-t-linux-talos-1804** | generic-worker | 61.0.0 | simple | 3bd4419b4b | linux | amd64 | 1.22.1 | 0 | 0 |
| **releng-hardware/gecko-t-linux-talos-2404** | generic-worker | 88.0.2 | insecure | fcaf4a25fc | linux | amd64 | 1.24.5 | 0 | 0 |
| **releng-hardware/gecko-t-osx-1015-r8-staging** | generic-worker | 60.3.4 | simple | 943a6f2b0d | darwin | amd64 | 1.22.0 | 0 | 0 |
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
| **releng-hardware/win11-64-24h2-hw** | generic-worker | 91.1.0 | multiuser | b8f34332c3 | windows | amd64 | 1.25.3 | 0 | 0 |
| **releng-hardware/win11-64-24h2-hw-alpha** | generic-worker | 96.2.2 | multiuser | 773509947e | windows | amd64 | 1.25.7 | 0 | 0 |
| **releng-hardware/win11-64-24h2-hw-perf-debug** | generic-worker | 91.1.0 | multiuser | b8f34332c3 | windows | amd64 | 1.25.3 | 0 | 0 |
| **releng-hardware/win11-64-24h2-hw-perf-sheriff** | generic-worker | 91.1.0 | multiuser | b8f34332c3 | windows | amd64 | 1.25.3 | 0 | 0 |
| **releng-hardware/win11-64-24h2-hw-ref** | generic-worker | 91.1.0 | multiuser | b8f34332c3 | windows | amd64 | 1.25.3 | 0 | 0 |
| **releng-hardware/win11-64-24h2-hw-ref-alpha** | generic-worker | 96.2.2 | multiuser | 773509947e | windows | amd64 | 1.25.7 | 0 | 0 |
| **releng-t/linux-docker** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 474 | 474 |
| **relops-1/decision** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 2 | 2 |
| **relops-1/images** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 2 | 2 |
| **relops-3/decision** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 31 | 31 |
| **relops-3/images** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 2 | 2 |
| **scriptworker-1/decision** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 40 | 40 |
| **scriptworker-1/images** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 235 | 235 |
| **scriptworker-3/decision** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 26 | 26 |
| **scriptworker-3/images** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 197 | 197 |
| **taskgraph-1/decision** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 41 | 41 |
| **taskgraph-1/images** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 21 | 21 |
| **taskgraph-3/decision** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 17 | 17 |
| **taskgraph-3/images** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 13 | 13 |
| **taskgraph-t/linux-docker** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 398 | 398 |
| **taskgraph-t/win11-64-24h2** | generic-worker | 96.1.0 | multiuser | e25f647384 | windows | amd64 | 1.25.6 | 41 | 41 |
| **translations-1/b-linux-large-gcp-1tb-32-256-d2g** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 37 | 37 |
| **translations-1/b-linux-large-gcp-1tb-32-256-std-d2g** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 13 | 13 |
| **translations-1/b-linux-large-gcp-1tb-64-512-d2g** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 2 | 2 |
| **translations-1/b-linux-large-gcp-1tb-64-512-std-d2g** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 2 | 2 |
| **translations-1/b-linux-large-gcp-d2g** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 417 | 417 |
| **translations-1/b-linux-large-gcp-d2g-1tb** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 2 | 2 |
| **translations-1/b-linux-large-gcp-d2g-1tb-standard** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 2 | 2 |
| **translations-1/b-linux-large-gcp-d2g-300gb** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 320 | 320 |
| **translations-1/b-linux-v100-gpu-d2g** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 133 | 133 |
| **translations-1/b-linux-v100-gpu-d2g-4** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 187 | 187 |
| **translations-1/b-linux-v100-gpu-d2g-4-1tb** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 38 | 38 |
| **translations-1/b-linux-v100-gpu-d2g-4-1tb-alpha** | generic-worker | 84.1.2 | multiuser | 4a139f395a | linux | amd64 | 1.24.4 | 2 | 2 |
| **translations-1/b-linux-v100-gpu-d2g-4-1tb-standard** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 2 | 2 |
| **translations-1/b-linux-v100-gpu-d2g-4-1tb-std-alpha** | generic-worker | 84.1.2 | multiuser | 4a139f395a | linux | amd64 | 1.24.4 | 2 | 2 |
| **translations-1/b-linux-v100-gpu-d2g-4-2tb** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 15 | 15 |
| **translations-1/b-linux-v100-gpu-d2g-4-2tb-alpha** | generic-worker | 84.1.2 | multiuser | 4a139f395a | linux | amd64 | 1.24.4 | 2 | 2 |
| **translations-1/b-linux-v100-gpu-d2g-4-300gb** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 74 | 74 |
| **translations-1/b-linux-v100-gpu-d2g-4-300gb-alpha** | generic-worker | 84.1.2 | multiuser | 4a139f395a | linux | amd64 | 1.24.4 | 2 | 2 |
| **translations-1/b-linux-v100-gpu-d2g-4-300gb-standard** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 2 | 2 |
| **translations-1/b-linux-v100-gpu-d2g-4-300gb-std-alpha** | generic-worker | 84.1.2 | multiuser | 4a139f395a | linux | amd64 | 1.24.4 | 2 | 2 |
| **translations-1/b-linux-v100-gpu-d2g-4-alpha** | generic-worker | 84.1.2 | multiuser | 4a139f395a | linux | amd64 | 1.24.4 | 2 | 2 |
| **translations-1/b-linux-v100-gpu-d2g-alpha** | generic-worker | 84.1.2 | multiuser | 4a139f395a | linux | amd64 | 1.24.4 | 2 | 2 |
| **translations-1/decision** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 1238 | 1238 |
| **translations-1/images** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 5 | 5 |
| **translations-t/t-linux-gw-noscratch-amd** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 2 | 2 |
| **xpi-1/b-linux** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 86 | 86 |
| **xpi-1/decision** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 48 | 48 |
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
| projects/taskcluster-imaging/global/images/docker-firefoxci-gcp-l1-googlecompute-2025-06-13t18-31-38z | 85 |
| projects/fxci-production-level3-workers/global/images/docker-firefoxci-gcp-l3-googlecompute-2024-02-05t23-18-22z | 52 |
| projects/taskcluster-imaging/global/images/docker-worker-gcp-u14-04-2025-06-16 | 19 |
| projects/taskcluster-imaging/global/images/docker-firefoxci-gcp-lt-googlecompute-2023-04-13t21-30-28z | 1 |


| Worker Pool | Implementation | Version | Total Workers | Total Capacity |
| --- | --- | --- | ---: | ---: |
| **app-services-1/b-linux-amd** | docker-worker | 38.0.5 | 3 | 3 |
| **app-services-1/b-linux-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **app-services-1/decision-gcp** | docker-worker | 38.0.5 | 805 | 1610 |
| **app-services-1/images-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **app-services-3/b-linux-amd** | docker-worker | 38.0.5 | 2 | 2 |
| **app-services-3/b-linux-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **app-services-3/decision-gcp** | docker-worker | 38.0.5 | 2 | 4 |
| **app-services-3/images-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **code-analysis-1/linux-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **code-analysis-3/linux-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **comm-1/b-linux-amd** | docker-worker | 38.0.5 | 2 | 2 |
| **comm-1/b-linux-gcp** | docker-worker | 38.0.5 | 32 | 32 |
| **comm-1/b-linux-large-amd** | docker-worker | 38.0.5 | 2 | 2 |
| **comm-1/b-linux-large-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **comm-1/b-linux-xlarge-amd** | docker-worker | 38.0.5 | 2 | 2 |
| **comm-1/b-linux-xlarge-gcp** | docker-worker | 38.0.5 | 3 | 3 |
| **comm-1/decision-gcp** | docker-worker | 38.0.5 | 9 | 9 |
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
| **comm-3/b-linux-gcp** | docker-worker | 38.0.5 | 349 | 349 |
| **comm-3/b-linux-large-amd** | docker-worker | 38.0.5 | 2 | 2 |
| **comm-3/b-linux-large-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **comm-3/b-linux-xlarge-amd** | docker-worker | 38.0.5 | 2 | 2 |
| **comm-3/b-linux-xlarge-gcp** | docker-worker | 38.0.5 | 3 | 3 |
| **comm-3/decision-gcp** | docker-worker | 38.0.5 | 16 | 16 |
| **comm-3/images-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **comm-t/misc-gcp** | docker-worker | 38.0.5 | 2 | 16 |
| **comm-t/t-linux-large-gcp** | docker-worker | unknown version | 2 | 2 |
| **comm-t/t-linux-large-noscratch-gcp** | docker-worker | unknown version | 2 | 2 |
| **comm-t/t-linux-xlarge-gcp** | docker-worker | unknown version | 2 | 2 |
| **comm-t/t-linux-xlarge-noscratch-gcp** | docker-worker | unknown version | 2 | 2 |
| **comm-t/t-linux-xlarge-source-gcp** | docker-worker | unknown version | 2 | 2 |
| **comm-t/t-linux-xlarge-source-noscratch-gcp** | docker-worker | unknown version | 2 | 2 |
| **gecko-1/b-linux-amd** | docker-worker | 38.0.5 | 82 | 82 |
| **gecko-1/b-linux-gcp** | docker-worker | 38.0.5 | 57 | 57 |
| **gecko-1/b-linux-gcp-relops1411** | docker-worker | 38.0.5 | 2 | 2 |
| **gecko-1/b-linux-gcp-test-bug-1882320** | docker-worker | 38.0.5 | 2 | 2 |
| **gecko-1/b-linux-kvm-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **gecko-1/b-linux-large-amd** | docker-worker | 38.0.5 | 3 | 3 |
| **gecko-1/b-linux-large-gcp** | docker-worker | 38.0.5 | 7 | 7 |
| **gecko-1/b-linux-medium-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **gecko-1/b-linux-xlarge-amd** | docker-worker | 38.0.5 | 3 | 3 |
| **gecko-1/b-linux-xlarge-gcp** | docker-worker | 38.0.5 | 5 | 5 |
| **gecko-1/b-linux-xlarge-gcp-bug1797804-c2** | docker-worker | 38.0.5 | 2 | 2 |
| **gecko-1/decision-gcp** | docker-worker | 38.0.5 | 13 | 13 |
| **gecko-1/decision-gcp-c3d4-amd-bug1990935** | docker-worker | 38.0.5 | 2 | 2 |
| **gecko-1/decision-gcp-c3d8-amd-bug1990935** | docker-worker | 38.0.5 | 4 | 4 |
| **gecko-1/decision-gcp-c4d4-amd-bug1990935** | docker-worker | 38.0.5 | 2 | 2 |
| **gecko-1/decision-gcp-c4d4-hcpu-amd-bug1990935** | docker-worker | 38.0.5 | 2 | 2 |
| **gecko-1/decision-gcp-c4d8-amd-bug1990935** | docker-worker | 38.0.5 | 2 | 2 |
| **gecko-1/decision-gcp-c4d8-hcpu-amd-bug1990935** | docker-worker | 38.0.5 | 2 | 2 |
| **gecko-1/decision-gcp-c4d8-lssd-amd-bug1990935** | docker-worker | 38.0.5 | 2 | 2 |
| **gecko-1/images-gcp** | docker-worker | 38.0.5 | 26 | 26 |
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
| **gecko-3/b-linux-amd** | docker-worker | 38.0.5 | 55 | 55 |
| **gecko-3/b-linux-gcp** | docker-worker | 38.0.5 | 952 | 952 |
| **gecko-3/b-linux-kvm-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **gecko-3/b-linux-large-amd** | docker-worker | 38.0.5 | 2 | 2 |
| **gecko-3/b-linux-large-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **gecko-3/b-linux-medium-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **gecko-3/b-linux-xlarge-amd** | docker-worker | 38.0.5 | 2 | 2 |
| **gecko-3/b-linux-xlarge-gcp** | docker-worker | 38.0.5 | 7 | 7 |
| **gecko-3/decision-gcp** | docker-worker | 38.0.5 | 845 | 845 |
| **gecko-3/images-gcp** | docker-worker | 38.0.5 | 3 | 3 |
| **gecko-t/misc-gcp** | docker-worker | 38.0.5 | 30 | 240 |
| **gecko-t/t-linux-kvm-gcp** | docker-worker | unknown version | 2 | 2 |
| **gecko-t/t-linux-kvm-gcp-relops1411** | docker-worker | unknown version | 2 | 2 |
| **gecko-t/t-linux-kvm-noscratch-gcp** | docker-worker | unknown version | 17 | 17 |
| **gecko-t/t-linux-large-gcp** | docker-worker | unknown version | 2 | 2 |
| **gecko-t/t-linux-large-hostub2204-gcp** | docker-worker | 48.3.0 | 2 | 2 |
| **gecko-t/t-linux-large-noscratch-gcp** | docker-worker | unknown version | 4217 | 4217 |
| **gecko-t/t-linux-xlarge-gcp** | docker-worker | unknown version | 56 | 56 |
| **gecko-t/t-linux-xlarge-noscratch-gcp** | docker-worker | unknown version | 2693 | 2693 |
| **gecko-t/t-linux-xlarge-source-gcp** | docker-worker | unknown version | 219 | 219 |
| **gecko-t/t-linux-xlarge-source-noscratch-gcp** | docker-worker | unknown version | 2 | 2 |
| **glean-1/b-linux-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **glean-1/decision-gcp** | docker-worker | 38.0.5 | 791 | 1582 |
| **glean-1/images-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **glean-3/b-linux-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **glean-3/decision-gcp** | docker-worker | 38.0.5 | 2 | 4 |
| **glean-3/images-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **infra/build-decision** | docker-worker | 38.0.5 | 5 | 40 |
| **mobile-1/b-linux-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **mobile-1/decision-gcp** | docker-worker | 38.0.5 | 789 | 1578 |
| **mobile-1/images-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **mobile-3/b-linux-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **mobile-3/decision-gcp** | docker-worker | 38.0.5 | 2 | 4 |
| **mobile-3/images-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **mozilla-1/b-linux-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **mozilla-1/decision-gcp** | docker-worker | 38.0.5 | 787 | 1574 |
| **mozilla-1/images-gcp** | docker-worker | 38.0.5 | 9 | 9 |
| **mozilla-3/b-linux-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **mozilla-3/decision-gcp** | docker-worker | 38.0.5 | 2 | 4 |
| **mozilla-3/images-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **mozilla-t/t-linux-large-gcp** | docker-worker | unknown version | 20 | 20 |
| **mozilla-t/t-linux-large-noscratch-gcp** | docker-worker | unknown version | 2 | 2 |
| **mozillavpn-1/b-linux-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **mozillavpn-1/b-linux-large-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **mozillavpn-1/decision-gcp** | docker-worker | 38.0.5 | 792 | 1584 |
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
| **translations-1/decision-gcp** | docker-worker | 38.0.5 | 1760 | 3520 |
| **translations-1/images-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **xpi-1/b-linux-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **xpi-1/decision-gcp** | docker-worker | 38.0.5 | 2 | 4 |
| **xpi-1/images-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **xpi-3/b-linux-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **xpi-3/decision-gcp** | docker-worker | 38.0.5 | 2 | 4 |
| **xpi-3/images-gcp** | docker-worker | 38.0.5 | 2 | 2 |


## Script Worker

Total: `34`



| Worker Pool | Implementation | Version | Total Workers | Total Capacity |
| --- | --- | --- | ---: | ---: |
| **scriptworker-k8s/adhoc-t-signing** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/app-services-3-beetmover** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/app-services-3-signing** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/comm-3-balrog** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/comm-3-beetmover** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/comm-3-signing** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/comm-3-tree** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/comm-t-signing** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/enterprise-t-signing** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/gecko-1-beetmover** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
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
| **scriptworker-k8s/mozillavpn-3-signing** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
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
|  | 2 |
| projects/taskcluster-imaging/global/images/generic-2204-wayland-vm-gcp-googlecompute-2023-09-22t17-39-37z | 2 |


| Worker Pool | Implementation | Version | Total Workers | Total Capacity |
| --- | --- | --- | ---: | ---: |
| **built-in/fail** |  | No artifacts found | 0 | 0 |
| **built-in/succeed** |  | No artifacts found | 0 | 0 |
| **gecko-t/t-linux-vm-2204-wayland** |  | No artifacts found | 61 | 61 |
| **gecko-t/t-linux-vm-2204-wayland-snap** |  | No artifacts found | 2 | 2 |


## Version not determined [^2]

Total: `20`


Count by image:

| Version | Count |
| :--- | ---: |
|  | 7 |
| unknown | 11 |
| /subscriptions/108d46d5-fe9b-4850-9a7d-8c914aa6c1f0/resourceGroups/rg-packer-worker-images/providers/Microsoft.Compute/galleries/win11_a64_24h2_tester_alpha/images/win11_a64_24h2_tester_alpha/versions/1.0.0 | 1 |
| projects/taskcluster-imaging/global/images/gw-fxci-gcp-l1-2404-amd64-googlecompute-alpha | 1 |


| Worker Pool | Implementation | Version | Total Workers | Total Capacity |
| --- | --- | --- | ---: | ---: |
| **comm-t/win11-a64-24h2-alpha** |  | Version not determined; task not (yet) claimed | 21 | 21 |
| **enterprise-t/win11-64-24h2-alpha** |  | Version not determined; task not (yet) claimed | 5 | 5 |
| **enterprise-t/win11-64-24h2-source-alpha** |  | Version not determined; task not (yet) claimed | 5532 | 5532 |
| **gecko-t/t-linux-2404-relsre** |  | Version not determined; task not (yet) claimed | 3 | 3 |
| **gecko-t/win10-64-2009-alpha** |  | Version not determined; task not (yet) claimed | 306 | 306 |
| **gecko-t/win10-64-2009-gpu-alpha** |  | Version not determined; task not (yet) claimed | 250 | 250 |
| **gecko-t/win10-64-2009-source-alpha** |  | Version not determined; task not (yet) claimed | 285 | 285 |
| **gecko-t/win10-64-2009-webgpu-alpha** |  | Version not determined; task not (yet) claimed | 276 | 276 |
| **gecko-t/win11-64-24h2-alpha** |  | Version not determined; task not (yet) claimed | 2169 | 2169 |
| **gecko-t/win11-64-24h2-gpu-alpha** |  | Version not determined; task not (yet) claimed | 186 | 186 |
| **gecko-t/win11-64-24h2-large-alpha** |  | Version not determined; task not (yet) claimed | 6 | 6 |
| **gecko-t/win11-64-24h2-source-alpha** |  | Version not determined; task not (yet) claimed | 16 | 16 |
| **gecko-t/win11-64-24h2-webgpu-alpha** |  | Version not determined; task not (yet) claimed | 4 | 4 |
| **proj-autophone/gecko-t-bitbar-gw-perf-p6** |  | Version not determined; task not (yet) claimed | 0 | 0 |
| **proj-autophone/gecko-t-bitbar-gw-perf-s24** |  | Version not determined; task not (yet) claimed | 0 | 0 |
| **releng-hardware/gecko-3-b-osx-1015** |  | Version not determined; task not (yet) claimed | 0 | 0 |
| **releng-hardware/gecko-t-linux-netperf-1804** |  | Version not determined; task not (yet) claimed | 0 | 0 |
| **releng-hardware/gecko-t-osx-1015-r8** |  | Version not determined; task not (yet) claimed | 0 | 0 |
| **releng-hardware/gecko-t-osx-1400-r8** |  | Version not determined; task not (yet) claimed | 0 | 0 |
| **releng-hardware/win11-64-24h2-hw-relops1213** |  | Version not determined; task not (yet) claimed | 0 | 0 |



[^1]: Those are the pools whose tasks were claimed and resolved by a worker as expected, but the worker did not publish either artifact `public/logs/live_backing.log` nor `public/logs/chain_of_trust.log`, which is the source used to identify the worker implementation.

[^2]: Probing task remains pending after two hours. Those are the pools that were not able to start any worker to claim the task within two hours.

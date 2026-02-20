

# Worker Pool Versions


## Generic Worker

Total: `359`

Count by version:

| Version | Count |
| :--- | ---: |
| 36.0.0 | 4 |
| 45.0.0 | 1 |
| 60.3.4 | 18 |
| 61.0.0 | 1 |
| 64.3.0 | 26 |
| 84.1.2 | 7 |
| 87.0.0 | 1 |
| 88.0.2 | 2 |
| 91.0.2 | 2 |
| 91.1.0 | 5 |
| 95.1.3 | 211 |
| 96.2.2 | 2 |
| 96.2.3 | 79 |


Count by image:

| Version | Count |
| :--- | ---: |
| /subscriptions/108d46d5-fe9b-4850-9a7d-8c914aa6c1f0/resourceGroups/rg-packer-worker-images/providers/Microsoft.Compute/galleries/win11_a64_24h2_tester_alpha/images/win11_a64_24h2_tester_alpha/versions/1.0.0 | 4 |
| projects/taskcluster-imaging/global/images/gw-fxci-gcp-l1-gui-googlecompute-2024-08-22t22-48-09z | 8 |
| projects/fxci-production-level3-workers/global/images/gw-fxci-gcp-l3-2404-amd64-headless-googlecompute-2026-01-14 | 58 |
| projects/fxci-production-level3-workers/global/images/gw-fxci-gcp-l3-arm64-gui-googlecompute-2024-09-18t19-02-58z | 4 |
| projects/taskcluster-imaging/global/images/gw-fxci-gcp-l1-2404-amd64-gui-googlecompute-alpha | 1 |
| projects/taskcluster-imaging/global/images/gw-fxci-gcp-l1-2404-arm64-headless-googlecompute-2026-01-14 | 14 |
| projects/taskcluster-imaging/global/images/gw-fxci-gcp-l1-gui-googlecompute-2025-01-13t22-33-40z | 2 |
| unknown | 61 |
| /subscriptions/108d46d5-fe9b-4850-9a7d-8c914aa6c1f0/resourceGroups/rg-packer-worker-images/providers/Microsoft.Compute/galleries/win11_a64_24h2_tester/images/win11_a64_24h2_tester/versions/1.2.0 | 4 |
| projects/taskcluster-imaging/global/images/gw-fxci-gcp-l1-2404-amd64-headless-googlecompute-alpha-tc | 7 |
| /subscriptions/a30e97ab-734a-4f3b-a0e4-c51c0bff0701/resourceGroups/rg-packer-worker-images/providers/Microsoft.Compute/galleries/trusted_win11_a64_24h2_builder/images/trusted_win11_a64_24h2_builder/versions/1.2.0 | 4 |
| projects/fxci-production-level3-workers/global/images/gw-fxci-gcp-l3-gui-googlecompute-2024-09-18t05-46-31z | 3 |
| projects/fxci-production-level3-workers/global/images/gw-fxci-gcp-l3-2404-arm64-headless-googlecompute-2026-01-14 | 6 |
| projects/taskcluster-imaging/global/images/gw-fxci-gcp-l1-arm64-gui-googlecompute-2024-09-18t14-57-52z | 9 |
| projects/taskcluster-imaging/global/images/gw-fxci-gcp-l1-2404-arm64-headless-googlecompute-alpha | 1 |
| /subscriptions/108d46d5-fe9b-4850-9a7d-8c914aa6c1f0/resourceGroups/rg-packer-worker-images/providers/Microsoft.Compute/galleries/win11_a64_24h2_builder/images/win11_a64_24h2_builder/versions/1.2.0 | 4 |
|  | 36 |
| projects/taskcluster-imaging/global/images/gw-fxci-gcp-l1-2404-amd64-headless-googlecompute-2026-01-14 | 120 |
| /subscriptions/108d46d5-fe9b-4850-9a7d-8c914aa6c1f0/resourceGroups/rg-packer-worker-images/providers/Microsoft.Compute/galleries/win11_a64_24h2_builder_alpha/images/win11_a64_24h2_builder_alpha/versions/1.0.0 | 2 |
| projects/taskcluster-imaging/global/images/gw-fxci-gcp-l1-2404-amd64-headless-googlecompute-alpha | 6 |
| projects/taskcluster-imaging/global/images/gw-fxci-gcp-l1-2404-amd64-gui-googlecompute-2026-01-14 | 4 |
| projects/fxci-production-level3-workers/global/images/gw-fxci-gcp-l3-2404-amd64-headless-googlecompute-alpha | 1 |


| Worker Pool | Implementation | Version | Engine | Revision | OS | Arch | GO | Total Workers | Total Capacity |
| --- | --- | --- | --- | --- | --- | --- | --- | ---: | ---: |
| **adhoc-1/decision** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 15 | 15 |
| **adhoc-1/images** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 64 | 64 |
| **adhoc-3/decision** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 10 | 10 |
| **adhoc-3/images** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 7 | 7 |
| **app-services-1/b-linux** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 302 | 302 |
| **app-services-1/b-linux-docker-amd** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 2 | 2 |
| **app-services-1/decision** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 1261 | 1261 |
| **app-services-1/images** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 2 | 2 |
| **app-services-3/b-linux** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 642 | 642 |
| **app-services-3/b-linux-docker-amd** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 2 | 2 |
| **app-services-3/decision** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 22 | 22 |
| **app-services-3/images** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 2 | 2 |
| **code-analysis-1/linux-docker** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 2 | 2 |
| **code-analysis-1/linux-gw-gcp** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | amd64 | 1.22.2 | 60 | 60 |
| **code-analysis-3/linux-docker** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 2 | 2 |
| **code-analysis-3/linux-gw-gcp** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | amd64 | 1.22.2 | 16 | 16 |
| **code-coverage/bot** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 223 | 223 |
| **code-review/bot** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 6750 | 6750 |
| **comm-1/b-linux** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 781 | 781 |
| **comm-1/b-linux-aarch64** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | arm64 | 1.25.5 | 2 | 2 |
| **comm-1/b-linux-docker-amd** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 4 | 4 |
| **comm-1/b-linux-docker-large-amd** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 2 | 2 |
| **comm-1/b-linux-docker-xlarge-amd** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 2 | 2 |
| **comm-1/b-linux-gcp-aarch64** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | arm64 | 1.22.2 | 2 | 2 |
| **comm-1/b-linux-large** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 2 | 2 |
| **comm-1/b-linux-xlarge** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 2 | 2 |
| **comm-1/b-win2022** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 79 | 79 |
| **comm-1/decision** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 204 | 204 |
| **comm-1/images** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 2 | 2 |
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
| **comm-3/b-linux** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 2584 | 2584 |
| **comm-3/b-linux-aarch64** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | arm64 | 1.25.5 | 3 | 3 |
| **comm-3/b-linux-docker-amd** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 1751 | 1751 |
| **comm-3/b-linux-docker-large-amd** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 3 | 3 |
| **comm-3/b-linux-docker-xlarge-amd** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 4 | 4 |
| **comm-3/b-linux-gcp-aarch64** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | arm64 | 1.22.2 | 2 | 2 |
| **comm-3/b-linux-large** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 2 | 2 |
| **comm-3/b-linux-xlarge** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 2 | 2 |
| **comm-3/b-win2022** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 722 | 722 |
| **comm-3/decision** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 87 | 87 |
| **comm-3/images** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 3 | 3 |
| **comm-3/images-aarch64** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | arm64 | 1.25.5 | 2 | 2 |
| **comm-3/images-gcp-aarch64** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | arm64 | 1.22.2 | 2 | 2 |
| **comm-t/misc** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 2 | 2 |
| **comm-t/t-linux-arm64-docker** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | arm64 | 1.25.5 | 2 | 2 |
| **comm-t/t-linux-docker** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 17 | 17 |
| **comm-t/t-linux-docker-amd** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 2212 | 2212 |
| **comm-t/t-linux-docker-noscratch** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 73 | 73 |
| **comm-t/t-linux-docker-noscratch-amd** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 3720 | 3720 |
| **comm-t/win11-64-2009** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 2 | 2 |
| **comm-t/win11-64-2009-ssd** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 2 | 2 |
| **comm-t/win11-64-24h2** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 6086 | 6086 |
| **comm-t/win11-64-24h2-ssd** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 458 | 458 |
| **comm-t/win11-a64-24h2** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | arm64 | 1.25.7 | 2 | 2 |
| **comm-t/win11-a64-24h2-alpha** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | arm64 | 1.25.7 | 15 | 15 |
| **enterprise-1/b-linux** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 648 | 648 |
| **enterprise-1/b-linux-aarch64** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | arm64 | 1.25.5 | 3 | 3 |
| **enterprise-1/b-linux-docker-amd** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 297 | 297 |
| **enterprise-1/b-linux-docker-large-amd** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 7 | 7 |
| **enterprise-1/b-linux-docker-xlarge-amd** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 407 | 407 |
| **enterprise-1/b-win2022** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 165 | 165 |
| **enterprise-1/decision** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 1246 | 1246 |
| **enterprise-1/images** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 7 | 7 |
| **enterprise-1/images-aarch64** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | arm64 | 1.25.5 | 2 | 2 |
| **enterprise-1/win11-a64-24h2-builder** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | arm64 | 1.25.7 | 2 | 2 |
| **enterprise-3/b-linux** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 637 | 637 |
| **enterprise-3/b-linux-aarch64** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | arm64 | 1.25.5 | 3 | 3 |
| **enterprise-3/b-linux-docker-amd** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 285 | 285 |
| **enterprise-3/b-linux-docker-large-amd** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 31 | 31 |
| **enterprise-3/b-linux-docker-xlarge-amd** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 285 | 285 |
| **enterprise-3/b-win2022** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 176 | 176 |
| **enterprise-3/decision** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 33 | 33 |
| **enterprise-3/images** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 6 | 6 |
| **enterprise-3/images-aarch64** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | arm64 | 1.25.5 | 2 | 2 |
| **enterprise-3/win11-a64-24h2-builder** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | arm64 | 1.25.7 | 18 | 18 |
| **enterprise-t/misc** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 67 | 67 |
| **enterprise-t/t-linux-docker-16c32gb-amd** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 482 | 482 |
| **enterprise-t/t-linux-docker-amd** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 1901 | 1901 |
| **enterprise-t/t-linux-docker-noscratch-amd** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 2409 | 2409 |
| **enterprise-t/win11-64-24h2** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 1653 | 1653 |
| **enterprise-t/win11-64-24h2-alpha** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 7 | 7 |
| **enterprise-t/win11-64-24h2-source** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 3 | 3 |
| **enterprise-t/win11-64-24h2-source-alpha** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 8 | 8 |
| **enterprise-t/win11-64-25h2-alpha** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 1 | 1 |
| **enterprise-t/win11-64-25h2-source-alpha** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 1 | 1 |
| **gecko-1/b-linux** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 6318 | 6318 |
| **gecko-1/b-linux-2204-kvm-gcp** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | amd64 | 1.22.2 | 2 | 2 |
| **gecko-1/b-linux-aarch64** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | arm64 | 1.25.5 | 6 | 6 |
| **gecko-1/b-linux-docker-alpha** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 2 | 2 |
| **gecko-1/b-linux-docker-amd** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 9217 | 9217 |
| **gecko-1/b-linux-docker-large-amd** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 184 | 184 |
| **gecko-1/b-linux-docker-xlarge-amd** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 1131 | 1131 |
| **gecko-1/b-linux-gcp-aarch64** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | arm64 | 1.22.2 | 2 | 2 |
| **gecko-1/b-linux-kvm** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 303 | 303 |
| **gecko-1/b-linux-large** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 2 | 2 |
| **gecko-1/b-linux-medium** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 20704 | 20704 |
| **gecko-1/b-linux-xlarge** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 10 | 10 |
| **gecko-1/b-win2022** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 935 | 935 |
| **gecko-1/b-win2022-alpha** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 12 | 12 |
| **gecko-1/b-win2022-xxlarge** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 2 | 2 |
| **gecko-1/decision** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 4884 | 4884 |
| **gecko-1/images** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 26 | 26 |
| **gecko-1/images-aarch64** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | arm64 | 1.25.5 | 2 | 2 |
| **gecko-1/images-gcp-aarch64** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | arm64 | 1.22.2 | 2 | 2 |
| **gecko-1/win11-a64-24h2-builder** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | arm64 | 1.25.7 | 25 | 25 |
| **gecko-1/win11-a64-24h2-builder-alpha** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | arm64 | 1.25.7 | 3 | 3 |
| **gecko-2/b-linux** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 133 | 133 |
| **gecko-2/b-linux-aarch64** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | arm64 | 1.25.5 | 3 | 3 |
| **gecko-2/b-linux-docker-amd** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 117 | 117 |
| **gecko-2/b-linux-docker-large-amd** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 16 | 16 |
| **gecko-2/b-linux-docker-xlarge-amd** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 29 | 29 |
| **gecko-2/b-linux-gcp-aarch64** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | arm64 | 1.22.2 | 3 | 3 |
| **gecko-2/b-linux-kvm** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 12 | 12 |
| **gecko-2/b-linux-large** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 2 | 2 |
| **gecko-2/b-linux-medium** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 179 | 179 |
| **gecko-2/b-linux-xlarge** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 2 | 2 |
| **gecko-2/b-win2022** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 13 | 13 |
| **gecko-2/decision** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 13 | 13 |
| **gecko-2/images** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 2 | 2 |
| **gecko-2/images-aarch64** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | arm64 | 1.25.5 | 2 | 2 |
| **gecko-2/images-gcp-aarch64** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | arm64 | 1.22.2 | 2 | 2 |
| **gecko-3/b-linux** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 21225 | 21225 |
| **gecko-3/b-linux-2204-kvm-gcp** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | amd64 | 1.22.2 | 2 | 2 |
| **gecko-3/b-linux-aarch64** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | arm64 | 1.25.5 | 2 | 2 |
| **gecko-3/b-linux-docker-alpha** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 2 | 2 |
| **gecko-3/b-linux-docker-amd** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 23148 | 23148 |
| **gecko-3/b-linux-docker-large-amd** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 338 | 338 |
| **gecko-3/b-linux-docker-xlarge-amd** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 3089 | 3089 |
| **gecko-3/b-linux-gcp-aarch64** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | arm64 | 1.22.2 | 2 | 2 |
| **gecko-3/b-linux-kvm** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 602 | 602 |
| **gecko-3/b-linux-large** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 2 | 2 |
| **gecko-3/b-linux-medium** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 22370 | 22370 |
| **gecko-3/b-linux-xlarge** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 64 | 64 |
| **gecko-3/b-win2022** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 6499 | 6499 |
| **gecko-3/b-win2022-xxlarge** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 12 | 12 |
| **gecko-3/decision** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 5340 | 5340 |
| **gecko-3/images** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 17 | 17 |
| **gecko-3/images-aarch64** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | arm64 | 1.25.5 | 2 | 2 |
| **gecko-3/images-gcp-aarch64** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | arm64 | 1.22.2 | 2 | 2 |
| **gecko-3/win11-a64-24h2-builder** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | arm64 | 1.25.7 | 50 | 50 |
| **gecko-t/misc** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 1040 | 1040 |
| **gecko-t/t-linux-2204-wayland** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | amd64 | 1.22.2 | 8866 | 8866 |
| **gecko-t/t-linux-2204-wayland-arm64-relsre** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | arm64 | 1.22.2 | 3 | 3 |
| **gecko-t/t-linux-2204-wayland-relsre** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | amd64 | 1.22.2 | 3 | 3 |
| **gecko-t/t-linux-2204-wayland-root-exp** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | amd64 | 1.22.2 | 2 | 2 |
| **gecko-t/t-linux-2204-wayland-snap** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | amd64 | 1.22.2 | 128 | 128 |
| **gecko-t/t-linux-2404-headless-alpha** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 3 | 3 |
| **gecko-t/t-linux-2404-headless-arm64-alpha** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | arm64 | 1.25.5 | 3 | 3 |
| **gecko-t/t-linux-2404-headless-ssd-alpha** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 2 | 2 |
| **gecko-t/t-linux-2404-wayland** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 2 | 2 |
| **gecko-t/t-linux-2404-wayland-relsre** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 2 | 2 |
| **gecko-t/t-linux-2404-wayland-snap** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 125 | 125 |
| **gecko-t/t-linux-arm64-docker** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | arm64 | 1.25.5 | 27 | 27 |
| **gecko-t/t-linux-docker** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 1897 | 1897 |
| **gecko-t/t-linux-docker-16c32gb-amd** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 3903 | 3903 |
| **gecko-t/t-linux-docker-amd** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 30669 | 30669 |
| **gecko-t/t-linux-docker-amd-alpha** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 2 | 2 |
| **gecko-t/t-linux-docker-kvm** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 74442 | 74442 |
| **gecko-t/t-linux-docker-kvm-alpha** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 2 | 2 |
| **gecko-t/t-linux-docker-noscratch** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 5905 | 5905 |
| **gecko-t/t-linux-docker-noscratch-amd** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 119172 | 119172 |
| **gecko-t/t-linux-docker-noscratch-amd-alpha** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 2 | 2 |
| **gecko-t/t-linux-xlarge-2204-wayland** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | amd64 | 1.22.2 | 11422 | 11422 |
| **gecko-t/t-linux-xlarge-2404-wayland** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 3 | 3 |
| **gecko-t/win10-64-2009** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 4696 | 4696 |
| **gecko-t/win10-64-2009-gpu** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 480 | 480 |
| **gecko-t/win10-64-2009-source** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 56 | 56 |
| **gecko-t/win10-64-2009-webgpu** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 2 | 2 |
| **gecko-t/win11-64-2009** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 2246 | 2246 |
| **gecko-t/win11-64-2009-alpha** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 2 | 2 |
| **gecko-t/win11-64-2009-gpu** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 1431 | 1431 |
| **gecko-t/win11-64-2009-gpu-alpha** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 2 | 2 |
| **gecko-t/win11-64-2009-source** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 2 | 2 |
| **gecko-t/win11-64-2009-source-alpha** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 2 | 2 |
| **gecko-t/win11-64-2009-ssd** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 2 | 2 |
| **gecko-t/win11-64-2009-ssd-gpu** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 2 | 2 |
| **gecko-t/win11-64-2009-webgpu** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 2 | 2 |
| **gecko-t/win11-64-2009-webgpu-alpha** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 3 | 3 |
| **gecko-t/win11-64-24h2** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 47188 | 47188 |
| **gecko-t/win11-64-24h2-alpha** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 1412 | 1412 |
| **gecko-t/win11-64-24h2-amd** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 24442 | 24442 |
| **gecko-t/win11-64-24h2-gpu** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 17427 | 17427 |
| **gecko-t/win11-64-24h2-gpu-alpha** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 461 | 461 |
| **gecko-t/win11-64-24h2-large** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 2005 | 2005 |
| **gecko-t/win11-64-24h2-large-alpha** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 11 | 11 |
| **gecko-t/win11-64-24h2-source** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 2045 | 2045 |
| **gecko-t/win11-64-24h2-source-alpha** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 13 | 13 |
| **gecko-t/win11-64-24h2-ssd** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 2 | 2 |
| **gecko-t/win11-64-24h2-ssd-gpu** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 3 | 3 |
| **gecko-t/win11-64-24h2-webgpu** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 3990 | 3990 |
| **gecko-t/win11-64-24h2-webgpu-alpha** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 8 | 8 |
| **gecko-t/win11-64-25h2-alpha** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 1 | 1 |
| **gecko-t/win11-64-25h2-gpu-alpha** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 1 | 1 |
| **gecko-t/win11-64-25h2-large-alpha** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 3 | 3 |
| **gecko-t/win11-64-25h2-source-alpha** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 1 | 1 |
| **gecko-t/win11-64-25h2-webgpu-alpha** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 1 | 1 |
| **gecko-t/win11-a64-24h2** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | arm64 | 1.25.7 | 173 | 173 |
| **gecko-t/win11-a64-24h2-alpha** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | arm64 | 1.25.7 | 12 | 12 |
| **glean-1/b-linux** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 29 | 29 |
| **glean-1/decision** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 1230 | 1230 |
| **glean-1/images** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 2 | 2 |
| **glean-3/b-linux** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 32 | 32 |
| **glean-3/decision** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 9 | 9 |
| **glean-3/images** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 2 | 2 |
| **mobile-1/b-linux** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 129 | 129 |
| **mobile-1/decision** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 1201 | 1201 |
| **mobile-1/images** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 2 | 2 |
| **mobile-3/b-linux** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 216 | 216 |
| **mobile-3/decision** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 130 | 130 |
| **mobile-3/images** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 2 | 2 |
| **mozilla-1/b-linux** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 2 | 2 |
| **mozilla-1/b-win2022** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 2 | 2 |
| **mozilla-1/decision** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 1225 | 1225 |
| **mozilla-1/images** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 2 | 2 |
| **mozilla-3/b-linux** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 2 | 2 |
| **mozilla-3/b-win2022** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 3 | 3 |
| **mozilla-3/decision** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 429 | 429 |
| **mozilla-3/images** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 2 | 2 |
| **mozilla-t/t-linux-2204-wayland** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | amd64 | 1.22.2 | 602 | 602 |
| **mozilla-t/t-linux-2204-wayland-relsre** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | amd64 | 1.22.2 | 2 | 2 |
| **mozilla-t/t-linux-2404-wayland** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 3 | 3 |
| **mozilla-t/t-linux-arm64-docker** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | arm64 | 1.25.5 | 2 | 2 |
| **mozilla-t/t-linux-docker** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 36 | 36 |
| **mozilla-t/t-linux-docker-amd** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 2 | 2 |
| **mozilla-t/t-linux-docker-noscratch** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 21 | 21 |
| **mozilla-t/t-linux-docker-noscratch-amd** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 2 | 2 |
| **mozillavpn-1/b-linux** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 445 | 445 |
| **mozillavpn-1/b-linux-gcp-gw** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | amd64 | 1.22.2 | 2 | 2 |
| **mozillavpn-1/b-linux-large** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 321 | 321 |
| **mozillavpn-1/b-win2022** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 111 | 111 |
| **mozillavpn-1/b-win2022-alpha** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 2 | 2 |
| **mozillavpn-1/decision** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 1227 | 1227 |
| **mozillavpn-1/images** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 39 | 39 |
| **mozillavpn-1/win11-a64-24h2-builder** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | arm64 | 1.25.7 | 2 | 2 |
| **mozillavpn-3/b-linux** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 168 | 168 |
| **mozillavpn-3/b-linux-gcp-gw** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | amd64 | 1.22.2 | 2 | 2 |
| **mozillavpn-3/b-linux-large** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 164 | 164 |
| **mozillavpn-3/b-win2022** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 52 | 52 |
| **mozillavpn-3/decision** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 22 | 22 |
| **mozillavpn-3/images** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 9 | 9 |
| **mozillavpn-3/win11-a64-24h2-builder** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | arm64 | 1.25.7 | 2 | 2 |
| **nss-1/b-win2022** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 89 | 89 |
| **nss-1/b-win2022-alpha** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 3 | 3 |
| **nss-1/decision** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 10 | 10 |
| **nss-1/images** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 2 | 2 |
| **nss-1/linux-docker** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 356 | 356 |
| **nss-1/win11-a64-24h2** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | arm64 | 1.25.7 | 2 | 2 |
| **nss-1/win11-a64-24h2-alpha** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | arm64 | 1.25.7 | 13 | 13 |
| **nss-1/win11-a64-24h2-builder** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | arm64 | 1.25.7 | 2 | 2 |
| **nss-1/win11-a64-24h2-builder-alpha** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | arm64 | 1.25.7 | 2 | 2 |
| **nss-3/b-win2022** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 169 | 169 |
| **nss-3/decision** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 11 | 11 |
| **nss-3/images** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 2 | 2 |
| **nss-3/linux-docker** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 613 | 613 |
| **nss-3/win11-a64-24h2-builder** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | arm64 | 1.25.7 | 2 | 2 |
| **nss-t/t-linux-arm64-docker** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | arm64 | 1.25.5 | 2 | 2 |
| **nss-t/t-linux-docker** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 1114 | 1114 |
| **nss-t/t-linux-docker-amd** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 2 | 2 |
| **nss-t/win11-a64-24h2** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | arm64 | 1.25.7 | 2 | 2 |
| **nss-t/win11-a64-24h2-alpha** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | arm64 | 1.25.7 | 4 | 4 |
| **proj-autophone/gecko-t-bitbar-gw-perf-a55** | generic-worker | 36.0.0 | simple | b9cb11293f | linux | amd64 | 1.13.7 | 0 | 0 |
| **proj-autophone/gecko-t-bitbar-gw-perf-p6** | generic-worker | 36.0.0 | simple | b9cb11293f | linux | amd64 | 1.13.7 | 0 | 0 |
| **proj-autophone/gecko-t-bitbar-gw-perf-s24** | generic-worker | 36.0.0 | simple | b9cb11293f | linux | amd64 | 1.13.7 | 0 | 0 |
| **proj-autophone/gecko-t-bitbar-gw-unit-p5** | generic-worker | 36.0.0 | simple | b9cb11293f | linux | amd64 | 1.13.7 | 0 | 0 |
| **proj-autophone/gecko-t-lambda-perf-a55** | generic-worker | 87.0.0 | insecure | 99a1fcafbb | linux | amd64 | 1.24.5 | 0 | 0 |
| **releng-1/decision** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 99 | 99 |
| **releng-1/linux-docker** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 114 | 114 |
| **releng-3/decision** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 49 | 49 |
| **releng-3/linux-docker** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 80 | 80 |
| **releng-hardware/applicationservices-b-1-osx1015** | generic-worker | 60.3.4 | multiuser | 943a6f2b0d | darwin | amd64 | 1.22.0 | 0 | 0 |
| **releng-hardware/applicationservices-b-3-osx1015** | generic-worker | 60.3.4 | multiuser | 943a6f2b0d | darwin | amd64 | 1.22.0 | 0 | 0 |
| **releng-hardware/enterprise-1-b-osx-arm64** | generic-worker | 60.3.4 | multiuser | 943a6f2b0d | darwin | arm64 | 1.22.0 | 0 | 0 |
| **releng-hardware/enterprise-3-b-osx-arm64** | generic-worker | 60.3.4 | multiuser | 943a6f2b0d | darwin | arm64 | 1.22.0 | 0 | 0 |
| **releng-hardware/gecko-1-b-osx-1015** | generic-worker | 60.3.4 | multiuser | 943a6f2b0d | darwin | amd64 | 1.22.0 | 0 | 0 |
| **releng-hardware/gecko-1-b-osx-1015-staging** | generic-worker | 60.3.4 | multiuser | 943a6f2b0d | darwin | amd64 | 1.22.0 | 0 | 0 |
| **releng-hardware/gecko-1-b-osx-arm64** | generic-worker | 60.3.4 | multiuser | 943a6f2b0d | darwin | arm64 | 1.22.0 | 0 | 0 |
| **releng-hardware/gecko-3-b-osx-1015** | generic-worker | 60.3.4 | multiuser | 943a6f2b0d | darwin | amd64 | 1.22.0 | 0 | 0 |
| **releng-hardware/gecko-3-b-osx-arm64** | generic-worker | 60.3.4 | multiuser | 943a6f2b0d | darwin | arm64 | 1.22.0 | 0 | 0 |
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
| **releng-hardware/nss-3-b-osx-1015** | generic-worker | 60.3.4 | multiuser | 943a6f2b0d | darwin | amd64 | 1.22.0 | 0 | 0 |
| **releng-hardware/win11-64-24h2-hw** | generic-worker | 91.1.0 | multiuser | b8f34332c3 | windows | amd64 | 1.25.3 | 0 | 0 |
| **releng-hardware/win11-64-24h2-hw-alpha** | generic-worker | 96.2.2 | multiuser | 773509947e | windows | amd64 | 1.25.7 | 0 | 0 |
| **releng-hardware/win11-64-24h2-hw-perf-debug** | generic-worker | 91.1.0 | multiuser | b8f34332c3 | windows | amd64 | 1.25.3 | 0 | 0 |
| **releng-hardware/win11-64-24h2-hw-perf-sheriff** | generic-worker | 91.1.0 | multiuser | b8f34332c3 | windows | amd64 | 1.25.3 | 0 | 0 |
| **releng-hardware/win11-64-24h2-hw-ref** | generic-worker | 91.1.0 | multiuser | b8f34332c3 | windows | amd64 | 1.25.3 | 0 | 0 |
| **releng-hardware/win11-64-24h2-hw-ref-alpha** | generic-worker | 96.2.2 | multiuser | 773509947e | windows | amd64 | 1.25.7 | 0 | 0 |
| **releng-hardware/win11-64-24h2-hw-relops1213** | generic-worker | 91.1.0 | multiuser | b8f34332c3 | windows | amd64 | 1.25.3 | 0 | 0 |
| **releng-t/linux-docker** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 499 | 499 |
| **relops-1/decision** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 2 | 2 |
| **relops-1/images** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 2 | 2 |
| **relops-3/decision** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 34 | 34 |
| **relops-3/images** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 2 | 2 |
| **scriptworker-1/decision** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 17 | 17 |
| **scriptworker-1/images** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 103 | 103 |
| **scriptworker-3/decision** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 13 | 13 |
| **scriptworker-3/images** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 100 | 100 |
| **taskgraph-1/decision** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 36 | 36 |
| **taskgraph-1/images** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 21 | 21 |
| **taskgraph-3/decision** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 21 | 21 |
| **taskgraph-3/images** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 19 | 19 |
| **taskgraph-t/linux-docker** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 357 | 357 |
| **taskgraph-t/win11-64-24h2** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 38 | 38 |
| **translations-1/b-linux-large-gcp-1tb-32-256-d2g** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 23 | 23 |
| **translations-1/b-linux-large-gcp-1tb-32-256-std-d2g** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 22 | 22 |
| **translations-1/b-linux-large-gcp-1tb-64-512-d2g** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 2 | 2 |
| **translations-1/b-linux-large-gcp-1tb-64-512-std-d2g** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 2 | 2 |
| **translations-1/b-linux-large-gcp-d2g** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 271 | 271 |
| **translations-1/b-linux-large-gcp-d2g-1tb** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 2 | 2 |
| **translations-1/b-linux-large-gcp-d2g-1tb-standard** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 2 | 2 |
| **translations-1/b-linux-large-gcp-d2g-300gb** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 357 | 357 |
| **translations-1/b-linux-v100-gpu-d2g** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 176 | 176 |
| **translations-1/b-linux-v100-gpu-d2g-4** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 297 | 297 |
| **translations-1/b-linux-v100-gpu-d2g-4-1tb** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 100 | 100 |
| **translations-1/b-linux-v100-gpu-d2g-4-1tb-alpha** | generic-worker | 84.1.2 | multiuser | 4a139f395a | linux | amd64 | 1.24.4 | 2 | 2 |
| **translations-1/b-linux-v100-gpu-d2g-4-1tb-standard** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 2 | 2 |
| **translations-1/b-linux-v100-gpu-d2g-4-1tb-std-alpha** | generic-worker | 84.1.2 | multiuser | 4a139f395a | linux | amd64 | 1.24.4 | 2 | 2 |
| **translations-1/b-linux-v100-gpu-d2g-4-2tb** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 19 | 19 |
| **translations-1/b-linux-v100-gpu-d2g-4-2tb-alpha** | generic-worker | 84.1.2 | multiuser | 4a139f395a | linux | amd64 | 1.24.4 | 2 | 2 |
| **translations-1/b-linux-v100-gpu-d2g-4-300gb** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 123 | 123 |
| **translations-1/b-linux-v100-gpu-d2g-4-300gb-alpha** | generic-worker | 84.1.2 | multiuser | 4a139f395a | linux | amd64 | 1.24.4 | 2 | 2 |
| **translations-1/b-linux-v100-gpu-d2g-4-300gb-standard** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 2 | 2 |
| **translations-1/b-linux-v100-gpu-d2g-4-300gb-std-alpha** | generic-worker | 84.1.2 | multiuser | 4a139f395a | linux | amd64 | 1.24.4 | 3 | 3 |
| **translations-1/b-linux-v100-gpu-d2g-4-alpha** | generic-worker | 84.1.2 | multiuser | 4a139f395a | linux | amd64 | 1.24.4 | 2 | 2 |
| **translations-1/b-linux-v100-gpu-d2g-alpha** | generic-worker | 84.1.2 | multiuser | 4a139f395a | linux | amd64 | 1.24.4 | 2 | 2 |
| **translations-1/decision** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 1248 | 1248 |
| **translations-1/images** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 4 | 4 |
| **translations-t/t-linux-gw-noscratch-amd** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 2 | 2 |
| **xpi-1/b-linux** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 122 | 122 |
| **xpi-1/decision** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 62 | 62 |
| **xpi-1/images** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 2 | 2 |
| **xpi-3/b-linux** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 4 | 4 |
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
| **app-services-1/b-linux-amd** | docker-worker | 38.0.5 | 2 | 2 |
| **app-services-1/b-linux-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **app-services-1/decision-gcp** | docker-worker | 38.0.5 | 808 | 1616 |
| **app-services-1/images-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **app-services-3/b-linux-amd** | docker-worker | 38.0.5 | 2 | 2 |
| **app-services-3/b-linux-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **app-services-3/decision-gcp** | docker-worker | 38.0.5 | 2 | 4 |
| **app-services-3/images-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **code-analysis-1/linux-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **code-analysis-3/linux-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **comm-1/b-linux-amd** | docker-worker | 38.0.5 | 2 | 2 |
| **comm-1/b-linux-gcp** | docker-worker | 38.0.5 | 6 | 6 |
| **comm-1/b-linux-large-amd** | docker-worker | 38.0.5 | 2 | 2 |
| **comm-1/b-linux-large-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **comm-1/b-linux-xlarge-amd** | docker-worker | 38.0.5 | 2 | 2 |
| **comm-1/b-linux-xlarge-gcp** | docker-worker | 38.0.5 | 3 | 3 |
| **comm-1/decision-gcp** | docker-worker | 38.0.5 | 10 | 10 |
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
| **comm-3/b-linux-gcp** | docker-worker | 38.0.5 | 1862 | 1862 |
| **comm-3/b-linux-large-amd** | docker-worker | 38.0.5 | 2 | 2 |
| **comm-3/b-linux-large-gcp** | docker-worker | 38.0.5 | 3 | 3 |
| **comm-3/b-linux-xlarge-amd** | docker-worker | 38.0.5 | 2 | 2 |
| **comm-3/b-linux-xlarge-gcp** | docker-worker | 38.0.5 | 8 | 8 |
| **comm-3/decision-gcp** | docker-worker | 38.0.5 | 43 | 43 |
| **comm-3/images-gcp** | docker-worker | 38.0.5 | 26 | 26 |
| **comm-t/misc-gcp** | docker-worker | 38.0.5 | 2 | 16 |
| **comm-t/t-linux-large-gcp** | docker-worker | unknown version | 2 | 2 |
| **comm-t/t-linux-large-noscratch-gcp** | docker-worker | unknown version | 2 | 2 |
| **comm-t/t-linux-xlarge-gcp** | docker-worker | unknown version | 2 | 2 |
| **comm-t/t-linux-xlarge-noscratch-gcp** | docker-worker | unknown version | 2 | 2 |
| **comm-t/t-linux-xlarge-source-gcp** | docker-worker | unknown version | 2 | 2 |
| **comm-t/t-linux-xlarge-source-noscratch-gcp** | docker-worker | unknown version | 2 | 2 |
| **gecko-1/b-linux-amd** | docker-worker | 38.0.5 | 12 | 12 |
| **gecko-1/b-linux-gcp** | docker-worker | 38.0.5 | 11 | 11 |
| **gecko-1/b-linux-gcp-relops1411** | docker-worker | 38.0.5 | 2 | 2 |
| **gecko-1/b-linux-gcp-test-bug-1882320** | docker-worker | 38.0.5 | 2 | 2 |
| **gecko-1/b-linux-kvm-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **gecko-1/b-linux-large-amd** | docker-worker | 38.0.5 | 2 | 2 |
| **gecko-1/b-linux-large-gcp** | docker-worker | 38.0.5 | 4 | 4 |
| **gecko-1/b-linux-medium-gcp** | docker-worker | 38.0.5 | 4 | 4 |
| **gecko-1/b-linux-xlarge-amd** | docker-worker | 38.0.5 | 10 | 10 |
| **gecko-1/b-linux-xlarge-gcp** | docker-worker | 38.0.5 | 3 | 3 |
| **gecko-1/b-linux-xlarge-gcp-bug1797804-c2** | docker-worker | 38.0.5 | 2 | 2 |
| **gecko-1/decision-gcp** | docker-worker | 38.0.5 | 13 | 13 |
| **gecko-1/decision-gcp-c3d4-amd-bug1990935** | docker-worker | 38.0.5 | 2 | 2 |
| **gecko-1/decision-gcp-c3d8-amd-bug1990935** | docker-worker | 38.0.5 | 2 | 2 |
| **gecko-1/decision-gcp-c4d4-amd-bug1990935** | docker-worker | 38.0.5 | 2 | 2 |
| **gecko-1/decision-gcp-c4d4-hcpu-amd-bug1990935** | docker-worker | 38.0.5 | 2 | 2 |
| **gecko-1/decision-gcp-c4d8-amd-bug1990935** | docker-worker | 38.0.5 | 2 | 2 |
| **gecko-1/decision-gcp-c4d8-hcpu-amd-bug1990935** | docker-worker | 38.0.5 | 2 | 2 |
| **gecko-1/decision-gcp-c4d8-lssd-amd-bug1990935** | docker-worker | 38.0.5 | 2 | 2 |
| **gecko-1/images-gcp** | docker-worker | 38.0.5 | 4 | 4 |
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
| **gecko-3/b-linux-amd** | docker-worker | 38.0.5 | 47 | 47 |
| **gecko-3/b-linux-gcp** | docker-worker | 38.0.5 | 2888 | 2888 |
| **gecko-3/b-linux-kvm-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **gecko-3/b-linux-large-amd** | docker-worker | 38.0.5 | 2 | 2 |
| **gecko-3/b-linux-large-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **gecko-3/b-linux-medium-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **gecko-3/b-linux-xlarge-amd** | docker-worker | 38.0.5 | 2 | 2 |
| **gecko-3/b-linux-xlarge-gcp** | docker-worker | 38.0.5 | 8 | 8 |
| **gecko-3/decision-gcp** | docker-worker | 38.0.5 | 818 | 818 |
| **gecko-3/images-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **gecko-t/misc-gcp** | docker-worker | 38.0.5 | 29 | 232 |
| **gecko-t/t-linux-kvm-gcp** | docker-worker | unknown version | 2 | 2 |
| **gecko-t/t-linux-kvm-gcp-relops1411** | docker-worker | unknown version | 2 | 2 |
| **gecko-t/t-linux-kvm-noscratch-gcp** | docker-worker | unknown version | 17 | 17 |
| **gecko-t/t-linux-large-gcp** | docker-worker | unknown version | 2 | 2 |
| **gecko-t/t-linux-large-hostub2204-gcp** | docker-worker | 48.3.0 | 2 | 2 |
| **gecko-t/t-linux-large-noscratch-gcp** | docker-worker | unknown version | 3845 | 3845 |
| **gecko-t/t-linux-xlarge-gcp** | docker-worker | unknown version | 51 | 51 |
| **gecko-t/t-linux-xlarge-noscratch-gcp** | docker-worker | unknown version | 2449 | 2449 |
| **gecko-t/t-linux-xlarge-source-gcp** | docker-worker | unknown version | 192 | 192 |
| **gecko-t/t-linux-xlarge-source-noscratch-gcp** | docker-worker | unknown version | 2 | 2 |
| **glean-1/b-linux-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **glean-1/decision-gcp** | docker-worker | 38.0.5 | 799 | 1598 |
| **glean-1/images-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **glean-3/b-linux-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **glean-3/decision-gcp** | docker-worker | 38.0.5 | 2 | 4 |
| **glean-3/images-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **infra/build-decision** | docker-worker | 38.0.5 | 1 | 8 |
| **mobile-1/b-linux-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **mobile-1/decision-gcp** | docker-worker | 38.0.5 | 798 | 1596 |
| **mobile-1/images-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **mobile-3/b-linux-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **mobile-3/decision-gcp** | docker-worker | 38.0.5 | 2 | 4 |
| **mobile-3/images-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **mozilla-1/b-linux-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **mozilla-1/decision-gcp** | docker-worker | 38.0.5 | 794 | 1588 |
| **mozilla-1/images-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **mozilla-3/b-linux-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **mozilla-3/decision-gcp** | docker-worker | 38.0.5 | 2 | 4 |
| **mozilla-3/images-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **mozilla-t/t-linux-large-gcp** | docker-worker | unknown version | 2 | 2 |
| **mozilla-t/t-linux-large-noscratch-gcp** | docker-worker | unknown version | 2 | 2 |
| **mozillavpn-1/b-linux-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **mozillavpn-1/b-linux-large-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **mozillavpn-1/decision-gcp** | docker-worker | 38.0.5 | 798 | 1596 |
| **mozillavpn-1/images-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **mozillavpn-3/b-linux-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **mozillavpn-3/b-linux-large-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **mozillavpn-3/decision-gcp** | docker-worker | 38.0.5 | 2 | 4 |
| **mozillavpn-3/images-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **nss-1/decision-gcp** | docker-worker | 38.0.5 | 3 | 6 |
| **nss-1/images-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **nss-1/linux-gcp** | docker-worker | 38.0.5 | 83 | 83 |
| **nss-3/decision-gcp** | docker-worker | 38.0.5 | 3 | 6 |
| **nss-3/images-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **nss-3/linux-gcp** | docker-worker | 38.0.5 | 83 | 83 |
| **nss-t/t-linux-xlarge-gcp** | docker-worker | unknown version | 97 | 97 |
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
| **translations-1/decision-gcp** | docker-worker | 38.0.5 | 1785 | 3570 |
| **translations-1/images-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **xpi-1/b-linux-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **xpi-1/decision-gcp** | docker-worker | 38.0.5 | 2 | 4 |
| **xpi-1/images-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **xpi-3/b-linux-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **xpi-3/decision-gcp** | docker-worker | 38.0.5 | 2 | 4 |
| **xpi-3/images-gcp** | docker-worker | 38.0.5 | 2 | 2 |


## Script Worker

Total: `44`



| Worker Pool | Implementation | Version | Total Workers | Total Capacity |
| --- | --- | --- | ---: | ---: |
| **scriptworker-k8s/app-services-3-signing** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/comm-1-balrog** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/comm-1-beetmover** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/comm-1-bouncer** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/comm-1-tree** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/comm-3-balrog** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/comm-3-beetmover** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/comm-3-bouncer** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/comm-3-pushmsix** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/comm-3-shipit** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/comm-3-signing** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/comm-3-tree** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/comm-t-signing** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/enterprise-3-signing** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
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
| **scriptworker-k8s/mobile-t-signing** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/mozillavpn-3-signing** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/translations-1-beetmover** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/xpi-3-github** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/xpi-3-shipit** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/xpi-3-signing** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
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
| **gecko-t/t-linux-vm-2204-wayland** |  | No artifacts found | 44 | 44 |
| **gecko-t/t-linux-vm-2204-wayland-snap** |  | No artifacts found | 2 | 2 |


## Version not determined [^2]

Total: `37`


Count by image:

| Version | Count |
| :--- | ---: |
| projects/taskcluster-imaging/global/images/gw-fxci-gcp-l1-2404-amd64-googlecompute-alpha | 1 |
| unknown | 16 |
|  | 2 |
| /subscriptions/a30e97ab-734a-4f3b-a0e4-c51c0bff0701/resourceGroups/rg-packer-worker-images/providers/Microsoft.Compute/galleries/trusted_win11_a64_25h2_builder/images/trusted_win11_a64_25h2_builder/versions/1.0.0 | 4 |
| /subscriptions/108d46d5-fe9b-4850-9a7d-8c914aa6c1f0/resourceGroups/rg-packer-worker-images/providers/Microsoft.Compute/galleries/win11a6425h2builder/images/win11a6425h2builder/versions/1.0.0 | 4 |
| /subscriptions/108d46d5-fe9b-4850-9a7d-8c914aa6c1f0/resourceGroups/rg-packer-worker-images/providers/Microsoft.Compute/galleries/win11a6425h2tester/images/win11a6425h2tester/versions/1.0.0 | 4 |
| /subscriptions/108d46d5-fe9b-4850-9a7d-8c914aa6c1f0/resourceGroups/rg-packer-worker-images/providers/Microsoft.Compute/galleries/win11a6425h2testeralpha/images/win11a6425h2testeralpha/versions/1.0.0 | 4 |
| /subscriptions/108d46d5-fe9b-4850-9a7d-8c914aa6c1f0/resourceGroups/rg-packer-worker-images/providers/Microsoft.Compute/galleries/win11a6425h2builderalpha/images/win11a6425h2builderalpha/versions/1.0.0 | 2 |


| Worker Pool | Implementation | Version | Total Workers | Total Capacity |
| --- | --- | --- | ---: | ---: |
| **comm-t/win11-64-25h2** |  | Version not determined; task not (yet) claimed | 80 | 80 |
| **comm-t/win11-64-25h2-ssd** |  | Version not determined; task not (yet) claimed | 80 | 80 |
| **comm-t/win11-a64-25h2** |  | Version not determined; task not (yet) claimed | 28 | 28 |
| **comm-t/win11-a64-25h2-alpha** |  | Version not determined; task not (yet) claimed | 28 | 28 |
| **enterprise-1/win11-a64-25h2-builder** |  | Version not determined; task not (yet) claimed | 28 | 28 |
| **enterprise-3/win11-a64-25h2-builder** |  | Version not determined; task not (yet) claimed | 28 | 28 |
| **enterprise-t/win11-64-25h2** |  | Version not determined; task not (yet) claimed | 80 | 80 |
| **enterprise-t/win11-64-25h2-source** |  | Version not determined; task not (yet) claimed | 80 | 80 |
| **gecko-1/win11-a64-25h2-builder** |  | Version not determined; task not (yet) claimed | 27 | 27 |
| **gecko-1/win11-a64-25h2-builder-alpha** |  | Version not determined; task not (yet) claimed | 27 | 27 |
| **gecko-3/win11-a64-25h2-builder** |  | Version not determined; task not (yet) claimed | 27 | 27 |
| **gecko-t/t-linux-2404-relsre** |  | Version not determined; task not (yet) claimed | 4 | 4 |
| **gecko-t/win10-64-2009-alpha** |  | Version not determined; task not (yet) claimed | 301 | 301 |
| **gecko-t/win10-64-2009-gpu-alpha** |  | Version not determined; task not (yet) claimed | 86 | 86 |
| **gecko-t/win10-64-2009-source-alpha** |  | Version not determined; task not (yet) claimed | 125 | 125 |
| **gecko-t/win10-64-2009-webgpu-alpha** |  | Version not determined; task not (yet) claimed | 157 | 157 |
| **gecko-t/win11-64-25h2** |  | Version not determined; task not (yet) claimed | 79 | 79 |
| **gecko-t/win11-64-25h2-gpu** |  | Version not determined; task not (yet) claimed | 79 | 79 |
| **gecko-t/win11-64-25h2-large** |  | Version not determined; task not (yet) claimed | 79 | 79 |
| **gecko-t/win11-64-25h2-source** |  | Version not determined; task not (yet) claimed | 79 | 79 |
| **gecko-t/win11-64-25h2-ssd** |  | Version not determined; task not (yet) claimed | 79 | 79 |
| **gecko-t/win11-64-25h2-ssd-gpu** |  | Version not determined; task not (yet) claimed | 79 | 79 |
| **gecko-t/win11-64-25h2-webgpu** |  | Version not determined; task not (yet) claimed | 79 | 79 |
| **gecko-t/win11-a64-25h2** |  | Version not determined; task not (yet) claimed | 24 | 24 |
| **gecko-t/win11-a64-25h2-alpha** |  | Version not determined; task not (yet) claimed | 23 | 23 |
| **mozillavpn-1/win11-a64-25h2-builder** |  | Version not determined; task not (yet) claimed | 20 | 20 |
| **mozillavpn-3/win11-a64-25h2-builder** |  | Version not determined; task not (yet) claimed | 19 | 19 |
| **nss-1/win11-a64-25h2** |  | Version not determined; task not (yet) claimed | 26 | 26 |
| **nss-1/win11-a64-25h2-alpha** |  | Version not determined; task not (yet) claimed | 24 | 24 |
| **nss-1/win11-a64-25h2-builder** |  | Version not determined; task not (yet) claimed | 18 | 18 |
| **nss-1/win11-a64-25h2-builder-alpha** |  | Version not determined; task not (yet) claimed | 18 | 18 |
| **nss-3/win11-a64-25h2-builder** |  | Version not determined; task not (yet) claimed | 19 | 19 |
| **nss-t/win11-a64-25h2** |  | Version not determined; task not (yet) claimed | 30 | 30 |
| **nss-t/win11-a64-25h2-alpha** |  | Version not determined; task not (yet) claimed | 23 | 23 |
| **releng-hardware/gecko-t-linux-netperf-1804** |  | Version not determined; task not (yet) claimed | 0 | 0 |
| **releng-hardware/nss-1-b-osx-1015** |  | Version not determined; task not (yet) claimed | 0 | 0 |
| **taskgraph-t/win11-64-25h2** |  | Version not determined; task not (yet) claimed | 66 | 66 |



[^1]: Those are the pools whose tasks were claimed and resolved by a worker as expected, but the worker did not publish either artifact `public/logs/live_backing.log` nor `public/logs/chain_of_trust.log`, which is the source used to identify the worker implementation.

[^2]: Probing task remains pending after two hours. Those are the pools that were not able to start any worker to claim the task within two hours.

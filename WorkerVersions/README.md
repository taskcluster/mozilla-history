

# Worker Pool Versions


## Generic Worker

Total: `378`

Count by version:

| Version | Count |
| :--- | ---: |
| 36.0.0 | 4 |
| 45.0.0 | 1 |
| 60.3.4 | 15 |
| 61.0.0 | 1 |
| 64.3.0 | 26 |
| 72.0.1 | 1 |
| 84.1.2 | 7 |
| 87.0.0 | 1 |
| 88.0.2 | 2 |
| 91.0.2 | 2 |
| 91.1.0 | 7 |
| 95.1.3 | 202 |
| 96.2.3 | 100 |
| 96.5.2 | 9 |


Count by image:

| Version | Count |
| :--- | ---: |
| projects/taskcluster-imaging/global/images/gw-fxci-gcp-l1-2404-arm64-headless-googlecompute-2026-01-14 | 14 |
| projects/fxci-production-level3-workers/global/images/gw-fxci-gcp-l3-2404-arm64-headless-googlecompute-2026-01-14 | 6 |
| projects/taskcluster-imaging/global/images/gw-fxci-gcp-l1-arm64-gui-googlecompute-2024-09-18t14-57-52z | 9 |
| /subscriptions/108d46d5-fe9b-4850-9a7d-8c914aa6c1f0/resourceGroups/rg-packer-worker-images/providers/Microsoft.Compute/galleries/win11a6425h2testeralpha/images/win11a6425h2testeralpha/versions/1.0.0 | 4 |
| /subscriptions/108d46d5-fe9b-4850-9a7d-8c914aa6c1f0/resourceGroups/rg-packer-worker-images/providers/Microsoft.Compute/galleries/win11_a64_24h2_builder_alpha/images/win11_a64_24h2_builder_alpha/versions/1.0.0 | 2 |
| projects/fxci-production-level3-workers/global/images/gw-fxci-gcp-l3-gui-googlecompute-2024-09-18t05-46-31z | 3 |
| /subscriptions/108d46d5-fe9b-4850-9a7d-8c914aa6c1f0/resourceGroups/rg-packer-worker-images/providers/Microsoft.Compute/galleries/win11_a64_24h2_tester_alpha/images/win11_a64_24h2_tester_alpha/versions/1.0.0 | 4 |
| projects/taskcluster-imaging/global/images/gw-fxci-gcp-l1-gui-googlecompute-2025-01-13t22-33-40z | 2 |
| projects/fxci-production-level3-workers/global/images/gw-fxci-gcp-l3-arm64-gui-googlecompute-2024-09-18t19-02-58z | 4 |
| projects/taskcluster-imaging/global/images/gw-fxci-gcp-l1-2404-amd64-gui-googlecompute-alpha | 1 |
| projects/taskcluster-imaging/global/images/gw-fxci-gcp-l1-2404-arm64-headless-googlecompute-alpha | 1 |
| unknown | 78 |
| /subscriptions/a30e97ab-734a-4f3b-a0e4-c51c0bff0701/resourceGroups/rg-packer-worker-images/providers/Microsoft.Compute/galleries/trusted_win11_a64_24h2_builder/images/trusted_win11_a64_24h2_builder/versions/1.2.0 | 4 |
| projects/taskcluster-imaging/global/images/gw-fxci-gcp-l1-2404-amd64-headless-googlecompute-2026-01-14 | 120 |
| /subscriptions/108d46d5-fe9b-4850-9a7d-8c914aa6c1f0/resourceGroups/rg-packer-worker-images/providers/Microsoft.Compute/galleries/win11_a64_24h2_builder/images/win11_a64_24h2_builder/versions/1.2.0 | 4 |
| /subscriptions/108d46d5-fe9b-4850-9a7d-8c914aa6c1f0/resourceGroups/rg-packer-worker-images/providers/Microsoft.Compute/galleries/win11_a64_24h2_tester/images/win11_a64_24h2_tester/versions/1.2.0 | 4 |
| projects/taskcluster-imaging/global/images/gw-fxci-gcp-l1-2404-amd64-gui-googlecompute-2026-01-14 | 4 |
|  | 34 |
| projects/fxci-production-level3-workers/global/images/gw-fxci-gcp-l3-2404-amd64-headless-googlecompute-2026-01-14 | 58 |
| projects/taskcluster-imaging/global/images/gw-fxci-gcp-l1-gui-googlecompute-2024-08-22t22-48-09z | 8 |
| projects/fxci-production-level3-workers/global/images/gw-fxci-gcp-l3-2404-amd64-headless-googlecompute-alpha | 1 |
| projects/taskcluster-imaging/global/images/gw-fxci-gcp-l1-2404-amd64-headless-googlecompute-alpha | 6 |
| projects/taskcluster-imaging/global/images/gw-fxci-gcp-l1-2404-amd64-headless-googlecompute-alpha-tc | 7 |


| Worker Pool | Implementation | Version | Engine | Revision | OS | Arch | GO | Total Workers | Total Capacity |
| --- | --- | --- | --- | --- | --- | --- | --- | ---: | ---: |
| **adhoc-1/decision** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 5 | 5 |
| **adhoc-1/images** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 16 | 16 |
| **adhoc-3/decision** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 4 | 4 |
| **adhoc-3/images** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 4 | 4 |
| **app-services-1/b-linux** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 301 | 301 |
| **app-services-1/b-linux-docker-amd** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 2 | 2 |
| **app-services-1/decision** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 1229 | 1229 |
| **app-services-1/images** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 3 | 3 |
| **app-services-3/b-linux** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 1225 | 1225 |
| **app-services-3/b-linux-docker-amd** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 2 | 2 |
| **app-services-3/decision** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 38 | 38 |
| **app-services-3/images** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 3 | 3 |
| **code-analysis-1/linux-docker** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 2 | 2 |
| **code-analysis-1/linux-gw-gcp** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | amd64 | 1.22.2 | 54 | 54 |
| **code-analysis-3/linux-docker** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 2 | 2 |
| **code-analysis-3/linux-gw-gcp** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | amd64 | 1.22.2 | 54 | 54 |
| **code-coverage/bot** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 280 | 280 |
| **code-review/bot** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 6502 | 6502 |
| **comm-1/b-linux** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 801 | 801 |
| **comm-1/b-linux-aarch64** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | arm64 | 1.25.5 | 2 | 2 |
| **comm-1/b-linux-docker-amd** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 2 | 2 |
| **comm-1/b-linux-docker-large-amd** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 5 | 5 |
| **comm-1/b-linux-docker-xlarge-amd** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 2 | 2 |
| **comm-1/b-linux-gcp-aarch64** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | arm64 | 1.22.2 | 2 | 2 |
| **comm-1/b-linux-large** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 2 | 2 |
| **comm-1/b-linux-xlarge** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 2 | 2 |
| **comm-1/b-win2022** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 22 | 22 |
| **comm-1/decision** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 312 | 312 |
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
| **comm-3/b-linux** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 2434 | 2434 |
| **comm-3/b-linux-aarch64** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | arm64 | 1.25.5 | 2 | 2 |
| **comm-3/b-linux-docker-amd** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 1370 | 1370 |
| **comm-3/b-linux-docker-large-amd** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 4 | 4 |
| **comm-3/b-linux-docker-xlarge-amd** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 5 | 5 |
| **comm-3/b-linux-gcp-aarch64** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | arm64 | 1.22.2 | 2 | 2 |
| **comm-3/b-linux-large** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 2 | 2 |
| **comm-3/b-linux-xlarge** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 2 | 2 |
| **comm-3/b-win2022** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 556 | 556 |
| **comm-3/decision** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 83 | 83 |
| **comm-3/images** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 2 | 2 |
| **comm-3/images-aarch64** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | arm64 | 1.25.5 | 2 | 2 |
| **comm-3/images-gcp-aarch64** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | arm64 | 1.22.2 | 2 | 2 |
| **comm-t/misc** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 2 | 2 |
| **comm-t/t-linux-arm64-docker** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | arm64 | 1.25.5 | 2 | 2 |
| **comm-t/t-linux-docker** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 14 | 14 |
| **comm-t/t-linux-docker-amd** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 2972 | 2972 |
| **comm-t/t-linux-docker-noscratch** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 33 | 33 |
| **comm-t/t-linux-docker-noscratch-amd** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 3544 | 3544 |
| **comm-t/win11-64-2009** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 2 | 2 |
| **comm-t/win11-64-2009-ssd** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 2 | 2 |
| **comm-t/win11-64-24h2** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 5839 | 5839 |
| **comm-t/win11-64-24h2-ssd** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 327 | 327 |
| **comm-t/win11-64-25h2** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 167 | 167 |
| **comm-t/win11-64-25h2-ssd** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 167 | 167 |
| **comm-t/win11-a64-24h2** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | arm64 | 1.25.7 | 2 | 2 |
| **comm-t/win11-a64-24h2-alpha** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | arm64 | 1.25.7 | 8 | 8 |
| **comm-t/win11-a64-25h2-alpha** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | arm64 | 1.25.7 | 57 | 57 |
| **enterprise-1/b-linux** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 643 | 643 |
| **enterprise-1/b-linux-aarch64** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | arm64 | 1.25.5 | 2 | 2 |
| **enterprise-1/b-linux-docker-amd** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 316 | 316 |
| **enterprise-1/b-linux-docker-large-amd** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 19 | 19 |
| **enterprise-1/b-linux-docker-xlarge-amd** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 424 | 424 |
| **enterprise-1/b-win2022** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 147 | 147 |
| **enterprise-1/decision** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 1220 | 1220 |
| **enterprise-1/images** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 3 | 3 |
| **enterprise-1/images-aarch64** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | arm64 | 1.25.5 | 2 | 2 |
| **enterprise-1/win11-a64-24h2-builder** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | arm64 | 1.25.7 | 2 | 2 |
| **enterprise-3/b-linux** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 793 | 793 |
| **enterprise-3/b-linux-aarch64** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | arm64 | 1.25.5 | 2 | 2 |
| **enterprise-3/b-linux-docker-amd** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 267 | 267 |
| **enterprise-3/b-linux-docker-large-amd** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 61 | 61 |
| **enterprise-3/b-linux-docker-xlarge-amd** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 289 | 289 |
| **enterprise-3/b-win2022** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 239 | 239 |
| **enterprise-3/decision** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 60 | 60 |
| **enterprise-3/images** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 3 | 3 |
| **enterprise-3/images-aarch64** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | arm64 | 1.25.5 | 2 | 2 |
| **enterprise-3/win11-a64-24h2-builder** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | arm64 | 1.25.7 | 5 | 5 |
| **enterprise-t/misc** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 180 | 180 |
| **enterprise-t/t-linux-docker-16c32gb-amd** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 505 | 505 |
| **enterprise-t/t-linux-docker-amd** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 1939 | 1939 |
| **enterprise-t/t-linux-docker-noscratch-amd** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 2657 | 2657 |
| **enterprise-t/win11-64-24h2** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 1546 | 1546 |
| **enterprise-t/win11-64-24h2-alpha** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 2 | 2 |
| **enterprise-t/win11-64-24h2-source** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 2 | 2 |
| **enterprise-t/win11-64-24h2-source-alpha** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 2 | 2 |
| **enterprise-t/win11-64-25h2** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 167 | 167 |
| **enterprise-t/win11-64-25h2-alpha** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 2 | 2 |
| **enterprise-t/win11-64-25h2-source** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 167 | 167 |
| **enterprise-t/win11-64-25h2-source-alpha** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 2 | 2 |
| **gecko-1/b-linux** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 6127 | 6127 |
| **gecko-1/b-linux-2204-kvm-gcp** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | amd64 | 1.22.2 | 2 | 2 |
| **gecko-1/b-linux-aarch64** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | arm64 | 1.25.5 | 11 | 11 |
| **gecko-1/b-linux-docker-alpha** | generic-worker | 96.5.2 | multiuser | 9892d6e257 | linux | amd64 | 1.26.0 | 2 | 2 |
| **gecko-1/b-linux-docker-amd** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 10624 | 10624 |
| **gecko-1/b-linux-docker-large-amd** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 201 | 201 |
| **gecko-1/b-linux-docker-xlarge-amd** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 1292 | 1292 |
| **gecko-1/b-linux-gcp-aarch64** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | arm64 | 1.22.2 | 2 | 2 |
| **gecko-1/b-linux-kvm** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 314 | 314 |
| **gecko-1/b-linux-large** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 2 | 2 |
| **gecko-1/b-linux-medium** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 18579 | 18579 |
| **gecko-1/b-linux-xlarge** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 11 | 11 |
| **gecko-1/b-win2022** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 1068 | 1068 |
| **gecko-1/b-win2022-alpha** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 2 | 2 |
| **gecko-1/b-win2022-xxlarge** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 18 | 18 |
| **gecko-1/decision** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 4371 | 4371 |
| **gecko-1/images** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 102 | 102 |
| **gecko-1/images-aarch64** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | arm64 | 1.25.5 | 8 | 8 |
| **gecko-1/images-gcp-aarch64** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | arm64 | 1.22.2 | 2 | 2 |
| **gecko-1/win11-a64-24h2-builder** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | arm64 | 1.25.7 | 41 | 41 |
| **gecko-1/win11-a64-24h2-builder-alpha** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | arm64 | 1.25.7 | 3 | 3 |
| **gecko-2/b-linux** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 199 | 199 |
| **gecko-2/b-linux-aarch64** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | arm64 | 1.25.5 | 2 | 2 |
| **gecko-2/b-linux-docker-amd** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 268 | 268 |
| **gecko-2/b-linux-docker-large-amd** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 46 | 46 |
| **gecko-2/b-linux-docker-xlarge-amd** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 53 | 53 |
| **gecko-2/b-linux-gcp-aarch64** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | arm64 | 1.22.2 | 2 | 2 |
| **gecko-2/b-linux-kvm** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 19 | 19 |
| **gecko-2/b-linux-large** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 2 | 2 |
| **gecko-2/b-linux-medium** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 379 | 379 |
| **gecko-2/b-linux-xlarge** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 2 | 2 |
| **gecko-2/b-win2022** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 21 | 21 |
| **gecko-2/decision** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 20 | 20 |
| **gecko-2/images** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 2 | 2 |
| **gecko-2/images-aarch64** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | arm64 | 1.25.5 | 2 | 2 |
| **gecko-2/images-gcp-aarch64** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | arm64 | 1.22.2 | 2 | 2 |
| **gecko-3/b-linux** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 17077 | 17077 |
| **gecko-3/b-linux-2204-kvm-gcp** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | amd64 | 1.22.2 | 2 | 2 |
| **gecko-3/b-linux-aarch64** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | arm64 | 1.25.5 | 2 | 2 |
| **gecko-3/b-linux-docker-alpha** | generic-worker | 96.5.2 | multiuser | 9892d6e257 | linux | amd64 | 1.26.0 | 2 | 2 |
| **gecko-3/b-linux-docker-amd** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 22940 | 22940 |
| **gecko-3/b-linux-docker-large-amd** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 475 | 475 |
| **gecko-3/b-linux-docker-xlarge-amd** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 3085 | 3085 |
| **gecko-3/b-linux-gcp-aarch64** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | arm64 | 1.22.2 | 2 | 2 |
| **gecko-3/b-linux-kvm** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 561 | 561 |
| **gecko-3/b-linux-large** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 2 | 2 |
| **gecko-3/b-linux-medium** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 22079 | 22079 |
| **gecko-3/b-linux-xlarge** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 31 | 31 |
| **gecko-3/b-win2022** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 4983 | 4983 |
| **gecko-3/b-win2022-xxlarge** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 13 | 13 |
| **gecko-3/decision** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 5173 | 5173 |
| **gecko-3/images** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 11 | 11 |
| **gecko-3/images-aarch64** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | arm64 | 1.25.5 | 2 | 2 |
| **gecko-3/images-gcp-aarch64** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | arm64 | 1.22.2 | 2 | 2 |
| **gecko-3/win11-a64-24h2-builder** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | arm64 | 1.25.7 | 43 | 43 |
| **gecko-t/misc** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 1053 | 1053 |
| **gecko-t/t-linux-2204-wayland** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | amd64 | 1.22.2 | 8719 | 8719 |
| **gecko-t/t-linux-2204-wayland-arm64-relsre** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | arm64 | 1.22.2 | 2 | 2 |
| **gecko-t/t-linux-2204-wayland-relsre** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | amd64 | 1.22.2 | 2 | 2 |
| **gecko-t/t-linux-2204-wayland-root-exp** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | amd64 | 1.22.2 | 2 | 2 |
| **gecko-t/t-linux-2204-wayland-snap** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | amd64 | 1.22.2 | 135 | 135 |
| **gecko-t/t-linux-2404-headless-alpha** | generic-worker | 96.5.2 | multiuser | 9892d6e257 | linux | amd64 | 1.26.0 | 3 | 3 |
| **gecko-t/t-linux-2404-headless-arm64-alpha** | generic-worker | 96.5.2 | multiuser | 9892d6e257 | linux | arm64 | 1.26.0 | 2 | 2 |
| **gecko-t/t-linux-2404-headless-ssd-alpha** | generic-worker | 96.5.2 | multiuser | 9892d6e257 | linux | amd64 | 1.26.0 | 2 | 2 |
| **gecko-t/t-linux-2404-wayland** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 2 | 2 |
| **gecko-t/t-linux-2404-wayland-relsre** | generic-worker | 96.5.2 | multiuser | 9892d6e257 | linux | amd64 | 1.26.0 | 2 | 2 |
| **gecko-t/t-linux-2404-wayland-snap** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 130 | 130 |
| **gecko-t/t-linux-arm64-docker** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | arm64 | 1.25.5 | 26 | 26 |
| **gecko-t/t-linux-docker** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 1701 | 1701 |
| **gecko-t/t-linux-docker-16c32gb-amd** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 3707 | 3707 |
| **gecko-t/t-linux-docker-amd** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 30139 | 30139 |
| **gecko-t/t-linux-docker-amd-alpha** | generic-worker | 96.5.2 | multiuser | 9892d6e257 | linux | amd64 | 1.26.0 | 2 | 2 |
| **gecko-t/t-linux-docker-kvm** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 69581 | 69581 |
| **gecko-t/t-linux-docker-kvm-alpha** | generic-worker | 96.5.2 | multiuser | 9892d6e257 | linux | amd64 | 1.26.0 | 2 | 2 |
| **gecko-t/t-linux-docker-noscratch** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 2034 | 2034 |
| **gecko-t/t-linux-docker-noscratch-amd** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 119749 | 119749 |
| **gecko-t/t-linux-docker-noscratch-amd-alpha** | generic-worker | 96.5.2 | multiuser | 9892d6e257 | linux | amd64 | 1.26.0 | 2 | 2 |
| **gecko-t/t-linux-xlarge-2204-wayland** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | amd64 | 1.22.2 | 12701 | 12701 |
| **gecko-t/t-linux-xlarge-2404-wayland** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 2 | 2 |
| **gecko-t/win10-64-2009** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 4327 | 4327 |
| **gecko-t/win10-64-2009-alpha** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 172 | 172 |
| **gecko-t/win10-64-2009-gpu** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 423 | 423 |
| **gecko-t/win10-64-2009-gpu-alpha** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 166 | 166 |
| **gecko-t/win10-64-2009-source** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 41 | 41 |
| **gecko-t/win10-64-2009-source-alpha** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 167 | 167 |
| **gecko-t/win10-64-2009-webgpu** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 2 | 2 |
| **gecko-t/win10-64-2009-webgpu-alpha** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 172 | 172 |
| **gecko-t/win11-64-2009** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 471 | 471 |
| **gecko-t/win11-64-2009-alpha** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 31 | 31 |
| **gecko-t/win11-64-2009-gpu** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 282 | 282 |
| **gecko-t/win11-64-2009-gpu-alpha** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 7 | 7 |
| **gecko-t/win11-64-2009-source** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 2 | 2 |
| **gecko-t/win11-64-2009-source-alpha** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 4 | 4 |
| **gecko-t/win11-64-2009-ssd** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 2 | 2 |
| **gecko-t/win11-64-2009-ssd-gpu** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 2 | 2 |
| **gecko-t/win11-64-2009-webgpu** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 2 | 2 |
| **gecko-t/win11-64-2009-webgpu-alpha** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 3 | 3 |
| **gecko-t/win11-64-24h2** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 54917 | 54917 |
| **gecko-t/win11-64-24h2-alpha** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 37 | 37 |
| **gecko-t/win11-64-24h2-amd** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 22005 | 22005 |
| **gecko-t/win11-64-24h2-gpu** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 21314 | 21314 |
| **gecko-t/win11-64-24h2-gpu-alpha** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 7 | 7 |
| **gecko-t/win11-64-24h2-large** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 1797 | 1797 |
| **gecko-t/win11-64-24h2-large-alpha** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 3 | 3 |
| **gecko-t/win11-64-24h2-source** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 2353 | 2353 |
| **gecko-t/win11-64-24h2-source-alpha** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 5 | 5 |
| **gecko-t/win11-64-24h2-ssd** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 2 | 2 |
| **gecko-t/win11-64-24h2-ssd-gpu** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 3 | 3 |
| **gecko-t/win11-64-24h2-webgpu** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 3905 | 3905 |
| **gecko-t/win11-64-24h2-webgpu-alpha** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 3 | 3 |
| **gecko-t/win11-64-25h2** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 181 | 181 |
| **gecko-t/win11-64-25h2-alpha** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 612 | 612 |
| **gecko-t/win11-64-25h2-amd** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 1192 | 1192 |
| **gecko-t/win11-64-25h2-gpu** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 13446 | 13446 |
| **gecko-t/win11-64-25h2-gpu-alpha** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 165 | 165 |
| **gecko-t/win11-64-25h2-large** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 165 | 165 |
| **gecko-t/win11-64-25h2-large-alpha** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 4 | 4 |
| **gecko-t/win11-64-25h2-source** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 165 | 165 |
| **gecko-t/win11-64-25h2-source-alpha** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 59 | 59 |
| **gecko-t/win11-64-25h2-ssd** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 165 | 165 |
| **gecko-t/win11-64-25h2-ssd-gpu** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 165 | 165 |
| **gecko-t/win11-64-25h2-webgpu** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 165 | 165 |
| **gecko-t/win11-64-25h2-webgpu-alpha** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 2 | 2 |
| **gecko-t/win11-a64-24h2** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | arm64 | 1.25.7 | 181 | 181 |
| **gecko-t/win11-a64-24h2-alpha** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | arm64 | 1.25.7 | 19 | 19 |
| **gecko-t/win11-a64-25h2-alpha** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | arm64 | 1.25.7 | 49 | 49 |
| **glean-1/b-linux** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 35 | 35 |
| **glean-1/decision** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 1213 | 1213 |
| **glean-1/images** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 2 | 2 |
| **glean-3/b-linux** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 26 | 26 |
| **glean-3/decision** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 8 | 8 |
| **glean-3/images** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 2 | 2 |
| **mobile-1/b-linux** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 100 | 100 |
| **mobile-1/decision** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 1173 | 1173 |
| **mobile-1/images** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 2 | 2 |
| **mobile-3/b-linux** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 203 | 203 |
| **mobile-3/decision** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 133 | 133 |
| **mobile-3/images** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 2 | 2 |
| **mozilla-1/b-linux** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 2 | 2 |
| **mozilla-1/b-win2022** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 2 | 2 |
| **mozilla-1/decision** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 1200 | 1200 |
| **mozilla-1/images** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 8 | 8 |
| **mozilla-3/b-linux** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 2 | 2 |
| **mozilla-3/b-win2022** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 3 | 3 |
| **mozilla-3/decision** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 431 | 431 |
| **mozilla-3/images** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 2 | 2 |
| **mozilla-t/t-linux-2204-wayland** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | amd64 | 1.22.2 | 596 | 596 |
| **mozilla-t/t-linux-2204-wayland-relsre** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | amd64 | 1.22.2 | 2 | 2 |
| **mozilla-t/t-linux-2404-wayland** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 2 | 2 |
| **mozilla-t/t-linux-arm64-docker** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | arm64 | 1.25.5 | 2 | 2 |
| **mozilla-t/t-linux-docker** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 36 | 36 |
| **mozilla-t/t-linux-docker-amd** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 3 | 3 |
| **mozilla-t/t-linux-docker-noscratch** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 34 | 34 |
| **mozilla-t/t-linux-docker-noscratch-amd** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 2 | 2 |
| **mozillavpn-1/b-linux** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 281 | 281 |
| **mozillavpn-1/b-linux-gcp-gw** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | amd64 | 1.22.2 | 2 | 2 |
| **mozillavpn-1/b-linux-large** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 185 | 185 |
| **mozillavpn-1/b-win2022** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 68 | 68 |
| **mozillavpn-1/b-win2022-alpha** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 3 | 3 |
| **mozillavpn-1/decision** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 1198 | 1198 |
| **mozillavpn-1/images** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 3 | 3 |
| **mozillavpn-1/win11-a64-24h2-builder** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | arm64 | 1.25.7 | 2 | 2 |
| **mozillavpn-3/b-linux** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 142 | 142 |
| **mozillavpn-3/b-linux-gcp-gw** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | amd64 | 1.22.2 | 2 | 2 |
| **mozillavpn-3/b-linux-large** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 142 | 142 |
| **mozillavpn-3/b-win2022** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 55 | 55 |
| **mozillavpn-3/decision** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 25 | 25 |
| **mozillavpn-3/images** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 2 | 2 |
| **mozillavpn-3/win11-a64-24h2-builder** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | arm64 | 1.25.7 | 2 | 2 |
| **nss-1/b-win2022** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 459 | 459 |
| **nss-1/b-win2022-alpha** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 3 | 3 |
| **nss-1/decision** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 49 | 49 |
| **nss-1/images** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 2 | 2 |
| **nss-1/linux-docker** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 1461 | 1461 |
| **nss-1/win11-a64-24h2** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | arm64 | 1.25.7 | 2 | 2 |
| **nss-1/win11-a64-24h2-alpha** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | arm64 | 1.25.7 | 7 | 7 |
| **nss-1/win11-a64-24h2-builder** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | arm64 | 1.25.7 | 2 | 2 |
| **nss-1/win11-a64-24h2-builder-alpha** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | arm64 | 1.25.7 | 2 | 2 |
| **nss-1/win11-a64-25h2-alpha** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | arm64 | 1.25.7 | 51 | 51 |
| **nss-3/b-win2022** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 309 | 309 |
| **nss-3/decision** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 18 | 18 |
| **nss-3/images** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 2 | 2 |
| **nss-3/linux-docker** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 1131 | 1131 |
| **nss-3/win11-a64-24h2-builder** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | arm64 | 1.25.7 | 2 | 2 |
| **nss-t/t-linux-arm64-docker** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | arm64 | 1.25.5 | 2 | 2 |
| **nss-t/t-linux-docker** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 2423 | 2423 |
| **nss-t/t-linux-docker-amd** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 2 | 2 |
| **nss-t/win11-a64-24h2** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | arm64 | 1.25.7 | 2 | 2 |
| **nss-t/win11-a64-24h2-alpha** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | arm64 | 1.25.7 | 9 | 9 |
| **nss-t/win11-a64-25h2-alpha** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | arm64 | 1.25.7 | 47 | 47 |
| **proj-autophone/gecko-t-bitbar-gw-perf-a55** | generic-worker | 36.0.0 | simple | b9cb11293f | linux | amd64 | 1.13.7 | 0 | 0 |
| **proj-autophone/gecko-t-bitbar-gw-perf-p6** | generic-worker | 36.0.0 | simple | b9cb11293f | linux | amd64 | 1.13.7 | 0 | 0 |
| **proj-autophone/gecko-t-bitbar-gw-perf-s24** | generic-worker | 36.0.0 | simple | b9cb11293f | linux | amd64 | 1.13.7 | 0 | 0 |
| **proj-autophone/gecko-t-bitbar-gw-unit-p5** | generic-worker | 36.0.0 | simple | b9cb11293f | linux | amd64 | 1.13.7 | 0 | 0 |
| **proj-autophone/gecko-t-lambda-perf-a55** | generic-worker | 87.0.0 | insecure | 99a1fcafbb | linux | amd64 | 1.24.5 | 0 | 0 |
| **releng-1/decision** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 109 | 109 |
| **releng-1/linux-docker** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 93 | 93 |
| **releng-3/decision** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 51 | 51 |
| **releng-3/linux-docker** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 87 | 87 |
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
| **releng-hardware/win11-64-24h2-hw** | generic-worker | 91.1.0 | multiuser | b8f34332c3 | windows | amd64 | 1.25.3 | 0 | 0 |
| **releng-hardware/win11-64-24h2-hw-alpha** | generic-worker | 91.1.0 | multiuser | b8f34332c3 | windows | amd64 | 1.25.3 | 0 | 0 |
| **releng-hardware/win11-64-24h2-hw-perf-debug** | generic-worker | 91.1.0 | multiuser | b8f34332c3 | windows | amd64 | 1.25.3 | 0 | 0 |
| **releng-hardware/win11-64-24h2-hw-perf-sheriff** | generic-worker | 91.1.0 | multiuser | b8f34332c3 | windows | amd64 | 1.25.3 | 0 | 0 |
| **releng-hardware/win11-64-24h2-hw-ref** | generic-worker | 91.1.0 | multiuser | b8f34332c3 | windows | amd64 | 1.25.3 | 0 | 0 |
| **releng-hardware/win11-64-24h2-hw-ref-alpha** | generic-worker | 91.1.0 | multiuser | b8f34332c3 | windows | amd64 | 1.25.3 | 0 | 0 |
| **releng-hardware/win11-64-24h2-hw-relops1213** | generic-worker | 91.1.0 | multiuser | b8f34332c3 | windows | amd64 | 1.25.3 | 0 | 0 |
| **releng-t/linux-docker** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 616 | 616 |
| **relops-1/decision** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 2 | 2 |
| **relops-1/images** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 2 | 2 |
| **relops-3/decision** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 14 | 14 |
| **relops-3/images** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 2 | 2 |
| **scriptworker-1/decision** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 38 | 38 |
| **scriptworker-1/images** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 195 | 195 |
| **scriptworker-3/decision** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 25 | 25 |
| **scriptworker-3/images** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 171 | 171 |
| **taskgraph-1/decision** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 15 | 15 |
| **taskgraph-1/images** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 8 | 8 |
| **taskgraph-3/decision** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 9 | 9 |
| **taskgraph-3/images** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 9 | 9 |
| **taskgraph-t/linux-docker** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 100 | 100 |
| **taskgraph-t/win11-64-24h2** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 5 | 5 |
| **taskgraph-t/win11-64-25h2** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 144 | 144 |
| **translations-1/b-linux-large-gcp-1tb-32-256-d2g** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 19 | 19 |
| **translations-1/b-linux-large-gcp-1tb-32-256-std-d2g** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 14 | 14 |
| **translations-1/b-linux-large-gcp-1tb-64-512-d2g** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 2 | 2 |
| **translations-1/b-linux-large-gcp-1tb-64-512-std-d2g** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 2 | 2 |
| **translations-1/b-linux-large-gcp-d2g** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 144 | 144 |
| **translations-1/b-linux-large-gcp-d2g-1tb** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 2 | 2 |
| **translations-1/b-linux-large-gcp-d2g-1tb-standard** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 2 | 2 |
| **translations-1/b-linux-large-gcp-d2g-300gb** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 296 | 296 |
| **translations-1/b-linux-v100-gpu-d2g** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 212 | 212 |
| **translations-1/b-linux-v100-gpu-d2g-4** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 132 | 132 |
| **translations-1/b-linux-v100-gpu-d2g-4-1tb** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 67 | 67 |
| **translations-1/b-linux-v100-gpu-d2g-4-1tb-alpha** | generic-worker | 84.1.2 | multiuser | 4a139f395a | linux | amd64 | 1.24.4 | 2 | 2 |
| **translations-1/b-linux-v100-gpu-d2g-4-1tb-standard** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 2 | 2 |
| **translations-1/b-linux-v100-gpu-d2g-4-1tb-std-alpha** | generic-worker | 84.1.2 | multiuser | 4a139f395a | linux | amd64 | 1.24.4 | 2 | 2 |
| **translations-1/b-linux-v100-gpu-d2g-4-2tb** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 14 | 14 |
| **translations-1/b-linux-v100-gpu-d2g-4-2tb-alpha** | generic-worker | 84.1.2 | multiuser | 4a139f395a | linux | amd64 | 1.24.4 | 2 | 2 |
| **translations-1/b-linux-v100-gpu-d2g-4-300gb** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 48 | 48 |
| **translations-1/b-linux-v100-gpu-d2g-4-300gb-alpha** | generic-worker | 84.1.2 | multiuser | 4a139f395a | linux | amd64 | 1.24.4 | 2 | 2 |
| **translations-1/b-linux-v100-gpu-d2g-4-300gb-standard** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 2 | 2 |
| **translations-1/b-linux-v100-gpu-d2g-4-300gb-std-alpha** | generic-worker | 84.1.2 | multiuser | 4a139f395a | linux | amd64 | 1.24.4 | 3 | 3 |
| **translations-1/b-linux-v100-gpu-d2g-4-alpha** | generic-worker | 84.1.2 | multiuser | 4a139f395a | linux | amd64 | 1.24.4 | 2 | 2 |
| **translations-1/b-linux-v100-gpu-d2g-alpha** | generic-worker | 84.1.2 | multiuser | 4a139f395a | linux | amd64 | 1.24.4 | 2 | 2 |
| **translations-1/decision** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 1203 | 1203 |
| **translations-1/images** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 4 | 4 |
| **translations-t/t-linux-gw-noscratch-amd** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 2 | 2 |
| **xpi-1/b-linux** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 86 | 86 |
| **xpi-1/decision** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 51 | 51 |
| **xpi-1/images** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 2 | 2 |
| **xpi-3/b-linux** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 6 | 6 |
| **xpi-3/decision** | generic-worker | 95.1.3 | multiuser | 300ac23a15 | linux | amd64 | 1.25.5 | 7 | 7 |
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
| projects/fxci-production-level3-workers/global/images/docker-firefoxci-gcp-l3-googlecompute-2024-02-05t23-18-22z | 52 |
| projects/taskcluster-imaging/global/images/docker-firefoxci-gcp-l1-googlecompute-2025-06-13t18-31-38z | 85 |
| projects/taskcluster-imaging/global/images/docker-worker-gcp-u14-04-2025-06-16 | 19 |
| projects/taskcluster-imaging/global/images/docker-firefoxci-gcp-lt-googlecompute-2023-04-13t21-30-28z | 1 |


| Worker Pool | Implementation | Version | Total Workers | Total Capacity |
| --- | --- | --- | ---: | ---: |
| **app-services-1/b-linux-amd** | docker-worker | 38.0.5 | 2 | 2 |
| **app-services-1/b-linux-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **app-services-1/decision-gcp** | docker-worker | 38.0.5 | 803 | 1606 |
| **app-services-1/images-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **app-services-3/b-linux-amd** | docker-worker | 38.0.5 | 2 | 2 |
| **app-services-3/b-linux-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **app-services-3/decision-gcp** | docker-worker | 38.0.5 | 2 | 4 |
| **app-services-3/images-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **code-analysis-1/linux-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **code-analysis-3/linux-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **comm-1/b-linux-amd** | docker-worker | 38.0.5 | 2 | 2 |
| **comm-1/b-linux-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **comm-1/b-linux-large-amd** | docker-worker | 38.0.5 | 2 | 2 |
| **comm-1/b-linux-large-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **comm-1/b-linux-xlarge-amd** | docker-worker | 38.0.5 | 2 | 2 |
| **comm-1/b-linux-xlarge-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **comm-1/decision-gcp** | docker-worker | 38.0.5 | 4 | 4 |
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
| **comm-3/b-linux-gcp** | docker-worker | 38.0.5 | 895 | 895 |
| **comm-3/b-linux-large-amd** | docker-worker | 38.0.5 | 2 | 2 |
| **comm-3/b-linux-large-gcp** | docker-worker | 38.0.5 | 3 | 3 |
| **comm-3/b-linux-xlarge-amd** | docker-worker | 38.0.5 | 2 | 2 |
| **comm-3/b-linux-xlarge-gcp** | docker-worker | 38.0.5 | 8 | 8 |
| **comm-3/decision-gcp** | docker-worker | 38.0.5 | 20 | 20 |
| **comm-3/images-gcp** | docker-worker | 38.0.5 | 26 | 26 |
| **comm-t/misc-gcp** | docker-worker | 38.0.5 | 2 | 16 |
| **comm-t/t-linux-large-gcp** | docker-worker | unknown version | 2 | 2 |
| **comm-t/t-linux-large-noscratch-gcp** | docker-worker | unknown version | 2 | 2 |
| **comm-t/t-linux-xlarge-gcp** | docker-worker | unknown version | 2 | 2 |
| **comm-t/t-linux-xlarge-noscratch-gcp** | docker-worker | unknown version | 2 | 2 |
| **comm-t/t-linux-xlarge-source-gcp** | docker-worker | unknown version | 2 | 2 |
| **comm-t/t-linux-xlarge-source-noscratch-gcp** | docker-worker | unknown version | 2 | 2 |
| **gecko-1/b-linux-amd** | docker-worker | 38.0.5 | 2 | 2 |
| **gecko-1/b-linux-gcp** | docker-worker | 38.0.5 | 6 | 6 |
| **gecko-1/b-linux-gcp-relops1411** | docker-worker | 38.0.5 | 2 | 2 |
| **gecko-1/b-linux-gcp-test-bug-1882320** | docker-worker | 38.0.5 | 2 | 2 |
| **gecko-1/b-linux-kvm-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **gecko-1/b-linux-large-amd** | docker-worker | 38.0.5 | 2 | 2 |
| **gecko-1/b-linux-large-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **gecko-1/b-linux-medium-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **gecko-1/b-linux-xlarge-amd** | docker-worker | 38.0.5 | 2 | 2 |
| **gecko-1/b-linux-xlarge-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **gecko-1/b-linux-xlarge-gcp-bug1797804-c2** | docker-worker | 38.0.5 | 2 | 2 |
| **gecko-1/decision-gcp** | docker-worker | 38.0.5 | 6 | 6 |
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
| **gecko-3/b-linux-amd** | docker-worker | 38.0.5 | 2 | 2 |
| **gecko-3/b-linux-gcp** | docker-worker | 38.0.5 | 493 | 493 |
| **gecko-3/b-linux-kvm-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **gecko-3/b-linux-large-amd** | docker-worker | 38.0.5 | 2 | 2 |
| **gecko-3/b-linux-large-gcp** | docker-worker | 38.0.5 | 34 | 34 |
| **gecko-3/b-linux-medium-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **gecko-3/b-linux-xlarge-amd** | docker-worker | 38.0.5 | 2 | 2 |
| **gecko-3/b-linux-xlarge-gcp** | docker-worker | 38.0.5 | 13 | 13 |
| **gecko-3/decision-gcp** | docker-worker | 38.0.5 | 764 | 764 |
| **gecko-3/images-gcp** | docker-worker | 38.0.5 | 48 | 48 |
| **gecko-t/misc-gcp** | docker-worker | 38.0.5 | 21 | 168 |
| **gecko-t/t-linux-kvm-gcp** | docker-worker | unknown version | 2 | 2 |
| **gecko-t/t-linux-kvm-gcp-relops1411** | docker-worker | unknown version | 2 | 2 |
| **gecko-t/t-linux-kvm-noscratch-gcp** | docker-worker | unknown version | 4 | 4 |
| **gecko-t/t-linux-large-gcp** | docker-worker | unknown version | 2 | 2 |
| **gecko-t/t-linux-large-hostub2204-gcp** | docker-worker | 48.3.0 | 2 | 2 |
| **gecko-t/t-linux-large-noscratch-gcp** | docker-worker | unknown version | 775 | 775 |
| **gecko-t/t-linux-xlarge-gcp** | docker-worker | unknown version | 12 | 12 |
| **gecko-t/t-linux-xlarge-noscratch-gcp** | docker-worker | unknown version | 495 | 495 |
| **gecko-t/t-linux-xlarge-source-gcp** | docker-worker | unknown version | 66 | 66 |
| **gecko-t/t-linux-xlarge-source-noscratch-gcp** | docker-worker | unknown version | 2 | 2 |
| **glean-1/b-linux-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **glean-1/decision-gcp** | docker-worker | 38.0.5 | 786 | 1572 |
| **glean-1/images-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **glean-3/b-linux-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **glean-3/decision-gcp** | docker-worker | 38.0.5 | 2 | 4 |
| **glean-3/images-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **infra/build-decision** | docker-worker | 38.0.5 | 1 | 8 |
| **mobile-1/b-linux-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **mobile-1/decision-gcp** | docker-worker | 38.0.5 | 787 | 1574 |
| **mobile-1/images-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **mobile-3/b-linux-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **mobile-3/decision-gcp** | docker-worker | 38.0.5 | 2 | 4 |
| **mobile-3/images-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **mozilla-1/b-linux-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **mozilla-1/decision-gcp** | docker-worker | 38.0.5 | 789 | 1578 |
| **mozilla-1/images-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **mozilla-3/b-linux-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **mozilla-3/decision-gcp** | docker-worker | 38.0.5 | 2 | 4 |
| **mozilla-3/images-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **mozilla-t/t-linux-large-gcp** | docker-worker | unknown version | 2 | 2 |
| **mozilla-t/t-linux-large-noscratch-gcp** | docker-worker | unknown version | 2 | 2 |
| **mozillavpn-1/b-linux-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **mozillavpn-1/b-linux-large-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **mozillavpn-1/decision-gcp** | docker-worker | 38.0.5 | 786 | 1572 |
| **mozillavpn-1/images-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **mozillavpn-3/b-linux-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **mozillavpn-3/b-linux-large-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **mozillavpn-3/decision-gcp** | docker-worker | 38.0.5 | 2 | 4 |
| **mozillavpn-3/images-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **nss-1/decision-gcp** | docker-worker | 38.0.5 | 4 | 8 |
| **nss-1/images-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **nss-1/linux-gcp** | docker-worker | 38.0.5 | 183 | 183 |
| **nss-3/decision-gcp** | docker-worker | 38.0.5 | 2 | 4 |
| **nss-3/images-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **nss-3/linux-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **nss-t/t-linux-xlarge-gcp** | docker-worker | unknown version | 102 | 102 |
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
| **translations-1/decision-gcp** | docker-worker | 38.0.5 | 1737 | 3474 |
| **translations-1/images-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **xpi-1/b-linux-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **xpi-1/decision-gcp** | docker-worker | 38.0.5 | 2 | 4 |
| **xpi-1/images-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **xpi-3/b-linux-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **xpi-3/decision-gcp** | docker-worker | 38.0.5 | 2 | 4 |
| **xpi-3/images-gcp** | docker-worker | 38.0.5 | 2 | 2 |


## Script Worker

Total: `35`



| Worker Pool | Implementation | Version | Total Workers | Total Capacity |
| --- | --- | --- | ---: | ---: |
| **scriptworker-k8s/app-services-3-signing** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/comm-3-balrog** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/comm-3-beetmover** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/comm-3-signing** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/comm-3-tree** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/comm-t-signing** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/enterprise-3-signing** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/enterprise-t-signing** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/gecko-1-beetmover** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/gecko-1-bouncer** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/gecko-3-balrog** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/gecko-3-beetmover** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/gecko-3-bouncer** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/gecko-3-lando** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/gecko-3-pushapk** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/gecko-3-shipit** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/gecko-3-signing** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/gecko-t-signing** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/mobile-3-bitrise** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/mobile-t-signing** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/mozillavpn-1-beetmover** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/mozillavpn-3-beetmover** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
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
| projects/taskcluster-imaging/global/images/generic-2204-wayland-vm-gcp-googlecompute-2023-09-22t17-39-37z | 2 |
|  | 2 |


| Worker Pool | Implementation | Version | Total Workers | Total Capacity |
| --- | --- | --- | ---: | ---: |
| **built-in/fail** |  | No artifacts found | 0 | 0 |
| **built-in/succeed** |  | No artifacts found | 0 | 0 |
| **gecko-t/t-linux-vm-2204-wayland** |  | No artifacts found | 38 | 38 |
| **gecko-t/t-linux-vm-2204-wayland-snap** |  | No artifacts found | 2 | 2 |


## Version not determined [^2]

Total: `19`


Count by image:

| Version | Count |
| :--- | ---: |
| projects/taskcluster-imaging/global/images/gw-fxci-gcp-l1-2404-amd64-googlecompute-alpha | 1 |
|  | 4 |
| /subscriptions/108d46d5-fe9b-4850-9a7d-8c914aa6c1f0/resourceGroups/rg-packer-worker-images/providers/Microsoft.Compute/galleries/win11a6425h2builderalpha/images/win11a6425h2builderalpha/versions/1.0.0 | 2 |
| /subscriptions/108d46d5-fe9b-4850-9a7d-8c914aa6c1f0/resourceGroups/rg-packer-worker-images/providers/Microsoft.Compute/galleries/win11a6425h2builder/images/win11a6425h2builder/versions/1.0.0 | 4 |
| /subscriptions/108d46d5-fe9b-4850-9a7d-8c914aa6c1f0/resourceGroups/rg-packer-worker-images/providers/Microsoft.Compute/galleries/win11a6425h2tester/images/win11a6425h2tester/versions/1.0.0 | 4 |
| /subscriptions/a30e97ab-734a-4f3b-a0e4-c51c0bff0701/resourceGroups/rg-packer-worker-images/providers/Microsoft.Compute/galleries/trusted_win11_a64_25h2_builder/images/trusted_win11_a64_25h2_builder/versions/1.0.0 | 4 |


| Worker Pool | Implementation | Version | Total Workers | Total Capacity |
| --- | --- | --- | ---: | ---: |
| **comm-t/win11-a64-25h2** |  | Version not determined; task not (yet) claimed | 77 | 77 |
| **enterprise-1/win11-a64-25h2-builder** |  | Version not determined; task not (yet) claimed | 77 | 77 |
| **enterprise-3/win11-a64-25h2-builder** |  | Version not determined; task not (yet) claimed | 77 | 77 |
| **gecko-1/win11-a64-25h2-builder** |  | Version not determined; task not (yet) claimed | 75 | 75 |
| **gecko-1/win11-a64-25h2-builder-alpha** |  | Version not determined; task not (yet) claimed | 75 | 75 |
| **gecko-3/win11-a64-25h2-builder** |  | Version not determined; task not (yet) claimed | 75 | 75 |
| **gecko-t/t-linux-2404-relsre** |  | Version not determined; task not (yet) claimed | 3 | 3 |
| **gecko-t/win11-a64-25h2** |  | Version not determined; task not (yet) claimed | 68 | 68 |
| **mozillavpn-1/win11-a64-25h2-builder** |  | Version not determined; task not (yet) claimed | 66 | 66 |
| **mozillavpn-3/win11-a64-25h2-builder** |  | Version not determined; task not (yet) claimed | 63 | 63 |
| **nss-1/win11-a64-25h2** |  | Version not determined; task not (yet) claimed | 73 | 73 |
| **nss-1/win11-a64-25h2-builder** |  | Version not determined; task not (yet) claimed | 56 | 56 |
| **nss-1/win11-a64-25h2-builder-alpha** |  | Version not determined; task not (yet) claimed | 57 | 57 |
| **nss-3/win11-a64-25h2-builder** |  | Version not determined; task not (yet) claimed | 58 | 58 |
| **nss-t/win11-a64-25h2** |  | Version not determined; task not (yet) claimed | 77 | 77 |
| **releng-hardware/gecko-t-linux-netperf-1804** |  | Version not determined; task not (yet) claimed | 0 | 0 |
| **releng-hardware/gecko-t-osx-1015-r8** |  | Version not determined; task not (yet) claimed | 0 | 0 |
| **releng-hardware/gecko-t-osx-1400-r8** |  | Version not determined; task not (yet) claimed | 0 | 0 |
| **releng-hardware/nss-3-b-osx-1015** |  | Version not determined; task not (yet) claimed | 0 | 0 |



[^1]: Those are the pools whose tasks were claimed and resolved by a worker as expected, but the worker did not publish either artifact `public/logs/live_backing.log` nor `public/logs/chain_of_trust.log`, which is the source used to identify the worker implementation.

[^2]: Probing task remains pending after two hours. Those are the pools that were not able to start any worker to claim the task within two hours.

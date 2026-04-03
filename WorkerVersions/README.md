

# Worker Pool Versions


## Generic Worker

Total: `400`

Count by version:

| Version | Count |
| :--- | ---: |
| 36.0.0 | 4 |
| 45.0.0 | 1 |
| 60.3.4 | 15 |
| 61.0.0 | 1 |
| 64.3.0 | 26 |
| 84.1.2 | 7 |
| 87.0.0 | 3 |
| 88.0.2 | 2 |
| 91.0.2 | 2 |
| 96.2.2 | 7 |
| 96.2.3 | 118 |
| 97.0.1 | 214 |


Count by image:

| Version | Count |
| :--- | ---: |
| projects/taskcluster-imaging/global/images/gw-fxci-gcp-l1-arm64-gui-googlecompute-2024-09-18t14-57-52z | 9 |
| projects/taskcluster-imaging/global/images/gw-fxci-gcp-l1-2404-arm64-headless-googlecompute-alpha | 1 |
| /subscriptions/108d46d5-fe9b-4850-9a7d-8c914aa6c1f0/resourceGroups/rg-packer-worker-images/providers/Microsoft.Compute/galleries/win11_a64_24h2_builder_alpha/images/win11_a64_24h2_builder_alpha/versions/1.0.0 | 2 |
| projects/taskcluster-imaging/global/images/gw-fxci-gcp-l1-gui-googlecompute-2025-01-13t22-33-40z | 2 |
| unknown | 88 |
| projects/taskcluster-imaging/global/images/gw-fxci-gcp-l1-2404-amd64-gui-googlecompute-alpha | 1 |
| projects/fxci-production-level3-workers/global/images/gw-fxci-gcp-l3-2404-amd64-headless-googlecompute-alpha | 1 |
| projects/taskcluster-imaging/global/images/gw-fxci-gcp-l1-2404-amd64-headless-googlecompute-alpha | 6 |
| /subscriptions/108d46d5-fe9b-4850-9a7d-8c914aa6c1f0/resourceGroups/rg-packer-worker-images/providers/Microsoft.Compute/galleries/win11_a64_24h2_builder/images/win11_a64_24h2_builder/versions/1.3.0 | 4 |
| projects/fxci-production-level3-workers/global/images/gw-fxci-gcp-l3-arm64-gui-googlecompute-2024-09-18t19-02-58z | 4 |
| projects/taskcluster-imaging/global/images/gw-fxci-gcp-l1-2404-amd64-gui-googlecompute-2026-03-05 | 4 |
| /subscriptions/108d46d5-fe9b-4850-9a7d-8c914aa6c1f0/resourceGroups/rg-packer-worker-images/providers/Microsoft.Compute/galleries/win11_a64_24h2_tester_alpha/images/win11_a64_24h2_tester_alpha/versions/1.0.0 | 4 |
| projects/taskcluster-imaging/global/images/gw-fxci-gcp-l1-2404-amd64-headless-googlecompute-alpha-tc | 7 |
| projects/taskcluster-imaging/global/images/gw-fxci-gcp-l1-gui-googlecompute-2024-08-22t22-48-09z | 8 |
| projects/fxci-production-level3-workers/global/images/gw-fxci-gcp-l3-gui-googlecompute-2024-09-18t05-46-31z | 3 |
| projects/taskcluster-imaging/global/images/gw-fxci-gcp-l1-2404-arm64-headless-googlecompute-2026-03-05 | 14 |
| projects/taskcluster-imaging/global/images/gw-fxci-gcp-l1-2404-amd64-headless-googlecompute-2026-03-05 | 122 |
|  | 35 |
| /subscriptions/a30e97ab-734a-4f3b-a0e4-c51c0bff0701/resourceGroups/rg-packer-worker-images/providers/Microsoft.Compute/galleries/trusted_win11_a64_24h2_builder/images/trusted_win11_a64_24h2_builder/versions/1.3.0 | 4 |
| projects/fxci-production-level3-workers/global/images/gw-fxci-gcp-l3-2404-arm64-headless-googlecompute-2026-03-05 | 6 |
| /subscriptions/108d46d5-fe9b-4850-9a7d-8c914aa6c1f0/resourceGroups/rg-packer-worker-images/providers/Microsoft.Compute/galleries/win11_a64_25h2_builder/images/win11_a64_25h2_builder/versions/1.0.1 | 4 |
| /subscriptions/108d46d5-fe9b-4850-9a7d-8c914aa6c1f0/resourceGroups/rg-packer-worker-images/providers/Microsoft.Compute/galleries/win11_a64_25h2_tester_alpha/images/win11_a64_25h2_tester_alpha/versions/1.0.0 | 4 |
| /subscriptions/108d46d5-fe9b-4850-9a7d-8c914aa6c1f0/resourceGroups/rg-packer-worker-images/providers/Microsoft.Compute/galleries/win11_a64_25h2_tester/images/win11_a64_25h2_tester/versions/1.0.1 | 4 |
| /subscriptions/108d46d5-fe9b-4850-9a7d-8c914aa6c1f0/resourceGroups/rg-packer-worker-images/providers/Microsoft.Compute/galleries/win11_a64_24h2_tester/images/win11_a64_24h2_tester/versions/1.3.0 | 4 |
| projects/fxci-production-level3-workers/global/images/gw-fxci-gcp-l3-2404-amd64-headless-googlecompute-2026-03-05 | 59 |


| Worker Pool | Implementation | Version | Engine | Revision | OS | Arch | GO | Total Workers | Total Capacity |
| --- | --- | --- | --- | --- | --- | --- | --- | ---: | ---: |
| **adhoc-1/decision** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 3 | 3 |
| **adhoc-1/images** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 12 | 12 |
| **adhoc-3/decision** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 2 | 2 |
| **adhoc-3/images** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 2 | 2 |
| **app-services-1/b-linux** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 225 | 225 |
| **app-services-1/b-linux-docker-amd** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 2 | 2 |
| **app-services-1/decision** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 7 | 7 |
| **app-services-1/images** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 2 | 2 |
| **app-services-3/b-linux** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 524 | 524 |
| **app-services-3/b-linux-docker-amd** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 2 | 2 |
| **app-services-3/decision** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 21 | 21 |
| **app-services-3/images** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 2 | 2 |
| **code-analysis-1/decision** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 9 | 9 |
| **code-analysis-1/linux-docker** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 7 | 7 |
| **code-analysis-1/linux-gw-gcp** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | amd64 | 1.22.2 | 46 | 46 |
| **code-analysis-3/decision** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 2 | 2 |
| **code-analysis-3/linux-docker** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 2 | 2 |
| **code-analysis-3/linux-gw-gcp** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | amd64 | 1.22.2 | 2 | 2 |
| **code-coverage/bot** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 250 | 250 |
| **code-review/bot** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 10162 | 10162 |
| **comm-1/b-linux** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 697 | 697 |
| **comm-1/b-linux-aarch64** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | arm64 | 1.26.0 | 2 | 2 |
| **comm-1/b-linux-docker-amd** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 6 | 6 |
| **comm-1/b-linux-docker-large-amd** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 3 | 3 |
| **comm-1/b-linux-docker-xlarge-amd** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 2 | 2 |
| **comm-1/b-linux-gcp-aarch64** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | arm64 | 1.22.2 | 2 | 2 |
| **comm-1/b-linux-large** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 2 | 2 |
| **comm-1/b-linux-xlarge** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 2 | 2 |
| **comm-1/b-win2022** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 2 | 2 |
| **comm-1/decision** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 242 | 242 |
| **comm-1/images** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 2 | 2 |
| **comm-1/images-aarch64** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | arm64 | 1.26.0 | 2 | 2 |
| **comm-1/images-gcp-aarch64** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | arm64 | 1.22.2 | 2 | 2 |
| **comm-2/b-linux** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 2 | 2 |
| **comm-2/b-linux-aarch64** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | arm64 | 1.26.0 | 2 | 2 |
| **comm-2/b-linux-docker-amd** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 2 | 2 |
| **comm-2/b-linux-docker-large-amd** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 2 | 2 |
| **comm-2/b-linux-docker-xlarge-amd** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 2 | 2 |
| **comm-2/b-linux-gcp-aarch64** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | arm64 | 1.22.2 | 2 | 2 |
| **comm-2/b-linux-large** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 2 | 2 |
| **comm-2/b-linux-xlarge** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 2 | 2 |
| **comm-2/decision** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 2 | 2 |
| **comm-2/images** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 2 | 2 |
| **comm-2/images-aarch64** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | arm64 | 1.26.0 | 2 | 2 |
| **comm-2/images-gcp-aarch64** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | arm64 | 1.22.2 | 2 | 2 |
| **comm-3/b-linux** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 2476 | 2476 |
| **comm-3/b-linux-aarch64** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | arm64 | 1.26.0 | 2 | 2 |
| **comm-3/b-linux-docker-amd** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 1397 | 1397 |
| **comm-3/b-linux-docker-large-amd** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 3 | 3 |
| **comm-3/b-linux-docker-xlarge-amd** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 2 | 2 |
| **comm-3/b-linux-gcp-aarch64** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | arm64 | 1.22.2 | 2 | 2 |
| **comm-3/b-linux-large** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 2 | 2 |
| **comm-3/b-linux-xlarge** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 2 | 2 |
| **comm-3/b-win2022** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 508 | 508 |
| **comm-3/decision** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 88 | 88 |
| **comm-3/images** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 4 | 4 |
| **comm-3/images-aarch64** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | arm64 | 1.26.0 | 2 | 2 |
| **comm-3/images-gcp-aarch64** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | arm64 | 1.22.2 | 2 | 2 |
| **comm-t/misc** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 2 | 2 |
| **comm-t/t-linux-arm64-docker** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | arm64 | 1.26.0 | 2 | 2 |
| **comm-t/t-linux-docker** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 2 | 2 |
| **comm-t/t-linux-docker-amd** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 1870 | 1870 |
| **comm-t/t-linux-docker-noscratch** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 3 | 3 |
| **comm-t/t-linux-docker-noscratch-amd** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 4387 | 4387 |
| **comm-t/win11-64-2009** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 2 | 2 |
| **comm-t/win11-64-2009-ssd** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 3 | 3 |
| **comm-t/win11-64-24h2** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 5553 | 5553 |
| **comm-t/win11-64-24h2-ssd** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 371 | 371 |
| **comm-t/win11-64-25h2** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 2 | 2 |
| **comm-t/win11-64-25h2-ssd** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 2 | 2 |
| **comm-t/win11-a64-24h2** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | arm64 | 1.25.7 | 2 | 2 |
| **comm-t/win11-a64-24h2-alpha** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | arm64 | 1.25.7 | 12 | 12 |
| **comm-t/win11-a64-25h2** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | arm64 | 1.25.7 | 2 | 2 |
| **comm-t/win11-a64-25h2-alpha** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | arm64 | 1.25.7 | 4 | 4 |
| **enterprise-1/b-linux** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 2048 | 2048 |
| **enterprise-1/b-linux-aarch64** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | arm64 | 1.26.0 | 2 | 2 |
| **enterprise-1/b-linux-docker-amd** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 856 | 856 |
| **enterprise-1/b-linux-docker-large-amd** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 19 | 19 |
| **enterprise-1/b-linux-docker-xlarge-amd** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 1376 | 1376 |
| **enterprise-1/b-win2022** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 584 | 584 |
| **enterprise-1/decision** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 215 | 215 |
| **enterprise-1/images** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 5 | 5 |
| **enterprise-1/images-aarch64** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | arm64 | 1.26.0 | 2 | 2 |
| **enterprise-1/win11-a64-24h2-builder** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | arm64 | 1.25.7 | 2 | 2 |
| **enterprise-1/win11-a64-25h2-builder** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | arm64 | 1.25.7 | 2 | 2 |
| **enterprise-3/b-linux** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 2344 | 2344 |
| **enterprise-3/b-linux-aarch64** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | arm64 | 1.26.0 | 2 | 2 |
| **enterprise-3/b-linux-docker-amd** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 449 | 449 |
| **enterprise-3/b-linux-docker-large-amd** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 76 | 76 |
| **enterprise-3/b-linux-docker-xlarge-amd** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 519 | 519 |
| **enterprise-3/b-win2022** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 560 | 560 |
| **enterprise-3/decision** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 76 | 76 |
| **enterprise-3/images** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 5 | 5 |
| **enterprise-3/images-aarch64** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | arm64 | 1.26.0 | 2 | 2 |
| **enterprise-3/win11-a64-24h2-builder** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | arm64 | 1.25.7 | 2 | 2 |
| **enterprise-t/misc** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 220 | 220 |
| **enterprise-t/t-linux-docker-16c32gb-amd** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 1159 | 1159 |
| **enterprise-t/t-linux-docker-amd** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 4474 | 4474 |
| **enterprise-t/t-linux-docker-noscratch-amd** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 7072 | 7072 |
| **enterprise-t/win11-64-24h2** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 9322 | 9322 |
| **enterprise-t/win11-64-24h2-alpha** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 2 | 2 |
| **enterprise-t/win11-64-24h2-source** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 21 | 21 |
| **enterprise-t/win11-64-24h2-source-alpha** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 2 | 2 |
| **enterprise-t/win11-64-25h2** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 2 | 2 |
| **enterprise-t/win11-64-25h2-alpha** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 2 | 2 |
| **enterprise-t/win11-64-25h2-source** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 2 | 2 |
| **enterprise-t/win11-64-25h2-source-alpha** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 2 | 2 |
| **gecko-1/b-linux** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 9807 | 9807 |
| **gecko-1/b-linux-2204-kvm-gcp** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | amd64 | 1.22.2 | 2 | 2 |
| **gecko-1/b-linux-aarch64** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | arm64 | 1.26.0 | 2 | 2 |
| **gecko-1/b-linux-docker-alpha** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 2 | 2 |
| **gecko-1/b-linux-docker-amd** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 7665 | 7665 |
| **gecko-1/b-linux-docker-large-amd** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 66 | 66 |
| **gecko-1/b-linux-docker-xlarge-amd** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 864 | 864 |
| **gecko-1/b-linux-gcp-aarch64** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | arm64 | 1.22.2 | 2 | 2 |
| **gecko-1/b-linux-kvm** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 178 | 178 |
| **gecko-1/b-linux-large** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 2 | 2 |
| **gecko-1/b-linux-medium** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 16946 | 16946 |
| **gecko-1/b-linux-xlarge** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 2 | 2 |
| **gecko-1/b-win2022** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 818 | 818 |
| **gecko-1/b-win2022-alpha** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 2 | 2 |
| **gecko-1/b-win2022-headless** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 2 | 2 |
| **gecko-1/b-win2022-xxlarge** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 17 | 17 |
| **gecko-1/decision** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 1511 | 1511 |
| **gecko-1/images** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 38 | 38 |
| **gecko-1/images-aarch64** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | arm64 | 1.26.0 | 3 | 3 |
| **gecko-1/images-gcp-aarch64** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | arm64 | 1.22.2 | 2 | 2 |
| **gecko-1/win11-a64-24h2-builder** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | arm64 | 1.25.7 | 19 | 19 |
| **gecko-1/win11-a64-24h2-builder-alpha** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | arm64 | 1.25.7 | 3 | 3 |
| **gecko-1/win11-a64-25h2-builder** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | arm64 | 1.25.7 | 2 | 2 |
| **gecko-2/b-linux** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 42 | 42 |
| **gecko-2/b-linux-aarch64** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | arm64 | 1.26.0 | 2 | 2 |
| **gecko-2/b-linux-docker-amd** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 122 | 122 |
| **gecko-2/b-linux-docker-large-amd** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 24 | 24 |
| **gecko-2/b-linux-docker-xlarge-amd** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 25 | 25 |
| **gecko-2/b-linux-gcp-aarch64** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | arm64 | 1.22.2 | 2 | 2 |
| **gecko-2/b-linux-kvm** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 7 | 7 |
| **gecko-2/b-linux-large** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 2 | 2 |
| **gecko-2/b-linux-medium** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 135 | 135 |
| **gecko-2/b-linux-xlarge** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 2 | 2 |
| **gecko-2/b-win2022** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 5 | 5 |
| **gecko-2/decision** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 13 | 13 |
| **gecko-2/images** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 2 | 2 |
| **gecko-2/images-aarch64** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | arm64 | 1.26.0 | 2 | 2 |
| **gecko-2/images-gcp-aarch64** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | arm64 | 1.22.2 | 2 | 2 |
| **gecko-3/b-linux** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 21892 | 21892 |
| **gecko-3/b-linux-2204-kvm-gcp** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | amd64 | 1.22.2 | 2 | 2 |
| **gecko-3/b-linux-aarch64** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | arm64 | 1.26.0 | 2 | 2 |
| **gecko-3/b-linux-docker-alpha** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 2 | 2 |
| **gecko-3/b-linux-docker-amd** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 27904 | 27904 |
| **gecko-3/b-linux-docker-large-amd** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 623 | 623 |
| **gecko-3/b-linux-docker-xlarge-amd** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 3230 | 3230 |
| **gecko-3/b-linux-gcp-aarch64** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | arm64 | 1.22.2 | 2 | 2 |
| **gecko-3/b-linux-kvm** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 641 | 641 |
| **gecko-3/b-linux-large** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 16 | 16 |
| **gecko-3/b-linux-medium** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 27905 | 27905 |
| **gecko-3/b-linux-xlarge** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 311 | 311 |
| **gecko-3/b-win2022** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 6800 | 6800 |
| **gecko-3/b-win2022-headless** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 2 | 2 |
| **gecko-3/b-win2022-xxlarge** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 18 | 18 |
| **gecko-3/decision** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 6703 | 6703 |
| **gecko-3/images** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 53 | 53 |
| **gecko-3/images-aarch64** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | arm64 | 1.26.0 | 2 | 2 |
| **gecko-3/images-gcp-aarch64** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | arm64 | 1.22.2 | 2 | 2 |
| **gecko-3/win11-a64-24h2-builder** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | arm64 | 1.25.7 | 193 | 193 |
| **gecko-t/misc** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 955 | 955 |
| **gecko-t/t-linux-2204-wayland** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | amd64 | 1.22.2 | 10625 | 10625 |
| **gecko-t/t-linux-2204-wayland-arm64-relsre** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | arm64 | 1.22.2 | 2 | 2 |
| **gecko-t/t-linux-2204-wayland-relsre** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | amd64 | 1.22.2 | 2 | 2 |
| **gecko-t/t-linux-2204-wayland-root-exp** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | amd64 | 1.22.2 | 2 | 2 |
| **gecko-t/t-linux-2204-wayland-snap** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | amd64 | 1.22.2 | 73 | 73 |
| **gecko-t/t-linux-2404-headless-alpha** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 2 | 2 |
| **gecko-t/t-linux-2404-headless-arm64-alpha** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | arm64 | 1.26.0 | 2 | 2 |
| **gecko-t/t-linux-2404-headless-ssd-alpha** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 2 | 2 |
| **gecko-t/t-linux-2404-wayland** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 2 | 2 |
| **gecko-t/t-linux-2404-wayland-relsre** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 2 | 2 |
| **gecko-t/t-linux-2404-wayland-snap** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 83 | 83 |
| **gecko-t/t-linux-arm64-docker** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | arm64 | 1.26.0 | 28 | 28 |
| **gecko-t/t-linux-docker** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 2273 | 2273 |
| **gecko-t/t-linux-docker-16c32gb-amd** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 1641 | 1641 |
| **gecko-t/t-linux-docker-amd** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 40352 | 40352 |
| **gecko-t/t-linux-docker-amd-alpha** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 2 | 2 |
| **gecko-t/t-linux-docker-kvm** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 64501 | 64501 |
| **gecko-t/t-linux-docker-kvm-alpha** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 2 | 2 |
| **gecko-t/t-linux-docker-noscratch** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 33005 | 33005 |
| **gecko-t/t-linux-docker-noscratch-amd** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 176831 | 176831 |
| **gecko-t/t-linux-docker-noscratch-amd-alpha** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 2 | 2 |
| **gecko-t/t-linux-xlarge-2204-wayland** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | amd64 | 1.22.2 | 7982 | 7982 |
| **gecko-t/t-linux-xlarge-2404-wayland** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 2 | 2 |
| **gecko-t/win10-64-2009** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 7491 | 7491 |
| **gecko-t/win10-64-2009-alpha** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 2 | 2 |
| **gecko-t/win10-64-2009-gpu** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 1050 | 1050 |
| **gecko-t/win10-64-2009-gpu-alpha** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 5 | 5 |
| **gecko-t/win10-64-2009-source** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 54 | 54 |
| **gecko-t/win10-64-2009-source-alpha** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 4 | 4 |
| **gecko-t/win10-64-2009-webgpu** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 2 | 2 |
| **gecko-t/win10-64-2009-webgpu-alpha** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 11 | 11 |
| **gecko-t/win11-64-2009** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 6951 | 6951 |
| **gecko-t/win11-64-2009-alpha** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 2 | 2 |
| **gecko-t/win11-64-2009-gpu** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 4037 | 4037 |
| **gecko-t/win11-64-2009-gpu-alpha** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 2 | 2 |
| **gecko-t/win11-64-2009-source** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 13 | 13 |
| **gecko-t/win11-64-2009-source-alpha** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 2 | 2 |
| **gecko-t/win11-64-2009-ssd** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 2 | 2 |
| **gecko-t/win11-64-2009-ssd-gpu** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 2 | 2 |
| **gecko-t/win11-64-2009-webgpu** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 2 | 2 |
| **gecko-t/win11-64-2009-webgpu-alpha** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 3 | 3 |
| **gecko-t/win11-64-24h2** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 85952 | 85952 |
| **gecko-t/win11-64-24h2-alpha** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 2 | 2 |
| **gecko-t/win11-64-24h2-amd** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 3 | 3 |
| **gecko-t/win11-64-24h2-gpu** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 26267 | 26267 |
| **gecko-t/win11-64-24h2-gpu-alpha** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 2 | 2 |
| **gecko-t/win11-64-24h2-large** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 1662 | 1662 |
| **gecko-t/win11-64-24h2-large-alpha** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 3 | 3 |
| **gecko-t/win11-64-24h2-privileged** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 215 | 215 |
| **gecko-t/win11-64-24h2-privileged-alpha** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 3 | 3 |
| **gecko-t/win11-64-24h2-source** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 1001 | 1001 |
| **gecko-t/win11-64-24h2-source-alpha** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 2 | 2 |
| **gecko-t/win11-64-24h2-ssd** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 2 | 2 |
| **gecko-t/win11-64-24h2-ssd-gpu** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 2 | 2 |
| **gecko-t/win11-64-24h2-unprivileged** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 2 | 2 |
| **gecko-t/win11-64-24h2-unprivileged-alpha** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 16 | 16 |
| **gecko-t/win11-64-24h2-webgpu** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 110 | 110 |
| **gecko-t/win11-64-24h2-webgpu-alpha** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 2 | 2 |
| **gecko-t/win11-64-25h2** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 20682 | 20682 |
| **gecko-t/win11-64-25h2-alpha** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 40 | 40 |
| **gecko-t/win11-64-25h2-amd** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 2 | 2 |
| **gecko-t/win11-64-25h2-gpu** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 10148 | 10148 |
| **gecko-t/win11-64-25h2-gpu-alpha** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 32 | 32 |
| **gecko-t/win11-64-25h2-large** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 1549 | 1549 |
| **gecko-t/win11-64-25h2-large-alpha** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 18 | 18 |
| **gecko-t/win11-64-25h2-privileged** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 2 | 2 |
| **gecko-t/win11-64-25h2-privileged-alpha** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 13 | 13 |
| **gecko-t/win11-64-25h2-source** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 1251 | 1251 |
| **gecko-t/win11-64-25h2-source-alpha** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 6 | 6 |
| **gecko-t/win11-64-25h2-ssd** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 2 | 2 |
| **gecko-t/win11-64-25h2-ssd-gpu** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 2 | 2 |
| **gecko-t/win11-64-25h2-unprivileged** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 2 | 2 |
| **gecko-t/win11-64-25h2-unprivileged-alpha** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 10 | 10 |
| **gecko-t/win11-64-25h2-webgpu** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 5830 | 5830 |
| **gecko-t/win11-64-25h2-webgpu-alpha** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 2 | 2 |
| **gecko-t/win11-a64-24h2** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | arm64 | 1.25.7 | 320 | 320 |
| **gecko-t/win11-a64-24h2-alpha** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | arm64 | 1.25.7 | 9 | 9 |
| **gecko-t/win11-a64-25h2** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | arm64 | 1.25.7 | 196 | 196 |
| **gecko-t/win11-a64-25h2-alpha** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | arm64 | 1.25.7 | 17 | 17 |
| **glean-1/b-linux** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 23 | 23 |
| **glean-1/decision** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 247 | 247 |
| **glean-1/images** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 2 | 2 |
| **glean-3/b-linux** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 14 | 14 |
| **glean-3/decision** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 5 | 5 |
| **glean-3/images** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 2 | 2 |
| **infra/build-decision-alpha** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 2 | 2 |
| **mobile-1/b-linux** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 146 | 146 |
| **mobile-1/decision** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 276 | 276 |
| **mobile-1/images** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 2 | 2 |
| **mobile-3/b-linux** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 240 | 240 |
| **mobile-3/decision** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 144 | 144 |
| **mobile-3/images** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 2 | 2 |
| **mozilla-1/b-linux** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 2 | 2 |
| **mozilla-1/b-win2022** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 2 | 2 |
| **mozilla-1/decision** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 270 | 270 |
| **mozilla-1/images** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 5 | 5 |
| **mozilla-3/b-linux** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 2 | 2 |
| **mozilla-3/b-win2022** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 3 | 3 |
| **mozilla-3/decision** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 419 | 419 |
| **mozilla-3/images** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 2 | 2 |
| **mozilla-t/t-linux-2204-wayland** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | amd64 | 1.22.2 | 571 | 571 |
| **mozilla-t/t-linux-2204-wayland-relsre** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | amd64 | 1.22.2 | 2 | 2 |
| **mozilla-t/t-linux-2404-wayland** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 2 | 2 |
| **mozilla-t/t-linux-arm64-docker** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | arm64 | 1.26.0 | 2 | 2 |
| **mozilla-t/t-linux-docker** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 11 | 11 |
| **mozilla-t/t-linux-docker-amd** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 3 | 3 |
| **mozilla-t/t-linux-docker-noscratch** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 37 | 37 |
| **mozilla-t/t-linux-docker-noscratch-amd** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 3 | 3 |
| **mozillavpn-1/b-linux** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 407 | 407 |
| **mozillavpn-1/b-linux-gcp-gw** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | amd64 | 1.22.2 | 2 | 2 |
| **mozillavpn-1/b-linux-large** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 316 | 316 |
| **mozillavpn-1/b-win2022** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 62 | 62 |
| **mozillavpn-1/b-win2022-alpha** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 3 | 3 |
| **mozillavpn-1/decision** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 288 | 288 |
| **mozillavpn-1/images** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 28 | 28 |
| **mozillavpn-1/win11-a64-24h2-builder** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | arm64 | 1.25.7 | 3 | 3 |
| **mozillavpn-1/win11-a64-25h2-builder** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | arm64 | 1.25.7 | 2 | 2 |
| **mozillavpn-3/b-linux** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 156 | 156 |
| **mozillavpn-3/b-linux-gcp-gw** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | amd64 | 1.22.2 | 2 | 2 |
| **mozillavpn-3/b-linux-large** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 160 | 160 |
| **mozillavpn-3/b-win2022** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 66 | 66 |
| **mozillavpn-3/decision** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 22 | 22 |
| **mozillavpn-3/images** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 22 | 22 |
| **mozillavpn-3/win11-a64-24h2-builder** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | arm64 | 1.25.7 | 3 | 3 |
| **nss-1/b-win2022** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 261 | 261 |
| **nss-1/b-win2022-alpha** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 3 | 3 |
| **nss-1/decision** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 29 | 29 |
| **nss-1/images** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 2 | 2 |
| **nss-1/linux-docker** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 584 | 584 |
| **nss-1/win11-a64-24h2** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | arm64 | 1.25.7 | 3 | 3 |
| **nss-1/win11-a64-24h2-alpha** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | arm64 | 1.25.7 | 4 | 4 |
| **nss-1/win11-a64-24h2-builder** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | arm64 | 1.25.7 | 2 | 2 |
| **nss-1/win11-a64-24h2-builder-alpha** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | arm64 | 1.25.7 | 6 | 6 |
| **nss-1/win11-a64-25h2** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | arm64 | 1.25.7 | 3 | 3 |
| **nss-1/win11-a64-25h2-alpha** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | arm64 | 1.25.7 | 4 | 4 |
| **nss-1/win11-a64-25h2-builder** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | arm64 | 1.25.7 | 2 | 2 |
| **nss-3/b-win2022** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 369 | 369 |
| **nss-3/decision** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 19 | 19 |
| **nss-3/images** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 2 | 2 |
| **nss-3/linux-docker** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 1201 | 1201 |
| **nss-3/win11-a64-24h2-builder** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | arm64 | 1.25.7 | 2 | 2 |
| **nss-t/t-linux-arm64-docker** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | arm64 | 1.26.0 | 2 | 2 |
| **nss-t/t-linux-docker** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 2244 | 2244 |
| **nss-t/t-linux-docker-amd** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 2 | 2 |
| **nss-t/win11-a64-24h2** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | arm64 | 1.25.7 | 3 | 3 |
| **nss-t/win11-a64-24h2-alpha** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | arm64 | 1.25.7 | 4 | 4 |
| **nss-t/win11-a64-25h2** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | arm64 | 1.25.7 | 2 | 2 |
| **nss-t/win11-a64-25h2-alpha** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | arm64 | 1.25.7 | 28 | 28 |
| **proj-autophone/gecko-t-bitbar-gw-perf-a55** | generic-worker | 36.0.0 | simple | b9cb11293f | linux | amd64 | 1.13.7 | 0 | 0 |
| **proj-autophone/gecko-t-bitbar-gw-perf-p6** | generic-worker | 36.0.0 | simple | b9cb11293f | linux | amd64 | 1.13.7 | 0 | 0 |
| **proj-autophone/gecko-t-bitbar-gw-perf-s24** | generic-worker | 36.0.0 | simple | b9cb11293f | linux | amd64 | 1.13.7 | 0 | 0 |
| **proj-autophone/gecko-t-bitbar-gw-unit-p5** | generic-worker | 36.0.0 | simple | b9cb11293f | linux | amd64 | 1.13.7 | 0 | 0 |
| **proj-autophone/gecko-t-lambda-alpha-a55** | generic-worker | 87.0.0 | insecure | 99a1fcafbb | linux | amd64 | 1.24.5 | 0 | 0 |
| **proj-autophone/gecko-t-lambda-perf-a55** | generic-worker | 87.0.0 | insecure | 99a1fcafbb | linux | amd64 | 1.24.5 | 0 | 0 |
| **proj-autophone/gecko-t-lambda-test-1** | generic-worker | 87.0.0 | insecure | 99a1fcafbb | linux | amd64 | 1.24.5 | 0 | 0 |
| **releng-1/decision** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 69 | 69 |
| **releng-1/linux-docker** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 71 | 71 |
| **releng-3/decision** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 34 | 34 |
| **releng-3/linux-docker** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 62 | 62 |
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
| **releng-hardware/gecko-t-osx-1015-r8-staging** | generic-worker | 60.3.4 | simple | 943a6f2b0d | darwin | amd64 | 1.22.0 | 0 | 0 |
| **releng-hardware/gecko-t-osx-1400-r8-staging** | generic-worker | 60.3.4 | simple | 943a6f2b0d | darwin | amd64 | 1.22.0 | 0 | 0 |
| **releng-hardware/gecko-t-osx-1500-m-vms** | generic-worker | 60.3.4 | simple | 943a6f2b0d | darwin | arm64 | 1.22.0 | 0 | 0 |
| **releng-hardware/gecko-t-osx-1500-m4** | generic-worker | 91.0.2 | multiuser | 06628b3721 | darwin | arm64 | 1.25.3 | 0 | 0 |
| **releng-hardware/gecko-t-osx-1500-m4-ipv6** | generic-worker | 60.3.4 | simple | 943a6f2b0d | darwin | arm64 | 1.22.0 | 0 | 0 |
| **releng-hardware/gecko-t-osx-1500-m4-staging** | generic-worker | 91.0.2 | multiuser | 06628b3721 | darwin | arm64 | 1.25.3 | 0 | 0 |
| **releng-hardware/gecko-t-win7-32-hw** | generic-worker | 45.0.0 | multiuser | 988e8100b3 | windows | 386 | 1.19.3 | 0 | 0 |
| **releng-hardware/mozillavpn-b-1-osx** | generic-worker | 60.3.4 | multiuser | 943a6f2b0d | darwin | amd64 | 1.22.0 | 0 | 0 |
| **releng-hardware/mozillavpn-b-3-osx** | generic-worker | 60.3.4 | multiuser | 943a6f2b0d | darwin | amd64 | 1.22.0 | 0 | 0 |
| **releng-hardware/win11-64-24h2-hw** | generic-worker | 96.2.2 | multiuser | 773509947e | windows | amd64 | 1.25.7 | 0 | 0 |
| **releng-hardware/win11-64-24h2-hw-alpha** | generic-worker | 96.2.2 | multiuser | 773509947e | windows | amd64 | 1.25.7 | 0 | 0 |
| **releng-hardware/win11-64-24h2-hw-perf-debug** | generic-worker | 96.2.2 | multiuser | 773509947e | windows | amd64 | 1.25.7 | 0 | 0 |
| **releng-hardware/win11-64-24h2-hw-perf-sheriff** | generic-worker | 96.2.2 | multiuser | 773509947e | windows | amd64 | 1.25.7 | 0 | 0 |
| **releng-hardware/win11-64-24h2-hw-ref** | generic-worker | 96.2.2 | multiuser | 773509947e | windows | amd64 | 1.25.7 | 0 | 0 |
| **releng-hardware/win11-64-24h2-hw-ref-alpha** | generic-worker | 96.2.2 | multiuser | 773509947e | windows | amd64 | 1.25.7 | 0 | 0 |
| **releng-hardware/win11-64-24h2-hw-relops1213** | generic-worker | 96.2.2 | multiuser | 773509947e | windows | amd64 | 1.25.7 | 0 | 0 |
| **releng-t/linux-docker** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 323 | 323 |
| **relops-1/decision** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 2 | 2 |
| **relops-1/images** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 2 | 2 |
| **relops-3/decision** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 3 | 3 |
| **relops-3/images** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 2 | 2 |
| **scriptworker-1/decision** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 15 | 15 |
| **scriptworker-1/images** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 81 | 81 |
| **scriptworker-3/decision** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 5 | 5 |
| **scriptworker-3/images** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 19 | 19 |
| **taskgraph-1/decision** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 26 | 26 |
| **taskgraph-1/images** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 19 | 19 |
| **taskgraph-3/decision** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 9 | 9 |
| **taskgraph-3/images** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 6 | 6 |
| **taskgraph-t/linux-docker** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 202 | 202 |
| **taskgraph-t/win11-64-24h2** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 18 | 18 |
| **taskgraph-t/win11-64-25h2** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 3 | 3 |
| **translations-1/b-linux-large-gcp-1tb-32-256-d2g** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 28 | 28 |
| **translations-1/b-linux-large-gcp-1tb-32-256-std-d2g** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 2 | 2 |
| **translations-1/b-linux-large-gcp-1tb-64-512-d2g** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 2 | 2 |
| **translations-1/b-linux-large-gcp-1tb-64-512-std-d2g** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 2 | 2 |
| **translations-1/b-linux-large-gcp-d2g** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 166 | 166 |
| **translations-1/b-linux-large-gcp-d2g-1tb** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 2 | 2 |
| **translations-1/b-linux-large-gcp-d2g-1tb-standard** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 2 | 2 |
| **translations-1/b-linux-large-gcp-d2g-300gb** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 93 | 93 |
| **translations-1/b-linux-v100-gpu-d2g** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 65 | 65 |
| **translations-1/b-linux-v100-gpu-d2g-4** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 19 | 19 |
| **translations-1/b-linux-v100-gpu-d2g-4-1tb** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 12 | 12 |
| **translations-1/b-linux-v100-gpu-d2g-4-1tb-alpha** | generic-worker | 84.1.2 | multiuser | 4a139f395a | linux | amd64 | 1.24.4 | 2 | 2 |
| **translations-1/b-linux-v100-gpu-d2g-4-1tb-standard** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 2 | 2 |
| **translations-1/b-linux-v100-gpu-d2g-4-1tb-std-alpha** | generic-worker | 84.1.2 | multiuser | 4a139f395a | linux | amd64 | 1.24.4 | 2 | 2 |
| **translations-1/b-linux-v100-gpu-d2g-4-2tb** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 8 | 8 |
| **translations-1/b-linux-v100-gpu-d2g-4-2tb-alpha** | generic-worker | 84.1.2 | multiuser | 4a139f395a | linux | amd64 | 1.24.4 | 2 | 2 |
| **translations-1/b-linux-v100-gpu-d2g-4-300gb** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 30 | 30 |
| **translations-1/b-linux-v100-gpu-d2g-4-300gb-alpha** | generic-worker | 84.1.2 | multiuser | 4a139f395a | linux | amd64 | 1.24.4 | 2 | 2 |
| **translations-1/b-linux-v100-gpu-d2g-4-300gb-standard** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 2 | 2 |
| **translations-1/b-linux-v100-gpu-d2g-4-300gb-std-alpha** | generic-worker | 84.1.2 | multiuser | 4a139f395a | linux | amd64 | 1.24.4 | 2 | 2 |
| **translations-1/b-linux-v100-gpu-d2g-4-alpha** | generic-worker | 84.1.2 | multiuser | 4a139f395a | linux | amd64 | 1.24.4 | 2 | 2 |
| **translations-1/b-linux-v100-gpu-d2g-alpha** | generic-worker | 84.1.2 | multiuser | 4a139f395a | linux | amd64 | 1.24.4 | 2 | 2 |
| **translations-1/decision** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 421 | 421 |
| **translations-1/images** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 2 | 2 |
| **translations-t/t-linux-gw-noscratch-amd** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 2 | 2 |
| **xpi-1/b-linux** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 162 | 162 |
| **xpi-1/decision** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 62 | 62 |
| **xpi-1/images** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 6 | 6 |
| **xpi-3/b-linux** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 27 | 27 |
| **xpi-3/decision** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 13 | 13 |
| **xpi-3/images** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 6 | 6 |


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
| **app-services-1/decision-gcp** | docker-worker | 38.0.5 | 750 | 1500 |
| **app-services-1/images-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **app-services-3/b-linux-amd** | docker-worker | 38.0.5 | 2 | 2 |
| **app-services-3/b-linux-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **app-services-3/decision-gcp** | docker-worker | 38.0.5 | 2 | 4 |
| **app-services-3/images-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **code-analysis-1/linux-gcp** | docker-worker | 38.0.5 | 4 | 4 |
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
| **comm-3/b-linux-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **comm-3/b-linux-large-amd** | docker-worker | 38.0.5 | 2 | 2 |
| **comm-3/b-linux-large-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **comm-3/b-linux-xlarge-amd** | docker-worker | 38.0.5 | 2 | 2 |
| **comm-3/b-linux-xlarge-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **comm-3/decision-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **comm-3/images-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **comm-t/misc-gcp** | docker-worker | 38.0.5 | 2 | 16 |
| **comm-t/t-linux-large-gcp** | docker-worker | unknown version | 2 | 2 |
| **comm-t/t-linux-large-noscratch-gcp** | docker-worker | unknown version | 2 | 2 |
| **comm-t/t-linux-xlarge-gcp** | docker-worker | unknown version | 2 | 2 |
| **comm-t/t-linux-xlarge-noscratch-gcp** | docker-worker | unknown version | 2 | 2 |
| **comm-t/t-linux-xlarge-source-gcp** | docker-worker | unknown version | 2 | 2 |
| **comm-t/t-linux-xlarge-source-noscratch-gcp** | docker-worker | unknown version | 2 | 2 |
| **gecko-1/b-linux-amd** | docker-worker | 38.0.5 | 2 | 2 |
| **gecko-1/b-linux-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **gecko-1/b-linux-gcp-relops1411** | docker-worker | 38.0.5 | 2 | 2 |
| **gecko-1/b-linux-gcp-test-bug-1882320** | docker-worker | 38.0.5 | 2 | 2 |
| **gecko-1/b-linux-kvm-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **gecko-1/b-linux-large-amd** | docker-worker | 38.0.5 | 2 | 2 |
| **gecko-1/b-linux-large-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **gecko-1/b-linux-medium-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **gecko-1/b-linux-xlarge-amd** | docker-worker | 38.0.5 | 2 | 2 |
| **gecko-1/b-linux-xlarge-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **gecko-1/b-linux-xlarge-gcp-bug1797804-c2** | docker-worker | 38.0.5 | 2 | 2 |
| **gecko-1/decision-gcp** | docker-worker | 38.0.5 | 2 | 2 |
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
| **gecko-2/b-linux-xlarge-gcp** | docker-worker | 38.0.5 | 4 | 4 |
| **gecko-2/decision-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **gecko-2/images-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **gecko-3/b-linux-amd** | docker-worker | 38.0.5 | 2 | 2 |
| **gecko-3/b-linux-gcp** | docker-worker | 38.0.5 | 544 | 544 |
| **gecko-3/b-linux-kvm-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **gecko-3/b-linux-large-amd** | docker-worker | 38.0.5 | 2 | 2 |
| **gecko-3/b-linux-large-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **gecko-3/b-linux-medium-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **gecko-3/b-linux-xlarge-amd** | docker-worker | 38.0.5 | 2 | 2 |
| **gecko-3/b-linux-xlarge-gcp** | docker-worker | 38.0.5 | 4 | 4 |
| **gecko-3/decision-gcp** | docker-worker | 38.0.5 | 783 | 783 |
| **gecko-3/images-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **gecko-t/misc-gcp** | docker-worker | 38.0.5 | 17 | 136 |
| **gecko-t/t-linux-kvm-gcp** | docker-worker | unknown version | 2 | 2 |
| **gecko-t/t-linux-kvm-gcp-relops1411** | docker-worker | unknown version | 2 | 2 |
| **gecko-t/t-linux-kvm-noscratch-gcp** | docker-worker | unknown version | 11 | 11 |
| **gecko-t/t-linux-large-gcp** | docker-worker | unknown version | 2 | 2 |
| **gecko-t/t-linux-large-hostub2204-gcp** | docker-worker | 48.3.0 | 2 | 2 |
| **gecko-t/t-linux-large-noscratch-gcp** | docker-worker | unknown version | 3087 | 3087 |
| **gecko-t/t-linux-xlarge-gcp** | docker-worker | unknown version | 53 | 53 |
| **gecko-t/t-linux-xlarge-noscratch-gcp** | docker-worker | unknown version | 1911 | 1911 |
| **gecko-t/t-linux-xlarge-source-gcp** | docker-worker | unknown version | 80 | 80 |
| **gecko-t/t-linux-xlarge-source-noscratch-gcp** | docker-worker | unknown version | 2 | 2 |
| **glean-1/b-linux-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **glean-1/decision-gcp** | docker-worker | 38.0.5 | 729 | 1458 |
| **glean-1/images-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **glean-3/b-linux-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **glean-3/decision-gcp** | docker-worker | 38.0.5 | 2 | 4 |
| **glean-3/images-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **infra/build-decision** | docker-worker | 38.0.5 | 2 | 16 |
| **mobile-1/b-linux-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **mobile-1/decision-gcp** | docker-worker | 38.0.5 | 726 | 1452 |
| **mobile-1/images-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **mobile-3/b-linux-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **mobile-3/decision-gcp** | docker-worker | 38.0.5 | 2 | 4 |
| **mobile-3/images-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **mozilla-1/b-linux-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **mozilla-1/decision-gcp** | docker-worker | 38.0.5 | 728 | 1456 |
| **mozilla-1/images-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **mozilla-3/b-linux-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **mozilla-3/decision-gcp** | docker-worker | 38.0.5 | 2 | 4 |
| **mozilla-3/images-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **mozilla-t/t-linux-large-gcp** | docker-worker | unknown version | 2 | 2 |
| **mozilla-t/t-linux-large-noscratch-gcp** | docker-worker | unknown version | 2 | 2 |
| **mozillavpn-1/b-linux-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **mozillavpn-1/b-linux-large-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **mozillavpn-1/decision-gcp** | docker-worker | 38.0.5 | 735 | 1470 |
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
| **releng-1/b-linux-gcp** | docker-worker | 38.0.5 | 3 | 3 |
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
| **translations-1/decision-gcp** | docker-worker | 38.0.5 | 1507 | 3014 |
| **translations-1/images-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **xpi-1/b-linux-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **xpi-1/decision-gcp** | docker-worker | 38.0.5 | 5 | 10 |
| **xpi-1/images-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **xpi-3/b-linux-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **xpi-3/decision-gcp** | docker-worker | 38.0.5 | 2 | 4 |
| **xpi-3/images-gcp** | docker-worker | 38.0.5 | 2 | 2 |


## Script Worker

Total: `42`



| Worker Pool | Implementation | Version | Total Workers | Total Capacity |
| --- | --- | --- | ---: | ---: |
| **scriptworker-k8s/comm-3-balrog** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/comm-3-beetmover** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/comm-3-bouncer** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/comm-3-pushflatpak** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/comm-3-pushmsix** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
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
| **gecko-t/t-linux-vm-2204-wayland** |  | No artifacts found | 134 | 134 |
| **gecko-t/t-linux-vm-2204-wayland-snap** |  | No artifacts found | 2 | 2 |


## Version not determined [^2]

Total: `12`


Count by image:

| Version | Count |
| :--- | ---: |
| /subscriptions/108d46d5-fe9b-4850-9a7d-8c914aa6c1f0/resourceGroups/rg-packer-worker-images/providers/Microsoft.Compute/galleries/win11_a64_25h2_builder_alpha/images/win11_a64_25h2_builder_alpha/versions/1.0.0 | 2 |
|  | 5 |
| projects/taskcluster-imaging/global/images/gw-fxci-gcp-l1-2404-amd64-googlecompute-alpha | 1 |
| /subscriptions/a30e97ab-734a-4f3b-a0e4-c51c0bff0701/resourceGroups/rg-packer-worker-images/providers/Microsoft.Compute/galleries/trusted_win11_a64_25h2_builder/images/trusted_win11_a64_25h2_builder/versions/1.0.0 | 4 |


| Worker Pool | Implementation | Version | Total Workers | Total Capacity |
| --- | --- | --- | ---: | ---: |
| **enterprise-3/win11-a64-25h2-builder** |  | Version not determined; task not (yet) claimed | 75 | 75 |
| **gecko-1/win11-a64-25h2-builder-alpha** |  | Version not determined; task not (yet) claimed | 74 | 74 |
| **gecko-3/win11-a64-25h2-builder** |  | Version not determined; task not (yet) claimed | 74 | 74 |
| **gecko-t/t-linux-2404-relsre** |  | Version not determined; task not (yet) claimed | 3 | 3 |
| **mozillavpn-3/win11-a64-25h2-builder** |  | Version not determined; task not (yet) claimed | 69 | 69 |
| **nss-1/win11-a64-25h2-builder-alpha** |  | Version not determined; task not (yet) claimed | 66 | 66 |
| **nss-3/win11-a64-25h2-builder** |  | Version not determined; task not (yet) claimed | 61 | 61 |
| **releng-hardware/gecko-t-linux-netperf-1804** |  | Version not determined; task not (yet) claimed | 0 | 0 |
| **releng-hardware/gecko-t-osx-1015-r8** |  | Version not determined; task not (yet) claimed | 0 | 0 |
| **releng-hardware/gecko-t-osx-1400-r8** |  | Version not determined; task not (yet) claimed | 0 | 0 |
| **releng-hardware/nss-1-b-osx-1015** |  | Version not determined; task not (yet) claimed | 0 | 0 |
| **releng-hardware/nss-3-b-osx-1015** |  | Version not determined; task not (yet) claimed | 0 | 0 |



[^1]: Those are the pools whose tasks were claimed and resolved by a worker as expected, but the worker did not publish either artifact `public/logs/live_backing.log` nor `public/logs/chain_of_trust.log`, which is the source used to identify the worker implementation.

[^2]: Probing task remains pending after two hours. Those are the pools that were not able to start any worker to claim the task within two hours.

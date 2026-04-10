

# Worker Pool Versions


## Generic Worker

Total: `380`

Count by version:

| Version | Count |
| :--- | ---: |
| 36.0.0 | 4 |
| 45.0.0 | 1 |
| 60.3.4 | 15 |
| 61.0.0 | 1 |
| 64.3.0 | 26 |
| 65.1.0 | 1 |
| 84.1.2 | 7 |
| 88.0.2 | 2 |
| 91.0.2 | 1 |
| 95.0.3 | 1 |
| 96.2.2 | 7 |
| 96.2.3 | 100 |
| 97.0.1 | 214 |


Count by image:

| Version | Count |
| :--- | ---: |
|  | 33 |
| /subscriptions/108d46d5-fe9b-4850-9a7d-8c914aa6c1f0/resourceGroups/rg-packer-worker-images/providers/Microsoft.Compute/galleries/win11_a64_25h2_builder/images/win11_a64_25h2_builder/versions/1.0.1 | 2 |
| /subscriptions/108d46d5-fe9b-4850-9a7d-8c914aa6c1f0/resourceGroups/rg-packer-worker-images/providers/Microsoft.Compute/galleries/win11_a64_25h2_tester_alpha/images/win11_a64_25h2_tester_alpha/versions/1.0.0 | 1 |
| unknown | 88 |
| projects/taskcluster-imaging/global/images/gw-fxci-gcp-l1-2404-amd64-headless-googlecompute-2026-03-05 | 122 |
| projects/taskcluster-imaging/global/images/gw-fxci-gcp-l1-2404-amd64-headless-googlecompute-alpha | 6 |
| projects/taskcluster-imaging/global/images/gw-fxci-gcp-l1-2404-amd64-headless-googlecompute-alpha-tc | 7 |
| projects/taskcluster-imaging/global/images/gw-fxci-gcp-l1-2404-arm64-headless-googlecompute-alpha | 1 |
| projects/fxci-production-level3-workers/global/images/gw-fxci-gcp-l3-2404-amd64-headless-googlecompute-alpha | 1 |
| projects/fxci-production-level3-workers/global/images/gw-fxci-gcp-l3-arm64-gui-googlecompute-2024-09-18t19-02-58z | 4 |
| /subscriptions/a30e97ab-734a-4f3b-a0e4-c51c0bff0701/resourceGroups/rg-packer-worker-images/providers/Microsoft.Compute/galleries/trusted_win11_a64_24h2_builder/images/trusted_win11_a64_24h2_builder/versions/1.3.0 | 1 |
| projects/fxci-production-level3-workers/global/images/gw-fxci-gcp-l3-2404-amd64-headless-googlecompute-2026-03-05 | 59 |
| /subscriptions/108d46d5-fe9b-4850-9a7d-8c914aa6c1f0/resourceGroups/rg-packer-worker-images/providers/Microsoft.Compute/galleries/win11_a64_25h2_tester/images/win11_a64_25h2_tester/versions/1.0.1 | 2 |
| /subscriptions/108d46d5-fe9b-4850-9a7d-8c914aa6c1f0/resourceGroups/rg-packer-worker-images/providers/Microsoft.Compute/galleries/win11_a64_24h2_tester/images/win11_a64_24h2_tester/versions/1.3.0 | 2 |
| /subscriptions/108d46d5-fe9b-4850-9a7d-8c914aa6c1f0/resourceGroups/rg-packer-worker-images/providers/Microsoft.Compute/galleries/win11_a64_24h2_builder_alpha/images/win11_a64_24h2_builder_alpha/versions/1.0.0 | 2 |
| projects/fxci-production-level3-workers/global/images/gw-fxci-gcp-l3-2404-arm64-headless-googlecompute-2026-03-05 | 6 |
| projects/fxci-production-level3-workers/global/images/gw-fxci-gcp-l3-gui-googlecompute-2024-09-18t05-46-31z | 3 |
| projects/taskcluster-imaging/global/images/gw-fxci-gcp-l1-2404-amd64-gui-googlecompute-2026-03-05 | 4 |
| projects/taskcluster-imaging/global/images/gw-fxci-gcp-l1-arm64-gui-googlecompute-2024-09-18t14-57-52z | 9 |
| /subscriptions/108d46d5-fe9b-4850-9a7d-8c914aa6c1f0/resourceGroups/rg-packer-worker-images/providers/Microsoft.Compute/galleries/win11_a64_24h2_builder/images/win11_a64_24h2_builder/versions/1.3.0 | 2 |
| projects/taskcluster-imaging/global/images/gw-fxci-gcp-l1-gui-googlecompute-2025-01-13t22-33-40z | 2 |
| projects/taskcluster-imaging/global/images/gw-fxci-gcp-l1-gui-googlecompute-2024-08-22t22-48-09z | 8 |
| projects/taskcluster-imaging/global/images/gw-fxci-gcp-l1-2404-arm64-headless-googlecompute-2026-03-05 | 14 |
| projects/taskcluster-imaging/global/images/gw-fxci-gcp-l1-2404-amd64-gui-googlecompute-alpha | 1 |


| Worker Pool | Implementation | Version | Engine | Revision | OS | Arch | GO | Total Workers | Total Capacity |
| --- | --- | --- | --- | --- | --- | --- | --- | ---: | ---: |
| **adhoc-1/decision** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 5 | 5 |
| **adhoc-1/images** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 31 | 31 |
| **adhoc-3/decision** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 3 | 3 |
| **adhoc-3/images** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 3 | 3 |
| **app-services-1/b-linux** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 283 | 283 |
| **app-services-1/b-linux-docker-amd** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 2 | 2 |
| **app-services-1/decision** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 12 | 12 |
| **app-services-1/images** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 2 | 2 |
| **app-services-3/b-linux** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 582 | 582 |
| **app-services-3/b-linux-docker-amd** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 2 | 2 |
| **app-services-3/decision** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 24 | 24 |
| **app-services-3/images** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 2 | 2 |
| **code-analysis-1/decision** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 2 | 2 |
| **code-analysis-1/linux-docker** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 2 | 2 |
| **code-analysis-1/linux-gw-gcp** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | amd64 | 1.22.2 | 31 | 31 |
| **code-analysis-3/decision** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 2 | 2 |
| **code-analysis-3/linux-docker** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 2 | 2 |
| **code-analysis-3/linux-gw-gcp** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | amd64 | 1.22.2 | 6 | 6 |
| **code-coverage/bot** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 248 | 248 |
| **code-review/bot** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 10708 | 10708 |
| **comm-1/b-linux** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 581 | 581 |
| **comm-1/b-linux-aarch64** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | arm64 | 1.26.0 | 6 | 6 |
| **comm-1/b-linux-docker-amd** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 76 | 76 |
| **comm-1/b-linux-docker-large-amd** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 6 | 6 |
| **comm-1/b-linux-docker-xlarge-amd** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 8 | 8 |
| **comm-1/b-linux-gcp-aarch64** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | arm64 | 1.22.2 | 2 | 2 |
| **comm-1/b-linux-large** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 2 | 2 |
| **comm-1/b-linux-xlarge** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 5 | 5 |
| **comm-1/b-win2022** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 9 | 9 |
| **comm-1/decision** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 175 | 175 |
| **comm-1/images** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 16 | 16 |
| **comm-1/images-aarch64** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | arm64 | 1.26.0 | 3 | 3 |
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
| **comm-3/b-linux** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 3402 | 3402 |
| **comm-3/b-linux-aarch64** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | arm64 | 1.26.0 | 2 | 2 |
| **comm-3/b-linux-docker-amd** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 2014 | 2014 |
| **comm-3/b-linux-docker-large-amd** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 6 | 6 |
| **comm-3/b-linux-docker-xlarge-amd** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 6 | 6 |
| **comm-3/b-linux-gcp-aarch64** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | arm64 | 1.22.2 | 2 | 2 |
| **comm-3/b-linux-large** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 2 | 2 |
| **comm-3/b-linux-xlarge** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 3 | 3 |
| **comm-3/b-win2022** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 699 | 699 |
| **comm-3/decision** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 96 | 96 |
| **comm-3/images** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 2 | 2 |
| **comm-3/images-aarch64** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | arm64 | 1.26.0 | 2 | 2 |
| **comm-3/images-gcp-aarch64** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | arm64 | 1.22.2 | 2 | 2 |
| **comm-t/misc** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 2 | 2 |
| **comm-t/t-linux-arm64-docker** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | arm64 | 1.26.0 | 2 | 2 |
| **comm-t/t-linux-docker** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 5 | 5 |
| **comm-t/t-linux-docker-amd** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 1638 | 1638 |
| **comm-t/t-linux-docker-noscratch** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 28 | 28 |
| **comm-t/t-linux-docker-noscratch-amd** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 3926 | 3926 |
| **comm-t/win11-64-2009** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 2 | 2 |
| **comm-t/win11-64-2009-ssd** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 3 | 3 |
| **comm-t/win11-64-24h2** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 4069 | 4069 |
| **comm-t/win11-64-24h2-ssd** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 368 | 368 |
| **comm-t/win11-64-25h2** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 2 | 2 |
| **comm-t/win11-64-25h2-ssd** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 2 | 2 |
| **comm-t/win11-a64-24h2** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | arm64 | 1.25.7 | 2 | 2 |
| **comm-t/win11-a64-25h2** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | arm64 | 1.25.7 | 2 | 2 |
| **enterprise-1/b-linux** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 1675 | 1675 |
| **enterprise-1/b-linux-aarch64** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | arm64 | 1.26.0 | 5 | 5 |
| **enterprise-1/b-linux-docker-amd** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 688 | 688 |
| **enterprise-1/b-linux-docker-large-amd** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 12 | 12 |
| **enterprise-1/b-linux-docker-xlarge-amd** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 946 | 946 |
| **enterprise-1/b-win2022** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 381 | 381 |
| **enterprise-1/decision** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 320 | 320 |
| **enterprise-1/images** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 4 | 4 |
| **enterprise-1/images-aarch64** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | arm64 | 1.26.0 | 2 | 2 |
| **enterprise-1/win11-a64-24h2-builder** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | arm64 | 1.25.7 | 3 | 3 |
| **enterprise-1/win11-a64-25h2-builder** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | arm64 | 1.25.7 | 2 | 2 |
| **enterprise-3/b-linux** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 1825 | 1825 |
| **enterprise-3/b-linux-aarch64** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | arm64 | 1.26.0 | 4 | 4 |
| **enterprise-3/b-linux-docker-amd** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 378 | 378 |
| **enterprise-3/b-linux-docker-large-amd** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 44 | 44 |
| **enterprise-3/b-linux-docker-xlarge-amd** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 395 | 395 |
| **enterprise-3/b-win2022** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 400 | 400 |
| **enterprise-3/decision** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 57 | 57 |
| **enterprise-3/images** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 4 | 4 |
| **enterprise-3/images-aarch64** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | arm64 | 1.26.0 | 2 | 2 |
| **enterprise-3/win11-a64-24h2-builder** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | arm64 | 1.25.7 | 2 | 2 |
| **enterprise-t/misc** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 168 | 168 |
| **enterprise-t/t-linux-docker-16c32gb-amd** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 769 | 769 |
| **enterprise-t/t-linux-docker-amd** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 3378 | 3378 |
| **enterprise-t/t-linux-docker-noscratch-amd** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 5350 | 5350 |
| **enterprise-t/win11-64-24h2** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 6329 | 6329 |
| **enterprise-t/win11-64-24h2-alpha** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 2 | 2 |
| **enterprise-t/win11-64-24h2-source** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 2 | 2 |
| **enterprise-t/win11-64-24h2-source-alpha** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 2 | 2 |
| **enterprise-t/win11-64-25h2** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 2 | 2 |
| **enterprise-t/win11-64-25h2-alpha** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 2 | 2 |
| **enterprise-t/win11-64-25h2-source** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 41 | 41 |
| **enterprise-t/win11-64-25h2-source-alpha** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 2 | 2 |
| **gecko-1/b-linux** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 9697 | 9697 |
| **gecko-1/b-linux-2204-kvm-gcp** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | amd64 | 1.22.2 | 2 | 2 |
| **gecko-1/b-linux-aarch64** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | arm64 | 1.26.0 | 13 | 13 |
| **gecko-1/b-linux-docker-alpha** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 2 | 2 |
| **gecko-1/b-linux-docker-amd** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 8977 | 8977 |
| **gecko-1/b-linux-docker-large-amd** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 156 | 156 |
| **gecko-1/b-linux-docker-xlarge-amd** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 965 | 965 |
| **gecko-1/b-linux-gcp-aarch64** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | arm64 | 1.22.2 | 2 | 2 |
| **gecko-1/b-linux-kvm** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 181 | 181 |
| **gecko-1/b-linux-large** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 5 | 5 |
| **gecko-1/b-linux-medium** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 20039 | 20039 |
| **gecko-1/b-linux-xlarge** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 11 | 11 |
| **gecko-1/b-win2022** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 883 | 883 |
| **gecko-1/b-win2022-alpha** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 2 | 2 |
| **gecko-1/b-win2022-headless** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 2 | 2 |
| **gecko-1/b-win2022-xxlarge** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 14 | 14 |
| **gecko-1/decision** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 1077 | 1077 |
| **gecko-1/images** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 32 | 32 |
| **gecko-1/images-aarch64** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | arm64 | 1.26.0 | 3 | 3 |
| **gecko-1/images-gcp-aarch64** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | arm64 | 1.22.2 | 3 | 3 |
| **gecko-1/win11-a64-24h2-builder** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | arm64 | 1.25.7 | 30 | 30 |
| **gecko-1/win11-a64-24h2-builder-alpha** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | arm64 | 1.25.7 | 2 | 2 |
| **gecko-1/win11-a64-25h2-builder** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | arm64 | 1.25.7 | 2 | 2 |
| **gecko-2/b-linux** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 118 | 118 |
| **gecko-2/b-linux-aarch64** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | arm64 | 1.26.0 | 2 | 2 |
| **gecko-2/b-linux-docker-amd** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 264 | 264 |
| **gecko-2/b-linux-docker-large-amd** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 36 | 36 |
| **gecko-2/b-linux-docker-xlarge-amd** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 37 | 37 |
| **gecko-2/b-linux-gcp-aarch64** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | arm64 | 1.22.2 | 2 | 2 |
| **gecko-2/b-linux-kvm** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 9 | 9 |
| **gecko-2/b-linux-large** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 2 | 2 |
| **gecko-2/b-linux-medium** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 201 | 201 |
| **gecko-2/b-linux-xlarge** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 2 | 2 |
| **gecko-2/b-win2022** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 11 | 11 |
| **gecko-2/decision** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 27 | 27 |
| **gecko-2/images** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 2 | 2 |
| **gecko-2/images-aarch64** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | arm64 | 1.26.0 | 2 | 2 |
| **gecko-2/images-gcp-aarch64** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | arm64 | 1.22.2 | 2 | 2 |
| **gecko-3/b-linux** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 20660 | 20660 |
| **gecko-3/b-linux-2204-kvm-gcp** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | amd64 | 1.22.2 | 2 | 2 |
| **gecko-3/b-linux-aarch64** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | arm64 | 1.26.0 | 4 | 4 |
| **gecko-3/b-linux-docker-alpha** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 2 | 2 |
| **gecko-3/b-linux-docker-amd** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 26303 | 26303 |
| **gecko-3/b-linux-docker-large-amd** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 523 | 523 |
| **gecko-3/b-linux-docker-xlarge-amd** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 3342 | 3342 |
| **gecko-3/b-linux-gcp-aarch64** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | arm64 | 1.22.2 | 3 | 3 |
| **gecko-3/b-linux-kvm** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 595 | 595 |
| **gecko-3/b-linux-large** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 3 | 3 |
| **gecko-3/b-linux-medium** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 24280 | 24280 |
| **gecko-3/b-linux-xlarge** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 279 | 279 |
| **gecko-3/b-win2022** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 6718 | 6718 |
| **gecko-3/b-win2022-headless** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 2 | 2 |
| **gecko-3/b-win2022-xxlarge** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 14 | 14 |
| **gecko-3/decision** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 6171 | 6171 |
| **gecko-3/images** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 8 | 8 |
| **gecko-3/images-aarch64** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | arm64 | 1.26.0 | 2 | 2 |
| **gecko-3/images-gcp-aarch64** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | arm64 | 1.22.2 | 2 | 2 |
| **gecko-t/misc** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 1073 | 1073 |
| **gecko-t/t-linux-2204-wayland** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | amd64 | 1.22.2 | 8910 | 8910 |
| **gecko-t/t-linux-2204-wayland-arm64-relsre** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | arm64 | 1.22.2 | 2 | 2 |
| **gecko-t/t-linux-2204-wayland-relsre** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | amd64 | 1.22.2 | 2 | 2 |
| **gecko-t/t-linux-2204-wayland-root-exp** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | amd64 | 1.22.2 | 2 | 2 |
| **gecko-t/t-linux-2204-wayland-snap** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | amd64 | 1.22.2 | 84 | 84 |
| **gecko-t/t-linux-2404-headless-alpha** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 2 | 2 |
| **gecko-t/t-linux-2404-headless-arm64-alpha** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | arm64 | 1.26.0 | 2 | 2 |
| **gecko-t/t-linux-2404-headless-ssd-alpha** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 2 | 2 |
| **gecko-t/t-linux-2404-wayland** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 2 | 2 |
| **gecko-t/t-linux-2404-wayland-relsre** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 2 | 2 |
| **gecko-t/t-linux-2404-wayland-snap** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 96 | 96 |
| **gecko-t/t-linux-arm64-docker** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | arm64 | 1.26.0 | 30 | 30 |
| **gecko-t/t-linux-docker** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 1828 | 1828 |
| **gecko-t/t-linux-docker-16c32gb-amd** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 1442 | 1442 |
| **gecko-t/t-linux-docker-amd** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 33500 | 33500 |
| **gecko-t/t-linux-docker-amd-alpha** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 2 | 2 |
| **gecko-t/t-linux-docker-kvm** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 62163 | 62163 |
| **gecko-t/t-linux-docker-kvm-alpha** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 2 | 2 |
| **gecko-t/t-linux-docker-noscratch** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 30588 | 30588 |
| **gecko-t/t-linux-docker-noscratch-amd** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 138817 | 138817 |
| **gecko-t/t-linux-docker-noscratch-amd-alpha** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 2 | 2 |
| **gecko-t/t-linux-xlarge-2204-wayland** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | amd64 | 1.22.2 | 6674 | 6674 |
| **gecko-t/t-linux-xlarge-2404-wayland** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 3 | 3 |
| **gecko-t/win10-64-2009** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 5419 | 5419 |
| **gecko-t/win10-64-2009-alpha** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 2 | 2 |
| **gecko-t/win10-64-2009-gpu** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 760 | 760 |
| **gecko-t/win10-64-2009-gpu-alpha** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 7 | 7 |
| **gecko-t/win10-64-2009-source** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 91 | 91 |
| **gecko-t/win10-64-2009-source-alpha** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 17 | 17 |
| **gecko-t/win10-64-2009-webgpu** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 2 | 2 |
| **gecko-t/win10-64-2009-webgpu-alpha** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 8 | 8 |
| **gecko-t/win11-64-2009** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 7184 | 7184 |
| **gecko-t/win11-64-2009-alpha** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 2 | 2 |
| **gecko-t/win11-64-2009-gpu** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 4260 | 4260 |
| **gecko-t/win11-64-2009-gpu-alpha** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 2 | 2 |
| **gecko-t/win11-64-2009-source** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 78 | 78 |
| **gecko-t/win11-64-2009-source-alpha** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 2 | 2 |
| **gecko-t/win11-64-2009-ssd** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 2 | 2 |
| **gecko-t/win11-64-2009-ssd-gpu** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 2 | 2 |
| **gecko-t/win11-64-2009-webgpu** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 2 | 2 |
| **gecko-t/win11-64-2009-webgpu-alpha** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 3 | 3 |
| **gecko-t/win11-64-24h2** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 41569 | 41569 |
| **gecko-t/win11-64-24h2-alpha** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 2 | 2 |
| **gecko-t/win11-64-24h2-amd** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 4 | 4 |
| **gecko-t/win11-64-24h2-gpu** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 20229 | 20229 |
| **gecko-t/win11-64-24h2-gpu-alpha** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 2 | 2 |
| **gecko-t/win11-64-24h2-large** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 1354 | 1354 |
| **gecko-t/win11-64-24h2-large-alpha** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 3 | 3 |
| **gecko-t/win11-64-24h2-privileged** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 190 | 190 |
| **gecko-t/win11-64-24h2-privileged-alpha** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 14 | 14 |
| **gecko-t/win11-64-24h2-source** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 348 | 348 |
| **gecko-t/win11-64-24h2-source-alpha** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 2 | 2 |
| **gecko-t/win11-64-24h2-ssd** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 2 | 2 |
| **gecko-t/win11-64-24h2-ssd-gpu** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 2 | 2 |
| **gecko-t/win11-64-24h2-unprivileged** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 2 | 2 |
| **gecko-t/win11-64-24h2-unprivileged-alpha** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 17 | 17 |
| **gecko-t/win11-64-24h2-webgpu** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 2 | 2 |
| **gecko-t/win11-64-24h2-webgpu-alpha** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 3 | 3 |
| **gecko-t/win11-64-25h2** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 61040 | 61040 |
| **gecko-t/win11-64-25h2-alpha** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 2 | 2 |
| **gecko-t/win11-64-25h2-amd** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 3 | 3 |
| **gecko-t/win11-64-25h2-gpu** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 4888 | 4888 |
| **gecko-t/win11-64-25h2-gpu-alpha** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 2 | 2 |
| **gecko-t/win11-64-25h2-large** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 4946 | 4946 |
| **gecko-t/win11-64-25h2-large-alpha** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 2 | 2 |
| **gecko-t/win11-64-25h2-privileged** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 2 | 2 |
| **gecko-t/win11-64-25h2-privileged-alpha** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 12 | 12 |
| **gecko-t/win11-64-25h2-source** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 1437 | 1437 |
| **gecko-t/win11-64-25h2-source-alpha** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 2 | 2 |
| **gecko-t/win11-64-25h2-ssd** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 2 | 2 |
| **gecko-t/win11-64-25h2-ssd-gpu** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 2 | 2 |
| **gecko-t/win11-64-25h2-unprivileged** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 2 | 2 |
| **gecko-t/win11-64-25h2-unprivileged-alpha** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 14 | 14 |
| **gecko-t/win11-64-25h2-webgpu** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 3889 | 3889 |
| **gecko-t/win11-64-25h2-webgpu-alpha** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 2 | 2 |
| **gecko-t/win11-a64-24h2** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | arm64 | 1.25.7 | 152 | 152 |
| **gecko-t/win11-a64-25h2** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | arm64 | 1.25.7 | 190 | 190 |
| **glean-1/b-linux** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 16 | 16 |
| **glean-1/decision** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 276 | 276 |
| **glean-1/images** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 2 | 2 |
| **glean-3/b-linux** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 18 | 18 |
| **glean-3/decision** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 6 | 6 |
| **glean-3/images** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 2 | 2 |
| **infra/build-decision-alpha** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 2 | 2 |
| **mobile-1/b-linux** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 135 | 135 |
| **mobile-1/decision** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 278 | 278 |
| **mobile-1/images** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 2 | 2 |
| **mobile-3/b-linux** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 225 | 225 |
| **mobile-3/decision** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 120 | 120 |
| **mobile-3/images** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 2 | 2 |
| **mozilla-1/b-linux** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 2 | 2 |
| **mozilla-1/b-win2022** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 2 | 2 |
| **mozilla-1/decision** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 292 | 292 |
| **mozilla-1/images** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 9 | 9 |
| **mozilla-3/b-linux** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 2 | 2 |
| **mozilla-3/b-win2022** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 3 | 3 |
| **mozilla-3/decision** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 432 | 432 |
| **mozilla-3/images** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 2 | 2 |
| **mozilla-t/t-linux-2204-wayland** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | amd64 | 1.22.2 | 595 | 595 |
| **mozilla-t/t-linux-2204-wayland-relsre** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | amd64 | 1.22.2 | 2 | 2 |
| **mozilla-t/t-linux-2404-wayland** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 2 | 2 |
| **mozilla-t/t-linux-arm64-docker** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | arm64 | 1.26.0 | 2 | 2 |
| **mozilla-t/t-linux-docker** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 12 | 12 |
| **mozilla-t/t-linux-docker-amd** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 3 | 3 |
| **mozilla-t/t-linux-docker-noscratch** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 35 | 35 |
| **mozilla-t/t-linux-docker-noscratch-amd** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 2 | 2 |
| **mozillavpn-1/b-linux** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 333 | 333 |
| **mozillavpn-1/b-linux-gcp-gw** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | amd64 | 1.22.2 | 3 | 3 |
| **mozillavpn-1/b-linux-large** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 249 | 249 |
| **mozillavpn-1/b-win2022** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 55 | 55 |
| **mozillavpn-1/b-win2022-alpha** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 4 | 4 |
| **mozillavpn-1/decision** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 301 | 301 |
| **mozillavpn-1/images** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 12 | 12 |
| **mozillavpn-3/b-linux** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 157 | 157 |
| **mozillavpn-3/b-linux-gcp-gw** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | amd64 | 1.22.2 | 2 | 2 |
| **mozillavpn-3/b-linux-large** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 150 | 150 |
| **mozillavpn-3/b-win2022** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 56 | 56 |
| **mozillavpn-3/decision** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 20 | 20 |
| **mozillavpn-3/images** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 12 | 12 |
| **nss-1/b-win2022** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 334 | 334 |
| **nss-1/b-win2022-alpha** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 2 | 2 |
| **nss-1/decision** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 35 | 35 |
| **nss-1/images** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 2 | 2 |
| **nss-1/linux-docker** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 824 | 824 |
| **nss-1/win11-a64-24h2-builder-alpha** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | arm64 | 1.25.7 | 7 | 7 |
| **nss-1/win11-a64-25h2-alpha** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | arm64 | 1.25.7 | 6 | 6 |
| **nss-3/b-win2022** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 339 | 339 |
| **nss-3/decision** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 27 | 27 |
| **nss-3/images** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 2 | 2 |
| **nss-3/linux-docker** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 1484 | 1484 |
| **nss-t/t-linux-arm64-docker** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | arm64 | 1.26.0 | 2 | 2 |
| **nss-t/t-linux-docker** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 2342 | 2342 |
| **nss-t/t-linux-docker-amd** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 2 | 2 |
| **proj-autophone/gecko-t-bitbar-gw-perf-a55** | generic-worker | 36.0.0 | simple | b9cb11293f | linux | amd64 | 1.13.7 | 0 | 0 |
| **proj-autophone/gecko-t-bitbar-gw-perf-p6** | generic-worker | 36.0.0 | simple | b9cb11293f | linux | amd64 | 1.13.7 | 0 | 0 |
| **proj-autophone/gecko-t-bitbar-gw-perf-s24** | generic-worker | 36.0.0 | simple | b9cb11293f | linux | amd64 | 1.13.7 | 0 | 0 |
| **proj-autophone/gecko-t-bitbar-gw-unit-p5** | generic-worker | 36.0.0 | simple | b9cb11293f | linux | amd64 | 1.13.7 | 0 | 0 |
| **releng-1/decision** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 44 | 44 |
| **releng-1/linux-docker** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 57 | 57 |
| **releng-3/decision** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 27 | 27 |
| **releng-3/linux-docker** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 74 | 74 |
| **releng-hardware/applicationservices-b-1-osx1015** | generic-worker | 60.3.4 | multiuser | 943a6f2b0d | darwin | amd64 | 1.22.0 | 0 | 0 |
| **releng-hardware/applicationservices-b-3-osx1015** | generic-worker | 60.3.4 | multiuser | 943a6f2b0d | darwin | amd64 | 1.22.0 | 0 | 0 |
| **releng-hardware/enterprise-1-b-osx-arm64** | generic-worker | 60.3.4 | multiuser | 943a6f2b0d | darwin | arm64 | 1.22.0 | 0 | 0 |
| **releng-hardware/enterprise-3-b-osx-arm64** | generic-worker | 60.3.4 | multiuser | 943a6f2b0d | darwin | arm64 | 1.22.0 | 0 | 0 |
| **releng-hardware/gecko-1-b-osx-1015** | generic-worker | 60.3.4 | multiuser | 943a6f2b0d | darwin | amd64 | 1.22.0 | 0 | 0 |
| **releng-hardware/gecko-1-b-osx-1015-staging** | generic-worker | 60.3.4 | multiuser | 943a6f2b0d | darwin | amd64 | 1.22.0 | 0 | 0 |
| **releng-hardware/gecko-3-b-osx-1015** | generic-worker | 60.3.4 | multiuser | 943a6f2b0d | darwin | amd64 | 1.22.0 | 0 | 0 |
| **releng-hardware/gecko-3-b-osx-arm64** | generic-worker | 60.3.4 | multiuser | 943a6f2b0d | darwin | arm64 | 1.22.0 | 0 | 0 |
| **releng-hardware/gecko-t-linux-netperf-1804** | generic-worker | 65.1.0 | insecure | 1a085daa37 | linux | amd64 | 1.22.3 | 0 | 0 |
| **releng-hardware/gecko-t-linux-netperf-2404** | generic-worker | 88.0.2 | insecure | fcaf4a25fc | linux | amd64 | 1.24.5 | 0 | 0 |
| **releng-hardware/gecko-t-linux-talos-1804** | generic-worker | 61.0.0 | simple | 3bd4419b4b | linux | amd64 | 1.22.1 | 0 | 0 |
| **releng-hardware/gecko-t-linux-talos-2404** | generic-worker | 88.0.2 | insecure | fcaf4a25fc | linux | amd64 | 1.24.5 | 0 | 0 |
| **releng-hardware/gecko-t-osx-1015-r8-staging** | generic-worker | 60.3.4 | simple | 943a6f2b0d | darwin | amd64 | 1.22.0 | 0 | 0 |
| **releng-hardware/gecko-t-osx-1400-r8-staging** | generic-worker | 60.3.4 | simple | 943a6f2b0d | darwin | amd64 | 1.22.0 | 0 | 0 |
| **releng-hardware/gecko-t-osx-1500-m-vms** | generic-worker | 60.3.4 | simple | 943a6f2b0d | darwin | arm64 | 1.22.0 | 0 | 0 |
| **releng-hardware/gecko-t-osx-1500-m4** | generic-worker | 95.0.3 | multiuser | a996b8269b | darwin | arm64 | 1.25.5 | 0 | 0 |
| **releng-hardware/gecko-t-osx-1500-m4-ipv6** | generic-worker | 60.3.4 | simple | 943a6f2b0d | darwin | arm64 | 1.22.0 | 0 | 0 |
| **releng-hardware/gecko-t-osx-1500-m4-staging** | generic-worker | 91.0.2 | multiuser | 06628b3721 | darwin | arm64 | 1.25.3 | 0 | 0 |
| **releng-hardware/gecko-t-win7-32-hw** | generic-worker | 45.0.0 | multiuser | 988e8100b3 | windows | 386 | 1.19.3 | 0 | 0 |
| **releng-hardware/mozillavpn-b-1-osx** | generic-worker | 60.3.4 | multiuser | 943a6f2b0d | darwin | amd64 | 1.22.0 | 0 | 0 |
| **releng-hardware/mozillavpn-b-3-osx** | generic-worker | 60.3.4 | multiuser | 943a6f2b0d | darwin | amd64 | 1.22.0 | 0 | 0 |
| **releng-hardware/nss-1-b-osx-1015** | generic-worker | 60.3.4 | multiuser | 943a6f2b0d | darwin | amd64 | 1.22.0 | 0 | 0 |
| **releng-hardware/win11-64-24h2-hw** | generic-worker | 96.2.2 | multiuser | 773509947e | windows | amd64 | 1.25.7 | 0 | 0 |
| **releng-hardware/win11-64-24h2-hw-alpha** | generic-worker | 96.2.2 | multiuser | 773509947e | windows | amd64 | 1.25.7 | 0 | 0 |
| **releng-hardware/win11-64-24h2-hw-perf-debug** | generic-worker | 96.2.2 | multiuser | 773509947e | windows | amd64 | 1.25.7 | 0 | 0 |
| **releng-hardware/win11-64-24h2-hw-perf-sheriff** | generic-worker | 96.2.2 | multiuser | 773509947e | windows | amd64 | 1.25.7 | 0 | 0 |
| **releng-hardware/win11-64-24h2-hw-ref** | generic-worker | 96.2.2 | multiuser | 773509947e | windows | amd64 | 1.25.7 | 0 | 0 |
| **releng-hardware/win11-64-24h2-hw-ref-alpha** | generic-worker | 96.2.2 | multiuser | 773509947e | windows | amd64 | 1.25.7 | 0 | 0 |
| **releng-hardware/win11-64-24h2-hw-relops1213** | generic-worker | 96.2.2 | multiuser | 773509947e | windows | amd64 | 1.25.7 | 0 | 0 |
| **releng-t/linux-docker** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 247 | 247 |
| **relops-1/decision** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 2 | 2 |
| **relops-1/images** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 2 | 2 |
| **relops-3/decision** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 2 | 2 |
| **relops-3/images** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 2 | 2 |
| **scriptworker-1/decision** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 16 | 16 |
| **scriptworker-1/images** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 83 | 83 |
| **scriptworker-3/decision** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 10 | 10 |
| **scriptworker-3/images** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 52 | 52 |
| **taskgraph-1/decision** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 22 | 22 |
| **taskgraph-1/images** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 21 | 21 |
| **taskgraph-3/decision** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 8 | 8 |
| **taskgraph-3/images** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 8 | 8 |
| **taskgraph-t/linux-docker** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 135 | 135 |
| **taskgraph-t/win11-64-24h2** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 10 | 10 |
| **taskgraph-t/win11-64-25h2** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 2 | 2 |
| **translations-1/b-linux-large-gcp-1tb-32-256-d2g** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 8 | 8 |
| **translations-1/b-linux-large-gcp-1tb-32-256-std-d2g** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 2 | 2 |
| **translations-1/b-linux-large-gcp-1tb-64-512-d2g** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 2 | 2 |
| **translations-1/b-linux-large-gcp-1tb-64-512-std-d2g** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 2 | 2 |
| **translations-1/b-linux-large-gcp-d2g** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 54 | 54 |
| **translations-1/b-linux-large-gcp-d2g-1tb** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 2 | 2 |
| **translations-1/b-linux-large-gcp-d2g-1tb-standard** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 2 | 2 |
| **translations-1/b-linux-large-gcp-d2g-300gb** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 32 | 32 |
| **translations-1/b-linux-v100-gpu-d2g** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 15 | 15 |
| **translations-1/b-linux-v100-gpu-d2g-4** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 5 | 5 |
| **translations-1/b-linux-v100-gpu-d2g-4-1tb** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 4 | 4 |
| **translations-1/b-linux-v100-gpu-d2g-4-1tb-alpha** | generic-worker | 84.1.2 | multiuser | 4a139f395a | linux | amd64 | 1.24.4 | 2 | 2 |
| **translations-1/b-linux-v100-gpu-d2g-4-1tb-standard** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 2 | 2 |
| **translations-1/b-linux-v100-gpu-d2g-4-1tb-std-alpha** | generic-worker | 84.1.2 | multiuser | 4a139f395a | linux | amd64 | 1.24.4 | 2 | 2 |
| **translations-1/b-linux-v100-gpu-d2g-4-2tb** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 3 | 3 |
| **translations-1/b-linux-v100-gpu-d2g-4-2tb-alpha** | generic-worker | 84.1.2 | multiuser | 4a139f395a | linux | amd64 | 1.24.4 | 2 | 2 |
| **translations-1/b-linux-v100-gpu-d2g-4-300gb** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 11 | 11 |
| **translations-1/b-linux-v100-gpu-d2g-4-300gb-alpha** | generic-worker | 84.1.2 | multiuser | 4a139f395a | linux | amd64 | 1.24.4 | 2 | 2 |
| **translations-1/b-linux-v100-gpu-d2g-4-300gb-standard** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 2 | 2 |
| **translations-1/b-linux-v100-gpu-d2g-4-300gb-std-alpha** | generic-worker | 84.1.2 | multiuser | 4a139f395a | linux | amd64 | 1.24.4 | 2 | 2 |
| **translations-1/b-linux-v100-gpu-d2g-4-alpha** | generic-worker | 84.1.2 | multiuser | 4a139f395a | linux | amd64 | 1.24.4 | 2 | 2 |
| **translations-1/b-linux-v100-gpu-d2g-alpha** | generic-worker | 84.1.2 | multiuser | 4a139f395a | linux | amd64 | 1.24.4 | 2 | 2 |
| **translations-1/decision** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 529 | 529 |
| **translations-1/images** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 2 | 2 |
| **translations-t/t-linux-gw-noscratch-amd** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 2 | 2 |
| **xpi-1/b-linux** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 118 | 118 |
| **xpi-1/decision** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 67 | 67 |
| **xpi-1/images** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 2 | 2 |
| **xpi-3/b-linux** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 6 | 6 |
| **xpi-3/decision** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 8 | 8 |
| **xpi-3/images** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 2 | 2 |


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
| **app-services-1/decision-gcp** | docker-worker | 38.0.5 | 769 | 1538 |
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
| **comm-1/decision-gcp** | docker-worker | 38.0.5 | 2 | 2 |
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
| **comm-3/b-linux-gcp** | docker-worker | 38.0.5 | 362 | 362 |
| **comm-3/b-linux-large-amd** | docker-worker | 38.0.5 | 2 | 2 |
| **comm-3/b-linux-large-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **comm-3/b-linux-xlarge-amd** | docker-worker | 38.0.5 | 2 | 2 |
| **comm-3/b-linux-xlarge-gcp** | docker-worker | 38.0.5 | 3 | 3 |
| **comm-3/decision-gcp** | docker-worker | 38.0.5 | 6 | 6 |
| **comm-3/images-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **comm-t/misc-gcp** | docker-worker | 38.0.5 | 2 | 16 |
| **comm-t/t-linux-large-gcp** | docker-worker | unknown version | 2 | 2 |
| **comm-t/t-linux-large-noscratch-gcp** | docker-worker | unknown version | 2 | 2 |
| **comm-t/t-linux-xlarge-gcp** | docker-worker | unknown version | 2 | 2 |
| **comm-t/t-linux-xlarge-noscratch-gcp** | docker-worker | unknown version | 2 | 2 |
| **comm-t/t-linux-xlarge-source-gcp** | docker-worker | unknown version | 2 | 2 |
| **comm-t/t-linux-xlarge-source-noscratch-gcp** | docker-worker | unknown version | 2 | 2 |
| **gecko-1/b-linux-amd** | docker-worker | 38.0.5 | 2 | 2 |
| **gecko-1/b-linux-gcp** | docker-worker | 38.0.5 | 60 | 60 |
| **gecko-1/b-linux-gcp-relops1411** | docker-worker | 38.0.5 | 2 | 2 |
| **gecko-1/b-linux-gcp-test-bug-1882320** | docker-worker | 38.0.5 | 2 | 2 |
| **gecko-1/b-linux-kvm-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **gecko-1/b-linux-large-amd** | docker-worker | 38.0.5 | 2 | 2 |
| **gecko-1/b-linux-large-gcp** | docker-worker | 38.0.5 | 3 | 3 |
| **gecko-1/b-linux-medium-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **gecko-1/b-linux-xlarge-amd** | docker-worker | 38.0.5 | 2 | 2 |
| **gecko-1/b-linux-xlarge-gcp** | docker-worker | 38.0.5 | 8 | 8 |
| **gecko-1/b-linux-xlarge-gcp-bug1797804-c2** | docker-worker | 38.0.5 | 2 | 2 |
| **gecko-1/decision-gcp** | docker-worker | 38.0.5 | 4 | 4 |
| **gecko-1/decision-gcp-c3d4-amd-bug1990935** | docker-worker | 38.0.5 | 2 | 2 |
| **gecko-1/decision-gcp-c3d8-amd-bug1990935** | docker-worker | 38.0.5 | 2 | 2 |
| **gecko-1/decision-gcp-c4d4-amd-bug1990935** | docker-worker | 38.0.5 | 2 | 2 |
| **gecko-1/decision-gcp-c4d4-hcpu-amd-bug1990935** | docker-worker | 38.0.5 | 2 | 2 |
| **gecko-1/decision-gcp-c4d8-amd-bug1990935** | docker-worker | 38.0.5 | 2 | 2 |
| **gecko-1/decision-gcp-c4d8-hcpu-amd-bug1990935** | docker-worker | 38.0.5 | 2 | 2 |
| **gecko-1/decision-gcp-c4d8-lssd-amd-bug1990935** | docker-worker | 38.0.5 | 2 | 2 |
| **gecko-1/images-gcp** | docker-worker | 38.0.5 | 32 | 32 |
| **gecko-2/b-linux-amd** | docker-worker | 38.0.5 | 2 | 2 |
| **gecko-2/b-linux-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **gecko-2/b-linux-kvm-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **gecko-2/b-linux-large-amd** | docker-worker | 38.0.5 | 2 | 2 |
| **gecko-2/b-linux-large-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **gecko-2/b-linux-medium-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **gecko-2/b-linux-xlarge-amd** | docker-worker | 38.0.5 | 2 | 2 |
| **gecko-2/b-linux-xlarge-gcp** | docker-worker | 38.0.5 | 8 | 8 |
| **gecko-2/decision-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **gecko-2/images-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **gecko-3/b-linux-amd** | docker-worker | 38.0.5 | 2 | 2 |
| **gecko-3/b-linux-gcp** | docker-worker | 38.0.5 | 1049 | 1049 |
| **gecko-3/b-linux-kvm-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **gecko-3/b-linux-large-amd** | docker-worker | 38.0.5 | 2 | 2 |
| **gecko-3/b-linux-large-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **gecko-3/b-linux-medium-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **gecko-3/b-linux-xlarge-amd** | docker-worker | 38.0.5 | 2 | 2 |
| **gecko-3/b-linux-xlarge-gcp** | docker-worker | 38.0.5 | 4 | 4 |
| **gecko-3/decision-gcp** | docker-worker | 38.0.5 | 807 | 807 |
| **gecko-3/images-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **gecko-t/misc-gcp** | docker-worker | 38.0.5 | 26 | 208 |
| **gecko-t/t-linux-kvm-gcp** | docker-worker | unknown version | 2 | 2 |
| **gecko-t/t-linux-kvm-gcp-relops1411** | docker-worker | unknown version | 2 | 2 |
| **gecko-t/t-linux-kvm-noscratch-gcp** | docker-worker | unknown version | 8 | 8 |
| **gecko-t/t-linux-large-gcp** | docker-worker | unknown version | 2 | 2 |
| **gecko-t/t-linux-large-hostub2204-gcp** | docker-worker | 48.3.0 | 2 | 2 |
| **gecko-t/t-linux-large-noscratch-gcp** | docker-worker | unknown version | 2683 | 2683 |
| **gecko-t/t-linux-xlarge-gcp** | docker-worker | unknown version | 42 | 42 |
| **gecko-t/t-linux-xlarge-noscratch-gcp** | docker-worker | unknown version | 1645 | 1645 |
| **gecko-t/t-linux-xlarge-source-gcp** | docker-worker | unknown version | 92 | 92 |
| **gecko-t/t-linux-xlarge-source-noscratch-gcp** | docker-worker | unknown version | 2 | 2 |
| **glean-1/b-linux-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **glean-1/decision-gcp** | docker-worker | 38.0.5 | 753 | 1506 |
| **glean-1/images-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **glean-3/b-linux-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **glean-3/decision-gcp** | docker-worker | 38.0.5 | 2 | 4 |
| **glean-3/images-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **infra/build-decision** | docker-worker | 38.0.5 | 1 | 8 |
| **mobile-1/b-linux-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **mobile-1/decision-gcp** | docker-worker | 38.0.5 | 755 | 1510 |
| **mobile-1/images-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **mobile-3/b-linux-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **mobile-3/decision-gcp** | docker-worker | 38.0.5 | 2 | 4 |
| **mobile-3/images-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **mozilla-1/b-linux-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **mozilla-1/decision-gcp** | docker-worker | 38.0.5 | 755 | 1510 |
| **mozilla-1/images-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **mozilla-3/b-linux-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **mozilla-3/decision-gcp** | docker-worker | 38.0.5 | 2 | 4 |
| **mozilla-3/images-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **mozilla-t/t-linux-large-gcp** | docker-worker | unknown version | 2 | 2 |
| **mozilla-t/t-linux-large-noscratch-gcp** | docker-worker | unknown version | 3 | 3 |
| **mozillavpn-1/b-linux-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **mozillavpn-1/b-linux-large-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **mozillavpn-1/decision-gcp** | docker-worker | 38.0.5 | 760 | 1520 |
| **mozillavpn-1/images-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **mozillavpn-3/b-linux-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **mozillavpn-3/b-linux-large-gcp** | docker-worker | 38.0.5 | 3 | 3 |
| **mozillavpn-3/decision-gcp** | docker-worker | 38.0.5 | 2 | 4 |
| **mozillavpn-3/images-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **nss-1/decision-gcp** | docker-worker | 38.0.5 | 2 | 4 |
| **nss-1/images-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **nss-1/linux-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **nss-3/decision-gcp** | docker-worker | 38.0.5 | 2 | 4 |
| **nss-3/images-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **nss-3/linux-gcp** | docker-worker | 38.0.5 | 11 | 11 |
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
| **scriptworker-3/b-linux-gcp** | docker-worker | 38.0.5 | 62 | 62 |
| **scriptworker-3/decision-gcp** | docker-worker | 38.0.5 | 5 | 10 |
| **scriptworker-3/images-gcp** | docker-worker | 38.0.5 | 14 | 14 |
| **taskgraph-1/decision-gcp** | docker-worker | 38.0.5 | 2 | 4 |
| **taskgraph-1/images-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **taskgraph-3/decision-gcp** | docker-worker | 38.0.5 | 2 | 4 |
| **taskgraph-3/images-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **translations-1/decision-gcp** | docker-worker | 38.0.5 | 1624 | 3248 |
| **translations-1/images-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **xpi-1/b-linux-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **xpi-1/decision-gcp** | docker-worker | 38.0.5 | 2 | 4 |
| **xpi-1/images-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **xpi-3/b-linux-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **xpi-3/decision-gcp** | docker-worker | 38.0.5 | 2 | 4 |
| **xpi-3/images-gcp** | docker-worker | 38.0.5 | 2 | 2 |


## Script Worker

Total: `51`



| Worker Pool | Implementation | Version | Total Workers | Total Capacity |
| --- | --- | --- | ---: | ---: |
| **scriptworker-k8s/adhoc-t-signing** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/app-services-3-beetmover** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/app-services-3-signing** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
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
| **scriptworker-k8s/xpi-1-lando-dev** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/xpi-3-balrog** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/xpi-3-beetmover** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
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
|  | 2 |
| projects/taskcluster-imaging/global/images/generic-2204-wayland-vm-gcp-googlecompute-2023-09-22t17-39-37z | 2 |


| Worker Pool | Implementation | Version | Total Workers | Total Capacity |
| --- | --- | --- | ---: | ---: |
| **built-in/fail** |  | No artifacts found | 0 | 0 |
| **built-in/succeed** |  | No artifacts found | 0 | 0 |
| **gecko-t/t-linux-vm-2204-wayland** |  | No artifacts found | 262 | 262 |
| **gecko-t/t-linux-vm-2204-wayland-snap** |  | No artifacts found | 2 | 2 |


## Version not determined [^2]

Total: `31`


Count by image:

| Version | Count |
| :--- | ---: |
| /subscriptions/a30e97ab-734a-4f3b-a0e4-c51c0bff0701/resourceGroups/rg-packer-worker-images/providers/Microsoft.Compute/galleries/trusted_win11_a64_25h2_builder/images/trusted_win11_a64_25h2_builder/versions/1.0.0 | 4 |
| /subscriptions/108d46d5-fe9b-4850-9a7d-8c914aa6c1f0/resourceGroups/rg-packer-worker-images/providers/Microsoft.Compute/galleries/win11_a64_25h2_builder/images/win11_a64_25h2_builder/versions/1.0.1 | 2 |
| projects/taskcluster-imaging/global/images/gw-fxci-gcp-l1-2404-amd64-googlecompute-alpha | 1 |
| /subscriptions/108d46d5-fe9b-4850-9a7d-8c914aa6c1f0/resourceGroups/rg-packer-worker-images/providers/Microsoft.Compute/galleries/win11_a64_24h2_builder/images/win11_a64_24h2_builder/versions/1.3.0 | 2 |
| /subscriptions/108d46d5-fe9b-4850-9a7d-8c914aa6c1f0/resourceGroups/rg-packer-worker-images/providers/Microsoft.Compute/galleries/win11_a64_24h2_tester_alpha/images/win11_a64_24h2_tester_alpha/versions/1.0.0 | 4 |
|  | 6 |
| /subscriptions/108d46d5-fe9b-4850-9a7d-8c914aa6c1f0/resourceGroups/rg-packer-worker-images/providers/Microsoft.Compute/galleries/win11_a64_24h2_tester/images/win11_a64_24h2_tester/versions/1.3.0 | 2 |
| /subscriptions/a30e97ab-734a-4f3b-a0e4-c51c0bff0701/resourceGroups/rg-packer-worker-images/providers/Microsoft.Compute/galleries/trusted_win11_a64_24h2_builder/images/trusted_win11_a64_24h2_builder/versions/1.3.0 | 3 |
| /subscriptions/108d46d5-fe9b-4850-9a7d-8c914aa6c1f0/resourceGroups/rg-packer-worker-images/providers/Microsoft.Compute/galleries/win11_a64_25h2_tester_alpha/images/win11_a64_25h2_tester_alpha/versions/1.0.0 | 3 |
| /subscriptions/108d46d5-fe9b-4850-9a7d-8c914aa6c1f0/resourceGroups/rg-packer-worker-images/providers/Microsoft.Compute/galleries/win11_a64_25h2_builder_alpha/images/win11_a64_25h2_builder_alpha/versions/1.0.0 | 2 |
| /subscriptions/108d46d5-fe9b-4850-9a7d-8c914aa6c1f0/resourceGroups/rg-packer-worker-images/providers/Microsoft.Compute/galleries/win11_a64_25h2_tester/images/win11_a64_25h2_tester/versions/1.0.1 | 2 |


| Worker Pool | Implementation | Version | Total Workers | Total Capacity |
| --- | --- | --- | ---: | ---: |
| **comm-t/win11-a64-24h2-alpha** |  | Version not determined; task not (yet) claimed | 8 | 8 |
| **comm-t/win11-a64-25h2-alpha** |  | Version not determined; task not (yet) claimed | 3 | 3 |
| **enterprise-3/win11-a64-25h2-builder** |  | Version not determined; task not (yet) claimed | 79 | 79 |
| **gecko-1/win11-a64-25h2-builder-alpha** |  | Version not determined; task not (yet) claimed | 79 | 79 |
| **gecko-3/win11-a64-24h2-builder** |  | Version not determined; task not (yet) claimed | 102 | 102 |
| **gecko-3/win11-a64-25h2-builder** |  | Version not determined; task not (yet) claimed | 78 | 78 |
| **gecko-t/t-linux-2404-relsre** |  | Version not determined; task not (yet) claimed | 3 | 3 |
| **gecko-t/win11-a64-24h2-alpha** |  | Version not determined; task not (yet) claimed | 8 | 8 |
| **gecko-t/win11-a64-25h2-alpha** |  | Version not determined; task not (yet) claimed | 11 | 11 |
| **mozillavpn-1/win11-a64-24h2-builder** |  | Version not determined; task not (yet) claimed | 3 | 3 |
| **mozillavpn-1/win11-a64-25h2-builder** |  | Version not determined; task not (yet) claimed | 3 | 3 |
| **mozillavpn-3/win11-a64-24h2-builder** |  | Version not determined; task not (yet) claimed | 3 | 3 |
| **mozillavpn-3/win11-a64-25h2-builder** |  | Version not determined; task not (yet) claimed | 76 | 76 |
| **nss-1/win11-a64-24h2** |  | Version not determined; task not (yet) claimed | 4 | 4 |
| **nss-1/win11-a64-24h2-alpha** |  | Version not determined; task not (yet) claimed | 4 | 4 |
| **nss-1/win11-a64-24h2-builder** |  | Version not determined; task not (yet) claimed | 3 | 3 |
| **nss-1/win11-a64-25h2** |  | Version not determined; task not (yet) claimed | 4 | 4 |
| **nss-1/win11-a64-25h2-builder** |  | Version not determined; task not (yet) claimed | 3 | 3 |
| **nss-1/win11-a64-25h2-builder-alpha** |  | Version not determined; task not (yet) claimed | 71 | 71 |
| **nss-3/win11-a64-24h2-builder** |  | Version not determined; task not (yet) claimed | 3 | 3 |
| **nss-3/win11-a64-25h2-builder** |  | Version not determined; task not (yet) claimed | 67 | 67 |
| **nss-t/win11-a64-24h2** |  | Version not determined; task not (yet) claimed | 6 | 6 |
| **nss-t/win11-a64-24h2-alpha** |  | Version not determined; task not (yet) claimed | 5 | 5 |
| **nss-t/win11-a64-25h2** |  | Version not determined; task not (yet) claimed | 4 | 4 |
| **nss-t/win11-a64-25h2-alpha** |  | Version not determined; task not (yet) claimed | 12 | 12 |
| **proj-autophone/gecko-t-lambda-alpha-a55** |  | Version not determined; task not (yet) claimed | 0 | 0 |
| **proj-autophone/gecko-t-lambda-perf-a55** |  | Version not determined; task not (yet) claimed | 0 | 0 |
| **releng-hardware/gecko-1-b-osx-arm64** |  | Version not determined; task not (yet) claimed | 0 | 0 |
| **releng-hardware/gecko-t-osx-1015-r8** |  | Version not determined; task not (yet) claimed | 0 | 0 |
| **releng-hardware/gecko-t-osx-1400-r8** |  | Version not determined; task not (yet) claimed | 0 | 0 |
| **releng-hardware/nss-3-b-osx-1015** |  | Version not determined; task not (yet) claimed | 0 | 0 |



[^1]: Those are the pools whose tasks were claimed and resolved by a worker as expected, but the worker did not publish either artifact `public/logs/live_backing.log` nor `public/logs/chain_of_trust.log`, which is the source used to identify the worker implementation.

[^2]: Probing task remains pending after two hours. Those are the pools that were not able to start any worker to claim the task within two hours.

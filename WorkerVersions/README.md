

# Worker Pool Versions


## Generic Worker

Total: `390`

Count by version:

| Version | Count |
| :--- | ---: |
| 36.0.0 | 4 |
| 45.0.0 | 1 |
| 60.3.4 | 17 |
| 61.0.0 | 1 |
| 64.3.0 | 26 |
| 84.1.2 | 7 |
| 87.0.0 | 2 |
| 88.0.2 | 2 |
| 91.0.2 | 1 |
| 95.0.3 | 1 |
| 96.2.2 | 7 |
| 96.2.3 | 110 |
| 97.0.1 | 211 |


Count by image:

| Version | Count |
| :--- | ---: |
| projects/taskcluster-imaging/global/images/gw-fxci-gcp-l1-gui-googlecompute-2025-01-13t22-33-40z | 2 |
| projects/fxci-production-level3-workers/global/images/gw-fxci-gcp-l3-gui-googlecompute-2024-09-18t05-46-31z | 3 |
| projects/taskcluster-imaging/global/images/gw-fxci-gcp-l1-2404-amd64-headless-googlecompute-alpha-tc | 7 |
| /subscriptions/108d46d5-fe9b-4850-9a7d-8c914aa6c1f0/resourceGroups/rg-packer-worker-images/providers/Microsoft.Compute/galleries/win11_a64_24h2_tester/images/win11_a64_24h2_tester/versions/1.3.0 | 4 |
| projects/taskcluster-imaging/global/images/gw-fxci-gcp-l1-2404-arm64-headless-googlecompute-2026-03-05 | 14 |
| projects/fxci-production-level3-workers/global/images/gw-fxci-gcp-l3-2404-arm64-headless-googlecompute-2026-03-05 | 6 |
| projects/taskcluster-imaging/global/images/gw-fxci-gcp-l1-arm64-gui-googlecompute-2024-09-18t14-57-52z | 9 |
| projects/fxci-production-level3-workers/global/images/gw-fxci-gcp-l3-arm64-gui-googlecompute-2024-09-18t19-02-58z | 4 |
|  | 36 |
| unknown | 88 |
| /subscriptions/108d46d5-fe9b-4850-9a7d-8c914aa6c1f0/resourceGroups/rg-packer-worker-images/providers/Microsoft.Compute/galleries/win11_a64_24h2_builder_alpha/images/win11_a64_24h2_builder_alpha/versions/1.0.0 | 2 |
| /subscriptions/108d46d5-fe9b-4850-9a7d-8c914aa6c1f0/resourceGroups/rg-packer-worker-images/providers/Microsoft.Compute/galleries/win11_a64_25h2_tester_alpha/images/win11_a64_25h2_tester_alpha/versions/1.0.0 | 4 |
| projects/taskcluster-imaging/global/images/gw-fxci-gcp-l1-2404-arm64-headless-googlecompute-alpha | 1 |
| projects/taskcluster-imaging/global/images/gw-fxci-gcp-l1-2404-amd64-gui-googlecompute-2026-03-05 | 4 |
| projects/taskcluster-imaging/global/images/gw-fxci-gcp-l1-2404-amd64-headless-googlecompute-2026-03-05 | 120 |
| /subscriptions/108d46d5-fe9b-4850-9a7d-8c914aa6c1f0/resourceGroups/rg-packer-worker-images/providers/Microsoft.Compute/galleries/win11_a64_24h2_tester_alpha/images/win11_a64_24h2_tester_alpha/versions/1.0.0 | 4 |
| projects/taskcluster-imaging/global/images/gw-fxci-gcp-l1-2404-amd64-headless-googlecompute-alpha | 6 |
| projects/taskcluster-imaging/global/images/gw-fxci-gcp-l1-gui-googlecompute-2024-08-22t22-48-09z | 8 |
| /subscriptions/a30e97ab-734a-4f3b-a0e4-c51c0bff0701/resourceGroups/rg-packer-worker-images/providers/Microsoft.Compute/galleries/trusted_win11_a64_24h2_builder/images/trusted_win11_a64_24h2_builder/versions/1.3.0 | 4 |
| projects/taskcluster-imaging/global/images/gw-fxci-gcp-l1-2404-amd64-gui-googlecompute-alpha | 1 |
| projects/fxci-production-level3-workers/global/images/gw-fxci-gcp-l3-2404-amd64-headless-googlecompute-alpha | 1 |
| /subscriptions/108d46d5-fe9b-4850-9a7d-8c914aa6c1f0/resourceGroups/rg-packer-worker-images/providers/Microsoft.Compute/galleries/win11_a64_24h2_builder/images/win11_a64_24h2_builder/versions/1.3.0 | 4 |
| projects/fxci-production-level3-workers/global/images/gw-fxci-gcp-l3-2404-amd64-headless-googlecompute-2026-03-05 | 58 |


| Worker Pool | Implementation | Version | Engine | Revision | OS | Arch | GO | Total Workers | Total Capacity |
| --- | --- | --- | --- | --- | --- | --- | --- | ---: | ---: |
| **adhoc-1/decision** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 7 | 7 |
| **adhoc-1/images** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 34 | 34 |
| **adhoc-3/decision** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 5 | 5 |
| **adhoc-3/images** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 5 | 5 |
| **app-services-1/b-linux** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 200 | 200 |
| **app-services-1/b-linux-docker-amd** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 2 | 2 |
| **app-services-1/decision** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 5 | 5 |
| **app-services-1/images** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 2 | 2 |
| **app-services-3/b-linux** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 339 | 339 |
| **app-services-3/b-linux-docker-amd** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 2 | 2 |
| **app-services-3/decision** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 15 | 15 |
| **app-services-3/images** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 2 | 2 |
| **code-analysis-1/linux-docker** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 2 | 2 |
| **code-analysis-1/linux-gw-gcp** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | amd64 | 1.22.2 | 30 | 30 |
| **code-analysis-3/linux-docker** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 2 | 2 |
| **code-analysis-3/linux-gw-gcp** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | amd64 | 1.22.2 | 2 | 2 |
| **code-coverage/bot** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 465 | 465 |
| **code-review/bot** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 11363 | 11363 |
| **comm-1/b-linux** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 805 | 805 |
| **comm-1/b-linux-aarch64** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | arm64 | 1.26.0 | 2 | 2 |
| **comm-1/b-linux-docker-amd** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 10 | 10 |
| **comm-1/b-linux-docker-large-amd** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 15 | 15 |
| **comm-1/b-linux-docker-xlarge-amd** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 2 | 2 |
| **comm-1/b-linux-gcp-aarch64** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | arm64 | 1.22.2 | 2 | 2 |
| **comm-1/b-linux-large** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 2 | 2 |
| **comm-1/b-linux-xlarge** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 2 | 2 |
| **comm-1/b-win2022** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 5 | 5 |
| **comm-1/decision** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 291 | 291 |
| **comm-1/images** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 21 | 21 |
| **comm-1/images-aarch64** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | arm64 | 1.26.0 | 2 | 2 |
| **comm-1/images-gcp-aarch64** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | arm64 | 1.22.2 | 2 | 2 |
| **comm-2/b-linux** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 2 | 2 |
| **comm-2/b-linux-aarch64** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | arm64 | 1.26.0 | 2 | 2 |
| **comm-2/b-linux-docker-amd** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 2 | 2 |
| **comm-2/b-linux-docker-large-amd** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 2 | 2 |
| **comm-2/b-linux-docker-xlarge-amd** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 2 | 2 |
| **comm-2/b-linux-gcp-aarch64** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | arm64 | 1.22.2 | 2 | 2 |
| **comm-2/b-linux-large** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 3 | 3 |
| **comm-2/b-linux-xlarge** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 2 | 2 |
| **comm-2/decision** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 2 | 2 |
| **comm-2/images** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 2 | 2 |
| **comm-2/images-aarch64** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | arm64 | 1.26.0 | 2 | 2 |
| **comm-2/images-gcp-aarch64** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | arm64 | 1.22.2 | 2 | 2 |
| **comm-3/b-linux** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 2631 | 2631 |
| **comm-3/b-linux-aarch64** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | arm64 | 1.26.0 | 2 | 2 |
| **comm-3/b-linux-docker-amd** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 1698 | 1698 |
| **comm-3/b-linux-docker-large-amd** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 7 | 7 |
| **comm-3/b-linux-docker-xlarge-amd** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 6 | 6 |
| **comm-3/b-linux-gcp-aarch64** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | arm64 | 1.22.2 | 2 | 2 |
| **comm-3/b-linux-large** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 3 | 3 |
| **comm-3/b-linux-xlarge** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 2 | 2 |
| **comm-3/b-win2022** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 549 | 549 |
| **comm-3/decision** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 90 | 90 |
| **comm-3/images** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 22 | 22 |
| **comm-3/images-aarch64** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | arm64 | 1.26.0 | 2 | 2 |
| **comm-3/images-gcp-aarch64** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | arm64 | 1.22.2 | 2 | 2 |
| **comm-t/misc** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 2 | 2 |
| **comm-t/t-linux-arm64-docker** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | arm64 | 1.26.0 | 2 | 2 |
| **comm-t/t-linux-docker** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 23 | 23 |
| **comm-t/t-linux-docker-amd** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 2460 | 2460 |
| **comm-t/t-linux-docker-noscratch** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 2 | 2 |
| **comm-t/t-linux-docker-noscratch-amd** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 4854 | 4854 |
| **comm-t/win11-64-2009** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 2 | 2 |
| **comm-t/win11-64-2009-ssd** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 2 | 2 |
| **comm-t/win11-64-24h2** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 6104 | 6104 |
| **comm-t/win11-64-24h2-ssd** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 388 | 388 |
| **comm-t/win11-64-25h2** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 2 | 2 |
| **comm-t/win11-64-25h2-ssd** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 2 | 2 |
| **comm-t/win11-a64-24h2** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | arm64 | 1.25.7 | 2 | 2 |
| **comm-t/win11-a64-24h2-alpha** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | arm64 | 1.25.7 | 10 | 10 |
| **comm-t/win11-a64-25h2-alpha** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | arm64 | 1.25.7 | 2 | 2 |
| **enterprise-1/b-linux** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 2234 | 2234 |
| **enterprise-1/b-linux-aarch64** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | arm64 | 1.26.0 | 5 | 5 |
| **enterprise-1/b-linux-docker-amd** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 929 | 929 |
| **enterprise-1/b-linux-docker-large-amd** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 23 | 23 |
| **enterprise-1/b-linux-docker-xlarge-amd** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 1118 | 1118 |
| **enterprise-1/b-win2022** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 459 | 459 |
| **enterprise-1/decision** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 53 | 53 |
| **enterprise-1/images** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 41 | 41 |
| **enterprise-1/images-aarch64** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | arm64 | 1.26.0 | 5 | 5 |
| **enterprise-1/win11-a64-24h2-builder** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | arm64 | 1.25.7 | 3 | 3 |
| **enterprise-3/b-linux** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 980 | 980 |
| **enterprise-3/b-linux-aarch64** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | arm64 | 1.26.0 | 2 | 2 |
| **enterprise-3/b-linux-docker-amd** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 254 | 254 |
| **enterprise-3/b-linux-docker-large-amd** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 9 | 9 |
| **enterprise-3/b-linux-docker-xlarge-amd** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 294 | 294 |
| **enterprise-3/b-win2022** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 254 | 254 |
| **enterprise-3/decision** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 55 | 55 |
| **enterprise-3/images** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 5 | 5 |
| **enterprise-3/images-aarch64** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | arm64 | 1.26.0 | 2 | 2 |
| **enterprise-3/win11-a64-24h2-builder** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | arm64 | 1.25.7 | 2 | 2 |
| **enterprise-t/misc** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 162 | 162 |
| **enterprise-t/t-linux-docker-16c32gb-amd** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 918 | 918 |
| **enterprise-t/t-linux-docker-amd** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 4123 | 4123 |
| **enterprise-t/t-linux-docker-noscratch-amd** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 6258 | 6258 |
| **enterprise-t/win11-64-24h2** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 2459 | 2459 |
| **enterprise-t/win11-64-24h2-alpha** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 2 | 2 |
| **enterprise-t/win11-64-24h2-source** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 66 | 66 |
| **enterprise-t/win11-64-24h2-source-alpha** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 2 | 2 |
| **enterprise-t/win11-64-25h2** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 2 | 2 |
| **enterprise-t/win11-64-25h2-alpha** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 2 | 2 |
| **enterprise-t/win11-64-25h2-source** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 2 | 2 |
| **enterprise-t/win11-64-25h2-source-alpha** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 2 | 2 |
| **gecko-1/b-linux** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 8065 | 8065 |
| **gecko-1/b-linux-2204-kvm-gcp** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | amd64 | 1.22.2 | 2 | 2 |
| **gecko-1/b-linux-aarch64** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | arm64 | 1.26.0 | 12 | 12 |
| **gecko-1/b-linux-docker-alpha** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 2 | 2 |
| **gecko-1/b-linux-docker-amd** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 7457 | 7457 |
| **gecko-1/b-linux-docker-large-amd** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 121 | 121 |
| **gecko-1/b-linux-docker-xlarge-amd** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 1006 | 1006 |
| **gecko-1/b-linux-gcp-aarch64** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | arm64 | 1.22.2 | 2 | 2 |
| **gecko-1/b-linux-kvm** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 252 | 252 |
| **gecko-1/b-linux-large** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 2 | 2 |
| **gecko-1/b-linux-medium** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 16947 | 16947 |
| **gecko-1/b-linux-xlarge** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 9 | 9 |
| **gecko-1/b-win2022** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 896 | 896 |
| **gecko-1/b-win2022-alpha** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 2 | 2 |
| **gecko-1/b-win2022-headless** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 2 | 2 |
| **gecko-1/b-win2022-xxlarge** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 8 | 8 |
| **gecko-1/decision** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 1465 | 1465 |
| **gecko-1/images** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 42 | 42 |
| **gecko-1/images-aarch64** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | arm64 | 1.26.0 | 5 | 5 |
| **gecko-1/images-gcp-aarch64** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | arm64 | 1.22.2 | 2 | 2 |
| **gecko-1/win11-a64-24h2-builder** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | arm64 | 1.25.7 | 26 | 26 |
| **gecko-1/win11-a64-24h2-builder-alpha** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | arm64 | 1.25.7 | 3 | 3 |
| **gecko-2/b-linux** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 2 | 2 |
| **gecko-2/b-linux-aarch64** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | arm64 | 1.26.0 | 2 | 2 |
| **gecko-2/b-linux-docker-amd** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 2 | 2 |
| **gecko-2/b-linux-docker-large-amd** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 2 | 2 |
| **gecko-2/b-linux-docker-xlarge-amd** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 2 | 2 |
| **gecko-2/b-linux-gcp-aarch64** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | arm64 | 1.22.2 | 2 | 2 |
| **gecko-2/b-linux-kvm** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 2 | 2 |
| **gecko-2/b-linux-large** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 2 | 2 |
| **gecko-2/b-linux-medium** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 2 | 2 |
| **gecko-2/b-linux-xlarge** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 2 | 2 |
| **gecko-2/b-win2022** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 2 | 2 |
| **gecko-2/decision** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 10 | 10 |
| **gecko-2/images** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 2 | 2 |
| **gecko-2/images-aarch64** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | arm64 | 1.26.0 | 2 | 2 |
| **gecko-2/images-gcp-aarch64** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | arm64 | 1.22.2 | 2 | 2 |
| **gecko-3/b-linux** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 20353 | 20353 |
| **gecko-3/b-linux-2204-kvm-gcp** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | amd64 | 1.22.2 | 2 | 2 |
| **gecko-3/b-linux-aarch64** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | arm64 | 1.26.0 | 9 | 9 |
| **gecko-3/b-linux-docker-alpha** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 2 | 2 |
| **gecko-3/b-linux-docker-amd** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 21661 | 21661 |
| **gecko-3/b-linux-docker-large-amd** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 464 | 464 |
| **gecko-3/b-linux-docker-xlarge-amd** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 2411 | 2411 |
| **gecko-3/b-linux-gcp-aarch64** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | arm64 | 1.22.2 | 2 | 2 |
| **gecko-3/b-linux-kvm** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 470 | 470 |
| **gecko-3/b-linux-large** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 2 | 2 |
| **gecko-3/b-linux-medium** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 23140 | 23140 |
| **gecko-3/b-linux-xlarge** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 85 | 85 |
| **gecko-3/b-win2022** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 6307 | 6307 |
| **gecko-3/b-win2022-headless** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 2 | 2 |
| **gecko-3/b-win2022-xxlarge** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 18 | 18 |
| **gecko-3/decision** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 6104 | 6104 |
| **gecko-3/images** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 106 | 106 |
| **gecko-3/images-aarch64** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | arm64 | 1.26.0 | 11 | 11 |
| **gecko-3/images-gcp-aarch64** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | arm64 | 1.22.2 | 2 | 2 |
| **gecko-3/win11-a64-24h2-builder** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | arm64 | 1.25.7 | 87 | 87 |
| **gecko-t/misc** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 879 | 879 |
| **gecko-t/t-linux-2204-wayland** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | amd64 | 1.22.2 | 8445 | 8445 |
| **gecko-t/t-linux-2204-wayland-arm64-relsre** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | arm64 | 1.22.2 | 2 | 2 |
| **gecko-t/t-linux-2204-wayland-relsre** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | amd64 | 1.22.2 | 2 | 2 |
| **gecko-t/t-linux-2204-wayland-root-exp** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | amd64 | 1.22.2 | 2 | 2 |
| **gecko-t/t-linux-2204-wayland-snap** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | amd64 | 1.22.2 | 143 | 143 |
| **gecko-t/t-linux-2404-headless-alpha** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 2 | 2 |
| **gecko-t/t-linux-2404-headless-arm64-alpha** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | arm64 | 1.26.0 | 2 | 2 |
| **gecko-t/t-linux-2404-headless-ssd-alpha** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 2 | 2 |
| **gecko-t/t-linux-2404-wayland** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 2 | 2 |
| **gecko-t/t-linux-2404-wayland-relsre** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 2 | 2 |
| **gecko-t/t-linux-2404-wayland-snap** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 152 | 152 |
| **gecko-t/t-linux-arm64-docker** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | arm64 | 1.26.0 | 34 | 34 |
| **gecko-t/t-linux-docker** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 1747 | 1747 |
| **gecko-t/t-linux-docker-16c32gb-amd** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 2267 | 2267 |
| **gecko-t/t-linux-docker-amd** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 37734 | 37734 |
| **gecko-t/t-linux-docker-amd-alpha** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 2 | 2 |
| **gecko-t/t-linux-docker-kvm** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 66501 | 66501 |
| **gecko-t/t-linux-docker-kvm-alpha** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 2 | 2 |
| **gecko-t/t-linux-docker-noscratch** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 11053 | 11053 |
| **gecko-t/t-linux-docker-noscratch-amd** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 127890 | 127890 |
| **gecko-t/t-linux-docker-noscratch-amd-alpha** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 2 | 2 |
| **gecko-t/t-linux-xlarge-2204-wayland** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | amd64 | 1.22.2 | 10597 | 10597 |
| **gecko-t/t-linux-xlarge-2404-wayland** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 2 | 2 |
| **gecko-t/win10-64-2009** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 5182 | 5182 |
| **gecko-t/win10-64-2009-alpha** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 2 | 2 |
| **gecko-t/win10-64-2009-gpu** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 630 | 630 |
| **gecko-t/win10-64-2009-gpu-alpha** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 2 | 2 |
| **gecko-t/win10-64-2009-source** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 108 | 108 |
| **gecko-t/win10-64-2009-source-alpha** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 6 | 6 |
| **gecko-t/win10-64-2009-webgpu** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 2 | 2 |
| **gecko-t/win10-64-2009-webgpu-alpha** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 22 | 22 |
| **gecko-t/win11-64-2009** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 2488 | 2488 |
| **gecko-t/win11-64-2009-alpha** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 2 | 2 |
| **gecko-t/win11-64-2009-gpu** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 1526 | 1526 |
| **gecko-t/win11-64-2009-gpu-alpha** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 2 | 2 |
| **gecko-t/win11-64-2009-source** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 2 | 2 |
| **gecko-t/win11-64-2009-source-alpha** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 2 | 2 |
| **gecko-t/win11-64-2009-ssd** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 2 | 2 |
| **gecko-t/win11-64-2009-ssd-gpu** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 2 | 2 |
| **gecko-t/win11-64-2009-webgpu** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 2 | 2 |
| **gecko-t/win11-64-2009-webgpu-alpha** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 8 | 8 |
| **gecko-t/win11-64-24h2** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 71371 | 71371 |
| **gecko-t/win11-64-24h2-alpha** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 2 | 2 |
| **gecko-t/win11-64-24h2-amd** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 3 | 3 |
| **gecko-t/win11-64-24h2-gpu** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 26778 | 26778 |
| **gecko-t/win11-64-24h2-gpu-alpha** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 2 | 2 |
| **gecko-t/win11-64-24h2-large** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 2612 | 2612 |
| **gecko-t/win11-64-24h2-large-alpha** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 6 | 6 |
| **gecko-t/win11-64-24h2-privileged** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 244 | 244 |
| **gecko-t/win11-64-24h2-privileged-alpha** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 11 | 11 |
| **gecko-t/win11-64-24h2-source** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 2545 | 2545 |
| **gecko-t/win11-64-24h2-source-alpha** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 2 | 2 |
| **gecko-t/win11-64-24h2-ssd** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 2 | 2 |
| **gecko-t/win11-64-24h2-ssd-gpu** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 2 | 2 |
| **gecko-t/win11-64-24h2-unprivileged** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 2 | 2 |
| **gecko-t/win11-64-24h2-unprivileged-alpha** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 7 | 7 |
| **gecko-t/win11-64-24h2-webgpu** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 3837 | 3837 |
| **gecko-t/win11-64-24h2-webgpu-alpha** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 2 | 2 |
| **gecko-t/win11-64-25h2** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 4732 | 4732 |
| **gecko-t/win11-64-25h2-alpha** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 1657 | 1657 |
| **gecko-t/win11-64-25h2-amd** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 2 | 2 |
| **gecko-t/win11-64-25h2-gpu** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 1447 | 1447 |
| **gecko-t/win11-64-25h2-gpu-alpha** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 430 | 430 |
| **gecko-t/win11-64-25h2-large** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 31 | 31 |
| **gecko-t/win11-64-25h2-large-alpha** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 116 | 116 |
| **gecko-t/win11-64-25h2-privileged** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 2 | 2 |
| **gecko-t/win11-64-25h2-privileged-alpha** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 11 | 11 |
| **gecko-t/win11-64-25h2-source** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 64 | 64 |
| **gecko-t/win11-64-25h2-source-alpha** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 72 | 72 |
| **gecko-t/win11-64-25h2-ssd** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 2 | 2 |
| **gecko-t/win11-64-25h2-ssd-gpu** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 2 | 2 |
| **gecko-t/win11-64-25h2-unprivileged** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 2 | 2 |
| **gecko-t/win11-64-25h2-unprivileged-alpha** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 3 | 3 |
| **gecko-t/win11-64-25h2-webgpu** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 849 | 849 |
| **gecko-t/win11-64-25h2-webgpu-alpha** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 103 | 103 |
| **gecko-t/win11-a64-24h2** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | arm64 | 1.25.7 | 243 | 243 |
| **gecko-t/win11-a64-24h2-alpha** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | arm64 | 1.25.7 | 14 | 14 |
| **gecko-t/win11-a64-25h2-alpha** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | arm64 | 1.25.7 | 6 | 6 |
| **glean-1/b-linux** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 18 | 18 |
| **glean-1/decision** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 249 | 249 |
| **glean-1/images** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 2 | 2 |
| **glean-3/b-linux** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 22 | 22 |
| **glean-3/decision** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 8 | 8 |
| **glean-3/images** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 2 | 2 |
| **mobile-1/b-linux** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 145 | 145 |
| **mobile-1/decision** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 282 | 282 |
| **mobile-1/images** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 2 | 2 |
| **mobile-3/b-linux** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 263 | 263 |
| **mobile-3/decision** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 147 | 147 |
| **mobile-3/images** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 2 | 2 |
| **mozilla-1/b-linux** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 2 | 2 |
| **mozilla-1/b-win2022** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 2 | 2 |
| **mozilla-1/decision** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 281 | 281 |
| **mozilla-1/images** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 24 | 24 |
| **mozilla-3/b-linux** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 2 | 2 |
| **mozilla-3/b-win2022** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 3 | 3 |
| **mozilla-3/decision** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 440 | 440 |
| **mozilla-3/images** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 2 | 2 |
| **mozilla-t/t-linux-2204-wayland** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | amd64 | 1.22.2 | 621 | 621 |
| **mozilla-t/t-linux-2204-wayland-relsre** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | amd64 | 1.22.2 | 2 | 2 |
| **mozilla-t/t-linux-2404-wayland** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 2 | 2 |
| **mozilla-t/t-linux-arm64-docker** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | arm64 | 1.26.0 | 2 | 2 |
| **mozilla-t/t-linux-docker** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 22 | 22 |
| **mozilla-t/t-linux-docker-amd** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 2 | 2 |
| **mozilla-t/t-linux-docker-noscratch** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 23 | 23 |
| **mozilla-t/t-linux-docker-noscratch-amd** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 2 | 2 |
| **mozillavpn-1/b-linux** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 264 | 264 |
| **mozillavpn-1/b-linux-gcp-gw** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | amd64 | 1.22.2 | 2 | 2 |
| **mozillavpn-1/b-linux-large** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 205 | 205 |
| **mozillavpn-1/b-win2022** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 46 | 46 |
| **mozillavpn-1/b-win2022-alpha** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 3 | 3 |
| **mozillavpn-1/decision** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 275 | 275 |
| **mozillavpn-1/images** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 35 | 35 |
| **mozillavpn-1/win11-a64-24h2-builder** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | arm64 | 1.25.7 | 3 | 3 |
| **mozillavpn-3/b-linux** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 150 | 150 |
| **mozillavpn-3/b-linux-gcp-gw** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | amd64 | 1.22.2 | 3 | 3 |
| **mozillavpn-3/b-linux-large** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 120 | 120 |
| **mozillavpn-3/b-win2022** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 51 | 51 |
| **mozillavpn-3/decision** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 22 | 22 |
| **mozillavpn-3/images** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 16 | 16 |
| **mozillavpn-3/win11-a64-24h2-builder** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | arm64 | 1.25.7 | 2 | 2 |
| **nss-1/b-win2022** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 207 | 207 |
| **nss-1/b-win2022-alpha** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 3 | 3 |
| **nss-1/decision** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 33 | 33 |
| **nss-1/images** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 2 | 2 |
| **nss-1/linux-docker** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 551 | 551 |
| **nss-1/win11-a64-24h2** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | arm64 | 1.25.7 | 3 | 3 |
| **nss-1/win11-a64-24h2-alpha** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | arm64 | 1.25.7 | 12 | 12 |
| **nss-1/win11-a64-24h2-builder** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | arm64 | 1.25.7 | 2 | 2 |
| **nss-1/win11-a64-24h2-builder-alpha** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | arm64 | 1.25.7 | 3 | 3 |
| **nss-1/win11-a64-25h2-alpha** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | arm64 | 1.25.7 | 5 | 5 |
| **nss-3/b-win2022** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 276 | 276 |
| **nss-3/decision** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 15 | 15 |
| **nss-3/images** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 2 | 2 |
| **nss-3/linux-docker** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 692 | 692 |
| **nss-3/win11-a64-24h2-builder** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | arm64 | 1.25.7 | 2 | 2 |
| **nss-t/t-linux-arm64-docker** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | arm64 | 1.26.0 | 2 | 2 |
| **nss-t/t-linux-docker** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 1346 | 1346 |
| **nss-t/t-linux-docker-amd** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 2 | 2 |
| **nss-t/win11-a64-24h2** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | arm64 | 1.25.7 | 2 | 2 |
| **nss-t/win11-a64-24h2-alpha** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | arm64 | 1.25.7 | 9 | 9 |
| **nss-t/win11-a64-25h2-alpha** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | arm64 | 1.25.7 | 14 | 14 |
| **proj-autophone/gecko-t-bitbar-gw-perf-a55** | generic-worker | 36.0.0 | simple | b9cb11293f | linux | amd64 | 1.13.7 | 0 | 0 |
| **proj-autophone/gecko-t-bitbar-gw-perf-p6** | generic-worker | 36.0.0 | simple | b9cb11293f | linux | amd64 | 1.13.7 | 0 | 0 |
| **proj-autophone/gecko-t-bitbar-gw-perf-s24** | generic-worker | 36.0.0 | simple | b9cb11293f | linux | amd64 | 1.13.7 | 0 | 0 |
| **proj-autophone/gecko-t-bitbar-gw-unit-p5** | generic-worker | 36.0.0 | simple | b9cb11293f | linux | amd64 | 1.13.7 | 0 | 0 |
| **proj-autophone/gecko-t-lambda-alpha-a55** | generic-worker | 87.0.0 | insecure | 99a1fcafbb | linux | amd64 | 1.24.5 | 0 | 0 |
| **proj-autophone/gecko-t-lambda-perf-a55** | generic-worker | 87.0.0 | insecure | 99a1fcafbb | linux | amd64 | 1.24.5 | 0 | 0 |
| **releng-1/decision** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 109 | 109 |
| **releng-1/linux-docker** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 107 | 107 |
| **releng-3/decision** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 58 | 58 |
| **releng-3/linux-docker** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 96 | 96 |
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
| **releng-t/linux-docker** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 549 | 549 |
| **relops-1/decision** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 2 | 2 |
| **relops-1/images** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 2 | 2 |
| **relops-3/decision** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 2 | 2 |
| **relops-3/images** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 2 | 2 |
| **scriptworker-1/decision** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 52 | 52 |
| **scriptworker-1/images** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 294 | 294 |
| **scriptworker-3/decision** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 27 | 27 |
| **scriptworker-3/images** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 199 | 199 |
| **taskgraph-1/decision** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 23 | 23 |
| **taskgraph-1/images** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 17 | 17 |
| **taskgraph-3/decision** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 12 | 12 |
| **taskgraph-3/images** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 12 | 12 |
| **taskgraph-t/linux-docker** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 191 | 191 |
| **taskgraph-t/win11-64-24h2** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 13 | 13 |
| **taskgraph-t/win11-64-25h2** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 2 | 2 |
| **translations-1/b-linux-large-gcp-1tb-32-256-d2g** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 22 | 22 |
| **translations-1/b-linux-large-gcp-1tb-32-256-std-d2g** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 8 | 8 |
| **translations-1/b-linux-large-gcp-1tb-64-512-d2g** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 2 | 2 |
| **translations-1/b-linux-large-gcp-1tb-64-512-std-d2g** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 2 | 2 |
| **translations-1/b-linux-large-gcp-d2g** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 209 | 209 |
| **translations-1/b-linux-large-gcp-d2g-1tb** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 2 | 2 |
| **translations-1/b-linux-large-gcp-d2g-1tb-standard** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 2 | 2 |
| **translations-1/b-linux-large-gcp-d2g-300gb** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 148 | 148 |
| **translations-1/b-linux-v100-gpu-d2g** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 175 | 175 |
| **translations-1/b-linux-v100-gpu-d2g-4** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 56 | 56 |
| **translations-1/b-linux-v100-gpu-d2g-4-1tb** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 48 | 48 |
| **translations-1/b-linux-v100-gpu-d2g-4-1tb-alpha** | generic-worker | 84.1.2 | multiuser | 4a139f395a | linux | amd64 | 1.24.4 | 2 | 2 |
| **translations-1/b-linux-v100-gpu-d2g-4-1tb-standard** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 2 | 2 |
| **translations-1/b-linux-v100-gpu-d2g-4-1tb-std-alpha** | generic-worker | 84.1.2 | multiuser | 4a139f395a | linux | amd64 | 1.24.4 | 2 | 2 |
| **translations-1/b-linux-v100-gpu-d2g-4-2tb** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 14 | 14 |
| **translations-1/b-linux-v100-gpu-d2g-4-2tb-alpha** | generic-worker | 84.1.2 | multiuser | 4a139f395a | linux | amd64 | 1.24.4 | 2 | 2 |
| **translations-1/b-linux-v100-gpu-d2g-4-300gb** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 18 | 18 |
| **translations-1/b-linux-v100-gpu-d2g-4-300gb-alpha** | generic-worker | 84.1.2 | multiuser | 4a139f395a | linux | amd64 | 1.24.4 | 2 | 2 |
| **translations-1/b-linux-v100-gpu-d2g-4-300gb-standard** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 2 | 2 |
| **translations-1/b-linux-v100-gpu-d2g-4-300gb-std-alpha** | generic-worker | 84.1.2 | multiuser | 4a139f395a | linux | amd64 | 1.24.4 | 2 | 2 |
| **translations-1/b-linux-v100-gpu-d2g-4-alpha** | generic-worker | 84.1.2 | multiuser | 4a139f395a | linux | amd64 | 1.24.4 | 2 | 2 |
| **translations-1/b-linux-v100-gpu-d2g-alpha** | generic-worker | 84.1.2 | multiuser | 4a139f395a | linux | amd64 | 1.24.4 | 2 | 2 |
| **translations-1/decision** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 527 | 527 |
| **translations-1/images** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 32 | 32 |
| **translations-t/t-linux-gw-noscratch-amd** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 2 | 2 |
| **xpi-1/b-linux** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 225 | 225 |
| **xpi-1/decision** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 138 | 138 |
| **xpi-1/images** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 2 | 2 |
| **xpi-3/b-linux** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 2 | 2 |
| **xpi-3/decision** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 2 | 2 |
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
| projects/fxci-production-level3-workers/global/images/docker-firefoxci-gcp-l3-googlecompute-2024-02-05t23-18-22z | 52 |
| projects/taskcluster-imaging/global/images/docker-firefoxci-gcp-l1-googlecompute-2025-06-13t18-31-38z | 85 |
| projects/taskcluster-imaging/global/images/docker-worker-gcp-u14-04-2025-06-16 | 19 |
| projects/taskcluster-imaging/global/images/docker-firefoxci-gcp-lt-googlecompute-2023-04-13t21-30-28z | 1 |


| Worker Pool | Implementation | Version | Total Workers | Total Capacity |
| --- | --- | --- | ---: | ---: |
| **app-services-1/b-linux-amd** | docker-worker | 38.0.5 | 2 | 2 |
| **app-services-1/b-linux-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **app-services-1/decision-gcp** | docker-worker | 38.0.5 | 794 | 1588 |
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
| **comm-1/decision-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **comm-1/images-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **comm-2/b-linux-amd** | docker-worker | 38.0.5 | 2 | 2 |
| **comm-2/b-linux-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **comm-2/b-linux-large-amd** | docker-worker | 38.0.5 | 2 | 2 |
| **comm-2/b-linux-large-gcp** | docker-worker | 38.0.5 | 3 | 3 |
| **comm-2/b-linux-xlarge-amd** | docker-worker | 38.0.5 | 2 | 2 |
| **comm-2/b-linux-xlarge-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **comm-2/decision-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **comm-2/images-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **comm-3/b-linux-amd** | docker-worker | 38.0.5 | 2 | 2 |
| **comm-3/b-linux-gcp** | docker-worker | 38.0.5 | 38 | 38 |
| **comm-3/b-linux-large-amd** | docker-worker | 38.0.5 | 2 | 2 |
| **comm-3/b-linux-large-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **comm-3/b-linux-xlarge-amd** | docker-worker | 38.0.5 | 2 | 2 |
| **comm-3/b-linux-xlarge-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **comm-3/decision-gcp** | docker-worker | 38.0.5 | 4 | 4 |
| **comm-3/images-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **comm-t/misc-gcp** | docker-worker | 38.0.5 | 2 | 16 |
| **comm-t/t-linux-large-gcp** | docker-worker | unknown version | 2 | 2 |
| **comm-t/t-linux-large-noscratch-gcp** | docker-worker | unknown version | 2 | 2 |
| **comm-t/t-linux-xlarge-gcp** | docker-worker | unknown version | 2 | 2 |
| **comm-t/t-linux-xlarge-noscratch-gcp** | docker-worker | unknown version | 2 | 2 |
| **comm-t/t-linux-xlarge-source-gcp** | docker-worker | unknown version | 2 | 2 |
| **comm-t/t-linux-xlarge-source-noscratch-gcp** | docker-worker | unknown version | 2 | 2 |
| **gecko-1/b-linux-amd** | docker-worker | 38.0.5 | 2 | 2 |
| **gecko-1/b-linux-gcp** | docker-worker | 38.0.5 | 3 | 3 |
| **gecko-1/b-linux-gcp-relops1411** | docker-worker | 38.0.5 | 2 | 2 |
| **gecko-1/b-linux-gcp-test-bug-1882320** | docker-worker | 38.0.5 | 2 | 2 |
| **gecko-1/b-linux-kvm-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **gecko-1/b-linux-large-amd** | docker-worker | 38.0.5 | 2 | 2 |
| **gecko-1/b-linux-large-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **gecko-1/b-linux-medium-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **gecko-1/b-linux-xlarge-amd** | docker-worker | 38.0.5 | 2 | 2 |
| **gecko-1/b-linux-xlarge-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **gecko-1/b-linux-xlarge-gcp-bug1797804-c2** | docker-worker | 38.0.5 | 2 | 2 |
| **gecko-1/decision-gcp** | docker-worker | 38.0.5 | 4 | 4 |
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
| **gecko-2/decision-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **gecko-2/images-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **gecko-3/b-linux-amd** | docker-worker | 38.0.5 | 2 | 2 |
| **gecko-3/b-linux-gcp** | docker-worker | 38.0.5 | 1724 | 1724 |
| **gecko-3/b-linux-kvm-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **gecko-3/b-linux-large-amd** | docker-worker | 38.0.5 | 2 | 2 |
| **gecko-3/b-linux-large-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **gecko-3/b-linux-medium-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **gecko-3/b-linux-xlarge-amd** | docker-worker | 38.0.5 | 2 | 2 |
| **gecko-3/b-linux-xlarge-gcp** | docker-worker | 38.0.5 | 3 | 3 |
| **gecko-3/decision-gcp** | docker-worker | 38.0.5 | 836 | 836 |
| **gecko-3/images-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **gecko-t/misc-gcp** | docker-worker | 38.0.5 | 29 | 232 |
| **gecko-t/t-linux-kvm-gcp** | docker-worker | unknown version | 2 | 2 |
| **gecko-t/t-linux-kvm-gcp-relops1411** | docker-worker | unknown version | 2 | 2 |
| **gecko-t/t-linux-kvm-noscratch-gcp** | docker-worker | unknown version | 15 | 15 |
| **gecko-t/t-linux-large-gcp** | docker-worker | unknown version | 2 | 2 |
| **gecko-t/t-linux-large-hostub2204-gcp** | docker-worker | 48.3.0 | 2 | 2 |
| **gecko-t/t-linux-large-noscratch-gcp** | docker-worker | unknown version | 4161 | 4161 |
| **gecko-t/t-linux-xlarge-gcp** | docker-worker | unknown version | 59 | 59 |
| **gecko-t/t-linux-xlarge-noscratch-gcp** | docker-worker | unknown version | 2679 | 2679 |
| **gecko-t/t-linux-xlarge-source-gcp** | docker-worker | unknown version | 154 | 154 |
| **gecko-t/t-linux-xlarge-source-noscratch-gcp** | docker-worker | unknown version | 2 | 2 |
| **glean-1/b-linux-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **glean-1/decision-gcp** | docker-worker | 38.0.5 | 777 | 1554 |
| **glean-1/images-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **glean-3/b-linux-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **glean-3/decision-gcp** | docker-worker | 38.0.5 | 2 | 4 |
| **glean-3/images-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **infra/build-decision** | docker-worker | 38.0.5 | 3 | 24 |
| **mobile-1/b-linux-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **mobile-1/decision-gcp** | docker-worker | 38.0.5 | 784 | 1568 |
| **mobile-1/images-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **mobile-3/b-linux-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **mobile-3/decision-gcp** | docker-worker | 38.0.5 | 2 | 4 |
| **mobile-3/images-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **mozilla-1/b-linux-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **mozilla-1/decision-gcp** | docker-worker | 38.0.5 | 779 | 1558 |
| **mozilla-1/images-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **mozilla-3/b-linux-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **mozilla-3/decision-gcp** | docker-worker | 38.0.5 | 2 | 4 |
| **mozilla-3/images-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **mozilla-t/t-linux-large-gcp** | docker-worker | unknown version | 2 | 2 |
| **mozilla-t/t-linux-large-noscratch-gcp** | docker-worker | unknown version | 2 | 2 |
| **mozillavpn-1/b-linux-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **mozillavpn-1/b-linux-large-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **mozillavpn-1/decision-gcp** | docker-worker | 38.0.5 | 783 | 1566 |
| **mozillavpn-1/images-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **mozillavpn-3/b-linux-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **mozillavpn-3/b-linux-large-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **mozillavpn-3/decision-gcp** | docker-worker | 38.0.5 | 2 | 4 |
| **mozillavpn-3/images-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **nss-1/decision-gcp** | docker-worker | 38.0.5 | 2 | 4 |
| **nss-1/images-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **nss-1/linux-gcp** | docker-worker | 38.0.5 | 11 | 11 |
| **nss-3/decision-gcp** | docker-worker | 38.0.5 | 2 | 4 |
| **nss-3/images-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **nss-3/linux-gcp** | docker-worker | 38.0.5 | 18 | 18 |
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
| **translations-1/decision-gcp** | docker-worker | 38.0.5 | 1724 | 3448 |
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
| **scriptworker-k8s/adhoc-t-signing** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/app-services-3-beetmover** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/app-services-3-signing** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/comm-1-beetmover** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
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
| **scriptworker-k8s/gecko-1-balrog-dev** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
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
|  | 2 |
| projects/taskcluster-imaging/global/images/generic-2204-wayland-vm-gcp-googlecompute-2023-09-22t17-39-37z | 2 |


| Worker Pool | Implementation | Version | Total Workers | Total Capacity |
| --- | --- | --- | ---: | ---: |
| **built-in/fail** |  | No artifacts found | 0 | 0 |
| **built-in/succeed** |  | No artifacts found | 0 | 0 |
| **gecko-t/t-linux-vm-2204-wayland** |  | No artifacts found | 79 | 79 |
| **gecko-t/t-linux-vm-2204-wayland-snap** |  | No artifacts found | 2 | 2 |


## Version not determined [^2]

Total: `18`


Count by image:

| Version | Count |
| :--- | ---: |
| /subscriptions/108d46d5-fe9b-4850-9a7d-8c914aa6c1f0/resourceGroups/rg-packer-worker-images/providers/Microsoft.Compute/galleries/win11_a64_25h2_builder_alpha/images/win11_a64_25h2_builder_alpha/versions/1.0.0 | 2 |
| /subscriptions/108d46d5-fe9b-4850-9a7d-8c914aa6c1f0/resourceGroups/rg-packer-worker-images/providers/Microsoft.Compute/galleries/win11_a64_25h2_tester/images/win11_a64_25h2_tester/versions/1.0.0 | 4 |
|  | 3 |
| projects/taskcluster-imaging/global/images/gw-fxci-gcp-l1-2404-amd64-googlecompute-alpha | 1 |
| /subscriptions/108d46d5-fe9b-4850-9a7d-8c914aa6c1f0/resourceGroups/rg-packer-worker-images/providers/Microsoft.Compute/galleries/win11_a64_25h2_builder/images/win11_a64_25h2_builder/versions/1.0.0 | 4 |
| /subscriptions/a30e97ab-734a-4f3b-a0e4-c51c0bff0701/resourceGroups/rg-packer-worker-images/providers/Microsoft.Compute/galleries/trusted_win11_a64_25h2_builder/images/trusted_win11_a64_25h2_builder/versions/1.0.0 | 4 |


| Worker Pool | Implementation | Version | Total Workers | Total Capacity |
| --- | --- | --- | ---: | ---: |
| **comm-t/win11-a64-25h2** |  | Version not determined; task not (yet) claimed | 58 | 58 |
| **enterprise-1/win11-a64-25h2-builder** |  | Version not determined; task not (yet) claimed | 58 | 58 |
| **enterprise-3/win11-a64-25h2-builder** |  | Version not determined; task not (yet) claimed | 58 | 58 |
| **gecko-1/win11-a64-25h2-builder** |  | Version not determined; task not (yet) claimed | 58 | 58 |
| **gecko-1/win11-a64-25h2-builder-alpha** |  | Version not determined; task not (yet) claimed | 58 | 58 |
| **gecko-3/win11-a64-25h2-builder** |  | Version not determined; task not (yet) claimed | 58 | 58 |
| **gecko-t/t-linux-2404-relsre** |  | Version not determined; task not (yet) claimed | 4 | 4 |
| **gecko-t/win11-a64-25h2** |  | Version not determined; task not (yet) claimed | 767 | 767 |
| **mozillavpn-1/win11-a64-25h2-builder** |  | Version not determined; task not (yet) claimed | 58 | 58 |
| **mozillavpn-3/win11-a64-25h2-builder** |  | Version not determined; task not (yet) claimed | 56 | 56 |
| **nss-1/win11-a64-25h2** |  | Version not determined; task not (yet) claimed | 62 | 62 |
| **nss-1/win11-a64-25h2-builder** |  | Version not determined; task not (yet) claimed | 52 | 52 |
| **nss-1/win11-a64-25h2-builder-alpha** |  | Version not determined; task not (yet) claimed | 54 | 54 |
| **nss-3/win11-a64-25h2-builder** |  | Version not determined; task not (yet) claimed | 53 | 53 |
| **nss-t/win11-a64-25h2** |  | Version not determined; task not (yet) claimed | 60 | 60 |
| **releng-hardware/gecko-t-linux-netperf-1804** |  | Version not determined; task not (yet) claimed | 0 | 0 |
| **releng-hardware/gecko-t-osx-1400-r8** |  | Version not determined; task not (yet) claimed | 0 | 0 |
| **releng-hardware/nss-3-b-osx-1015** |  | Version not determined; task not (yet) claimed | 0 | 0 |



[^1]: Those are the pools whose tasks were claimed and resolved by a worker as expected, but the worker did not publish either artifact `public/logs/live_backing.log` nor `public/logs/chain_of_trust.log`, which is the source used to identify the worker implementation.

[^2]: Probing task remains pending after two hours. Those are the pools that were not able to start any worker to claim the task within two hours.

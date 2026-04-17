

# Worker Pool Versions


## Generic Worker

Total: `399`

Count by version:

| Version | Count |
| :--- | ---: |
| 36.0.0 | 4 |
| 45.0.0 | 1 |
| 60.3.4 | 17 |
| 61.0.0 | 1 |
| 64.3.0 | 26 |
| 65.1.0 | 1 |
| 84.1.2 | 7 |
| 87.0.0 | 1 |
| 88.0.2 | 1 |
| 91.0.2 | 2 |
| 96.2.2 | 7 |
| 96.2.3 | 117 |
| 97.0.1 | 214 |


Count by image:

| Version | Count |
| :--- | ---: |
| projects/fxci-production-level3-workers/global/images/gw-fxci-gcp-l3-2404-amd64-headless-googlecompute-2026-03-05 | 59 |
| projects/taskcluster-imaging/global/images/gw-fxci-gcp-l1-gui-googlecompute-2024-08-22t22-48-09z | 8 |
| projects/fxci-production-level3-workers/global/images/gw-fxci-gcp-l3-arm64-gui-googlecompute-2024-09-18t19-02-58z | 4 |
| projects/taskcluster-imaging/global/images/gw-fxci-gcp-l1-2404-amd64-headless-googlecompute-2026-03-05 | 122 |
| projects/fxci-production-level3-workers/global/images/gw-fxci-gcp-l3-2404-arm64-headless-googlecompute-2026-03-05 | 6 |
| projects/fxci-production-level3-workers/global/images/gw-fxci-gcp-l3-gui-googlecompute-2024-09-18t05-46-31z | 3 |
| projects/taskcluster-imaging/global/images/gw-fxci-gcp-l1-2404-amd64-gui-googlecompute-2026-03-05 | 4 |
| projects/taskcluster-imaging/global/images/gw-fxci-gcp-l1-gui-googlecompute-2025-01-13t22-33-40z | 2 |
| unknown | 117 |
| projects/fxci-production-level3-workers/global/images/gw-fxci-gcp-l3-2404-amd64-headless-googlecompute-alpha | 1 |
| projects/taskcluster-imaging/global/images/gw-fxci-gcp-l1-2404-amd64-gui-googlecompute-alpha | 1 |
| projects/taskcluster-imaging/global/images/gw-fxci-gcp-l1-2404-amd64-headless-googlecompute-alpha-tc | 7 |
| projects/taskcluster-imaging/global/images/gw-fxci-gcp-l1-2404-arm64-headless-googlecompute-alpha | 1 |
|  | 35 |
| projects/taskcluster-imaging/global/images/gw-fxci-gcp-l1-arm64-gui-googlecompute-2024-09-18t14-57-52z | 9 |
| projects/taskcluster-imaging/global/images/gw-fxci-gcp-l1-2404-amd64-headless-googlecompute-alpha | 6 |
| projects/taskcluster-imaging/global/images/gw-fxci-gcp-l1-2404-arm64-headless-googlecompute-2026-03-05 | 14 |


| Worker Pool | Implementation | Version | Engine | Revision | OS | Arch | GO | Total Workers | Total Capacity |
| --- | --- | --- | --- | --- | --- | --- | --- | ---: | ---: |
| **adhoc-1/decision** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 4 | 4 |
| **adhoc-1/images** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 21 | 21 |
| **adhoc-3/decision** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 3 | 3 |
| **adhoc-3/images** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 3 | 3 |
| **app-services-1/b-linux** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 305 | 305 |
| **app-services-1/b-linux-docker-amd** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 2 | 2 |
| **app-services-1/decision** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 1 | 1 |
| **app-services-1/images** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 4 | 4 |
| **app-services-3/b-linux** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 709 | 709 |
| **app-services-3/b-linux-docker-amd** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 2 | 2 |
| **app-services-3/decision** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 22 | 22 |
| **app-services-3/images** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 3 | 3 |
| **code-analysis-1/decision** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 2 | 2 |
| **code-analysis-1/linux-docker** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 2 | 2 |
| **code-analysis-1/linux-gw-gcp** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | amd64 | 1.22.2 | 23 | 23 |
| **code-analysis-3/decision** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 2 | 2 |
| **code-analysis-3/linux-docker** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 2 | 2 |
| **code-analysis-3/linux-gw-gcp** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | amd64 | 1.22.2 | 4 | 4 |
| **code-coverage/bot** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 234 | 234 |
| **code-review/bot** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 10776 | 10776 |
| **comm-1/b-linux** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 617 | 617 |
| **comm-1/b-linux-aarch64** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | arm64 | 1.26.0 | 3 | 3 |
| **comm-1/b-linux-docker-amd** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 3 | 3 |
| **comm-1/b-linux-docker-large-amd** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 4 | 4 |
| **comm-1/b-linux-docker-xlarge-amd** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 2 | 2 |
| **comm-1/b-linux-gcp-aarch64** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | arm64 | 1.22.2 | 3 | 3 |
| **comm-1/b-linux-large** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 2 | 2 |
| **comm-1/b-linux-xlarge** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 5 | 5 |
| **comm-1/b-win2022** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 2 | 2 |
| **comm-1/decision** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 277 | 277 |
| **comm-1/images** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 5 | 5 |
| **comm-1/images-aarch64** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | arm64 | 1.26.0 | 2 | 2 |
| **comm-1/images-gcp-aarch64** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | arm64 | 1.22.2 | 2 | 2 |
| **comm-2/b-linux** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 2 | 2 |
| **comm-2/b-linux-aarch64** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | arm64 | 1.26.0 | 6 | 6 |
| **comm-2/b-linux-docker-amd** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 2 | 2 |
| **comm-2/b-linux-docker-large-amd** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 2 | 2 |
| **comm-2/b-linux-docker-xlarge-amd** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 2 | 2 |
| **comm-2/b-linux-gcp-aarch64** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | arm64 | 1.22.2 | 3 | 3 |
| **comm-2/b-linux-large** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 2 | 2 |
| **comm-2/b-linux-xlarge** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 2 | 2 |
| **comm-2/decision** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 2 | 2 |
| **comm-2/images** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 2 | 2 |
| **comm-2/images-aarch64** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | arm64 | 1.26.0 | 2 | 2 |
| **comm-2/images-gcp-aarch64** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | arm64 | 1.22.2 | 2 | 2 |
| **comm-3/b-linux** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 3202 | 3202 |
| **comm-3/b-linux-aarch64** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | arm64 | 1.26.0 | 3 | 3 |
| **comm-3/b-linux-docker-amd** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 1824 | 1824 |
| **comm-3/b-linux-docker-large-amd** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 6 | 6 |
| **comm-3/b-linux-docker-xlarge-amd** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 2 | 2 |
| **comm-3/b-linux-gcp-aarch64** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | arm64 | 1.22.2 | 4 | 4 |
| **comm-3/b-linux-large** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 2 | 2 |
| **comm-3/b-linux-xlarge** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 3 | 3 |
| **comm-3/b-win2022** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 645 | 645 |
| **comm-3/decision** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 90 | 90 |
| **comm-3/images** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 5 | 5 |
| **comm-3/images-aarch64** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | arm64 | 1.26.0 | 2 | 2 |
| **comm-3/images-gcp-aarch64** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | arm64 | 1.22.2 | 2 | 2 |
| **comm-t/misc** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 2 | 2 |
| **comm-t/t-linux-arm64-docker** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | arm64 | 1.26.0 | 2 | 2 |
| **comm-t/t-linux-docker** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 11 | 11 |
| **comm-t/t-linux-docker-amd** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 2087 | 2087 |
| **comm-t/t-linux-docker-noscratch** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 98 | 98 |
| **comm-t/t-linux-docker-noscratch-amd** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 5295 | 5295 |
| **comm-t/win11-64-2009** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 2 | 2 |
| **comm-t/win11-64-2009-ssd** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 2 | 2 |
| **comm-t/win11-64-24h2** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 4684 | 4684 |
| **comm-t/win11-64-24h2-ssd** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 346 | 346 |
| **comm-t/win11-64-25h2** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 2 | 2 |
| **comm-t/win11-64-25h2-ssd** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 2 | 2 |
| **comm-t/win11-a64-24h2** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | arm64 | 1.25.7 | 2 | 2 |
| **comm-t/win11-a64-24h2-alpha** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | arm64 | 1.25.7 | 6 | 6 |
| **comm-t/win11-a64-25h2** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | arm64 | 1.25.7 | 2 | 2 |
| **comm-t/win11-a64-25h2-alpha** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | arm64 | 1.25.7 | 12 | 12 |
| **enterprise-1/b-linux** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 2261 | 2261 |
| **enterprise-1/b-linux-aarch64** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | arm64 | 1.26.0 | 6 | 6 |
| **enterprise-1/b-linux-docker-amd** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 778 | 778 |
| **enterprise-1/b-linux-docker-large-amd** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 31 | 31 |
| **enterprise-1/b-linux-docker-xlarge-amd** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 1045 | 1045 |
| **enterprise-1/b-win2022** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 401 | 401 |
| **enterprise-1/decision** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 682 | 682 |
| **enterprise-1/images** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 8 | 8 |
| **enterprise-1/images-aarch64** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | arm64 | 1.26.0 | 2 | 2 |
| **enterprise-1/win11-a64-24h2-builder** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | arm64 | 1.25.7 | 7 | 7 |
| **enterprise-1/win11-a64-25h2-builder** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | arm64 | 1.25.7 | 2 | 2 |
| **enterprise-3/b-linux** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 1388 | 1388 |
| **enterprise-3/b-linux-aarch64** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | arm64 | 1.26.0 | 5 | 5 |
| **enterprise-3/b-linux-docker-amd** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 305 | 305 |
| **enterprise-3/b-linux-docker-large-amd** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 72 | 72 |
| **enterprise-3/b-linux-docker-xlarge-amd** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 293 | 293 |
| **enterprise-3/b-win2022** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 292 | 292 |
| **enterprise-3/decision** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 47 | 47 |
| **enterprise-3/images** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 8 | 8 |
| **enterprise-3/images-aarch64** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | arm64 | 1.26.0 | 2 | 2 |
| **enterprise-3/win11-a64-24h2-builder** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | arm64 | 1.25.7 | 2 | 2 |
| **enterprise-t/misc** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 244 | 244 |
| **enterprise-t/t-linux-docker-16c32gb-amd** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 820 | 820 |
| **enterprise-t/t-linux-docker-amd** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 4279 | 4279 |
| **enterprise-t/t-linux-docker-noscratch-amd** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 5300 | 5300 |
| **enterprise-t/win11-64-24h2** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 5710 | 5710 |
| **enterprise-t/win11-64-24h2-alpha** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 2 | 2 |
| **enterprise-t/win11-64-24h2-source** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 2 | 2 |
| **enterprise-t/win11-64-24h2-source-alpha** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 2 | 2 |
| **enterprise-t/win11-64-25h2** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 44 | 44 |
| **enterprise-t/win11-64-25h2-alpha** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 2 | 2 |
| **enterprise-t/win11-64-25h2-source** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 91 | 91 |
| **enterprise-t/win11-64-25h2-source-alpha** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 2 | 2 |
| **gecko-1/b-linux** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 6355 | 6355 |
| **gecko-1/b-linux-2204-kvm-gcp** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | amd64 | 1.22.2 | 2 | 2 |
| **gecko-1/b-linux-aarch64** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | arm64 | 1.26.0 | 31 | 31 |
| **gecko-1/b-linux-docker-alpha** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 2 | 2 |
| **gecko-1/b-linux-docker-amd** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 10237 | 10237 |
| **gecko-1/b-linux-docker-large-amd** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 175 | 175 |
| **gecko-1/b-linux-docker-xlarge-amd** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 1080 | 1080 |
| **gecko-1/b-linux-gcp-aarch64** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | arm64 | 1.22.2 | 6 | 6 |
| **gecko-1/b-linux-kvm** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 175 | 175 |
| **gecko-1/b-linux-large** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 5 | 5 |
| **gecko-1/b-linux-medium** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 18506 | 18506 |
| **gecko-1/b-linux-xlarge** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 4 | 4 |
| **gecko-1/b-win2022** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 926 | 926 |
| **gecko-1/b-win2022-alpha** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 2 | 2 |
| **gecko-1/b-win2022-headless** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 2 | 2 |
| **gecko-1/b-win2022-xxlarge** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 3 | 3 |
| **gecko-1/decision** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 1480 | 1480 |
| **gecko-1/images** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 48 | 48 |
| **gecko-1/images-aarch64** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | arm64 | 1.26.0 | 2 | 2 |
| **gecko-1/images-gcp-aarch64** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | arm64 | 1.22.2 | 2 | 2 |
| **gecko-1/win11-a64-24h2-builder** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | arm64 | 1.25.7 | 23 | 23 |
| **gecko-1/win11-a64-24h2-builder-alpha** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | arm64 | 1.25.7 | 4 | 4 |
| **gecko-1/win11-a64-25h2-builder** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | arm64 | 1.25.7 | 2 | 2 |
| **gecko-2/b-linux** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 86 | 86 |
| **gecko-2/b-linux-aarch64** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | arm64 | 1.26.0 | 4 | 4 |
| **gecko-2/b-linux-docker-amd** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 210 | 210 |
| **gecko-2/b-linux-docker-large-amd** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 35 | 35 |
| **gecko-2/b-linux-docker-xlarge-amd** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 51 | 51 |
| **gecko-2/b-linux-gcp-aarch64** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | arm64 | 1.22.2 | 2 | 2 |
| **gecko-2/b-linux-kvm** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 14 | 14 |
| **gecko-2/b-linux-large** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 2 | 2 |
| **gecko-2/b-linux-medium** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 245 | 245 |
| **gecko-2/b-linux-xlarge** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 2 | 2 |
| **gecko-2/b-win2022** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 8 | 8 |
| **gecko-2/decision** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 17 | 17 |
| **gecko-2/images** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 2 | 2 |
| **gecko-2/images-aarch64** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | arm64 | 1.26.0 | 2 | 2 |
| **gecko-2/images-gcp-aarch64** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | arm64 | 1.22.2 | 2 | 2 |
| **gecko-3/b-linux** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 20226 | 20226 |
| **gecko-3/b-linux-2204-kvm-gcp** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | amd64 | 1.22.2 | 2 | 2 |
| **gecko-3/b-linux-aarch64** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | arm64 | 1.26.0 | 2 | 2 |
| **gecko-3/b-linux-docker-alpha** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 2 | 2 |
| **gecko-3/b-linux-docker-amd** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 30639 | 30639 |
| **gecko-3/b-linux-docker-large-amd** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 538 | 538 |
| **gecko-3/b-linux-docker-xlarge-amd** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 3931 | 3931 |
| **gecko-3/b-linux-gcp-aarch64** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | arm64 | 1.22.2 | 3 | 3 |
| **gecko-3/b-linux-kvm** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 490 | 490 |
| **gecko-3/b-linux-large** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 2 | 2 |
| **gecko-3/b-linux-medium** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 26568 | 26568 |
| **gecko-3/b-linux-xlarge** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 188 | 188 |
| **gecko-3/b-win2022** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 5930 | 5930 |
| **gecko-3/b-win2022-headless** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 2 | 2 |
| **gecko-3/b-win2022-xxlarge** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 14 | 14 |
| **gecko-3/decision** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 6307 | 6307 |
| **gecko-3/images** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 20 | 20 |
| **gecko-3/images-aarch64** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | arm64 | 1.26.0 | 2 | 2 |
| **gecko-3/images-gcp-aarch64** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | arm64 | 1.22.2 | 2 | 2 |
| **gecko-3/win11-a64-24h2-builder** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | arm64 | 1.25.7 | 87 | 87 |
| **gecko-t/misc** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 1240 | 1240 |
| **gecko-t/t-linux-2204-wayland** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | amd64 | 1.22.2 | 9278 | 9278 |
| **gecko-t/t-linux-2204-wayland-arm64-relsre** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | arm64 | 1.22.2 | 2 | 2 |
| **gecko-t/t-linux-2204-wayland-relsre** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | amd64 | 1.22.2 | 2 | 2 |
| **gecko-t/t-linux-2204-wayland-root-exp** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | amd64 | 1.22.2 | 2 | 2 |
| **gecko-t/t-linux-2204-wayland-snap** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | amd64 | 1.22.2 | 110 | 110 |
| **gecko-t/t-linux-2404-headless-alpha** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 2 | 2 |
| **gecko-t/t-linux-2404-headless-arm64-alpha** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | arm64 | 1.26.0 | 2 | 2 |
| **gecko-t/t-linux-2404-headless-ssd-alpha** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 2 | 2 |
| **gecko-t/t-linux-2404-wayland** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 2 | 2 |
| **gecko-t/t-linux-2404-wayland-relsre** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 2 | 2 |
| **gecko-t/t-linux-2404-wayland-snap** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 112 | 112 |
| **gecko-t/t-linux-arm64-docker** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | arm64 | 1.26.0 | 34 | 34 |
| **gecko-t/t-linux-docker** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 1874 | 1874 |
| **gecko-t/t-linux-docker-16c32gb-amd** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 1771 | 1771 |
| **gecko-t/t-linux-docker-amd** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 35134 | 35134 |
| **gecko-t/t-linux-docker-amd-alpha** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 2 | 2 |
| **gecko-t/t-linux-docker-kvm** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 62807 | 62807 |
| **gecko-t/t-linux-docker-kvm-alpha** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 2 | 2 |
| **gecko-t/t-linux-docker-noscratch** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 26473 | 26473 |
| **gecko-t/t-linux-docker-noscratch-amd** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 147490 | 147490 |
| **gecko-t/t-linux-docker-noscratch-amd-alpha** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 2 | 2 |
| **gecko-t/t-linux-xlarge-2204-wayland** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | amd64 | 1.22.2 | 7351 | 7351 |
| **gecko-t/t-linux-xlarge-2404-wayland** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 3 | 3 |
| **gecko-t/win10-64-2009** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 4530 | 4530 |
| **gecko-t/win10-64-2009-alpha** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 5 | 5 |
| **gecko-t/win10-64-2009-gpu** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 678 | 678 |
| **gecko-t/win10-64-2009-gpu-alpha** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 7 | 7 |
| **gecko-t/win10-64-2009-source** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 90 | 90 |
| **gecko-t/win10-64-2009-source-alpha** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 19 | 19 |
| **gecko-t/win10-64-2009-webgpu** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 2 | 2 |
| **gecko-t/win10-64-2009-webgpu-alpha** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 5 | 5 |
| **gecko-t/win11-64-2009** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 5737 | 5737 |
| **gecko-t/win11-64-2009-alpha** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 32 | 32 |
| **gecko-t/win11-64-2009-gpu** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 3641 | 3641 |
| **gecko-t/win11-64-2009-gpu-alpha** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 7 | 7 |
| **gecko-t/win11-64-2009-source** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 130 | 130 |
| **gecko-t/win11-64-2009-source-alpha** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 4 | 4 |
| **gecko-t/win11-64-2009-ssd** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 2 | 2 |
| **gecko-t/win11-64-2009-ssd-gpu** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 2 | 2 |
| **gecko-t/win11-64-2009-webgpu** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 2 | 2 |
| **gecko-t/win11-64-2009-webgpu-alpha** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 2 | 2 |
| **gecko-t/win11-64-24h2** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 27997 | 27997 |
| **gecko-t/win11-64-24h2-alpha** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 32 | 32 |
| **gecko-t/win11-64-24h2-amd** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 4 | 4 |
| **gecko-t/win11-64-24h2-gpu** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 15150 | 15150 |
| **gecko-t/win11-64-24h2-gpu-alpha** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 7 | 7 |
| **gecko-t/win11-64-24h2-large** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 1290 | 1290 |
| **gecko-t/win11-64-24h2-large-alpha** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 2 | 2 |
| **gecko-t/win11-64-24h2-privileged** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 242 | 242 |
| **gecko-t/win11-64-24h2-privileged-alpha** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 31 | 31 |
| **gecko-t/win11-64-24h2-source** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 215 | 215 |
| **gecko-t/win11-64-24h2-source-alpha** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 4 | 4 |
| **gecko-t/win11-64-24h2-ssd** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 2 | 2 |
| **gecko-t/win11-64-24h2-ssd-gpu** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 2 | 2 |
| **gecko-t/win11-64-24h2-unprivileged** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 2 | 2 |
| **gecko-t/win11-64-24h2-unprivileged-alpha** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 6 | 6 |
| **gecko-t/win11-64-24h2-webgpu** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 2 | 2 |
| **gecko-t/win11-64-24h2-webgpu-alpha** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 3 | 3 |
| **gecko-t/win11-64-25h2** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 58858 | 58858 |
| **gecko-t/win11-64-25h2-alpha** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 32 | 32 |
| **gecko-t/win11-64-25h2-amd** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 4 | 4 |
| **gecko-t/win11-64-25h2-gpu** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 8074 | 8074 |
| **gecko-t/win11-64-25h2-gpu-alpha** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 8 | 8 |
| **gecko-t/win11-64-25h2-large-alpha** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 42 | 42 |
| **gecko-t/win11-64-25h2-privileged** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 16 | 16 |
| **gecko-t/win11-64-25h2-privileged-alpha** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 7 | 7 |
| **gecko-t/win11-64-25h2-source** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 1577 | 1577 |
| **gecko-t/win11-64-25h2-source-alpha** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 4 | 4 |
| **gecko-t/win11-64-25h2-ssd** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 2 | 2 |
| **gecko-t/win11-64-25h2-ssd-gpu** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 2 | 2 |
| **gecko-t/win11-64-25h2-unprivileged** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 2 | 2 |
| **gecko-t/win11-64-25h2-unprivileged-alpha** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 10 | 10 |
| **gecko-t/win11-64-25h2-webgpu** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 2909 | 2909 |
| **gecko-t/win11-64-25h2-webgpu-alpha** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 2 | 2 |
| **gecko-t/win11-a64-24h2** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | arm64 | 1.25.7 | 148 | 148 |
| **gecko-t/win11-a64-24h2-alpha** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | arm64 | 1.25.7 | 11 | 11 |
| **gecko-t/win11-a64-25h2** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | arm64 | 1.25.7 | 562 | 562 |
| **gecko-t/win11-a64-25h2-alpha** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | arm64 | 1.25.7 | 16 | 16 |
| **glean-1/b-linux** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 35 | 35 |
| **glean-1/decision** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 276 | 276 |
| **glean-1/images** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 2 | 2 |
| **glean-3/b-linux** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 26 | 26 |
| **glean-3/decision** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 8 | 8 |
| **glean-3/images** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 2 | 2 |
| **infra/build-decision-alpha** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 2 | 2 |
| **mobile-1/b-linux** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 123 | 123 |
| **mobile-1/decision** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 291 | 291 |
| **mobile-1/images** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 2 | 2 |
| **mobile-3/b-linux** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 250 | 250 |
| **mobile-3/decision** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 147 | 147 |
| **mobile-3/images** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 2 | 2 |
| **mozilla-1/b-linux** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 2 | 2 |
| **mozilla-1/b-win2022** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 2 | 2 |
| **mozilla-1/decision** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 279 | 279 |
| **mozilla-1/images** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 7 | 7 |
| **mozilla-3/b-linux** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 2 | 2 |
| **mozilla-3/b-win2022** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 2 | 2 |
| **mozilla-3/decision** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 438 | 438 |
| **mozilla-3/images** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 2 | 2 |
| **mozilla-t/t-linux-2204-wayland** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | amd64 | 1.22.2 | 591 | 591 |
| **mozilla-t/t-linux-2204-wayland-relsre** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | amd64 | 1.22.2 | 2 | 2 |
| **mozilla-t/t-linux-2404-wayland** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 2 | 2 |
| **mozilla-t/t-linux-arm64-docker** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | arm64 | 1.26.0 | 2 | 2 |
| **mozilla-t/t-linux-docker** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 16 | 16 |
| **mozilla-t/t-linux-docker-amd** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 2 | 2 |
| **mozilla-t/t-linux-docker-noscratch** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 16 | 16 |
| **mozilla-t/t-linux-docker-noscratch-amd** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 2 | 2 |
| **mozillavpn-1/b-linux** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 435 | 435 |
| **mozillavpn-1/b-linux-gcp-gw** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | amd64 | 1.22.2 | 3 | 3 |
| **mozillavpn-1/b-linux-large** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 333 | 333 |
| **mozillavpn-1/b-win2022** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 101 | 101 |
| **mozillavpn-1/b-win2022-alpha** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 3 | 3 |
| **mozillavpn-1/decision** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 291 | 291 |
| **mozillavpn-1/images** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 22 | 22 |
| **mozillavpn-1/win11-a64-24h2-builder** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | arm64 | 1.25.7 | 7 | 7 |
| **mozillavpn-1/win11-a64-25h2-builder** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | arm64 | 1.25.7 | 3 | 3 |
| **mozillavpn-3/b-linux** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 147 | 147 |
| **mozillavpn-3/b-linux-gcp-gw** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | amd64 | 1.22.2 | 2 | 2 |
| **mozillavpn-3/b-linux-large** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 150 | 150 |
| **mozillavpn-3/b-win2022** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 49 | 49 |
| **mozillavpn-3/decision** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 21 | 21 |
| **mozillavpn-3/images** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 8 | 8 |
| **mozillavpn-3/win11-a64-24h2-builder** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | arm64 | 1.25.7 | 3 | 3 |
| **nss-1/b-win2022** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 232 | 232 |
| **nss-1/b-win2022-alpha** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 3 | 3 |
| **nss-1/decision** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 18 | 18 |
| **nss-1/images** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 2 | 2 |
| **nss-1/linux-docker** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 587 | 587 |
| **nss-1/win11-a64-24h2** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | arm64 | 1.25.7 | 3 | 3 |
| **nss-1/win11-a64-24h2-alpha** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | arm64 | 1.25.7 | 11 | 11 |
| **nss-1/win11-a64-24h2-builder** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | arm64 | 1.25.7 | 3 | 3 |
| **nss-1/win11-a64-24h2-builder-alpha** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | arm64 | 1.25.7 | 3 | 3 |
| **nss-1/win11-a64-25h2** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | arm64 | 1.25.7 | 5 | 5 |
| **nss-1/win11-a64-25h2-alpha** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | arm64 | 1.25.7 | 13 | 13 |
| **nss-1/win11-a64-25h2-builder** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | arm64 | 1.25.7 | 3 | 3 |
| **nss-3/b-win2022** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 373 | 373 |
| **nss-3/decision** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 33 | 33 |
| **nss-3/images** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 2 | 2 |
| **nss-3/linux-docker** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 1916 | 1916 |
| **nss-3/win11-a64-24h2-builder** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | arm64 | 1.25.7 | 3 | 3 |
| **nss-t/t-linux-arm64-docker** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | arm64 | 1.26.0 | 2 | 2 |
| **nss-t/t-linux-docker** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 2759 | 2759 |
| **nss-t/t-linux-docker-amd** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 2 | 2 |
| **nss-t/win11-a64-24h2** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | arm64 | 1.25.7 | 5 | 5 |
| **nss-t/win11-a64-24h2-alpha** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | arm64 | 1.25.7 | 14 | 14 |
| **nss-t/win11-a64-25h2** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | arm64 | 1.25.7 | 4 | 4 |
| **nss-t/win11-a64-25h2-alpha** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | arm64 | 1.25.7 | 13 | 13 |
| **proj-autophone/gecko-t-bitbar-gw-perf-a55** | generic-worker | 36.0.0 | simple | b9cb11293f | linux | amd64 | 1.13.7 | 0 | 0 |
| **proj-autophone/gecko-t-bitbar-gw-perf-p6** | generic-worker | 36.0.0 | simple | b9cb11293f | linux | amd64 | 1.13.7 | 0 | 0 |
| **proj-autophone/gecko-t-bitbar-gw-perf-s24** | generic-worker | 36.0.0 | simple | b9cb11293f | linux | amd64 | 1.13.7 | 0 | 0 |
| **proj-autophone/gecko-t-bitbar-gw-unit-p5** | generic-worker | 36.0.0 | simple | b9cb11293f | linux | amd64 | 1.13.7 | 0 | 0 |
| **proj-autophone/gecko-t-lambda-perf-a55** | generic-worker | 87.0.0 | insecure | 99a1fcafbb | linux | amd64 | 1.24.5 | 0 | 0 |
| **releng-1/decision** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 63 | 63 |
| **releng-1/linux-docker** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 48 | 48 |
| **releng-3/decision** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 35 | 35 |
| **releng-3/linux-docker** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 51 | 51 |
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
| **releng-hardware/gecko-t-osx-1015-r8** | generic-worker | 60.3.4 | simple | 943a6f2b0d | darwin | amd64 | 1.22.0 | 0 | 0 |
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
| **releng-hardware/win11-64-24h2-hw** | generic-worker | 96.2.2 | multiuser | 773509947e | windows | amd64 | 1.25.7 | 0 | 0 |
| **releng-hardware/win11-64-24h2-hw-alpha** | generic-worker | 96.2.2 | multiuser | 773509947e | windows | amd64 | 1.25.7 | 0 | 0 |
| **releng-hardware/win11-64-24h2-hw-perf-debug** | generic-worker | 96.2.2 | multiuser | 773509947e | windows | amd64 | 1.25.7 | 0 | 0 |
| **releng-hardware/win11-64-24h2-hw-perf-sheriff** | generic-worker | 96.2.2 | multiuser | 773509947e | windows | amd64 | 1.25.7 | 0 | 0 |
| **releng-hardware/win11-64-24h2-hw-ref** | generic-worker | 96.2.2 | multiuser | 773509947e | windows | amd64 | 1.25.7 | 0 | 0 |
| **releng-hardware/win11-64-24h2-hw-ref-alpha** | generic-worker | 96.2.2 | multiuser | 773509947e | windows | amd64 | 1.25.7 | 0 | 0 |
| **releng-hardware/win11-64-24h2-hw-relops1213** | generic-worker | 96.2.2 | multiuser | 773509947e | windows | amd64 | 1.25.7 | 0 | 0 |
| **releng-t/linux-docker** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 252 | 252 |
| **relops-1/decision** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 2 | 2 |
| **relops-1/images** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 2 | 2 |
| **relops-3/decision** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 5 | 5 |
| **relops-3/images** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 2 | 2 |
| **scriptworker-1/decision** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 18 | 18 |
| **scriptworker-1/images** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 102 | 102 |
| **scriptworker-3/decision** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 13 | 13 |
| **scriptworker-3/images** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 83 | 83 |
| **taskgraph-1/decision** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 27 | 27 |
| **taskgraph-1/images** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 20 | 20 |
| **taskgraph-3/decision** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 14 | 14 |
| **taskgraph-3/images** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 16 | 16 |
| **taskgraph-t/linux-docker** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 202 | 202 |
| **taskgraph-t/win11-64-24h2** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 25 | 25 |
| **taskgraph-t/win11-64-25h2** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 2 | 2 |
| **translations-1/b-linux-large-gcp-1tb-32-256-d2g** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 2 | 2 |
| **translations-1/b-linux-large-gcp-1tb-32-256-std-d2g** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 4 | 4 |
| **translations-1/b-linux-large-gcp-1tb-64-512-d2g** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 2 | 2 |
| **translations-1/b-linux-large-gcp-1tb-64-512-std-d2g** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 2 | 2 |
| **translations-1/b-linux-large-gcp-d2g** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 25 | 25 |
| **translations-1/b-linux-large-gcp-d2g-1tb** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 2 | 2 |
| **translations-1/b-linux-large-gcp-d2g-1tb-standard** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 2 | 2 |
| **translations-1/b-linux-large-gcp-d2g-300gb** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 32 | 32 |
| **translations-1/b-linux-v100-gpu-d2g** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 2 | 2 |
| **translations-1/b-linux-v100-gpu-d2g-4** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 22 | 22 |
| **translations-1/b-linux-v100-gpu-d2g-4-1tb** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 3 | 3 |
| **translations-1/b-linux-v100-gpu-d2g-4-1tb-alpha** | generic-worker | 84.1.2 | multiuser | 4a139f395a | linux | amd64 | 1.24.4 | 2 | 2 |
| **translations-1/b-linux-v100-gpu-d2g-4-1tb-standard** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 3 | 3 |
| **translations-1/b-linux-v100-gpu-d2g-4-1tb-std-alpha** | generic-worker | 84.1.2 | multiuser | 4a139f395a | linux | amd64 | 1.24.4 | 2 | 2 |
| **translations-1/b-linux-v100-gpu-d2g-4-2tb** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 3 | 3 |
| **translations-1/b-linux-v100-gpu-d2g-4-2tb-alpha** | generic-worker | 84.1.2 | multiuser | 4a139f395a | linux | amd64 | 1.24.4 | 2 | 2 |
| **translations-1/b-linux-v100-gpu-d2g-4-300gb** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 14 | 14 |
| **translations-1/b-linux-v100-gpu-d2g-4-300gb-alpha** | generic-worker | 84.1.2 | multiuser | 4a139f395a | linux | amd64 | 1.24.4 | 2 | 2 |
| **translations-1/b-linux-v100-gpu-d2g-4-300gb-standard** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 2 | 2 |
| **translations-1/b-linux-v100-gpu-d2g-4-300gb-std-alpha** | generic-worker | 84.1.2 | multiuser | 4a139f395a | linux | amd64 | 1.24.4 | 2 | 2 |
| **translations-1/b-linux-v100-gpu-d2g-4-alpha** | generic-worker | 84.1.2 | multiuser | 4a139f395a | linux | amd64 | 1.24.4 | 2 | 2 |
| **translations-1/b-linux-v100-gpu-d2g-alpha** | generic-worker | 84.1.2 | multiuser | 4a139f395a | linux | amd64 | 1.24.4 | 2 | 2 |
| **translations-1/decision** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 517 | 517 |
| **translations-1/images** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 2 | 2 |
| **translations-t/t-linux-gw-noscratch-amd** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 2 | 2 |
| **xpi-1/b-linux** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 112 | 112 |
| **xpi-1/decision** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 55 | 55 |
| **xpi-1/images** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 2 | 2 |
| **xpi-3/b-linux** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 3 | 3 |
| **xpi-3/decision** | generic-worker | 97.0.1 | multiuser | 6519c127aa | linux | amd64 | 1.26.0 | 6 | 6 |
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
| projects/taskcluster-imaging/global/images/docker-worker-gcp-u14-04-2025-06-16 | 19 |
| projects/taskcluster-imaging/global/images/docker-firefoxci-gcp-lt-googlecompute-2023-04-13t21-30-28z | 1 |
| projects/taskcluster-imaging/global/images/docker-firefoxci-gcp-l1-googlecompute-2025-06-13t18-31-38z | 85 |
| projects/fxci-production-level3-workers/global/images/docker-firefoxci-gcp-l3-googlecompute-2024-02-05t23-18-22z | 52 |


| Worker Pool | Implementation | Version | Total Workers | Total Capacity |
| --- | --- | --- | ---: | ---: |
| **app-services-1/b-linux-amd** | docker-worker | 38.0.5 | 2 | 2 |
| **app-services-1/b-linux-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **app-services-1/decision-gcp** | docker-worker | 38.0.5 | 745 | 1490 |
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
| **comm-2/b-linux-large-gcp** | docker-worker | 38.0.5 | 3 | 3 |
| **comm-2/b-linux-xlarge-amd** | docker-worker | 38.0.5 | 2 | 2 |
| **comm-2/b-linux-xlarge-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **comm-2/decision-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **comm-2/images-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **comm-3/b-linux-amd** | docker-worker | 38.0.5 | 2 | 2 |
| **comm-3/b-linux-gcp** | docker-worker | 38.0.5 | 86 | 86 |
| **comm-3/b-linux-large-amd** | docker-worker | 38.0.5 | 2 | 2 |
| **comm-3/b-linux-large-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **comm-3/b-linux-xlarge-amd** | docker-worker | 38.0.5 | 2 | 2 |
| **comm-3/b-linux-xlarge-gcp** | docker-worker | 38.0.5 | 3 | 3 |
| **comm-3/decision-gcp** | docker-worker | 38.0.5 | 5 | 5 |
| **comm-3/images-gcp** | docker-worker | 38.0.5 | 2 | 2 |
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
| **gecko-1/b-linux-large-gcp** | docker-worker | 38.0.5 | 3 | 3 |
| **gecko-1/b-linux-medium-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **gecko-1/b-linux-xlarge-amd** | docker-worker | 38.0.5 | 2 | 2 |
| **gecko-1/b-linux-xlarge-gcp** | docker-worker | 38.0.5 | 3 | 3 |
| **gecko-1/b-linux-xlarge-gcp-bug1797804-c2** | docker-worker | 38.0.5 | 2 | 2 |
| **gecko-1/decision-gcp** | docker-worker | 38.0.5 | 7 | 7 |
| **gecko-1/decision-gcp-c3d4-amd-bug1990935** | docker-worker | 38.0.5 | 2 | 2 |
| **gecko-1/decision-gcp-c3d8-amd-bug1990935** | docker-worker | 38.0.5 | 2 | 2 |
| **gecko-1/decision-gcp-c4d4-amd-bug1990935** | docker-worker | 38.0.5 | 2 | 2 |
| **gecko-1/decision-gcp-c4d4-hcpu-amd-bug1990935** | docker-worker | 38.0.5 | 2 | 2 |
| **gecko-1/decision-gcp-c4d8-amd-bug1990935** | docker-worker | 38.0.5 | 2 | 2 |
| **gecko-1/decision-gcp-c4d8-hcpu-amd-bug1990935** | docker-worker | 38.0.5 | 2 | 2 |
| **gecko-1/decision-gcp-c4d8-lssd-amd-bug1990935** | docker-worker | 38.0.5 | 2 | 2 |
| **gecko-1/images-gcp** | docker-worker | 38.0.5 | 3 | 3 |
| **gecko-2/b-linux-amd** | docker-worker | 38.0.5 | 2 | 2 |
| **gecko-2/b-linux-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **gecko-2/b-linux-kvm-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **gecko-2/b-linux-large-amd** | docker-worker | 38.0.5 | 2 | 2 |
| **gecko-2/b-linux-large-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **gecko-2/b-linux-medium-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **gecko-2/b-linux-xlarge-amd** | docker-worker | 38.0.5 | 2 | 2 |
| **gecko-2/b-linux-xlarge-gcp** | docker-worker | 38.0.5 | 7 | 7 |
| **gecko-2/decision-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **gecko-2/images-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **gecko-3/b-linux-amd** | docker-worker | 38.0.5 | 2 | 2 |
| **gecko-3/b-linux-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **gecko-3/b-linux-kvm-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **gecko-3/b-linux-large-amd** | docker-worker | 38.0.5 | 2 | 2 |
| **gecko-3/b-linux-large-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **gecko-3/b-linux-medium-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **gecko-3/b-linux-xlarge-amd** | docker-worker | 38.0.5 | 2 | 2 |
| **gecko-3/b-linux-xlarge-gcp** | docker-worker | 38.0.5 | 3 | 3 |
| **gecko-3/decision-gcp** | docker-worker | 38.0.5 | 711 | 711 |
| **gecko-3/images-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **gecko-t/misc-gcp** | docker-worker | 38.0.5 | 2 | 16 |
| **gecko-t/t-linux-kvm-gcp** | docker-worker | unknown version | 2 | 2 |
| **gecko-t/t-linux-kvm-gcp-relops1411** | docker-worker | unknown version | 2 | 2 |
| **gecko-t/t-linux-kvm-noscratch-gcp** | docker-worker | unknown version | 2 | 2 |
| **gecko-t/t-linux-large-gcp** | docker-worker | unknown version | 2 | 2 |
| **gecko-t/t-linux-large-hostub2204-gcp** | docker-worker | 48.3.0 | 2 | 2 |
| **gecko-t/t-linux-large-noscratch-gcp** | docker-worker | unknown version | 3 | 3 |
| **gecko-t/t-linux-xlarge-gcp** | docker-worker | unknown version | 2 | 2 |
| **gecko-t/t-linux-xlarge-noscratch-gcp** | docker-worker | unknown version | 2 | 2 |
| **gecko-t/t-linux-xlarge-source-gcp** | docker-worker | unknown version | 2 | 2 |
| **gecko-t/t-linux-xlarge-source-noscratch-gcp** | docker-worker | unknown version | 2 | 2 |
| **glean-1/b-linux-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **glean-1/decision-gcp** | docker-worker | 38.0.5 | 728 | 1456 |
| **glean-1/images-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **glean-3/b-linux-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **glean-3/decision-gcp** | docker-worker | 38.0.5 | 2 | 4 |
| **glean-3/images-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **infra/build-decision** | docker-worker | 38.0.5 | 12 | 96 |
| **mobile-1/b-linux-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **mobile-1/decision-gcp** | docker-worker | 38.0.5 | 726 | 1452 |
| **mobile-1/images-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **mobile-3/b-linux-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **mobile-3/decision-gcp** | docker-worker | 38.0.5 | 2 | 4 |
| **mobile-3/images-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **mozilla-1/b-linux-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **mozilla-1/decision-gcp** | docker-worker | 38.0.5 | 729 | 1458 |
| **mozilla-1/images-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **mozilla-3/b-linux-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **mozilla-3/decision-gcp** | docker-worker | 38.0.5 | 2 | 4 |
| **mozilla-3/images-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **mozilla-t/t-linux-large-gcp** | docker-worker | unknown version | 2 | 2 |
| **mozilla-t/t-linux-large-noscratch-gcp** | docker-worker | unknown version | 3 | 3 |
| **mozillavpn-1/b-linux-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **mozillavpn-1/b-linux-large-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **mozillavpn-1/decision-gcp** | docker-worker | 38.0.5 | 731 | 1462 |
| **mozillavpn-1/images-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **mozillavpn-3/b-linux-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **mozillavpn-3/b-linux-large-gcp** | docker-worker | 38.0.5 | 3 | 3 |
| **mozillavpn-3/decision-gcp** | docker-worker | 38.0.5 | 2 | 4 |
| **mozillavpn-3/images-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **nss-1/decision-gcp** | docker-worker | 38.0.5 | 2 | 4 |
| **nss-1/images-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **nss-1/linux-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **nss-3/decision-gcp** | docker-worker | 38.0.5 | 3 | 6 |
| **nss-3/images-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **nss-3/linux-gcp** | docker-worker | 38.0.5 | 110 | 110 |
| **nss-t/t-linux-xlarge-gcp** | docker-worker | unknown version | 108 | 108 |
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
| **scriptworker-3/b-linux-gcp** | docker-worker | 38.0.5 | 42 | 42 |
| **scriptworker-3/decision-gcp** | docker-worker | 38.0.5 | 5 | 10 |
| **scriptworker-3/images-gcp** | docker-worker | 38.0.5 | 4 | 4 |
| **taskgraph-1/decision-gcp** | docker-worker | 38.0.5 | 2 | 4 |
| **taskgraph-1/images-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **taskgraph-3/decision-gcp** | docker-worker | 38.0.5 | 2 | 4 |
| **taskgraph-3/images-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **translations-1/decision-gcp** | docker-worker | 38.0.5 | 1545 | 3090 |
| **translations-1/images-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **xpi-1/b-linux-gcp** | docker-worker | 38.0.5 | 3 | 3 |
| **xpi-1/decision-gcp** | docker-worker | 38.0.5 | 3 | 6 |
| **xpi-1/images-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **xpi-3/b-linux-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **xpi-3/decision-gcp** | docker-worker | 38.0.5 | 2 | 4 |
| **xpi-3/images-gcp** | docker-worker | 38.0.5 | 2 | 2 |


## Script Worker

Total: `42`



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
| **scriptworker-k8s/mozillavpn-1-beetmover** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/mozillavpn-3-beetmover** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/mozillavpn-3-signing** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/translations-1-beetmover** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/xpi-1-beetmover** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/xpi-1-lando-dev** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/xpi-1-shipit** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
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
| **gecko-t/t-linux-vm-2204-wayland** |  | No artifacts found | 242 | 242 |
| **gecko-t/t-linux-vm-2204-wayland-snap** |  | No artifacts found | 2 | 2 |


## Version not determined [^2]

Total: `12`


Count by image:

| Version | Count |
| :--- | ---: |
| projects/taskcluster-imaging/global/images/gw-fxci-gcp-l1-2404-amd64-googlecompute-alpha | 1 |
|  | 4 |
| unknown | 7 |


| Worker Pool | Implementation | Version | Total Workers | Total Capacity |
| --- | --- | --- | ---: | ---: |
| **enterprise-3/win11-a64-25h2-builder** |  | Version not determined; task not (yet) claimed | 44 | 44 |
| **gecko-1/win11-a64-25h2-builder-alpha** |  | Version not determined; task not (yet) claimed | 153 | 153 |
| **gecko-3/win11-a64-25h2-builder** |  | Version not determined; task not (yet) claimed | 152 | 152 |
| **gecko-t/t-linux-2404-relsre** |  | Version not determined; task not (yet) claimed | 4 | 4 |
| **gecko-t/win11-64-25h2-large** |  | Version not determined; task not (yet) claimed | 3304 | 3304 |
| **mozillavpn-3/win11-a64-25h2-builder** |  | Version not determined; task not (yet) claimed | 39 | 39 |
| **nss-1/win11-a64-25h2-builder-alpha** |  | Version not determined; task not (yet) claimed | 84 | 84 |
| **nss-3/win11-a64-25h2-builder** |  | Version not determined; task not (yet) claimed | 83 | 83 |
| **proj-autophone/gecko-t-lambda-test-1** |  | Version not determined; task not (yet) claimed | 0 | 0 |
| **releng-hardware/gecko-t-linux-talos-2404** |  | Version not determined; task not (yet) claimed | 0 | 0 |
| **releng-hardware/gecko-t-osx-1400-r8** |  | Version not determined; task not (yet) claimed | 0 | 0 |
| **releng-hardware/nss-3-b-osx-1015** |  | Version not determined; task not (yet) claimed | 0 | 0 |



[^1]: Those are the pools whose tasks were claimed and resolved by a worker as expected, but the worker did not publish either artifact `public/logs/live_backing.log` nor `public/logs/chain_of_trust.log`, which is the source used to identify the worker implementation.

[^2]: Probing task remains pending after two hours. Those are the pools that were not able to start any worker to claim the task within two hours.



# Worker Pool Versions


## Generic Worker

Total: `435`

Count by version:

| Version | Count |
| :--- | ---: |
| 100.5.0 | 2 |
| 108.0.0 | 1 |
| 108.1.0 | 239 |
| 109.0.0 | 1 |
| 110.0.0 | 128 |
| 36.0.0 | 3 |
| 45.0.0 | 1 |
| 60.3.4 | 14 |
| 61.0.0 | 1 |
| 64.3.0 | 20 |
| 65.1.0 | 1 |
| 67.1.0 | 1 |
| 84.1.2 | 7 |
| 88.0.2 | 4 |
| 91.0.2 | 3 |
| 96.2.3 | 9 |


Count by image:

| Version | Count |
| :--- | ---: |
| unknown | 140 |
| projects/taskcluster-imaging/global/images/gw-fxci-gcp-l1-2404-amd64-gui-googlecompute-2026-09-09 | 5 |
| projects/fxci-production-level3-workers/global/images/gw-fxci-gcp-l3-arm64-gui-googlecompute-2024-09-18t19-02-58z | 2 |
| projects/taskcluster-imaging/global/images/generic-worker-ubuntu-24-04-9ed5a812d3d844ac9e4e | 4 |
| projects/fxci-production-level3-workers/global/images/gw-fxci-gcp-l3-2404-amd64-headless-googlecompute-alpha | 1 |
| projects/taskcluster-imaging/global/images/gw-fxci-gcp-l1-2404-amd64-headless-googlecompute-2026-09-09 | 128 |
| projects/taskcluster-imaging/global/images/gw-fxci-gcp-l1-2404-amd64-gui-googlecompute-alpha | 4 |
| projects/taskcluster-imaging/global/images/gw-fxci-gcp-l1-2404-amd64-headless-googlecompute-alpha | 6 |
| projects/taskcluster-imaging/global/images/gw-fxci-gcp-l1-gui-googlecompute-2024-08-22t22-48-09z | 9 |
| projects/taskcluster-imaging/global/images/gw-fxci-gcp-l1-gui-googlecompute-2025-01-13t22-33-40z | 2 |
| projects/fxci-production-level3-workers/global/images/gw-fxci-gcp-l3-2404-amd64-headless-googlecompute-2026-09-09 | 60 |
| projects/taskcluster-imaging/global/images/gw-fxci-gcp-l1-2404-arm64-headless-googlecompute-2026-09-09 | 16 |
| projects/taskcluster-imaging/global/images/gw-fxci-gcp-l1-2404-amd64-headless-googlecompute-alpha-tc | 7 |
|  | 35 |
| projects/fxci-production-level3-workers/global/images/gw-fxci-gcp-l3-2404-arm64-headless-googlecompute-2026-09-09 | 8 |
| projects/taskcluster-imaging/global/images/gw-fxci-gcp-l1-2404-arm64-headless-googlecompute-alpha | 1 |
| projects/taskcluster-imaging/global/images/gw-fxci-gcp-l1-arm64-gui-googlecompute-2024-09-18t14-57-52z | 5 |
| projects/fxci-production-level3-workers/global/images/gw-fxci-gcp-l3-gui-googlecompute-2024-09-18t05-46-31z | 2 |


| Worker Pool | Implementation | Version | Engine | Revision | OS | Arch | GO | Total Workers | Total Capacity |
| --- | --- | --- | --- | --- | --- | --- | --- | ---: | ---: |
| **adhoc-1/decision** | generic-worker | 108.1.0 | multiuser | f86624a8bf | linux | amd64 | 1.27.1 | 5 | 5 |
| **adhoc-1/images** | generic-worker | 108.1.0 | multiuser | f86624a8bf | linux | amd64 | 1.27.1 | 10 | 10 |
| **adhoc-3/decision** | generic-worker | 108.1.0 | multiuser | f86624a8bf | linux | amd64 | 1.27.1 | 5 | 5 |
| **adhoc-3/images** | generic-worker | 108.1.0 | multiuser | f86624a8bf | linux | amd64 | 1.27.1 | 4 | 4 |
| **app-services-1/b-linux** | generic-worker | 108.1.0 | multiuser | f86624a8bf | linux | amd64 | 1.27.1 | 440 | 440 |
| **app-services-1/b-linux-docker-amd** | generic-worker | 108.1.0 | multiuser | f86624a8bf | linux | amd64 | 1.27.1 | 2 | 2 |
| **app-services-1/decision** | generic-worker | 108.1.0 | multiuser | f86624a8bf | linux | amd64 | 1.27.1 | 5 | 5 |
| **app-services-1/images** | generic-worker | 108.1.0 | multiuser | f86624a8bf | linux | amd64 | 1.27.1 | 2 | 2 |
| **app-services-3/b-linux** | generic-worker | 108.1.0 | multiuser | f86624a8bf | linux | amd64 | 1.27.1 | 821 | 821 |
| **app-services-3/b-linux-docker-amd** | generic-worker | 108.1.0 | multiuser | f86624a8bf | linux | amd64 | 1.27.1 | 2 | 2 |
| **app-services-3/decision** | generic-worker | 108.1.0 | multiuser | f86624a8bf | linux | amd64 | 1.27.1 | 28 | 28 |
| **app-services-3/images** | generic-worker | 108.1.0 | multiuser | f86624a8bf | linux | amd64 | 1.27.1 | 2 | 2 |
| **code-analysis-1/decision** | generic-worker | 108.1.0 | multiuser | f86624a8bf | linux | amd64 | 1.27.1 | 2 | 2 |
| **code-analysis-1/linux-docker** | generic-worker | 108.1.0 | multiuser | f86624a8bf | linux | amd64 | 1.27.1 | 2 | 2 |
| **code-analysis-1/linux-gw-gcp** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | amd64 | 1.22.2 | 107 | 107 |
| **code-analysis-3/decision** | generic-worker | 108.1.0 | multiuser | f86624a8bf | linux | amd64 | 1.27.1 | 2 | 2 |
| **code-analysis-3/linux-docker** | generic-worker | 108.1.0 | multiuser | f86624a8bf | linux | amd64 | 1.27.1 | 2 | 2 |
| **code-analysis-3/linux-gw-gcp** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | amd64 | 1.22.2 | 112 | 112 |
| **code-coverage/bot** | generic-worker | 108.1.0 | multiuser | f86624a8bf | linux | amd64 | 1.27.1 | 377 | 377 |
| **code-review/bot** | generic-worker | 108.1.0 | multiuser | f86624a8bf | linux | amd64 | 1.27.1 | 7017 | 7017 |
| **comm-1/b-linux** | generic-worker | 108.1.0 | multiuser | f86624a8bf | linux | amd64 | 1.27.1 | 647 | 647 |
| **comm-1/b-linux-aarch64** | generic-worker | 108.1.0 | multiuser | f86624a8bf | linux | arm64 | 1.27.1 | 4 | 4 |
| **comm-1/b-linux-docker-amd** | generic-worker | 108.1.0 | multiuser | f86624a8bf | linux | amd64 | 1.27.1 | 2012 | 2012 |
| **comm-1/b-linux-docker-large-amd** | generic-worker | 108.1.0 | multiuser | f86624a8bf | linux | amd64 | 1.27.1 | 10 | 10 |
| **comm-1/b-linux-docker-xlarge-amd** | generic-worker | 108.1.0 | multiuser | f86624a8bf | linux | amd64 | 1.27.1 | 3 | 3 |
| **comm-1/b-linux-gcp-aarch64** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | arm64 | 1.22.2 | 2 | 2 |
| **comm-1/b-linux-large** | generic-worker | 108.1.0 | multiuser | f86624a8bf | linux | amd64 | 1.27.1 | 2 | 2 |
| **comm-1/b-linux-xlarge** | generic-worker | 108.1.0 | multiuser | f86624a8bf | linux | amd64 | 1.27.1 | 3 | 3 |
| **comm-1/b-win2022** | generic-worker | 110.0.0 | multiuser | 906277e693 | windows | amd64 | 1.27.1 | 4 | 4 |
| **comm-1/b-win2025** | generic-worker | 110.0.0 | multiuser | 906277e693 | windows | amd64 | 1.27.1 | 193 | 193 |
| **comm-1/decision** | generic-worker | 108.1.0 | multiuser | f86624a8bf | linux | amd64 | 1.27.1 | 341 | 341 |
| **comm-1/images** | generic-worker | 108.1.0 | multiuser | f86624a8bf | linux | amd64 | 1.27.1 | 9 | 9 |
| **comm-1/images-aarch64** | generic-worker | 108.1.0 | multiuser | f86624a8bf | linux | arm64 | 1.27.1 | 4 | 4 |
| **comm-2/b-linux** | generic-worker | 108.1.0 | multiuser | f86624a8bf | linux | amd64 | 1.27.1 | 2 | 2 |
| **comm-2/b-linux-aarch64** | generic-worker | 108.1.0 | multiuser | f86624a8bf | linux | arm64 | 1.27.1 | 2 | 2 |
| **comm-2/b-linux-docker-amd** | generic-worker | 108.1.0 | multiuser | f86624a8bf | linux | amd64 | 1.27.1 | 3 | 3 |
| **comm-2/b-linux-docker-large-amd** | generic-worker | 108.1.0 | multiuser | f86624a8bf | linux | amd64 | 1.27.1 | 2 | 2 |
| **comm-2/b-linux-docker-xlarge-amd** | generic-worker | 108.1.0 | multiuser | f86624a8bf | linux | amd64 | 1.27.1 | 4 | 4 |
| **comm-2/b-linux-gcp-aarch64** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | arm64 | 1.22.2 | 2 | 2 |
| **comm-2/b-linux-large** | generic-worker | 108.1.0 | multiuser | f86624a8bf | linux | amd64 | 1.27.1 | 2 | 2 |
| **comm-2/b-linux-xlarge** | generic-worker | 108.1.0 | multiuser | f86624a8bf | linux | amd64 | 1.27.1 | 3 | 3 |
| **comm-2/decision** | generic-worker | 108.1.0 | multiuser | f86624a8bf | linux | amd64 | 1.27.1 | 2 | 2 |
| **comm-2/images** | generic-worker | 108.1.0 | multiuser | f86624a8bf | linux | amd64 | 1.27.1 | 2 | 2 |
| **comm-2/images-aarch64** | generic-worker | 108.1.0 | multiuser | f86624a8bf | linux | arm64 | 1.27.1 | 2 | 2 |
| **comm-3/b-linux** | generic-worker | 108.1.0 | multiuser | f86624a8bf | linux | amd64 | 1.27.1 | 5756 | 5756 |
| **comm-3/b-linux-aarch64** | generic-worker | 108.1.0 | multiuser | f86624a8bf | linux | arm64 | 1.27.1 | 8 | 8 |
| **comm-3/b-linux-docker-amd** | generic-worker | 108.1.0 | multiuser | f86624a8bf | linux | amd64 | 1.27.1 | 3606 | 3606 |
| **comm-3/b-linux-docker-large-amd** | generic-worker | 108.1.0 | multiuser | f86624a8bf | linux | amd64 | 1.27.1 | 88 | 88 |
| **comm-3/b-linux-docker-xlarge-amd** | generic-worker | 108.1.0 | multiuser | f86624a8bf | linux | amd64 | 1.27.1 | 10 | 10 |
| **comm-3/b-linux-gcp-aarch64** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | arm64 | 1.22.2 | 2 | 2 |
| **comm-3/b-linux-large** | generic-worker | 108.1.0 | multiuser | f86624a8bf | linux | amd64 | 1.27.1 | 2 | 2 |
| **comm-3/b-linux-xlarge** | generic-worker | 108.1.0 | multiuser | f86624a8bf | linux | amd64 | 1.27.1 | 2 | 2 |
| **comm-3/b-win2022** | generic-worker | 110.0.0 | multiuser | 906277e693 | windows | amd64 | 1.27.1 | 883 | 883 |
| **comm-3/b-win2025** | generic-worker | 110.0.0 | multiuser | 906277e693 | windows | amd64 | 1.27.1 | 193 | 193 |
| **comm-3/decision** | generic-worker | 108.1.0 | multiuser | f86624a8bf | linux | amd64 | 1.27.1 | 181 | 181 |
| **comm-3/images** | generic-worker | 108.1.0 | multiuser | f86624a8bf | linux | amd64 | 1.27.1 | 10 | 10 |
| **comm-3/images-aarch64** | generic-worker | 108.1.0 | multiuser | f86624a8bf | linux | arm64 | 1.27.1 | 2 | 2 |
| **comm-t/misc** | generic-worker | 108.1.0 | multiuser | f86624a8bf | linux | amd64 | 1.27.1 | 240 | 240 |
| **comm-t/t-linux-arm64-docker** | generic-worker | 108.1.0 | multiuser | f86624a8bf | linux | arm64 | 1.27.1 | 30 | 30 |
| **comm-t/t-linux-docker** | generic-worker | 108.1.0 | multiuser | f86624a8bf | linux | amd64 | 1.27.1 | 80 | 80 |
| **comm-t/t-linux-docker-amd** | generic-worker | 108.1.0 | multiuser | f86624a8bf | linux | amd64 | 1.27.1 | 3307 | 3307 |
| **comm-t/t-linux-docker-noscratch** | generic-worker | 108.1.0 | multiuser | f86624a8bf | linux | amd64 | 1.27.1 | 102 | 102 |
| **comm-t/t-linux-docker-noscratch-amd** | generic-worker | 108.1.0 | multiuser | f86624a8bf | linux | amd64 | 1.27.1 | 8166 | 8166 |
| **comm-t/win11-64-2009** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 2 | 2 |
| **comm-t/win11-64-2009-ssd** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 2 | 2 |
| **comm-t/win11-64-24h2** | generic-worker | 110.0.0 | multiuser | 906277e693 | windows | amd64 | 1.27.1 | 121 | 121 |
| **comm-t/win11-64-24h2-ssd** | generic-worker | 110.0.0 | multiuser | 906277e693 | windows | amd64 | 1.27.1 | 2 | 2 |
| **comm-t/win11-64-25h2** | generic-worker | 110.0.0 | multiuser | 906277e693 | windows | amd64 | 1.27.1 | 8138 | 8138 |
| **comm-t/win11-64-25h2-ssd** | generic-worker | 110.0.0 | multiuser | 906277e693 | windows | amd64 | 1.27.1 | 2 | 2 |
| **comm-t/win11-a64-25h2** | generic-worker | 110.0.0 | multiuser | 906277e693 | windows | arm64 | 1.27.1 | 44 | 44 |
| **comm-t/win11-a64-25h2-alpha** | generic-worker | 110.0.0 | multiuser | 906277e693 | windows | arm64 | 1.27.1 | 21 | 21 |
| **enterprise-1/b-linux** | generic-worker | 108.1.0 | multiuser | f86624a8bf | linux | amd64 | 1.27.1 | 4856 | 4856 |
| **enterprise-1/b-linux-aarch64** | generic-worker | 108.1.0 | multiuser | f86624a8bf | linux | arm64 | 1.27.1 | 17 | 17 |
| **enterprise-1/b-linux-docker-amd** | generic-worker | 108.1.0 | multiuser | f86624a8bf | linux | amd64 | 1.27.1 | 1352 | 1352 |
| **enterprise-1/b-linux-docker-large-amd** | generic-worker | 108.1.0 | multiuser | f86624a8bf | linux | amd64 | 1.27.1 | 55 | 55 |
| **enterprise-1/b-linux-docker-xlarge-amd** | generic-worker | 108.1.0 | multiuser | f86624a8bf | linux | amd64 | 1.27.1 | 2523 | 2523 |
| **enterprise-1/b-win2022** | generic-worker | 110.0.0 | multiuser | 906277e693 | windows | amd64 | 1.27.1 | 366 | 366 |
| **enterprise-1/b-win2025** | generic-worker | 110.0.0 | multiuser | 906277e693 | windows | amd64 | 1.27.1 | 193 | 193 |
| **enterprise-1/decision** | generic-worker | 108.1.0 | multiuser | f86624a8bf | linux | amd64 | 1.27.1 | 176 | 176 |
| **enterprise-1/images** | generic-worker | 108.1.0 | multiuser | f86624a8bf | linux | amd64 | 1.27.1 | 72 | 72 |
| **enterprise-1/images-aarch64** | generic-worker | 108.1.0 | multiuser | f86624a8bf | linux | arm64 | 1.27.1 | 7 | 7 |
| **enterprise-1/win11-a64-25h2-builder** | generic-worker | 110.0.0 | multiuser | 906277e693 | windows | arm64 | 1.27.1 | 2 | 2 |
| **enterprise-3/b-linux** | generic-worker | 108.1.0 | multiuser | f86624a8bf | linux | amd64 | 1.27.1 | 9372 | 9372 |
| **enterprise-3/b-linux-aarch64** | generic-worker | 108.1.0 | multiuser | f86624a8bf | linux | arm64 | 1.27.1 | 4 | 4 |
| **enterprise-3/b-linux-docker-amd** | generic-worker | 108.1.0 | multiuser | f86624a8bf | linux | amd64 | 1.27.1 | 1442 | 1442 |
| **enterprise-3/b-linux-docker-large-amd** | generic-worker | 108.1.0 | multiuser | f86624a8bf | linux | amd64 | 1.27.1 | 90 | 90 |
| **enterprise-3/b-linux-docker-xlarge-amd** | generic-worker | 108.1.0 | multiuser | f86624a8bf | linux | amd64 | 1.27.1 | 2143 | 2143 |
| **enterprise-3/b-win2022** | generic-worker | 110.0.0 | multiuser | 906277e693 | windows | amd64 | 1.27.1 | 1213 | 1213 |
| **enterprise-3/b-win2025** | generic-worker | 110.0.0 | multiuser | 906277e693 | windows | amd64 | 1.27.1 | 193 | 193 |
| **enterprise-3/decision** | generic-worker | 108.1.0 | multiuser | f86624a8bf | linux | amd64 | 1.27.1 | 102 | 102 |
| **enterprise-3/images** | generic-worker | 108.1.0 | multiuser | f86624a8bf | linux | amd64 | 1.27.1 | 24 | 24 |
| **enterprise-3/images-aarch64** | generic-worker | 108.1.0 | multiuser | f86624a8bf | linux | arm64 | 1.27.1 | 2 | 2 |
| **enterprise-3/win11-a64-25h2-builder** | generic-worker | 110.0.0 | multiuser | 906277e693 | windows | arm64 | 1.27.1 | 2 | 2 |
| **enterprise-t/misc** | generic-worker | 108.1.0 | multiuser | f86624a8bf | linux | amd64 | 1.27.1 | 547 | 547 |
| **enterprise-t/t-linux-2204-wayland** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | amd64 | 1.22.2 | 1188 | 1188 |
| **enterprise-t/t-linux-2204-wayland-snap** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | amd64 | 1.22.2 | 2 | 2 |
| **enterprise-t/t-linux-2404-wayland-snap** | generic-worker | 108.1.0 | multiuser | f86624a8bf | linux | amd64 | 1.27.1 | 2 | 2 |
| **enterprise-t/t-linux-2404-wayland-snap-alpha** | generic-worker | 108.1.0 | multiuser | f86624a8bf | linux | amd64 | 1.27.1 | 2 | 2 |
| **enterprise-t/t-linux-docker** | generic-worker | 108.1.0 | multiuser | f86624a8bf | linux | amd64 | 1.27.1 | 2 | 2 |
| **enterprise-t/t-linux-docker-16c32gb-amd** | generic-worker | 108.1.0 | multiuser | f86624a8bf | linux | amd64 | 1.27.1 | 1433 | 1433 |
| **enterprise-t/t-linux-docker-amd** | generic-worker | 108.1.0 | multiuser | f86624a8bf | linux | amd64 | 1.27.1 | 10163 | 10163 |
| **enterprise-t/t-linux-docker-kvm** | generic-worker | 108.1.0 | multiuser | f86624a8bf | linux | amd64 | 1.27.1 | 2 | 2 |
| **enterprise-t/t-linux-docker-noscratch-amd** | generic-worker | 108.1.0 | multiuser | f86624a8bf | linux | amd64 | 1.27.1 | 15685 | 15685 |
| **enterprise-t/t-linux-xlarge-2204-wayland** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | amd64 | 1.22.2 | 2 | 2 |
| **enterprise-t/win10-64-2009** | generic-worker | 110.0.0 | multiuser | 906277e693 | windows | amd64 | 1.27.1 | 810 | 810 |
| **enterprise-t/win11-64-24h2** | generic-worker | 110.0.0 | multiuser | 906277e693 | windows | amd64 | 1.27.1 | 2 | 2 |
| **enterprise-t/win11-64-24h2-alpha** | generic-worker | 110.0.0 | multiuser | 906277e693 | windows | amd64 | 1.27.1 | 3 | 3 |
| **enterprise-t/win11-64-24h2-large** | generic-worker | 110.0.0 | multiuser | 906277e693 | windows | amd64 | 1.27.1 | 2 | 2 |
| **enterprise-t/win11-64-24h2-source** | generic-worker | 110.0.0 | multiuser | 906277e693 | windows | amd64 | 1.27.1 | 2 | 2 |
| **enterprise-t/win11-64-24h2-source-alpha** | generic-worker | 110.0.0 | multiuser | 906277e693 | windows | amd64 | 1.27.1 | 3 | 3 |
| **enterprise-t/win11-64-25h2** | generic-worker | 110.0.0 | multiuser | 906277e693 | windows | amd64 | 1.27.1 | 8282 | 8282 |
| **enterprise-t/win11-64-25h2-alpha** | generic-worker | 110.0.0 | multiuser | 
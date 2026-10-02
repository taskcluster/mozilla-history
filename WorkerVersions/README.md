

# Worker Pool Versions


## Generic Worker

Total: `428`

Count by version:

| Version | Count |
| :--- | ---: |
| 100.5.0 | 2 |
| 108.1.0 | 10 |
| 109.0.0 | 1 |
| 110.0.0 | 4 |
| 110.1.0 | 347 |
| 36.0.0 | 3 |
| 45.0.0 | 1 |
| 60.3.4 | 15 |
| 61.0.0 | 1 |
| 64.3.0 | 21 |
| 67.1.0 | 1 |
| 84.1.2 | 7 |
| 87.0.0 | 1 |
| 88.0.2 | 2 |
| 91.0.2 | 3 |
| 96.2.3 | 9 |


Count by image:

| Version | Count |
| :--- | ---: |
| unknown | 133 |
| projects/taskcluster-imaging/global/images/gw-fxci-gcp-l1-2404-amd64-gui-googlecompute-2026-09-30 | 5 |
| projects/taskcluster-imaging/global/images/gw-fxci-gcp-l1-arm64-gui-googlecompute-2024-09-18t14-57-52z | 5 |
| projects/taskcluster-imaging/global/images/gw-fxci-gcp-l1-gui-googlecompute-2024-08-22t22-48-09z | 10 |
| projects/taskcluster-imaging/global/images/generic-worker-ubuntu-24-04-9ed5a812d3d844ac9e4e | 4 |
| projects/taskcluster-imaging/global/images/gw-fxci-gcp-l1-2404-amd64-gui-googlecompute-alpha | 4 |
| projects/taskcluster-imaging/global/images/gw-fxci-gcp-l1-gui-googlecompute-2025-01-13t22-33-40z | 2 |
| projects/fxci-production-level3-workers/global/images/gw-fxci-gcp-l3-arm64-gui-googlecompute-2024-09-18t19-02-58z | 2 |
|  | 34 |
| projects/taskcluster-imaging/global/images/gw-fxci-gcp-l1-2404-amd64-headless-googlecompute-2026-09-30 | 128 |
| projects/taskcluster-imaging/global/images/gw-fxci-gcp-l1-2404-amd64-headless-googlecompute-alpha-tc | 7 |
| projects/taskcluster-imaging/global/images/gw-fxci-gcp-l1-2404-arm64-headless-googlecompute-alpha | 1 |
| projects/fxci-production-level3-workers/global/images/gw-fxci-gcp-l3-gui-googlecompute-2024-09-18t05-46-31z | 2 |
| projects/taskcluster-imaging/global/images/gw-fxci-gcp-l1-2404-amd64-headless-googlecompute-alpha | 6 |
| projects/fxci-production-level3-workers/global/images/gw-fxci-gcp-l3-2404-arm64-headless-googlecompute-2026-09-30 | 8 |
| projects/fxci-production-level3-workers/global/images/gw-fxci-gcp-l3-2404-amd64-headless-googlecompute-alpha | 1 |
| projects/fxci-production-level3-workers/global/images/gw-fxci-gcp-l3-2404-amd64-headless-googlecompute-2026-09-30 | 60 |
| projects/taskcluster-imaging/global/images/gw-fxci-gcp-l1-2404-arm64-headless-googlecompute-2026-09-30 | 16 |


| Worker Pool | Implementation | Version | Engine | Revision | OS | Arch | GO | Total Workers | Total Capacity |
| --- | --- | --- | --- | --- | --- | --- | --- | ---: | ---: |
| **adhoc-1/decision** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 4 | 4 |
| **adhoc-1/images** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 4 | 4 |
| **adhoc-3/decision** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 2 | 2 |
| **adhoc-3/images** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 2 | 2 |
| **app-services-1/b-linux** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 355 | 355 |
| **app-services-1/b-linux-docker-amd** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 2 | 2 |
| **app-services-1/decision** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 7 | 7 |
| **app-services-1/images** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 2 | 2 |
| **app-services-3/b-linux** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 785 | 785 |
| **app-services-3/b-linux-docker-amd** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 3 | 3 |
| **app-services-3/decision** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 25 | 25 |
| **app-services-3/images** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 2 | 2 |
| **code-analysis-1/decision** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 2 | 2 |
| **code-analysis-1/linux-docker** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 2 | 2 |
| **code-analysis-1/linux-gw-gcp** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | amd64 | 1.22.2 | 92 | 92 |
| **code-analysis-3/decision** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 2 | 2 |
| **code-analysis-3/linux-docker** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 2 | 2 |
| **code-analysis-3/linux-gw-gcp** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | amd64 | 1.22.2 | 66 | 66 |
| **code-coverage/bot** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 264 | 264 |
| **code-review/bot** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 10104 | 10104 |
| **comm-1/b-linux** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 550 | 550 |
| **comm-1/b-linux-aarch64** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | arm64 | 1.27.1 | 2 | 2 |
| **comm-1/b-linux-docker-amd** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 2678 | 2678 |
| **comm-1/b-linux-docker-large-amd** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 7 | 7 |
| **comm-1/b-linux-docker-xlarge-amd** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 3 | 3 |
| **comm-1/b-linux-gcp-aarch64** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | arm64 | 1.22.2 | 2 | 2 |
| **comm-1/b-linux-large** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 2 | 2 |
| **comm-1/b-linux-xlarge** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 2 | 2 |
| **comm-1/b-win2022** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | windows | amd64 | 1.27.1 | 6 | 6 |
| **comm-1/b-win2025** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | windows | amd64 | 1.27.1 | 2 | 2 |
| **comm-1/decision** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 363 | 363 |
| **comm-1/images** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 4 | 4 |
| **comm-1/images-aarch64** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | arm64 | 1.27.1 | 2 | 2 |
| **comm-2/b-linux** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 2 | 2 |
| **comm-2/b-linux-aarch64** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | arm64 | 1.27.1 | 2 | 2 |
| **comm-2/b-linux-docker-amd** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 2 | 2 |
| **comm-2/b-linux-docker-large-amd** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 2 | 2 |
| **comm-2/b-linux-docker-xlarge-amd** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 2 | 2 |
| **comm-2/b-linux-gcp-aarch64** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | arm64 | 1.22.2 | 2 | 2 |
| **comm-2/b-linux-large** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 2 | 2 |
| **comm-2/b-linux-xlarge** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 2 | 2 |
| **comm-2/decision** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 2 | 2 |
| **comm-2/images** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 2 | 2 |
| **comm-2/images-aarch64** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | arm64 | 1.27.1 | 2 | 2 |
| **comm-3/b-linux** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 6696 | 6696 |
| **comm-3/b-linux-aarch64** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | arm64 | 1.27.1 | 48 | 48 |
| **comm-3/b-linux-docker-amd** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 5549 | 5549 |
| **comm-3/b-linux-docker-large-amd** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 146 | 146 |
| **comm-3/b-linux-docker-xlarge-amd** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 51 | 51 |
| **comm-3/b-linux-gcp-aarch64** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | arm64 | 1.22.2 | 3 | 3 |
| **comm-3/b-linux-large** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 3 | 3 |
| **comm-3/b-linux-xlarge** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 2 | 2 |
| **comm-3/b-win2022** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | windows | amd64 | 1.27.1 | 857 | 857 |
| **comm-3/b-win2025** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | windows | amd64 | 1.27.1 | 2 | 2 |
| **comm-3/decision** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 243 | 243 |
| **comm-3/images** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 15 | 15 |
| **comm-3/images-aarch64** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | arm64 | 1.27.1 | 2 | 2 |
| **comm-t/misc** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 323 | 323 |
| **comm-t/t-linux-arm64-docker** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | arm64 | 1.27.1 | 2 | 2 |
| **comm-t/t-linux-docker** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 38 | 38 |
| **comm-t/t-linux-docker-amd** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 3292 | 3292 |
| **comm-t/t-linux-docker-noscratch** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 116 | 116 |
| **comm-t/t-linux-docker-noscratch-amd** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 11076 | 11076 |
| **comm-t/win11-64-2009** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 2 | 2 |
| **comm-t/win11-64-2009-ssd** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 2 | 2 |
| **comm-t/win11-64-24h2** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | windows | amd64 | 1.27.1 | 218 | 218 |
| **comm-t/win11-64-24h2-ssd** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | windows | amd64 | 1.27.1 | 2 | 2 |
| **comm-t/win11-64-25h2** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | windows | amd64 | 1.27.1 | 7864 | 7864 |
| **comm-t/win11-64-25h2-ssd** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | windows | amd64 | 1.27.1 | 2 | 2 |
| **comm-t/win11-a64-25h2** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | windows | arm64 | 1.27.1 | 2 | 2 |
| **comm-t/win11-a64-25h2-alpha** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | windows | arm64 | 1.27.1 | 94 | 94 |
| **enterprise-1/b-linux** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 1259 | 1259 |
| **enterprise-1/b-linux-aarch64** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | arm64 | 1.27.1 | 2 | 2 |
| **enterprise-1/b-linux-docker-amd** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 854 | 854 |
| **enterprise-1/b-linux-docker-large-amd** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 24 | 24 |
| **enterprise-1/b-linux-docker-xlarge-amd** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 2613 | 2613 |
| **enterprise-1/b-win2022** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | windows | amd64 | 1.27.1 | 81 | 81 |
| **enterprise-1/b-win2025** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | windows | amd64 | 1.27.1 | 2 | 2 |
| **enterprise-1/decision** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 168 | 168 |
| **enterprise-1/images** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 13 | 13 |
| **enterprise-1/images-aarch64** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | arm64 | 1.27.1 | 2 | 2 |
| **enterprise-1/win11-a64-25h2-builder** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | windows | arm64 | 1.27.1 | 2 | 2 |
| **enterprise-3/b-linux** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 8938 | 8938 |
| **enterprise-3/b-linux-aarch64** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | arm64 | 1.27.1 | 8 | 8 |
| **enterprise-3/b-linux-docker-amd** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 2027 | 2027 |
| **enterprise-3/b-linux-docker-large-amd** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 150 | 150 |
| **enterprise-3/b-linux-docker-xlarge-amd** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 3423 | 3423 |
| **enterprise-3/b-win2022** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | windows | amd64 | 1.27.1 | 1193 | 1193 |
| **enterprise-3/b-win2025** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | windows | amd64 | 1.27.1 | 2 | 2 |
| **enterprise-3/decision** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 112 | 112 |
| **enterprise-3/images** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 7 | 7 |
| **enterprise-3/images-aarch64** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | arm64 | 1.27.1 | 2 | 2 |
| **enterprise-3/win11-a64-25h2-builder** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | windows | arm64 | 1.27.1 | 2 | 2 |
| **enterprise-t/misc** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 369 | 369 |
| **enterprise-t/t-linux-2204-wayland** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | amd64 | 1.22.2 | 1003 | 1003 |
| **enterprise-t/t-linux-2204-wayland-snap** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | amd64 | 1.22.2 | 2 | 2 |
| **enterprise-t/t-linux-2404-wayland-snap** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 2 | 2 |
| **enterprise-t/t-linux-2404-wayland-snap-alpha** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 2 | 2 |
| **enterprise-t/t-linux-docker** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 2 | 2 |
| **enterprise-t/t-linux-docker-16c32gb-amd** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 1321 | 1321 |
| **enterprise-t/t-linux-docker-amd** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 6835 | 6835 |
| **enterprise-t/t-linux-docker-kvm** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 2 | 2 |
| **enterprise-t/t-linux-docker-noscratch-amd** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 14793 | 14793 |
| **enterprise-t/t-linux-xlarge-2204-wayland** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | amd64 | 1.22.2 | 2 | 2 |
| **enterprise-t/win10-64-2009** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | windows | amd64 | 1.27.1 | 603 | 603 |
| **enterprise-t/win11-64-24h2** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | windows | amd64 | 1.27.1 | 2 | 2 |
| **enterprise-t/win11-64-24h2-alpha** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | windows | amd64 | 1.27.1 | 94 | 94 |
| **enterprise-t/win11-64-24h2-large** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | windows | amd64 | 1.27.1 | 2 | 2 |
| **enterprise-t/win11-64-24h2-source** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | windows | amd64 | 1.27.1 | 2 | 2 |
| **enterprise-t/win11-64-24h2-source-alpha** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | windows | amd64 | 1.27.1 | 94 | 94 |
| **enterprise-t/win11-64-25h2** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | windows | amd64 | 1.27.1 | 5077 | 5077 |
| **enterprise-t/win11-64-25h2-alpha** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | windows | amd64 | 1.27.1 | 94 | 94 |
| **enterprise-t/win11-64-25h2-gpu** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | windows | amd64 | 1.27.1 | 278 | 278 |
| **enterprise-t/win11-64-25h2-large** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | windows | amd64 | 1.27.1 | 2 | 2 |
| **enterprise-t/win11-64-25h2-privileged** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | windows | amd64 | 1.27.1 | 2 | 2 |
| **enterprise-t/win11-64-25h2-source** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | windows | amd64 | 1.27.1 | 409 | 409 |
| **enterprise-t/win11-64-25h2-source-alpha** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | windows | amd64 | 1.27.1 | 94 | 94 |
| **enterprise-t/win11-64-25h2-webgpu** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | windows | amd64 | 1.27.1 | 3 | 3 |
| **gecko-1/b-linux** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 7131 | 7131 |
| **gecko-1/b-linux-2204-kvm-gcp** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | amd64 | 1.22.2 | 2 | 2 |
| **gecko-1/b-linux-aarch64** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | arm64 | 1.27.1 | 11 | 11 |
| **gecko-1/b-linux-docker-alpha** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 2 | 2 |
| **gecko-1/b-linux-docker-amd** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 16860 | 16860 |
| **gecko-1/b-linux-docker-large-amd** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 2266 | 2266 |
| **gecko-1/b-linux-docker-updatebot-amd** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 2 | 2 |
| **gecko-1/b-linux-docker-xlarge-amd** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 2621 | 2621 |
| **gecko-1/b-linux-gcp-aarch64** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | arm64 | 1.22.2 | 2 | 2 |
| **gecko-1/b-linux-kvm** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 230 | 230 |
| **gecko-1/b-linux-large** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 2 | 2 |
| **gecko-1/b-linux-medium** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 11331 | 11331 |
| **gecko-1/b-linux-xlarge** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 15 | 15 |
| **gecko-1/b-win2022** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | windows | amd64 | 1.27.1 | 773 | 773 |
| **gecko-1/b-win2022-alpha** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | windows | amd64 | 1.27.1 | 73 | 73 |
| **gecko-1/b-win2022-updatebot** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | windows | amd64 | 1.27.1 | 2 | 2 |
| **gecko-1/b-win2022-xxlarge** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | windows | amd64 | 1.27.1 | 12 | 12 |
| **gecko-1/b-win2025** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | windows | amd64 | 1.27.1 | 31 | 31 |
| **gecko-1/b-win2025-updatebot** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | windows | amd64 | 1.27.1 | 2 | 2 |
| **gecko-1/b-win2025-xxlarge** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | windows | amd64 | 1.27.1 | 3 | 3 |
| **gecko-1/decision** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 1439 | 1439 |
| **gecko-1/images** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 221 | 221 |
| **gecko-1/images-aarch64** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | arm64 | 1.27.1 | 34 | 34 |
| **gecko-1/win11-a64-25h2-builder** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | windows | arm64 | 1.27.1 | 21 | 21 |
| **gecko-1/win11-a64-25h2-builder-alpha** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | windows | arm64 | 1.27.1 | 94 | 94 |
| **gecko-2/b-linux** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 108 | 108 |
| **gecko-2/b-linux-aarch64** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | arm64 | 1.27.1 | 3 | 3 |
| **gecko-2/b-linux-docker-amd** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 170 | 170 |
| **gecko-2/b-linux-docker-large-amd** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 66 | 66 |
| **gecko-2/b-linux-docker-updatebot-amd** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 4 | 4 |
| **gecko-2/b-linux-docker-xlarge-amd** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 34 | 34 |
| **gecko-2/b-linux-gcp-aarch64** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | arm64 | 1.22.2 | 2 | 2 |
| **gecko-2/b-linux-kvm** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 8 | 8 |
| **gecko-2/b-linux-large** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 2 | 2 |
| **gecko-2/b-linux-medium** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 140 | 140 |
| **gecko-2/b-linux-xlarge** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 3 | 3 |
| **gecko-2/b-win2022** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | windows | amd64 | 1.27.1 | 9 | 9 |
| **gecko-2/b-win2022-updatebot** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | windows | amd64 | 1.27.1 | 3 | 3 |
| **gecko-2/b-win2025** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | windows | amd64 | 1.27.1 | 2 | 2 |
| **gecko-2/b-win2025-updatebot** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | windows | amd64 | 1.27.1 | 2 | 2 |
| **gecko-2/decision** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 13 | 13 |
| **gecko-2/images** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 2 | 2 |
| **gecko-2/images-aarch64** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | arm64 | 1.27.1 | 2 | 2 |
| **gecko-3/b-linux** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 25811 | 25811 |
| **gecko-3/b-linux-2204-kvm-gcp** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | amd64 | 1.22.2 | 2 | 2 |
| **gecko-3/b-linux-aarch64** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | arm64 | 1.27.1 | 86 | 86 |
| **gecko-3/b-linux-docker-alpha** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 2 | 2 |
| **gecko-3/b-linux-docker-amd** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 46070 | 46070 |
| **gecko-3/b-linux-docker-large-amd** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 5456 | 5456 |
| **gecko-3/b-linux-docker-updatebot-amd** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 74 | 74 |
| **gecko-3/b-linux-docker-xlarge-amd** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 7740 | 7740 |
| **gecko-3/b-linux-gcp-aarch64** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | arm64 | 1.22.2 | 2 | 2 |
| **gecko-3/b-linux-kvm** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 633 | 633 |
| **gecko-3/b-linux-large** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 2 | 2 |
| **gecko-3/b-linux-medium** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 28814 | 28814 |
| **gecko-3/b-linux-xlarge** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 45 | 45 |
| **gecko-3/b-win2022** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | windows | amd64 | 1.27.1 | 5015 | 5015 |
| **gecko-3/b-win2022-updatebot** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | windows | amd64 | 1.27.1 | 2 | 2 |
| **gecko-3/b-win2022-xxlarge** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | windows | amd64 | 1.27.1 | 13 | 13 |
| **gecko-3/b-win2025** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | windows | amd64 | 1.27.1 | 2 | 2 |
| **gecko-3/b-win2025-updatebot** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | windows | amd64 | 1.27.1 | 2 | 2 |
| **gecko-3/b-win2025-xxlarge** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | windows | amd64 | 1.27.1 | 2 | 2 |
| **gecko-3/decision** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 6151 | 6151 |
| **gecko-3/images** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 36 | 36 |
| **gecko-3/images-aarch64** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | arm64 | 1.27.1 | 2 | 2 |
| **gecko-3/win11-a64-25h2-builder** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | windows | arm64 | 1.27.1 | 111 | 111 |
| **gecko-t/misc** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 1568 | 1568 |
| **gecko-t/t-linux-2204-wayland** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | amd64 | 1.22.2 | 8188 | 8188 |
| **gecko-t/t-linux-2204-wayland-arm64-relsre** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | arm64 | 1.22.2 | 2 | 2 |
| **gecko-t/t-linux-2204-wayland-relsre** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | amd64 | 1.22.2 | 2 | 2 |
| **gecko-t/t-linux-2204-wayland-root-exp** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | amd64 | 1.22.2 | 2 | 2 |
| **gecko-t/t-linux-2204-wayland-snap** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | amd64 | 1.22.2 | 108 | 108 |
| **gecko-t/t-linux-2404-headless-alpha** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 15 | 15 |
| **gecko-t/t-linux-2404-headless-arm64-alpha** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | arm64 | 1.27.1 | 3 | 3 |
| **gecko-t/t-linux-2404-headless-ssd-alpha** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 2 | 2 |
| **gecko-t/t-linux-2404-wayland** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 2 | 2 |
| **gecko-t/t-linux-2404-wayland-alpha** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 2 | 2 |
| **gecko-t/t-linux-2404-wayland-snap** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 115 | 115 |
| **gecko-t/t-linux-2404-wayland-snap-alpha** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 166 | 166 |
| **gecko-t/t-linux-arm64-docker** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | arm64 | 1.27.1 | 48 | 48 |
| **gecko-t/t-linux-docker** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 677 | 677 |
| **gecko-t/t-linux-docker-16c32gb-amd** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 5229 | 5229 |
| **gecko-t/t-linux-docker-amd** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 54451 | 54451 |
| **gecko-t/t-linux-docker-amd-alpha** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 7 | 7 |
| **gecko-t/t-linux-docker-kvm** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 71813 | 71813 |
| **gecko-t/t-linux-docker-kvm-alpha** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 59 | 59 |
| **gecko-t/t-linux-docker-noscratch** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 4668 | 4668 |
| **gecko-t/t-linux-docker-noscratch-amd** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 342700 | 342700 |
| **gecko-t/t-linux-docker-noscratch-amd-alpha** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 959 | 959 |
| **gecko-t/t-linux-xlarge-2204-wayland** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | amd64 | 1.22.2 | 6205 | 6205 |
| **gecko-t/t-linux-xlarge-2404-wayland** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 2 | 2 |
| **gecko-t/win10-64-2009** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | windows | amd64 | 1.27.1 | 5280 | 5280 |
| **gecko-t/win10-64-2009-alpha** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | windows | amd64 | 1.27.1 | 1502 | 1502 |
| **gecko-t/win10-64-2009-gpu** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | windows | amd64 | 1.27.1 | 1103 | 1103 |
| **gecko-t/win10-64-2009-gpu-alpha** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | windows | amd64 | 1.27.1 | 94 | 94 |
| **gecko-t/win10-64-2009-source** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | windows | amd64 | 1.27.1 | 62 | 62 |
| **gecko-t/win10-64-2009-source-alpha** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | windows | amd64 | 1.27.1 | 94 | 94 |
| **gecko-t/win10-64-2009-webgpu** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | windows | amd64 | 1.27.1 | 2 | 2 |
| **gecko-t/win10-64-2009-webgpu-alpha** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | windows | amd64 | 1.27.1 | 94 | 94 |
| **gecko-t/win11-64-2009** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 1910 | 1910 |
| **gecko-t/win11-64-2009-alpha** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 2 | 2 |
| **gecko-t/win11-64-2009-source** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 23 | 23 |
| **gecko-t/win11-64-2009-source-alpha** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 3 | 3 |
| **gecko-t/win11-64-2009-ssd** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 2 | 2 |
| **gecko-t/win11-64-2009-webgpu** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 2 | 2 |
| **gecko-t/win11-64-2009-webgpu-alpha** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 5 | 5 |
| **gecko-t/win11-64-24h2** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | windows | amd64 | 1.27.1 | 273 | 273 |
| **gecko-t/win11-64-24h2-alpha** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | windows | amd64 | 1.27.1 | 5766 | 5766 |
| **gecko-t/win11-64-24h2-gpu** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | windows | amd64 | 1.27.1 | 2 | 2 |
| **gecko-t/win11-64-24h2-gpu-alpha** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | windows | amd64 | 1.27.1 | 599 | 599 |
| **gecko-t/win11-64-24h2-large** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | windows | amd64 | 1.27.1 | 24 | 24 |
| **gecko-t/win11-64-24h2-large-alpha** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | windows | amd64 | 1.27.1 | 94 | 94 |
| **gecko-t/win11-64-24h2-privileged** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | windows | amd64 | 1.27.1 | 2 | 2 |
| **gecko-t/win11-64-24h2-privileged-alpha** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | windows | amd64 | 1.27.1 | 94 | 94 |
| **gecko-t/win11-64-24h2-source** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | windows | amd64 | 1.27.1 | 3 | 3 |
| **gecko-t/win11-64-24h2-source-alpha** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | windows | amd64 | 1.27.1 | 299 | 299 |
| **gecko-t/win11-64-24h2-ssd** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | windows | amd64 | 1.27.1 | 2 | 2 |
| **gecko-t/win11-64-24h2-unprivileged** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | windows | amd64 | 1.27.1 | 2 | 2 |
| **gecko-t/win11-64-24h2-unprivileged-alpha** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | windows | amd64 | 1.27.1 | 94 | 94 |
| **gecko-t/win11-64-24h2-webgpu** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | windows | amd64 | 1.27.1 | 2 | 2 |
| **gecko-t/win11-64-24h2-webgpu-alpha** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | windows | amd64 | 1.27.1 | 94 | 94 |
| **gecko-t/win11-64-25h2** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | windows | amd64 | 1.27.1 | 92602 | 92602 |
| **gecko-t/win11-64-25h2-alpha** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | windows | amd64 | 1.27.1 | 2578 | 2578 |
| **gecko-t/win11-64-25h2-amd** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | windows | amd64 | 1.27.1 | 94 | 94 |
| **gecko-t/win11-64-25h2-gpu** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | windows | amd64 | 1.27.1 | 21485 | 21485 |
| **gecko-t/win11-64-25h2-gpu-alpha** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | windows | amd64 | 1.27.1 | 649 | 649 |
| **gecko-t/win11-64-25h2-large** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | windows | amd64 | 1.27.1 | 12578 | 12578 |
| **gecko-t/win11-64-25h2-large-alpha** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | windows | amd64 | 1.27.1 | 1242 | 1242 |
| **gecko-t/win11-64-25h2-privileged** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | windows | amd64 | 1.27.1 | 245 | 245 |
| **gecko-t/win11-64-25h2-privileged-alpha** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | windows | amd64 | 1.27.1 | 94 | 94 |
| **gecko-t/win11-64-25h2-source** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | windows | amd64 | 1.27.1 | 2554 | 2554 |
| **gecko-t/win11-64-25h2-source-alpha** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | windows | amd64 | 1.27.1 | 170 | 170 |
| **gecko-t/win11-64-25h2-ssd** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | windows | amd64 | 1.27.1 | 2 | 2 |
| **gecko-t/win11-64-25h2-ssd-gpu** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | windows | amd64 | 1.27.1 | 2 | 2 |
| **gecko-t/win11-64-25h2-unprivileged** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | windows | amd64 | 1.27.1 | 2 | 2 |
| **gecko-t/win11-64-25h2-unprivileged-alpha** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | windows | amd64 | 1.27.1 | 93 | 93 |
| **gecko-t/win11-64-25h2-webgpu** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | windows | amd64 | 1.27.1 | 3785 | 3785 |
| **gecko-t/win11-64-25h2-webgpu-alpha** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | windows | amd64 | 1.27.1 | 93 | 93 |
| **gecko-t/win11-a64-25h2** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | windows | arm64 | 1.27.1 | 297 | 297 |
| **gecko-t/win11-a64-25h2-alpha** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | windows | arm64 | 1.27.1 | 280 | 280 |
| **glean-1/b-linux** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 43 | 43 |
| **glean-1/decision** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 7 | 7 |
| **glean-1/images** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 2 | 2 |
| **glean-3/b-linux** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 47 | 47 |
| **glean-3/decision** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 13 | 13 |
| **glean-3/images** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 2 | 2 |
| **infra/build-decision-alpha** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 2 | 2 |
| **mobile-1/b-linux** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 164 | 164 |
| **mobile-1/decision** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 35 | 35 |
| **mobile-1/images** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 2 | 2 |
| **mobile-3/b-linux** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 273 | 273 |
| **mobile-3/decision** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 136 | 136 |
| **mobile-3/images** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 2 | 2 |
| **mozilla-1/b-linux** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 2 | 2 |
| **mozilla-1/b-win2022** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | windows | amd64 | 1.27.1 | 2 | 2 |
| **mozilla-1/b-win2025** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | windows | amd64 | 1.27.1 | 2 | 2 |
| **mozilla-1/decision** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 15 | 15 |
| **mozilla-1/images** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 78 | 78 |
| **mozilla-3/b-linux** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 2 | 2 |
| **mozilla-3/b-win2022** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | windows | amd64 | 1.27.1 | 2 | 2 |
| **mozilla-3/b-win2025** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | windows | amd64 | 1.27.1 | 2 | 2 |
| **mozilla-3/decision** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 460 | 460 |
| **mozilla-3/images** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 7 | 7 |
| **mozilla-t/pre-commit** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 65 | 65 |
| **mozilla-t/t-linux-2204-wayland** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | amd64 | 1.22.2 | 3 | 3 |
| **mozilla-t/t-linux-2204-wayland-relsre** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | amd64 | 1.22.2 | 2 | 2 |
| **mozilla-t/t-linux-2404-wayland** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 482 | 482 |
| **mozilla-t/t-linux-2404-wayland-alpha** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 2146 | 2146 |
| **mozilla-t/t-linux-arm64-docker** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | arm64 | 1.27.1 | 2 | 2 |
| **mozilla-t/t-linux-docker** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 411 | 411 |
| **mozilla-t/t-linux-docker-amd** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 2 | 2 |
| **mozilla-t/t-linux-docker-noscratch** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 32 | 32 |
| **mozilla-t/t-linux-docker-noscratch-amd** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 2 | 2 |
| **mozillavpn-1/b-linux** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 216 | 216 |
| **mozillavpn-1/b-linux-large** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 224 | 224 |
| **mozillavpn-1/b-win2022** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | windows | amd64 | 1.27.1 | 71 | 71 |
| **mozillavpn-1/b-win2022-alpha** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | windows | amd64 | 1.27.1 | 67 | 67 |
| **mozillavpn-1/b-win2025** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | windows | amd64 | 1.27.1 | 2 | 2 |
| **mozillavpn-1/decision** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 13 | 13 |
| **mozillavpn-1/images** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 54 | 54 |
| **mozillavpn-1/win11-a64-25h2-builder** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | windows | arm64 | 1.27.1 | 9 | 9 |
| **mozillavpn-3/b-linux** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 123 | 123 |
| **mozillavpn-3/b-linux-large** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 134 | 134 |
| **mozillavpn-3/b-win2022** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | windows | amd64 | 1.27.1 | 30 | 30 |
| **mozillavpn-3/b-win2025** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | windows | amd64 | 1.27.1 | 2 | 2 |
| **mozillavpn-3/decision** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 19 | 19 |
| **mozillavpn-3/images** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 3 | 3 |
| **mozillavpn-3/win11-a64-25h2-builder** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | windows | arm64 | 1.27.1 | 2 | 2 |
| **nss-1/b-linux-aarch64** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | arm64 | 1.27.1 | 509 | 509 |
| **nss-1/b-win2022** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | windows | amd64 | 1.27.1 | 354 | 354 |
| **nss-1/b-win2022-alpha** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | windows | amd64 | 1.27.1 | 67 | 67 |
| **nss-1/b-win2025** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | windows | amd64 | 1.27.1 | 2 | 2 |
| **nss-1/decision** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 30 | 30 |
| **nss-1/images** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 2 | 2 |
| **nss-1/images-aarch64** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | arm64 | 1.27.1 | 2 | 2 |
| **nss-1/linux-docker** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 754 | 754 |
| **nss-1/win11-a64-25h2** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | windows | arm64 | 1.27.1 | 2 | 2 |
| **nss-1/win11-a64-25h2-alpha** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | windows | arm64 | 1.27.1 | 93 | 93 |
| **nss-1/win11-a64-25h2-builder** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | windows | arm64 | 1.27.1 | 8 | 8 |
| **nss-1/win11-a64-25h2-builder-alpha** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | windows | arm64 | 1.27.1 | 93 | 93 |
| **nss-3/b-linux-aarch64** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | arm64 | 1.27.1 | 496 | 496 |
| **nss-3/b-win2022** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | windows | amd64 | 1.27.1 | 285 | 285 |
| **nss-3/b-win2025** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | windows | amd64 | 1.27.1 | 2 | 2 |
| **nss-3/decision** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 19 | 19 |
| **nss-3/images** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 2 | 2 |
| **nss-3/images-aarch64** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | arm64 | 1.27.1 | 2 | 2 |
| **nss-3/linux-docker** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 1252 | 1252 |
| **nss-3/win11-a64-25h2-builder** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | windows | arm64 | 1.27.1 | 2 | 2 |
| **nss-t/t-linux-arm64-docker** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | arm64 | 1.27.1 | 1858 | 1858 |
| **nss-t/t-linux-docker** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 2310 | 2310 |
| **nss-t/t-linux-docker-amd** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 2 | 2 |
| **nss-t/win11-a64-25h2** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | windows | arm64 | 1.27.1 | 68 | 68 |
| **nss-t/win11-a64-25h2-alpha** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | windows | arm64 | 1.27.1 | 93 | 93 |
| **proj-autophone/gecko-t-bitbar-gw-perf-a55** | generic-worker | 36.0.0 | simple | b9cb11293f | linux | amd64 | 1.13.7 | 0 | 0 |
| **proj-autophone/gecko-t-bitbar-gw-perf-p6** | generic-worker | 36.0.0 | simple | b9cb11293f | linux | amd64 | 1.13.7 | 0 | 0 |
| **proj-autophone/gecko-t-bitbar-gw-perf-s24** | generic-worker | 36.0.0 | simple | b9cb11293f | linux | amd64 | 1.13.7 | 0 | 0 |
| **proj-autophone/gecko-t-lambda-perf-a55** | generic-worker | 87.0.0 | insecure | 99a1fcafbb | linux | amd64 | 1.24.5 | 0 | 0 |
| **proj-fuzzing/bugmon-processor-windows** | generic-worker | 108.1.0 | multiuser | f86624a8bf | windows | amd64 | 1.27.1 | 3 | 3 |
| **proj-fuzzing/ci-windows** | generic-worker | 108.1.0 | multiuser | f86624a8bf | windows | amd64 | 1.27.1 | 2 | 2 |
| **proj-fuzzing/grizzly-reduce-worker-windows** | generic-worker | 108.1.0 | multiuser | f86624a8bf | windows | amd64 | 1.27.1 | 2 | 2 |
| **proj-fuzzing/grizzly-reduce-worker-windows-ngpu** | generic-worker | 108.1.0 | multiuser | f86624a8bf | windows | amd64 | 1.27.1 | 2 | 2 |
| **proj-taskcluster/ci** | generic-worker | 108.1.0 | multiuser | f86624a8bf | linux | amd64 | 1.27.1 | 86 | 86 |
| **proj-taskcluster/gw-ci-macos** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | darwin | arm64 | 1.27.1 | 2 | 2 |
| **proj-taskcluster/gw-ubuntu-24-04** | generic-worker | 108.1.0 | multiuser | f86624a8bf | linux | amd64 | 1.27.1 | 2842 | 2842 |
| **proj-taskcluster/gw-ubuntu-24-04-gui** | generic-worker | 108.1.0 | multiuser | f86624a8bf | linux | amd64 | 1.27.1 | 147 | 147 |
| **proj-taskcluster/gw-windows-2022** | generic-worker | 108.1.0 | multiuser | f86624a8bf | windows | amd64 | 1.27.1 | 158 | 158 |
| **proj-taskcluster/gw-windows-2022-gui** | generic-worker | 108.1.0 | multiuser | f86624a8bf | windows | amd64 | 1.27.1 | 156 | 156 |
| **proj-taskcluster/release** | generic-worker | 108.1.0 | multiuser | f86624a8bf | linux | amd64 | 1.27.1 | 2 | 2 |
| **releng-1/decision** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 91 | 91 |
| **releng-1/linux-docker** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 69 | 69 |
| **releng-3/decision** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 38 | 38 |
| **releng-3/linux-docker** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 64 | 64 |
| **releng-hardware/applicationservices-b-1-osx1015** | generic-worker | 60.3.4 | multiuser | 943a6f2b0d | darwin | amd64 | 1.22.0 | 0 | 0 |
| **releng-hardware/applicationservices-b-3-osx1015** | generic-worker | 60.3.4 | multiuser | 943a6f2b0d | darwin | amd64 | 1.22.0 | 0 | 0 |
| **releng-hardware/enterprise-1-b-osx-arm64** | generic-worker | 60.3.4 | multiuser | 943a6f2b0d | darwin | arm64 | 1.22.0 | 0 | 0 |
| **releng-hardware/enterprise-3-b-osx-arm64** | generic-worker | 60.3.4 | multiuser | 943a6f2b0d | darwin | arm64 | 1.22.0 | 0 | 0 |
| **releng-hardware/gecko-1-b-osx-1015** | generic-worker | 60.3.4 | multiuser | 943a6f2b0d | darwin | amd64 | 1.22.0 | 0 | 0 |
| **releng-hardware/gecko-1-b-osx-1500-staging** | generic-worker | 60.3.4 | multiuser | 943a6f2b0d | darwin | amd64 | 1.22.0 | 0 | 0 |
| **releng-hardware/gecko-1-b-osx-arm64** | generic-worker | 60.3.4 | multiuser | 943a6f2b0d | darwin | arm64 | 1.22.0 | 0 | 0 |
| **releng-hardware/gecko-3-b-osx-1015** | generic-worker | 60.3.4 | multiuser | 943a6f2b0d | darwin | amd64 | 1.22.0 | 0 | 0 |
| **releng-hardware/gecko-3-b-osx-arm64** | generic-worker | 60.3.4 | multiuser | 943a6f2b0d | darwin | arm64 | 1.22.0 | 0 | 0 |
| **releng-hardware/gecko-t-linux-talos-1804** | generic-worker | 61.0.0 | simple | 3bd4419b4b | linux | amd64 | 1.22.1 | 0 | 0 |
| **releng-hardware/gecko-t-linux-talos-2404** | generic-worker | 88.0.2 | insecure | fcaf4a25fc | linux | amd64 | 1.24.5 | 0 | 0 |
| **releng-hardware/gecko-t-linux-talos-2404-relops-aje** | generic-worker | 88.0.2 | insecure | fcaf4a25fc | linux | amd64 | 1.24.5 | 0 | 0 |
| **releng-hardware/gecko-t-osx-1015-r8** | generic-worker | 60.3.4 | simple | 943a6f2b0d | darwin | amd64 | 1.22.0 | 0 | 0 |
| **releng-hardware/gecko-t-osx-1015-r8-staging** | generic-worker | 67.1.0 | insecure | 0e62d3bf79 | darwin | amd64 | 1.22.4 | 0 | 0 |
| **releng-hardware/gecko-t-osx-1400-r8** | generic-worker | 60.3.4 | simple | 943a6f2b0d | darwin | amd64 | 1.22.0 | 0 | 0 |
| **releng-hardware/gecko-t-osx-1400-r8-staging** | generic-worker | 100.5.0 | insecure | f5fc37cc8f | darwin | amd64 | 1.26.4 | 0 | 0 |
| **releng-hardware/gecko-t-osx-1500-m-vms** | generic-worker | 91.0.2 | multiuser | 06628b3721 | darwin | arm64 | 1.25.3 | 0 | 0 |
| **releng-hardware/gecko-t-osx-1500-m4** | generic-worker | 91.0.2 | multiuser | 06628b3721 | darwin | arm64 | 1.25.3 | 0 | 0 |
| **releng-hardware/gecko-t-osx-1500-m4-staging** | generic-worker | 100.5.0 | multiuser | f5fc37cc8f | darwin | arm64 | 1.26.4 | 0 | 0 |
| **releng-hardware/gecko-t-osx-2700-m4** | generic-worker | 91.0.2 | multiuser | 06628b3721 | darwin | arm64 | 1.25.3 | 0 | 0 |
| **releng-hardware/gecko-t-win7-32-hw** | generic-worker | 45.0.0 | multiuser | 988e8100b3 | windows | 386 | 1.19.3 | 0 | 0 |
| **releng-hardware/mozillavpn-b-1-osx** | generic-worker | 60.3.4 | multiuser | 943a6f2b0d | darwin | amd64 | 1.22.0 | 0 | 0 |
| **releng-hardware/mozillavpn-b-3-osx** | generic-worker | 60.3.4 | multiuser | 943a6f2b0d | darwin | amd64 | 1.22.0 | 0 | 0 |
| **releng-hardware/nss-1-b-osx-1015** | generic-worker | 60.3.4 | multiuser | 943a6f2b0d | darwin | amd64 | 1.22.0 | 0 | 0 |
| **releng-hardware/nss-3-b-osx-1015** | generic-worker | 60.3.4 | multiuser | 943a6f2b0d | darwin | amd64 | 1.22.0 | 0 | 0 |
| **releng-hardware/win11-64-24h2-hw** | generic-worker | 110.0.0 | multiuser | 906277e693 | windows | amd64 | 1.27.1 | 0 | 0 |
| **releng-hardware/win11-64-24h2-hw-alpha** | generic-worker | 110.0.0 | multiuser | 906277e693 | windows | amd64 | 1.27.1 | 0 | 0 |
| **releng-hardware/win11-64-24h2-hw-perf-debug** | generic-worker | 109.0.0 | multiuser | 5b6d78a2d7 | windows | amd64 | 1.27.1 | 0 | 0 |
| **releng-hardware/win11-64-24h2-hw-ref** | generic-worker | 110.0.0 | multiuser | 906277e693 | windows | amd64 | 1.27.1 | 0 | 0 |
| **releng-hardware/win11-64-24h2-hw-ref-alpha** | generic-worker | 110.0.0 | multiuser | 906277e693 | windows | amd64 | 1.27.1 | 0 | 0 |
| **releng-t/linux-docker** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 311 | 311 |
| **relops-1/decision** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 2 | 2 |
| **relops-1/images** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 2 | 2 |
| **relops-3/decision** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 44 | 44 |
| **relops-3/images** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 2 | 2 |
| **reviewer-assignment/bot** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 7 | 7 |
| **scriptworker-1/decision** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 26 | 26 |
| **scriptworker-1/images** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 153 | 153 |
| **scriptworker-3/decision** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 18 | 18 |
| **scriptworker-3/images** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 122 | 122 |
| **taskgraph-1/decision** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 26 | 26 |
| **taskgraph-1/images** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 32 | 32 |
| **taskgraph-3/decision** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 12 | 12 |
| **taskgraph-3/images** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 13 | 13 |
| **taskgraph-t/linux-docker** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 175 | 175 |
| **taskgraph-t/win11-64-24h2** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | windows | amd64 | 1.27.1 | 26 | 26 |
| **taskgraph-t/win11-64-25h2** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | windows | amd64 | 1.27.1 | 2 | 2 |
| **translations-1/b-linux-large-gcp-1tb-32-256-d2g** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 13 | 13 |
| **translations-1/b-linux-large-gcp-1tb-32-256-std-d2g** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 4 | 4 |
| **translations-1/b-linux-large-gcp-1tb-64-512-d2g** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 2 | 2 |
| **translations-1/b-linux-large-gcp-1tb-64-512-std-d2g** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 2 | 2 |
| **translations-1/b-linux-large-gcp-d2g** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 39 | 39 |
| **translations-1/b-linux-large-gcp-d2g-1tb** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 2 | 2 |
| **translations-1/b-linux-large-gcp-d2g-1tb-standard** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 2 | 2 |
| **translations-1/b-linux-large-gcp-d2g-300gb** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 82 | 82 |
| **translations-1/b-linux-v100-gpu-d2g** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 138 | 138 |
| **translations-1/b-linux-v100-gpu-d2g-4** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 32 | 32 |
| **translations-1/b-linux-v100-gpu-d2g-4-1tb** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 31 | 31 |
| **translations-1/b-linux-v100-gpu-d2g-4-1tb-alpha** | generic-worker | 84.1.2 | multiuser | 4a139f395a | linux | amd64 | 1.24.4 | 2 | 2 |
| **translations-1/b-linux-v100-gpu-d2g-4-1tb-standard** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 2 | 2 |
| **translations-1/b-linux-v100-gpu-d2g-4-1tb-std-alpha** | generic-worker | 84.1.2 | multiuser | 4a139f395a | linux | amd64 | 1.24.4 | 2 | 2 |
| **translations-1/b-linux-v100-gpu-d2g-4-2tb** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 7 | 7 |
| **translations-1/b-linux-v100-gpu-d2g-4-2tb-alpha** | generic-worker | 84.1.2 | multiuser | 4a139f395a | linux | amd64 | 1.24.4 | 2 | 2 |
| **translations-1/b-linux-v100-gpu-d2g-4-300gb** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 6 | 6 |
| **translations-1/b-linux-v100-gpu-d2g-4-300gb-alpha** | generic-worker | 84.1.2 | multiuser | 4a139f395a | linux | amd64 | 1.24.4 | 3 | 3 |
| **translations-1/b-linux-v100-gpu-d2g-4-300gb-standard** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 3 | 3 |
| **translations-1/b-linux-v100-gpu-d2g-4-300gb-std-alpha** | generic-worker | 84.1.2 | multiuser | 4a139f395a | linux | amd64 | 1.24.4 | 2 | 2 |
| **translations-1/b-linux-v100-gpu-d2g-4-alpha** | generic-worker | 84.1.2 | multiuser | 4a139f395a | linux | amd64 | 1.24.4 | 2 | 2 |
| **translations-1/b-linux-v100-gpu-d2g-alpha** | generic-worker | 84.1.2 | multiuser | 4a139f395a | linux | amd64 | 1.24.4 | 2 | 2 |
| **translations-1/decision** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 21 | 21 |
| **translations-1/images** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 2 | 2 |
| **translations-t/t-linux-gw-noscratch-amd** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 2 | 2 |
| **xpi-1/b-linux** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 75 | 75 |
| **xpi-1/decision** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 39 | 39 |
| **xpi-1/images** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 2 | 2 |
| **xpi-3/b-linux** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 3 | 3 |
| **xpi-3/decision** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 7 | 7 |
| **xpi-3/images** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 2 | 2 |


## Docker Worker

Total: `2`

Count by version:

| Version | Count |
| :--- | ---: |
| 38.0.5 | 1 |
| 44.23.4 | 1 |


Count by image:

| Version | Count |
| :--- | ---: |
| ami-03e4f8db63254ce7e,ami-0a6e926238859761c,ami-0b5dd0bbb670ec80e | 1 |
| projects/taskcluster-imaging/global/images/docker-firefoxci-gcp-l1-googlecompute-2025-06-13t18-31-38z | 1 |


| Worker Pool | Implementation | Version | Total Workers | Total Capacity |
| --- | --- | --- | ---: | ---: |
| **infra/build-decision** | docker-worker | 38.0.5 | 5 | 40 |
| **proj-fuzzing/bugmon-pernosco** | docker-worker | 44.23.4 | 2 | 2 |


## Script Worker

Total: `42`



| Worker Pool | Implementation | Version | Total Workers | Total Capacity |
| --- | --- | --- | ---: | ---: |
| **scriptworker-k8s/comm-3-balrog** | Scriptworker | <no value> | 0 | 0 |
| **scriptworker-k8s/comm-3-beetmover** | Scriptworker | <no value> | 0 | 0 |
| **scriptworker-k8s/comm-3-bouncer** | Scriptworker | <no value> | 0 | 0 |
| **scriptworker-k8s/comm-3-lando** | Scriptworker | <no value> | 0 | 0 |
| **scriptworker-k8s/comm-3-pushflatpak** | Scriptworker | <no value> | 0 | 0 |
| **scriptworker-k8s/comm-3-pushmsix** | Scriptworker | <no value> | 0 | 0 |
| **scriptworker-k8s/comm-3-shipit** | Scriptworker | <no value> | 0 | 0 |
| **scriptworker-k8s/comm-3-signing** | Scriptworker | <no value> | 0 | 0 |
| **scriptworker-k8s/comm-t-signing** | Scriptworker | <no value> | 0 | 0 |
| **scriptworker-k8s/enterprise-3-signing** | Scriptworker | <no value> | 0 | 0 |
| **scriptworker-k8s/enterprise-t-signing** | Scriptworker | <no value> | 0 | 0 |
| **scriptworker-k8s/gecko-1-balrog** | Scriptworker | <no value> | 0 | 0 |
| **scriptworker-k8s/gecko-1-beetmover** | Scriptworker | <no value> | 0 | 0 |
| **scriptworker-k8s/gecko-1-bouncer** | Scriptworker | <no value> | 0 | 0 |
| **scriptworker-k8s/gecko-1-lando** | Scriptworker | <no value> | 0 | 0 |
| **scriptworker-k8s/gecko-3-balrog** | Scriptworker | <no value> | 0 | 0 |
| **scriptworker-k8s/gecko-3-beetmover** | Scriptworker | <no value> | 0 | 0 |
| **scriptworker-k8s/gecko-3-bouncer** | Scriptworker | <no value> | 0 | 0 |
| **scriptworker-k8s/gecko-3-lando** | Scriptworker | <no value> | 0 | 0 |
| **scriptworker-k8s/gecko-3-pushapk** | Scriptworker | <no value> | 0 | 0 |
| **scriptworker-k8s/gecko-3-shipit** | Scriptworker | <no value> | 0 | 0 |
| **scriptworker-k8s/gecko-3-signing** | Scriptworker | <no value> | 0 | 0 |
| **scriptworker-k8s/gecko-t-signing** | Scriptworker | <no value> | 0 | 0 |
| **scriptworker-k8s/mobile-3-bitrise** | Scriptworker | <no value> | 0 | 0 |
| **scriptworker-k8s/mobile-3-pushapk** | Scriptworker | <no value> | 0 | 0 |
| **scriptworker-k8s/mobile-3-signing** | Scriptworker | <no value> | 0 | 0 |
| **scriptworker-k8s/mobile-t-signing** | Scriptworker | <no value> | 0 | 0 |
| **scriptworker-k8s/mozillavpn-1-beetmover** | Scriptworker | <no value> | 0 | 0 |
| **scriptworker-k8s/mozillavpn-3-beetmover** | Scriptworker | <no value> | 0 | 0 |
| **scriptworker-k8s/mozillavpn-3-signing** | Scriptworker | <no value> | 0 | 0 |
| **scriptworker-k8s/translations-1-beetmover** | Scriptworker | <no value> | 0 | 0 |
| **scriptworker-k8s/xpi-t-signing** | Scriptworker | <no value> | 0 | 0 |
| **scriptworker-prov-v1/adhoc-signing-mac14m2** | Scriptworker | <no value> | 0 | 0 |
| **scriptworker-prov-v1/comm-signing-mac14m2** | Scriptworker | <no value> | 0 | 0 |
| **scriptworker-prov-v1/dep-adhoc-signing-mac14m2** | Scriptworker | <no value> | 0 | 0 |
| **scriptworker-prov-v1/dep-comm-signing-mac14m2** | Scriptworker | <no value> | 0 | 0 |
| **scriptworker-prov-v1/dep-enterprise-signing-mac14m2** | Scriptworker | <no value> | 0 | 0 |
| **scriptworker-prov-v1/dep-gecko-signing-mac14m2** | Scriptworker | <no value> | 0 | 0 |
| **scriptworker-prov-v1/dep-mozillavpn-signing-mac14m2** | Scriptworker | <no value> | 0 | 0 |
| **scriptworker-prov-v1/enterprise-signing-mac14m2** | Scriptworker | <no value> | 0 | 0 |
| **scriptworker-prov-v1/gecko-signing-mac14m2** | Scriptworker | <no value> | 0 | 0 |
| **scriptworker-prov-v1/mozillavpn-signing-mac14m2** | Scriptworker | <no value> | 0 | 0 |


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
| **gecko-t/t-linux-vm-2204-wayland** |  | No artifacts found | 23 | 23 |
| **gecko-t/t-linux-vm-2204-wayland-snap** |  | No artifacts found | 2 | 2 |


## Version not determined [^2]

Total: `24`


Count by image:

| Version | Count |
| :--- | ---: |
| ami-02619e55246806e8d | 1 |
| projects/taskcluster-imaging/global/images/generic-worker-ubuntu-24-04-arm64-mqscbutmsbawtytlroxk | 1 |
| projects/taskcluster-imaging/global/images/generic-worker-ubuntu-24-04-9ed5a812d3d844ac9e4e | 7 |
| projects/taskcluster-imaging/global/images/gw-fxci-gcp-l1-2404-amd64-googlecompute-alpha | 1 |
| unknown | 7 |
|  | 7 |


| Worker Pool | Implementation | Version | Total Workers | Total Capacity |
| --- | --- | --- | ---: | ---: |
| **gecko-1/b-win2022-headless** |  | Version not determined; task not (yet) claimed | 2 | 2 |
| **gecko-1/b-win2025-alpha** |  | Version not determined; task not (yet) claimed | 272 | 272 |
| **gecko-1/b-win2025-headless** |  | Version not determined; task not (yet) claimed | 2 | 2 |
| **gecko-3/b-win2022-headless** |  | Version not determined; task not (yet) claimed | 2 | 2 |
| **gecko-3/b-win2025-headless** |  | Version not determined; task not (yet) claimed | 2 | 2 |
| **gecko-t/dev-drive-validation-arm64** |  | Version not determined; task not (yet) claimed | 0 | 0 |
| **gecko-t/dev-drive-validation-arm64-fallback** |  | Version not determined; task not (yet) claimed | 0 | 0 |
| **gecko-t/dev-drive-validation-win10** |  | Version not determined; task not (yet) claimed | 0 | 0 |
| **gecko-t/dev-drive-validation-x64** |  | Version not determined; task not (yet) claimed | 0 | 0 |
| **gecko-t/t-linux-2404-relsre** |  | Version not determined; task not (yet) claimed | 3 | 3 |
| **mozillavpn-1/b-win2025-alpha** |  | Version not determined; task not (yet) claimed | 270 | 270 |
| **nss-1/b-win2025-alpha** |  | Version not determined; task not (yet) claimed | 269 | 269 |
| **proj-fuzzing/bugmon-monitor** |  | Version not determined; task not (yet) claimed | 96 | 96 |
| **proj-fuzzing/bugmon-pernosco-staging** |  | Version not determined; task not (yet) claimed | 3 | 3 |
| **proj-fuzzing/bugmon-processor** |  | Version not determined; task not (yet) claimed | 3 | 3 |
| **proj-fuzzing/ci** |  | Version not determined; task not (yet) claimed | 98 | 98 |
| **proj-fuzzing/ci-arm64** |  | Version not determined; task not (yet) claimed | 3 | 3 |
| **proj-fuzzing/decision** |  | Version not determined; task not (yet) claimed | 231 | 231 |
| **proj-fuzzing/grizzly-reduce-worker** |  | Version not determined; task not (yet) claimed | 3 | 3 |
| **proj-fuzzing/grizzly-reduce-worker-android** |  | Version not determined; task not (yet) claimed | 4 | 4 |
| **proj-fuzzing/nss-corpus-update-worker** |  | Version not determined; task not (yet) claimed | 6 | 6 |
| **releng-hardware/gecko-t-linux-netperf-1804** |  | Version not determined; task not (yet) claimed | 0 | 0 |
| **releng-hardware/gecko-t-linux-netperf-2404** |  | Version not determined; task not (yet) claimed | 0 | 0 |
| **releng-hardware/gecko-t-linux-talos-1804-relops-aje** |  | Version not determined; task not (yet) claimed | 0 | 0 |



[^1]: Those are the pools whose tasks were claimed and resolved by a worker as expected, but the worker did not publish either artifact `public/logs/live_backing.log` nor `public/logs/chain_of_trust.log`, which is the source used to identify the worker implementation.

[^2]: Probing task remains pending after two hours. Those are the pools that were not able to start any worker to claim the task within two hours.

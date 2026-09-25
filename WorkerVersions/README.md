

# Worker Pool Versions


## Generic Worker

Total: `398`

Count by version:

| Version | Count |
| :--- | ---: |
| 100.5.0 | 1 |
| 108.1.0 | 10 |
| 109.0.0 | 1 |
| 110.0.0 | 3 |
| 110.1.0 | 319 |
| 36.0.0 | 3 |
| 45.0.0 | 1 |
| 60.3.4 | 13 |
| 61.0.0 | 2 |
| 64.3.0 | 21 |
| 65.1.0 | 1 |
| 67.1.0 | 1 |
| 84.1.2 | 7 |
| 88.0.2 | 3 |
| 91.0.2 | 3 |
| 96.2.3 | 9 |


Count by image:

| Version | Count |
| :--- | ---: |
| unknown | 105 |
| projects/fxci-production-level3-workers/global/images/gw-fxci-gcp-l3-2404-arm64-headless-googlecompute-2026-09-23 | 8 |
| projects/taskcluster-imaging/global/images/gw-fxci-gcp-l1-arm64-gui-googlecompute-2024-09-18t14-57-52z | 5 |
| projects/taskcluster-imaging/global/images/gw-fxci-gcp-l1-gui-googlecompute-2025-01-13t22-33-40z | 2 |
| projects/fxci-production-level3-workers/global/images/gw-fxci-gcp-l3-arm64-gui-googlecompute-2024-09-18t19-02-58z | 2 |
| projects/taskcluster-imaging/global/images/gw-fxci-gcp-l1-2404-arm64-headless-googlecompute-alpha | 1 |
| projects/fxci-production-level3-workers/global/images/gw-fxci-gcp-l3-2404-amd64-headless-googlecompute-2026-09-23 | 60 |
| projects/taskcluster-imaging/global/images/gw-fxci-gcp-l1-gui-googlecompute-2024-08-22t22-48-09z | 10 |
| projects/taskcluster-imaging/global/images/gw-fxci-gcp-l1-2404-amd64-headless-googlecompute-2026-09-23 | 128 |
| projects/taskcluster-imaging/global/images/gw-fxci-gcp-l1-2404-amd64-headless-googlecompute-alpha-tc | 7 |
| projects/taskcluster-imaging/global/images/generic-worker-ubuntu-24-04-9ed5a812d3d844ac9e4e | 4 |
| projects/fxci-production-level3-workers/global/images/gw-fxci-gcp-l3-gui-googlecompute-2024-09-18t05-46-31z | 2 |
| projects/taskcluster-imaging/global/images/gw-fxci-gcp-l1-2404-amd64-gui-googlecompute-2026-09-23 | 5 |
| projects/taskcluster-imaging/global/images/gw-fxci-gcp-l1-2404-amd64-headless-googlecompute-alpha | 6 |
|  | 32 |
| projects/taskcluster-imaging/global/images/gw-fxci-gcp-l1-2404-amd64-gui-googlecompute-alpha | 4 |
| projects/taskcluster-imaging/global/images/gw-fxci-gcp-l1-2404-arm64-headless-googlecompute-2026-09-23 | 16 |
| projects/fxci-production-level3-workers/global/images/gw-fxci-gcp-l3-2404-amd64-headless-googlecompute-alpha | 1 |


| Worker Pool | Implementation | Version | Engine | Revision | OS | Arch | GO | Total Workers | Total Capacity |
| --- | --- | --- | --- | --- | --- | --- | --- | ---: | ---: |
| **adhoc-1/decision** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 3 | 3 |
| **adhoc-1/images** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 10 | 10 |
| **adhoc-3/decision** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 5 | 5 |
| **adhoc-3/images** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 4 | 4 |
| **app-services-1/b-linux** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 551 | 551 |
| **app-services-1/b-linux-docker-amd** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 2 | 2 |
| **app-services-1/decision** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 6 | 6 |
| **app-services-1/images** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 2 | 2 |
| **app-services-3/b-linux** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 1353 | 1353 |
| **app-services-3/b-linux-docker-amd** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 3 | 3 |
| **app-services-3/decision** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 33 | 33 |
| **app-services-3/images** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 2 | 2 |
| **code-analysis-1/decision** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 2 | 2 |
| **code-analysis-1/linux-docker** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 2 | 2 |
| **code-analysis-1/linux-gw-gcp** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | amd64 | 1.22.2 | 48 | 48 |
| **code-analysis-3/decision** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 2 | 2 |
| **code-analysis-3/linux-docker** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 2 | 2 |
| **code-analysis-3/linux-gw-gcp** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | amd64 | 1.22.2 | 69 | 69 |
| **code-coverage/bot** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 257 | 257 |
| **code-review/bot** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 7029 | 7029 |
| **comm-1/b-linux** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 570 | 570 |
| **comm-1/b-linux-aarch64** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | arm64 | 1.27.1 | 2 | 2 |
| **comm-1/b-linux-docker-amd** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 2412 | 2412 |
| **comm-1/b-linux-docker-large-amd** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 24 | 24 |
| **comm-1/b-linux-docker-xlarge-amd** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 8 | 8 |
| **comm-1/b-linux-gcp-aarch64** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | arm64 | 1.22.2 | 2 | 2 |
| **comm-1/b-linux-large** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 2 | 2 |
| **comm-1/b-linux-xlarge** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 3 | 3 |
| **comm-1/b-win2022** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | windows | amd64 | 1.27.1 | 2 | 2 |
| **comm-1/b-win2025** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | windows | amd64 | 1.27.1 | 2 | 2 |
| **comm-1/decision** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 427 | 427 |
| **comm-1/images** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 44 | 44 |
| **comm-1/images-aarch64** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | arm64 | 1.27.1 | 2 | 2 |
| **comm-2/b-linux** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 2 | 2 |
| **comm-2/b-linux-aarch64** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | arm64 | 1.27.1 | 2 | 2 |
| **comm-2/b-linux-docker-amd** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 3 | 3 |
| **comm-2/b-linux-docker-large-amd** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 2 | 2 |
| **comm-2/b-linux-docker-xlarge-amd** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 3 | 3 |
| **comm-2/b-linux-gcp-aarch64** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | arm64 | 1.22.2 | 2 | 2 |
| **comm-2/b-linux-large** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 2 | 2 |
| **comm-2/b-linux-xlarge** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 3 | 3 |
| **comm-2/decision** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 2 | 2 |
| **comm-2/images** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 2 | 2 |
| **comm-2/images-aarch64** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | arm64 | 1.27.1 | 2 | 2 |
| **comm-3/b-linux** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 4269 | 4269 |
| **comm-3/b-linux-aarch64** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | arm64 | 1.27.1 | 48 | 48 |
| **comm-3/b-linux-docker-amd** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 5236 | 5236 |
| **comm-3/b-linux-docker-large-amd** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 138 | 138 |
| **comm-3/b-linux-docker-xlarge-amd** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 40 | 40 |
| **comm-3/b-linux-gcp-aarch64** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | arm64 | 1.22.2 | 2 | 2 |
| **comm-3/b-linux-large** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 3 | 3 |
| **comm-3/b-linux-xlarge** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 2 | 2 |
| **comm-3/b-win2022** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | windows | amd64 | 1.27.1 | 672 | 672 |
| **comm-3/b-win2025** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | windows | amd64 | 1.27.1 | 2 | 2 |
| **comm-3/decision** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 183 | 183 |
| **comm-3/images** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 38 | 38 |
| **comm-3/images-aarch64** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | arm64 | 1.27.1 | 2 | 2 |
| **comm-t/misc** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 182 | 182 |
| **comm-t/t-linux-arm64-docker** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | arm64 | 1.27.1 | 16 | 16 |
| **comm-t/t-linux-docker** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 2 | 2 |
| **comm-t/t-linux-docker-amd** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 3680 | 3680 |
| **comm-t/t-linux-docker-noscratch** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 3 | 3 |
| **comm-t/t-linux-docker-noscratch-amd** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 13361 | 13361 |
| **comm-t/win11-64-2009** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 2 | 2 |
| **comm-t/win11-64-2009-ssd** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 2 | 2 |
| **comm-t/win11-64-24h2** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | windows | amd64 | 1.27.1 | 2 | 2 |
| **comm-t/win11-64-24h2-ssd** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | windows | amd64 | 1.27.1 | 2 | 2 |
| **comm-t/win11-64-25h2** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | windows | amd64 | 1.27.1 | 5982 | 5982 |
| **comm-t/win11-64-25h2-ssd** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | windows | amd64 | 1.27.1 | 2 | 2 |
| **comm-t/win11-a64-25h2** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | windows | arm64 | 1.27.1 | 9 | 9 |
| **enterprise-1/b-linux** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 4228 | 4228 |
| **enterprise-1/b-linux-aarch64** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | arm64 | 1.27.1 | 4 | 4 |
| **enterprise-1/b-linux-docker-amd** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 1187 | 1187 |
| **enterprise-1/b-linux-docker-large-amd** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 30 | 30 |
| **enterprise-1/b-linux-docker-xlarge-amd** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 4048 | 4048 |
| **enterprise-1/b-win2022** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | windows | amd64 | 1.27.1 | 270 | 270 |
| **enterprise-1/b-win2025** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | windows | amd64 | 1.27.1 | 2 | 2 |
| **enterprise-1/decision** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 334 | 334 |
| **enterprise-1/images** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 86 | 86 |
| **enterprise-1/images-aarch64** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | arm64 | 1.27.1 | 13 | 13 |
| **enterprise-1/win11-a64-25h2-builder** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | windows | arm64 | 1.27.1 | 2 | 2 |
| **enterprise-3/b-linux** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 12817 | 12817 |
| **enterprise-3/b-linux-aarch64** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | arm64 | 1.27.1 | 49 | 49 |
| **enterprise-3/b-linux-docker-amd** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 2210 | 2210 |
| **enterprise-3/b-linux-docker-large-amd** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 142 | 142 |
| **enterprise-3/b-linux-docker-xlarge-amd** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 3857 | 3857 |
| **enterprise-3/b-win2022** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | windows | amd64 | 1.27.1 | 1581 | 1581 |
| **enterprise-3/b-win2025** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | windows | amd64 | 1.27.1 | 2 | 2 |
| **enterprise-3/decision** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 149 | 149 |
| **enterprise-3/images** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 44 | 44 |
| **enterprise-3/images-aarch64** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | arm64 | 1.27.1 | 12 | 12 |
| **enterprise-3/win11-a64-25h2-builder** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | windows | arm64 | 1.27.1 | 2 | 2 |
| **enterprise-t/misc** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 404 | 404 |
| **enterprise-t/t-linux-2204-wayland** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | amd64 | 1.22.2 | 1565 | 1565 |
| **enterprise-t/t-linux-2204-wayland-snap** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | amd64 | 1.22.2 | 2 | 2 |
| **enterprise-t/t-linux-2404-wayland-snap** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 2 | 2 |
| **enterprise-t/t-linux-2404-wayland-snap-alpha** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 2 | 2 |
| **enterprise-t/t-linux-docker** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 2 | 2 |
| **enterprise-t/t-linux-docker-16c32gb-amd** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 2769 | 2769 |
| **enterprise-t/t-linux-docker-amd** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 12574 | 12574 |
| **enterprise-t/t-linux-docker-kvm** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 2 | 2 |
| **enterprise-t/t-linux-docker-noscratch-amd** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 27989 | 27989 |
| **enterprise-t/t-linux-xlarge-2204-wayland** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | amd64 | 1.22.2 | 2 | 2 |
| **enterprise-t/win10-64-2009** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | windows | amd64 | 1.27.1 | 1032 | 1032 |
| **enterprise-t/win11-64-24h2** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | windows | amd64 | 1.27.1 | 2 | 2 |
| **enterprise-t/win11-64-24h2-large** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | windows | amd64 | 1.27.1 | 2 | 2 |
| **enterprise-t/win11-64-24h2-source** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | windows | amd64 | 1.27.1 | 2 | 2 |
| **enterprise-t/win11-64-25h2** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | windows | amd64 | 1.27.1 | 8181 | 8181 |
| **enterprise-t/win11-64-25h2-gpu** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | windows | amd64 | 1.27.1 | 404 | 404 |
| **enterprise-t/win11-64-25h2-large** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | windows | amd64 | 1.27.1 | 294 | 294 |
| **enterprise-t/win11-64-25h2-privileged** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | windows | amd64 | 1.27.1 | 2 | 2 |
| **enterprise-t/win11-64-25h2-source** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | windows | amd64 | 1.27.1 | 1060 | 1060 |
| **enterprise-t/win11-64-25h2-webgpu** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | windows | amd64 | 1.27.1 | 3 | 3 |
| **gecko-1/b-linux** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 8022 | 8022 |
| **gecko-1/b-linux-2204-kvm-gcp** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | amd64 | 1.22.2 | 2 | 2 |
| **gecko-1/b-linux-aarch64** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | arm64 | 1.27.1 | 356 | 356 |
| **gecko-1/b-linux-docker-alpha** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 2 | 2 |
| **gecko-1/b-linux-docker-amd** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 20188 | 20188 |
| **gecko-1/b-linux-docker-large-amd** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 2685 | 2685 |
| **gecko-1/b-linux-docker-updatebot-amd** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 2 | 2 |
| **gecko-1/b-linux-docker-xlarge-amd** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 2941 | 2941 |
| **gecko-1/b-linux-gcp-aarch64** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | arm64 | 1.22.2 | 2 | 2 |
| **gecko-1/b-linux-kvm** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 247 | 247 |
| **gecko-1/b-linux-large** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 2 | 2 |
| **gecko-1/b-linux-medium** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 17260 | 17260 |
| **gecko-1/b-linux-xlarge** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 7 | 7 |
| **gecko-1/b-win2022** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | windows | amd64 | 1.27.1 | 999 | 999 |
| **gecko-1/b-win2022-headless** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | windows | amd64 | 1.27.1 | 3 | 3 |
| **gecko-1/b-win2022-updatebot** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | windows | amd64 | 1.27.1 | 3 | 3 |
| **gecko-1/b-win2022-xxlarge** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | windows | amd64 | 1.27.1 | 2 | 2 |
| **gecko-1/b-win2025** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | windows | amd64 | 1.27.1 | 2 | 2 |
| **gecko-1/b-win2025-headless** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | windows | amd64 | 1.27.1 | 2 | 2 |
| **gecko-1/b-win2025-updatebot** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | windows | amd64 | 1.27.1 | 2 | 2 |
| **gecko-1/b-win2025-xxlarge** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | windows | amd64 | 1.27.1 | 2 | 2 |
| **gecko-1/decision** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 1660 | 1660 |
| **gecko-1/images** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 224 | 224 |
| **gecko-1/images-aarch64** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | arm64 | 1.27.1 | 37 | 37 |
| **gecko-1/win11-a64-25h2-builder** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | windows | arm64 | 1.27.1 | 27 | 27 |
| **gecko-2/b-linux** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 210 | 210 |
| **gecko-2/b-linux-aarch64** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | arm64 | 1.27.1 | 2 | 2 |
| **gecko-2/b-linux-docker-amd** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 1012 | 1012 |
| **gecko-2/b-linux-docker-large-amd** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 130 | 130 |
| **gecko-2/b-linux-docker-updatebot-amd** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 4 | 4 |
| **gecko-2/b-linux-docker-xlarge-amd** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 209 | 209 |
| **gecko-2/b-linux-gcp-aarch64** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | arm64 | 1.22.2 | 2 | 2 |
| **gecko-2/b-linux-kvm** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 20 | 20 |
| **gecko-2/b-linux-large** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 2 | 2 |
| **gecko-2/b-linux-medium** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 338 | 338 |
| **gecko-2/b-linux-xlarge** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 2 | 2 |
| **gecko-2/b-win2022** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | windows | amd64 | 1.27.1 | 17 | 17 |
| **gecko-2/b-win2022-updatebot** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | windows | amd64 | 1.27.1 | 2 | 2 |
| **gecko-2/b-win2025** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | windows | amd64 | 1.27.1 | 2 | 2 |
| **gecko-2/b-win2025-updatebot** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | windows | amd64 | 1.27.1 | 2 | 2 |
| **gecko-2/decision** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 18 | 18 |
| **gecko-2/images** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 2 | 2 |
| **gecko-2/images-aarch64** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | arm64 | 1.27.1 | 2 | 2 |
| **gecko-3/b-linux** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 33738 | 33738 |
| **gecko-3/b-linux-2204-kvm-gcp** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | amd64 | 1.22.2 | 2 | 2 |
| **gecko-3/b-linux-aarch64** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | arm64 | 1.27.1 | 9 | 9 |
| **gecko-3/b-linux-docker-alpha** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 2 | 2 |
| **gecko-3/b-linux-docker-amd** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 62696 | 62696 |
| **gecko-3/b-linux-docker-large-amd** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 8001 | 8001 |
| **gecko-3/b-linux-docker-updatebot-amd** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 75 | 75 |
| **gecko-3/b-linux-docker-xlarge-amd** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 10503 | 10503 |
| **gecko-3/b-linux-gcp-aarch64** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | arm64 | 1.22.2 | 2 | 2 |
| **gecko-3/b-linux-kvm** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 1000 | 1000 |
| **gecko-3/b-linux-large** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 2 | 2 |
| **gecko-3/b-linux-medium** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 33524 | 33524 |
| **gecko-3/b-linux-xlarge** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 213 | 213 |
| **gecko-3/b-win2022** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | windows | amd64 | 1.27.1 | 7151 | 7151 |
| **gecko-3/b-win2022-headless** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | windows | amd64 | 1.27.1 | 2 | 2 |
| **gecko-3/b-win2022-updatebot** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | windows | amd64 | 1.27.1 | 2 | 2 |
| **gecko-3/b-win2022-xxlarge** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | windows | amd64 | 1.27.1 | 12 | 12 |
| **gecko-3/b-win2025** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | windows | amd64 | 1.27.1 | 2 | 2 |
| **gecko-3/b-win2025-headless** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | windows | amd64 | 1.27.1 | 2 | 2 |
| **gecko-3/b-win2025-updatebot** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | windows | amd64 | 1.27.1 | 2 | 2 |
| **gecko-3/b-win2025-xxlarge** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | windows | amd64 | 1.27.1 | 2 | 2 |
| **gecko-3/decision** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 7494 | 7494 |
| **gecko-3/images** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 106 | 106 |
| **gecko-3/images-aarch64** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | arm64 | 1.27.1 | 19 | 19 |
| **gecko-3/win11-a64-25h2-builder** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | windows | arm64 | 1.27.1 | 173 | 173 |
| **gecko-t/misc** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 1529 | 1529 |
| **gecko-t/t-linux-2204-wayland** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | amd64 | 1.22.2 | 8144 | 8144 |
| **gecko-t/t-linux-2204-wayland-arm64-relsre** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | arm64 | 1.22.2 | 2 | 2 |
| **gecko-t/t-linux-2204-wayland-relsre** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | amd64 | 1.22.2 | 2 | 2 |
| **gecko-t/t-linux-2204-wayland-root-exp** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | amd64 | 1.22.2 | 2 | 2 |
| **gecko-t/t-linux-2204-wayland-snap** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | amd64 | 1.22.2 | 115 | 115 |
| **gecko-t/t-linux-2404-headless-alpha** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 2 | 2 |
| **gecko-t/t-linux-2404-headless-arm64-alpha** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | arm64 | 1.27.1 | 2 | 2 |
| **gecko-t/t-linux-2404-headless-ssd-alpha** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 2 | 2 |
| **gecko-t/t-linux-2404-wayland** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 3 | 3 |
| **gecko-t/t-linux-2404-wayland-alpha** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 2 | 2 |
| **gecko-t/t-linux-2404-wayland-snap** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 143 | 143 |
| **gecko-t/t-linux-2404-wayland-snap-alpha** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 4 | 4 |
| **gecko-t/t-linux-arm64-docker** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | arm64 | 1.27.1 | 40 | 40 |
| **gecko-t/t-linux-docker** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 966 | 966 |
| **gecko-t/t-linux-docker-16c32gb-amd** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 8404 | 8404 |
| **gecko-t/t-linux-docker-amd** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 65356 | 65356 |
| **gecko-t/t-linux-docker-amd-alpha** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 4 | 4 |
| **gecko-t/t-linux-docker-kvm** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 94662 | 94662 |
| **gecko-t/t-linux-docker-kvm-alpha** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 22 | 22 |
| **gecko-t/t-linux-docker-noscratch** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 11701 | 11701 |
| **gecko-t/t-linux-docker-noscratch-amd** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 455636 | 455636 |
| **gecko-t/t-linux-docker-noscratch-amd-alpha** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 840 | 840 |
| **gecko-t/t-linux-xlarge-2204-wayland** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | amd64 | 1.22.2 | 6797 | 6797 |
| **gecko-t/t-linux-xlarge-2404-wayland** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 2 | 2 |
| **gecko-t/win10-64-2009** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | windows | amd64 | 1.27.1 | 6996 | 6996 |
| **gecko-t/win10-64-2009-gpu** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | windows | amd64 | 1.27.1 | 1370 | 1370 |
| **gecko-t/win10-64-2009-source** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | windows | amd64 | 1.27.1 | 125 | 125 |
| **gecko-t/win10-64-2009-webgpu** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | windows | amd64 | 1.27.1 | 2 | 2 |
| **gecko-t/win11-64-2009** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 2715 | 2715 |
| **gecko-t/win11-64-2009-alpha** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 3 | 3 |
| **gecko-t/win11-64-2009-source** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 13 | 13 |
| **gecko-t/win11-64-2009-source-alpha** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 2 | 2 |
| **gecko-t/win11-64-2009-ssd** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 2 | 2 |
| **gecko-t/win11-64-2009-webgpu** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 2 | 2 |
| **gecko-t/win11-64-2009-webgpu-alpha** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 5 | 5 |
| **gecko-t/win11-64-24h2** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | windows | amd64 | 1.27.1 | 3081 | 3081 |
| **gecko-t/win11-64-24h2-gpu** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | windows | amd64 | 1.27.1 | 2 | 2 |
| **gecko-t/win11-64-24h2-large** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | windows | amd64 | 1.27.1 | 244 | 244 |
| **gecko-t/win11-64-24h2-privileged** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | windows | amd64 | 1.27.1 | 2 | 2 |
| **gecko-t/win11-64-24h2-source** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | windows | amd64 | 1.27.1 | 13 | 13 |
| **gecko-t/win11-64-24h2-ssd** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | windows | amd64 | 1.27.1 | 2 | 2 |
| **gecko-t/win11-64-24h2-unprivileged** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | windows | amd64 | 1.27.1 | 2 | 2 |
| **gecko-t/win11-64-24h2-webgpu** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | windows | amd64 | 1.27.1 | 2 | 2 |
| **gecko-t/win11-64-25h2** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | windows | amd64 | 1.27.1 | 111985 | 111985 |
| **gecko-t/win11-64-25h2-gpu** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | windows | amd64 | 1.27.1 | 25363 | 25363 |
| **gecko-t/win11-64-25h2-large** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | windows | amd64 | 1.27.1 | 14192 | 14192 |
| **gecko-t/win11-64-25h2-privileged** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | windows | amd64 | 1.27.1 | 412 | 412 |
| **gecko-t/win11-64-25h2-source** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | windows | amd64 | 1.27.1 | 3392 | 3392 |
| **gecko-t/win11-64-25h2-ssd** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | windows | amd64 | 1.27.1 | 2 | 2 |
| **gecko-t/win11-64-25h2-ssd-gpu** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | windows | amd64 | 1.27.1 | 2 | 2 |
| **gecko-t/win11-64-25h2-unprivileged** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | windows | amd64 | 1.27.1 | 2 | 2 |
| **gecko-t/win11-64-25h2-webgpu** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | windows | amd64 | 1.27.1 | 4738 | 4738 |
| **gecko-t/win11-a64-25h2** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | windows | arm64 | 1.27.1 | 354 | 354 |
| **glean-1/b-linux** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 87 | 87 |
| **glean-1/decision** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 3 | 3 |
| **glean-1/images** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 2 | 2 |
| **glean-3/b-linux** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 92 | 92 |
| **glean-3/decision** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 26 | 26 |
| **glean-3/images** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 2 | 2 |
| **infra/build-decision-alpha** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 2 | 2 |
| **mobile-1/b-linux** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 254 | 254 |
| **mobile-1/decision** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 27 | 27 |
| **mobile-1/images** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 2 | 2 |
| **mobile-3/b-linux** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 268 | 268 |
| **mobile-3/decision** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 147 | 147 |
| **mobile-3/images** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 2 | 2 |
| **mozilla-1/b-linux** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 2 | 2 |
| **mozilla-1/b-win2022** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | windows | amd64 | 1.27.1 | 2 | 2 |
| **mozilla-1/b-win2025** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | windows | amd64 | 1.27.1 | 2 | 2 |
| **mozilla-1/decision** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 13 | 13 |
| **mozilla-1/images** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 19 | 19 |
| **mozilla-3/b-linux** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 2 | 2 |
| **mozilla-3/b-win2022** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | windows | amd64 | 1.27.1 | 2 | 2 |
| **mozilla-3/b-win2025** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | windows | amd64 | 1.27.1 | 2 | 2 |
| **mozilla-3/decision** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 467 | 467 |
| **mozilla-3/images** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 6 | 6 |
| **mozilla-t/pre-commit** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 86 | 86 |
| **mozilla-t/t-linux-2204-wayland** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | amd64 | 1.22.2 | 2 | 2 |
| **mozilla-t/t-linux-2204-wayland-relsre** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | amd64 | 1.22.2 | 2 | 2 |
| **mozilla-t/t-linux-2404-wayland** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 2 | 2 |
| **mozilla-t/t-linux-2404-wayland-alpha** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 1079 | 1079 |
| **mozilla-t/t-linux-arm64-docker** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | arm64 | 1.27.1 | 2 | 2 |
| **mozilla-t/t-linux-docker** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 371 | 371 |
| **mozilla-t/t-linux-docker-amd** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 2 | 2 |
| **mozilla-t/t-linux-docker-noscratch** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 2 | 2 |
| **mozilla-t/t-linux-docker-noscratch-amd** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 2 | 2 |
| **mozillavpn-1/b-linux** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 288 | 288 |
| **mozillavpn-1/b-linux-large** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 265 | 265 |
| **mozillavpn-1/b-win2022** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | windows | amd64 | 1.27.1 | 85 | 85 |
| **mozillavpn-1/b-win2025** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | windows | amd64 | 1.27.1 | 2 | 2 |
| **mozillavpn-1/decision** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 9 | 9 |
| **mozillavpn-1/images** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 35 | 35 |
| **mozillavpn-1/win11-a64-25h2-builder** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | windows | arm64 | 1.27.1 | 6 | 6 |
| **mozillavpn-3/b-linux** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 144 | 144 |
| **mozillavpn-3/b-linux-large** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 149 | 149 |
| **mozillavpn-3/b-win2022** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | windows | amd64 | 1.27.1 | 33 | 33 |
| **mozillavpn-3/b-win2025** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | windows | amd64 | 1.27.1 | 2 | 2 |
| **mozillavpn-3/decision** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 20 | 20 |
| **mozillavpn-3/images** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 5 | 5 |
| **mozillavpn-3/win11-a64-25h2-builder** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | windows | arm64 | 1.27.1 | 2 | 2 |
| **nss-1/b-linux-aarch64** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | arm64 | 1.27.1 | 259 | 259 |
| **nss-1/b-win2022** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | windows | amd64 | 1.27.1 | 336 | 336 |
| **nss-1/b-win2025** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | windows | amd64 | 1.27.1 | 2 | 2 |
| **nss-1/decision** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 42 | 42 |
| **nss-1/images** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 2 | 2 |
| **nss-1/images-aarch64** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | arm64 | 1.27.1 | 2 | 2 |
| **nss-1/linux-docker** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 1042 | 1042 |
| **nss-1/win11-a64-25h2** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | windows | arm64 | 1.27.1 | 3 | 3 |
| **nss-1/win11-a64-25h2-builder** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | windows | arm64 | 1.27.1 | 2 | 2 |
| **nss-3/b-linux-aarch64** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | arm64 | 1.27.1 | 564 | 564 |
| **nss-3/b-win2022** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | windows | amd64 | 1.27.1 | 256 | 256 |
| **nss-3/b-win2025** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | windows | amd64 | 1.27.1 | 2 | 2 |
| **nss-3/decision** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 23 | 23 |
| **nss-3/images** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 2 | 2 |
| **nss-3/images-aarch64** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | arm64 | 1.27.1 | 2 | 2 |
| **nss-3/linux-docker** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 1464 | 1464 |
| **nss-3/win11-a64-25h2-builder** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | windows | arm64 | 1.27.1 | 2 | 2 |
| **nss-t/t-linux-arm64-docker** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | arm64 | 1.27.1 | 2059 | 2059 |
| **nss-t/t-linux-docker** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 3116 | 3116 |
| **nss-t/t-linux-docker-amd** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 2 | 2 |
| **nss-t/win11-a64-25h2** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | windows | arm64 | 1.27.1 | 2 | 2 |
| **proj-autophone/gecko-t-bitbar-gw-perf-a55** | generic-worker | 36.0.0 | simple | b9cb11293f | linux | amd64 | 1.13.7 | 0 | 0 |
| **proj-autophone/gecko-t-bitbar-gw-perf-p6** | generic-worker | 36.0.0 | simple | b9cb11293f | linux | amd64 | 1.13.7 | 0 | 0 |
| **proj-autophone/gecko-t-bitbar-gw-perf-s24** | generic-worker | 36.0.0 | simple | b9cb11293f | linux | amd64 | 1.13.7 | 0 | 0 |
| **proj-fuzzing/bugmon-processor-windows** | generic-worker | 108.1.0 | multiuser | f86624a8bf | windows | amd64 | 1.27.1 | 3 | 3 |
| **proj-fuzzing/ci-windows** | generic-worker | 108.1.0 | multiuser | f86624a8bf | windows | amd64 | 1.27.1 | 2 | 2 |
| **proj-fuzzing/grizzly-reduce-worker-windows** | generic-worker | 108.1.0 | multiuser | f86624a8bf | windows | amd64 | 1.27.1 | 2 | 2 |
| **proj-fuzzing/grizzly-reduce-worker-windows-ngpu** | generic-worker | 108.1.0 | multiuser | f86624a8bf | windows | amd64 | 1.27.1 | 2 | 2 |
| **proj-taskcluster/ci** | generic-worker | 108.1.0 | multiuser | f86624a8bf | linux | amd64 | 1.27.1 | 60 | 60 |
| **proj-taskcluster/gw-ci-macos** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | darwin | arm64 | 1.27.1 | 2 | 2 |
| **proj-taskcluster/gw-ubuntu-24-04** | generic-worker | 108.1.0 | multiuser | f86624a8bf | linux | amd64 | 1.27.1 | 4225 | 4225 |
| **proj-taskcluster/gw-ubuntu-24-04-gui** | generic-worker | 108.1.0 | multiuser | f86624a8bf | linux | amd64 | 1.27.1 | 168 | 168 |
| **proj-taskcluster/gw-windows-2022** | generic-worker | 108.1.0 | multiuser | f86624a8bf | windows | amd64 | 1.27.1 | 84 | 84 |
| **proj-taskcluster/gw-windows-2022-gui** | generic-worker | 108.1.0 | multiuser | f86624a8bf | windows | amd64 | 1.27.1 | 87 | 87 |
| **proj-taskcluster/release** | generic-worker | 108.1.0 | multiuser | f86624a8bf | linux | amd64 | 1.27.1 | 4 | 4 |
| **releng-1/decision** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 72 | 72 |
| **releng-1/linux-docker** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 51 | 51 |
| **releng-3/decision** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 37 | 37 |
| **releng-3/linux-docker** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 68 | 68 |
| **releng-hardware/applicationservices-b-1-osx1015** | generic-worker | 60.3.4 | multiuser | 943a6f2b0d | darwin | amd64 | 1.22.0 | 0 | 0 |
| **releng-hardware/applicationservices-b-3-osx1015** | generic-worker | 60.3.4 | multiuser | 943a6f2b0d | darwin | amd64 | 1.22.0 | 0 | 0 |
| **releng-hardware/enterprise-1-b-osx-arm64** | generic-worker | 60.3.4 | multiuser | 943a6f2b0d | darwin | arm64 | 1.22.0 | 0 | 0 |
| **releng-hardware/enterprise-3-b-osx-arm64** | generic-worker | 60.3.4 | multiuser | 943a6f2b0d | darwin | arm64 | 1.22.0 | 0 | 0 |
| **releng-hardware/gecko-1-b-osx-1015** | generic-worker | 60.3.4 | multiuser | 943a6f2b0d | darwin | amd64 | 1.22.0 | 0 | 0 |
| **releng-hardware/gecko-1-b-osx-arm64** | generic-worker | 60.3.4 | multiuser | 943a6f2b0d | darwin | arm64 | 1.22.0 | 0 | 0 |
| **releng-hardware/gecko-3-b-osx-1015** | generic-worker | 60.3.4 | multiuser | 943a6f2b0d | darwin | amd64 | 1.22.0 | 0 | 0 |
| **releng-hardware/gecko-3-b-osx-arm64** | generic-worker | 60.3.4 | multiuser | 943a6f2b0d | darwin | arm64 | 1.22.0 | 0 | 0 |
| **releng-hardware/gecko-t-linux-netperf-1804** | generic-worker | 65.1.0 | insecure | 1a085daa37 | linux | amd64 | 1.22.3 | 0 | 0 |
| **releng-hardware/gecko-t-linux-netperf-2404** | generic-worker | 88.0.2 | insecure | fcaf4a25fc | linux | amd64 | 1.24.5 | 0 | 0 |
| **releng-hardware/gecko-t-linux-talos-1804** | generic-worker | 61.0.0 | simple | 3bd4419b4b | linux | amd64 | 1.22.1 | 0 | 0 |
| **releng-hardware/gecko-t-linux-talos-1804-relops-aje** | generic-worker | 61.0.0 | simple | 3bd4419b4b | linux | amd64 | 1.22.1 | 0 | 0 |
| **releng-hardware/gecko-t-linux-talos-2404** | generic-worker | 88.0.2 | insecure | fcaf4a25fc | linux | amd64 | 1.24.5 | 0 | 0 |
| **releng-hardware/gecko-t-linux-talos-2404-relops-aje** | generic-worker | 88.0.2 | insecure | fcaf4a25fc | linux | amd64 | 1.24.5 | 0 | 0 |
| **releng-hardware/gecko-t-osx-1015-r8** | generic-worker | 60.3.4 | simple | 943a6f2b0d | darwin | amd64 | 1.22.0 | 0 | 0 |
| **releng-hardware/gecko-t-osx-1015-r8-staging** | generic-worker | 67.1.0 | insecure | 0e62d3bf79 | darwin | amd64 | 1.22.4 | 0 | 0 |
| **releng-hardware/gecko-t-osx-1400-r8** | generic-worker | 60.3.4 | simple | 943a6f2b0d | darwin | amd64 | 1.22.0 | 0 | 0 |
| **releng-hardware/gecko-t-osx-1400-r8-staging** | generic-worker | 100.5.0 | insecure | f5fc37cc8f | darwin | amd64 | 1.26.4 | 0 | 0 |
| **releng-hardware/gecko-t-osx-1500-m-vms** | generic-worker | 91.0.2 | multiuser | 06628b3721 | darwin | arm64 | 1.25.3 | 0 | 0 |
| **releng-hardware/gecko-t-osx-1500-m4-staging** | generic-worker | 91.0.2 | multiuser | 06628b3721 | darwin | arm64 | 1.25.3 | 0 | 0 |
| **releng-hardware/gecko-t-osx-2700-m4** | generic-worker | 91.0.2 | multiuser | 06628b3721 | darwin | arm64 | 1.25.3 | 0 | 0 |
| **releng-hardware/gecko-t-win7-32-hw** | generic-worker | 45.0.0 | multiuser | 988e8100b3 | windows | 386 | 1.19.3 | 0 | 0 |
| **releng-hardware/mozillavpn-b-1-osx** | generic-worker | 60.3.4 | multiuser | 943a6f2b0d | darwin | amd64 | 1.22.0 | 0 | 0 |
| **releng-hardware/mozillavpn-b-3-osx** | generic-worker | 60.3.4 | multiuser | 943a6f2b0d | darwin | amd64 | 1.22.0 | 0 | 0 |
| **releng-hardware/nss-1-b-osx-1015** | generic-worker | 60.3.4 | multiuser | 943a6f2b0d | darwin | amd64 | 1.22.0 | 0 | 0 |
| **releng-hardware/win11-64-24h2-hw-alpha** | generic-worker | 110.0.0 | multiuser | 906277e693 | windows | amd64 | 1.27.1 | 0 | 0 |
| **releng-hardware/win11-64-24h2-hw-perf-debug** | generic-worker | 109.0.0 | multiuser | 5b6d78a2d7 | windows | amd64 | 1.27.1 | 0 | 0 |
| **releng-hardware/win11-64-24h2-hw-ref** | generic-worker | 110.0.0 | multiuser | 906277e693 | windows | amd64 | 1.27.1 | 0 | 0 |
| **releng-hardware/win11-64-24h2-hw-ref-alpha** | generic-worker | 110.0.0 | multiuser | 906277e693 | windows | amd64 | 1.27.1 | 0 | 0 |
| **releng-t/linux-docker** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 282 | 282 |
| **relops-1/decision** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 2 | 2 |
| **relops-1/images** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 2 | 2 |
| **relops-3/decision** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 42 | 42 |
| **relops-3/images** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 2 | 2 |
| **reviewer-assignment/bot** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 15 | 15 |
| **scriptworker-1/decision** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 21 | 21 |
| **scriptworker-1/images** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 133 | 133 |
| **scriptworker-3/decision** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 22 | 22 |
| **scriptworker-3/images** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 177 | 177 |
| **taskgraph-1/decision** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 22 | 22 |
| **taskgraph-1/images** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 39 | 39 |
| **taskgraph-3/decision** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 7 | 7 |
| **taskgraph-3/images** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 11 | 11 |
| **taskgraph-t/linux-docker** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 243 | 243 |
| **taskgraph-t/win11-64-24h2** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | windows | amd64 | 1.27.1 | 25 | 25 |
| **taskgraph-t/win11-64-25h2** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | windows | amd64 | 1.27.1 | 2 | 2 |
| **translations-1/b-linux-large-gcp-1tb-32-256-d2g** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 72 | 72 |
| **translations-1/b-linux-large-gcp-1tb-32-256-std-d2g** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 30 | 30 |
| **translations-1/b-linux-large-gcp-1tb-64-512-d2g** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 2 | 2 |
| **translations-1/b-linux-large-gcp-1tb-64-512-std-d2g** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 3 | 3 |
| **translations-1/b-linux-large-gcp-d2g** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 321 | 321 |
| **translations-1/b-linux-large-gcp-d2g-1tb** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 2 | 2 |
| **translations-1/b-linux-large-gcp-d2g-1tb-standard** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 2 | 2 |
| **translations-1/b-linux-large-gcp-d2g-300gb** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 436 | 436 |
| **translations-1/b-linux-v100-gpu-d2g** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 333 | 333 |
| **translations-1/b-linux-v100-gpu-d2g-4** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 265 | 265 |
| **translations-1/b-linux-v100-gpu-d2g-4-1tb** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 120 | 120 |
| **translations-1/b-linux-v100-gpu-d2g-4-1tb-alpha** | generic-worker | 84.1.2 | multiuser | 4a139f395a | linux | amd64 | 1.24.4 | 2 | 2 |
| **translations-1/b-linux-v100-gpu-d2g-4-1tb-standard** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 4 | 4 |
| **translations-1/b-linux-v100-gpu-d2g-4-1tb-std-alpha** | generic-worker | 84.1.2 | multiuser | 4a139f395a | linux | amd64 | 1.24.4 | 2 | 2 |
| **translations-1/b-linux-v100-gpu-d2g-4-2tb** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 15 | 15 |
| **translations-1/b-linux-v100-gpu-d2g-4-2tb-alpha** | generic-worker | 84.1.2 | multiuser | 4a139f395a | linux | amd64 | 1.24.4 | 2 | 2 |
| **translations-1/b-linux-v100-gpu-d2g-4-300gb** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 118 | 118 |
| **translations-1/b-linux-v100-gpu-d2g-4-300gb-alpha** | generic-worker | 84.1.2 | multiuser | 4a139f395a | linux | amd64 | 1.24.4 | 3 | 3 |
| **translations-1/b-linux-v100-gpu-d2g-4-300gb-standard** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 2 | 2 |
| **translations-1/b-linux-v100-gpu-d2g-4-300gb-std-alpha** | generic-worker | 84.1.2 | multiuser | 4a139f395a | linux | amd64 | 1.24.4 | 2 | 2 |
| **translations-1/b-linux-v100-gpu-d2g-4-alpha** | generic-worker | 84.1.2 | multiuser | 4a139f395a | linux | amd64 | 1.24.4 | 2 | 2 |
| **translations-1/b-linux-v100-gpu-d2g-alpha** | generic-worker | 84.1.2 | multiuser | 4a139f395a | linux | amd64 | 1.24.4 | 2 | 2 |
| **translations-1/decision** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 29 | 29 |
| **translations-1/images** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 12 | 12 |
| **translations-t/t-linux-gw-noscratch-amd** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 2 | 2 |
| **xpi-1/b-linux** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 222 | 222 |
| **xpi-1/decision** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 86 | 86 |
| **xpi-1/images** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 2 | 2 |
| **xpi-3/b-linux** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 4 | 4 |
| **xpi-3/decision** | generic-worker | 110.1.0 | multiuser | 8522ee4ea6 | linux | amd64 | 1.27.1 | 8 | 8 |
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
| projects/taskcluster-imaging/global/images/docker-firefoxci-gcp-l1-googlecompute-2025-06-13t18-31-38z | 1 |
| ami-03e4f8db63254ce7e,ami-0a6e926238859761c,ami-0b5dd0bbb670ec80e | 1 |


| Worker Pool | Implementation | Version | Total Workers | Total Capacity |
| --- | --- | --- | ---: | ---: |
| **infra/build-decision** | docker-worker | 38.0.5 | 6 | 48 |
| **proj-fuzzing/bugmon-pernosco** | docker-worker | 44.23.4 | 2 | 2 |


## Script Worker

Total: `52`



| Worker Pool | Implementation | Version | Total Workers | Total Capacity |
| --- | --- | --- | ---: | ---: |
| **scriptworker-k8s/app-services-3-beetmover** | Scriptworker | <no value> | 0 | 0 |
| **scriptworker-k8s/app-services-3-shipit** | Scriptworker | <no value> | 0 | 0 |
| **scriptworker-k8s/app-services-3-signing** | Scriptworker | <no value> | 0 | 0 |
| **scriptworker-k8s/comm-3-balrog** | Scriptworker | <no value> | 0 | 0 |
| **scriptworker-k8s/comm-3-beetmover** | Scriptworker | <no value> | 0 | 0 |
| **scriptworker-k8s/comm-3-bouncer** | Scriptworker | <no value> | 0 | 0 |
| **scriptworker-k8s/comm-3-lando** | Scriptworker | <no value> | 0 | 0 |
| **scriptworker-k8s/comm-3-shipit** | Scriptworker | <no value> | 0 | 0 |
| **scriptworker-k8s/comm-3-signing** | Scriptworker | <no value> | 0 | 0 |
| **scriptworker-k8s/comm-t-signing** | Scriptworker | <no value> | 0 | 0 |
| **scriptworker-k8s/enterprise-3-signing** | Scriptworker | <no value> | 0 | 0 |
| **scriptworker-k8s/enterprise-t-signing** | Scriptworker | <no value> | 0 | 0 |
| **scriptworker-k8s/gecko-1-balrog** | Scriptworker | <no value> | 0 | 0 |
| **scriptworker-k8s/gecko-1-balrog-dev** | Scriptworker | <no value> | 0 | 0 |
| **scriptworker-k8s/gecko-1-beetmover** | Scriptworker | <no value> | 0 | 0 |
| **scriptworker-k8s/gecko-1-beetmover-dev** | Scriptworker | <no value> | 0 | 0 |
| **scriptworker-k8s/gecko-1-bouncer** | Scriptworker | <no value> | 0 | 0 |
| **scriptworker-k8s/gecko-1-bouncer-dev** | Scriptworker | <no value> | 0 | 0 |
| **scriptworker-k8s/gecko-1-lando** | Scriptworker | <no value> | 0 | 0 |
| **scriptworker-k8s/gecko-3-addon** | Scriptworker | <no value> | 0 | 0 |
| **scriptworker-k8s/gecko-3-balrog** | Scriptworker | <no value> | 0 | 0 |
| **scriptworker-k8s/gecko-3-beetmover** | Scriptworker | <no value> | 0 | 0 |
| **scriptworker-k8s/gecko-3-bouncer** | Scriptworker | <no value> | 0 | 0 |
| **scriptworker-k8s/gecko-3-lando** | Scriptworker | <no value> | 0 | 0 |
| **scriptworker-k8s/gecko-3-pushapk** | Scriptworker | <no value> | 0 | 0 |
| **scriptworker-k8s/gecko-3-pushmsix** | Scriptworker | <no value> | 0 | 0 |
| **scriptworker-k8s/gecko-3-shipit** | Scriptworker | <no value> | 0 | 0 |
| **scriptworker-k8s/gecko-3-signing** | Scriptworker | <no value> | 0 | 0 |
| **scriptworker-k8s/gecko-t-signing** | Scriptworker | <no value> | 0 | 0 |
| **scriptworker-k8s/gecko-t-signing-dev** | Scriptworker | <no value> | 0 | 0 |
| **scriptworker-k8s/mobile-3-bitrise** | Scriptworker | <no value> | 0 | 0 |
| **scriptworker-k8s/mobile-3-pushapk** | Scriptworker | <no value> | 0 | 0 |
| **scriptworker-k8s/mobile-3-signing** | Scriptworker | <no value> | 0 | 0 |
| **scriptworker-k8s/mobile-t-signing** | Scriptworker | <no value> | 0 | 0 |
| **scriptworker-k8s/mozillavpn-3-signing** | Scriptworker | <no value> | 0 | 0 |
| **scriptworker-k8s/mozillavpn-t-signing** | Scriptworker | <no value> | 0 | 0 |
| **scriptworker-k8s/mozillavpn-t-signing-dev** | Scriptworker | <no value> | 0 | 0 |
| **scriptworker-k8s/translations-1-beetmover** | Scriptworker | <no value> | 0 | 0 |
| **scriptworker-k8s/xpi-3-github** | Scriptworker | <no value> | 0 | 0 |
| **scriptworker-k8s/xpi-3-shipit** | Scriptworker | <no value> | 0 | 0 |
| **scriptworker-k8s/xpi-3-signing** | Scriptworker | <no value> | 0 | 0 |
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
| **gecko-t/t-linux-vm-2204-wayland** |  | No artifacts found | 63 | 63 |
| **gecko-t/t-linux-vm-2204-wayland-snap** |  | No artifacts found | 2 | 2 |


## Version not determined [^2]

Total: `50`


Count by image:

| Version | Count |
| :--- | ---: |
|  | 5 |
| projects/taskcluster-imaging/global/images/gw-fxci-gcp-l1-2404-amd64-googlecompute-alpha | 1 |
| projects/taskcluster-imaging/global/images/generic-worker-ubuntu-24-04-arm64-mqscbutmsbawtytlroxk | 1 |
| ami-02619e55246806e8d | 1 |
| unknown | 35 |
| projects/taskcluster-imaging/global/images/generic-worker-ubuntu-24-04-9ed5a812d3d844ac9e4e | 7 |


| Worker Pool | Implementation | Version | Total Workers | Total Capacity |
| --- | --- | --- | ---: | ---: |
| **comm-t/win11-a64-25h2-alpha** |  | Version not determined; task not (yet) claimed | 60 | 60 |
| **enterprise-t/win11-64-24h2-alpha** |  | Version not determined; task not (yet) claimed | 45 | 45 |
| **enterprise-t/win11-64-24h2-source-alpha** |  | Version not determined; task not (yet) claimed | 45 | 45 |
| **enterprise-t/win11-64-25h2-alpha** |  | Version not determined; task not (yet) claimed | 45 | 45 |
| **enterprise-t/win11-64-25h2-source-alpha** |  | Version not determined; task not (yet) claimed | 45 | 45 |
| **gecko-1/b-win2022-alpha** |  | Version not determined; task not (yet) claimed | 52 | 52 |
| **gecko-1/b-win2025-alpha** |  | Version not determined; task not (yet) claimed | 45 | 45 |
| **gecko-1/win11-a64-25h2-builder-alpha** |  | Version not determined; task not (yet) claimed | 47 | 47 |
| **gecko-t/t-linux-2404-relsre** |  | Version not determined; task not (yet) claimed | 4 | 4 |
| **gecko-t/win10-64-2009-alpha** |  | Version not determined; task not (yet) claimed | 188 | 188 |
| **gecko-t/win10-64-2009-gpu-alpha** |  | Version not determined; task not (yet) claimed | 45 | 45 |
| **gecko-t/win10-64-2009-source-alpha** |  | Version not determined; task not (yet) claimed | 51 | 51 |
| **gecko-t/win10-64-2009-webgpu-alpha** |  | Version not determined; task not (yet) claimed | 51 | 51 |
| **gecko-t/win11-64-24h2-alpha** |  | Version not determined; task not (yet) claimed | 479 | 479 |
| **gecko-t/win11-64-24h2-gpu-alpha** |  | Version not determined; task not (yet) claimed | 92 | 92 |
| **gecko-t/win11-64-24h2-large-alpha** |  | Version not determined; task not (yet) claimed | 46 | 46 |
| **gecko-t/win11-64-24h2-privileged-alpha** |  | Version not determined; task not (yet) claimed | 45 | 45 |
| **gecko-t/win11-64-24h2-source-alpha** |  | Version not determined; task not (yet) claimed | 58 | 58 |
| **gecko-t/win11-64-24h2-unprivileged-alpha** |  | Version not determined; task not (yet) claimed | 46 | 46 |
| **gecko-t/win11-64-24h2-webgpu-alpha** |  | Version not determined; task not (yet) claimed | 45 | 45 |
| **gecko-t/win11-64-25h2-alpha** |  | Version not determined; task not (yet) claimed | 765 | 765 |
| **gecko-t/win11-64-25h2-amd** |  | Version not determined; task not (yet) claimed | 45 | 45 |
| **gecko-t/win11-64-25h2-gpu-alpha** |  | Version not determined; task not (yet) claimed | 154 | 154 |
| **gecko-t/win11-64-25h2-large-alpha** |  | Version not determined; task not (yet) claimed | 256 | 256 |
| **gecko-t/win11-64-25h2-privileged-alpha** |  | Version not determined; task not (yet) claimed | 45 | 45 |
| **gecko-t/win11-64-25h2-source-alpha** |  | Version not determined; task not (yet) claimed | 85 | 85 |
| **gecko-t/win11-64-25h2-unprivileged-alpha** |  | Version not determined; task not (yet) claimed | 46 | 46 |
| **gecko-t/win11-64-25h2-webgpu-alpha** |  | Version not determined; task not (yet) claimed | 43 | 43 |
| **gecko-t/win11-a64-25h2-alpha** |  | Version not determined; task not (yet) claimed | 76 | 76 |
| **mozillavpn-1/b-win2022-alpha** |  | Version not determined; task not (yet) claimed | 44 | 44 |
| **mozillavpn-1/b-win2025-alpha** |  | Version not determined; task not (yet) claimed | 43 | 43 |
| **nss-1/b-win2022-alpha** |  | Version not determined; task not (yet) claimed | 43 | 43 |
| **nss-1/b-win2025-alpha** |  | Version not determined; task not (yet) claimed | 43 | 43 |
| **nss-1/win11-a64-25h2-alpha** |  | Version not determined; task not (yet) claimed | 56 | 56 |
| **nss-1/win11-a64-25h2-builder-alpha** |  | Version not determined; task not (yet) claimed | 44 | 44 |
| **nss-t/win11-a64-25h2-alpha** |  | Version not determined; task not (yet) claimed | 48 | 48 |
| **proj-autophone/gecko-t-lambda-perf-a55** |  | Version not determined; task not (yet) claimed | 0 | 0 |
| **proj-fuzzing/bugmon-monitor** |  | Version not determined; task not (yet) claimed | 94 | 94 |
| **proj-fuzzing/bugmon-pernosco-staging** |  | Version not determined; task not (yet) claimed | 3 | 3 |
| **proj-fuzzing/bugmon-processor** |  | Version not determined; task not (yet) claimed | 3 | 3 |
| **proj-fuzzing/ci** |  | Version not determined; task not (yet) claimed | 101 | 101 |
| **proj-fuzzing/ci-arm64** |  | Version not determined; task not (yet) claimed | 3 | 3 |
| **proj-fuzzing/decision** |  | Version not determined; task not (yet) claimed | 228 | 228 |
| **proj-fuzzing/grizzly-reduce-worker** |  | Version not determined; task not (yet) claimed | 3 | 3 |
| **proj-fuzzing/grizzly-reduce-worker-android** |  | Version not determined; task not (yet) claimed | 3 | 3 |
| **proj-fuzzing/nss-corpus-update-worker** |  | Version not determined; task not (yet) claimed | 3 | 3 |
| **releng-hardware/gecko-t-osx-1500-m4** |  | Version not determined; task not (yet) claimed | 0 | 0 |
| **releng-hardware/gecko-t-osx-2600-m4** |  | Version not determined; task not (yet) claimed | 0 | 0 |
| **releng-hardware/nss-3-b-osx-1015** |  | Version not determined; task not (yet) claimed | 0 | 0 |
| **releng-hardware/win11-64-24h2-hw** |  | Version not determined; task not (yet) claimed | 0 | 0 |



[^1]: Those are the pools whose tasks were claimed and resolved by a worker as expected, but the worker did not publish either artifact `public/logs/live_backing.log` nor `public/logs/chain_of_trust.log`, which is the source used to identify the worker implementation.

[^2]: Probing task remains pending after two hours. Those are the pools that were not able to start any worker to claim the task within two hours.

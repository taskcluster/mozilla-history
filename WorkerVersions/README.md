

# Worker Pool Versions


## Generic Worker

Total: `347`

Count by version:

| Version | Count |
| :--- | ---: |
| 36.0.0 | 4 |
| 45.0.0 | 1 |
| 60.3.4 | 18 |
| 61.0.0 | 1 |
| 64.3.0 | 26 |
| 84.1.2 | 7 |
| 86.0.2 | 3 |
| 87.0.0 | 1 |
| 88.0.2 | 2 |
| 88.1.3 | 47 |
| 91.0.2 | 2 |
| 91.1.0 | 4 |
| 91.1.1 | 2 |
| 93.1.2 | 211 |
| 93.1.4 | 13 |
| 94.0.1 | 5 |


Count by image:

| Version | Count |
| :--- | ---: |
| /subscriptions/108d46d5-fe9b-4850-9a7d-8c914aa6c1f0/resourceGroups/rg-packer-worker-images/providers/Microsoft.Compute/galleries/win11_a64_24h2_tester_alpha/images/win11_a64_24h2_tester_alpha/versions/1.0.0 | 4 |
| /subscriptions/108d46d5-fe9b-4850-9a7d-8c914aa6c1f0/resourceGroups/rg-packer-worker-images/providers/Microsoft.Compute/galleries/win11_a64_24h2_tester/images/win11_a64_24h2_tester/versions/1.0.8 | 4 |
| projects/fxci-production-level3-workers/global/images/gw-fxci-gcp-l3-arm64-gui-googlecompute-2024-09-18t19-02-58z | 4 |
| projects/fxci-production-level3-workers/global/images/gw-fxci-gcp-l3-2404-arm64-headless-googlecompute-2025-11-14 | 6 |
| projects/taskcluster-imaging/global/images/gw-fxci-gcp-l1-2404-arm64-headless-googlecompute-alpha | 1 |
| projects/taskcluster-imaging/global/images/gw-fxci-gcp-l1-2404-amd64-gui-googlecompute-2025-11-14 | 4 |
| /subscriptions/108d46d5-fe9b-4850-9a7d-8c914aa6c1f0/resourceGroups/rg-packer-worker-images/providers/Microsoft.Compute/galleries/win11_a64_24h2_builder_alpha/images/win11_a64_24h2_builder_alpha/versions/1.0.0 | 2 |
| /subscriptions/a30e97ab-734a-4f3b-a0e4-c51c0bff0701/resourceGroups/rg-packer-worker-images/providers/Microsoft.Compute/galleries/trusted_win11_a64_24h2_builder/images/trusted_win11_a64_24h2_builder/versions/1.0.8 | 3 |
| projects/taskcluster-imaging/global/images/gw-fxci-gcp-l1-2404-amd64-headless-googlecompute-alpha | 6 |
| unknown | 53 |
|  | 37 |
| projects/fxci-production-level3-workers/global/images/gw-fxci-gcp-l3-gui-googlecompute-2024-09-18t05-46-31z | 3 |
| projects/taskcluster-imaging/global/images/gw-fxci-gcp-l1-2404-amd64-headless-googlecompute-alpha-tc | 7 |
| /subscriptions/108d46d5-fe9b-4850-9a7d-8c914aa6c1f0/resourceGroups/rg-packer-worker-images/providers/Microsoft.Compute/galleries/win11_a64_24h2_builder/images/win11_a64_24h2_builder/versions/1.0.8 | 3 |
| projects/taskcluster-imaging/global/images/gw-fxci-gcp-l1-arm64-gui-googlecompute-2024-09-18t14-57-52z | 9 |
| projects/taskcluster-imaging/global/images/gw-fxci-gcp-l1-2404-amd64-headless-googlecompute-2025-11-14 | 117 |
| projects/fxci-production-level3-workers/global/images/gw-fxci-gcp-l3-2404-amd64-headless-googlecompute-2025-11-14 | 58 |
| projects/taskcluster-imaging/global/images/gw-fxci-gcp-l1-2404-arm64-headless-googlecompute-2025-11-14 | 14 |
| projects/taskcluster-imaging/global/images/gw-fxci-gcp-l1-2404-amd64-gui-googlecompute-alpha | 1 |
| projects/fxci-production-level3-workers/global/images/gw-fxci-gcp-l3-2404-amd64-headless-googlecompute-alpha | 1 |
| projects/taskcluster-imaging/global/images/gw-fxci-gcp-l1-gui-googlecompute-2025-01-13t22-33-40z | 2 |
| projects/taskcluster-imaging/global/images/gw-fxci-gcp-l1-gui-googlecompute-2024-08-22t22-48-09z | 8 |


| Worker Pool | Implementation | Version | Engine | Revision | OS | Arch | GO | Total Workers | Total Capacity |
| --- | --- | --- | --- | --- | --- | --- | --- | ---: | ---: |
| **adhoc-1/decision** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 1 | 1 |
| **adhoc-1/images** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 2 | 2 |
| **adhoc-3/decision** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 1 | 1 |
| **adhoc-3/images** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 1 | 1 |
| **app-services-1/b-linux** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 253 | 253 |
| **app-services-1/b-linux-docker-amd** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 1 | 1 |
| **app-services-1/decision** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 1182 | 1182 |
| **app-services-1/images** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 3 | 3 |
| **app-services-3/b-linux** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 627 | 627 |
| **app-services-3/b-linux-docker-amd** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 1 | 1 |
| **app-services-3/decision** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 23 | 23 |
| **app-services-3/images** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 2 | 2 |
| **code-analysis-1/linux-docker** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 1 | 1 |
| **code-analysis-1/linux-gw-gcp** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | amd64 | 1.22.2 | 9 | 9 |
| **code-analysis-3/linux-docker** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 1 | 1 |
| **code-analysis-3/linux-gw-gcp** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | amd64 | 1.22.2 | 1 | 1 |
| **code-coverage/bot** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 215 | 215 |
| **code-review/bot** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 7217 | 7217 |
| **comm-1/b-linux** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 1 | 1 |
| **comm-1/b-linux-aarch64** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | arm64 | 1.25.4 | 1 | 1 |
| **comm-1/b-linux-docker-amd** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 19 | 19 |
| **comm-1/b-linux-docker-large-amd** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 2 | 2 |
| **comm-1/b-linux-docker-xlarge-amd** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 4 | 4 |
| **comm-1/b-linux-gcp-aarch64** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | arm64 | 1.22.2 | 1 | 1 |
| **comm-1/b-linux-large** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 1 | 1 |
| **comm-1/b-linux-xlarge** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 1 | 1 |
| **comm-1/b-win2022** | generic-worker | 88.1.3 | multiuser | ffc99d77c0 | windows | amd64 | 1.25.1 | 1 | 1 |
| **comm-1/decision** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 1 | 1 |
| **comm-1/images** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 1 | 1 |
| **comm-1/images-aarch64** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | arm64 | 1.25.4 | 1 | 1 |
| **comm-1/images-gcp-aarch64** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | arm64 | 1.22.2 | 1 | 1 |
| **comm-2/b-linux** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 1 | 1 |
| **comm-2/b-linux-aarch64** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | arm64 | 1.25.4 | 1 | 1 |
| **comm-2/b-linux-docker-amd** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 1 | 1 |
| **comm-2/b-linux-docker-large-amd** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 1 | 1 |
| **comm-2/b-linux-docker-xlarge-amd** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 1 | 1 |
| **comm-2/b-linux-gcp-aarch64** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | arm64 | 1.22.2 | 1 | 1 |
| **comm-2/b-linux-large** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 1 | 1 |
| **comm-2/b-linux-xlarge** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 1 | 1 |
| **comm-2/decision** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 1 | 1 |
| **comm-2/images** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 1 | 1 |
| **comm-2/images-aarch64** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | arm64 | 1.25.4 | 1 | 1 |
| **comm-2/images-gcp-aarch64** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | arm64 | 1.22.2 | 1 | 1 |
| **comm-3/b-linux** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 2 | 2 |
| **comm-3/b-linux-aarch64** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | arm64 | 1.25.4 | 1 | 1 |
| **comm-3/b-linux-docker-amd** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 1830 | 1830 |
| **comm-3/b-linux-docker-large-amd** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 7 | 7 |
| **comm-3/b-linux-docker-xlarge-amd** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 8 | 8 |
| **comm-3/b-linux-gcp-aarch64** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | arm64 | 1.22.2 | 1 | 1 |
| **comm-3/b-linux-large** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 1 | 1 |
| **comm-3/b-linux-xlarge** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 1 | 1 |
| **comm-3/b-win2022** | generic-worker | 88.1.3 | multiuser | ffc99d77c0 | windows | amd64 | 1.25.1 | 580 | 580 |
| **comm-3/decision** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 1 | 1 |
| **comm-3/images** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 1 | 1 |
| **comm-3/images-aarch64** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | arm64 | 1.25.4 | 1 | 1 |
| **comm-3/images-gcp-aarch64** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | arm64 | 1.22.2 | 1 | 1 |
| **comm-t/misc** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 1 | 1 |
| **comm-t/t-linux-arm64-docker** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | arm64 | 1.25.4 | 1 | 1 |
| **comm-t/t-linux-docker** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 17 | 17 |
| **comm-t/t-linux-docker-amd** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 1174 | 1174 |
| **comm-t/t-linux-docker-noscratch** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 57 | 57 |
| **comm-t/t-linux-docker-noscratch-amd** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 3106 | 3106 |
| **comm-t/win11-64-2009** | generic-worker | 88.1.3 | multiuser | ffc99d77c0 | windows | amd64 | 1.25.1 | 1 | 1 |
| **comm-t/win11-64-2009-ssd** | generic-worker | 88.1.3 | multiuser | ffc99d77c0 | windows | amd64 | 1.25.1 | 1 | 1 |
| **comm-t/win11-64-24h2** | generic-worker | 88.1.3 | multiuser | ffc99d77c0 | windows | amd64 | 1.25.1 | 2993 | 2993 |
| **comm-t/win11-64-24h2-ssd** | generic-worker | 88.1.3 | multiuser | ffc99d77c0 | windows | amd64 | 1.25.1 | 384 | 384 |
| **comm-t/win11-a64-24h2** | generic-worker | 88.1.3 | multiuser | ffc99d77c0 | windows | amd64 | 1.25.1 | 1 | 1 |
| **comm-t/win11-a64-24h2-alpha** | generic-worker | 93.1.4 | multiuser | 1088d85c3b | windows | arm64 | 1.25.4 | 1 | 1 |
| **enterprise-1/b-linux** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 258 | 258 |
| **enterprise-1/b-linux-aarch64** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | arm64 | 1.25.4 | 1 | 1 |
| **enterprise-1/b-linux-docker-amd** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 256 | 256 |
| **enterprise-1/b-linux-docker-large-amd** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 11 | 11 |
| **enterprise-1/b-linux-docker-xlarge-amd** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 158 | 158 |
| **enterprise-1/b-win2022** | generic-worker | 88.1.3 | multiuser | ffc99d77c0 | windows | amd64 | 1.25.1 | 82 | 82 |
| **enterprise-1/decision** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 1181 | 1181 |
| **enterprise-1/images** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 1 | 1 |
| **enterprise-1/images-aarch64** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | arm64 | 1.25.4 | 1 | 1 |
| **enterprise-3/b-linux** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 83 | 83 |
| **enterprise-3/b-linux-aarch64** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | arm64 | 1.25.4 | 3 | 3 |
| **enterprise-3/b-linux-docker-amd** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 267 | 267 |
| **enterprise-3/b-linux-docker-large-amd** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 17 | 17 |
| **enterprise-3/b-linux-docker-xlarge-amd** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 57 | 57 |
| **enterprise-3/b-win2022** | generic-worker | 88.1.3 | multiuser | ffc99d77c0 | windows | amd64 | 1.25.1 | 36 | 36 |
| **enterprise-3/decision** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 7 | 7 |
| **enterprise-3/images** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 1 | 1 |
| **enterprise-3/images-aarch64** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | arm64 | 1.25.4 | 1 | 1 |
| **enterprise-t/misc** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 120 | 120 |
| **enterprise-t/t-linux-docker-amd** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 1085 | 1085 |
| **gecko-1/b-linux** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 6003 | 6003 |
| **gecko-1/b-linux-2204-kvm-gcp** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | amd64 | 1.22.2 | 1 | 1 |
| **gecko-1/b-linux-aarch64** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | arm64 | 1.25.4 | 15 | 15 |
| **gecko-1/b-linux-docker-alpha** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 1 | 1 |
| **gecko-1/b-linux-docker-amd** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 6460 | 6460 |
| **gecko-1/b-linux-docker-large-amd** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 71 | 71 |
| **gecko-1/b-linux-docker-xlarge-amd** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 714 | 714 |
| **gecko-1/b-linux-gcp-aarch64** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | arm64 | 1.22.2 | 1 | 1 |
| **gecko-1/b-linux-kvm** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 186 | 186 |
| **gecko-1/b-linux-large** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 1 | 1 |
| **gecko-1/b-linux-medium** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 7957 | 7957 |
| **gecko-1/b-linux-xlarge** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 1 | 1 |
| **gecko-1/b-win2022** | generic-worker | 88.1.3 | multiuser | ffc99d77c0 | windows | amd64 | 1.25.1 | 544 | 544 |
| **gecko-1/b-win2022-alpha** | generic-worker | 93.1.2 | multiuser | e954a7555f | windows | amd64 | 1.25.4 | 1 | 1 |
| **gecko-1/b-win2022-xxlarge** | generic-worker | 88.1.3 | multiuser | ffc99d77c0 | windows | amd64 | 1.25.1 | 7 | 7 |
| **gecko-1/decision** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 2177 | 2177 |
| **gecko-1/images** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 12 | 12 |
| **gecko-1/images-aarch64** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | arm64 | 1.25.4 | 3 | 3 |
| **gecko-1/images-gcp-aarch64** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | arm64 | 1.22.2 | 1 | 1 |
| **gecko-1/win11-a64-24h2-builder** | generic-worker | 88.1.3 | multiuser | ffc99d77c0 | windows | amd64 | 1.25.1 | 18 | 18 |
| **gecko-1/win11-a64-24h2-builder-alpha** | generic-worker | 93.1.4 | multiuser | 1088d85c3b | windows | arm64 | 1.25.4 | 1 | 1 |
| **gecko-2/b-linux** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 257 | 257 |
| **gecko-2/b-linux-aarch64** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | arm64 | 1.25.4 | 1 | 1 |
| **gecko-2/b-linux-docker-amd** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 445 | 445 |
| **gecko-2/b-linux-docker-large-amd** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 19 | 19 |
| **gecko-2/b-linux-docker-xlarge-amd** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 80 | 80 |
| **gecko-2/b-linux-gcp-aarch64** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | arm64 | 1.22.2 | 1 | 1 |
| **gecko-2/b-linux-kvm** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 20 | 20 |
| **gecko-2/b-linux-large** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 1 | 1 |
| **gecko-2/b-linux-medium** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 311 | 311 |
| **gecko-2/b-linux-xlarge** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 1 | 1 |
| **gecko-2/b-win2022** | generic-worker | 88.1.3 | multiuser | ffc99d77c0 | windows | amd64 | 1.25.1 | 15 | 15 |
| **gecko-2/decision** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 19 | 19 |
| **gecko-2/images** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 1 | 1 |
| **gecko-2/images-aarch64** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | arm64 | 1.25.4 | 1 | 1 |
| **gecko-2/images-gcp-aarch64** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | arm64 | 1.22.2 | 1 | 1 |
| **gecko-3/b-linux** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 12662 | 12662 |
| **gecko-3/b-linux-2204-kvm-gcp** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | amd64 | 1.22.2 | 1 | 1 |
| **gecko-3/b-linux-aarch64** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | arm64 | 1.25.4 | 7 | 7 |
| **gecko-3/b-linux-docker-alpha** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 1 | 1 |
| **gecko-3/b-linux-docker-amd** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 15032 | 15032 |
| **gecko-3/b-linux-docker-large-amd** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 166 | 166 |
| **gecko-3/b-linux-docker-xlarge-amd** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 1942 | 1942 |
| **gecko-3/b-linux-gcp-aarch64** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | arm64 | 1.22.2 | 1 | 1 |
| **gecko-3/b-linux-kvm** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 369 | 369 |
| **gecko-3/b-linux-large** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 1 | 1 |
| **gecko-3/b-linux-medium** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 16041 | 16041 |
| **gecko-3/b-linux-xlarge** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 40 | 40 |
| **gecko-3/b-win2022** | generic-worker | 88.1.3 | multiuser | ffc99d77c0 | windows | amd64 | 1.25.1 | 4780 | 4780 |
| **gecko-3/b-win2022-xxlarge** | generic-worker | 88.1.3 | multiuser | ffc99d77c0 | windows | amd64 | 1.25.1 | 11 | 11 |
| **gecko-3/decision** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 3771 | 3771 |
| **gecko-3/images** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 82 | 82 |
| **gecko-3/images-aarch64** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | arm64 | 1.25.4 | 10 | 10 |
| **gecko-3/images-gcp-aarch64** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | arm64 | 1.22.2 | 1 | 1 |
| **gecko-3/win11-a64-24h2-builder** | generic-worker | 88.1.3 | multiuser | ffc99d77c0 | windows | amd64 | 1.25.1 | 41 | 41 |
| **gecko-t/misc** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 589 | 589 |
| **gecko-t/t-linux-2204-wayland** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | amd64 | 1.22.2 | 8251 | 8251 |
| **gecko-t/t-linux-2204-wayland-arm64-relsre** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | arm64 | 1.22.2 | 1 | 1 |
| **gecko-t/t-linux-2204-wayland-relsre** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | amd64 | 1.22.2 | 1 | 1 |
| **gecko-t/t-linux-2204-wayland-root-exp** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | amd64 | 1.22.2 | 1 | 1 |
| **gecko-t/t-linux-2204-wayland-snap** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | amd64 | 1.22.2 | 124 | 124 |
| **gecko-t/t-linux-2404-headless-alpha** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 1 | 1 |
| **gecko-t/t-linux-2404-headless-arm64-alpha** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | arm64 | 1.25.4 | 1 | 1 |
| **gecko-t/t-linux-2404-headless-ssd-alpha** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 1 | 1 |
| **gecko-t/t-linux-2404-wayland** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 1 | 1 |
| **gecko-t/t-linux-2404-wayland-relsre** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 1 | 1 |
| **gecko-t/t-linux-2404-wayland-snap** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 125 | 125 |
| **gecko-t/t-linux-arm64-docker** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | arm64 | 1.25.4 | 22 | 22 |
| **gecko-t/t-linux-docker** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 1389 | 1389 |
| **gecko-t/t-linux-docker-16c32gb-amd** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 3007 | 3007 |
| **gecko-t/t-linux-docker-amd** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 21791 | 21791 |
| **gecko-t/t-linux-docker-amd-alpha** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 1 | 1 |
| **gecko-t/t-linux-docker-kvm** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 61938 | 61938 |
| **gecko-t/t-linux-docker-kvm-alpha** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 1 | 1 |
| **gecko-t/t-linux-docker-noscratch** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 3630 | 3630 |
| **gecko-t/t-linux-docker-noscratch-amd** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 81865 | 81865 |
| **gecko-t/t-linux-docker-noscratch-amd-alpha** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 1 | 1 |
| **gecko-t/t-linux-xlarge-2204-wayland** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | amd64 | 1.22.2 | 9915 | 9915 |
| **gecko-t/t-linux-xlarge-2404-wayland** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 1 | 1 |
| **gecko-t/win10-64-2009** | generic-worker | 88.1.3 | multiuser | ffc99d77c0 | windows | amd64 | 1.25.1 | 3563 | 3563 |
| **gecko-t/win10-64-2009-alpha** | generic-worker | 91.1.0 | multiuser | b8f34332c3 | windows | amd64 | 1.25.3 | 1 | 1 |
| **gecko-t/win10-64-2009-gpu** | generic-worker | 88.1.3 | multiuser | ffc99d77c0 | windows | amd64 | 1.25.1 | 404 | 404 |
| **gecko-t/win10-64-2009-gpu-alpha** | generic-worker | 91.1.0 | multiuser | b8f34332c3 | windows | amd64 | 1.25.3 | 1 | 1 |
| **gecko-t/win10-64-2009-source** | generic-worker | 88.1.3 | multiuser | ffc99d77c0 | windows | amd64 | 1.25.1 | 16 | 16 |
| **gecko-t/win10-64-2009-source-alpha** | generic-worker | 91.1.0 | multiuser | b8f34332c3 | windows | amd64 | 1.25.3 | 2 | 2 |
| **gecko-t/win10-64-2009-webgpu** | generic-worker | 88.1.3 | multiuser | ffc99d77c0 | windows | amd64 | 1.25.1 | 1 | 1 |
| **gecko-t/win10-64-2009-webgpu-alpha** | generic-worker | 91.1.0 | multiuser | b8f34332c3 | windows | amd64 | 1.25.3 | 1 | 1 |
| **gecko-t/win11-64-2009** | generic-worker | 88.1.3 | multiuser | ffc99d77c0 | windows | amd64 | 1.25.1 | 782 | 782 |
| **gecko-t/win11-64-2009-alpha** | generic-worker | 93.1.4 | multiuser | 1088d85c3b | windows | amd64 | 1.25.4 | 1 | 1 |
| **gecko-t/win11-64-2009-gpu** | generic-worker | 88.1.3 | multiuser | ffc99d77c0 | windows | amd64 | 1.25.1 | 469 | 469 |
| **gecko-t/win11-64-2009-gpu-alpha** | generic-worker | 93.1.4 | multiuser | 1088d85c3b | windows | amd64 | 1.25.4 | 1 | 1 |
| **gecko-t/win11-64-2009-source** | generic-worker | 88.1.3 | multiuser | ffc99d77c0 | windows | amd64 | 1.25.1 | 1 | 1 |
| **gecko-t/win11-64-2009-source-alpha** | generic-worker | 93.1.4 | multiuser | 1088d85c3b | windows | amd64 | 1.25.4 | 1 | 1 |
| **gecko-t/win11-64-2009-ssd** | generic-worker | 88.1.3 | multiuser | ffc99d77c0 | windows | amd64 | 1.25.1 | 1 | 1 |
| **gecko-t/win11-64-2009-ssd-gpu** | generic-worker | 88.1.3 | multiuser | ffc99d77c0 | windows | amd64 | 1.25.1 | 1 | 1 |
| **gecko-t/win11-64-2009-webgpu** | generic-worker | 88.1.3 | multiuser | ffc99d77c0 | windows | amd64 | 1.25.1 | 1 | 1 |
| **gecko-t/win11-64-2009-webgpu-alpha** | generic-worker | 93.1.4 | multiuser | 1088d85c3b | windows | amd64 | 1.25.4 | 2 | 2 |
| **gecko-t/win11-64-24h2** | generic-worker | 88.1.3 | multiuser | ffc99d77c0 | windows | amd64 | 1.25.1 | 40405 | 40405 |
| **gecko-t/win11-64-24h2-alpha** | generic-worker | 94.0.1 | multiuser | e38b5ecf97 | windows | amd64 | 1.25.4 | 1 | 1 |
| **gecko-t/win11-64-24h2-gpu** | generic-worker | 88.1.3 | multiuser | ffc99d77c0 | windows | amd64 | 1.25.1 | 14907 | 14907 |
| **gecko-t/win11-64-24h2-gpu-alpha** | generic-worker | 94.0.1 | multiuser | e38b5ecf97 | windows | amd64 | 1.25.4 | 1 | 1 |
| **gecko-t/win11-64-24h2-large** | generic-worker | 88.1.3 | multiuser | ffc99d77c0 | windows | amd64 | 1.25.1 | 1561 | 1561 |
| **gecko-t/win11-64-24h2-large-alpha** | generic-worker | 94.0.1 | multiuser | e38b5ecf97 | windows | amd64 | 1.25.4 | 1 | 1 |
| **gecko-t/win11-64-24h2-source** | generic-worker | 88.1.3 | multiuser | ffc99d77c0 | windows | amd64 | 1.25.1 | 1964 | 1964 |
| **gecko-t/win11-64-24h2-source-alpha** | generic-worker | 94.0.1 | multiuser | e38b5ecf97 | windows | amd64 | 1.25.4 | 1 | 1 |
| **gecko-t/win11-64-24h2-ssd** | generic-worker | 88.1.3 | multiuser | ffc99d77c0 | windows | amd64 | 1.25.1 | 1 | 1 |
| **gecko-t/win11-64-24h2-ssd-gpu** | generic-worker | 88.1.3 | multiuser | ffc99d77c0 | windows | amd64 | 1.25.1 | 1 | 1 |
| **gecko-t/win11-64-24h2-webgpu** | generic-worker | 88.1.3 | multiuser | ffc99d77c0 | windows | amd64 | 1.25.1 | 2834 | 2834 |
| **gecko-t/win11-64-24h2-webgpu-alpha** | generic-worker | 94.0.1 | multiuser | e38b5ecf97 | windows | amd64 | 1.25.4 | 1 | 1 |
| **gecko-t/win11-a64-24h2** | generic-worker | 88.1.3 | multiuser | ffc99d77c0 | windows | amd64 | 1.25.1 | 148 | 148 |
| **gecko-t/win11-a64-24h2-alpha** | generic-worker | 93.1.4 | multiuser | 1088d85c3b | windows | arm64 | 1.25.4 | 1 | 1 |
| **glean-1/b-linux** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 15 | 15 |
| **glean-1/decision** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 1177 | 1177 |
| **glean-1/images** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 1 | 1 |
| **glean-3/b-linux** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 5 | 5 |
| **glean-3/decision** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 2 | 2 |
| **glean-3/images** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 1 | 1 |
| **mobile-1/b-linux** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 122 | 122 |
| **mobile-1/decision** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 1159 | 1159 |
| **mobile-1/images** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 1 | 1 |
| **mobile-3/b-linux** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 192 | 192 |
| **mobile-3/decision** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 62 | 62 |
| **mobile-3/images** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 1 | 1 |
| **mozilla-1/b-linux** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 1 | 1 |
| **mozilla-1/b-win2022** | generic-worker | 88.1.3 | multiuser | ffc99d77c0 | windows | amd64 | 1.25.1 | 1 | 1 |
| **mozilla-1/decision** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 1166 | 1166 |
| **mozilla-1/images** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 18 | 18 |
| **mozilla-3/b-linux** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 1 | 1 |
| **mozilla-3/b-win2022** | generic-worker | 88.1.3 | multiuser | ffc99d77c0 | windows | amd64 | 1.25.1 | 1 | 1 |
| **mozilla-3/decision** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 421 | 421 |
| **mozilla-3/images** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 1 | 1 |
| **mozilla-t/t-linux-2204-wayland** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | amd64 | 1.22.2 | 511 | 511 |
| **mozilla-t/t-linux-2204-wayland-relsre** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | amd64 | 1.22.2 | 1 | 1 |
| **mozilla-t/t-linux-2404-wayland** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 1 | 1 |
| **mozilla-t/t-linux-arm64-docker** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | arm64 | 1.25.4 | 1 | 1 |
| **mozilla-t/t-linux-docker** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 37 | 37 |
| **mozilla-t/t-linux-docker-amd** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 1 | 1 |
| **mozilla-t/t-linux-docker-noscratch** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 23 | 23 |
| **mozilla-t/t-linux-docker-noscratch-amd** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 1 | 1 |
| **mozillavpn-1/b-linux** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 320 | 320 |
| **mozillavpn-1/b-linux-gcp-gw** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | amd64 | 1.22.2 | 1 | 1 |
| **mozillavpn-1/b-linux-large** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 223 | 223 |
| **mozillavpn-1/b-win2022** | generic-worker | 88.1.3 | multiuser | ffc99d77c0 | windows | amd64 | 1.25.1 | 105 | 105 |
| **mozillavpn-1/b-win2022-alpha** | generic-worker | 93.1.2 | multiuser | e954a7555f | windows | amd64 | 1.25.4 | 1 | 1 |
| **mozillavpn-1/decision** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 1159 | 1159 |
| **mozillavpn-1/images** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 4 | 4 |
| **mozillavpn-1/win11-a64-24h2-builder** | generic-worker | 88.1.3 | multiuser | ffc99d77c0 | windows | amd64 | 1.25.1 | 6 | 6 |
| **mozillavpn-3/b-linux** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 112 | 112 |
| **mozillavpn-3/b-linux-gcp-gw** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | amd64 | 1.22.2 | 1 | 1 |
| **mozillavpn-3/b-linux-large** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 114 | 114 |
| **mozillavpn-3/b-win2022** | generic-worker | 88.1.3 | multiuser | ffc99d77c0 | windows | amd64 | 1.25.1 | 35 | 35 |
| **mozillavpn-3/decision** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 13 | 13 |
| **mozillavpn-3/images** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 4 | 4 |
| **mozillavpn-3/win11-a64-24h2-builder** | generic-worker | 88.1.3 | multiuser | ffc99d77c0 | windows | amd64 | 1.25.1 | 1 | 1 |
| **nss-1/b-win2022** | generic-worker | 88.1.3 | multiuser | ffc99d77c0 | windows | amd64 | 1.25.1 | 40 | 40 |
| **nss-1/b-win2022-alpha** | generic-worker | 93.1.2 | multiuser | e954a7555f | windows | amd64 | 1.25.4 | 1 | 1 |
| **nss-1/decision** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 2 | 2 |
| **nss-1/images** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 1 | 1 |
| **nss-1/linux-docker** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 74 | 74 |
| **nss-1/win11-a64-24h2** | generic-worker | 88.1.3 | multiuser | ffc99d77c0 | windows | amd64 | 1.25.1 | 1 | 1 |
| **nss-1/win11-a64-24h2-alpha** | generic-worker | 93.1.4 | multiuser | 1088d85c3b | windows | arm64 | 1.25.4 | 1 | 1 |
| **nss-1/win11-a64-24h2-builder** | generic-worker | 88.1.3 | multiuser | ffc99d77c0 | windows | amd64 | 1.25.1 | 1 | 1 |
| **nss-1/win11-a64-24h2-builder-alpha** | generic-worker | 93.1.4 | multiuser | 1088d85c3b | windows | arm64 | 1.25.4 | 1 | 1 |
| **nss-3/b-win2022** | generic-worker | 88.1.3 | multiuser | ffc99d77c0 | windows | amd64 | 1.25.1 | 203 | 203 |
| **nss-3/decision** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 8 | 8 |
| **nss-3/images** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 1 | 1 |
| **nss-3/linux-docker** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 432 | 432 |
| **nss-3/win11-a64-24h2-builder** | generic-worker | 88.1.3 | multiuser | ffc99d77c0 | windows | amd64 | 1.25.1 | 1 | 1 |
| **nss-t/t-linux-arm64-docker** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | arm64 | 1.25.4 | 1 | 1 |
| **nss-t/t-linux-docker** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 577 | 577 |
| **nss-t/t-linux-docker-amd** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 1 | 1 |
| **nss-t/win11-a64-24h2** | generic-worker | 88.1.3 | multiuser | ffc99d77c0 | windows | amd64 | 1.25.1 | 1 | 1 |
| **nss-t/win11-a64-24h2-alpha** | generic-worker | 93.1.4 | multiuser | 1088d85c3b | windows | arm64 | 1.25.4 | 1 | 1 |
| **proj-autophone/gecko-t-bitbar-gw-perf-a55** | generic-worker | 36.0.0 | simple | b9cb11293f | linux | amd64 | 1.13.7 | 0 | 0 |
| **proj-autophone/gecko-t-bitbar-gw-perf-p6** | generic-worker | 36.0.0 | simple | b9cb11293f | linux | amd64 | 1.13.7 | 0 | 0 |
| **proj-autophone/gecko-t-bitbar-gw-perf-s24** | generic-worker | 36.0.0 | simple | b9cb11293f | linux | amd64 | 1.13.7 | 0 | 0 |
| **proj-autophone/gecko-t-bitbar-gw-unit-p5** | generic-worker | 36.0.0 | simple | b9cb11293f | linux | amd64 | 1.13.7 | 0 | 0 |
| **proj-autophone/gecko-t-lambda-perf-a55** | generic-worker | 87.0.0 | insecure | 99a1fcafbb | linux | amd64 | 1.24.5 | 0 | 0 |
| **releng-1/decision** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 22 | 22 |
| **releng-1/linux-docker** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 28 | 28 |
| **releng-3/decision** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 15 | 15 |
| **releng-3/linux-docker** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 23 | 23 |
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
| **releng-hardware/nss-1-b-osx-1015** | generic-worker | 60.3.4 | multiuser | 943a6f2b0d | darwin | amd64 | 1.22.0 | 0 | 0 |
| **releng-hardware/win10-64-2009-hw** | generic-worker | 86.0.2 | multiuser | 9ce326432e | windows | amd64 | 1.24.4 | 0 | 0 |
| **releng-hardware/win10-64-2009-hw-alpha** | generic-worker | 86.0.2 | multiuser | 9ce326432e | windows | amd64 | 1.24.4 | 0 | 0 |
| **releng-hardware/win11-64-24h2-hw** | generic-worker | 91.1.1 | multiuser | 9ee80dc4d0 | windows | amd64 | 1.25.3 | 0 | 0 |
| **releng-hardware/win11-64-24h2-hw-alpha** | generic-worker | 93.1.4 | multiuser | 1088d85c3b | windows | amd64 | 1.25.4 | 0 | 0 |
| **releng-hardware/win11-64-24h2-hw-perf-sheriff** | generic-worker | 86.0.2 | multiuser | 9ce326432e | windows | amd64 | 1.24.4 | 0 | 0 |
| **releng-hardware/win11-64-24h2-hw-ref** | generic-worker | 91.1.1 | multiuser | 9ee80dc4d0 | windows | amd64 | 1.25.3 | 0 | 0 |
| **releng-hardware/win11-64-24h2-hw-ref-alpha** | generic-worker | 93.1.4 | multiuser | 1088d85c3b | windows | amd64 | 1.25.4 | 0 | 0 |
| **releng-hardware/win11-64-24h2-hw-relops1213** | generic-worker | 93.1.4 | multiuser | 1088d85c3b | windows | amd64 | 1.25.4 | 0 | 0 |
| **releng-t/linux-docker** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 189 | 189 |
| **relops-1/decision** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 1 | 1 |
| **relops-1/images** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 1 | 1 |
| **relops-3/decision** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 1 | 1 |
| **relops-3/images** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 1 | 1 |
| **scriptworker-1/decision** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 5 | 5 |
| **scriptworker-1/images** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 26 | 26 |
| **scriptworker-3/decision** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 5 | 5 |
| **scriptworker-3/images** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 26 | 26 |
| **taskgraph-1/decision** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 12 | 12 |
| **taskgraph-1/images** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 3 | 3 |
| **taskgraph-3/decision** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 6 | 6 |
| **taskgraph-3/images** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 5 | 5 |
| **taskgraph-t/linux-docker** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 72 | 72 |
| **taskgraph-t/win11-64-24h2** | generic-worker | 88.1.3 | multiuser | ffc99d77c0 | windows | amd64 | 1.25.1 | 3 | 3 |
| **translations-1/b-linux-large-gcp-1tb-32-256-d2g** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 8 | 8 |
| **translations-1/b-linux-large-gcp-1tb-32-256-std-d2g** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 3 | 3 |
| **translations-1/b-linux-large-gcp-1tb-64-512-d2g** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 1 | 1 |
| **translations-1/b-linux-large-gcp-1tb-64-512-std-d2g** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 1 | 1 |
| **translations-1/b-linux-large-gcp-d2g** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 173 | 173 |
| **translations-1/b-linux-large-gcp-d2g-1tb** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 1 | 1 |
| **translations-1/b-linux-large-gcp-d2g-1tb-standard** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 1 | 1 |
| **translations-1/b-linux-large-gcp-d2g-300gb** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 22 | 22 |
| **translations-1/b-linux-v100-gpu-d2g** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 57 | 57 |
| **translations-1/b-linux-v100-gpu-d2g-4** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 8 | 8 |
| **translations-1/b-linux-v100-gpu-d2g-4-1tb** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 5 | 5 |
| **translations-1/b-linux-v100-gpu-d2g-4-1tb-alpha** | generic-worker | 84.1.2 | multiuser | 4a139f395a | linux | amd64 | 1.24.4 | 1 | 1 |
| **translations-1/b-linux-v100-gpu-d2g-4-1tb-standard** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 2 | 2 |
| **translations-1/b-linux-v100-gpu-d2g-4-1tb-std-alpha** | generic-worker | 84.1.2 | multiuser | 4a139f395a | linux | amd64 | 1.24.4 | 1 | 1 |
| **translations-1/b-linux-v100-gpu-d2g-4-2tb** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 8 | 8 |
| **translations-1/b-linux-v100-gpu-d2g-4-2tb-alpha** | generic-worker | 84.1.2 | multiuser | 4a139f395a | linux | amd64 | 1.24.4 | 1 | 1 |
| **translations-1/b-linux-v100-gpu-d2g-4-300gb** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 7 | 7 |
| **translations-1/b-linux-v100-gpu-d2g-4-300gb-alpha** | generic-worker | 84.1.2 | multiuser | 4a139f395a | linux | amd64 | 1.24.4 | 1 | 1 |
| **translations-1/b-linux-v100-gpu-d2g-4-300gb-standard** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 1 | 1 |
| **translations-1/b-linux-v100-gpu-d2g-4-300gb-std-alpha** | generic-worker | 84.1.2 | multiuser | 4a139f395a | linux | amd64 | 1.24.4 | 1 | 1 |
| **translations-1/b-linux-v100-gpu-d2g-4-alpha** | generic-worker | 84.1.2 | multiuser | 4a139f395a | linux | amd64 | 1.24.4 | 1 | 1 |
| **translations-1/b-linux-v100-gpu-d2g-alpha** | generic-worker | 84.1.2 | multiuser | 4a139f395a | linux | amd64 | 1.24.4 | 1 | 1 |
| **translations-1/decision** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 1153 | 1153 |
| **translations-1/images** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 2 | 2 |
| **xpi-1/b-linux** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 80 | 80 |
| **xpi-1/decision** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 36 | 36 |
| **xpi-1/images** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 1 | 1 |
| **xpi-3/b-linux** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 1 | 1 |
| **xpi-3/decision** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 1 | 1 |
| **xpi-3/images** | generic-worker | 93.1.2 | multiuser | e954a7555f | linux | amd64 | 1.25.4 | 1 | 1 |


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
| projects/taskcluster-imaging/global/images/docker-firefoxci-gcp-lt-googlecompute-2023-04-13t21-30-28z | 1 |
| projects/fxci-production-level3-workers/global/images/docker-firefoxci-gcp-l3-googlecompute-2024-02-05t23-18-22z | 52 |
| projects/taskcluster-imaging/global/images/docker-firefoxci-gcp-l1-googlecompute-2025-06-13t18-31-38z | 85 |
| projects/taskcluster-imaging/global/images/docker-worker-gcp-u14-04-2025-06-16 | 19 |


| Worker Pool | Implementation | Version | Total Workers | Total Capacity |
| --- | --- | --- | ---: | ---: |
| **app-services-1/b-linux-amd** | docker-worker | 38.0.5 | 1 | 1 |
| **app-services-1/b-linux-gcp** | docker-worker | 38.0.5 | 174 | 174 |
| **app-services-1/decision-gcp** | docker-worker | 38.0.5 | 828 | 1656 |
| **app-services-1/images-gcp** | docker-worker | 38.0.5 | 1 | 1 |
| **app-services-3/b-linux-amd** | docker-worker | 38.0.5 | 1 | 1 |
| **app-services-3/b-linux-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **app-services-3/decision-gcp** | docker-worker | 38.0.5 | 1 | 2 |
| **app-services-3/images-gcp** | docker-worker | 38.0.5 | 1 | 1 |
| **code-analysis-1/linux-gcp** | docker-worker | 38.0.5 | 1 | 1 |
| **code-analysis-3/linux-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **comm-1/b-linux-amd** | docker-worker | 38.0.5 | 1 | 1 |
| **comm-1/b-linux-gcp** | docker-worker | 38.0.5 | 236 | 236 |
| **comm-1/b-linux-large-amd** | docker-worker | 38.0.5 | 1 | 1 |
| **comm-1/b-linux-large-gcp** | docker-worker | 38.0.5 | 1 | 1 |
| **comm-1/b-linux-xlarge-amd** | docker-worker | 38.0.5 | 1 | 1 |
| **comm-1/b-linux-xlarge-gcp** | docker-worker | 38.0.5 | 1 | 1 |
| **comm-1/decision-gcp** | docker-worker | 38.0.5 | 131 | 131 |
| **comm-1/images-gcp** | docker-worker | 38.0.5 | 1 | 1 |
| **comm-2/b-linux-amd** | docker-worker | 38.0.5 | 1 | 1 |
| **comm-2/b-linux-gcp** | docker-worker | 38.0.5 | 1 | 1 |
| **comm-2/b-linux-large-amd** | docker-worker | 38.0.5 | 1 | 1 |
| **comm-2/b-linux-large-gcp** | docker-worker | 38.0.5 | 1 | 1 |
| **comm-2/b-linux-xlarge-amd** | docker-worker | 38.0.5 | 1 | 1 |
| **comm-2/b-linux-xlarge-gcp** | docker-worker | 38.0.5 | 1 | 1 |
| **comm-2/decision-gcp** | docker-worker | 38.0.5 | 1 | 1 |
| **comm-2/images-gcp** | docker-worker | 38.0.5 | 1 | 1 |
| **comm-3/b-linux-amd** | docker-worker | 38.0.5 | 1 | 1 |
| **comm-3/b-linux-gcp** | docker-worker | 38.0.5 | 3511 | 3511 |
| **comm-3/b-linux-large-amd** | docker-worker | 38.0.5 | 1 | 1 |
| **comm-3/b-linux-large-gcp** | docker-worker | 38.0.5 | 4 | 4 |
| **comm-3/b-linux-xlarge-amd** | docker-worker | 38.0.5 | 2 | 2 |
| **comm-3/b-linux-xlarge-gcp** | docker-worker | 38.0.5 | 5 | 5 |
| **comm-3/decision-gcp** | docker-worker | 38.0.5 | 92 | 92 |
| **comm-3/images-gcp** | docker-worker | 38.0.5 | 1 | 1 |
| **comm-t/misc-gcp** | docker-worker | 38.0.5 | 1 | 8 |
| **comm-t/t-linux-large-gcp** | docker-worker | unknown version | 1 | 1 |
| **comm-t/t-linux-large-noscratch-gcp** | docker-worker | unknown version | 1 | 1 |
| **comm-t/t-linux-xlarge-gcp** | docker-worker | unknown version | 1 | 1 |
| **comm-t/t-linux-xlarge-noscratch-gcp** | docker-worker | unknown version | 1 | 1 |
| **comm-t/t-linux-xlarge-source-gcp** | docker-worker | unknown version | 1 | 1 |
| **comm-t/t-linux-xlarge-source-noscratch-gcp** | docker-worker | unknown version | 1 | 1 |
| **gecko-1/b-linux-amd** | docker-worker | 38.0.5 | 90 | 90 |
| **gecko-1/b-linux-gcp** | docker-worker | 38.0.5 | 23 | 23 |
| **gecko-1/b-linux-gcp-relops1411** | docker-worker | 38.0.5 | 1 | 1 |
| **gecko-1/b-linux-gcp-test-bug-1882320** | docker-worker | 38.0.5 | 1 | 1 |
| **gecko-1/b-linux-kvm-gcp** | docker-worker | 38.0.5 | 1 | 1 |
| **gecko-1/b-linux-large-amd** | docker-worker | 38.0.5 | 1 | 1 |
| **gecko-1/b-linux-large-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **gecko-1/b-linux-medium-gcp** | docker-worker | 38.0.5 | 1 | 1 |
| **gecko-1/b-linux-xlarge-amd** | docker-worker | 38.0.5 | 1 | 1 |
| **gecko-1/b-linux-xlarge-gcp** | docker-worker | 38.0.5 | 1 | 1 |
| **gecko-1/b-linux-xlarge-gcp-bug1797804-c2** | docker-worker | 38.0.5 | 1 | 1 |
| **gecko-1/decision-gcp** | docker-worker | 38.0.5 | 820 | 820 |
| **gecko-1/decision-gcp-c3d4-amd-bug1990935** | docker-worker | 38.0.5 | 1 | 1 |
| **gecko-1/decision-gcp-c3d8-amd-bug1990935** | docker-worker | 38.0.5 | 1 | 1 |
| **gecko-1/decision-gcp-c4d4-amd-bug1990935** | docker-worker | 38.0.5 | 1 | 1 |
| **gecko-1/decision-gcp-c4d4-hcpu-amd-bug1990935** | docker-worker | 38.0.5 | 1 | 1 |
| **gecko-1/decision-gcp-c4d8-amd-bug1990935** | docker-worker | 38.0.5 | 1 | 1 |
| **gecko-1/decision-gcp-c4d8-hcpu-amd-bug1990935** | docker-worker | 38.0.5 | 1 | 1 |
| **gecko-1/decision-gcp-c4d8-lssd-amd-bug1990935** | docker-worker | 38.0.5 | 1 | 1 |
| **gecko-1/images-gcp** | docker-worker | 38.0.5 | 1 | 1 |
| **gecko-2/b-linux-amd** | docker-worker | 38.0.5 | 1 | 1 |
| **gecko-2/b-linux-gcp** | docker-worker | 38.0.5 | 1 | 1 |
| **gecko-2/b-linux-kvm-gcp** | docker-worker | 38.0.5 | 1 | 1 |
| **gecko-2/b-linux-large-amd** | docker-worker | 38.0.5 | 1 | 1 |
| **gecko-2/b-linux-large-gcp** | docker-worker | 38.0.5 | 1 | 1 |
| **gecko-2/b-linux-medium-gcp** | docker-worker | 38.0.5 | 1 | 1 |
| **gecko-2/b-linux-xlarge-amd** | docker-worker | 38.0.5 | 1 | 1 |
| **gecko-2/b-linux-xlarge-gcp** | docker-worker | 38.0.5 | 1 | 1 |
| **gecko-2/decision-gcp** | docker-worker | 38.0.5 | 9 | 9 |
| **gecko-2/images-gcp** | docker-worker | 38.0.5 | 1 | 1 |
| **gecko-3/b-linux-amd** | docker-worker | 38.0.5 | 421 | 421 |
| **gecko-3/b-linux-gcp** | docker-worker | 38.0.5 | 1270 | 1270 |
| **gecko-3/b-linux-kvm-gcp** | docker-worker | 38.0.5 | 4 | 4 |
| **gecko-3/b-linux-large-amd** | docker-worker | 38.0.5 | 2 | 2 |
| **gecko-3/b-linux-large-gcp** | docker-worker | 38.0.5 | 1 | 1 |
| **gecko-3/b-linux-medium-gcp** | docker-worker | 38.0.5 | 145 | 145 |
| **gecko-3/b-linux-xlarge-amd** | docker-worker | 38.0.5 | 13 | 13 |
| **gecko-3/b-linux-xlarge-gcp** | docker-worker | 38.0.5 | 6 | 6 |
| **gecko-3/decision-gcp** | docker-worker | 38.0.5 | 819 | 819 |
| **gecko-3/images-gcp** | docker-worker | 38.0.5 | 1 | 1 |
| **gecko-t/misc-gcp** | docker-worker | 38.0.5 | 13 | 104 |
| **gecko-t/t-linux-kvm-gcp** | docker-worker | unknown version | 1 | 1 |
| **gecko-t/t-linux-kvm-gcp-relops1411** | docker-worker | unknown version | 1 | 1 |
| **gecko-t/t-linux-kvm-noscratch-gcp** | docker-worker | unknown version | 5 | 5 |
| **gecko-t/t-linux-large-gcp** | docker-worker | unknown version | 1 | 1 |
| **gecko-t/t-linux-large-hostub2204-gcp** | docker-worker | 48.3.0 | 1 | 1 |
| **gecko-t/t-linux-large-noscratch-gcp** | docker-worker | unknown version | 1448 | 1448 |
| **gecko-t/t-linux-xlarge-gcp** | docker-worker | unknown version | 19 | 19 |
| **gecko-t/t-linux-xlarge-noscratch-gcp** | docker-worker | unknown version | 937 | 937 |
| **gecko-t/t-linux-xlarge-source-gcp** | docker-worker | unknown version | 109 | 109 |
| **gecko-t/t-linux-xlarge-source-noscratch-gcp** | docker-worker | unknown version | 1 | 1 |
| **glean-1/b-linux-gcp** | docker-worker | 38.0.5 | 1 | 1 |
| **glean-1/decision-gcp** | docker-worker | 38.0.5 | 822 | 1644 |
| **glean-1/images-gcp** | docker-worker | 38.0.5 | 1 | 1 |
| **glean-3/b-linux-gcp** | docker-worker | 38.0.5 | 1 | 1 |
| **glean-3/decision-gcp** | docker-worker | 38.0.5 | 1 | 2 |
| **glean-3/images-gcp** | docker-worker | 38.0.5 | 1 | 1 |
| **infra/build-decision** | docker-worker | 38.0.5 | 3 | 24 |
| **mobile-1/b-linux-gcp** | docker-worker | 38.0.5 | 59 | 59 |
| **mobile-1/decision-gcp** | docker-worker | 38.0.5 | 778 | 1556 |
| **mobile-1/images-gcp** | docker-worker | 38.0.5 | 1 | 1 |
| **mobile-3/b-linux-gcp** | docker-worker | 38.0.5 | 32 | 32 |
| **mobile-3/decision-gcp** | docker-worker | 38.0.5 | 40 | 80 |
| **mobile-3/images-gcp** | docker-worker | 38.0.5 | 1 | 1 |
| **mozilla-1/b-linux-gcp** | docker-worker | 38.0.5 | 1 | 1 |
| **mozilla-1/decision-gcp** | docker-worker | 38.0.5 | 814 | 1628 |
| **mozilla-1/images-gcp** | docker-worker | 38.0.5 | 3 | 3 |
| **mozilla-3/b-linux-gcp** | docker-worker | 38.0.5 | 1 | 1 |
| **mozilla-3/decision-gcp** | docker-worker | 38.0.5 | 1 | 2 |
| **mozilla-3/images-gcp** | docker-worker | 38.0.5 | 1 | 1 |
| **mozilla-t/t-linux-large-gcp** | docker-worker | unknown version | 3 | 3 |
| **mozilla-t/t-linux-large-noscratch-gcp** | docker-worker | unknown version | 1 | 1 |
| **mozillavpn-1/b-linux-gcp** | docker-worker | 38.0.5 | 40 | 40 |
| **mozillavpn-1/b-linux-large-gcp** | docker-worker | 38.0.5 | 25 | 25 |
| **mozillavpn-1/decision-gcp** | docker-worker | 38.0.5 | 818 | 1636 |
| **mozillavpn-1/images-gcp** | docker-worker | 38.0.5 | 2 | 2 |
| **mozillavpn-3/b-linux-gcp** | docker-worker | 38.0.5 | 1 | 1 |
| **mozillavpn-3/b-linux-large-gcp** | docker-worker | 38.0.5 | 1 | 1 |
| **mozillavpn-3/decision-gcp** | docker-worker | 38.0.5 | 1 | 2 |
| **mozillavpn-3/images-gcp** | docker-worker | 38.0.5 | 1 | 1 |
| **nss-1/decision-gcp** | docker-worker | 38.0.5 | 1 | 2 |
| **nss-1/images-gcp** | docker-worker | 38.0.5 | 1 | 1 |
| **nss-1/linux-gcp** | docker-worker | 38.0.5 | 1 | 1 |
| **nss-3/decision-gcp** | docker-worker | 38.0.5 | 1 | 2 |
| **nss-3/images-gcp** | docker-worker | 38.0.5 | 1 | 1 |
| **nss-3/linux-gcp** | docker-worker | 38.0.5 | 1 | 1 |
| **nss-t/t-linux-xlarge-gcp** | docker-worker | unknown version | 1 | 1 |
| **nss-t/t-linux-xlarge-noscratch-gcp** | docker-worker | unknown version | 1 | 1 |
| **releng-1/b-linux-gcp** | docker-worker | 38.0.5 | 1 | 1 |
| **releng-1/decision-gcp** | docker-worker | 38.0.5 | 1 | 2 |
| **releng-1/linux-gcp** | docker-worker | 38.0.5 | 1 | 1 |
| **releng-3/b-linux-gcp** | docker-worker | 38.0.5 | 1 | 1 |
| **releng-3/decision-gcp** | docker-worker | 38.0.5 | 1 | 2 |
| **releng-3/linux-gcp** | docker-worker | 38.0.5 | 1 | 1 |
| **relops-1/decision-gcp** | docker-worker | 38.0.5 | 1 | 2 |
| **relops-1/images-gcp** | docker-worker | 38.0.5 | 1 | 1 |
| **relops-3/decision-gcp** | docker-worker | 38.0.5 | 1 | 2 |
| **relops-3/images-gcp** | docker-worker | 38.0.5 | 1 | 1 |
| **scriptworker-1/b-linux-gcp** | docker-worker | 38.0.5 | 1 | 1 |
| **scriptworker-1/decision-gcp** | docker-worker | 38.0.5 | 1 | 2 |
| **scriptworker-1/images-gcp** | docker-worker | 38.0.5 | 1 | 1 |
| **scriptworker-3/b-linux-gcp** | docker-worker | 38.0.5 | 1 | 1 |
| **scriptworker-3/decision-gcp** | docker-worker | 38.0.5 | 1 | 2 |
| **scriptworker-3/images-gcp** | docker-worker | 38.0.5 | 1 | 1 |
| **taskgraph-1/decision-gcp** | docker-worker | 38.0.5 | 1 | 2 |
| **taskgraph-1/images-gcp** | docker-worker | 38.0.5 | 1 | 1 |
| **taskgraph-3/decision-gcp** | docker-worker | 38.0.5 | 1 | 2 |
| **taskgraph-3/images-gcp** | docker-worker | 38.0.5 | 1 | 1 |
| **translations-1/decision-gcp** | docker-worker | 38.0.5 | 1906 | 3812 |
| **translations-1/images-gcp** | docker-worker | 38.0.5 | 1 | 1 |
| **xpi-1/b-linux-gcp** | docker-worker | 38.0.5 | 1 | 1 |
| **xpi-1/decision-gcp** | docker-worker | 38.0.5 | 7 | 14 |
| **xpi-1/images-gcp** | docker-worker | 38.0.5 | 1 | 1 |
| **xpi-3/b-linux-gcp** | docker-worker | 38.0.5 | 1 | 1 |
| **xpi-3/decision-gcp** | docker-worker | 38.0.5 | 1 | 2 |
| **xpi-3/images-gcp** | docker-worker | 38.0.5 | 1 | 1 |


## Script Worker

Total: `38`



| Worker Pool | Implementation | Version | Total Workers | Total Capacity |
| --- | --- | --- | ---: | ---: |
| **scriptworker-k8s/app-services-3-beetmover** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/app-services-3-signing** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
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
| **scriptworker-k8s/mobile-1-pushapk** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/mobile-3-bitrise** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/mobile-3-pushapk** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/mobile-3-signing** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/mobile-t-signing** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/mozillavpn-3-signing** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/translations-1-beetmover** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/xpi-t-signing** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
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

Total: `3`


Count by image:

| Version | Count |
| :--- | ---: |
|  | 2 |
| projects/taskcluster-imaging/global/images/generic-2204-wayland-vm-gcp-googlecompute-2023-09-22t17-39-37z | 1 |


| Worker Pool | Implementation | Version | Total Workers | Total Capacity |
| --- | --- | --- | ---: | ---: |
| **built-in/fail** |  | No artifacts found | 0 | 0 |
| **built-in/succeed** |  | No artifacts found | 0 | 0 |
| **gecko-t/t-linux-vm-2204-wayland** |  | No artifacts found | 37 | 37 |


## Version not determined [^2]

Total: `7`


Count by image:

| Version | Count |
| :--- | ---: |
| projects/taskcluster-imaging/global/images/generic-2204-wayland-vm-gcp-googlecompute-2023-09-22t17-39-37z | 1 |
|  | 3 |
| unknown | 2 |
| projects/taskcluster-imaging/global/images/gw-fxci-gcp-l1-2404-amd64-googlecompute-alpha | 1 |


| Worker Pool | Implementation | Version | Total Workers | Total Capacity |
| --- | --- | --- | ---: | ---: |
| **enterprise-t/win11-64-24h2-source** |  | Version not determined; task not (yet) claimed | 95 | 95 |
| **gecko-t/t-linux-2404-relsre** |  | Version not determined; task not (yet) claimed | 1 | 1 |
| **gecko-t/t-linux-vm-2204-wayland-snap** |  | Version not determined; task not (yet) claimed | 1 | 1 |
| **gecko-t/win11-64-24h2-amd** |  | Version not determined; task not (yet) claimed | 34 | 34 |
| **releng-hardware/gecko-t-linux-netperf-1804** |  | Version not determined; task not (yet) claimed | 0 | 0 |
| **releng-hardware/nss-3-b-osx-1015** |  | Version not determined; task not (yet) claimed | 0 | 0 |
| **scriptworker-prov-v1/adhoc-signing-mac14m2** |  | Version not determined; task not (yet) claimed | 0 | 0 |



[^1]: Those are the pools whose tasks were claimed and resolved by a worker as expected, but the worker did not publish either artifact `public/logs/live_backing.log` nor `public/logs/chain_of_trust.log`, which is the source used to identify the worker implementation.

[^2]: Probing task remains pending after two hours. Those are the pools that were not able to start any worker to claim the task within two hours.

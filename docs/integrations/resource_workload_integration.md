# Instruction on how to create a Workload Security Integration

---
page_title: "autobotai_workload_security_integration"
description: |-
  Provides an autobotAI Workload Security Integration.
---

# Resource `autobotai_workload_security_integration`
Provides an autobotAI Workload Security Integration.


## Required things 
1. apikey and url
2. alias,groups,api_version,secretkey,url

## Steps 
1. Create terraform "main.tf" inside "example/integrations/workload_security" folder for autobotAI Workload Security Integration, create the file if not exists and add the following values:
    ## Example Usage 
    ```
    resource "autobotai_workload_security_integration" "wl_security" {
    alias            = "MySEcurityTesting"
    groups=["test"]
    api_version="v1"
    secretkey="cdscwewvewew"
    url="url"
    } 

    ```
2. To use the autobotAI Workload Security Integration follow the "index.md"

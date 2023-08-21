# Instruction for how to create a Fetcher

---
page_title: "autobotai_fetcher"
description: |-
  Provides to create an autobotAI Fetcher
---

# Resource `autobotai_fetcher`
 Provides to create an autobotAI Fetcher

## Required things 
1. apikey and url
2. name, clients,code,integration_type

## Steps 
1. Create terraform "main.tf" inside "example/bot_building_block/fetcher" folder for autobotAI Fetcher, create the file if not  exists and add the following values:
    ## Example Usage 
    ```
    resource "autobotai_fetcher" "fetcher" {
        integration_type            = "azure"
        clients       =  ["ComputeManagementClient","ResourceManagementClient"]
        code     = "#code"
        name = "testings-pulkit"
    }
  
    ```
2. To use the autobotAI Fetcher follow the "index.md"
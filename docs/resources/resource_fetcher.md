# Guidance on Creating a Fetcher

---
    page_title: "autobotai_fetcher"
    description: Enables the creation of an autobotAI Fetcher.
---

## Resource `autobotai_fetcher`
 Discover how to set up an autobotAI Fetcher

## Prerequisites
1. apikey and url
2. name, clients,code,integration_type

## Steps 
1. Generate a Terraform file named "main.tf" within the "example/bot_building_block/fetcher" directory to create an autobotAI Fetcher. If the file does not exist, create it, and include the following content:
    ## Example Usage 
    ```
    resource "autobotai_fetcher" "fetcher" {
        integration_type            = "azure"
        clients       =  ["ComputeManagementClient","ResourceManagementClient"]
        code     = "#code"
        name = "testings-pulkit"
    }
  
    ```
2. For implementing autobotAI fetcher, refer to the instructions provided in the [autobotAI-Provider-Guidance](provider_guidance.md) document.
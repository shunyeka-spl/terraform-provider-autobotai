# Guidance on Creating a Workload Integration

---
    page_title: "autobotai_workload_security_integration"
    description: Enables the creation of an autobotAI Workload Integration
---

## Resource `autobotai_workload_security_integration`
Discover how to set up an autobotAI Workload Integration


## Prerequisites 
1. apikey and url
2. alias,groups,api_version,secretkey,url

## Steps 
1. Create terraform "main.tf" inside "example/integrations/workload_security" folder for autobotAI Workload Security Integration. If the file doesn't exist, create it and input the following values:
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
2. For implementing autobotAI Workload Integration, refer to the instructions provided in the [autobotAI-Provider-Guidance](../autobotAI_provider_guidance.md) document.

# Guidance on Creating a Conformity Integration

---
    page_title: "autobotai_conformity_integration"
    description: Enables the creation of an autobotAI Conformity Integration
---

## Resource `autobotai_conformity_integration`
Discover how to set up an autobotAI Conformity Integration


## Prerequisites
1. apikey and url
2. alias, api_key,region,groups

## Steps 
1. Create terraform "main.tf" inside "example/integrations/conformity" folder for autobotAI Conformity Integration. If the file doesn't exist, create it and input the following values:
    ## Example Usage 
    ```
    resource "autobotai_conformity_integration" "conformity" {
    alias            = "Conformity"
    groups=["test"]
    api_key="test@123"
    region=["ap-south-1"]
    }
    ```
2. For implementing autobotAI Conformity Integration, refer to the instructions provided in the [autobotAI-Provider-Guidance](../autobotAI_provider_guidance.md) document.

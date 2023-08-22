# Guidance on Creating an Azure Integration

---
    page_title: "autobotai_azure_integration"
    description: Enables the creation of an autobotAI Azure Integration
---

## Resource `autobotai_azure_integration`
 Discover how to set up an autobotAI Azure Integration

## Prerequisites
1. apikey and url
2. alias, client_id,tenant_id,secret_key,subscription_id,groups,cspname

## Steps 
1. Create terraform "main.tf" inside "example/integrations/azure" folder for autobotAI Azure Integration. If the file doesn't exist, create it and input the following values:
    ## Example Usage 
    ```
        resource "autobotai_azure_integration" "azure" {  
        alias            = "MyProject"
        client_id       = "autobot647"
        tenant_id     = "12345678"
        secret_key = "123345678"
        subscription_id = "1se345f4"
        groups=["test"]
        cspname="azure"
    }   
    ```
2. For implementing autobotAI Azure Integration, refer to the instructions provided in the [autobotAI-Provider-Guidance](../autobotAI_provider_guidance.md) document.
---
page_title: "autobotai_azure_integration"
description: |-
  Provides an autobotAI Azure Integration.
---

# Resource `autobotai_azure_integration`
Provides an autobotAI Azure Integration.

# Instruction for how to create a Azure Integration
## Required things 
1. apikey and url
2. alias, client_id,tenant_id,secret_key,subscription_id,groups,cspname

## Steps 
1. Create terraform "main.tf" inside "example/integrations/azure" folder for autobotAI Azure Integration, create the file if not  exists and add the following values:
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
2. To use the autobotAI Azure Integration follow the "index.md"
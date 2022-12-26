---
page_title: "autobot_azure_integration"
description: |-
  Provides a Autobot Azure Integration.
---

# Resource `autobot_azure_integration`
Provides a Autobot Azure Integration.

# Instruction for how to create a Azure Integration
## Required things 
1. apikey
2. alias, client_id,tenant_id,secret_key,subscription_id,groups,cspname

## Steps 
1. Create terraform "main.tf" inside "example/azure" folder for Autobot Azure Integration, create the file if not exists and add the following vallues:
    ## Example Usage 
    ```
        resource "autobot_azure_integration" "azure" {  
        alias            = "MyProject"
        client_id       = "autobot647"
        tenant_id     = "12345678"
        secret_key = "123345678"
        subscription_id = "1se345f4"
        groups=["test"]
        cspname="azure"
    }   
    ```
2. To use the Autobot Azure Integration add the API key to the varibale.tf 
3. create "variable.tf" inside "example/azure" folder and add the following values:
    ## Example Usage 
    ```
    variable "apikey"{
    type    = string
    default = "API Key"
    }

    ```
Note: You can always change the values declared according to your choice.
4. Run Terraform 
    1. Navigate to the Project Directory:
        ```
        cd /path/autobot_integration

        ``` 
    2. Create the Artifact:
        ```
        make install

        ```
    3. Now, you can run the terraform code by navigating to "example/azure" folder:
        ```
        cd example/azure
        terraform init
        terraform apply

        ```


---
page_title: "autobot_conformity_integration"
description: |-
  Provides a Autobot Conformity Integration.
---

# Resource `autobot_conformity_integration`
Provides a Autobot Conformity Integration.

# Instruction on how to create a Conformity Integration

## Required things 
1. apikey
2. alias, api_key,region,groups

## Steps 
1. Create terraform "main.tf" inside "example/conformity" folder for Autobot Conformity Integration, create the file if not exists and add the following vallues:
    ## Example Usage 
    ```
    resource "autobot_conformity_integration" "conformity" {
    alias            = "Conformity"
    groups=["test"]
    api_key="test@123"
    region=["ap-south-1"]
    }
    ```
2. To use the Autobot Conformity Integration add the API key to the varibale.tf 
3. create "variable.tf" inside "example/conformity" folder and add the following values:
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
    3. Now, you can run the terraform code by navigating to "example/conformity" folder:
        ```
        cd example/conformity
        terraform init
        terraform apply

        ```
---
page_title: "autobotai_workload_security_integration"
description: |-
  Provides a autobotAI Workload Security Integration.
---

# Resource `autobotai_workload_security_integration`
Provides a autobotAI Workload Security Integration.

# Instruction on how to create a Workload Security Integration

## Required things 
1. apikey
2. alias,groups,api_version,secretkey,url

## Steps 
1. Create terraform "main.tf" inside "example/workload_security" folder for autobotAI Workload Security Integration, create the file if not exists and add the following vallues:
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
2. To use the autobotAI Workload Security Integration add the API key to the varibale.tf 
3. create "variable.tf" inside "example/git" folder and add the following values:
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
        cd /path/terraform-provider-autobotai

        ``` 
    2. Create the Artifact:
        ```
        make install

        ```
    3. Now, you can run the terraform code by navigating to "example/workload_Security" folder:
        ```
        cd example/workload_security
        terraform init
        terraform apply

        ```
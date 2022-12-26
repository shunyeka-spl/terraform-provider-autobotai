---
page_title: "autobot_git_integration"
description: |-
  Provides a Autobot Git Integration.
---

# Resource `autobot_git_integration`
Provides a Autobot Git  Integration.

# Instruction on how to create a Git Integration

## Required things 
1. apikey
2. alias,groups,host,token

## Steps 
1. Create terraform "main.tf" inside "example/git" folder for Autobot Git Integration, create the file if not exists and add the following vallues:
    ## Example Usage 
    ```
    resource "autobot_git_integration" "git" {
    alias            = "MyGit"
    groups=["test"]
    host="aseweewewwew"
    token="cdscdscsc"

    }
    ```
2. To use the Autobot Git  Integration add the API key to the varibale.tf 
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
        cd /path/autobot_integration

        ``` 
    2. Create the Artifact:
        ```
        make install

        ```
    3. Now, you can run the terraform code by navigating to "example/git" folder:
        ```
        cd example/git
        terraform init
        terraform apply

        ```
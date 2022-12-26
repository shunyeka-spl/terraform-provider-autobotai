---
page_title: "autobotai_ms_teams_integration"
description: |-
  Provides a autobotAI Ms Teams Integration.
---

# Resource `autobotai_ms_teams_integration`
Provides a autobotAI Ms Teams Integration.

# Instruction on how to create a Ms Teams Integration

## Required things 
1. apikey
2. alias,groups,webhook

## Steps 
1. Create terraform "main.tf" inside "example/ms_teams" folder for autobotAI Ms Teams Integration, create the file if not exists and add the following vallues:
    ## Example Usage 
    ```
    resource "autobotai_ms_teams_integration" "ms_teams" {
    alias            = "MyTeams"
    groups=["test"]
    webhook="webhookurl"
    }   

    ```
2. To use the autobotAI Ms Teams  Integration add the API key to the varibale.tf 
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
    3. Now, you can run the terraform code by navigating to "example/ms_teams" folder:
        ```
        cd example/ms_teams
        terraform init
        terraform apply

        ```
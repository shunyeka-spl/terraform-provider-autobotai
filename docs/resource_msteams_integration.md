---
page_title: "autobot_ms_teams_integration"
description: |-
  Provides a Autobot Ms Teams Integration.
---

# Resource `autobot_ms_teams_integration`
Provides a Autobot Ms Teams Integration.

# Instruction on how to create a Ms Teams Integration

## Required things 
1. apikey
2. alias,groups,webhook

## Steps 
1. Create terraform "main.tf" inside "example/ms_teams" folder for Autobot Ms Teams Integration, create the file if not exists and add the following vallues:
    ## Example Usage 
    ```
    resource "autobot_ms_teams_integration" "ms_teams" {
    alias            = "MyTeams"
    groups=["test"]
    webhook="webhookurl"
    }   

    ```
2. To use the Autobot Ms Teams  Integration add the API key to the varibale.tf 
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
    3. Now, you can run the terraform code by navigating to "example/ms_teams" folder:
        ```
        cd example/ms_teams
        terraform init
        terraform apply

        ```
# Instruction for how to use autobotAI Integrations,Bot Building Block and Inventory
    1. To use the autobotAI  Integrations, Bot Building Block and Inventory add the ApiKey and url to the varibale.tf 
    2. create "variable.tf" inside "example/integrations/(aws,azure,gcp,git,ms_teams,conformity,workload_security)" or "example/bot_building_block/(automation,evaluator,fetcher,listener)" or "example/inventory/(inventory_schedule)" folder and add the following values:
        ## Example Usage 
        ```
        variable "apikey"{
        type    = string
        default = "API Key"
        }

        variable "url" {
            type =string
            default ="url"
        }

        ```
    3. Create the "provider.tf" and configure the required providers as well as the credentials to it
        ## Example Usage
        ```
            terraform {
                required_providers {
                    autobotai = {
                        source  = "shunyeka-spl/autobotai"
                    }
                }
            }
            provider "autobotai" {
                apikey = var.apikey
                url =var.url
            }
        ```
    Note: You can always change the values declared according to your choice.
    3. Run Terraform 
        1. Navigate to the Project Directory:
            ```
             cd /path/terraform-provider-autobotai

             ``` 
        2. Create the Artifact:
            ```
             make install

            ```
        3. Now, you can run the terraform code by navigating to "example/integrations/" or "example/bot_building_block" or "example/inventory" folder:
            ```
            cd example/integrations/(aws,azure,gcp,git,ms_teams,conformity,workload_security)
            terraform init
            terraform apply

            OR

            cd example/bot_building_block/(automation,evaluator,fetcher,listener)
            terraform init
            terraform apply

            OR

            cd example/inventory/(inventory_schedule)
            terraform init
            terraform apply


            ```
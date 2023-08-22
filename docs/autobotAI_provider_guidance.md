# Guidance on utilizing autobotAI Integrations, Bot Building Blocks, and Inventory functionalities.

1. To harness the capabilities of autobotAI Integrations, Bot Building Blocks, and Inventory, follow these steps:

    - Incorporate the API Key and URL into the variable.tf file.
    - Inside the relevant folders—such as example/integrations/(aws,azure,gcp,git,ms_teams,conformity,workload_security) or example/bot_building_block/(automation,evaluator,fetcher,listener) or example/inventory/(inventory_schedule)—create a variable.tf file. Include the following contents:

        ```hcl
        variable "apikey" {
            type    = string
            default = "API Key"
        }

        variable "url" {
            type    = string
            default = "url"
        }
        ```
    - Configure the essential providers and credentials by creating a provider.tf file. Here's an example setup:

         ```hcl
            terraform {
                required_providers {
                    autobotai = {
                        source  = "shunyeka-spl/autobotai"
                    }
                }
            }

            provider "autobotai" {
                apikey = var.apikey
                url    = var.url
            }
        ```
        - Keep in mind that you can modify the declared values based on your preferences.

2. Execute Terraform operations:
    -  Navigate to the project directory:

            ```
                cd /path/terraform-provider-autobotai
            ```
    - Generate the artifact:

            ```
                make install
            ```
    - Run the Terraform code by moving to the appropriate folder—either example/integrations/, example/bot_building_block/, or example/inventory/:

        ```
            cd example/integrations/(aws,azure,gcp,git,ms_teams,conformity,workload_security)
            terraform init
            terraform apply
            ```

            OR

            ```
            cd example/bot_building_block/(automation,evaluator,fetcher,listener)
            terraform init
            terraform apply
            ```

            OR

            ```
            cd example/inventory/(inventory_schedule)
            terraform init
            terraform apply

         ```






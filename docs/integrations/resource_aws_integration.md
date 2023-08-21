# Instruction for how to create an  AWS Integration
---
page_title: "autobotai_aws_integration"
description: |-
  Provides an autobotAI Aws Integration.
---

# Resource `autobotai_aws_integration`
Provides an autobotAI Aws Integration.

# Instruction for how to create a Aws Integration
## Required things for Fetching the ExternlId and TargetPrincipal
1. apikey and url
2. alias,groups and deploy_bots

## Steps 
1. Create terraform "main.tf" inside "example/integrations/aws" folder for autobotAI Aws Integration, create the file if not exists and add the following values:
    ## Example Usage 
    ```
        resource "autobotai_external_id" "aws"{
        alias="My-Project-Testing"
        groups=["test"]
    } 
    ```
2. For creating the Cloud Formation Stack add the following values to "main.tf" and also add the aws required providers and credentials to "provider.tf"
    ## Example Usage(in main.tf)
    ```
        resource "aws_cloudformation_stack" "autobotai_cloud_formation" {
        name         = "Testing"
        template_url = "template_url"
        parameters = {
            APIURL="https://api-test.autobot.live/"
            Environment="test"
            TrustPrincipal  = autobotai_external_id.aws.target_principal
            ExternalID = autobotai_external_id.aws.external_id
        }
        capabilities = ["CAPABILITY_NAMED_IAM"]
    }

    ```
    ## provider.tf
    ```
        terraform {
            required_providers {
                aws = {
                source  = "hashicorp/aws"
                version = "version"
                }
            }
        }
        provider "aws" {
            access_key = "access_key"
            secret_key = "secret_key"
            region = "region"
        }

    ``` 
3. To use the autobotAI Aws Integration follow the "index.md"
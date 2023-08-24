# Guidance on Creating an AWS Integration

---
    page_title: "autobotai_aws_integration"
    description: Enables the creation of an autobotAI AWS Integration
---

## Resource `autobotai_aws_integration`
 Discover how to set up an autobotAI AWS Integration

## Prerequisites for fetching the ExternlId and TargetPrincipal
1. apikey and url
2. alias,groups and deploy_bots

## Steps 
1. Create a Terraform file named "main.tf" within the "example/integrations/aws" directory for setting up autobotAI AWS Integration. If the file doesn't exist, create it and input the following values:
    ## Example Usage 
    ```
        resource "autobotai_external_id" "aws"{
        alias="My-Project-Testing"
        groups=["test"]
    } 
    ```
2. For creating the Cloud Formation Stack, add the following values to the "main.tf" file, and also include the required AWS providers and credentials in the "provider.tf" file:
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
3. For implementing autobotAI AWS Integration, refer to the instructions provided in the [autobotAI-Provider-Guidance](provider_guidance.md) document.
# Guidance on Creating an AWS SES Integration

---
    page_title: "autobotai_aws_ses_configure_integration"
    description: Enables the creation of an autobotAI AWS SES Integration
---

## Resource `autobotai_aws_ses_configure_integration`
 Discover how to set up an autobotAI AWS SES Integration

## Prerequisites
1. apikey and url
2. alias, region,integration_id

## Steps 
1. Create terraform "main.tf" inside "example/integrations/aws_ses_configure" folder for autobotAI AWS SES Configure Integration.If the file doesn't exist, create it and input the following values:
    ## Example Usage 
    ```
        resource "autobotai_aws_ses_configure_integration" "aws_ses" {
        alias            = "MyProjecttesting"
        region       = "us-east-1"
        integration_id     = "16548497741465"
        
    }   
    ```
2. For implementing autobotAI AWS SES Integration, refer to the instructions provided in the [autobotAI-Provider-Guidance](provider_guidance.md) document.
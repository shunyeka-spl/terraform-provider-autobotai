# Instruction for how to create an AWS SES  Integration

---
page_title: "autobotai_aws_ses_configure_integration"
description: |-
  Provides an autobotAI Aws Ses Configure Integration.
---

# Resource `autobotai_aws_ses_configure_integration`
Provides an autobotAI Aws Ses Configure Integration.

# Instruction for how to create a Aws Ses Configure Integration
## Required things 
1. apikey and url
2. alias, region,integration_id

## Steps 
1. Create terraform "main.tf" inside "example/integrations/aws_ses_configure" folder for autobotAI Aws Ses Configure Integration, create the file if not  exists and add the following values:
    ## Example Usage 
    ```
        resource "autobotai_aws_ses_configure_integration" "aws_ses" {
        alias            = "MyProjecttesting"
        region       = "us-east-1"
        integration_id     = "16548497741465"
        
    }   
    ```
2. To use the autobotAI Azure Integration follow the "index.md"
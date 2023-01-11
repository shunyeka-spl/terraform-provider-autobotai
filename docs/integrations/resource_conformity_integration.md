---
page_title: "autobotai_conformity_integration"
description: |-
  Provides an autobotAI Conformity Integration.
---

# Resource `autobotai_conformity_integration`
Provides an autobotAI Conformity Integration.

# Instruction on how to create a Conformity Integration

## Required things 
1. apikey and url
2. alias, api_key,region,groups

## Steps 
1. Create terraform "main.tf" inside "example/integrations/conformity" folder for autobotAI Conformity Integration, create the file if not exists and add the following vallues:
    ## Example Usage 
    ```
    resource "autobotai_conformity_integration" "conformity" {
    alias            = "Conformity"
    groups=["test"]
    api_key="test@123"
    region=["ap-south-1"]
    }
    ```
2. To use the autobotAI Conformity Integration follow the "provider_setup.md"

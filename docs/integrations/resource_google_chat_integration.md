# Instruction on how to create a Google Chat Integration
---
page_title: "autobotai_google_chat_integration"
description: |-
  Provides an autobotAI Google Chat Integration.
---

# Resource `autobotai_google_chat_integration`
Provides an autobotAI Google Chat  Integration.


## Required things 
1. apikey and url
2. alias,groups,webhook

## Steps 
1. Create terraform "main.tf" inside "example/integrations/google_chat" folder for autobotAI Google Chat Integration, create the file if not exists and add the following vallues:
    ## Example Usage 
    ```
    resource "autobotai_google_chat_integration" "google_chat" {
        alias            = "MyProject"
        groups=["test"]
        webhook="webhook"
    }
    ```
2. To use the autobotAI Google Chat  Integration follow the "provider_setup.md"
# Guidance on Creating a Google Chat Integration
---
    page_title: "autobotai_google_chat_integration"
    description: Enables the creation of an autobotAI Google Chat Integration
---

## Resource `autobotai_google_chat_integration`
Discover how to set up an autobotAI Google Chat Integration


## Prerequisites
1. apikey and url
2. alias,groups,webhook

## Steps 
1. Create terraform "main.tf" inside "example/integrations/google_chat" folder for autobotAI Google Chat Integration. If the file doesn't exist, create it and input the following values:
    ## Example Usage 
    ```
    resource "autobotai_google_chat_integration" "google_chat" {
        alias            = "MyProject"
        groups=["test"]
        webhook="webhook"
    }
    ```
2. For implementing autobotAI Google Chat Integration, refer to the instructions provided in the [autobotAI-Provider-Guidance](../autobotAI_provider_guidance.md) document.
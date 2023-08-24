# Guidance on Creating a Listener

---
    page_title: "autobotai_listener"
    description: Enables the creation of an autobotAI Listener
---

## Resource `autobotai_listener`
 Discover how to set up an autobotAI Fetcher

## Prerequisites
1. apikey and url
2. description, webhook_source_integration_type,name

## Steps 
1. Craft a Terraform file named "main.tf" within the "example/bot_building_block/listener" directory to establish an autobotAI Listener. If the file does not already exist, create it and incorporate the subsequent values:
    ## Example Usage 
    ```
    resource "autobotai_listener" "listener" {
        description            = "Testing-pulkit-Testing-Pulkit"
        name       = "testing-pulkit"
        webhook_source_integration_type="conformity"
    }
  
    ```
2. For implementing autobotAI Listener, refer to the instructions provided in the [autobotAI-Provider-Guidance](provider_guidance.md) document.
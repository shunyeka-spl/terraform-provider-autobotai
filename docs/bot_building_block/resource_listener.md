# Instruction for how to create a Listener

---
page_title: "autobotai_listener"
description: |-
  Provides to create an autobotAI Listener
---

# Resource `autobotai_listener`
 Provides to create an autobotAI Listener

## Required things 
1. apikey and url
2. description, webhook_source_integration_type,name

## Steps 
1. Create terraform "main.tf" inside "example/bot_building_block/listener" folder for autobotAI Listener, create the file if not  exists and add the following values:
    ## Example Usage 
    ```
    resource "autobotai_listener" "listener" {
        description            = "Testing-pulkit-Testing-Pulkit"
        name       = "testing-pulkit"
        webhook_source_integration_type="conformity"
    }
  
    ```
2. To use the autobotAI Fetcher follow the "index.md"
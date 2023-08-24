# Guidance on Creating an Automation
---
    page_title: "autobotai_automation"
    description: Enables the creation of an autobotAI Automation.
---

## Resource `autobotai_automation`
Discover how to set up an autobotAI Automation


## Prerequisites
1. apikey and url
2. name, clients,code,title,type,integration_type,approval_required

## Steps 
1. Begin by crafting a file named "main.tf" within the "example/bot_building_block/automation" folder. If this file isn't present, create it and input the subsequent content:
    ## Example Usage 
    ```
        resource "autobotai_automation" "automation" {
            name            = "testing"
            clients       =  ["NetworkManagementClient"]
            code= "#code"
            title ="testing"
            type="communication"
            integration_type= "azure"
            approval_required="true"
        }
  
    ```
2. For implementing autobotAI Automation, refer to the instructions provided in the[autobotAI-Provider-Guidance](provider_guidance.md) document.

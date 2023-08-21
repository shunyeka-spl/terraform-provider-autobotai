# Instruction for how to create an Automation

---
page_title: "autobotai_automation"
description: |-
  Provides to create an autobotAI Automation
---

# Resource `autobotai_automation`
Provides to create an autobotAI Automation

## Required things 
1. apikey and url
2. name, clients,code,title,type,integration_type,approval_required

## Steps 
1. Create terraform "main.tf" inside "example/bot_building_block/automation" folder for autobotAI Automation, create the file if not  exists and add the following values:
    ## Example Usage 
    ```
        resource "autobotai_automation" "automation" {
            name            = "testing-pulkit-testing-pulkit-bhagwate"
            clients       =  ["NetworkManagementClient","NetworkManagementClient"]
            code= "#code"
            title ="testing"
            type="communication"
            integration_type= "azure"
            approval_required="true"
        }
  
    ```
2. To use the autobotAI Automation follow the "index.md"

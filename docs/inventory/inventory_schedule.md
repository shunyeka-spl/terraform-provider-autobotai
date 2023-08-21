# Instruction for how to Schedule an Inventory
---
page_title: "autobotai_inventory_schedule"
description: |-
  Provides an autobotAI Inventory Schedule.
---

# Resource `autobotai_inventory_schedule`
Provides an autobotAI Inventory Schedule.

# Instruction for how to create a Inventory Schedule
## Required things 
1. apikey and url
2. integration_id,integration_type and cron_expression

## Steps 
1. Create terraform "main.tf" inside "example/inventory/inventory_schedule" folder for autobotAI Inventory Schedule, create the file if not  exists and add the following values:
    ## Example Usage 
    ```
        resource "autobotai_inventory_schedule" "inventory" {
        integration_id="qwedwqdwq_sfdfvdcsdcdsdsc"
        integration_type="azure"
        cron_expression="0 0 * * *"
} 
    ```
2. To use the autobotAI Inventory Schedule follow the "index.md"

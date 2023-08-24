# Guidance on Scheduling the Inventory
---
    page_title: "autobotai_inventory_schedule"
    description: Enables the creation of an autobotAI Inventory Schedule
---

## Resource `autobotai_inventory_schedule`
 Discover how to set up an autobotAI Inventory Schedule

## Prerequisites
1. apikey and url
2. integration_id,integration_type and cron_expression

## Steps 
1. Create terraform "main.tf" inside "example/inventory/inventory_schedule" folder for autobotAI Inventory Schedule. If the file doesn't exist, create it and input the following values:
    ## Example Usage 
    ```
        resource "autobotai_inventory_schedule" "inventory" {
          integration_id="qwedwqdwq_sfdfvdcsdcdsdsc"
          integration_type="azure"
          cron_expression="0 0 * * *"
        } 
    ```
    
2. For implementing autobotAI Inventory Schedule, refer to the instructions provided in the [autobotAI-Provider-Guidance](provider_guidance.md) document.

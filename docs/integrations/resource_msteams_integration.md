# Instruction on how to create a Ms Teams Integration

---
page_title: "autobotai_ms_teams_integration"
description: |-
  Provides an autobotAI Ms Teams Integration.
---

# Resource `autobotai_ms_teams_integration`
Provides an autobotAI Ms Teams Integration.


## Required things 
1. apikey and url
2. alias,groups,webhook

## Steps 
1. Create terraform "main.tf" inside "example/integrations/ms_teams" folder for autobotAI Ms Teams Integration, create the file if not exists and add the following vallues:
    ## Example Usage 
    ```
    resource "autobotai_ms_teams_integration" "ms_teams" {
    alias            = "MyTeams"
    groups=["test"]
    webhook="webhookurl"
    }   

    ```
2. To use the autobotAI Ms Teams  Integration follow the "provider_setup.md"

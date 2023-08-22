# Guidance on Creating a MS Teams Integration

---
    page_title: "autobotai_ms_teams_integration"
    description: Enables the creation of an autobotAI MS Teams Integration
---

## Resource `autobotai_ms_teams_integration`
Discover how to set up an autobotAI MS Teams Integration


## Prerequisites
1. apikey and url
2. alias,groups,webhook

## Steps 
1. Create terraform "main.tf" inside "example/integrations/ms_teams" folder for autobotAI Ms Teams Integration. If the file doesn't exist, create it and input the following values:
    ## Example Usage 
    ```
      resource "autobotai_ms_teams_integration" "ms_teams" {
        alias            = "MyTeams"
        groups=["test"]
        webhook="webhookurl"
      }   

    ```
2. For implementing autobotAI MS Teams Integration, refer to the instructions provided in the [autobotAI-Provider-Guidance](../autobotAI_provider_guidance.md) document.

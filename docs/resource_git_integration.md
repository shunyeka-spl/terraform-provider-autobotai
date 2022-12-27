---
page_title: "autobotai_git_integration"
description: |-
  Provides an autobotAI Git Integration.
---

# Resource `autobotai_git_integration`
Provides an autobotAI Git  Integration.

# Instruction on how to create a Git Integration

## Required things 
1. apikey
2. alias,groups,host,token

## Steps 
1. Create terraform "main.tf" inside "example/integrations/git" folder for autobotAI Git Integration, create the file if not exists and add the following vallues:
    ## Example Usage 
    ```
    resource "autobotai_git_integration" "git" {
    alias            = "MyGit"
    groups=["test"]
    host="aseweewewwew"
    token="cdscdscsc"

    }
    ```
2. To use the autobotAI Git  Integration follow the "provider_setup.md"
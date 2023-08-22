# Guidance on Creating a Git Integration

---
    page_title: "autobotai_git_integration"
    description: Enables the creation of an autobotAI Git Integration
---

## Resource `autobotai_git_integration`
Discover how to set up an autobotAI Git Integration


## Prerequisites
1. apikey and url
2. alias,groups,host,token

## Steps 
1. Create terraform "main.tf" inside "example/integrations/git" folder for autobotAI Git Integration. If the file doesn't exist, create it and input the following values:
    ## Example Usage 
    ```
      resource "autobotai_git_integration" "git" {
        alias            = "MyGit"
        groups=["test"]
        host="aseweewewwew"
        token="cdscdscsc"

      }
    ```
2. For implementing autobotAI Git Integration, refer to the instructions provided in the [autobotAI-Provider-Guidance](../autobotAI_provider_guidance.md) document.
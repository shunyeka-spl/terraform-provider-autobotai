# Guidance on Creating a Bot

---
    page_title: "autobotai_bot"
    description: Enables the creation of an autobotAI Bot
---

## Resource `autobotai_bot`
Discover how to set up an autobotAI Bot

## Prerequisites 
1. apikey and url
2. name,topic,category,importance,fetcher_id,integration_id,integration_type,actions{type,automation_id,params{description,name,type,values},rer\quired,action_id}evaluator_id(you can pass the id from evaluator)


## Steps 
1. Create a "main.tf" Terraform file within the "example/bot_building_block/evaluator" directory for autobotAI Bot. If the file doesn't exist, create it, and input the following values:

2. Notably, ensure you place the bot resource after the evaluator resource in the "evaluator/main.tf" file.

    ## Example Usage
    ```
    resource "autobotai_bot" "bot"{
        name="bot-testing"
        topic="testing-bots"
        category="Reliability"
        importance="Medium"
        fetcher_id="fetcherid"
        integration_id="integration_id"
        integration_type="azure"
        actions{
            type="mutation"
            automation_id="automation_id"
            params{
                description="Name of the bucket"
                name="bucket_name"
                type="string"
                values="autobot-bucket-automation-test" 
            }
            required=false
            action_id="0"
        }
        evaluator_id=autobotai_evaluator.evaluator.id
    }

    ```   
3. For implementing autobotAI Bot, refer to the instructions provided in the [autobotAI-Provider-Guidance](../autobotAI_provider_guidance.md) document.
---
page_title: "autobotai_bot"
description: |-
  Provides to create an autobotAI Bot
---

# Resource `autobotai_bot`
Provides to create an autobotAI Bot

# Instruction for how to create a Bot
## Required things 
1. apikey and url
2. name,topic,category,importance,fetcher_id,integration_id,integration_type,actions{type,automation_id,params{description,name,type,values},rer\quired,action_id}evaluator_id(you can pass the id from evaluator)


## Steps 
1. Create terraform "main.tf" inside "example/bot_building_block/evaluator" folder for autobotAI Bot, create the file if not  exists and add the following values:

2. Note - you need to add the bot resource  after the evaluator resource in evaluator/main.tf

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
3. To use the autobotAI Bot  follow the "index.md"
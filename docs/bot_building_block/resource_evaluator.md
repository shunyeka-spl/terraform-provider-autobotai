# Guidance on Creating an Evaluator

---
    page_title: "autobotai_evaluator"
    description: Enables the creation of an autobotAI Evaluator.
---

## Resource `autobotai_evaluator`
Discover how to set up an autobotAI Evaluator

## Prerequisites
1. apikey and url
2. name, id(qb_rules), rules{id(rule_id),field,operator,value,value_score}, combinator,not,preference
Note - you can define multiple rules 

## Steps 
1. To create an autobotAI Evaluator, generate a "main.tf" Terraform file within the "example/bot_building_block/evaluator" directory. If the file doesn't exist, generate it and input the subsequent values:
    ## Example Usage 
    ```
        resource "autobotai_evaluator" "evaluator" {
            name="testing-bot"
            code="null"        
            qb_rules_id="id"
                rules{
                    rules_id="id"
                    field="name"
                    operator="is_not_empty"
                    value_source="value"  
                    value=""
                }
                rules{
                    rules_id="id"
                    field="type"
                    operator="is_not_empty"
                    value_source="value"  
                    value=""
                }
                rules{
                    qb_rules_id="id"
                    rules{
                            rules_id="id"
                            field="region"
                            operator="is_not_empty"
                            value_source="value"  
                            value=""
                        }
                    combinator="and"
                    not=false
                }
            combinator="and"
            not=false
            preference="qb_rules"
        }

    ```
    - After the evaluator is created, you'll need to create a bot. 
    - Refer to the guidance provided in the [Bot-Creation](resource_bot.md) document for step-by-step instructions on how to create a Bot.
 
2. For implementing autobotAI Evaluator, refer to the instructions provided in the [autobotAI-Provider-Guidance](../autobotAI_provider_guidance.md) document.
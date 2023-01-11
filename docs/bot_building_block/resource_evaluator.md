---
page_title: "autobotai_evaluator"
description: |-
  Provides to create an autobotAI Evaluator
---

# Resource `autobotai_evaluator`
Provides to create an autobotAI Evaluator

# Instruction for how to create a Evaluator
## Required things 
1. apikey and url
2. name, id(qb_rules), rules{id(rule_id),field,operator,value,value_score}, combinator,not,preference
Note - you can define multiple rules 

## Steps 
1. Create terraform "main.tf" inside "example/bot_building_block/evaluator" folder for autobotAI Evaluator, create the file if not  exists and add the following values:
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
 # After the evaluator is created you need to create a bot 
 # Instruction for how to create a Bot -> follow "resource_bot.md"
 
2. To use the autobotAI Evaluator  follow the "index.md"
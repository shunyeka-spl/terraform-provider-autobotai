resource "autobotai_evaluator" "evaluator" {
name="testing-bot"
code="null"    
qb_rules_id="g-0.04594995406042268"

    rules{
        rules_id="r-0.9009325597055955"
        field="name"
        operator="is_not_empty"
        value_source="value"  
        value=""
    }
    rules{
        rules_id="r-0.03355383276017343"
        field="type"
        operator="is_not_empty"
        value_source="value"  
        value=""
    }

    rules{
        qb_rules_id="g-0.2095897105507767"
        rules{
                rules_id="r-0.697227662162232"
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

resource "autobotai_bot" "bot"{
name="testing"
topic="testing-bots"
category="Reliability"
importance="Medium"
fetcher_id="63454da5a6ac05bbfbfg0023697494"
integration_id="e9cb5bbfg62a-1f12-42c1-8e68-586e140aa6c1"
integration_type="azure"
    actions{
        type="mutation"
        automation_id="63171dbgbfgbfg962268c20751622346"
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


output "evaluatorRead"{
        value=autobotai_evaluator.evaluator
}
output "botRead"{
        value=autobotai_bot.bot
}



package autobotai

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"testing"
)

func TestAccResourceEvaluator(t *testing.T) {

	name := "testing-bot"
	code := "null"
	qbRulesId := "63bafb3f47a9176af4a97a93"
	preference := "qb_rules"
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccAutobotAIPreCheck(t) },
		Providers: testAccAutobotAIProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckEvaluatorConfigBasic(name, code, qbRulesId, preference),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("autobotai_evaluator.evaluator", "name", name),
					resource.TestCheckResourceAttr("autobotai_evaluator.evaluator", "code", code),
					resource.TestCheckResourceAttr("autobotai_evaluator.evaluator", "qb_rules_id", qbRulesId),

					resource.TestCheckResourceAttr("autobotai_evaluator.evaluator", "rules.0.rules_id", "r-0.9009325597055955"),
					resource.TestCheckResourceAttr("autobotai_evaluator.evaluator", "rules.0.field", "name"),
					resource.TestCheckResourceAttr("autobotai_evaluator.evaluator", "rules.0.operator", "is_not_empty"),
					resource.TestCheckResourceAttr("autobotai_evaluator.evaluator", "rules.0.value_source", "value"),
					resource.TestCheckResourceAttr("autobotai_evaluator.evaluator", "rules.0.value", ""),

					resource.TestCheckResourceAttr("autobotai_evaluator.evaluator", "rules.1.rules_id", "r-0.03355383276017343"),
					resource.TestCheckResourceAttr("autobotai_evaluator.evaluator", "rules.1.field", "type"),
					resource.TestCheckResourceAttr("autobotai_evaluator.evaluator", "rules.1.operator", "is_not_empty"),
					resource.TestCheckResourceAttr("autobotai_evaluator.evaluator", "rules.1.value_source", "value"),
					resource.TestCheckResourceAttr("autobotai_evaluator.evaluator", "rules.1.value", ""),

					resource.TestCheckResourceAttr("autobotai_evaluator.evaluator", "combinator", "and"),
					resource.TestCheckResourceAttr("autobotai_evaluator.evaluator", "preference", "qb_rules"),
				),
			},
		},
	})
}

func testAccCheckEvaluatorConfigBasic(name, code, qbRulesId, preference string) string {
	return fmt.Sprintf(`
	resource "autobotai_evaluator" "evaluator" {
		name="%s"
		code="%s"
		qb_rules_id="%s"
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
		
		
		combinator="and"
		not=false
		preference="%s"
	
		}
		output "evaluatorRead"{
			value=autobotai_evaluator.evaluator
	}
	`,
		name, code, qbRulesId, preference)
}

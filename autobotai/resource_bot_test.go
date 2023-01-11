package autobotai

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"testing"
)

func TestAccResourceBot(t *testing.T) {
	name := "testing"
	topic := "testing"
	category := "Reliability"
	importance := "Medium"
	fetcherId := "63454da5a6ac05002369749412345"
	integrationId := "e9cb562a-1f12-42c1-8e68-586e140aa6c12345"
	integrationType := "azure"
	evaluatorId := "63b6bc1b0e5f0e6ae94dc5e11234"
	automationId := "63171d962268c207516223463456"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccAutobotAIPreCheck(t) },
		Providers: testAccAutobotAIProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckBotConfigBasic(name, topic, category, importance, fetcherId, integrationId, integrationType, automationId, evaluatorId),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("autobotai_bot.bot", "name", name),
					resource.TestCheckResourceAttr("autobotai_bot.bot", "topic", topic),
					resource.TestCheckResourceAttr("autobotai_bot.bot", "category", category),
					resource.TestCheckResourceAttr("autobotai_bot.bot", "importance", importance),
					resource.TestCheckResourceAttr("autobotai_bot.bot", "fetcher_id", fetcherId),
					resource.TestCheckResourceAttr("autobotai_bot.bot", "integration_id", integrationId),
					resource.TestCheckResourceAttr("autobotai_bot.bot", "integration_type", integrationType),

					resource.TestCheckResourceAttr("autobotai_bot.bot", "actions.0.type", "mutation"),
					resource.TestCheckResourceAttr("autobotai_bot.bot", "actions.0.automation_id", "63171d962268c207516223463456"),
					resource.TestCheckResourceAttr("autobotai_bot.bot", "actions.0.params.0.description", "Name of the bucket"),
					resource.TestCheckResourceAttr("autobotai_bot.bot", "actions.0.params.0.name", "bucket_name"),
					resource.TestCheckResourceAttr("autobotai_bot.bot", "actions.0.params.0.type", "string"),
					resource.TestCheckResourceAttr("autobotai_bot.bot", "actions.0.params.0.values", "autobot-bucket-automation-test"),
					resource.TestCheckResourceAttr("autobotai_bot.bot", "actions.0.action_id", ""),
					resource.TestCheckResourceAttr("autobotai_bot.bot", "evaluator_id", evaluatorId),
				),
			},
		},
	})
}

func testAccCheckBotConfigBasic(name, topic, category, importance, fetcherId, integrationId, integrationType, automationId, evaluatorId string) string {
	return fmt.Sprintf(`
	resource "autobotai_bot" "bot"{
		name="%s"
		topic="%s"
		category="%s"
		importance="%s"
		fetcher_id="%s"
		integration_id="%s"
		integration_type="%s"
		actions{
			type="mutation"
			automation_id="%s"
			params{
				description="Name of the bucket"
				name="bucket_name"
				type="string"
				values="autobot-bucket-automation-test" 
			}
			required=false
			action_id=""
		}
		evaluator_id="%s"
	}
	output "botRead"{
        value=autobotai_bot.bot
}
	`,
		name, topic, category, importance, fetcherId, integrationId, integrationType, automationId, evaluatorId)
}

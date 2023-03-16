package autobotai

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"testing"
)

func TestAccResourceInventorySchedule(t *testing.T) {
	integrationType := "azure"
	integrationId := "qwedwqdwq_sdnbnghncsdcdsdsc"
	cronExpression := "0 0 * * *"
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccAutobotAIPreCheck(t) },
		Providers: testAccAutobotAIProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckInventoryScheduleConfigBasic(integrationType, integrationId, cronExpression),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("autobotai_inventory_schedule.inventory", "integration_type", integrationType),
					resource.TestCheckResourceAttr("autobotai_inventory_schedule.inventory", "integration_id", integrationId),
					resource.TestCheckResourceAttr("autobotai_inventory_schedule.inventory", "cron_expression", cronExpression),
				),
			},
		},
	})
}

func testAccCheckInventoryScheduleConfigBasic(integrationType, integrationId, cronExpression string) string {
	return fmt.Sprintf(`
	resource "autobotai_inventory_schedule" "inventory" {
       
        integration_type="%s"
		integration_id="%s"
        cron_expression="%s"
	}
	`,
		integrationType, integrationId, cronExpression)
}

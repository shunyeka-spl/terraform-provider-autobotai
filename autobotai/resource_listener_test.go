package autobotai

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"testing"
)

func TestAccResourceListener(t *testing.T) {
	description := "testing"
	name := "testing"
	webhookSourceIntegrationType := "conformity"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccAutobotAIPreCheck(t) },
		Providers: testAccAutobotAIProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckListenerConfigBasic(description, name, webhookSourceIntegrationType),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("autobotai_listener.listener", "description", description),
					resource.TestCheckResourceAttr("autobotai_listener.listener", "name", name),
					resource.TestCheckResourceAttr("autobotai_listener.listener", "webhook_source_integration_type", webhookSourceIntegrationType),
				),
			},
		},
	})
}

func testAccCheckListenerConfigBasic(description, name, webhookSourceIntegrationType string) string {
	return fmt.Sprintf(`
		resource "autobotai_listener" "listener" {
			description = "%s"
			name = "%s"
			webhook_source_integration_type = "%s"
		}
		output "listenerRead"{
			value=autobotai_listener.listener
		}
	`,
		description, name, webhookSourceIntegrationType)
}

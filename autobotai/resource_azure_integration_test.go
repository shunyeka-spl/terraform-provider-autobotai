package autobotai

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"testing"
)

func TestAccResourceAzureIntegration(t *testing.T) {
	alias := "MyProject"
	client_id := "autobotai"
	tenant_id := "12233345687sr"
	secret_key := "112840099457dcdce455417995"
	subscription_id := "cewewwe"
	groups := "test"
	cspname := "azure"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccAutobotAIPreCheck(t) },
		Providers: testAccAutobotAIProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckAzureIntegrationConfigBasic(alias, client_id, tenant_id, secret_key, subscription_id, groups, cspname),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("autobotai_azure_integration.azure", "alias", alias),
					resource.TestCheckResourceAttr("autobotai_azure_integration.azure", "client_id", client_id),
					resource.TestCheckResourceAttr("autobotai_azure_integration.azure", "tenant_id", tenant_id),
					resource.TestCheckResourceAttr("autobotai_azure_integration.azure", "secret_key", secret_key),
					resource.TestCheckResourceAttr("autobotai_azure_integration.azure", "subscription_id", subscription_id),
					resource.TestCheckResourceAttr("autobotai_azure_integration.azure", "groups.0", groups),
					resource.TestCheckResourceAttr("autobotai_azure_integration.azure", "cspname", cspname),
				),
			},
		},
	})
}

func testAccCheckAzureIntegrationConfigBasic(alias, client_id, tenant_id, secret_key, subscription_id, groups, cspname string) string {
	return fmt.Sprintf(`
		resource "autobotai_azure_integration" "azure" {
			alias = "%s"
			client_id = "%s"
			tenant_id = "%s"
			secret_key = "%s"
			subscription_id = "%s"
			groups = ["%s"]
			cspname = "%s"  

		}
		output "azureRead"{
			value=autobotai_azure_integration.azure
		}
	`,
		alias, client_id, tenant_id, secret_key, subscription_id, groups, cspname)
}

package autobotai

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"testing"
)

func TestAccResourceFetcher(t *testing.T) {
	integrationType := "gcp"
	name := "testing-fetcher-gcp"
	clients := "gcp_audit_policy"
	code := ""
	types := "no_code"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccAutobotAIPreCheck(t) },
		Providers: testAccAutobotAIProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckFetcherConfigBasic(integrationType, clients, name, code, types),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("autobotai_fetcher.fetcher", "integration_type", integrationType),
					resource.TestCheckResourceAttr("autobotai_fetcher.fetcher", "name", name),
					resource.TestCheckResourceAttr("autobotai_fetcher.fetcher", "clients.0", clients),
					resource.TestCheckResourceAttr("autobotai_fetcher.fetcher", "code", code),
					resource.TestCheckResourceAttr("autobotai_fetcher.fetcher", "type", types),
				),
			},
		},
	})
}

func testAccCheckFetcherConfigBasic(integrationType, clients, name, code, types string) string {
	return fmt.Sprintf(`
	resource "autobotai_fetcher" "fetcher" {
        integration_type = "%s"
        clients =  ["%s"]
        name = "%s"
		 code  = "%s"
		 type = "%s"
		 data_keys=[{"name"= "service.gcp_audit_policy","type"= "str"},{"name"= "audit_log_configs.gcp_audit_policy","type"= "str"},{"name"= "akas.gcp_audit_policy","type"= "str"},{"name"= "location.gcp_audit_policy","type"= "str"},{"name"= "project.gcp_audit_policy","type"= "str"}]

	 }

		output "fetcherRead"{
			value=autobotai_fetcher.fetcher
		}
	`,
		integrationType, clients, name, code, types)
}

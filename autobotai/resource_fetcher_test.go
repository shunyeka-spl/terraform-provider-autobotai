package autobotai

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"testing"
)

func TestAccResourceFetcher(t *testing.T) {
	integrationType := "azure"
	name := "testing"
	clients := "NetworkManagementClient"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccAutobotAIPreCheck(t) },
		Providers: testAccAutobotAIProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckFetcherConfigBasic(integrationType, clients, name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("autobotai_fetcher.fetcher", "integration_type", integrationType),
					resource.TestCheckResourceAttr("autobotai_fetcher.fetcher", "name", name),
					resource.TestCheckResourceAttr("autobotai_fetcher.fetcher", "clients.0", clients),
					resource.TestCheckResourceAttr("autobotai_fetcher.fetcher", "code", "# 'clients' is a dict that contains all the clients mentioned in fetcher setup\n# 'test' is a boolean that could be used for conditional testing\n\ndef fetch(clients, test=False):\n    # To get the provided clients, proceed as per Example :\n    resource_client = clients['ResourceManagementClient']\n    # Retrieve the list of resource groups\n    resource_group_response = resource_client.resource_groups.list()\n    # <> Your Code goes here.\n\n    # resources=[]\n    # for resource_group in list(resource_group_response):\n    # # Retrieve the list of resources for each resource group\n    #     resource_list = resource_client.resources.list_by_resource_group(resource_group.name)\n    #     for resource in list(resource_list):\n    #         if resource.type=='Microsoft.Network/networkSecurityGroups':\n    #             resources.append({\n    #                 'id': resource.id,\n    #                 'name':resource.name,\n    #                 'type':resource.type,\n    #                 'region':resource.location,\n    #                 'resource_group_name':resource_group.name\n    #             })\n    # return resources\n"),
				),
			},
		},
	})
}

func testAccCheckFetcherConfigBasic(integrationType, clients, name string) string {
	return fmt.Sprintf(`
	resource "autobotai_fetcher" "fetcher" {
        integration_type = "%s"
        clients =  ["%s"]
        code  = "# 'clients' is a dict that contains all the clients mentioned in fetcher setup\n# 'test' is a boolean that could be used for conditional testing\n\ndef fetch(clients, test=False):\n    # To get the provided clients, proceed as per Example :\n    resource_client = clients['ResourceManagementClient']\n    # Retrieve the list of resource groups\n    resource_group_response = resource_client.resource_groups.list()\n    # <> Your Code goes here.\n\n    # resources=[]\n    # for resource_group in list(resource_group_response):\n    # # Retrieve the list of resources for each resource group\n    #     resource_list = resource_client.resources.list_by_resource_group(resource_group.name)\n    #     for resource in list(resource_list):\n    #         if resource.type=='Microsoft.Network/networkSecurityGroups':\n    #             resources.append({\n    #                 'id': resource.id,\n    #                 'name':resource.name,\n    #                 'type':resource.type,\n    #                 'region':resource.location,\n    #                 'resource_group_name':resource_group.name\n    #             })\n    # return resources\n"
        name = "%s"
	 }

		output "fetcherRead"{
			value=autobotai_fetcher.fetcher
		}
	`,
		integrationType, clients, name)
}

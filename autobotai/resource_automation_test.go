package autobotai

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"testing"
)

func TestAccResourceAutomation(t *testing.T) {
	integrationType := "azure"
	name := "testing"
	clients := "ComputeManagementClient"
	title := "testing"
	approvalRequired := "true"
	types := "communication"
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccAutobotAIPreCheck(t) },
		Providers: testAccAutobotAIProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckAutomationConfigBasic(name, clients, title, types, integrationType, approvalRequired),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("autobotai_automation.automation", "integration_type", integrationType),
					resource.TestCheckResourceAttr("autobotai_automation.automation", "name", name),
					resource.TestCheckResourceAttr("autobotai_automation.automation", "clients.0", clients),
					resource.TestCheckResourceAttr("autobotai_automation.automation", "code", "# This automation fetch all resources.\n\n# 'clients' is a dict that contains all the clients mentioned in automation setup\n# 'params' is a dict where key is mentioned in automation setup and value is provided in bot setup\n# 'resources' is a list of resources fetched by fetcher and filtered via evaluator\n# 'test' is a boolean that could be used for conditional testing\n\ndef execute(clients, params, resources=[], test=False):\n    # To get the provided clients, proceed as per Example :\n    resource_client = clients['ResourceManagementClient']\n    # Retrieve the list of resource groups\n    resource_group_response = resource_client.resource_groups.list()\n    # <> Your Code goes here.\n\n    # resources=[]\n    # for resource_group in list(resource_group_response):\n    # # Retrieve the list of resources for each resource group\n    #     resource_list = resource_client.resources.list_by_resource_group(resource_group.name)\n    #     for resource in list(resource_list):\n    #         if resource.type=='Microsoft.Network/networkSecurityGroups':\n    #             resources.append({\n    #                 'id': resource.id,\n    #                 'name':resource.name,\n    #                 'type':resource.type,\n    #                 'region':resource.location,\n    #                 'resource_group_name':resource_group.name\n    #             })\n    # return resources\n"),
					resource.TestCheckResourceAttr("autobotai_automation.automation", "title", title),
					resource.TestCheckResourceAttr("autobotai_automation.automation", "type", types),
					resource.TestCheckResourceAttr("autobotai_automation.automation", "approval_required", approvalRequired),
				),
			},
		},
	})
}

func testAccCheckAutomationConfigBasic(name, clients, title, types, integrationType, approvalRequired string) string {
	return fmt.Sprintf(`
	resource "autobotai_automation" "automation" {
        name            = "%s"
        clients       =  ["%s"]
        code= "# This automation fetch all resources.\n\n# 'clients' is a dict that contains all the clients mentioned in automation setup\n# 'params' is a dict where key is mentioned in automation setup and value is provided in bot setup\n# 'resources' is a list of resources fetched by fetcher and filtered via evaluator\n# 'test' is a boolean that could be used for conditional testing\n\ndef execute(clients, params, resources=[], test=False):\n    # To get the provided clients, proceed as per Example :\n    resource_client = clients['ResourceManagementClient']\n    # Retrieve the list of resource groups\n    resource_group_response = resource_client.resource_groups.list()\n    # <> Your Code goes here.\n\n    # resources=[]\n    # for resource_group in list(resource_group_response):\n    # # Retrieve the list of resources for each resource group\n    #     resource_list = resource_client.resources.list_by_resource_group(resource_group.name)\n    #     for resource in list(resource_list):\n    #         if resource.type=='Microsoft.Network/networkSecurityGroups':\n    #             resources.append({\n    #                 'id': resource.id,\n    #                 'name':resource.name,\n    #                 'type':resource.type,\n    #                 'region':resource.location,\n    #                 'resource_group_name':resource_group.name\n    #             })\n    # return resources\n"
        title ="%s"
        type="%s"
        integration_type= "%s"
        approval_required="%s"
	}

	`,
		name, clients, title, types, integrationType, approvalRequired)
}

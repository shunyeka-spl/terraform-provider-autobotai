resource "autobotai_fetcher" "fetcher" {
        integration_type            = "azure"
        clients       =  ["ComputeManagementClient","ResourceManagementClient"]
        code     = " # 'clients' is a dict that contains all the clients mentioned in fetcher setup\n# 'test' is a boolean that could be used for conditional testing\n\ndef fetch(clients, test=False):\n    # To get the provided clients, proceed as per Example :\n    resource_client = clients['ResourceManagementClient']\n    # Retrieve the list of resource groups\n    resource_group_response = resource_client.resource_groups.list()\n    # <> Your Code goes here.\n\n    # resources=[]\n    # for resource_group in list(resource_group_response):\n    # # Retrieve the list of resources for each resource group\n    #     resource_list = resource_client.resources.list_by_resource_group(resource_group.name)\n    #     for resource in list(resource_list):\n    #         if resource.type=='Microsoft.Network/networkSecurityGroups':\n    #             resources.append({\n    #                 'id': resource.id,\n    #                 'name':resource.name,\n    #                 'type':resource.type,\n    #                 'region':resource.location,\n    #                 'resource_group_name':resource_group.name\n    #             })\n    # return resources\n"
        name = "testings"
}

output "fetherRead"{
        value=autobotai_fetcher.fetcher
}



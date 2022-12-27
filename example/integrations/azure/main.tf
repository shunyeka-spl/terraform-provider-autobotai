resource "autobotai_azure_integration" "azure" {
        alias            = "MyProject"
        client_id       = "autobot647"
        tenant_id     = "12233345687sr"
        secret_key = "11284009fvfvfvgf9457dcdce4554"
        subscription_id = "cewewwe"
        groups=["test"]
        cspname="azure"
}
output "azureRead"{
    value=autobotai_azure_integration.azure
}


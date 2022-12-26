resource "autobot_azure_integration" "azure" {
        alias            = "MyProject"
        client_id       = "autobot647"
        tenant_id     = "12233345687sr"
        secret_key = "11284009fvfvfvgf9457dcdce455417995"
        subscription_id = "cewewwe"
        groups=["test"]
        cspname="azure"
}
output "azureRead"{
    value=autobot_azure_integration.azure
}


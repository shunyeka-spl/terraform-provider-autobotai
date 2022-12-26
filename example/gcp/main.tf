
resource "autobot_gcp_integration" "gcp"{
    type="service_account"
    project_id="conformity-346910"
    private_key_id="016b4789c9544b1abf8243c81e0271f826330a2f"
    private_key="-----BEGIN PRIVATE KEY-----\nprivateKey\n-----END PRIVATE KEY-----\n"
    client_email="client_email"
    client_id="client_id"
    auth_uri="auth_uri"
    token_uri="token_uri"
    auth_provider_x_cert_url="auth_provider_x_cert_url"
    client_x_cert_url="client_x_cert_url"
    alias="MyGCPProject1-testing"
    groups=["test"]
    
}

output "gcp"{
    value=autobot_gcp_integration.gcp
}
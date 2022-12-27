---
page_title: "autobotai_gcp_integration"
description: |-
  Provides an autobotAI Gcp Integration.
---

# Resource `autobotai_gcp_integration`
Provides an autobotAI Conformity Integration.

# Instruction on how to create a Gcp Integration

## Required things 
1. apikey
2. type,project_id,private_key_id,private_key,client_email,client_id,auth_uri,token_uri,auth_provider_x_cert_url,client_x_cert_url,alias,groups

## Steps 
1. Create terraform "main.tf" inside "example/integrations/gcp" folder for autobotAI Gcp Integration, create the file if not exists and add the following vallues:
    ## Example Usage 
    ```
    resource "autobot_gcp_integration" "gcp"{
    type="service_account"
    project_id="conformity-346910"
    private_key_id="123456gf78uhy"
    private_key="-----BEGIN PRIVATE KEY-----privatekey-----END PRIVATE KEY-----\n"
    client_email=client_email"
    client_id="123456789"
    auth_uri="auth_uri"
    token_uri="token_uri"
    auth_provider_x_cert_url="auth_provider_x_cert_url"
    client_x_cert_url="client_x_cert_url"
    alias="MyGCPProject1-testing"
    groups=["test"]

    }
    ```
2. To use the autobotAI Gcp Integration follow the "provider_setup.md"

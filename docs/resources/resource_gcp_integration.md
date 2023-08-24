# Guidance on Creating a GCP Integration

---
    page_title: "autobotai_gcp_integration"
    description: Enables the creation of an autobotAI GCP Integration
---

## Resource `autobotai_gcp_integration`
Discover how to set up an autobotAI GCP Integration


## Prerequisites
1. apikey and url
2. type,project_id,private_key_id,private_key,client_email,client_id,auth_uri,token_uri,auth_provider_x_cert_url,client_x_cert_url,alias,groups

## Steps 
1. Create terraform "main.tf" inside "example/integrations/gcp" folder for autobotAI GCP Integration. If the file doesn't exist, create it and input the following values:
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
2. For implementing autobotAI GCP Integration, refer to the instructions provided in the [autobotAI-Provider-Guidance](provider_guidance.md) document.

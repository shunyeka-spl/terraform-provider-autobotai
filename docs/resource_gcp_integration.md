---
page_title: "autobot_gcp_integration"
description: |-
  Provides a Autobot Gcp Integration.
---

# Resource `autobot_gcp_integration`
Provides a Autobot Conformity Integration.

# Instruction on how to create a Gcp Integration

## Required things 
1. apikey
2. type,project_id,private_key_id,private_key,client_email,client_id,auth_uri,token_uri,auth_provider_x_cert_url,client_x_cert_url,alias,groups

## Steps 
1. Create terraform "main.tf" inside "example/gcp" folder for Autobot Gcp Integration, create the file if not exists and add the following vallues:
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
2. To use the Autobot Gcp Integration add the API key to the varibale.tf 
3. create "variable.tf" inside "example/gcp" folder and add the following values:
    ## Example Usage 
    ```
    variable "apikey"{
    type    = string
    default = "API Key"
    }

    ```
Note: You can always change the values declared according to your choice.
4. Run Terraform 
    1. Navigate to the Project Directory:
        ```
        cd /path/autobot_integration

        ``` 
    2. Create the Artifact:
        ```
        make install

        ```
    3. Now, you can run the terraform code by navigating to "example/gcp" folder:
        ```
        cd example/gcp
        terraform init
        terraform apply

        ```
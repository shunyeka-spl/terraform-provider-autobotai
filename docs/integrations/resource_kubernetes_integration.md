# Instruction on how to create a Kubernetes Integration

---
page_title: "autobotai_kubernetes_integration"
description: |-
  Provides an autobotAI Kubernetes Integration.
---

# Resource `autobotai_kubernetes_integration`
Provides an autobotAI Kubernetes Integration.


## Required things 
1. apikey and url
2. alias,groups,deploy_bots
3. name (for agent)

## Steps 
1. Create terraform "main.tf" inside "example/integrations/kubernetes" folder for autobotAI kubernetes Integration, create the file if not exists and add the following vallues:
    ## Example Usage 
    ```
        resource "autobotai_kubernetes_integration" "kubernetes" {
            alias            = "MyProjecttesting"
            deploy_bots       = true
            groups=["k8s"]
        }

        resource "autobotai_agent" "agent" {
            integration_id= autobotai_kubernetes_integration.kubernetes.id
            integration_type="kubernetes"
            name= "terraformAgentTesting"
        }

        resource "autobotai_apikey" "apikey" {
            name = autobotai_agent.agent.name
        }
    ```
2. To use the autobotAI kubernetes Integration follow the "provider_setup.md"
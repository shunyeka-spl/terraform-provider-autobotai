# Guidance on Creating a Kubernetes Integration

---
    page_title: "autobotai_kubernetes_integration"
    description: Enables the creation of an autobotAI Kubernetes  Integration
---

## Resource `autobotai_kubernetes_integration`
Discover how to set up an autobotAI Kubernetes Integration


## Prerequisites 
1. apikey and url
2. alias,groups,deploy_bots
3. name (for agent)

## Steps 
1. Create terraform "main.tf" inside "example/integrations/kubernetes" folder for autobotAI kubernetes Integration. If the file doesn't exist, create it and input the following values:
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
2. For implementing autobotAI Kubernetes Integration, refer to the instructions provided in the [autobotAI-Provider-Guidance](provider_guidance.md) document.
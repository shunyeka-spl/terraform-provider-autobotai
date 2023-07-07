
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

output "Agent_Installation"{
    value= "kubectl create secret generic abai-secrets --from-literal=apiKey=${autobotai_apikey.apikey.api_key_id}:${autobotai_apikey.apikey.api_key_secret} --from-literal=agentId=${autobotai_agent.agent.id} --from-literal=apiUrl=${var.url}"
}

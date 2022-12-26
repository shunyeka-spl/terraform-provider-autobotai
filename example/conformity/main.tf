resource "autobot_conformity_integration" "conformity" {
    alias            = "Conformity"
    groups=["test"]
    api_key="test@123"
    region=["ap-south-1"]
}

output "conformity"{
    value=autobot_conformity_integration.conformity
}
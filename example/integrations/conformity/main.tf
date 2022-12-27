resource "autobotai_conformity_integration" "conformity" {
    alias            = "Conformity"
    groups=["test"]
    api_key="test@123"
    region=["ap-south-1"]
}

output "conformity"{
    value=autobotai_conformity_integration.conformity
}
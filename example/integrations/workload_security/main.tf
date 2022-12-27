resource "autobotai_workload_security_integration" "wl_security" {
    alias            = "MySecurityTesting"
    groups=["test"]
    api_version="v1"
    secretkey="scdscdcdcdcdcd"
    url="url"
}

output "workload_security"{
    value=autobotai_workload_security_integration.wl_security
}
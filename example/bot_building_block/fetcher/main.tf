resource "autobotai_fetcher" "fetcher" {
        integration_type            = "gcp"
        clients       =  ["gcp_audit_policy"]
        code     = " "
        name = "testings"
        type="no_code"
        data_keys=[
                {
            "name"= "service.gcp_audit_policy",
            "type"= "str"
        },
        {
            "name"= "audit_log_configs.gcp_audit_policy",
            "type"= "str"
        },
        {
            "name"= "akas.gcp_audit_policy",
            "type"= "str"
        },
        {
            "name"= "location.gcp_audit_policy",
            "type"= "str"
        },
        {
            "name"= "project.gcp_audit_policy",
            "type"= "str"
        }
        ]
}

output "fetherRead"{
        value=autobotai_fetcher.fetcher
}



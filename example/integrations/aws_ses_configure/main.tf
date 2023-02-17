resource "autobotai_aws_ses_configure_integration" "aws_ses" {
        alias            = "MyProjecttesting"
        region       = "us-east-1"
        integration_id     = "16548497741465"
        
}
output "awsSesRead"{
    value=autobotai_aws_ses_configure_integration.aws_ses
}


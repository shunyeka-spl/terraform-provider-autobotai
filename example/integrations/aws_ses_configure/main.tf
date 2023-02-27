resource "autobotai_aws_ses_configure_integration" "aws_ses" {
        alias            = "MyProjecttesting"
        region       = "us-east-1"
        dependson="789465414178"
        
}
output "awsSesRead"{
    value=autobotai_aws_ses_configure_integration.aws_ses
}


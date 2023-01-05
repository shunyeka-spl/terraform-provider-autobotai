resource "autobotai_external_id" "aws"{
    alias="My-Project-Testing"
    groups=["test"]
}

output "aws_details"{
    value=autobotai_external_id.aws.external_id
}


 resource "aws_cloudformation_stack" "autobotai_cloud_formation" {
  name         = "autobotAIAWS"
  template_url = "template_url"
  parameters = {
      APIURL="https://api-test.autobot.live/"
      Environment="test"
      TrustPrincipal  = autobotai_external_id.aws.target_principal
      ExternalID = autobotai_external_id.aws.external_id
  }
  capabilities = ["CAPABILITY_NAMED_IAM"]
}

data "aws_caller_identity" "current" {}


resource "autobotai_aws_integration" "aws_integration" {
   stack_id=aws_cloudformation_stack.autobotai_cloud_formation.id
   external_id = autobotai_external_id.aws.external_id
   account_id = data.aws_caller_identity.current.account_id 
}

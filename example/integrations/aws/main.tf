data "autobotai_external_id" "aws"{
    alias="MyAWSProject"
    groups=["test"]
}

output "aws_details"{
    value=data.autobot_external_id.aws.external_id
}

 resource "aws_cloudformation_stack" "autobotai_cloud_formation" {
  name         = "AutobotAIAWS"
  template_url = "template_url"
  parameters = {
      APIURL="https://api-test.autobot.live/"
      Environment="test"
      TrustPrincipal  = data.autobotai_external_id.aws.target_principal
      ExternalID = data.autobotai_external_id.aws.external_id
  }
  capabilities = ["CAPABILITY_NAMED_IAM"]
}

data "autobot_external_id" "aws"{
    alias="MyAWSProject"
    groups=["test"]
}

output "aws_details"{
    value=data.autobot_external_id.aws.external_id
}

 resource "aws_cloudformation_stack" "autobot_cloud_formation" {
  name         = "AutobotAWS"
  template_url = "template_url"
  parameters = {
      APIURL="https://api-test.autobot.live/"
      Environment="test"
      TrustPrincipal  = data.autobot_external_id.aws.target_principal
      ExternalID = data.autobot_external_id.aws.external_id
  }
  capabilities = ["CAPABILITY_NAMED_IAM"]
}

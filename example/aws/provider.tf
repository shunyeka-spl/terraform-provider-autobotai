
terraform {
  required_providers {
    autobot = {
      source  = "shunyeka/autobot/autobot"
    }
    aws = {
      source  = "hashicorp/aws"
      version = "~> 3.0"
    }
  }
}
provider "autobot" {
  apikey = var.apikey
}

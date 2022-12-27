
terraform {
  required_providers {
    autobotai = {
      source  = "shunyeka-spl/autobotai/autobotai"
    }
    aws = {
      source  = "hashicorp/aws"
      version = "~> 3.0"
    }
  }
}
provider "autobotai" {
  apikey = var.apikey
  url =var.url
}

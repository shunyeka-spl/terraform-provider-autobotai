
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
provider "aws" {
  access_key = "access_key"
  secret_key = "secret_key"
  region = "region"
  
}
provider "autobotai" {
  apikey = var.apikey
  url =var.url
}

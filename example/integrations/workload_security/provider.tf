terraform {
  required_providers {
    autobotai = {
      source  = "shunyeka-spl/autobotai/autobotai"
    }
  }
}

provider "autobotai" {
  apikey = var.apikey
  url = var.url
}

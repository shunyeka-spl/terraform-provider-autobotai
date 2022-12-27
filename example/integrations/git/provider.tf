terraform {
  required_providers {
    autobotai = {
      source  = "shunyeka-spl/autobotAI/autobotAI"
    }
  }
}

provider "autobotai" {
  apikey = var.apikey
  url = var.url
}

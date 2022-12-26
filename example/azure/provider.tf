terraform {
  required_providers {
    autobotai = {
      source  = "shunyeka/autobot/autobot"
    }
  }
}

provider "autobotai" {
  apikey = var.apikey
}

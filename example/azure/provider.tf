terraform {
  required_providers {
    autobot = {
      source  = "shunyeka/autobot/autobot"
    }
  }
}

provider "autobot" {
  apikey = var.apikey
}

terraform {
  required_providers {
    ohdear = {
      source  = "bax-energy/ohdear"
      version = "0.0.10"
    }
  }
}

provider "ohdear" {} # needs : OHDEAR_APY_KEY and OHDEAR_TEAM_ID

provider "ohdear" {
  api_key = var.api_key # or use environment variable OHDEAR_APY_KEY
  team_id = var.team_id # or use environment variable OHDEAR_TEAM_ID
}
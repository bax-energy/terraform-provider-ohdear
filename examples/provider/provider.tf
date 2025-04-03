terraform {
  required_providers {
    ohdear = {
      source  = "app.terraform.io/bax-energy/ohdear"
      version = "0.1.0"
    }
  }
}

provider "ohdear" {} # needs : OHDEAR_TEAM_ID and OHDEAR_TEAM_ID

provider "ohdear" {
  api_key = var.api_key # or use environment variable OHDEAR_APY_KEY
  team_id = var.team_id # or use environment variable OHDEAR_TEAM_ID
  api_url = var.api_url # Optional: or use environment variable OHDEAR_API_URL
}

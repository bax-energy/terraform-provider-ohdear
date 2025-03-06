terraform {
  required_providers {
    farsight = {
      source  = "example.com/local/farsight"
      version = "0.1.0"
    }
  }
}

provider "farsight" {
  api_key = var.api_key
  team_id = var.team_id
}


resource "farsight_ohdear_site" "test_code" {
  url     = "https://farsight-otel-dev.bax.energy/v1/logs"
  friendly_name = "Dev - Otel"

  checks {
    uptime        = true
  }
  
  uptime {
    check_valid_status_codes = ["2*","405"]

  }
  tags = [ "falsight" , "dev" ]
  
}

resource "farsight_ohdear_site" "test_code2" {
  url     = "https://api-dev.bax.energy/swagger/index.html"
  friendly_name = "Dev - api public"

  checks {
    uptime        = true
  }
  
  uptime {
    check_max_redirect_count = 2
  }
  tags = [ "falsight" , "dev" ]
  
}
terraform {
  required_providers {
    farsight = {
      source  = "bax-energy/farsight"
      version = "0.1.0"
    }
  }
}

provider "farsight" {} # needs : OHDEAR_TEAM_ID and OHDEAR_TEAM_ID

provider "farsight" {
  api_key = var.api_key # or use environment variable OHDEAR_APY_KEY
  team_id = var.team_id #or use environment variable OHDEAR_TEAM_ID
  api_url = var.api_url #or use environment variable OHDEAR_API_URL
}


resource "farsight_ohdear_site" "example" {
  url     = "https://yoururl.com"

  #all checks are enabled by default
}

resource "farsight_ohdear_site" "example2" {
  url     = "https://api-dev.bax.energy/swagger/index.html"
  friendly_name = "Dev - api public"

  #specify which checks to enable
  checks {
    uptime        = true
  }

  #uptime options
  uptime {
    check_max_redirect_count = 2
    
  }
  tags = [ "falsight" , "dev" ]
  
}
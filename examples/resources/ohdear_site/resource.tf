resource "ohdear_site" "example" {
  url = "https://yoururl.com"

  #all checks are enabled by default
}

resource "ohdear_site" "example2" {
  url           = "https://yoururl.com"
  friendly_name = "display name on ohdear"

  #specify which checks to enable
  checks {
    uptime = true
  }

  #uptime options
  uptime {
    check_max_redirect_count = 2

  }
  tags = ["tag1", "tag2"]

}
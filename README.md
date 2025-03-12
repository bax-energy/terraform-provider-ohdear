# How to use this provider

To install this provider, copy and paste this code into your Terraform configuration. Then, run terraform init.act with the Oh Dear! platform directly from Terraform.

### Terraform 0.13+


   ```hcl
   terraform {
     required_providers {
       ohdear = {
         source = "app.terraform.io/bax-energy/ohdear"
         version = "0.0.5"
       }
     }
   }
   ```


### Example Usage

```hcl
provider "ohdear" {
  api_token = var.token   # optionally use OHDEAR_TOKEN env var
  api_url   = var.api_url # optionally use OHDEAR_API_URL env var
  team_id   = var.team_id # optionally use OHDEAR_TEAM_ID env var
}
```

### Schema

#### Required
- **api_key** (String) Oh Dear API token. If not set, uses OHDEAR_TOKEN env var.
- **team_id** (String) The default team ID to use for sitesIf not set, uses OHDEAR_TEAM_ID env var.

#### Optional
- **api_url** (String) Oh Dear API URL. If not set, uses OHDEAR_API_URL env var. Defaults to ```https://ohdear.app```.


# ohdear_site (Resource)

ohdear_site manages a site in Oh Dear via [their API](https://ohdear.app/docs/integrations/the-oh-dear-api#sites).

Example Usage


```hcl
resource "ohdear_site" "example" {
  url     = "https://yoururl.com"

  # all checks are enabled by default
}


resource "ohdear_site" "example2" {
  url           = "https://yoururl.com"
  friendly_name = "display name on ohdear"

  # specify which checks to enable
  checks {
    uptime = true
  }

  # uptime options
  uptime {
    check_max_redirect_count = 2
  }
  tags = ["tag1", "tag2"]
}
```

### Schema

#### Required

- `url` (String) - The URL of the site to be monitored.

#### Optional

- `team_id` (String) - The ID of the team that owns the site.
- `friendly_name` (String) - If you specify a friendly name, we'll display this instead of the URL.
- `tags` (List of String) - We'll display these tags across our UI and will send them along when requesting sites via the API.
- `checks` (Block List, Max: 1) - The list of checks to be performed on the site. If block is not present, it will enable all checks. (see below for nested schema)
  - `uptime` (Boolean) - Check if the site is up and running.
  - `performance` (Boolean) - Check the performance of the site.
  - `broken_links` (Boolean) - Check for broken links on the site.
  - `mixed_content` (Boolean) - Check for mixed content on the site.
  - `lighthouse` (Boolean) - Run Lighthouse checks on the site.
  - `cron` (Boolean) - Check the cron jobs of the site.
  - `application_health` (Boolean) - Check the health of the application running on the site.
  - `sitemap` (Boolean) - Check the sitemap of the site.
  - `dns` (Boolean) - Check the DNS configuration of the site.
  - `domain` (Boolean) - Check the domain configuration of the site.
  - `certificate_health` (Boolean) - Check the health of the SSL certificate of the site.
  - `certificate_transparency` (Boolean) - Check the certificate transparency logs of the site.
- `uptime` (Block List, Max: 1) - Uptime check configuration. (see below for nested schema)
  - `check_valid_status_codes` (List of String) - A list of valid status codes for the uptime check. You can specify a comma-separated list and use wildcards. '2*' means everything in the 200 range.
  - `http_client_headers` (List of Object) - A list of HTTP client headers to be sent with the requests.
    - `name` (String) - The name of the HTTP header.
    - `value` (String) - The value of the HTTP header.
  - `check_location` (String) - We can check your server from all over the world. Default: "paris".
  - `check_failed_notification_threshold` (Number) - The threshold for failed notifications. Minutes. Default: 2.
  - `check_http_verb` (String) - The HTTP verb to use for the check. Values: GET, POST, PUT, PATCH. Default: "get".
  - `check_timeout` (Number) - The timeout for the check. Seconds. Default: 5.
  - `check_max_redirect_count` (Number) - The maximum number of redirects to follow. Default: 5.
  - `check_payload` (List of Object) - The payload to send with the check.
    - `name` (String) - The name of the payload.
    - `value` (String) - The value of the payload.
  - `check_look_for_string` (String) - Verify text on response. The string to look for in the response. Default: "".
  - `check_absent_string` (String) - Verify absence of text on response. The string that should be absent in the response. Default: "".
  - `check_expected_response_headers` (List of Object) - Verify headers on response. The expected response headers.
    - `name` (String) - The name of the response header.
    - `condition` (String) - The condition to check for the response header. Values: contains, not contains, equals, matches pattern.
    - `value` (String) - The value of the response header.
- `broken_links` (Block List, Max: 1) - Broken links configuration. (see below for nested schema)
  - `crawler_headers` (List of Object) - A list of HTTP client headers to be sent with the requests.
    - `name` (String) - The name of the HTTP header.
    - `value` (String) - The value of the HTTP header.

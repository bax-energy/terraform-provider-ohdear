package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccOhdearSite_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccOhdearSiteConfigBasic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("ohdear_site.test", "url", "https://example.com"),
					resource.TestCheckResourceAttr("ohdear_site.test", "friendly_name", "Example Site"),
				),
			},
			// {
			// 	ResourceName:      "ohdear_site.test",
			// 	ImportState:       true,
			// 	ImportStateVerify: true,
			// },
			// {
			// 	Config: "",
			// 	Check:  resource.ComposeTestCheckFunc(),
			// },
		},
	})
}

func testAccOhdearSiteConfigBasic() string {
	return `
resource "ohdear_site" "test" {
  url           = "https://example.com"
  friendly_name = "Example Site"
  checks {
    uptime = true
  }
  tags = ["test", "example"]
}`
}

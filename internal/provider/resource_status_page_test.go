package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccOhdearStatusPage_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccOhdearStatusPageConfigBasic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("ohdear_status_page.test", "title", "Example Status Page"),
				),
			},
			{
				ResourceName:      "ohdear_status_page.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccOhdearStatusPageConfigBasic() string {
	return `
resource "ohdear_status_page" "test" {
  title = "Example Status Page"
  sites = [
    {
      id        = "12345"
      clickable = true
    }
  ]
}`
}

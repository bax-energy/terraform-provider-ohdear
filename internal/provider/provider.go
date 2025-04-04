package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type Config struct {
	APIKey string
	APIURL string
	teamID string
}

func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"api_key": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("OHDEAR_TOKEN", nil),
				Description: "Oh Dear API token",
			},
			"api_url": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("OHDEAR_API_URL", "https://ohdear.app"),
				Description: "Oh Dear API URL",
			},
			"team_id": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("OHDEAR_TEAM_ID", nil),
				Description: "The default team ID to use for sites",
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"ohdear_site": resourceSite(),
		},
		ConfigureContextFunc: providerConfigure,
	}
}

func providerConfigure(_ context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	return &Config{
		APIKey: d.Get("api_key").(string),
		APIURL: d.Get("api_url").(string),
		teamID: d.Get("team_id").(string),
	}, nil
}

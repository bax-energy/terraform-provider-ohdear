package provider

import (
	"context"
	"fmt"
	"runtime"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func init() {
	schema.DescriptionKind = schema.StringMarkdown

	// add defaults on to the exported descriptions if present
	schema.SchemaDescriptionBuilder = func(s *schema.Schema) string {
		desc := s.Description
		if s.Default != nil {
			desc += fmt.Sprintf(" Defaults to `%v`.", s.Default)
		}
		if s.Deprecated != "" {
			desc += " __Deprecated__: " + s.Deprecated
		}
		return strings.TrimSpace(desc)
	}
}

func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"api_key": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("OHDEAR_APY_KEY", nil),
				Description: "Oh Dear API token. If not set, uses `OHDEAR_APY_KEY` env var",
			},
			"api_url": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("OHDEAR_API_URL", "https://ohdear.app"),
				Description: "Oh Dear API URL. If not set, uses `OHDEAR_API_URL` env var. Defaults to `https://ohdear.app`.",
			},
			"team_id": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("OHDEAR_TEAM_ID", nil),
				Description: "The default team ID to use for sites. If not set, uses `OHDEAR_TEAM_ID` env var.",
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"ohdear_site": resourceSite(),
		},
		ConfigureContextFunc: providerConfigure,
	}
}

func providerConfigure(_ context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	ua := fmt.Sprintf(
		"terraform-provider-ohdear/%s (https://github.com/bax-energy/terraform-provider-ohdear)",
		runtime.Version(),
	)

	apiURL, ok := d.Get("api_url").(string)
	if !ok || apiURL == "" {
		return nil, diag.Errorf("api_url must be a non-empty string")
	}

	apiKey, ok := d.Get("api_key").(string)
	if !ok || apiKey == "" {
		return nil, diag.Errorf("api_key must be a non-empty string")
	}

	teamID, ok := d.Get("team_id").(string)
	if !ok || teamID == "" {
		return nil, diag.Errorf("team_id must be a non-empty string")
	}

	client := NewClient(apiURL, apiKey)
	client.SetUserAgent(ua)

	return &Config{
		client: client,
		teamID: teamID,
	}, nil
}

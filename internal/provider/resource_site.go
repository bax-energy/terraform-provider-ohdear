package provider

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSite() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceSiteCreate,
		ReadContext:   resourceSiteRead,
		UpdateContext: resourceSiteUpdate,
		DeleteContext: resourceSiteDelete,

		Schema: map[string]*schema.Schema{
			"url": {
				Type:         schema.TypeString,
				Required:     true,
				Description:  "The URL of the site to be monitored.",
				ForceNew:     true,
				ValidateFunc: validation.IsURLWithHTTPorHTTPS,
			},
			"team_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				ForceNew:    true,
				Description: "The ID of the team that owns the site.",
			},
			"friendly_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "If you specify a friendly name we'll display this instead of the url.",
			},
			"tags": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "We'll display these tags across our UI and will send them along when requesting sites via the API.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"checks": {
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
				Computed:    true,
				Description: "The list of checks to be performed on the site.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"uptime": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Check if the site is up and running.",
						},
						"performance": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Check the performance of the site.",
						},
						"broken_links": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Check for broken links on the site.",
						},
						"mixed_content": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Check for mixed content on the site.",
						},
						"lighthouse": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Run Lighthouse checks on the site.",
						},
						"cron": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Check the cron jobs of the site.",
						},
						"application_health": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Check the health of the application running on the site.",
						},
						"sitemap": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Check the sitemap of the site.",
						},
						"dns": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Check the DNS configuration of the site.",
						},
						"domain": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Check the domain configuration of the site.",
						},
						"certificate_health": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Check the health of the SSL certificate of the site.",
						},
						"certificate_transparency": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Check the certificate transparency logs of the site.",
						},
					},
				},
			},
			"uptime": {
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
				Description: "Uptime check configuration.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"check_valid_status_codes": {
							Type:        schema.TypeList,
							Optional:    true,
							Description: "A list of valid status codes for the uptime check. You can specify a comma separated list and use wildcards. '2*' means everything in the 200 range.",
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"http_client_headers": {
							Type:        schema.TypeList,
							Optional:    true,
							Description: "A list of HTTP client headers to be sent with the requests.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": {
										Type:        schema.TypeString,
										Required:    true,
										Description: "The name of the HTTP header.",
									},
									"value": {
										Type:        schema.TypeString,
										Required:    true,
										Description: "The value of the HTTP header.",
									},
								},
							},
						},
						"check_location": {
							Type:        schema.TypeString,
							Optional:    true,
							Default:     "paris",
							Description: "We can check your server from all over the world.",
						},
						"check_failed_notification_threshold": {
							Type:        schema.TypeInt,
							Optional:    true,
							Default:     2,
							Description: "The threshold for failed notifications. Minutes",
						},
						"check_http_verb": {
							Type:        schema.TypeString,
							Optional:    true,
							Default:     "get",
							Description: "The HTTP verb to use for the check. Values: GET, POST, PUT, PATCH",
						},
						"check_timeout": {
							Type:        schema.TypeInt,
							Optional:    true,
							Default:     5,
							Description: "The timeout for the check. Seconds",
						},
						"check_max_redirect_count": {
							Type:        schema.TypeInt,
							Optional:    true,
							Default:     5,
							Description: "The maximum number of redirects to follow.",
						},
						"check_payload": {
							Type:        schema.TypeList,
							Optional:    true,
							Description: "The payload to send with the check.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": {
										Type:        schema.TypeString,
										Required:    true,
										Description: "The name of the payload.",
									},
									"value": {
										Type:        schema.TypeString,
										Required:    true,
										Description: "The value of the payload.",
									},
								},
							},
						},
						"check_look_for_string": {
							Type:        schema.TypeString,
							Optional:    true,
							Default:     "",
							Description: "Verify text on response. The string to look for in the response.",
						},
						"check_absent_string": {
							Type:        schema.TypeString,
							Optional:    true,
							Default:     "",
							Description: "Verify absence of text on response. The string that should be absent in the response.",
						},
						"check_expected_response_headers": {
							Type:        schema.TypeList,
							Optional:    true,
							Description: "Verify headers on response. The expected response headers.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": {
										Type:        schema.TypeString,
										Required:    true,
										Description: "The name of the response header.",
									},
									"condition": {
										Type:        schema.TypeString,
										Required:    true,
										Description: "The condition to check for the response header. Values: contains,not contains,equals,matches pattern",
									},
									"value": {
										Type:        schema.TypeString,
										Required:    true,
										Description: "The value of the response header.",
									},
								},
							},
						},
					},
				},
			},
			"broken_links": {
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
				Description: "broken_links configuration.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"crawler_headers": {
							Type:        schema.TypeList,
							Optional:    true,
							Description: "A list of HTTP client headers to be sent with the requests.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": {
										Type:        schema.TypeString,
										Required:    true,
										Description: "The name of the HTTP header.",
									},
									"value": {
										Type:        schema.TypeString,
										Required:    true,
										Description: "The value of the HTTP header.",
									},
								},
							},
						},
					},
				},
			},
			"application_health": {
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
				Description: "application_health configuration.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"check_result_url": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "The URL to check the application health.",
						},
						"secret": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "the secret to use for the application health check.",
						},
						"headers": {
							Type:        schema.TypeList,
							Optional:    true,
							Description: "A list of HTTP client headers to be sent with the requests.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": {
										Type:        schema.TypeString,
										Required:    true,
										Description: "The name of the HTTP header.",
									},
									"value": {
										Type:        schema.TypeString,
										Required:    true,
										Description: "The value of the HTTP header.",
									},
								},
							},
						},
					},
				},
			},
		},
		CustomizeDiff: resourceOhdearSiteDiff,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}
func resourceOhdearSiteDiff(_ context.Context, d *schema.ResourceDiff, m interface{}) error {
	checks := d.Get("checks").([]interface{})
	if len(checks) == 0 {
		isHTTPS := strings.HasPrefix(d.Get("url").(string), "https")
		checks = append(checks, map[string]bool{
			"uptime":                   true,
			"broken_links":             true,
			"performance":              true,
			"mixed_content":            isHTTPS,
			"lighthouse":               true,
			"cron":                     true,
			"application_health":       true,
			"sitemap":                  true,
			"dns":                      true,
			"domain":                   true,
			"certificate_health":       isHTTPS,
			"certificate_transparency": isHTTPS,
		})

		if err := d.SetNew("checks", checks); err != nil {
			return err
		}
	}
	// set team_id from provider default if not provided
	if v, ok := d.GetOk("team_id"); !ok || v.(string) == "" {
		if err := d.SetNew("team_id", m.(*Config).teamID); err != nil {
			return err
		}
	}

	return nil
}
func resourceSiteCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	client := m.(*Config).client

	payload := BuildPayload(d, "create")

	site, err := client.AddSite(payload)
	if err != nil {
		return diagErrorf(err, "Could not add site to Oh Dear")
	}

	d.SetId(fmt.Sprintf("%d", site.ID))

	return resourceSiteRead(ctx, d, m)

}

func resourceSiteRead(_ context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Implementa la logica di lettura della risorsa qui
	return diag.Diagnostics{}
}

func resourceSiteUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	client := m.(*Config).client

	payload := BuildPayload(d, "update")

	site, err := client.UpdateSite(d.Id(), payload)
	if err != nil {
		return diagErrorf(err, "Could not add site to Oh Dear")
	}

	d.SetId(fmt.Sprintf("%d", site.ID))

	return resourceSiteRead(ctx, d, m)
}

func getSiteID(d *schema.ResourceData) (int, error) {
	id, err := strconv.Atoi(d.Id())
	if err != nil {
		return id, fmt.Errorf("corrupted resource ID in terraform state, Oh Dear only supports integer IDs. Err: %w", err)
	}
	return id, err
}

func resourceSiteDelete(_ context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	id, err := getSiteID(d)
	if err != nil {
		return diag.FromErr(err)
	}

	client := m.(*Config).client
	if err = client.RemoveSite(id); err != nil {
		return diagErrorf(err, "Could not remove site %d from Oh Dear", id)
	}

	return nil
}

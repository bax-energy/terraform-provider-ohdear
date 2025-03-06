package provider

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceOhDearSite() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceOhDearSiteCreate,
		ReadContext:   resourceOhDearSiteRead,
		UpdateContext: resourceOhDearSiteUpdate,
		DeleteContext: resourceOhDearSiteDelete,

		Schema: map[string]*schema.Schema{
			"url": {
				Type:         schema.TypeString,
				Required:     true,
				Description:  "The URL of the site to be monitored.",
				ValidateFunc: validation.IsURLWithHTTPorHTTPS,
			},
			"team_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The ID of the team that owns the site.",
			},
			"checks": {
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
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
		},
	}
}

func resourceOhDearSiteCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	payload := map[string]any{
		"url": d.Get("url").(string),
	}

	d.Set("team_id", m.(*Config).teamID)
	teamID, err := strconv.Atoi(m.(*Config).teamID)
	if err != nil {
		return diag.FromErr(fmt.Errorf("failed to convert team_id to integer: %s", err))
	}
	payload["team_id"] = teamID

	payload["checks"] = buildChecksPayload(d)

	if v, ok := d.GetOk("friendly_name"); ok {
		payload["friendly_name"] = v.(string)
	}
	if v, ok := d.GetOk("tags"); ok {
		payload["tags"] = v.([]interface{})
	}

	if uptimeConfig, ok := d.GetOk("uptime"); ok {
		uptimeList := uptimeConfig.([]interface{})
		if len(uptimeList) > 0 {
			uptimeMap := uptimeList[0].(map[string]any)
			if v, ok := uptimeMap["check_valid_status_codes"]; ok {
				payload["uptime_check_valid_status_codes"] = v.([]any)
			} else {
				payload["uptime_check_valid_status_codes"] = []interface{}{"2*"}
			}
			if v, ok := uptimeMap["http_client_headers"]; ok {
				headers := v.([]interface{})
				headerPayload := []map[string]string{}
				for _, header := range headers {
					headerMap := header.(map[string]interface{})
					headerPayload = append(headerPayload, map[string]string{
						"name":  headerMap["name"].(string),
						"value": headerMap["value"].(string),
					})
				}
				payload["http_client_headers"] = headerPayload
			}
			if v, ok := uptimeMap["check_location"]; ok {
				payload["uptime_check_location"] = v.(string)
			}
			if v, ok := uptimeMap["check_failed_notification_threshold"]; ok {
				payload["uptime_check_failed_notification_threshold"] = v.(int)
			}
			if v, ok := uptimeMap["check_http_verb"]; ok {
				payload["uptime_check_http_verb"] = v.(string)
			}
			if v, ok := uptimeMap["check_timeout"]; ok {
				payload["uptime_check_timeout"] = v.(int)
			}
			if v, ok := uptimeMap["check_max_redirect_count"]; ok {
				payload["uptime_check_max_redirect_count"] = v.(int)
			}
			if v, ok := uptimeMap["check_payload"]; ok {
				payload["uptime_check_payload"] = v.([]interface{})
			}
			if v, ok := uptimeMap["check_look_for_string"]; ok {
				payload["uptime_check_look_for_string"] = v.(string)
			}
			if v, ok := uptimeMap["check_absent_string"]; ok {
				payload["uptime_check_absent_string"] = v.(string)
			}
			if v, ok := uptimeMap["check_expected_response_headers"]; ok {
				headers := v.([]interface{})
				headerPayload := []map[string]string{}
				for _, header := range headers {
					headerMap := header.(map[string]interface{})
					headerPayload = append(headerPayload, map[string]string{
						"name":      headerMap["name"].(string),
						"condition": headerMap["condition"].(string),
						"value":     headerMap["value"].(string),
					})
				}
				payload["uptime_check_expected_response_headers"] = headerPayload
			}
		}
	}

	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return diag.FromErr(fmt.Errorf("failed to marshal payload: %s", err))
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/api/sites", m.(*Config).APIURL), bytes.NewBuffer(jsonPayload))
	if err != nil {
		return diag.FromErr(fmt.Errorf("failed to create request: %s", err))
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", m.(*Config).APIKey))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return diag.FromErr(fmt.Errorf("failed to create site: %s", err))
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		var apiErrResp struct {
			Message string              `json:"message"`
			Errors  map[string][]string `json:"errors"`
		}
		if err := json.NewDecoder(resp.Body).Decode(&apiErrResp); err != nil {
			return diag.FromErr(fmt.Errorf("failed to create site, status code %d, unable to parse error response: %s", resp.StatusCode, err))
		}

		errMsg := "API errors:"
		for field, messages := range apiErrResp.Errors {
			for _, msg := range messages {
				errMsg += fmt.Sprintf("\n- %s: %s", field, msg)
			}
		}
		return diag.FromErr(fmt.Errorf("API returned errors: %s", errMsg))
	}

	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return diag.FromErr(fmt.Errorf("failed to parse response: %s", err))
	}

	if id, ok := result["id"].(float64); ok {
		d.SetId(fmt.Sprintf("%d", int(id)))
	} else {
		return diag.FromErr(fmt.Errorf("API response does not contain a valid ID"))
	}

	return resourceOhDearSiteRead(ctx, d, m)
}

func resourceOhDearSiteRead(_ context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Implementa la logica di lettura della risorsa qui
	return diag.Diagnostics{}
}

func resourceOhDearSiteUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	payload := map[string]any{
		"url": d.Get("url").(string),
	}

	payload["checks"] = buildChecksPayload(d)

	if v, ok := d.GetOk("friendly_name"); ok {
		payload["friendly_name"] = v.(string)
	}
	if v, ok := d.GetOk("tags"); ok {
		payload["tags"] = v.([]interface{})
	}

	if uptimeConfig, ok := d.GetOk("uptime"); ok {
		uptimeList := uptimeConfig.([]interface{})
		if len(uptimeList) > 0 {
			uptimeMap := uptimeList[0].(map[string]interface{})
			if v, ok := uptimeMap["check_valid_status_codes"]; ok {
				payload["uptime_check_valid_status_codes"] = v.([]interface{})
			} else {
				payload["uptime_check_valid_status_codes"] = []interface{}{"2*"}
			}

			if v, ok := uptimeMap["http_client_headers"]; ok {
				headers := v.([]interface{})
				headerPayload := []map[string]string{}
				for _, header := range headers {
					headerMap := header.(map[string]interface{})
					headerPayload = append(headerPayload, map[string]string{
						"name":  headerMap["name"].(string),
						"value": headerMap["value"].(string),
					})
				}
				payload["http_client_headers"] = headerPayload
			}
			if v, ok := uptimeMap["check_location"]; ok {
				payload["uptime_check_location"] = v.(string)
			}
			if v, ok := uptimeMap["check_failed_notification_threshold"]; ok {
				payload["uptime_check_failed_notification_threshold"] = v.(int)
			}
			if v, ok := uptimeMap["check_http_verb"]; ok {
				payload["uptime_check_http_verb"] = v.(string)
			}
			if v, ok := uptimeMap["check_timeout"]; ok {
				payload["uptime_check_timeout"] = v.(int)
			}
			if v, ok := uptimeMap["check_max_redirect_count"]; ok {
				payload["uptime_check_max_redirect_count"] = v.(int)
			}
			if v, ok := uptimeMap["check_payload"]; ok {
				payload["uptime_check_payload"] = v.([]interface{})
			}
			if v, ok := uptimeMap["check_look_for_string"]; ok {
				payload["uptime_check_look_for_string"] = v.(string)
			}
			if v, ok := uptimeMap["check_absent_string"]; ok {
				payload["uptime_check_absent_string"] = v.(string)
			}
			if v, ok := uptimeMap["check_expected_response_headers"]; ok {
				headers := v.([]interface{})
				headerPayload := []map[string]string{}
				for _, header := range headers {
					headerMap := header.(map[string]interface{})
					headerPayload = append(headerPayload, map[string]string{
						"name":      headerMap["name"].(string),
						"condition": headerMap["condition"].(string),
						"value":     headerMap["value"].(string),
					})
				}
				payload["uptime_check_expected_response_headers"] = headerPayload
			}
		}
	}

	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return diag.FromErr(fmt.Errorf("failed to marshal payload: %s", err))
	}

	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/api/sites/%s", m.(*Config).APIURL, d.Id()), bytes.NewBuffer(jsonPayload))
	if err != nil {
		return diag.FromErr(fmt.Errorf("failed to create request: %s", err))
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", m.(*Config).APIKey))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return diag.FromErr(fmt.Errorf("failed to update site: %s", err))
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		var apiErrResp struct {
			Message string              `json:"message"`
			Errors  map[string][]string `json:"errors"`
		}
		if err := json.NewDecoder(resp.Body).Decode(&apiErrResp); err != nil {
			return diag.FromErr(fmt.Errorf("failed to update site, status code %d, unable to parse error response: %s", resp.StatusCode, err))
		}

		errMsg := "API errors:"
		for field, messages := range apiErrResp.Errors {
			for _, msg := range messages {
				errMsg += fmt.Sprintf("\n- %s: %s", field, msg)
			}
		}
		return diag.FromErr(fmt.Errorf("API returned errors: %s", errMsg))
	}

	return resourceOhDearSiteRead(ctx, d, m)
}

func resourceOhDearSiteDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	config := m.(*Config)

	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/api/sites/%s", config.APIURL, d.Id()), nil)
	if err != nil {
		return diag.FromErr(fmt.Errorf("failed to create request: %s", err))
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", config.APIKey))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return diag.FromErr(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusNoContent {
		var apiErrResp struct {
			Message string              `json:"message"`
			Errors  map[string][]string `json:"errors"`
		}
		if err := json.NewDecoder(resp.Body).Decode(&apiErrResp); err != nil {
			return diag.FromErr(fmt.Errorf("failed to delete site, status code %d, unable to parse error response: %s", resp.StatusCode, err))
		}

		errMsg := "API errors:"
		for field, messages := range apiErrResp.Errors {
			for _, msg := range messages {
				errMsg += fmt.Sprintf("\n- %s: %s", field, msg)
			}
		}
		return diag.FromErr(fmt.Errorf("API returned errors: %s", errMsg))
	}

	d.SetId("")
	return nil
}

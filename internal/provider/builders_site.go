package provider

import (
	"errors"
	"fmt"
	"strconv"

	"terraform-provider-ohdear/ohdear"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var (
	ErrURLRequired   = errors.New("url is required")
	ErrInvalidTeamID = errors.New("team_id is invalid")
)

func BuildSiteCreateRequest(d *schema.ResourceData) (ohdear.SiteCreateRequest, error) {
	var req ohdear.SiteCreateRequest

	url := d.Get("url").(string)
	if url == "" {
		return req, fmt.Errorf("building site create request: %w", ErrURLRequired)
	}
	req.URL = url

	teamIDStr, ok := d.GetOk("team_id")
	if !ok || teamIDStr.(string) == "" {
		return req, fmt.Errorf("building site create request: %w", ErrInvalidTeamID)
	}
	teamID, err := strconv.Atoi(teamIDStr.(string))
	if err != nil || teamID <= 0 {
		return req, fmt.Errorf("building site create request: %w", ErrInvalidTeamID)
	}
	req.TeamID = teamID

	if v, ok := d.GetOk("friendly_name"); ok {
		req.FriendlyName = v.(string)
	}
	if v, ok := d.GetOk("tags"); ok {
		req.Tags = expandStringList(v.([]interface{}))
	}

	if v, ok := d.GetOk("checks"); ok {
		req.Checks = expandChecksBlock(v.([]interface{}))
	}

	if v, ok := d.GetOk("uptime"); ok {
		if m := firstBlockMap(v.([]interface{})); m != nil {
			if vv, ok := m["check_valid_status_codes"]; ok {
				req.UptimeCheckValidStatusCodes = expandStringList(vv.([]interface{}))
			}
			if vv, ok := m["http_client_headers"]; ok {
				req.HTTPClientHeaders = expandHeaders(vv.([]interface{}))
			}
			if vv, ok := m["check_location"]; ok {
				req.UptimeCheckLocation = vv.(string)
			}
			if vv, ok := m["check_failed_notification_threshold"]; ok {
				req.UptimeCheckFailedNotificationThreshold = vv.(int)
			}
			if vv, ok := m["check_http_verb"]; ok {
				req.UptimeCheckHTTPVerb = vv.(string)
			}
			if vv, ok := m["check_timeout"]; ok {
				req.UptimeCheckTimeout = vv.(int)
			}
			if vv, ok := m["check_max_redirect_count"]; ok {
				req.UptimeCheckMaxRedirectCount = vv.(int)
			}
			if vv, ok := m["check_payload"]; ok {
				req.UptimeCheckPayload = expandPayloads(vv.([]interface{}))
			}
			// look_for / absent strings: send even if "", so use Exists if you switch to pointers later
			if vv, ok := m["check_look_for_string"]; ok {
				s := vv.(string)
				if s != "" { // omit empty to keep JSON minimal; remove this guard if API distinguishes "" vs omitted
					req.UptimeCheckLookForString = &s
				}
			}
			if vv, ok := m["check_absent_string"]; ok {
				s := vv.(string)
				if s != "" {
					req.UptimeCheckAbsentString = &s
				}
			}
			if vv, ok := m["check_expected_response_headers"]; ok {
				req.UptimeCheckExpectedResponseHeaders = expandExpectedHeaders(vv.([]interface{}))
			}
		}
	}

	// --- broken_links block -> crawler_headers ---
	if v, ok := d.GetOk("broken_links"); ok {
		if m := firstBlockMap(v.([]interface{})); m != nil {
			if vv, ok := m["crawler_headers"]; ok {
				req.CrawlerHeaders = expandHeaders(vv.([]interface{}))
			}
		}
	}

	// --- application_health block -> result_url, secret, headers ---
	if v, ok := d.GetOk("application_health"); ok {
		if m := firstBlockMap(v.([]interface{})); m != nil {
			if vv, ok := m["check_result_url"]; ok {
				req.ApplicationHealthCheckResultURL = vv.(string)
			}
			if vv, ok := m["secret"]; ok {
				if s, ok := vv.(string); ok && s != "" {
					req.ApplicationHealthCheckSecret = &s
				}
			}
			if vv, ok := m["headers"]; ok {
				req.ApplicationHealthHeaders = expandHeaders(vv.([]interface{}))
			}
		}
	}

	return req, nil
}

// BuildSiteUpdateRequest converts Terraform ResourceData -> SDK SiteUpdateRequest.
// Only sets fields that are explicitly present in the config/state, so JSON stays minimal.
func BuildSiteUpdateRequest(d *schema.ResourceData) ohdear.SiteUpdateRequest {
	var req ohdear.SiteUpdateRequest

	// ---- Top-level simple fields (pointers) ----
	if v, ok := d.GetOk("url"); ok { // allow empty string if user really wants it
		s := v.(string)
		req.URL = &s
	}
	if v, ok := d.GetOk("friendly_name"); ok {
		s := v.(string)
		req.FriendlyName = &s
	}
	if v, ok := d.GetOk("tags"); ok {
		req.Tags = expandStringList(v.([]interface{}))
	}
	// checks block -> []string
	if v, ok := d.GetOk("checks"); ok {
		req.Checks = expandChecksBlock(v.([]interface{}))
	}

	// ---- Uptime block (single item) ----
	if v, ok := d.GetOk("uptime"); ok {
		if m := firstBlockMap(v.([]interface{})); m != nil {
			if vv, ok := m["check_valid_status_codes"]; ok {
				req.UptimeCheckValidStatusCodes = expandStringList(vv.([]interface{}))
			}
			if vv, ok := m["http_client_headers"]; ok {
				req.HTTPClientHeaders = expandHeaders(vv.([]interface{}))
			}
			if vv, ok := m["check_location"]; ok {
				s := vv.(string)
				req.UptimeCheckLocation = &s
			}
			if vv, ok := m["check_failed_notification_threshold"]; ok {
				i := vv.(int)
				req.UptimeCheckFailedNotificationThreshold = &i
			}
			if vv, ok := m["check_http_verb"]; ok {
				s := vv.(string)
				req.UptimeCheckHTTPVerb = &s
			}
			if vv, ok := m["check_timeout"]; ok {
				i := vv.(int)
				req.UptimeCheckTimeout = &i
			}
			if vv, ok := m["check_max_redirect_count"]; ok {
				i := vv.(int)
				req.UptimeCheckMaxRedirectCount = &i
			}
			if vv, ok := m["check_payload"]; ok {
				req.UptimeCheckPayload = expandPayloads(vv.([]interface{}))
			}
			if vv, ok := m["check_look_for_string"]; ok {
				s := vv.(string) // send even if "", API may treat "" meaningfully
				req.UptimeCheckLookForString = &s
			}
			if vv, ok := m["check_absent_string"]; ok {
				s := vv.(string)
				req.UptimeCheckAbsentString = &s
			}
			if vv, ok := m["check_expected_response_headers"]; ok {
				req.UptimeCheckExpectedResponseHeaders = expandExpectedHeaders(vv.([]interface{}))
			}
		}
	}

	// ---- Broken links block (crawler headers only from your schema) ----
	if v, ok := d.GetOk("broken_links"); ok {
		if m := firstBlockMap(v.([]interface{})); m != nil {
			if vv, ok := m["crawler_headers"]; ok {
				req.CrawlerHeaders = expandHeaders(vv.([]interface{}))
			}
		}
	}

	// ---- Application health block ----
	if v, ok := d.GetOk("application_health"); ok {
		if m := firstBlockMap(v.([]interface{})); m != nil {
			if vv, ok := m["check_result_url"]; ok {
				s := vv.(string)
				req.ApplicationHealthCheckResultURL = &s
			}
			if vv, ok := m["headers"]; ok {
				req.ApplicationHealthHeaders = expandHeaders(vv.([]interface{}))
			}
			// Secret semantics:
			// - not present  -> unchanged (omit)
			// - present ""   -> REMOVE (nil)
			// - present "x"  -> set to "x"
			if vv, ok := m["secret"]; ok {
				s := vv.(string)
				if s == "" {
					req.ApplicationHealthCheckSecret = nil
				} else {
					req.ApplicationHealthCheckSecret = &s
				}
			}
		}
	}

	return req
}

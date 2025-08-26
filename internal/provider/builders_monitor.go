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

func BuildCreateMonitorRequest(d *schema.ResourceData) (ohdear.CreateMonitorRequest, error) {
	var req ohdear.CreateMonitorRequest

	// Required
	req.URL = d.Get("url").(string)
	if req.URL == "" {
		return req, fmt.Errorf("building site create request: %w", ErrURLRequired)
	}
	teamIDStr, _ := d.Get("team_id").(string)
	if teamIDStr == "" {
		return req, fmt.Errorf("building site create request: %w", ErrInvalidTeamID)
	}
	teamID, err := strconv.Atoi(teamIDStr)
	if err != nil || teamID <= 0 {
		return req, fmt.Errorf("building site create request: %w", ErrInvalidTeamID)
	}
	req.TeamID = teamID
	// TODO: add monitor type to schema
	req.Type = ohdear.MonitorHTTP

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
			req.UptimeCheckSettings = expandUptimeSettings(m)
		}
	}

	if v, ok := d.GetOk("broken_links"); ok {
		if m := firstBlockMap(v.([]interface{})); m != nil {
			req.CrawlerHeaders = expandHeaders(getList(m, "crawler_headers"))
		}
	}

	if v, ok := d.GetOk("application_health"); ok {
		if m := firstBlockMap(v.([]interface{})); m != nil {
			s := &ohdear.ApplicationHealthSettings{}
			if x, ok := m["check_result_url"].(string); ok && x != "" {
				s.ResultURL = &x
			}
			if x, ok := m["secret"].(string); ok && x != "" {
				s.Secret = &x
			}
			if x, ok := m["headers"]; ok {
				s.Headers = expandHeaders(x.([]interface{}))
			}
			req.ApplicationHealthCheckSettings = s
		}
	}

	return req, nil
}

func BuildUpdateMonitorRequest(d *schema.ResourceData) ohdear.UpdateMonitorRequest {
	var req ohdear.UpdateMonitorRequest

	if v, ok := d.GetOk("url"); ok {
		s := v.(string)
		req.URL = &s
	} else if d.HasChange("url") {
		s := "" // explicit clear
		req.URL = &s
	}

	if v, ok := d.GetOk("friendly_name"); ok {
		s := v.(string)
		req.FriendlyName = &s
	} else if d.HasChange("friendly_name") {
		s := ""
		req.FriendlyName = &s
	}

	if v, ok := d.GetOk("tags"); ok {
		req.Tags = expandStringList(v.([]interface{}))
	} else if d.HasChange("tags") {
		req.Tags = []string{}
	}

	if v, ok := d.GetOk("checks"); ok {
		req.Checks = expandChecksBlock(v.([]interface{}))
	} else if d.HasChange("checks") {
		req.Checks = []string{}
	}

	if v, ok := d.GetOk("uptime"); ok {
		if m := firstBlockMap(v.([]interface{})); m != nil {
			req.UptimeCheckSettings = expandUptimeSettings(m)
		}
	} else if d.HasChange("uptime") {
		req.UptimeCheckSettings = &ohdear.UptimeSettings{}
	}

	// broken_links (your schema only exposes crawler_headers here)
	if v, ok := d.GetOk("broken_links"); ok {
		if m := firstBlockMap(v.([]interface{})); m != nil {
			// headers are sent at top-level
			req.CrawlerHeaders = expandHeaders(getList(m, "crawler_headers"))
		}
	} else if d.HasChange("broken_links") {
		req.CrawlerHeaders = []ohdear.Header{}
	}

	if v, ok := d.GetOk("application_health"); ok {
		if m := firstBlockMap(v.([]interface{})); m != nil {
			s := &ohdear.ApplicationHealthSettings{}

			if x, ok := m["check_result_url"].(string); ok {
				s.ResultURL = &x
			}

			if raw, exists := m["secret"]; exists {
				if x, ok := raw.(string); ok {
					s.Secret = &x
				}
			}

			if x, ok := m["headers"]; ok {
				s.Headers = expandHeaders(x.([]interface{}))
			}

			req.ApplicationHealthCheckSettings = s
		}
	}

	return req
}

package provider

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func buildChecks(d *schema.ResourceData) map[string]interface{} {
	payload := map[string]interface{}{}
	checks := []string{}
	schema := d.Get("checks").([]interface{})[0].(map[string]interface{})
	for check, enabled := range schema {
		if enabled.(bool) {
			checks = append(checks, check)
		}
	}
	payload["checks"] = checks
	return payload
}

// helper to extract generic headers
func extractHeaders(headers []interface{}) []map[string]string {
	headerPayload := []map[string]string{}
	for _, header := range headers {
		headerMap := header.(map[string]interface{})
		headerPayload = append(headerPayload, map[string]string{
			"name":  headerMap["name"].(string),
			"value": headerMap["value"].(string),
		})
	}
	return headerPayload
}

// helper to extract expected response headers
func extractExpectedResponseHeaders(headers []interface{}) []map[string]string {
	headerPayload := []map[string]string{}
	for _, header := range headers {
		headerMap := header.(map[string]interface{})
		headerPayload = append(headerPayload, map[string]string{
			"name":      headerMap["name"].(string),
			"condition": headerMap["condition"].(string),
			"value":     headerMap["value"].(string),
		})
	}
	return headerPayload
}

// helper to get value with default
func getOrDefault(m map[string]interface{}, key string, defaultValue interface{}) interface{} {
	if v, ok := m[key]; ok {
		return v
	}
	return defaultValue
}

func buildUptime(d *schema.ResourceData) map[string]interface{} {
	payload := map[string]interface{}{}

	uptimeConfig, ok := d.GetOk("uptime")
	if !ok || len(uptimeConfig.([]interface{})) == 0 {
		return payload
	}

	uptimeMap := uptimeConfig.([]interface{})[0].(map[string]interface{})

	payload["uptime_check_valid_status_codes"] = getOrDefault(uptimeMap, "check_valid_status_codes", []interface{}{"2*"})

	if headers, ok := uptimeMap["http_client_headers"]; ok {
		payload["http_client_headers"] = extractHeaders(headers.([]interface{}))
	}

	if headers, ok := uptimeMap["check_expected_response_headers"]; ok {
		payload["uptime_check_expected_response_headers"] = extractExpectedResponseHeaders(headers.([]interface{}))
	}

	// simple list of fields to copy directly if present
	simpleFields := []string{
		"check_location",
		"check_failed_notification_threshold",
		"check_http_verb",
		"check_timeout",
		"check_max_redirect_count",
		"check_payload",
		"check_look_for_string",
		"check_absent_string",
	}

	for _, field := range simpleFields {
		if v, ok := uptimeMap[field]; ok {
			payload["uptime_"+field] = v
		}
	}

	return payload
}

func buildBrokenlink(d *schema.ResourceData) map[string]interface{} {
	payload := map[string]interface{}{}

	brokenLinksConfig, ok := d.GetOk("broken_links")
	if !ok || len(brokenLinksConfig.([]interface{})) == 0 {
		return payload
	}

	brokenLinksMap := brokenLinksConfig.([]interface{})[0].(map[string]interface{})

	if headers, ok := brokenLinksMap["crawler_headers"]; ok {
		payload["crawler_headers"] = extractHeaders(headers.([]interface{}))
	}

	return payload
}

func buildApplicationHealth(d *schema.ResourceData) map[string]interface{} {
	payload := map[string]interface{}{}

	ApplicationHealthConfig, ok := d.GetOk("application_health")
	if !ok || len(ApplicationHealthConfig.([]interface{})) == 0 {
		return payload
	}

	ApplicationHealthMap := ApplicationHealthConfig.([]interface{})[0].(map[string]interface{})

	payload["application_health_check_result_url"] = d.Get("check_result_url").(string)

	if d.Get("secret") != nil {
		payload["application_health_check_secret "] = d.Get("secret").(string)
	} else {
		payload["application_health_check_secret "] = nil
	}

	if headers, ok := ApplicationHealthMap["headers"]; ok {
		payload["application_health_headers"] = extractHeaders(headers.([]interface{}))
	}

	return payload
}

func BuildPayload(d *schema.ResourceData, event string) map[string]interface{} {
	payloadFragments := []map[string]interface{}{}

	if d.Get("url") != nil {
		payloadFragments = append(payloadFragments, map[string]interface{}{"url": d.Get("url").(string)})
	}
	if d.Get("team_id") != nil && event == "create" {
		payloadFragments = append(payloadFragments, map[string]interface{}{"team_id": d.Get("team_id").(string)})

	}
	if v, ok := d.GetOk("friendly_name"); ok {
		payloadFragments = append(payloadFragments, map[string]interface{}{"friendly_name": v.(string)})
	}
	if v, ok := d.GetOk("tags"); ok {
		payloadFragments = append(payloadFragments, map[string]interface{}{"tags": v.([]interface{})})
	}

	if d.Get("checks") != nil {
		payloadChecks := buildChecks(d)
		payloadFragments = append(payloadFragments, payloadChecks)
	}
	if d.Get("uptime") != nil {
		uptimePayload := buildUptime(d)
		payloadFragments = append(payloadFragments, uptimePayload)
	}
	if d.Get("broken_links") != nil {
		brokenliknPayload := buildBrokenlink(d)
		payloadFragments = append(payloadFragments, brokenliknPayload)
	}
	if d.Get("application_health") != nil {
		buildApplicationHealth := buildApplicationHealth(d)
		payloadFragments = append(payloadFragments, buildApplicationHealth)
	}
	finalPayload := map[string]interface{}{}
	for _, fragment := range payloadFragments {
		for k, v := range fragment {
			finalPayload[k] = v
		}
	}

	return finalPayload
}

func diagErrorf(err error, format string, a ...interface{}) diag.Diagnostics {
	return diag.Diagnostics{
		diag.Diagnostic{
			Severity: diag.Error,
			Summary:  fmt.Sprintf(format, a...),
			Detail:   err.Error(),
		},
	}
}

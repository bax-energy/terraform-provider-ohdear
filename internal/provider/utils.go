package provider

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func buildChecksPayload(d *schema.ResourceData) []string {
	checksPayload := []string{}

	if checks, ok := d.GetOk("checks"); ok {
		checksList := checks.([]interface{})
		if len(checksList) > 0 {
			checkMap := checksList[0].(map[string]interface{})
			if checkMap["uptime"].(bool) {
				checksPayload = append(checksPayload, "uptime")
			}
			if checkMap["performance"].(bool) {
				checksPayload = append(checksPayload, "performance")
			}
			if checkMap["broken_links"].(bool) {
				checksPayload = append(checksPayload, "broken_links")
			}
			if checkMap["mixed_content"].(bool) {
				checksPayload = append(checksPayload, "mixed_content")
			}
			if checkMap["lighthouse"].(bool) {
				checksPayload = append(checksPayload, "lighthouse")
			}
			if checkMap["cron"].(bool) {
				checksPayload = append(checksPayload, "cron")
			}
			if checkMap["application_health"].(bool) {
				checksPayload = append(checksPayload, "application_health")
			}
			if checkMap["sitemap"].(bool) {
				checksPayload = append(checksPayload, "sitemap")
			}
			if checkMap["dns"].(bool) {
				checksPayload = append(checksPayload, "dns")
			}
			if checkMap["domain"].(bool) {
				checksPayload = append(checksPayload, "domain")
			}
			if checkMap["certificate_health"].(bool) {
				checksPayload = append(checksPayload, "certificate_health")
			}
			if checkMap["certificate_transparency"].(bool) {
				checksPayload = append(checksPayload, "certificate_transparency")
			}
		}
	} else {
		checksPayload = []string{
			"uptime", "performance", "broken_links", "mixed_content", "lighthouse",
			"cron", "application_health", "sitemap", "dns", "domain", "certificate_health",
			"certificate_transparency",
		}
	}

	return checksPayload
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}

	return false
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

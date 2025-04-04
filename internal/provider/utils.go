package provider

import (
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

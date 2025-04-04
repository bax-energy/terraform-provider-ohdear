package provider

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func buildChecks(d *schema.ResourceData) []string {
	checks := []string{}
	schema := d.Get("checks").([]interface{})[0].(map[string]interface{})
	for check, enabled := range schema {
		if enabled.(bool) {
			checks = append(checks, check)
		}
	}

	return checks
}

func BuildPayload(d *schema.ResourceData) map[string]interface{} {
	payload := map[string]interface{}{}

	if d.Get("checks") != nil {
		checks := buildChecks(d)
		payload["checks"] = checks
	}

	return payload
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

package provider

import "terraform-provider-ohdear/ohdear"

func expandChecksBlock(items []interface{}) []string {
	if len(items) == 0 || items[0] == nil {
		return nil
	}
	m, _ := items[0].(map[string]interface{})
	if m == nil {
		return nil
	}
	out := make([]string, 0, 12)

	// Map TF booleans to API check type strings
	addIf := func(key, api string) {
		if v, ok := m[key]; ok && v.(bool) {
			out = append(out, api)
		}
	}

	addIf("uptime", "uptime")
	addIf("performance", "performance")
	addIf("broken_links", "broken_links")
	addIf("mixed_content", "mixed_content")
	addIf("lighthouse", "lighthouse")
	addIf("cron", "cron")
	addIf("application_health", "application_health")
	addIf("sitemap", "sitemap")
	addIf("dns", "dns")
	addIf("domain", "domain")
	addIf("certificate_health", "certificate_health")

	if len(out) == 0 {
		return nil
	}
	return out
}

func firstBlockMap(items []interface{}) map[string]interface{} {
	if len(items) == 0 || items[0] == nil {
		return nil
	}
	m, _ := items[0].(map[string]interface{})
	return m
}

func expandStringList(list []interface{}) []string {
	res := make([]string, 0, len(list))
	for _, v := range list {
		if s, ok := v.(string); ok && s != "" {
			res = append(res, s)
		}
	}
	if len(res) == 0 {
		return nil
	}
	return res
}

func expandHeaders(list []interface{}) []ohdear.Header {
	out := make([]ohdear.Header, 0, len(list))
	for _, raw := range list {
		if m, ok := raw.(map[string]interface{}); ok {
			h := ohdear.Header{}
			if v, ok := m["name"].(string); ok {
				h.Name = v
			}
			if v, ok := m["value"].(string); ok {
				h.Value = v
			}
			if h.Name != "" {
				out = append(out, h)
			}
		}
	}
	if len(out) == 0 {
		return nil
	}
	return out
}

func expandExpectedHeaders(list []interface{}) []ohdear.ExpectedHeader {
	out := make([]ohdear.ExpectedHeader, 0, len(list))
	for _, raw := range list {
		if m, ok := raw.(map[string]interface{}); ok {
			h := ohdear.ExpectedHeader{}
			if v, ok := m["name"].(string); ok {
				h.Name = v
			}
			if v, ok := m["condition"].(string); ok {
				h.Condition = v
			}
			if v, ok := m["value"].(string); ok {
				h.Value = v
			}
			if h.Name != "" {
				out = append(out, h)
			}
		}
	}
	if len(out) == 0 {
		return nil
	}
	return out
}

func expandPayloads(list []interface{}) []ohdear.Payload {
	out := make([]ohdear.Payload, 0, len(list))
	for _, raw := range list {
		if m, ok := raw.(map[string]interface{}); ok {
			p := ohdear.Payload{}
			if v, ok := m["name"].(string); ok {
				p.Name = v
			}
			if v, ok := m["value"].(string); ok {
				p.Value = v
			}
			if p.Name != "" {
				out = append(out, p)
			}
		}
	}
	if len(out) == 0 {
		return nil
	}
	return out
}

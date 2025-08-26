package ohdear

type CreateMonitorRequest struct {
	URL    string      `json:"url"`
	TeamID int         `json:"team_id"`
	Type   MonitorType `json:"type"` // "http" | "ping" | "tcp"

	Checks       []string `json:"checks,omitempty"`
	GroupName    string   `json:"group_name,omitempty"`
	FriendlyName string   `json:"friendly_name,omitempty"`
	Tags         []string `json:"tags,omitempty"`
	Notes        string   `json:"notes,omitempty"`
	Description  string   `json:"description,omitempty"`

	UptimeCheckSettings            *UptimeSettings            `json:"uptime_check_settings,omitempty"`
	PerformanceCheckSettings       *PerformanceSettings       `json:"performance_check_settings,omitempty"`
	CertificateHealthCheckSettings *CertificateHealthSettings `json:"certificate_health_check_settings,omitempty"`
	BrokenLinksCheckSettings       *BrokenLinksSettings       `json:"broken_links_check_settings,omitempty"`
	LighthouseCheckSettings        *LighthouseSettings        `json:"lighthouse_check_settings,omitempty"`
	ApplicationHealthCheckSettings *ApplicationHealthSettings `json:"application_health_check_settings,omitempty"`
	SitemapCheckSettings           *SitemapSettings           `json:"sitemap_check_settings,omitempty"`
	DNSCheckSettings               *DNSSettings               `json:"dns_check_settings,omitempty"`
	DomainCheckSettings            *DomainSettings            `json:"domain_check_settings,omitempty"`

	CrawlerHeaders            []Header `json:"crawler_headers,omitempty"`
	SendReportToEmails        []string `json:"send_report_to_emails,omitempty"`
	IncludeCheckTypesInReport []string `json:"include_check_types_in_report,omitempty"`
}

type UpdateMonitorRequest struct {
	URL          *string      `json:"url,omitempty"`
	Type         *MonitorType `json:"type,omitempty"`
	Checks       []string     `json:"checks,omitempty"`
	GroupName    *string      `json:"group_name,omitempty"`
	FriendlyName *string      `json:"friendly_name,omitempty"`
	Tags         []string     `json:"tags,omitempty"`
	Notes        *string      `json:"notes,omitempty"`
	Description  *string      `json:"description,omitempty"`

	UptimeCheckSettings            *UptimeSettings            `json:"uptime_check_settings,omitempty"`
	PerformanceCheckSettings       *PerformanceSettings       `json:"performance_check_settings,omitempty"`
	CertificateHealthCheckSettings *CertificateHealthSettings `json:"certificate_health_check_settings,omitempty"`
	BrokenLinksCheckSettings       *BrokenLinksSettings       `json:"broken_links_check_settings,omitempty"`
	LighthouseCheckSettings        *LighthouseSettings        `json:"lighthouse_check_settings,omitempty"`
	ApplicationHealthCheckSettings *ApplicationHealthSettings `json:"application_health_check_settings,omitempty"`
	SitemapCheckSettings           *SitemapSettings           `json:"sitemap_check_settings,omitempty"`
	DNSCheckSettings               *DNSSettings               `json:"dns_check_settings,omitempty"`
	DomainCheckSettings            *DomainSettings            `json:"domain_check_settings,omitempty"`

	CrawlerHeaders            []Header `json:"crawler_headers,omitempty"`
	SendReportToEmails        []string `json:"send_report_to_emails,omitempty"`
	IncludeCheckTypesInReport []string `json:"include_check_types_in_report,omitempty"`
}

package ohdear

type UptimeSettings struct {
	Location                 *string          `json:"uptime_check_location,omitempty"`
	ExpectedFinalRedirectURL *string          `json:"uptime_check_expected_final_redirect_url,omitempty"`
	FailedNotificationThresh *int             `json:"uptime_check_failed_notification_threshold,omitempty"`
	HTTPVerb                 *string          `json:"uptime_check_http_verb,omitempty"`
	Timeout                  *int             `json:"uptime_check_timeout,omitempty"`
	MaxRedirectCount         *int             `json:"uptime_check_max_redirect_count,omitempty"`
	Payload                  []Payload        `json:"uptime_check_payload,omitempty"`
	ValidStatusCodes         []string         `json:"uptime_check_valid_status_codes,omitempty"`
	LookForString            *string          `json:"uptime_check_look_for_string,omitempty"`
	AbsentString             *string          `json:"uptime_check_absent_string,omitempty"`
	ExpectedResponseHeaders  []ExpectedHeader `json:"uptime_check_expected_response_headers,omitempty"`
	HTTPClientHeaders        []Header         `json:"http_client_headers,omitempty"`
}

type PerformanceSettings struct {
	ThresholdInMS    *int `json:"performance_threshold_in_ms,omitempty"`
	ChangePercentage *int `json:"performance_change_percentage,omitempty"`
}

type LighthouseSettings struct {
	Continent           *string `json:"lighthouse_check_continent,omitempty"`
	CPUSlowdownModifier *int    `json:"lighthouse_cpu_slowdown_modifier,omitempty"`
}

type DNSSettings struct {
	ExtraCNAMEs        []string `json:"dns_extra_cnames,omitempty"`
	IgnoredRecordTypes []string `json:"dns_ignored_record_types,omitempty"`
	NameserversInSync  *bool    `json:"dns_check_nameservers_in_sync,omitempty"`
	MonitorMainDomain  *bool    `json:"dns_monitor_main_domain,omitempty"`
}

type DomainSettings struct {
	ExpiresSoonDays *int `json:"domain_check_expires_soon_threshold_in_days,omitempty"`
}

type SitemapSettings struct {
	Path          *string `json:"sitemap_path,omitempty"`
	Speed         *string `json:"sitemap_speed,omitempty"`
	RespectRobots *bool   `json:"respect_robots,omitempty"`
}

type CertificateHealthSettings struct {
	ExpiresSoonDays *int `json:"certificate_health_check_expires_soon_threshold_in_days,omitempty"`
}

type ApplicationHealthSettings struct {
	ResultURL *string  `json:"application_health_check_result_url,omitempty"`
	Secret    *string  `json:"application_health_check_secret,omitempty"`
	Headers   []Header `json:"application_health_headers,omitempty"`
}

type BrokenLinksSettings struct {
	IncludeExternalLinks *bool    `json:"broken_links_check_include_external_links,omitempty"`
	BrokenLinkTypes      []string `json:"broken_link_types,omitempty"`
	WhitelistedURLs      []string `json:"broken_links_whitelisted_urls,omitempty"`
}

type ReportingSettings struct {
	SendReportToEmails []string `json:"send_report_to_emails,omitempty"`
	IncludeCheckTypes  []string `json:"include_check_types_in_report,omitempty"`
}

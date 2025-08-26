package ohdear

import "encoding/json"

type MonitorType string

const (
	MonitorHTTP MonitorType = "http"
	MonitorPing MonitorType = "ping"
	MonitorTCP  MonitorType = "tcp"
)

type HeaderMatch struct {
	Name      string `json:"name,omitempty"`
	Condition string `json:"condition,omitempty"`
	Value     string `json:"value,omitempty"`
}

type ExpectedHeader struct {
	Name      string `json:"name,omitempty"`
	Condition string `json:"condition,omitempty"`
	Value     string `json:"value,omitempty"`
}

type Monitor struct {
	ID                    int         `json:"id"`
	TeamID                int         `json:"team_id"`
	Type                  MonitorType `json:"type"` // "http" | "ping" | "tcp"
	URL                   string      `json:"url"`
	UsesHTTPS             bool        `json:"uses_https"`
	SortURL               string      `json:"sort_url"`
	Label                 string      `json:"label"`
	GroupName             string      `json:"group_name"`
	Tags                  []string    `json:"tags"`
	Description           *string     `json:"description"`
	Notes                 *string     `json:"notes"`
	LatestRunDate         *Time       `json:"latest_run_date"`
	SummarizedCheckResult string      `json:"summarized_check_result"`
	BadgeID               string      `json:"badge_id"`
	MarkedForDeletionAt   *Time       `json:"marked_for_deletion_at"`

	Checks    []Check `json:"checks"`
	CreatedAt Time    `json:"created_at"`
	UpdatedAt Time    `json:"updated_at"`

	UptimeCheckSettings            json.RawMessage `json:"uptime_check_settings"`
	PerformanceCheckSettings       json.RawMessage `json:"performance_check_settings"`
	CertificateHealthCheckSettings json.RawMessage `json:"certificate_health_check_settings"`
	BrokenLinksCheckSettings       json.RawMessage `json:"broken_links_check_settings"`
	MixedContentCheckSettings      json.RawMessage `json:"mixed_content_check_settings"`
	LighthouseCheckSettings        json.RawMessage `json:"lighthouse_check_settings"`
	CronCheckSettings              json.RawMessage `json:"cron_check_settings"`
	ApplicationHealthCheckSettings json.RawMessage `json:"application_health_check_settings"`
	SitemapCheckSettings           json.RawMessage `json:"sitemap_check_settings"`
	DNSCheckSettings               json.RawMessage `json:"dns_check_settings"`
	DomainCheckSettings            json.RawMessage `json:"domain_check_settings"`

	CrawlerHeaders            []Header `json:"crawler_headers"`
	SendReportToEmails        []string `json:"send_report_to_emails"`
	IncludeCheckTypesInReport []string `json:"include_check_types_in_report"`
}

func DecodeSettings[T any](raw json.RawMessage) (T, error) {
	var t T
	if len(raw) == 0 || string(raw) == "null" || string(raw) == "[]" {
		return t, nil
	}
	return t, json.Unmarshal(raw, &t)
}

package ohdear

type Site struct {
	ID                    int      `json:"id"`
	URL                   string   `json:"url"`
	SortURL               string   `json:"sort_url"`
	Label                 string   `json:"label"`
	TeamID                int      `json:"team_id"`
	GroupName             string   `json:"group_name"`
	Tags                  []string `json:"tags"`
	BadgeID               string   `json:"badge_id"`
	Description           *string  `json:"description"`
	Notes                 *string  `json:"notes"`
	LatestRunDate         *Time    `json:"latest_run_date"`
	SummarizedCheckResult string   `json:"summarized_check_result"`
	UsesHTTPS             bool     `json:"uses_https"`
	Checks                []Check  `json:"checks"`
	CreatedAt             Time     `json:"created_at"`
	UpdatedAt             Time     `json:"updated_at"`

	SendReportToEmails        string  `json:"send_report_to_emails"`
	IncludeCheckTypesInReport *string `json:"include_check_types_in_report"`

	BrokenLinksCheckIncludeExternal bool     `json:"broken_links_check_include_external_links"`
	BrokenLinksWhitelistedURLs      []string `json:"broken_links_whitelisted_urls"`

	MarkedForDeletionAt *Time `json:"marked_for_deletion_at"`

	UptimeCheckPayload       []Payload `json:"uptime_check_payload"`
	HTTPClientHeaders        []Header  `json:"http_client_headers"`
	CrawlerHeaders           []Header  `json:"crawler_headers"`
	ApplicationHealthHeaders []Header  `json:"application_health_headers"`

	SitemapPath string `json:"sitemap_path"`
}

type Header struct {
	Name  string `json:"name,omitempty"`
	Value string `json:"value,omitempty"`
}

type Payload struct {
	Name  string `json:"name,omitempty"`
	Value string `json:"value,omitempty"`
}

type Check struct {
	ID               int         `json:"id"`
	Type             string      `json:"type"`
	Label            string      `json:"label"`
	Enabled          bool        `json:"enabled"`
	LatestRunEndedAt *Time       `json:"latest_run_ended_at"`
	LatestRunResult  string      `json:"latest_run_result"`
	Summary          string      `json:"summary"`
	Settings         interface{} `json:"settings"`
	ActiveSnooze     interface{} `json:"active_snooze"`
}

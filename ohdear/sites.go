package ohdear

import (
	"context"
	"fmt"
)

type ExpectedHeader struct {
	Name      string `json:"name,omitempty"`
	Condition string `json:"condition,omitempty"`
	Value     string `json:"value,omitempty"`
}

type SiteCreateRequest struct {
	// Required
	URL    string `json:"url"`
	TeamID int    `json:"team_id"`

	// Optional
	Checks                                 []string         `json:"checks,omitempty"`
	GroupName                              string           `json:"group_name,omitempty"`
	FriendlyName                           string           `json:"friendly_name,omitempty"`
	Tags                                   []string         `json:"tags,omitempty"`
	Notes                                  string           `json:"notes,omitempty"`
	Description                            string           `json:"description,omitempty"`
	UptimeCheckLocation                    string           `json:"uptime_check_location,omitempty"`
	UptimeCheckExpectedFinalRedirectURL    string           `json:"uptime_check_expected_final_redirect_url,omitempty"`
	UptimeCheckFailedNotificationThreshold int              `json:"uptime_check_failed_notification_threshold,omitempty"`
	UptimeCheckHTTPVerb                    string           `json:"uptime_check_http_verb,omitempty"`
	UptimeCheckTimeout                     int              `json:"uptime_check_timeout,omitempty"`
	UptimeCheckMaxRedirectCount            int              `json:"uptime_check_max_redirect_count,omitempty"`
	UptimeCheckPayload                     []Payload        `json:"uptime_check_payload,omitempty"`
	UptimeCheckValidStatusCodes            []string         `json:"uptime_check_valid_status_codes,omitempty"`
	UptimeCheckLookForString               *string          `json:"uptime_check_look_for_string,omitempty"`
	UptimeCheckAbsentString                *string          `json:"uptime_check_absent_string,omitempty"`
	UptimeCheckExpectedResponseHeaders     []ExpectedHeader `json:"uptime_check_expected_response_headers,omitempty"`
	HTTPClientHeaders                      []Header         `json:"http_client_headers,omitempty"`
	PerformanceThresholdInMS               int              `json:"performance_threshold_in_ms,omitempty"`
	PerformanceChangePercentage            int              `json:"performance_change_percentage,omitempty"`
	CrawlerHeaders                         []Header         `json:"crawler_headers,omitempty"`
	BrokenLinksCheckIncludeExternalLinks   *bool            `json:"broken_links_check_include_external_links,omitempty"`
	BrokenLinkTypes                        []string         `json:"broken_link_types,omitempty"`
	BrokenLinksWhitelistedURLs             []string         `json:"broken_links_whitelisted_urls,omitempty"`
	RespectRobots                          *bool            `json:"respect_robots,omitempty"`
	SitemapPath                            string           `json:"sitemap_path,omitempty"`
	SitemapSpeed                           string           `json:"sitemap_speed,omitempty"`
	ApplicationHealthCheckResultURL        string           `json:"application_health_check_result_url,omitempty"`
	ApplicationHealthCheckSecret           *string          `json:"application_health_check_secret,omitempty"`
	ApplicationHealthHeaders               []Header         `json:"application_health_headers,omitempty"`
	CertificateHealthCheckExpiresSoonDays  int              `json:"certificate_health_check_expires_soon_threshold_in_days,omitempty"`
	DNSCheckNameserversInSync              *bool            `json:"dns_check_nameservers_in_sync,omitempty"`
	DNSMonitorMainDomain                   *bool            `json:"dns_monitor_main_domain,omitempty"`
	DNSExtraCNAMEs                         []string         `json:"dns_extra_cnames,omitempty"`
	DNSIgnoredRecordTypes                  []string         `json:"dns_ignored_record_types,omitempty"`
	DomainCheckExpiresSoonDays             int              `json:"domain_check_expires_soon_threshold_in_days,omitempty"`
	LighthouseCheckContinent               string           `json:"lighthouse_check_continent,omitempty"`
	LighthouseCPUSlowdownModifier          int              `json:"lighthouse_cpu_slowdown_modifier,omitempty"`
	SendReportToEmails                     []string         `json:"send_report_to_emails,omitempty"`
	IncludeCheckTypesInReport              []string         `json:"include_check_types_in_report,omitempty"`
}

type SiteUpdateRequest struct {
	URL                                    *string          `json:"url,omitempty"`
	Checks                                 []string         `json:"checks,omitempty"`
	GroupName                              *string          `json:"group_name,omitempty"`
	FriendlyName                           *string          `json:"friendly_name,omitempty"`
	Tags                                   []string         `json:"tags,omitempty"`
	Notes                                  *string          `json:"notes,omitempty"`
	Description                            *string          `json:"description,omitempty"`
	UptimeCheckLocation                    *string          `json:"uptime_check_location,omitempty"`
	UptimeCheckExpectedFinalRedirectURL    *string          `json:"uptime_check_expected_final_redirect_url,omitempty"`
	UptimeCheckFailedNotificationThreshold *int             `json:"uptime_check_failed_notification_threshold,omitempty"`
	UptimeCheckHTTPVerb                    *string          `json:"uptime_check_http_verb,omitempty"`
	UptimeCheckTimeout                     *int             `json:"uptime_check_timeout,omitempty"`
	UptimeCheckMaxRedirectCount            *int             `json:"uptime_check_max_redirect_count,omitempty"`
	UptimeCheckPayload                     []Payload        `json:"uptime_check_payload,omitempty"`
	UptimeCheckValidStatusCodes            []string         `json:"uptime_check_valid_status_codes,omitempty"`
	UptimeCheckLookForString               *string          `json:"uptime_check_look_for_string,omitempty"`
	UptimeCheckAbsentString                *string          `json:"uptime_check_absent_string,omitempty"`
	UptimeCheckExpectedResponseHeaders     []ExpectedHeader `json:"uptime_check_expected_response_headers,omitempty"`
	HTTPClientHeaders                      []Header         `json:"http_client_headers,omitempty"`
	PerformanceThresholdInMS               *int             `json:"performance_threshold_in_ms,omitempty"`
	PerformanceChangePercentage            *int             `json:"performance_change_percentage,omitempty"`
	CrawlerHeaders                         []Header         `json:"crawler_headers,omitempty"`
	BrokenLinksCheckIncludeExternalLinks   *bool            `json:"broken_links_check_include_external_links,omitempty"`
	BrokenLinkTypes                        []string         `json:"broken_link_types,omitempty"`
	BrokenLinksWhitelistedURLs             []string         `json:"broken_links_whitelisted_urls,omitempty"`
	RespectRobots                          *bool            `json:"respect_robots,omitempty"`
	SitemapPath                            *string          `json:"sitemap_path,omitempty"`
	SitemapSpeed                           *string          `json:"sitemap_speed,omitempty"`
	ApplicationHealthCheckResultURL        *string          `json:"application_health_check_result_url,omitempty"`
	ApplicationHealthCheckSecret           *string          `json:"application_health_check_secret,omitempty"`
	ApplicationHealthHeaders               []Header         `json:"application_health_headers,omitempty"`
	CertificateHealthCheckExpiresSoonDays  *int             `json:"certificate_health_check_expires_soon_threshold_in_days,omitempty"`
	DNSCheckNameserversInSync              *bool            `json:"dns_check_nameservers_in_sync,omitempty"`
	DNSMonitorMainDomain                   *bool            `json:"dns_monitor_main_domain,omitempty"`
	DNSExtraCNAMEs                         []string         `json:"dns_extra_cnames,omitempty"`
	DNSIgnoredRecordTypes                  []string         `json:"dns_ignored_record_types,omitempty"`
	DomainCheckExpiresSoonDays             *int             `json:"domain_check_expires_soon_threshold_in_days,omitempty"`
	LighthouseCheckContinent               *string          `json:"lighthouse_check_continent,omitempty"`
	LighthouseCPUSlowdownModifier          *int             `json:"lighthouse_cpu_slowdown_modifier,omitempty"`
	SendReportToEmails                     []string         `json:"send_report_to_emails,omitempty"`
	IncludeCheckTypesInReport              []string         `json:"include_check_types_in_report,omitempty"`
}

type SiteService interface {
	Get(ctx context.Context, id int) (*Site, error)
	Create(ctx context.Context, req SiteCreateRequest) (*Site, error)
	Update(ctx context.Context, siteID int, req SiteUpdateRequest) (*Site, error)
	Delete(ctx context.Context, siteID int) error
}

type siteService struct {
	client *Client
}

func (s *siteService) Get(ctx context.Context, id int) (*Site, error) {
	var Site Site
	err := s.client.do(ctx, "GET", fmt.Sprintf("/sites/%d", id), nil, &Site)
	return &Site, err
}

func (s *siteService) Create(ctx context.Context, req SiteCreateRequest) (*Site, error) {
	var Site Site
	err := s.client.do(ctx, "POST", "/sites", req, &Site)
	return &Site, err
}

func (s *siteService) Update(ctx context.Context, siteID int, req SiteUpdateRequest) (*Site, error) {
	var updatedSite Site
	endpoint := fmt.Sprintf("/sites/%d", siteID)

	if err := s.client.do(ctx, "PUT", endpoint, req, &updatedSite); err != nil {
		return nil, err
	}

	return &updatedSite, nil
}

func (s *siteService) Delete(ctx context.Context, siteID int) error {
	endpoint := fmt.Sprintf("/sites/%d", siteID)

	if err := s.client.do(ctx, "DELETE", endpoint, nil, nil); err != nil {
		return err
	}

	return nil
}

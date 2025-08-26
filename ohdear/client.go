package ohdear

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"
	"time"
)

type Client struct {
	httpClient *http.Client
	baseURL    string
	apiKey     string

	Sites SiteService
}

func NewClient(apiKey string, opts ...func(*Client)) *Client {
	c := &Client{
		httpClient: &http.Client{Timeout: 10 * time.Second},
		baseURL:    "https://ohdear.app/api",
		apiKey:     apiKey,
	}

	for _, opt := range opts {
		opt(c)
	}

	c.Sites = &siteService{client: c}
	return c
}

// UserAgentTransport is a custom RoundTripper that adds a User-Agent header
type userAgentTransport struct {
	Transport http.RoundTripper
	UserAgent string
}

func (uat *userAgentTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	req.Header.Set("User-Agent", uat.UserAgent)
	return uat.Transport.RoundTrip(req)
}

// WithUserAgent sets a custom User-Agent header
func WithUserAgent(ua string) func(*Client) {
	return func(c *Client) {
		c.httpClient.Transport = &userAgentTransport{
			Transport: http.DefaultTransport,
			UserAgent: ua,
		}
	}
}

// WithBaseURL overrides the default API base URL
func WithBaseURL(url string) func(*Client) {
	return func(c *Client) {
		c.baseURL = url
	}
}

func (c *Client) do(ctx context.Context, method, path string, body interface{}, out interface{}) error {
	var reqBody strings.Reader
	if body != nil {
		b, err := json.Marshal(body)
		if err != nil {
			return err
		}
		reqBody = *strings.NewReader(string(b))
	}

	req, err := http.NewRequestWithContext(ctx, method, c.baseURL+path, &reqBody)
	if err != nil {
		return err
	}
	req.Header.Set("Authorization", "Bearer "+c.apiKey)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 300 {
		return parseAPIError(resp)
	}

	if out != nil {
		return json.NewDecoder(resp.Body).Decode(out)
	}
	return nil
}

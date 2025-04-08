package provider

import "fmt"

type StatusPage struct {
	ID     int
	TeamID string
	Title  string
	Sites  []struct {
		ID        int
		Clickable bool
	}
}

func (c *Client) GetStatusPage(id int) (*StatusPage, error) {
	resp, err := c.R().
		SetResult(&StatusPage{}).
		Get(fmt.Sprintf("/api/status-pages/%d", id))
	if err != nil {
		return nil, err
	}

	return resp.Result().(*StatusPage), nil
}

func (c *Client) AddStatusPage(payload interface{}) (*StatusPage, error) {
	resp, err := c.R().
		SetBody(payload).
		SetResult(&StatusPage{}).
		Post("/api/status-pages")
	if err != nil {
		return nil, err
	}

	return resp.Result().(*StatusPage), nil
}

func (c *Client) AddSiteStatusPage(id string, payload interface{}) error {
	_, err := c.R().
		SetBody(payload).
		Post(fmt.Sprintf("/api/status-pages/%s/sites", id))
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) RemoveSiteStatusPage(id string, siteid int) error {
	_, err := c.R().
		Delete(fmt.Sprintf("/api/status-pages/%s/sites/%d", id, siteid))
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) RemoveStatusPage(id int) error {
	_, err := c.R().Delete(fmt.Sprintf("/api/status-pages/%d", id))
	return err
}

package provider

import "fmt"

type Site struct {
	ID     int
	URL    string
	TeamID int `json:"team_id"`
}

func (c *Client) GetSite(id int) (*Site, error) {
	resp, err := c.R().
		SetResult(&Site{}).
		Get(fmt.Sprintf("/api/sites/%d", id))
	if err != nil {
		return nil, err
	}

	return resp.Result().(*Site), nil
}

func (c *Client) AddSite(payload interface{}) (*Site, error) {
	fmt.Printf("Debug: Adding site : %+v\n", payload)
	resp, err := c.R().
		SetBody(payload).
		SetResult(&Site{}).
		Post("/api/sites")
	if err != nil {
		return nil, err
	}

	return resp.Result().(*Site), nil
}

func (c *Client) UpdateSite(id string, payload interface{}) (*Site, error) {
	resp, err := c.R().
		SetBody(payload).
		SetResult(&Site{}).
		Put(fmt.Sprintf("/api/sites/%s", id))
	if err != nil {
		return nil, err
	}

	return resp.Result().(*Site), nil
}

func (c *Client) RemoveSite(id int) error {
	_, err := c.R().Delete(fmt.Sprintf("/api/sites/%d", id))
	return err
}

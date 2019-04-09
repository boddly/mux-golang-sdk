package assets

import (
	"encoding/json"
	"errors"
	"net/http"
)

// ListData is the data object returned from the List request
type ListData struct {
	Data []Asset `json:"data"`
}

// List assets
func (c *Client) List(limit, page int32) ([]Asset, error) {
	req, err := http.NewRequest(http.MethodGet, AssetURL, nil)
	req.SetBasicAuth(c.AccessToken, c.SecretKey)

	resp, err := c.HTTP.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New(resp.Status)
	}

	r := &ListData{}
	return r.Data, json.NewDecoder(resp.Body).Decode(r)
}

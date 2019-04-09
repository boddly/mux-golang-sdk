package assets

import (
	"encoding/json"
	"net/http"
)

// Get an asset using the ID of the asset to be retrieved
func (c *Client) Get(id string) (*Asset, error) {
	req, err := http.NewRequest(http.MethodGet, AssetURL+id, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.SetBasicAuth(c.AccessToken, c.SecretKey)

	resp, err := c.HTTP.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	asset := &Asset{}
	return asset, json.NewDecoder(resp.Body).Decode(asset)
}

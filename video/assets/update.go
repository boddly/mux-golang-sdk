package assets

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

// UpdateResponse is the struct that is returned from update requests
type UpdateResponse struct {
	Data Asset `json:"data"`
}

// UpdateMP4 allows you add or remove mp4 support for an asset after it's created.
// Currently there are two values supported in this request: standard and none
func (c *Client) UpdateMP4(id, value string) (*Asset, error) {

	b := []byte(fmt.Sprintf(`{"mp4_support":"%s"}`, value))
	req, err := http.NewRequest(http.MethodPut, AssetURL+"/"+id+"/mp4-support", bytes.NewBuffer(b))
	req.SetBasicAuth(c.AccessToken, c.SecretKey)

	resp, err := c.HTTP.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New(resp.Status)
	}

	a := &UpdateResponse{}
	return &a.Data, json.NewDecoder(resp.Body).Decode(a)

}

// UpdateMaster allows you add temporary access to the master (highest-quality) version of the asset in MP4 format.
// A URL will be created that can be used to download the master version for 24 hours.
// After 24 hours Master Access will revert to "none".
func (c *Client) UpdateMaster(id, value string) (*Asset, error) {
	b := []byte(fmt.Sprintf(`{ "master_access": "%s" }`, value))
	req, err := http.NewRequest(http.MethodPut, AssetURL+"/"+id+"/master-access", bytes.NewBuffer(b))
	req.SetBasicAuth(c.AccessToken, c.SecretKey)

	resp, err := c.HTTP.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New(resp.Status)
	}

	a := &UpdateResponse{}
	return &a.Data, json.NewDecoder(resp.Body).Decode(a)
}

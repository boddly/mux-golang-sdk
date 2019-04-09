package assets

import (
	"encoding/json"
	"errors"
	"net/http"
)

// InputInfoResponse is the response struct from the input info request
type InputInfoResponse struct {
	Data []*InputInfoData `json:"data"`
}

// InputInfoData is an array of input objects, each of which contains details of an input (a file or existing Mux asset) that was used to create this asset.
type InputInfoData struct {
	Settings Input        `json:"settings"`
	File     InputDetails `json:"file"`
}

// InputInfo Returns a list of the input objects that were used to create the asset along with any settings that were applied to each input.
func (c *Client) InputInfo(id string) ([]*InputInfoData, error) {
	req, err := http.NewRequest(http.MethodGet, AssetURL+"/"+id+"/input-info", nil)
	if err != nil {
		return nil, err
	}
	req.SetBasicAuth(c.AccessToken, c.SecretKey)

	resp, err := c.HTTP.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New(resp.Status)
	}

	r := &InputInfoResponse{}
	return r.Data, json.NewDecoder(resp.Body).Decode(r)
}

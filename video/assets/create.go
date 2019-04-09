package assets

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
)

// Input is an array of objects that each describe an input file to be used to create the asset.
// As a shortcut, input can also be a string URL for a file when only one input file is used. See input[].url for requirements.
type Input struct {
	URL            string         `json:"url"`
	OverlaySetting OverlaySetting `json:"overlay_settings,omitempty"`
}

// OverlaySetting is an object that describes how the image file referenced in url should be placed over the video (i.e. watermarking).
// Currently overlays can only be images, and a video file must exist before the overlay image in the array of inputs.
type OverlaySetting struct {
	VerticalAlign    string `json:"vertical_align,omitempty"`
	VerticalMargin   string `json:"vertical_margin,omitempty"`
	HorizontalAlign  string `json:"horizontal_align,omitempty"`
	HorizontalMargin string `json:"horizontal_margin,omitempty"`
	Width            string `json:"width,omitempty"`
	Height           string `json:"height,omitempty"`
	Opacity          string `json:"opacity,omitempty"`
}

// CreateOptions is the request object for the asset.Create request
type CreateOptions struct {
	Input          []Input  `json:"input"`
	PlaybackPolicy []string `json:"playback_policy,omitempty"`
	Passthrough    string   `json:"passthrough,omitempty"`
	MP4Support     string   `json:"mp4_support,omitempty"`
}

// CreateResponse is the response from the create request
type CreateResponse struct {
	Data CreateData `json:"data"`
}

// CreateData is the Data object returned with the creation response
type CreateData struct {
	Status       string `json:"status"`
	MP4Support   string `json:"mp4_support"`
	MasterAccess string `json:"master_access"`
	ID           string `json:"id"`
	CreatedAt    string `json:"created_at"`
}

// Create an asset
func (c *Client) Create(opts CreateOptions) (*CreateResponse, error) {
	j, err := json.Marshal(opts)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest(http.MethodPost, AssetURL, bytes.NewBuffer(j))
	req.Header.Set("Content-Type", "application/json")
	req.SetBasicAuth(c.AccessToken, c.SecretKey)

	resp, err := c.HTTP.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		return nil, errors.New(resp.Status)
	}

	r := &CreateResponse{}
	return r, json.NewDecoder(resp.Body).Decode(r)

}

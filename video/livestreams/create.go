package livestreams

import "github.com/boddly/mux-golang-sdk/video/assets"

// CreateRequest is the object passed into the create request
type CreateRequest struct {
	ReconnectWindow  float64              `json:"reconnect_window"`
	PlaybackPolicy   []string             `json:"playback_policy"`
	NewAssetSettings assets.CreateOptions `json:"new_asset_settings"`
}

// Create a live stream
func (ls *LiveStream) Create() (*LiveStream, error) {
	return nil, nil
}

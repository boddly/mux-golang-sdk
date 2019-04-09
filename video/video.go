package video

import "github.com/boddly/mux-golang-sdk/video/assets"

// Video provides access to the mux video API
type Video struct {
	Assets assets.Client
}

// New is a video Constructor
func New(accessToken, secret string) *Video {
	return &Video{
		Assets: assets.Client{AccessToken: accessToken, SecretKey: secret},
	}
}

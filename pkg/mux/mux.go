package mux

import (
	"github.com/boddly/mux-golang-sdk/pkg/data"
	"github.com/boddly/mux-golang-sdk/pkg/video"
)

// Mux provides access to the Mux Video and Mux Data API
type Mux struct {
	Data  *data.Data
	Video video.Video
}

// New is a mux Constructor
func New(accessToken, secret string) *Mux {
	return &Mux{
		Data: data.New(accessToken, secret),
	}
}

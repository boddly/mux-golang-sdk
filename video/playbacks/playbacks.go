package playbacks

// IDs are used with stream.mux.com to create the source URL for a video player.
type IDs struct {
	ID     string `json:"id"`
	Policy string `json:"policy"`
}

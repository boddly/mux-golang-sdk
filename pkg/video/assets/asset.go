package assets

import "net/http"

const (
	// AssetURL is the base URL for asset requests
	AssetURL = "https://api.mux.com/video/v1/assets"
)

// Client handles all asset requests to the mux API
type Client struct {
	HTTP        http.Client
	AccessToken string
	SecretKey   string
}

// Asset refers to a piece of media content that is stored or is being live streamed through the Mux system.
// An asset always has a duration and one or more tracks (audio, video, and text data).
type Asset struct {
	ID                  string           `json:"id"`
	CreatedAt           string           `json:"created_at"`
	Status              string           `json:"status"`
	Duration            float64          `json:"duration"`
	MaxStoredResolution string           `json:"max_stored_resolution"`
	MaxStoredFrameRate  float64          `json:"max_stored_frame_rate"`
	AspectRatio         string           `json:"aspect_ratio"`
	PerTitleEncode      bool             `json:"per_title_encode"`
	Playbacks           []Playback       `json:"playback_ids"`
	Tracks              []Track          `json:"tracks"`
	MP4Support          string           `json:"mp4_support"`
	StaticRenditions    StaticRenditions `json:"static_renditions"`
	MasterAccess        string           `json:"master_access"`
	Master              Master           `json:"master"`
	Passthrough         string           `json:"passthrough"`
}

// Playback IDs are used with stream.mux.com to create the source URL for a video player.
type Playback struct {
	ID     string `json:"id"`
	Policy string `json:"policy"`
}

// Track is the individual media tracks that make up an asset.
type Track struct {
	ID               string  `json:"id"`
	Type             string  `json:"type"`
	Duration         float64 `json:"duration"`
	Language         string  `json:"language"`
	MaxWidth         int     `json:"max_width"`
	MaxHeight        int     `json:"max_height"`
	MaxFrameRate     float64 `json:"max_frame_rate"`
	MaxChannels      int     `json:"max_channels"`
	MaxChannelLayout string  `json:"max_channel_layout"`
	TextTrackType    string  `json:"text_track_type"`
}

// StaticRenditions are an object containing the current status of any static renditions such as mp4s.
// The object does not exist if no static renditions have been requested.
type StaticRenditions struct {
	Status string `json:"status"`
	Files  []File `json:"files"`
}

// File objects
type File struct {
	Name     string `json:"name"`
	Ext      string `json:"ext"`
	Height   int    `json:"height"`
	Width    int    `json:"width"`
	Bitrate  int    `json:"bitrate"`
	Filesize int    `json:"filesize"`
}

// Master is an object containing the current status of Master Access and the link to the Master MP4 file when ready.
// This object does not exist if master_acess is set to none and when the temporary URL expires.
type Master struct {
	Status string `json:"status"`
	URL    string `json:"url"`
}

// InputTracks are the tracks that were found in the input source.
type InputTracks struct {
	Type       string  `json:"type"` // Possible values are video, audio, and text.
	Duration   float64 `json:"duration"`
	Width      int     `json:"width"`
	Height     int     `json:"height"`
	FrameRate  float64 `json:"frame_rate"`
	Encoding   string  `json:"encoding"`
	SampleRate int     `json:"sample_rate"`
	SampleSize int     `json:"sample_size"`
}

// InputDetails is the inspected details of the input file. In the case of a Mux asset input this will be the available details for the input asset and its tracks.
type InputDetails struct {
	ContainerFormat string        `json:"container_format"`
	Tracks          []InputTracks `json:"tracks"`
}

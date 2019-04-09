package livestreams

import (
	"github.com/boddly/mux-golang-sdk/video/assets"
	"github.com/boddly/mux-golang-sdk/video/playbacks"
)

//LiveStream represents a unique live stream of video being pushed to Mux.
// It includes configuration details (a Stream Key) for live broadcasting software/hardware and a Playback ID for playing the stream anywhere.
// Currently, RTMP is the only supported method of ingesting video.
// Use rtmp://live.mux.com/app/${STREAM_KEY} with the Live Stream's Stream Key for getting the video into Mux, and use https://stream.mux.com with the Live Stream's Playback ID for playback.
// A Live Stream can be used once for a specific event, or re-used forever.
// If you're building an application like Facebook Live or Twitch, you could create one Live Stream per user.
// This allows them the configure their broadcasting software once, and the Live Stream Playback ID will always show their latest stream.
// Each RTMP session creates a new video asset automatically.
// You can set up a webhook to be notified each time a broadcast (or Live Stream RTMP session) begins or ends with the video.live_stream.active and video.live_stream.idle events respectively.
// Assets that are created from a Live Stream have the same behavior as other Assets you create.
type LiveStream struct {
	ID               string               `json:"id"`
	CreatedAt        int                  `json:"created_at"`
	Status           string               `json:"status"`
	StreamKey        string               `json:"stream_key"`
	ActiveAssetID    string               `json:"active_asset_id"`
	RecentAssetIDs   string               `json:"recent_asset_ids"`
	NewAssetSettings assets.CreateOptions `json:"new_asset_settings"`
	PlaybackIDs      []playbacks.IDs      `json:"playback_ids"`
}

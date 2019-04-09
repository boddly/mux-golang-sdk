package views

import "time"

// VideoViews provides access to the Mux Data Video Views API
type VideoViews struct{}

// ListData is the data in the response from the List method
type ListData struct {
	ViewerOSFamily        string      `json:"viewer_os_family"`
	ViewerApplicationName string      `json:"viewer_application_name"`
	ViewStart             time.Time   `json:"view_start"`
	ViewEnd               time.Time   `json:"view_end"`
	VideoTitle            string      `json:"video_title"`
	TotalRowCount         int         `json:"total_row_count"`
	PlayerErrorMessage    string      `json:"player_error_message"`
	PlayerErrorCode       interface{} `json:"player_error_code"`
	ID                    string      `json:"id"`
	ErrorTypeID           int         `json:"error_type_id"`
	CountryCode           string      `json:"country_code"`
}

// VideoData represents the set of data returned about a video
type VideoData struct {
	ViewSeekDuration               float64       `json:"view_seek_duration"`
	VideoLanguage                  string        `json:"video_language"`
	ViewerOSfamily                 string        `json:"viewer_os_family"`
	ViewerUserAgent                string        `json:"viewer_user_agent"`
	PlayerErrorCode                int           `json:"player_error_code"`
	PlayerHeight                   int           `json:"player_height"`
	PlayerPoster                   int           `json:"player_poster"`
	TimeToFirstFrame               float64       `json:"time_to_first_frame"`
	PlayerSourceType               string        `json:"player_source_type"`
	VideoStreamType                string        `json:"video_stream_type"`
	CDN                            string        `json:"cdn"`
	BufferingRate                  int           `json:"buffering_rate"`
	Metro                          string        `json:"metro"`
	PlayerInstanceID               string        `json:"player_instance_id"`
	ViewerApplicationVersion       string        `json:"viewer_application_version"`
	UpdatedAt                      time.Time     `json:"updated_at"`
	PlayerSourceHostName           string        `json:"player_source_host_name"`
	PrerollPlayed                  bool          `json:"preroll_played"`
	PlayerSourceDuration           string        `json:"player_source_duration"`
	ViewTotalDownscaling           int           `json:"view_total_downscaling"`
	PlayerSourceStreamType         string        `json:"player_source_stream_type"`
	ContinentCode                  string        `json:"continent_code"`
	ViewErrorID                    string        `json:"view_error_id"`
	PageLoadTime                   float64       `json:"page_load_time"`
	VideoVariantID                 string        `json:"video_variant_id"`
	ViewID                         string        `json:"view_id"`
	WatchTime                      float64       `json:"watch_time"`
	ISP                            string        `json:"isp"`
	InsertedAt                     time.Time     `json:"inserted_at"`
	ViewerApplicationName          string        `json:"viewer_application_name"`
	ExitBeforeVideoStart           bool          `json:"exit_before_video_start"`
	ViewSeekCount                  int           `json:"view_seek_count"`
	PlayerSoftware                 string        `json:"player_software"`
	PlayerPreload                  bool          `json:"player_preload"`
	ASN                            string        `json:"asn"`
	ViewEnd                        time.Time     `json:"view_end"`
	Events                         []interface{} `json:"events"`
	VideoStartupPrerollRequestTime interface{}   `json:"video_startup_preroll_request_time"`
	ViewTotalContentPlaybackTime   interface{}   `json:"view_total_content_playback_time"`
	BufferingDuration              interface{}   `json:"buffering_duration"`
	RequestsForFirstPreroll        interface{}   `json:"requests_for_first_preroll"`
	PlayerWidth                    interface{}   `json:"player_width"`
	PageURL                        string        `json:"page_url"`
	PrerollAdTagHostname           interface{}   `json:"preroll_ad_tag_hostname"`
	SessionID                      string        `json:"session_id"`
	ViewerOsArchitecture           interface{}   `json:"viewer_os_architecture"`
	PageType                       interface{}   `json:"page_type"`
	ViewMaxDownscalePercentage     interface{}   `json:"view_max_downscale_percentage"`
	PrerollAdAssetHostname         interface{}   `json:"preroll_ad_asset_hostname"`
	VideoVariantName               interface{}   `json:"video_variant_name"`
	ExperimentName                 interface{}   `json:"experiment_name"`
	Watched                        bool          `json:"watched"`
	MuxAPIVersion                  interface{}   `json:"mux_api_version"`
	PlayerLoadTime                 interface{}   `json:"player_load_time"`
	PrerollRequested               bool          `json:"preroll_requested"`
	Region                         string        `json:"region"`
	PlayerErrorMessage             bool          `json:"player_error_message"`
	CountryCode                    interface{}   `json:"country_code"`
	PlayerSourceDomain             interface{}   `json:"player_source_domain"`
	Longitude                      float64       `json:"longitude"`
	PlayerName                     string        `json:"player_name"`
	VideoProducer                  interface{}   `json:"video_producer"`
	VideoSeries                    interface{}   `json:"video_series"`
	CountryName                    string        `json:"country_name"`
	RebufferPercentage             float64       `json:"rebuffer_percentage"`
	UsedFullscreen                 bool          `json:"used_fullscreen"`
	VideoEncodingVariant           interface{}   `json:"video_encoding_variant"`
	PlayerLanguage                 string        `json:"player_language"`
	ViewerDeviceManufacturer       string        `json:"viewer_device_manufacturer"`
	ViewStart                      time.Time     `json:"view_start"`
	Latitude                       float64       `json:"latitude"`
	ViewerDeviceCategory           interface{}   `json:"viewer_device_category"`
	VideoID                        string        `json:"video_id"`
	PlayerSourceWidth              interface{}   `json:"player_source_width"`
	PropertyID                     int           `json:"property_id"`
	City                           string        `json:"city"`
	PlayerSoftwareVersion          interface{}   `json:"player_software_version"`
	PlayerAutoplay                 bool          `json:"player_autoplay"`
	VideoDuration                  float64       `json:"video_duration"`
	BufferingCount                 interface{}   `json:"buffering_count"`
	ViewMaxUpscalePercentage       interface{}   `json:"view_max_upscale_percentage"`
	PlatformSummary                interface{}   `json:"platform_summary"`
	ViewerApplicationEngine        interface{}   `json:"viewer_application_engine"`
	MuxEmbedVersion                interface{}   `json:"mux_embed_version"`
	ErrorTypeID                    interface{}   `json:"error_type_id"`
	ID                             interface{}   `json:"id"`
	PlayerVersion                  interface{}   `json:"player_version"`
	PlayerStartupTime              interface{}   `json:"player_startup_time"`
	PlayerSourceURL                interface{}   `json:"player_source_url"`
	ViewTotalUpscaling             interface{}   `json:"view_total_upscaling"`
	ShortTime                      interface{}   `json:"short_time"`
	ViewerUserID                   interface{}   `json:"viewer_user_id"`
	PlayerMuxPluginVersion         interface{}   `json:"player_mux_plugin_version"`
	VideoTitle                     interface{}   `json:"video_title"`
	MuxViewerID                    interface{}   `json:"mux_viewer_id"`
	ViewerDeviceName               interface{}   `json:"viewer_device_name"`
	ViewerOsVersion                interface{}   `json:"viewer_os_version"`
	VideoContentType               interface{}   `json:"video_content_type"`
	PlayerViewCount                interface{}   `json:"player_view_count"`
	PlayerSourceHeight             interface{}   `json:"player_source_height"`
	VideoStartupPrerollLoadTime    interface{}   `json:"video_startup_preroll_load_time"`
	PlatformDescription            interface{}   `json:"platform_description"`
}

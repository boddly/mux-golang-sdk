package filters

// Filter provides access to the Mux Data Filters API
type Filter struct {
}

// FilterID is a filter on the list command
type FilterID int

// The possible values for the filterID
const (
	ASN                    FilterID = 0
	Browser                FilterID = 1
	BowserVersion          FilterID = 2
	CDN                    FilterID = 3
	Country                FilterID = 4
	ExperimentName         FilterID = 5
	OperatingSyste         FilterID = 6
	OperatingSystemVersion FilterID = 7
	PlayerName             FilterID = 8
	PlayerSoftware         FilterID = 9
	PlayerSoftwareVersion  FilterID = 10
	PlayerVersion          FilterID = 11
	PrerollAdAssetHostname FilterID = 12
	PrerollAdTagHostname   FilterID = 13
	PrerollPlayed          FilterID = 14
	PrerollRequested       FilterID = 15
	SourceHostname         FilterID = 16
	SourceType             FilterID = 17
	StreamType             FilterID = 18
	SubPropertyID          FilterID = 19
	VideoSeries            FilterID = 20
	VideoTitle             FilterID = 21
)

func (f FilterID) String() string {
	filterNames := [...]string{
		"asn",
		"browser",
		"browser_version",
		"cdn",
		"country",
		"experiment_name",
		"operating_system",
		"operating_system_version",
		"player_name",
		"player_software",
		"player_software_version",
		"player_version",
		"preroll_ad_asset_hostname",
		"preroll_ad_tag_hostname",
		"preroll_played",
		"preroll_requested",
		"source_hostname",
		"source_type",
		"stream_type",
		"sub_property_id",
		"video_series",
		"video_title",
	}

	if f < ASN || f > VideoTitle {
		return "Unknown"
	}

	return filterNames[f]
}

package metrics

// Metrics provides access to the Mux Data Metrics API
type Metric struct {
}

// MetricID is the metric name
type MetricID int

// blah
const (
	AggregateStartupTime           MetricID = 0
	DownscalePercentage            MetricID = 1
	ExitsBeforeVideoStart          MetricID = 2
	MaxDownscalePercentage         MetricID = 3
	MaxUpscalePercentage           MetricID = 4
	PageLoadTime                   MetricID = 5
	PlaybackFailurePercentage      MetricID = 6
	PlaybackFailureScore           MetricID = 7
	PlayerStartupTime              MetricID = 8
	RebufferCount                  MetricID = 9
	RebufferDuration               MetricID = 10
	RebufferFrequency              MetricID = 11
	RebufferPercentage             MetricID = 12
	RebufferScore                  MetricID = 13
	RequestsForFirstPreroll        MetricID = 14
	SeekLatency                    MetricID = 15
	StartupTimeScore               MetricID = 16
	UpscalePercentage              MetricID = 17
	VideoQualityScore              MetricID = 18
	VideoStartupPrerollLoadTime    MetricID = 19
	VideoStartupPrerollRequestTime MetricID = 20
	VideoStartupTime               MetricID = 21
	ViewerExperienceScore          MetricID = 22
)

func (m MetricID) String() string {
	names := [...]string{
		"aggregate_startup_time",
		"downscale_percentage",
		"exits_before_video_start",
		"max_downscale_percentage",
		"max_upscale_percentage",
		"page_load_time",
		"playback_failure_percentage",
		"playback_failure_score",
		"player_startup_time",
		"rebuffer_count",
		"rebuffer_duration",
		"rebuffer_frequency",
		"rebuffer_percentage",
		"rebuffer_score",
		"requests_for_first_preroll",
		"seek_latency",
		"startup_time_score",
		"upscale_percentage",
		"video_quality_score",
		"video_startup_preroll_load_time",
		"video_startup_preroll_request_time",
		"video_startup_time",
		"viewer_experience_score",
	}

	if m < AggregateStartupTime || m > ViewerExperienceScore {
		return "Unknown"
	}

	return names[m]
}

// Dimension is the list of possible metric dimensions
type Dimension int

// blah
const (
	ASN                    Dimension = 0
	Browser                Dimension = 1
	BrowserVersion         Dimension = 2
	CDN                    Dimension = 3
	Country                Dimension = 4
	ExperimentName         Dimension = 5
	OperatingSystem        Dimension = 6
	OperatingSystemVersion Dimension = 7
	PlayerName             Dimension = 8
	PlayerSoftware         Dimension = 9
	PlayerSoftwareVersion  Dimension = 10
	PlayerVersion          Dimension = 11
	PrerollAdAssetHostname Dimension = 12
	PrerollAdTagHostname   Dimension = 13
	PrerollPlayed          Dimension = 14
	PrerollRequested       Dimension = 15
	SourceHostname         Dimension = 16
	SourceType             Dimension = 17
	StreamType             Dimension = 18
	SubPropertyID          Dimension = 19
	VideoSeries            Dimension = 20
	VideoTitle             Dimension = 21
)

func (d Dimension) String() string {
	names := [...]string{
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

	if d < ASN || d > VideoTitle {
		return "Unknown"
	}

	return names[d]
}

// Measurement is the possible metric measurments
type Measurement int

// blah
const (
	P95    Measurement = 0
	Median Measurement = 1
	AVG    Measurement = 2
)

func (m Measurement) String() string {
	names := [...]string{
		"95th",
		"median",
		"avg",
	}

	if m < P95 || m > AVG {
		return "Unknown"
	}

	return names[m]
}

// OrderBy is the value to order the results by
type OrderBy int

// blah
const (
	NegativeImpact OrderBy = 0
	Value          OrderBy = 1
	Views          OrderBy = 2
	Field          OrderBy = 3
)

func (o OrderBy) String() string {
	names := [...]string{
		"negative_impact",
		"value",
		"views",
		"field",
	}

	if o < NegativeImpact || o > Field {
		return "Unknown"
	}

	return names[o]
}

// OrderDirection is the sort order
type OrderDirection int

// blah
const (
	Ascending  OrderDirection = 0
	Descending OrderDirection = 1
)

func (o OrderDirection) String() string {
	names := [...]string{
		"asc",
		"desc",
	}

	if o < Ascending || o > Descending {
		return "Unknown"
	}

	return names[o]
}

// Duration is the ime granularity to group results by.
type Duration int

// blah
const (
	Hour Duration = 0
	Day  Duration = 1
)

func (d Duration) String() string {
	names := [...]string{
		"hour",
		"day",
	}

	if d < Hour || d > Day {
		return "Unknown"
	}

	return names[d]
}

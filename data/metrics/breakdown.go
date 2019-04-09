package metrics

// BreakdownOptions are the set of options that can be supplied to the Get request
type BreakdownOptions struct {
	GroupBy        Dimension
	Measurement    Measurement
	Filters        []string
	Limit          int
	Page           int
	OrderBy        OrderBy
	OrderDirection OrderDirection
	Timeframe      []string
}

// BreakdownResponse represents the total number of views that are applicable to the metric supplied in the request
type BreakdownResponse struct {
	TotalRowCount int32 `json:"total_row_count"`
	Timeframe     []int `json:"timeframe"`
}

// BreakdownData is the data object returned in the BreakdownResponse struct
type BreakdownData struct {
	Views          int32  `json:"views"`
	Value          int    `json:"value"`
	TotalWatchTime int    `json:"total_watch_time"`
	NegativeImpact int    `json:"negative_impact"`
	Field          string `json:"field"`
}

// Breakdown lists the breakdown values for a specific metric
// The value returned in views represents the total number of views that are applicable to the metric supplied;
// some metrics may report different view counts,
// if there are views that lack any information about the metric.
// For instance, in the case that we cannot retrieve player or video size we will not have upscaling and downscaling percentages, nor will we have a Quality Score.
// One specific example of this is that the value in views for video_upscale_percentage may be lower than the values for some other metrics, such as error_rate.
// In the case that you request the error_rate for a specific error (by providing &error= in the parameters), the value in views represents the number of views that encountered that error during playback.
func (m *Metric) Breakdown(id MetricID, opts BreakdownOptions) (*BreakdownResponse, error) {
	// TODO implement
	return nil, nil
}

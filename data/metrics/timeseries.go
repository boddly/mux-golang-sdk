package metrics

// TimeseriesOptions are the set of options that can be supplied to the Get request
type TimeseriesOptions struct {
	Filters     []string
	Measurement Measurement
	Timeframe   []string
	GroupBy     Duration
}

// TimeseriesResponse represents the total number of views that are applicable to the metric supplied in the request
type TimeseriesResponse struct {
	TotalRowCount int32         `json:"total_row_count"`
	Timeframe     []int         `json:"timeframe"`
	Data          []interface{} `json:"data"` // todo nned to handle custom parsing, implement MarshalJSON
}

// TimeseriesData is the data object returned in the TimeseriesResponse struct
type TimeseriesData struct {
}

// Timeseries returns timeseries data for a specific metric
func (m *Metric) Timeseries(metricID MetricID, opts TimeseriesOptions) (*OverallResponse, error) {
	// TODO implement
	return nil, nil
}

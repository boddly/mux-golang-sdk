package metrics

// OverallOptions are the set of options that can be supplied to the Get request
type OverallOptions struct {
	Filters     []string
	Measurement Measurement
	Timeframe   []string
}

// OverallResponse represents the total number of views that are applicable to the metric supplied in the request
type OverallResponse struct {
	TotalRowCount int32       `json:"total_row_count"`
	Timeframe     []int       `json:"timeframe"`
	Data          OverallData `json:"data"`
}

// OverallData is the data object returned in the OverallResponse struct
type OverallData struct {
	Views          int32   `json:"total_views"`
	Value          float64 `json:"value"`
	TotalWatchTime int     `json:"total_watch_time"`
	GlobalValue    string  `json:"global_value"`
}

// Overall returns the overall value for a specific metric, as well as the total view count, watch time, and the Mux Global metric value for the metric.
func (m *Metric) Overall(metricID MetricID, opts OverallOptions) (*OverallResponse, error) {
	// TODO implement
	return nil, nil
}

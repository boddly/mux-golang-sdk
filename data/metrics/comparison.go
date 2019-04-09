package metrics

// ComparisonOptions are the set of options that can be supplied to the Get request
type ComparisonOptions struct {
	Filters   []string
	Dimension Dimension
	Value     string
	Timeframe []string
}

// ComparisonResponse represents the total number of views that are applicable to the metric supplied in the request
type ComparisonResponse struct {
	TotalRowCount int32            `json:"total_row_count"`
	Timeframe     []int            `json:"timeframe"`
	Data          []ComparisonData `json:"data"`
}

// ComparisonData is the data object returned in the ComparisonResponse struct
type ComparisonData struct {
	Views          int32  `json:"views"`
	Value          int    `json:"value"`
	TotalWatchTime int    `json:"total_watch_time"`
	NegativeImpact int    `json:"negative_impact"`
	Field          string `json:"field"`
}

// Comparison lists all of the values across every breakdown for a specific metric
func (m *Metric) Comparison(opts ComparisonOptions) (*ComparisonResponse, error) {
	// TODO implement
	return nil, nil
}

package metrics

// InsightOptions are the set of options that can be supplied to the Get request
type InsightOptions struct {
	Measurement Measurement
	OrderBy     OrderBy
	Timeframe   []string
}

// InsightResponse representsthe list of insights for a metric
type InsightResponse struct {
	TotalRowCount int32         `json:"total_row_count"`
	Timeframe     []int         `json:"timeframe"`
	Data          []InsightData `json:"data"`
}

// InsightData is the data object returned in the InsightResponse struct
type InsightData struct {
	TotalWatchTime      int     `json:"total_watch_time"`
	TotalViews          int     `json:"total_views"`
	TotalRowCount       int     `json:"total_row_count"`
	NegativeImpactScore float64 `json:"negative_impact_score"`
	Metric              string  `json:"metric"`
	FilterValue         string  `json:"filter_value"`
	FilterColumn        string  `json:"filter_column"`
}

// Insights Returns a list of insights for a metric.
// These are the worst performing values across all breakdowns sorted by how much they negatively impact a specific metric.
func (m *Metric) Insights(metricID MetricID, opts InsightOptions) (*InsightResponse, error) {
	// TODO implement
	return nil, nil
}

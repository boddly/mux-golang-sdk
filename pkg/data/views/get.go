package views

// GetResponse contains the data from a single video view
type GetResponse struct {
	TotalRowCount int       `json:"total_row_count"`
	Timeframe     []string  `json:"timeframe"`
	Data          VideoData `json:"data"`
}

// Get returns the details for a single video view
func (v *VideoViews) Get(ID string) (*GetResponse, error) {
	return nil, nil
}

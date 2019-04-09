package exports

// Export provides access to the Mux Data Exports AP
type Export struct {
}

// ListResp contains the response from the Exports List API
type ListResp struct {
	TotalRowCount int32    `json:"total_row_count"`
	Timeframe     int64    `json:"timeframe"`
	Data          []string `json:"data"`
}

// List Lists the available video view exports along with URLs to retrieve them
func (e *Export) List() (*ListResp, error) {
	// TODO implement
	return nil, nil
}

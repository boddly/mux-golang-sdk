package filters

// GetReq is the struct containing parmeters for the Filter.Get method
type GetReq struct {
	TimeFrame []string `json:"timeframe"`
	Filters   []string `json:"filters"`
	Page      int32    `json:"page"`
	Limit     int32    `json:"limit"`
}

// GetResp is the response from the Filters.Get Request
type GetResp struct {
	TotalRowCount int32     `json:"total_row_count"`
	Timeframe     []int     `json:"timeframe"`
	Data          []GetData `json:"data"`
}

// GetData is the data object returned in the GetResp struct
type GetData struct {
	Value      string `json:"value"`
	TotalCount int32  `json:"total_count"`
}

// Get lists the values for a filter along with a total count of related views
func (f *Filter) Get(ID, query GetReq) (*GetResp, error) {
	// TODO implement
	return nil, nil
}

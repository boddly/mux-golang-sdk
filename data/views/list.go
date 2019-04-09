package views

import "github.com/boddly/mux-golang-sdk/data/metrics"

// ListResponse is the struct containing data from the List request
type ListResponse struct {
	TotalRowCount int      `json:"total_row_count"`
	Timeframe     []string `json:"timeframe"`
	Data          ListData `json:"data"`
}

// ListOptions is the set of options to pass into a list request
type ListOptions struct {
	Filters        []string
	ErrorID        string
	ViewerID       string
	Limit          int
	OrderDirection metrics.OrderDirection
	Timeframe      []string
}

// List returns a list of video views for a property that occurred within the specified timeframe.
// Results are ordered by view_end, according to what you provide for order_direction.
func (v *VideoViews) List() (*ListResponse, error) {
	return nil, nil
}

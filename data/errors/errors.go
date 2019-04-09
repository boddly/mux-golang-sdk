package errors

import "time"

// Error provides access to the Mux Data Errors API
type Error struct {
}

// Options contains the data needed to filter errors
type Options struct {
	Timeframe map[string]string `json:"timeframe"`
	Filters   map[string]string `json:"filters"`
}

// ListResp contains the response from the Data List Errors endpoint
type ListResp struct {
	TotalRowCount int64      `json:"total_row_count"`
	Timeframe     []int64    `json:"timeframe"`
	Data          []ListData `json:"data"`
}

// ListData is the data returned from the Data List Errors endpoint
type ListData struct {
	Percentage  float64   `json:"percentage"`
	Notes       string    `json:"notes"`
	Message     string    `json:"message"`
	LastSeen    time.Time `json:"last_seen"`
	ID          int64     `json:"id"`
	Description string    `json:"description"`
	Count       int32     `json:"count"`
	Code        int32     `json:"code"`
}

// List returns a list of playback errors filtered by the windows operating system
func (e *Error) List(Options) (*ListResp, error) {
	// TODO: implement
	return nil, nil
}

package filters

// ListResp is the response struct from the Filters.List operation
type ListResp struct {
}

// ListData is the data object returned with the list response
type ListData struct {
	Basic    []string `json:"basic"`
	Advanced []string `json:"advanced"`
}

// List lists all the filters broken out into basic and advanced
func (f *Filter) List() (*ListResp, error) {
	// TODO implement
	return nil, nil
}

package signing

// ListResponse is the response struct for the List request
type ListResponse struct {
	Data []URLSigningKey `json:"data"`
}

// List URL signing keys
func (s *URLSigningKey) List(limit, page int) ([]*URLSigningKey, error) {
	return nil, nil
}

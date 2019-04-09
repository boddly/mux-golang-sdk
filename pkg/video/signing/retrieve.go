package signing

// Retrieve the details of a URL signing key that has previously been created.
// Supply the unique signing key ID that was returned from your previous request, and Mux will return the corresponding signing key information.
// The private key is not returned in this response.
func (s *URLSigningKey) Retrieve(ID string) (*URLSigningKey, error) {
	return nil, nil
}

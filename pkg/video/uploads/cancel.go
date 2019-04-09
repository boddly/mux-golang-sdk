package uploads

// Cancel will not invalidate any signed upload URLs, but it will ensure that an asset is not created if a file is uploaded.
// This can only be called for assets still in the waiting state.
func (d *DirectUpload) Cancel(ID string) (*DirectUpload, error) {
	return nil, nil
}

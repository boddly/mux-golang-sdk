package livestreams

// Finish signals a live stream is finished
// When a live streamed event is finished you can use the live stream complete API to signal this to Mux,
// so that the recording of the live stream (a Mux Asset) is ready to play more quickly. Using this API is not required.
// When not used Mux will wait for the reconnect_window (default of 60 seconds) before considering the
// live streamed event finished and making the recording ready to play.
// One use case of this API would be to add a "Live Stream Complete" button to your UI for the person streaming, that they can click when finished.
func (ls *LiveStream) Finish(ID string) (*LiveStream, error) {
	return nil, nil
}

package image

// ThumbnailSizing is the set of query params used to appropriately size thumbnails
type ThumbnailSizing struct {
	Time    float64
	Width   int32
	Height  int32
	Rotate  int32
	FitMode string // preserve, stretch, smartcrop, pad
	FlipV   bool
	FlipH   bool
}

// Thumbnail pulls images from a Mux Video asset in real time.
// Any frame of an asset is available as a PNG or JPG image, to use as a thumbnail or poster image
// Thumbnails are usually used as a smaller representation of a video.
// Poster images are typically used in place of the video before the video is loaded or begins playing.
// From the perspective of Mux Video, thumbnails and poster images are functionally the same,
// and these docs will typically use the term "thumbnail" to cover any image that is extracted from a video.
func (i *Image) Thumbnail(id, ext string, opts ThumbnailSizing) error {
	return nil
}

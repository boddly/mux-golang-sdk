package image

import "fmt"

// Thumbnail pulls images from a Mux Video asset in real time.
// Any frame of an asset is available as a PNG or JPG image, to use as a thumbnail or poster image
// Thumbnails are usually used as a smaller representation of a video.
// Poster images are typically used in place of the video before the video is loaded or begins playing.
// From the perspective of Mux Video, thumbnails and poster images are functionally the same,
// and these docs will typically use the term "thumbnail" to cover any image that is extracted from a video.
type Thumbnail interface {
	SetTime(float64) Thumbnail
	SetWidth(int32) Thumbnail
	SetHeight(int32) Thumbnail
	SetRotate(int32) Thumbnail
	SetFitmode(string) Thumbnail
	SetFlipV(bool) Thumbnail
	SetFlipH(bool) Thumbnail
	fmt.Stringer
}

type thumbnailOptions struct {
	id        string
	extension string
	time      float64
	width     int32
	height    int32
	rotate    int32
	fitMode   string // preserve, stretch, smartcrop, pad
	flipV     bool
	flipH     bool
}

// NewThumbnail returns the thumbnail with default values
// use the setters to override the defaults
func NewThumbnail(id, ext string) (Thumbnail, error) {
	if ext != "png" && ext != "jpg" {
		return nil, fmt.Errorf("%s is an invalid extension, extensions must be either 'png' or 'jpg'", ext)
	}
	return &thumbnailOptions{
		id:        id,
		extension: ext,
		time:      -1,
		width:     -1,
		height:    -1,
		rotate:    -1,
		fitMode:   "preserve",
		flipV:     false,
		flipH:     false,
	}, nil
}

func (t *thumbnailOptions) String() string {
	s := fmt.Sprintf("https://image.mux.com/%s/thumbnail.%s?fit_mode=%s&flip_v=%t&flip_h=%t", t.id, t.extension, t.fitMode, t.flipV, t.flipH)

	if t.time >= 0 {
		s = fmt.Sprintf("%s&time=%f", s, t.time)
	}
	if t.width >= 0 {
		s = fmt.Sprintf("%s&width=%d", s, t.width)
	}
	if t.height >= 0 {
		s = fmt.Sprintf("%s&height=%d", s, t.height)
	}
	if t.rotate >= 0 {
		s = fmt.Sprintf("%s&rotate=%d", s, t.rotate)
	}

	return s
}

func (t *thumbnailOptions) SetTime(s float64) Thumbnail {
	t.time = s
	return t
}

func (t *thumbnailOptions) SetWidth(s int32) Thumbnail {
	t.width = s
	return t
}

func (t *thumbnailOptions) SetHeight(s int32) Thumbnail {
	t.height = s
	return t
}

func (t *thumbnailOptions) SetRotate(s int32) Thumbnail {
	if s != 90 && s != 180 && s != 270 {
		return t
	}
	t.rotate = s
	return t
}

func (t *thumbnailOptions) SetFitmode(s string) Thumbnail {
	if s != "preserve" && s != "stretch" && s != "crop" && s != "smartcrop" && s != "pad" {
		return t
	}

	t.fitMode = s
	return t
}

func (t *thumbnailOptions) SetFlipV(s bool) Thumbnail {
	t.flipV = s
	return t
}

func (t *thumbnailOptions) SetFlipH(s bool) Thumbnail {
	t.flipH = s
	return t
}

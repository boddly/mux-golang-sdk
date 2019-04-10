package image

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func assertThumbnailType(t *testing.T, tn Thumbnail) *thumbnailOptions {
	s, ok := tn.(*thumbnailOptions)
	require.True(t, ok)

	return s
}

func TestNewThumbnail(t *testing.T) {
	cases := []struct {
		id          string
		ext         string
		expectedEXT string
		err         error
	}{
		{
			id:  "foo",
			ext: "bar",
			err: errors.New("bar is an invalid extension, extensions must be either 'png' or 'jpg'"),
		},
		{
			id:          "foo",
			ext:         "png",
			expectedEXT: "png",
			err:         nil,
		},
		{
			id:          "foo",
			ext:         "JPG",
			expectedEXT: "jpg",
			err:         nil,
		},
	}

	for _, v := range cases {
		tn, err := NewThumbnail(v.id, v.ext)
		if v.err != nil {
			require.EqualError(t, err, v.err.Error())
			continue
		}

		s := assertThumbnailType(t, tn)

		assert.Equal(t, thumbnailOptions{
			id:        v.id,
			extension: v.expectedEXT,
			time:      -1,
			width:     -1,
			height:    -1,
			rotate:    -1,
			fitMode:   "preserve",
			flipV:     false,
			flipH:     false,
		}, *s)
	}
}

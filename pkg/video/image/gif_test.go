package image

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func assertGIFType(t *testing.T, g GIF) *options {
	s, ok := g.(*options)
	require.True(t, ok)

	return s
}

func TestNewGif(t *testing.T) {
	gif, err := NewGif("foo")
	require.NoError(t, err)

	s := assertGIFType(t, gif)

	assert.Equal(t, options{
		id:     "foo",
		start:  0,
		end:    5,
		width:  320,
		height: -1,
		fps:    15,
	}, *s)
}

func TestSetStart(t *testing.T) {
	cases := []struct {
		gif   GIF
		input float64
	}{
		{
			gif:   func() GIF { gif, err := NewGif("foo"); require.NoError(t, err); return gif }(),
			input: 6,
		},
		{
			gif:   func() GIF { gif, err := NewGif("foo"); require.NoError(t, err); return gif }(),
			input: 0,
		},
		{
			gif:   func() GIF { gif, err := NewGif("foo"); require.NoError(t, err); return gif }(),
			input: 1.2,
		},
	}

	for _, v := range cases {
		v.gif.SetStart(v.input)

		gif := assertGIFType(t, v.gif)
		assert.Equal(t, v.input, gif.start)
	}
}

func TestSetEnd(t *testing.T) {
	cases := []struct {
		gif   GIF
		input float64
		start float64
	}{
		{
			gif:   func() GIF { gif, err := NewGif("foo"); require.NoError(t, err); return gif }(),
			input: 6,
			start: 0,
		},
		{
			gif:   func() GIF { gif, err := NewGif("foo"); require.NoError(t, err); return gif }(),
			input: 0,
			start: 0,
		},
		{
			gif:   func() GIF { gif, err := NewGif("foo"); require.NoError(t, err); return gif }(),
			input: 1.2,
			start: 0,
		},
		{
			gif:   func() GIF { gif, err := NewGif("foo"); require.NoError(t, err); gif.SetStart(15); return gif }(),
			input: 50,
			start: 40,
		},
	}

	for _, v := range cases {
		v.gif.SetEnd(v.input)

		gif := assertGIFType(t, v.gif)
		assert.Equal(t, v.input, gif.end)
		assert.Equal(t, v.start, gif.start)
	}
}

func TestSetWidth(t *testing.T) {
	cases := []struct {
		gif      GIF
		input    int32
		expected int32
	}{
		{
			gif:      func() GIF { gif, err := NewGif("foo"); require.NoError(t, err); return gif }(),
			input:    720,
			expected: 640,
		},
		{
			gif:      func() GIF { gif, err := NewGif("foo"); require.NoError(t, err); return gif }(),
			input:    -1,
			expected: 320,
		},
		{
			gif:      func() GIF { gif, err := NewGif("foo"); require.NoError(t, err); return gif }(),
			input:    90,
			expected: 90,
		},
	}

	for _, v := range cases {
		v.gif.SetWidth(v.input)

		gif := assertGIFType(t, v.gif)
		assert.Equal(t, v.expected, gif.width)
	}
}

func TestSetHeight(t *testing.T) {
	cases := []struct {
		gif      GIF
		input    int32
		expected int32
	}{
		{
			gif:      func() GIF { gif, err := NewGif("foo"); require.NoError(t, err); return gif }(),
			input:    720,
			expected: 640,
		},
		{
			gif:      func() GIF { gif, err := NewGif("foo"); require.NoError(t, err); return gif }(),
			input:    0,
			expected: -1,
		},
		{
			gif:      func() GIF { gif, err := NewGif("foo"); require.NoError(t, err); return gif }(),
			input:    -10,
			expected: -1,
		},
		{
			gif:      func() GIF { gif, err := NewGif("foo"); require.NoError(t, err); return gif }(),
			input:    90,
			expected: 90,
		},
	}

	for _, v := range cases {
		v.gif.SetHeight(v.input)

		gif := assertGIFType(t, v.gif)
		assert.Equal(t, v.expected, gif.height)
	}
}

func TestSetFPS(t *testing.T) {
	cases := []struct {
		gif      GIF
		input    int32
		expected int32
	}{
		{
			gif:      func() GIF { gif, err := NewGif("foo"); require.NoError(t, err); return gif }(),
			input:    45,
			expected: 30,
		},
		{
			gif:      func() GIF { gif, err := NewGif("foo"); require.NoError(t, err); return gif }(),
			input:    0,
			expected: 15,
		},
		{
			gif:      func() GIF { gif, err := NewGif("foo"); require.NoError(t, err); return gif }(),
			input:    -1,
			expected: 15,
		},
		{
			gif:      func() GIF { gif, err := NewGif("foo"); require.NoError(t, err); return gif }(),
			input:    20,
			expected: 20,
		},
	}

	for _, v := range cases {
		v.gif.SetFPS(v.input)

		gif := assertGIFType(t, v.gif)
		assert.Equal(t, v.expected, gif.fps)
	}
}

func TestString(t *testing.T) {
	cases := []struct {
		gif      GIF
		expected string
	}{
		{
			gif:      func() GIF { gif, err := NewGif("foo"); require.NoError(t, err); return gif }(),
			expected: "https://image.mux.com/foo/animated.gif?&width=320&start=0.0&end=5.0&fps=15",
		},
		{
			gif:      func() GIF { gif, err := NewGif("foo"); require.NoError(t, err); gif.SetHeight(90); return gif }(),
			expected: "https://image.mux.com/foo/animated.gif?height=90&width=320&start=0.0&end=5.0&fps=15",
		},
	}

	for _, v := range cases {
		assert.Equal(t, v.expected, v.gif.String())
	}
}

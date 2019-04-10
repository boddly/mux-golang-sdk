package image

import (
	"errors"
	"fmt"
)

// GIF allows you to generate short animated GIFs from a video.
type GIF interface {
	SetStart(float64) GIF
	SetEnd(float64) GIF
	SetWidth(int32) GIF
	SetHeight(int32) GIF
	SetFPS(int32) GIF
	fmt.Stringer
}

type options struct {
	id     string
	start  float64
	end    float64
	width  int32
	height int32
	fps    int32
}

// NewGif returns the gif with default values
// use the setters to override the defaults
func NewGif(id string) (GIF, error) {
	if id == "" {
		return nil, errors.New("invalid id")
	}

	return &options{
		id:     id,
		start:  0,
		end:    5,
		width:  320,
		height: -1,
		fps:    15,
	}, nil
}

func (g *options) String() string {
	s := fmt.Sprintf("https://image.mux.com/%s/animated.gif?", g.id)

	if g.height != -1 {
		s = fmt.Sprintf("%sheight=%d", s, g.height)
	}
	return fmt.Sprintf("%s&width=%d&start=%f&end=%f&fps=%d", s, g.width, g.start, g.end, g.fps)
}

func (g *options) SetStart(s float64) GIF {
	g.start = s
	return g
}

func (g *options) SetEnd(s float64) GIF {
	g.end = s

	if (g.end - g.start) > 10 {
		g.start = g.end - 10
	}

	return g
}

func (g *options) SetWidth(s int32) GIF {
	if s > 640 {
		s = 640
	}

	if s < 0 {
		return g
	}

	g.width = s

	return g
}

func (g *options) SetHeight(s int32) GIF {
	if s > 640 {
		s = 640
	}

	if s <= 0 {
		return g
	}

	g.height = s

	return g
}

func (g *options) SetFPS(s int32) GIF {
	if s > 30 {
		s = 30
	}

	if s <= 0 {
		return g
	}

	g.fps = s

	return g
}

package data

import (
	"github.com/boddly/mux-golang-sdk/data/errors"
	"github.com/boddly/mux-golang-sdk/data/exports"
	"github.com/boddly/mux-golang-sdk/data/filters"
	"github.com/boddly/mux-golang-sdk/data/metrics"
	"github.com/boddly/mux-golang-sdk/data/views"
)

// Data provides access to the mux Data API
type Data struct {
	accessToken string
	secret      string
	Errors      errors.Error
	Exports     exports.Export
	Filters     filters.Filter
	Metrics     metrics.Metric
	Views       views.VideoViews
}

// New is a data Constructor
func New(accessToken, secret string) *Data {
	return &Data{
		accessToken: accessToken,
		secret:      secret,
		Errors:      errors.Error{},
		Exports:     exports.Export{},
		Filters:     filters.Filter{},
		Metrics:     metrics.Metric{},
		Views:       views.VideoViews{},
	}
}

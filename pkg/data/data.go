package data

import (
	"github.com/boddly/mux-golang-sdk/pkg/data/errors"
	"github.com/boddly/mux-golang-sdk/pkg/data/exports"
	"github.com/boddly/mux-golang-sdk/pkg/data/filters"
	"github.com/boddly/mux-golang-sdk/pkg/data/metrics"
	"github.com/boddly/mux-golang-sdk/pkg/data/views"
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

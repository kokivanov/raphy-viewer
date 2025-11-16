package core

import "context"

type QueryAllOptions struct {
	Page     int64    `json:"page,omitempty"`
	PerPage  int64    `json:"perPage,omitempty"`
	Sort     SortType `json:"sort,omitempty"`
	StatusIn []string `json:"statusIn,omitempty"`
	GenreIn  []string `json:"genreIn,omitempty"`
	Search   string   `json:"search,omitempty"`
	TagIn    []string `json:"tagIn,omitempty"`
	IsAdult  bool     `json:"isAdult,omitempty"`
}

type ApiProvider interface {
	QueryAll(ctx *context.Context, options *QueryAllOptions) (any, error)
	QueryById(ctx *context.Context, id int64) (any, error)
}

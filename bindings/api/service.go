package api

import (
	"raphyviewer/internals/api/core"

	"github.com/wailsapp/wails/v2/pkg/logger"
)

// TODO: Add cache service to automatically cache data
type ApiService struct {
	p core.ApiProvider
	l logger.Logger
}

func NewApiService(provider core.ApiProvider, l logger.Logger) (*ApiService, error) {
	return &ApiService{
		p: provider,
		l: l,
	}, nil
}

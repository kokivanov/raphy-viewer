package api

import (
	"raphyviewer/internals/api/core"
	"raphyviewer/internals/core/cache"

	"github.com/wailsapp/wails/v2/pkg/logger"
)

// TODO: Add cache service to automatically cache data
type ApiService struct {
	p core.ApiProvider
	c *cache.CacheManager
	l logger.Logger
}

func NewApiService(provider core.ApiProvider, l logger.Logger, c *cache.CacheManager) (*ApiService, error) {
	return &ApiService{
		p: provider,
		l: l,
		c: c,
	}, nil
}

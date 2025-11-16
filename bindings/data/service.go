package data

import (
	"raphyviewer/internals/core/localdata"

	"github.com/wailsapp/wails/v2/pkg/logger"
)

// TODO: Add cache service to automatically cache data
type DataService struct {
	dm *localdata.DataManager
	l  logger.Logger
}

func NewDataService(dataDir string, lg logger.Logger) (*DataService, error) {
	dm, err := localdata.NewDataManager(dataDir)
	if err != nil {
		return nil, err
	}

	return &DataService{
		dm: dm,
		l:  lg,
	}, nil
}

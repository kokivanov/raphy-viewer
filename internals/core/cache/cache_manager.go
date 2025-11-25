package cache

import (
	"context"
	"database/sql"
	"sync"
	"time"

	"net/http"
	"os"
	"path"

	_ "github.com/mattn/go-sqlite3"
	"github.com/wailsapp/wails/v2/pkg/logger"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/sqlitedialect"
	"github.com/uptrace/bun/driver/sqliteshim"
)

type CacheManager struct {
	dir string

	imgDir     string
	contentDir string

	db                   *bun.DB
	dbMutex              sync.Mutex
	CacheLifetime        time.Duration
	RequestCacheLifetime time.Duration
	imageCacheMaxSize    int64

	lg logger.Logger

	Client *http.Client
	ctx    context.Context
}

func NewCacheManager(cacheDir string, lg logger.Logger) (*CacheManager, error) {
	ctx := context.Background()

	dbPath := path.Join(cacheDir, "cache.db")

	imgDir := path.Join(cacheDir, "images")
	contentDir := path.Join(cacheDir, "media")

	if err := os.MkdirAll(path.Join(imgDir, "other"), 0o775); err != nil {
		return nil, err
	}

	if err := os.MkdirAll(contentDir, 0o775); err != nil {
		return nil, err
	}

	sqldb, err := sql.Open(sqliteshim.ShimName, dbPath)
	if err != nil {
		return nil, err
	}

	db := bun.NewDB(sqldb, sqlitedialect.New())

	_, err = db.NewCreateTable().Model((*ImageData)(nil)).IfNotExists().Exec(ctx)
	if err != nil {
		return nil, err
	}

	_, err = db.NewCreateTable().Model((*MetaData)(nil)).IfNotExists().Exec(ctx)
	if err != nil {
		return nil, err
	}

	_, err = db.NewCreateTable().Model((*ShortMetaData)(nil)).IfNotExists().Exec(ctx)
	if err != nil {
		return nil, err
	}

	_, err = db.NewCreateTable().Model((*RequestCache)(nil)).IfNotExists().Exec(ctx)
	if err != nil {
		return nil, err
	}

	cm := &CacheManager{
		db:                   db,
		dir:                  cacheDir,
		imgDir:               imgDir,
		contentDir:           contentDir,
		RequestCacheLifetime: 60 * time.Second,
		CacheLifetime:        7 * 24 * time.Hour,
		imageCacheMaxSize:    2 * 1024 * 1024 * 1024,
		Client: &http.Client{
			Timeout: 5 * time.Second,
		},
		lg:      lg,
		ctx:     ctx,
		dbMutex: sync.Mutex{},
	}

	cm.PurgeSortMetaDatum()

	return cm, nil
}

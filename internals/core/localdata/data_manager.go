package localdata

import (
	"context"
	"database/sql"
	"os"
	"path"
	"sync"

	_ "github.com/mattn/go-sqlite3"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/sqlitedialect"
	"github.com/uptrace/bun/driver/sqliteshim"
)

type DataManager struct {
	dir string

	db      *bun.DB
	ctx     context.Context
	dbMutex sync.Mutex
}

func NewDataManager(dataDir string) (*DataManager, error) {
	ctx := context.Background()

	err := os.MkdirAll(dataDir, 0o775)
	if err != nil {
		return nil, err
	}

	dataPath := path.Join(dataDir, "data.db")

	sqldb, err := sql.Open(sqliteshim.ShimName, dataPath)
	if err != nil {
		return nil, err
	}

	db := bun.NewDB(sqldb, sqlitedialect.New())

	_, err = db.NewCreateTable().Model((*MediaData)(nil)).IfNotExists().Exec(ctx)
	if err != nil {
		return nil, err
	}

	return &DataManager{
		dir:     dataDir,
		db:      db,
		ctx:     ctx,
		dbMutex: sync.Mutex{},
	}, nil
}

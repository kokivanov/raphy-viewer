package localdata

import (
	"time"

	"github.com/uptrace/bun"
)

func (dm *DataManager) AddFavorite(id int64) error {
	dm.dbMutex.Lock()
	defer dm.dbMutex.Unlock()

	favTime := time.Now().Unix()

	item := MediaData{
		ID:         id,
		IsFavorite: true,
		FavAt:      &favTime,
	}

	_, err := dm.db.NewInsert().Model(&item).On("CONFLICT (id) DO UPDATE").Set("isFavorite = ?", 1).Set("FavAt = ?", favTime).Exec(dm.ctx)
	return err
}

func (dm *DataManager) RemoveFavorite(id int64) error {
	dm.dbMutex.Lock()
	defer dm.dbMutex.Unlock()

	item := MediaData{
		ID:         id,
		IsFavorite: true,
		FavAt:      nil,
	}

	_, err := dm.db.NewUpdate().Model(&item).Where("id = ?", id).Exec(dm.ctx)
	return err
}

func (dm *DataManager) FavoritesJoin(ids []int64) ([]MediaData, error) {
	dm.dbMutex.Lock()
	defer dm.dbMutex.Unlock()

	var favJoin []MediaData

	err := dm.db.NewSelect().Model(&favJoin).Where("id IN (?)", bun.In(ids)).Scan(dm.ctx)
	if err != nil {
		return nil, err
	}

	return favJoin, nil
}

func (dm *DataManager) GetFavorite(id int64) (*MediaData, error) {
	dm.dbMutex.Lock()
	defer dm.dbMutex.Unlock()

	var fav MediaData
	err := dm.db.NewSelect().Model(&fav).Where("id = ?", id).Scan(dm.ctx)
	if err != nil {
		return nil, err
	}

	return &fav, nil
}

func (dm *DataManager) GetFavorites(page int, perPage int) (*[]MediaData, error) {
	dm.dbMutex.Lock()
	defer dm.dbMutex.Unlock()

	var favs []MediaData

	err := dm.db.NewSelect().Model(&favs).Limit(perPage).Offset((page - 1) * perPage).Scan(dm.ctx)
	if err != nil {
		return nil, err
	}

	return &favs, nil
}

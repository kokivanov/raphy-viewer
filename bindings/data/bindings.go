package data

import "raphyviewer/internals/core/localdata"

func (ds *DataService) AddFavorite(id int64) bool {
	err := ds.dm.AddFavorite(id)
	if err != nil {
		ds.l.Error(err.Error())
		return false
	}

	return true
}

func (ds *DataService) RemoveFavorite(id int64) bool {
	err := ds.dm.RemoveFavorite(id)
	if err != nil {
		ds.l.Error(err.Error())
		return true
	}

	return false
}

func (ds *DataService) GetFavoritesJoin(ids []int64) []localdata.MediaData {
	data, err := ds.dm.FavoritesJoin(ids)
	if err != nil {
		ds.l.Error(err.Error())
		return nil
	}

	return data
}

func (ds *DataService) IsFavorite(id int64) bool {
	fav, err := ds.dm.GetFavorite(id)
	if err != nil {
		return false
	}

	return fav.ID == id
}

package cache

import (
	"database/sql"
	"fmt"
	"os"
	"path"
	"strconv"
	"strings"
	"time"

	"github.com/uptrace/bun"
)

type RequestType string

type RequestOptions interface {
}

const (
	Insert RequestType = "INSERT"
	Delete RequestType = "DELETE"
	Select RequestType = "SELECT"
)

func (cm *CacheManager) AddMetaData(contentId int64, data string) error {
	cm.dbMutex.Lock()
	defer cm.dbMutex.Unlock()

	reqTime := time.Now().Unix()
	expTime := time.Now().Add(cm.CacheLifetime).Unix()
	var err error

	dataToAdd := MetaData{
		ID:        contentId,
		CreatedAt: reqTime,
		ExpiresAt: expTime,
		Size:      len(data),
		Data:      data,
	}

	_, err = cm.db.NewInsert().Model(&dataToAdd).Exec(cm.ctx)
	return err
}

func (cm *CacheManager) GetMetaData(id int64) (*MetaData, error) {
	cm.dbMutex.Lock()
	defer cm.dbMutex.Unlock()

	var metaData MetaData

	err := cm.db.NewSelect().Model(&metaData).Where("id = ?", id).Scan(cm.ctx)

	if err == nil {
		return &metaData, nil
	} else {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		return nil, err
	}
}

func (cm *CacheManager) RemoveMetaData(id int64) error {
	cm.dbMutex.Lock()
	defer cm.dbMutex.Unlock()

	_, err := cm.db.NewDelete().Model((*MetaData)(nil)).Where("id = ?", id).Exec(cm.ctx)
	if err != nil {
		return err
	}

	return nil
}

func (cm *CacheManager) RemoveMetaDatum(ids []int64) error {
	cm.dbMutex.Lock()
	defer cm.dbMutex.Unlock()

	_, err := cm.db.NewDelete().Model((*MetaData)(nil)).Where("hash IN (?)", bun.In(ids)).Exec(cm.ctx)
	if err != nil {
		return err
	}

	return nil
}

func (cm *CacheManager) PurgeMetaDatum() error {
	cm.dbMutex.Lock()
	defer cm.dbMutex.Unlock()

	_, err := cm.db.NewDelete().Model((*MetaData)(nil)).Exec(cm.ctx)
	if err != nil {
		return err
	}

	return nil
}

// TODO: Assign order

func (cm *CacheManager) AddRequestCache(hash string, data string) error {
	cm.dbMutex.Lock()
	defer cm.dbMutex.Unlock()

	reqTime := time.Now().Unix()
	expTime := time.Now().Add(cm.RequestCacheLifetime).Unix()
	var err error

	dataToAdd := RequestCache{
		Hash:      hash,
		CreatedAt: reqTime,
		ExpiresAt: expTime,
		Data:      data,
	}

	_, err = cm.db.NewInsert().Model(&dataToAdd).Exec(cm.ctx)
	return err
}

func (cm *CacheManager) GetRequestCache(hash string) (*RequestCache, error) {
	cm.dbMutex.Lock()
	defer cm.dbMutex.Unlock()

	var metaData RequestCache

	err := cm.db.NewSelect().Model(&metaData).Where("hash = ?", hash).Scan(cm.ctx)

	if err == nil {
		return &metaData, nil
	} else {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		return nil, err
	}
}

func (cm *CacheManager) RemoveRequestCache(hash string) error {
	cm.dbMutex.Lock()
	defer cm.dbMutex.Unlock()

	_, err := cm.db.NewDelete().Model((*RequestCache)(nil)).Where("hash = ?", hash).Exec(cm.ctx)
	if err != nil {
		return err
	}

	return nil
}

func (cm *CacheManager) RemoveRequestCaches(hashes []string) error {
	cm.dbMutex.Lock()
	defer cm.dbMutex.Unlock()

	_, err := cm.db.NewDelete().Model((*RequestCache)(nil)).Where("hash IN (?)", bun.In(hashes)).Exec(cm.ctx)
	if err != nil {
		return err
	}

	return nil
}

func (cm *CacheManager) PurgeRequestCache() error {
	cm.dbMutex.Lock()
	defer cm.dbMutex.Unlock()

	_, err := cm.db.NewDelete().Model((*RequestCache)(nil)).Exec(cm.ctx)
	if err != nil {
		return err
	}

	return nil
}

func (cm *CacheManager) AddShortMetaData(contentId int64, data string) error {
	cm.dbMutex.Lock()
	defer cm.dbMutex.Unlock()

	reqTime := time.Now().Unix()
	expTime := time.Now().Add(cm.CacheLifetime).Unix()
	var err error

	dataToAdd := ShortMetaData{
		ID:        contentId,
		CreatedAt: reqTime,
		ExpiresAt: expTime,
		Data:      data,
	}

	_, err = cm.db.NewInsert().Model(&dataToAdd).Exec(cm.ctx)
	return err
}

func (cm *CacheManager) AddShortMetaDatum(d []CachedShortMetaData) error {
	cm.dbMutex.Lock()
	defer cm.dbMutex.Unlock()

	reqTime := time.Now().Unix()
	expTime := time.Now().Add(cm.CacheLifetime).Unix()
	var err error

	datum := make([]ShortMetaData, len(d))

	for i := range datum {
		datum[i] = ShortMetaData{
			ID:        d[i].ID,
			CreatedAt: reqTime,
			ExpiresAt: expTime,
			Data:      d[i].Data,
		}
	}

	_, err = cm.db.NewInsert().Model(&datum).On("CONFLICT (id) DO UPDATE").Set("createdAt = EXCLUDED.createdAt, expiresAt = EXCLUDED.expiresAt").Exec(cm.ctx)
	return err
}

func (cm *CacheManager) GetShortMetaData(id int64) (*ShortMetaData, error) {
	cm.dbMutex.Lock()
	defer cm.dbMutex.Unlock()

	var metaData ShortMetaData

	err := cm.db.NewSelect().Model(&metaData).Where("id = ?", id).Scan(cm.ctx)

	if err == nil {
		return &metaData, nil
	} else {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		return nil, err
	}
}

func (cm *CacheManager) GetShortMetaDatum(ids []int64) (*[]ShortMetaData, error) {
	cm.dbMutex.Lock()
	defer cm.dbMutex.Unlock()

	var metaData []ShortMetaData

	fmt.Printf("%v", ids)
	err := cm.db.NewSelect().Model(&metaData).Where("id IN (?)", bun.In(ids)).Scan(cm.ctx)

	if err == nil {
		return &metaData, nil
	} else {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		return nil, err
	}
}

func (cm *CacheManager) RemoveShortMetaData(id int64) error {
	cm.dbMutex.Lock()
	defer cm.dbMutex.Unlock()

	_, err := cm.db.NewDelete().Model((*ShortMetaData)(nil)).Where("id = ?", id).Exec(cm.ctx)
	if err != nil {
		return err
	}

	return nil
}

func (cm *CacheManager) RemoveShortMetaDatum(ids []int64) error {
	cm.dbMutex.Lock()
	defer cm.dbMutex.Unlock()

	_, err := cm.db.NewDelete().Model((*ShortMetaData)(nil)).Where("hash IN (?)", bun.In(ids)).Exec(cm.ctx)
	if err != nil {
		return err
	}

	return nil
}

func (cm *CacheManager) PurgeSortMetaDatum() error {
	cm.dbMutex.Lock()
	defer cm.dbMutex.Unlock()

	_, err := cm.db.NewDelete().Model((*ShortMetaData)(nil)).Exec(cm.ctx)
	if err != nil {
		return err
	}

	return nil
}

func (cm *CacheManager) AddImageData(hash string, contentId *int64, size int64) error {
	cm.dbMutex.Lock()
	defer cm.dbMutex.Unlock()

	reqTime := time.Now().Unix()
	expTime := time.Now().Add(cm.CacheLifetime).Unix()

	var err error

	data := ImageData{
		ContentID: contentId,
		Hash:      hash,
		CreatedAt: reqTime,
		ExpiresAt: expTime,
		Size:      size,
	}

	_, err = cm.db.NewInsert().Model(&data).Exec(cm.ctx)
	return err
}

func (cm *CacheManager) GetImageData(hash string) (*ImageData, error) {
	cm.dbMutex.Lock()
	defer cm.dbMutex.Unlock()

	var imgData ImageData

	err := cm.db.NewSelect().Model(&imgData).Where("hash = ?", hash).Scan(cm.ctx)

	if err == nil {
		return &imgData, nil
	} else {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		return nil, err
	}
}

func (cm *CacheManager) GetImagesData(hashes []string) ([]ImageData, error) {
	cm.dbMutex.Lock()
	defer cm.dbMutex.Unlock()

	var (
		imagesData []ImageData
		err        error
	)

	if len(hashes) > 0 {
		err = cm.db.NewSelect().Model(&imagesData).Where("hash IN (?)", bun.In(hashes)).Scan(cm.ctx)
	} else {
		err = cm.db.NewSelect().Model(&imagesData).Order("expiresAt ASC").Limit(100).Scan(cm.ctx)
	}

	if err != nil {
		cm.lg.Error(err.Error())
		return nil, err
	}

	return imagesData, nil
}

func (cm *CacheManager) RemoveImageData(hash string) error {
	cm.dbMutex.Lock()
	defer cm.dbMutex.Unlock()

	_, err := cm.db.NewDelete().Model((*ImageData)(nil)).Where("hash = ?", hash).Exec(cm.ctx)
	if err != nil {
		return err
	}

	return nil
}

func (cm *CacheManager) RemoveImagesData(hashes []string) error {
	cm.dbMutex.Lock()
	defer cm.dbMutex.Unlock()

	_, err := cm.db.NewDelete().Model((*ImageData)(nil)).Where("hash IN (?)", bun.In(hashes)).Exec(cm.ctx)
	if err != nil {
		return err
	}

	return nil
}

func (cm *CacheManager) PurgeImageData() error {
	cm.dbMutex.Lock()
	defer cm.dbMutex.Unlock()

	_, err := cm.db.NewDelete().Model((*ImageData)(nil)).Exec(cm.ctx)
	if err != nil {
		return err
	}

	return nil
}

func (cm *CacheManager) GetCacheSize() (int64, error) {
	cm.dbMutex.Lock()
	defer cm.dbMutex.Unlock()

	res := cm.db.QueryRow(CACHE_SIZE_REQUEST)
	var size int64

	err := res.Scan(&size)

	if err != nil {
		if strings.Contains(err.Error(), "converting NULL to int64") {
			return 0, nil
		} else {
			return 0, err
		}
	}

	return size, nil
}

func (cm *CacheManager) addImageCache(fileHash string, contentId string, size int64) error {
	cacheSize, err := cm.GetCacheSize()
	if err != nil {
		return err
	}

	cd, err := (strconv.Atoi(contentId))

	contId := int64(cd)

	if err == nil {
		cm.AddImageData(fileHash, &contId, size)
	} else {
		cm.AddImageData(fileHash, nil, size)
	}

	if cacheSize > cm.imageCacheMaxSize {
		data, err := cm.GetImagesData([]string{})
		if err != nil {
			return err
		}

		datum, hashes := elemsToDelete(data, cacheSize-cm.imageCacheMaxSize)

		cm.RemoveImagesData(hashes)

		for i, img := range datum {
			if img.ContentID != nil {
				os.Remove(path.Join(cm.imgDir, fmt.Sprintf("%v", *(datum[i].ContentID)), img.Hash))
			}
		}
	}

	return nil
}

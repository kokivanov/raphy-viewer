package cache

const (
	CACHE_SIZE_REQUEST = ` 
		SELECT SUM(size) AS s FROM (
			SELECT size from meta_data
			UNION
			SELECT size FROM image_data
		)
	`
)

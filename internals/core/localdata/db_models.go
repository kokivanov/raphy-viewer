package localdata

import "github.com/uptrace/bun"

type MediaStatus string

const (
	Online      MediaStatus = "Online"
	Downloading MediaStatus = "Downloading"
	Offline     MediaStatus = "Offline"
)

type MediaData struct {
	bun.BaseModel `bun:"table:media_data,alias:md"`

	ID         int64                   `bun:"id,pk"`
	IsFavorite bool                    `bun:"isFavorite,type:integer,notnull,default:0"`
	FavAt      *int64                  `bun:"favAt"`
	Status     MediaStatus             `bun:"status,notnull,type:varchar(12),default:'Online'"`
	LoadedAt   *int64                  `bun:"loadedAt"`
	Size       *int64                  `bun:"size"`
	MetaData   *map[string]interface{} `bun:"metadata,type:jsonb"`
}

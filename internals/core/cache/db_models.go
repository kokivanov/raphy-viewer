package cache

import (
	"github.com/uptrace/bun"
)

type MetaData struct {
	bun.BaseModel `bun:"table:meta_data,alias:md"`

	ID        int64  `bun:"id,pk"`
	CreatedAt int64  `bun:"createdAt,notnull"`
	ExpiresAt int64  `bun:"expiresAt,notnull"`
	Size      int    `bun:"size,notnull"`
	Data      string `bun:"data,notnull"`
}

type ImageData struct {
	bun.BaseModel `bun:"table:image_data,alias:img_d"`

	Hash      string `bun:"hash,pk"`
	CreatedAt int64  `bun:"createdAt,notnull"`
	ExpiresAt int64  `bun:"expiresAt,notnull"`
	Size      int64  `bun:"size,notnull"`
	ContentID *int64 `bun:"contentId"`

	MetaData *MetaData `bun:"rel:belongs-to,join:contentId=id"`
}

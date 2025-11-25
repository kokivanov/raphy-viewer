package cache

type ShortMediaIDs struct {
	IDs []int64 `json:"ids"`
}

type CachedShortMetaData struct {
	ID   int64
	Data string
}

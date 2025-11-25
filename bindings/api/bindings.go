package api

import (
	"context"
	"encoding/json"
	"fmt"
	"raphyviewer/internals/api/core"
	cache "raphyviewer/internals/core/cache"
	"time"
)

func (a *ApiService) QueryAllContent(options *core.QueryAllOptions) []core.ShortMedia {
	cached, err := a.c.GetRequestCache(cache.GetObjectHash(*options))
	if cached != nil {
		fmt.Println("Cache hit!")
		if cached.ExpiresAt > time.Now().Unix() {
			fmt.Printf("%v = %v", cached.ExpiresAt, time.Now().Unix())
			var reqIDs cache.ShortMediaIDs
			err := json.Unmarshal([]byte(cached.Data), &reqIDs)
			if err == nil {
				fmt.Printf("Requesting ids %v\n", reqIDs)
				datum, err := a.c.GetShortMetaDatum(reqIDs.IDs)
				if err == nil {
					dataToReturn := make([]core.ShortMedia, len(*datum))
					var err error

					for i, d := range *datum {
						var unmarshalled core.ShortMedia
						err = json.Unmarshal([]byte(d.Data), &unmarshalled)
						if err == nil {
							dataToReturn[i] = unmarshalled
						} else {
							break
						}
					}

					if err == nil {
						return dataToReturn
					}
				}
			}
		} else {
			err := a.c.RemoveRequestCache(cache.GetObjectHash(*options))
			if err != nil {
				a.l.Error(err.Error())
			}
		}
	}

	fmt.Println("Requesting new")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)

	resp, err := a.p.QueryAll(&ctx, options)

	if err != nil {
		a.l.Error(err.Error())
	}

	cancel()

	if v, ok := resp.([]core.ShortMedia); ok {
		ids := make([]int64, len(v))
		for i, m := range v {
			ids[i] = m.IDMal
		}

		reqIDs := cache.ShortMediaIDs{
			IDs: ids,
		}
		idsObj, err := json.Marshal(reqIDs)
		if err == nil {
			a.c.AddRequestCache(cache.GetObjectHash(options), string(idsObj))

			d := make([]cache.CachedShortMetaData, len(v))
			var err error
			for i, m := range v {
				var marshalled []byte
				marshalled, err = json.Marshal(m)
				if err == nil {
					d[i] = cache.CachedShortMetaData{
						ID:   m.IDMal,
						Data: string(marshalled),
					}
				} else {
					break
				}
			}

			if err == nil {
				a.c.AddShortMetaDatum(d)
			}
		}

		return v
	} else {
		a.l.Error("got wrong return type during QuerryAll")
		return nil
	}
}

func (a *ApiService) QueryByID(id int64) *core.FullMedia {
	data, err := a.c.GetMetaData(id)

	var v core.FullMedia

	if data != nil {
		err := json.Unmarshal([]byte(data.Data), &v)
		a.l.Debug(fmt.Sprintf("Cache hit for %v", id))
		if err == nil {
			return &v
		}
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)

	resp, err := a.p.QueryById(&ctx, id)

	if err != nil {
		a.l.Error(err.Error())
	}

	cancel()

	if v, ok := resp.(core.FullMedia); ok {
		d, err := json.Marshal(v)
		if err == nil {
			a.c.AddMetaData(id, string(d[:]))
		} else {
			a.l.Error(fmt.Sprintf("Error adding cache for %v", id))
		}

		return &v
	} else {
		a.l.Error("got wrong return type during QuerryAll")
		return nil
	}
}

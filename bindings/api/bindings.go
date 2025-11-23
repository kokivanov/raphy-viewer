package api

import (
	"context"
	"encoding/json"
	"fmt"
	"raphyviewer/internals/api/core"
	"time"
)

func (m *ApiService) QueryAllContent(options *core.QueryAllOptions) []core.ShortMedia {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)

	resp, err := m.p.QueryAll(&ctx, options)

	if err != nil {
		m.l.Error(err.Error())
	}

	cancel()

	if v, ok := resp.([]core.ShortMedia); ok {
		return v
	} else {
		m.l.Error("got wrong return type during QuerryAll")
		return nil
	}
}

func (m *ApiService) QueryByID(id int64) *core.FullMedia {
	data, err := m.c.GetMetaData(id)

	var v core.FullMedia

	if data != nil {
		err := json.Unmarshal([]byte(data.Data), &v)
		m.l.Debug(fmt.Sprintf("Cache hit for %v", id))
		if err == nil {
			return &v
		}
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)

	resp, err := m.p.QueryById(&ctx, id)

	if err != nil {
		m.l.Error(err.Error())
	}

	cancel()

	if v, ok := resp.(core.FullMedia); ok {
		d, err := json.Marshal(v)
		if err == nil {
			m.c.AddMetaData(id, string(d[:]))
		} else {
			m.l.Error(fmt.Sprintf("Error adding cache for %v", id))
		}

		return &v
	} else {
		m.l.Error("got wrong return type during QuerryAll")
		return nil
	}
}

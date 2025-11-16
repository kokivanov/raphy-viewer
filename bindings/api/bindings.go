package api

import (
	"context"
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
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)

	resp, err := m.p.QueryById(&ctx, id)

	if err != nil {
		m.l.Error(err.Error())
	}

	cancel()

	if v, ok := resp.(core.FullMedia); ok {
		return &v
	} else {
		m.l.Error("got wrong return type during QuerryAll")
		return nil
	}
}

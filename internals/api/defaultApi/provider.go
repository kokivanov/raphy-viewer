package defaultApi

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"raphyviewer/internals/api/core"
	"time"
)

type DefaultAPIProvider struct {
	c *http.Client

	uri string
}

func NewDefaultProvider() core.ApiProvider {
	return DefaultAPIProvider{
		c: &http.Client{
			Timeout: time.Second * 3,
		},
		uri: "https://graphql.anilist.co",
	}
}

func (p DefaultAPIProvider) QueryAll(ctx *context.Context, options *core.QueryAllOptions) (any, error) {
	r := core.GQLRequest{
		Query:     QUERY_ALL_REQUEST,
		Variables: options,
	}

	b, err := json.Marshal(r)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(*ctx, http.MethodPost, p.uri, bytes.NewReader(b))
	if err != nil {
		return nil, fmt.Errorf("new request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	resp, err := p.c.Do(req)
	if err != nil {
		return nil, fmt.Errorf("do request: %w", err)
	}
	defer resp.Body.Close()

	var respObj core.ShortResponse
	if err := json.NewDecoder(resp.Body).Decode(&respObj); err != nil {
		return nil, fmt.Errorf("decode response: %w", err)
	}

	if len(respObj.Errors) > 0 {
		return nil, fmt.Errorf("graphql error: %s", respObj.Errors[0].Message)
	}

	return respObj.Data.Page.Media, nil
}

func (p DefaultAPIProvider) QueryById(ctx *context.Context, id int64) (any, error) {
	r := core.GQLRequest{
		Query: QUERY_BY_ID_REQUEST,
		Variables: map[string]string{
			"mediaId": fmt.Sprint(id),
		},
	}

	b, err := json.Marshal(r)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(*ctx, http.MethodPost, p.uri, bytes.NewReader(b))
	if err != nil {
		return nil, fmt.Errorf("new request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	resp, err := p.c.Do(req)
	if err != nil {
		return nil, fmt.Errorf("do request: %w", err)
	}
	defer resp.Body.Close()

	var respObj core.FullDataResponse
	if err := json.NewDecoder(resp.Body).Decode(&respObj); err != nil {
		return nil, fmt.Errorf("decode response: %w", err)
	}

	if len(respObj.Errors) > 0 {
		return nil, fmt.Errorf("graphql error: %s", respObj.Errors[0].Message)
	}

	return respObj.Data.Media, nil
}

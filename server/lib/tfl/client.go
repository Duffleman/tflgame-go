package tfl

import (
	"context"
	"fmt"
	"net/url"
	"strings"

	"tflgame/server/lib/cher"
	"tflgame/server/lib/jsonclient"
)

type Client struct {
	baseURL string
	apiKey  string
	*jsonclient.Client
}

func NewClient(baseURL string, apiKey string) *Client {
	jc := jsonclient.NewClient(baseURL, nil)

	return &Client{baseURL, apiKey, jc}
}

func (c *Client) DoWithKey(ctx context.Context, method, path string, params url.Values, src, dst interface{}) error {
	if c.apiKey != "" {
		if params == nil {
			params = url.Values{}
		}

		params.Add("app_key", c.apiKey)
	}

	return c.Do(ctx, method, path, params, src, dst)
}

func (c *Client) ListLines(ctx context.Context, modes ...string) (lines []*Line, err error) {
	if len(modes) < 1 {
		return nil, cher.New("must_provide_modes", nil)
	}

	modesStr := strings.Join(modes, ",")

	path := fmt.Sprintf("/Line/Mode/%s", modesStr)

	return lines, c.DoWithKey(ctx, "GET", path, nil, nil, &lines)
}

func (c *Client) ListStops(ctx context.Context, lineID string) (stops []*Stop, err error) {
	if lineID == "" {
		return nil, cher.New("empty_line", nil)
	}

	path := fmt.Sprintf("/Line/%s/StopPoints", lineID)

	return stops, c.DoWithKey(ctx, "GET", path, nil, nil, &stops)
}

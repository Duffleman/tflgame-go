package tfl

import (
	"context"
	"fmt"
	"strings"

	"tflgame/server/lib/cher"
	"tflgame/server/lib/jsonclient"
)

type Client struct {
	baseURL string
	*jsonclient.Client
}

func NewClient(baseURL string) *Client {
	jc := jsonclient.NewClient(baseURL, nil)

	return &Client{baseURL, jc}
}

func (c *Client) ListLines(ctx context.Context, modes ...string) (lines []*Line, err error) {
	if len(modes) < 1 {
		return nil, cher.New("must_provide_modes", nil)
	}

	modesStr := strings.Join(modes, ",")

	path := fmt.Sprintf("/Line/Mode/%s", modesStr)

	return lines, c.Do(ctx, "GET", path, nil, nil, &lines)
}

func (c *Client) ListStops(ctx context.Context, lineID string) (stops []*Stop, err error) {
	if lineID == "" {
		return nil, cher.New("empty_line", nil)
	}

	path := fmt.Sprintf("/Line/%s/StopPoints", lineID)

	return stops, c.Do(ctx, "GET", path, nil, nil, &stops)
}

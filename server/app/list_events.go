package app

import (
	"context"

	"tflgame"
)

func (a *App) ListEvents(ctx context.Context, req *tflgame.ListEventsRequest) ([]*tflgame.Event, error) {
	return a.db.Q.ListEvents(ctx, req)
}

package app

import (
	"context"

	"tflgame"
	"tflgame/server/lib/cher"
)

func (a *App) ChangeHandle(ctx context.Context, req *tflgame.ChangeHandleRequest) (*tflgame.PublicUser, error) {
	user, err := a.GetUserByID(ctx, req.ID)
	if err != nil {
		return nil, err
	}

	if user.Handle == req.NewHandle {
		return nil, cher.New("already_set", nil)
	}

	numeric, err := a.db.ChangeHandle(ctx, req.ID, req.NewHandle)
	if err != nil {
		return nil, err
	}

	return &tflgame.PublicUser{
		ID:      req.ID,
		Handle:  req.NewHandle,
		Numeric: numeric,
	}, nil
}

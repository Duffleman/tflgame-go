package rpc

import (
	"context"
)

func (r *RPC) SyncTFLData(ctx context.Context) error {
	return r.app.SyncTFLData(ctx)
}

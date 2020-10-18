package app

import (
	"context"

	"tflgame/server/lib/tfl"

	"golang.org/x/sync/errgroup"
)

func (a *App) SyncTFLData(ctx context.Context) error {
	err := a.db.TFLInsertModes(ctx, tfl.Tube, tfl.DLR, tfl.Overground, tfl.NationalRail, tfl.TFLRail)
	if err != nil {
		return nil
	}

	lines, err := a.tfl.ListLines(ctx, tfl.Tube, tfl.DLR, tfl.Overground, tfl.NationalRail, tfl.TFLRail)
	if err != nil {
		return err
	}

	err = a.db.TFLInsertLines(ctx, lines)
	if err != nil {
		return err
	}

	g, gctx := errgroup.WithContext(ctx)

	for _, line := range lines {
		line := line
		g.Go(func() error {
			a.Logger.Infof("line_update: %s", line.Name)

			stops, err := a.tfl.ListStops(gctx, line.ID)
			if err != nil {
				return err
			}

			a.Logger.Infof("line: %s, station_count: %d", line.Name, len(stops))

			err = a.db.TFLInsertStops(gctx, stops)
			if err != nil {
				return err
			}

			err = a.db.TFLInsertLineStops(gctx, stops)
			if err != nil {
				return err
			}

			return nil
		})
	}

	return g.Wait()
}

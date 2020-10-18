package db

import (
	"context"

	"tflgame/server/lib/tfl"
)

func (d *DB) TFLInsertLineStops(ctx context.Context, stops []*tfl.Stop) error {
	return d.DoTx(ctx, func(qw *QueryableWrapper) error {
		for _, stop := range stops {
			for _, lmg := range stop.LineModeGroups {
				if lmg.ModeName == tfl.Bus {
					continue
				}

				for _, line := range lmg.LineIdentifier {
					_, err := qw.q.ExecContext(ctx, `
						INSERT INTO tfl_lines_stops
						(line_id, stop_id, mode)
						VALUES($1, $2, $3) ON CONFLICT (line_id, stop_id) DO NOTHING
					`, line, stop.ID, lmg.ModeName)
					if err != nil {
						return err
					}
				}
			}
		}

		return nil
	})
}

package db

import (
	"context"
	"strings"

	"tflgame/server/lib/tfl"
)

func (d *DB) TFLInsertStopsZones(ctx context.Context, stops []*tfl.Stop) error {
	return d.DoTx(ctx, func(qw *QueryableWrapper) error {
		for _, stop := range stops {
			var zp *tfl.AdditionalProperty

			for _, ap := range stop.AdditionalProperties {
				if ap.Key == "Zone" {
					zp = &ap
					break
				}
			}

			if zp == nil {
				return nil
			}

			zones := strings.Split(zp.Value, "+")

			for _, zone := range zones {
				_, err := qw.q.ExecContext(ctx, `
					INSERT INTO tfl_stops_zones
					(stop_id, zone)
					VALUES($1, $2) ON CONFLICT (stop_id, zone) DO NOTHING
				`, stop.ID, zone)
				if err != nil {
					return err
				}
			}
		}

		return nil
	})
}

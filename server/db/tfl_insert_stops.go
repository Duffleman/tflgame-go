package db

import (
	"context"
	"strings"

	"tflgame/server/lib/tfl"
)

func (d *DB) TFLInsertStops(ctx context.Context, stops []*tfl.Stop) error {
	return d.DoTx(ctx, func(qw *QueryableWrapper) error {
		for _, stop := range stops {
			shortName := stop.Name

			if strings.HasSuffix(shortName, " Underground Station") {
				shortName = strings.TrimSuffix(shortName, " Underground Station")
			}

			if strings.HasSuffix(shortName, " Rail Station") {
				shortName = strings.TrimSuffix(shortName, " Rail Station")
			}

			if strings.HasSuffix(shortName, " DLR Station") {
				shortName = strings.TrimSuffix(shortName, " DLR Station")
			}

			_, err := qw.q.ExecContext(ctx, `
			INSERT INTO tfl_stops
			(id, name, short_name, ics_code, station_naptan, status, lat, lon)
			VALUES($1, $2, $3, $4, $5, $6, $7, $8) ON CONFLICT (id) DO UPDATE
			SET name=$2, short_name=$3, ics_code=$4, station_naptan=$5, status=$6, lat=$7, lon=$8
		`, stop.ID, stop.Name, shortName, stop.ICSCode, stop.StationNaptan, stop.Status, stop.Lat, stop.Lon)
			if err != nil {
				return err
			}
		}

		return nil
	})
}

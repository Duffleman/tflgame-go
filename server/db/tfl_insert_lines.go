package db

import (
	"context"

	"tflgame/server/lib/tfl"
)

func (d *DB) TFLInsertLines(ctx context.Context, lines []*tfl.Line) error {
	return d.DoTx(ctx, func(qw *QueryableWrapper) error {
		for _, line := range lines {
			_, err := qw.q.ExecContext(ctx, `
			INSERT INTO tfl_lines
			(id, name, mode_name, created_at, modified_at)
			VALUES($1, $2, $3, $4, $5) ON CONFLICT (id) DO UPDATE
			SET name=$2, mode_name=$3, created_at=$4, modified_at=$5
		`, line.ID, line.Name, line.ModeName, line.CreatedAt, line.ModifiedAt)
			if err != nil {
				return err
			}
		}

		return nil
	})
}

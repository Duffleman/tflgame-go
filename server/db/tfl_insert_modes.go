package db

import (
	"context"
	"strings"
)

func (d *DB) TFLInsertModes(ctx context.Context, modesStr ...string) error {
	modes := map[string]string{}

	for _, mode := range modesStr {
		title := mode

		switch mode {
		case "dlr":
			title = "DLR"
		case "tflrail":
			title = "TFL Rail"
		case "national-rail":
			title = "National Rail"
		default:
			title = strings.Title(mode)
		}

		modes[mode] = title
	}

	return d.DoTx(ctx, func(qw *QueryableWrapper) error {
		_, err := qw.q.ExecContext(ctx, `
		TRUNCATE tfl_modes
		`)
		if err != nil {
			return err
		}

		for mode, name := range modes {
			_, err = qw.q.ExecContext(ctx, `
			INSERT INTO tfl_modes
			(id, name)
			VALUES($1, $2)
			`, mode, name)
			if err != nil {
				return err
			}
		}

		return nil
	})
}

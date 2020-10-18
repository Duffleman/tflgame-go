package db

import (
	"context"

	"tflgame"

	sq "github.com/Masterminds/squirrel"
)

func (qw *QueryableWrapper) GetPossiblePrompts(ctx context.Context, options *tflgame.GameOptions) ([]string, error) {
	query, values, err := NewQueryBuilder().
		Select("DISTINCT(s.short_name)").
		From("tfl_stops s").
		Join("tfl_stops_zones zs on s.id = zs.stop_id").
		Join("tfl_lines_stops ls ON s.id = ls.stop_id").
		Join("tfl_lines l ON ls.line_id = l.id").
		Where(sq.Eq{
			"l.id":    options.Lines,
			"ls.mode": options.Modes,
			"zs.zone": options.Zones,
		}).
		OrderBy("s.short_name ASC").
		ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := qw.q.QueryContext(ctx, query, values...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	prompts := []string{}

	for rows.Next() {
		var p string

		err := rows.Scan(&p)
		if err != nil {
			return nil, err
		}

		prompts = append(prompts, p)
	}

	return prompts, nil
}

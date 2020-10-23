package db

import (
	"context"
	"fmt"
	"strings"

	"tflgame"

	sq "github.com/Masterminds/squirrel"
)

func (qw *QueryableWrapper) GetPossiblePrompts(ctx context.Context, options *tflgame.GameOptions, rounds int) ([]string, error) {
	qb := NewQueryBuilder().
		Select("DISTINCT(s.short_name), random()").
		From("tfl_stops s").
		LeftJoin("tfl_stops_zones zs on s.id = zs.stop_id").
		Join("tfl_lines_stops ls ON s.id = ls.stop_id").
		Join("tfl_lines l ON ls.line_id = l.id").
		Where(sq.Eq{
			"l.id": options.Lines,
		})

	if options.Zones != nil {
		qb = qb.Where(sq.Eq{
			"zs.zone": options.Zones,
		})
	}

	qb = qb.OrderBy("random()")

	if rounds > 0 {
		qb = qb.Limit(uint64(rounds))
	}

	query, values, err := qb.ToSql()
	if err != nil {
		return nil, err
	}

	fmt.Println(query)

	rows, err := qw.q.QueryContext(ctx, query, values...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	prompts := []string{}

	for rows.Next() {
		var p string
		var r interface{}

		err := rows.Scan(&p, &r)
		if err != nil {
			return nil, err
		}

		p = strings.TrimSpace(p)

		prompts = append(prompts, p)
	}

	return prompts, nil
}

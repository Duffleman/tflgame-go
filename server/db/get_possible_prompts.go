package db

import (
	"context"
	"strings"

	"tflgame"
	"tflgame/server/lib/cher"

	sq "github.com/Masterminds/squirrel"
)

func (qw *QueryableWrapper) GetPossiblePrompts(ctx context.Context, options *tflgame.GameOptions, rounds int) ([]string, error) {
	subQuery := NewQueryBuilder().
		Select("DISTINCT(s.short_name)").
		From("tfl_stops s").
		LeftJoin("tfl_stops_zones zs on s.id = zs.stop_id").
		Join("tfl_lines_stops ls ON s.id = ls.stop_id").
		Join("tfl_lines l ON ls.line_id = l.id").
		Where(sq.Eq{
			"l.id": options.Lines,
		})

	if options.Zones != nil {
		subQuery = subQuery.Where(sq.Eq{
			"zs.zone": options.Zones,
		})
	}

	qb := NewQueryBuilder().
		Select("t.short_name").
		FromSelect(subQuery, "t").
		OrderBy("random()")

	if rounds > 0 {
		qb = qb.Limit(uint64(rounds))
	}

	query, values, err := qb.ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := qw.q.QueryContext(ctx, query, values...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	promptMap := map[string]struct{}{}
	prompts := []string{}

	for rows.Next() {
		var p string

		err := rows.Scan(&p)
		if err != nil {
			return nil, err
		}

		p = strings.TrimSpace(p)

		promptMap[p] = struct{}{}
		prompts = append(prompts, p)
	}

	if rounds != 0 && len(promptMap) != rounds {
		return nil, cher.New("duplicates_generated", nil)
	}

	return prompts, nil
}

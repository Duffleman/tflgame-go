package db

import (
	"context"
	"fmt"
)

func (qw *QueryableWrapper) GetAllLines(ctx context.Context, stationName string) ([]string, error) {
	/*
		SELECT DISTINCT(ls.line_id)
		FROM tfl_stops s
		JOIN tfl_lines_stops ls ON s.id = ls.stop_id
		WHERE LOWER(short_name) LIKE LOWER('%finsbury park%')
	*/
	query, values, err := NewQueryBuilder().
		Select("DISTINCT(ls.line_id)").
		From("tfl_stops s").
		Join("tfl_lines_stops ls ON s.id = ls.stop_id").
		Where("LOWER(short_name) LIKE LOWER(?)", fmt.Sprint("%", stationName, "%")).
		ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := qw.q.QueryContext(ctx, query, values...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	lines := []string{}

	for rows.Next() {
		var line string

		err := rows.Scan(&line)
		if err != nil {
			return nil, err
		}

		lines = append(lines, line)
	}

	return lines, err
}

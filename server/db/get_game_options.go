package db

import (
	"context"

	"tflgame"
)

func (qw *QueryableWrapper) GetGameOptions(ctx context.Context) (*tflgame.GetGameOptionsResponse, error) {
	lines, err := qw.getLines(ctx)
	if err != nil {
		return nil, err
	}

	zones, err := qw.getZones(ctx)
	if err != nil {
		return nil, err
	}

	set := &tflgame.GetGameOptionsResponse{
		Lines: lines,
		Zones: zones,
	}

	return set, nil
}

func (qw *QueryableWrapper) getLines(ctx context.Context) (map[string][]tflgame.LineDisplay, error) {
	query, values, err := NewQueryBuilder().
		Select("id", "name", "mode_name").
		From("tfl_lines l").
		OrderBy("mode_name ASC").
		ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := qw.q.QueryContext(ctx, query, values...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	lines := map[string][]tflgame.LineDisplay{}

	for rows.Next() {
		var l struct {
			ID       string
			Name     string
			ModeName string
		}

		err := rows.Scan(&l.ID, &l.Name, &l.ModeName)
		if err != nil {
			return nil, err
		}

		if _, ok := lines[l.ModeName]; !ok {
			lines[l.ModeName] = []tflgame.LineDisplay{}
		}

		lines[l.ModeName] = append(lines[l.ModeName], tflgame.LineDisplay{
			ID:    l.ID,
			Name:  l.Name,
			Color: nil,
		})
	}

	return lines, nil
}

func (qw *QueryableWrapper) getZones(ctx context.Context) ([]string, error) {
	query, values, err := NewQueryBuilder().
		Select("DISTINCT(zone)").
		From("tfl_stops_zones").
		OrderBy("zone ASC").
		ToSql()
	if err != nil {
		return nil, err
	}
	if err != nil {
		return nil, err
	}

	rows, err := qw.q.QueryContext(ctx, query, values...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	zones := []string{}

	for rows.Next() {
		var zone string

		err := rows.Scan(&zone)
		if err != nil {
			return nil, err
		}

		zones = append(zones, zone)
	}

	return zones, nil
}

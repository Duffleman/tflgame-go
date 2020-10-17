package db

import (
	"fmt"

	"tflgame"
)

var mappedOrder = map[string]string{
	"oldest_first": "ASC",
	"newest_first": "DESC",
}

type Pagination struct {
	before, after *string
	order         string
	limit         int
}

var defaultPagination = Pagination{
	order: "ASC",
	limit: 100,
}

func NewPagination(p *tflgame.Pagination) (Pagination, error) {
	if p == nil {
		return defaultPagination, nil
	}

	order, ok := mappedOrder[p.Order]
	if !ok {
		return Pagination{}, fmt.Errorf("invalid pagination order: %s", p.Order)
	}

	return Pagination{
		before: p.Before,
		after:  p.After,
		order:  order,
		limit:  p.Limit,
	}, nil
}

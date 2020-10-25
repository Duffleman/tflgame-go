package db

import (
	"context"

	"tflgame"

	sq "github.com/Masterminds/squirrel"
)

func (qw *QueryableWrapper) ListPrompts(ctx context.Context, gameID string) ([]*tflgame.Prompt, error) {
	query, values, err := NewQueryBuilder().
		Select("p.id", "p.user_id", "p.game_id", "p.prompt", "p.answer", "p.answer_given", "p.correct", "p.created_at", "p.answered_at", "p.hint_given_at").
		From("proj_prompts p").
		Where(sq.Eq{
			"p.game_id": gameID,
		}).
		ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := qw.q.QueryContext(ctx, query, values...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	prompts := []*tflgame.Prompt{}

	for rows.Next() {
		var p tflgame.Prompt

		if err := rows.Scan(&p.ID, &p.UserID, &p.GameID, &p.Prompt, &p.Answer, &p.AnswerGiven, &p.Correct, &p.CreatedAt, &p.AnsweredAt, &p.HintGivenAt); err != nil {
			return nil, coerceNotFound(err)
		}

		prompts = append(prompts, &p)
	}

	return prompts, nil
}

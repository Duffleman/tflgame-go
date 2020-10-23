package db

import (
	"context"

	"tflgame"

	sq "github.com/Masterminds/squirrel"
)

func (qw *QueryableWrapper) GetPrompt(ctx context.Context, promptID string) (*tflgame.Prompt, error) {
	query, values, err := NewQueryBuilder().
		Select("id", "user_id", "game_id", "prompt", "answer", "answer_given", "correct", "created_at", "answered_at", "hint_given_at").
		From("proj_prompts p").
		Where(sq.Eq{
			"p.id": promptID,
		}).
		ToSql()
	if err != nil {
		return nil, err
	}

	row := qw.q.QueryRowContext(ctx, query, values...)

	var p tflgame.Prompt

	if err := row.Scan(&p.ID, &p.UserID, &p.GameID, &p.Prompt, &p.Answer, &p.AnswerGiven, &p.Correct, &p.CreatedAt, &p.AnsweredAt, &p.HintGivenAt); err != nil {
		return nil, coerceNotFound(err)
	}

	return &p, nil
}

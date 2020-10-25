package db

import (
	"context"

	"tflgame"
	"tflgame/server/lib/ptr"

	sq "github.com/Masterminds/squirrel"
)

func (qw *QueryableWrapper) GetNextPrompt(ctx context.Context, gameID string) (*tflgame.Prompt, error) {
	query, values, err := NewQueryBuilder().
		Select("p.id", "p.user_id", "p.game_id", "p.prompt", "p.answer", "p.answer_given", "p.correct", "p.created_at", "p.answered_at", "g.reveal_word_length").
		From("proj_prompts p").
		Join("proj_games g ON g.id = p.game_id").
		Where(sq.Eq{
			"p.game_id":     gameID,
			"p.answered_at": nil,
		}).
		Limit(1).
		ToSql()
	if err != nil {
		return nil, err
	}

	var revealWordLength bool

	row := qw.q.QueryRowContext(ctx, query, values...)

	var p tflgame.Prompt

	if err := row.Scan(&p.ID, &p.UserID, &p.GameID, &p.Prompt, &p.Answer, &p.AnswerGiven, &p.Correct, &p.CreatedAt, &p.AnsweredAt, &revealWordLength); err != nil {
		return nil, coerceNotFound(err)
	}

	if revealWordLength {
		p.Length = ptr.Int(len(p.Answer))
	}

	return &p, nil
}

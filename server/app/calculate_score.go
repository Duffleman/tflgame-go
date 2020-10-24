package app

import (
	"fmt"

	"tflgame"
	"tflgame/server/lib/ptr"
)

const (
	// Additions
	AdditionForCorrect = 2
	DeductionForHint   = -1

	// Multiplications
	IncludeRandomSpaces = 1.1
	RevealWordLength    = 0.8
	ChangeLetterOrder   = 2.5
)

func (a *App) CalculateScore(do tflgame.DifficultyOptions, prompts []*tflgame.Prompt) (int, *tflgame.Calculations, error) {
	c := &tflgame.Calculations{}

	var score float64 = 0

	c.Start = score
	c.Add(tflgame.CalculationEvent{
		Note:  "Start",
		Score: score,
	})

	for _, p := range prompts {
		if p.AnswerGiven == nil {
			c.Add(tflgame.CalculationEvent{
				PromptID: ptr.String(p.ID),
				Note:     "Prompt not answered",
			})

			continue
		}

		if p.Correct {
			score += AdditionForCorrect

			c.Add(tflgame.CalculationEvent{
				PromptID: ptr.String(p.ID),
				Effect:   ptr.String(fmt.Sprintf("%d", AdditionForCorrect)),
				Note:     "Answered & correct",
				Score:    score,
			})

			if p.HintGivenAt != nil {
				score += DeductionForHint

				c.Add(tflgame.CalculationEvent{
					PromptID: ptr.String(p.ID),
					Effect:   ptr.String(fmt.Sprintf("%d", DeductionForHint)),
					Note:     "Hint was given",
					Score:    score,
				})
			}
		} else {
			c.Add(tflgame.CalculationEvent{
				PromptID: ptr.String(p.ID),
				Note:     "Answered but incorrect",
				Score:    score,
			})
		}
	}

	c.Base = score

	c.Add(tflgame.CalculationEvent{
		Note:  "Base score calculated",
		Score: score,
	})

	if do.IncludeRandomSpaces {
		score *= IncludeRandomSpaces

		c.Add(tflgame.CalculationEvent{
			Note:   "do.include_random_spaces is ON",
			Effect: ptr.String(fmt.Sprintf("*%2f", IncludeRandomSpaces)),
			Score:  score,
		})
	}

	if do.RevealWordLength {
		score *= RevealWordLength

		c.Add(tflgame.CalculationEvent{
			Note:   "do.reveal_word_length is ON",
			Effect: ptr.String(fmt.Sprintf("*%2f", RevealWordLength)),
			Score:  score,
		})
	}

	if do.ChangeLetterOrder {
		score *= ChangeLetterOrder

		c.Add(tflgame.CalculationEvent{
			Note:   "do.change_letter_order is ON",
			Effect: ptr.String(fmt.Sprintf("*%2f", ChangeLetterOrder)),
			Score:  score,
		})
	}

	c.End = score

	c.Add(tflgame.CalculationEvent{
		Note:  "Score calculated",
		Score: score,
	})

	// this intentionally floors the score
	final := int(score)

	c.Final = final

	c.Add(tflgame.CalculationEvent{
		Note:  "Score floored",
		Score: float64(final),
	})

	return final, c, nil
}

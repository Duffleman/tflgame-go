package app

import (
	"fmt"

	"tflgame"
	"tflgame/server/lib/ptr"
)

const (
	// Additions
	start              = 0
	additionForCorrect = 2
	deductionForHint   = -1

	// Multiplications
	includeRandomSpaces = 1.1
	revealWordLength    = 0.8
	changeLetterOrder   = 2.5
)

func (a *App) CalculateGameScore(do tflgame.DifficultyOptions, prompts []*tflgame.Prompt) (int, *tflgame.Calculations, error) {
	c := &tflgame.Calculations{}

	var score float64 = start

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
			score += additionForCorrect

			c.Add(tflgame.CalculationEvent{
				PromptID: ptr.String(p.ID),
				Effect:   ptr.String(fmt.Sprintf("%d", additionForCorrect)),
				Note:     "Answered & correct",
				Score:    score,
			})

			if p.HintGivenAt != nil {
				score += deductionForHint

				c.Add(tflgame.CalculationEvent{
					PromptID: ptr.String(p.ID),
					Effect:   ptr.String(fmt.Sprintf("%d", deductionForHint)),
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
		score *= includeRandomSpaces

		c.Add(tflgame.CalculationEvent{
			Note:   "do.include_random_spaces is ON",
			Effect: ptr.String(fmt.Sprintf("*%2f", includeRandomSpaces)),
			Score:  score,
		})
	}

	if do.RevealWordLength {
		score *= revealWordLength

		c.Add(tflgame.CalculationEvent{
			Note:   "do.reveal_word_length is ON",
			Effect: ptr.String(fmt.Sprintf("*%2f", revealWordLength)),
			Score:  score,
		})
	}

	if do.ChangeLetterOrder {
		score *= changeLetterOrder

		c.Add(tflgame.CalculationEvent{
			Note:   "do.change_letter_order is ON",
			Effect: ptr.String(fmt.Sprintf("*%2f", changeLetterOrder)),
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

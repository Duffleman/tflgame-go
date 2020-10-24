package app

import (
	"strings"
)

var removeForComparison = []string{"'", "\"", "-"}

func (a *App) CompareAnswers(prompt, answer string) bool {
	answer = strings.ToUpper(answer)
	prompt = strings.ToUpper(prompt)

	answer = strings.Replace(answer, " ", "", -1)
	prompt = strings.Replace(prompt, " ", "", -1)

	for _, v := range removeForComparison {
		prompt = strings.Replace(prompt, v, "", -1)
	}

	for _, v := range removeForComparison {
		answer = strings.Replace(answer, v, "", -1)
	}

	return prompt == answer
}

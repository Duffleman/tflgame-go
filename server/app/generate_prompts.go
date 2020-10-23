package app

import (
	"math/rand"
	"strings"
	"time"

	"tflgame"
	"tflgame/server/lib/ptr"

	"github.com/cuvva/ksuid-go"
)

var alwaysRemove = []string{"A", "E", "I", "O", "U", "'", "\"", "-", "&"}

func (a *App) GeneratePrompts(promptsStr []string, do tflgame.DifficultyOptions) []tflgame.Prompt {
	prompts := []tflgame.Prompt{}

	promptsStr = promptsStr[:do.Rounds]

	for _, p := range promptsStr {
		id := ksuid.Generate("prompt").String()
		answer := strings.ToUpper(p)
		answerNoSpaces := strings.Replace(answer, " ", "", -1)
		prompt := answerNoSpaces

		for _, v := range alwaysRemove {
			prompt = strings.Replace(prompt, v, "", -1)
		}

		if do.ChangeLetterOrder {
			letters := strings.Split(prompt, "")
			letters = shuffleString(letters)
			prompt = strings.Join(letters, "")
		}

		if do.IncludeRandomSpaces {
			spacesToAdd := calculateSpacesToAdd(prompt)
			letters := strings.Split(prompt, "")
			indexesToAdd := generateRandomIndexes(spacesToAdd, prompt)

			for _, index := range indexesToAdd {
				letters = append(letters[:index+1], letters[index:]...)
				letters[index] = " "
			}

			prompt = strings.TrimSpace(strings.Join(letters, ""))
		}

		var length *int

		if do.RevealWordLength {
			length = ptr.Int(len(answer))
		}

		prompt = strings.TrimSpace(prompt)
		answer = strings.TrimSpace(answer)

		prompts = append(prompts, tflgame.Prompt{
			ID:     id,
			Prompt: prompt,
			Answer: answer,
			Length: length,
		})
	}

	return prompts
}

func shuffleString(vals []string) []string {
	r := rand.New(rand.NewSource(time.Now().UTC().UnixNano()))
	ret := make([]string, len(vals))
	n := len(vals)

	for i := 0; i < n; i++ {
		randIndex := r.Intn(len(vals))
		ret[i] = vals[randIndex]
		vals = append(vals[:randIndex], vals[randIndex+1:]...)
	}

	return ret
}

func shuffleInt(vals []int) []int {
	r := rand.New(rand.NewSource(time.Now().UTC().UnixNano()))
	ret := make([]int, len(vals))
	n := len(vals)

	for i := 0; i < n; i++ {
		randIndex := r.Intn(len(vals))
		ret[i] = vals[randIndex]
		vals = append(vals[:randIndex], vals[randIndex+1:]...)
	}

	return ret
}

func generateRandomIndexes(spaces int, prompt string) []int {
	possibleIndexes := []int{}

	for i := 1; i < len(prompt); i += 2 {
		possibleIndexes = append(possibleIndexes, i)
	}

	possibleIndexes = shuffleInt(possibleIndexes)

	return possibleIndexes[:spaces]
}

func calculateSpacesToAdd(prompt string) int {
	switch true {
	case len(prompt) < 5:
		return 1
	case len(prompt) < 8:
		return 2
	default:
		return 3
	}
}

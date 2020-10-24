package tflgame

import (
	"time"
)

type User struct {
	ID        string
	Handle    string
	Numeric   string
	Pin       *string
	Score     int
	CreatedAt time.Time
}

type PublicUser struct {
	UserID  string `json:"user_id"`
	Handle  string `json:"handle"`
	Numeric string `json:"numeric"`
}

type ListEventsRequest struct {
	UserID     string      `json:"user_id"`
	Pagination *Pagination `json:"pagination"`
}

type Pagination struct {
	Before *string `json:"before"`
	After  *string `json:"after"`
	Order  string  `json:"order"`
	Limit  int     `json:"limit"`
}

type CreateUserRequest struct {
	Handle string  `json:"handle"`
	Pin    *string `json:"pin"`
}

type CreateUserResponse struct {
	ID      string `json:"id"`
	Handle  string `json:"handle"`
	Numeric string `json:"numeric"`
	Token   string `json:"token"`
}

type AuthenticateRequest struct {
	Handle  string `json:"handle"`
	Numeric string `json:"numeric"`
	Pin     string `json:"pin"`
}

type AuthenticateResponse struct {
	UserID string `json:"user_id"`
	Token  string `json:"token"`
}

type ChangeHandleRequest struct {
	UserID    string `json:"user_id"`
	NewHandle string `json:"new_handle"`
}

type ReleaseHandleRequest struct {
	UserID string `json:"user_id"`
}

type ChangePinRequest struct {
	UserID     string `json:"user_id"`
	CurrentPin string `json:"current_pin"`
	NewPin     string `json:"new_pin"`
}

type GameOptions struct {
	Lines []string `json:"lines"`
	Zones []string `json:"zones"`
}

type TestGameOptionsResponse struct {
	PossiblePrompts int `json:"possible_prompts"`
}

type GetGameOptionsResponse struct {
	Lines map[string][]LineDisplay `json:"lines"`
	Zones []string                 `json:"zones"`
}

type LineDisplay struct {
	ID    string  `json:"id"`
	Name  string  `json:"name"`
	Color *string `json:"color"`
}

type CreateGameRequest struct {
	UserID            string            `json:"user_id"`
	DifficultyOptions DifficultyOptions `json:"difficulty_options"`
	GameOptions       GameOptions       `json:"game_options"`
}

type DifficultyOptions struct {
	Rounds              int  `json:"rounds"`
	IncludeRandomSpaces bool `json:"include_random_spaces"`
	ChangeLetterOrder   bool `json:"change_letter_order"`
	RevealWordLength    bool `json:"reveal_word_length"`
}

type CreateGameResponse struct {
	ID   string      `json:"id"`
	Next *NextPrompt `json:"next"`
}

// TODO(gm): this has some code smells of being used for too much
type Prompt struct {
	ID          string     `json:"id"`
	UserID      string     `json:"_"`
	GameID      string     `json:"_"`
	Prompt      string     `json:"prompt"`
	Answer      string     `json:"answer"`
	Length      *int       `json:"length"`
	AnswerGiven *string    `json:"_"`
	Correct     bool       `json:"_"`
	CreatedAt   time.Time  `json:"_"`
	AnsweredAt  *time.Time `json:"_"`
	HintGivenAt *time.Time `json:"_"`
}

type SubmitAnswerRequest struct {
	UserID   string `json:"user_id"`
	PromptID string `json:"prompt_id"`
	Answer   string `json:"answer"`
}

type SubmitAnswerResponse struct {
	UserID  string      `json:"user_id"`
	Correct bool        `json:"correct"`
	Answer  string      `json:"answer"`
	Next    *NextPrompt `json:"next"`
}

type NextPrompt struct {
	PromptID string `json:"prompt_id"`
	Prompt   string `json:"prompt"`
	Length   *int   `json:"length"`
}

type GetCurrentGameRequest struct {
	UserID string `json:"user_id"`
}

type GetCurrentGameResponse struct {
	GameID string      `json:"game_id"`
	Next   *NextPrompt `json:"next"`
}

type GetHintRequest struct {
	UserID   string `json:"user_id"`
	PromptID string `json:"prompt_id"`
}

type GetHintResponse struct {
	Prompt string   `json:"prompt"`
	Lines  []string `json:"lines"`
}

type GetGameStateRequest struct {
	UserID string `json:"user_id"`
	GameID string `json:"game_id"`
}

type GetGameStateResponse struct {
	InProgress        bool              `json:"in_progress"`
	Score             int               `json:"score"`
	GameTime          string            `json:"game_time"`
	DifficultyOptions DifficultyOptions `json:"difficulty_options"`
	GameOptions       GameOptions       `json:"game_options"`
}

type ExplainScoreRequest struct {
	UserID string  `json:"user_id"`
	GameID *string `json:"game_id"`
}

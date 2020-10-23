package tflgame

import (
	"encoding/json"
	"errors"
	"time"
)

type Event struct {
	ID        string      `json:"id"`
	Type      string      `json:"type"`
	UserID    string      `json:"user_id"`
	GameID    *string     `json:"game_id"`
	Payload   interface{} `json:"payload"`
	CreatedAt time.Time   `json:"created_at"`
}

type CreateUserPayload struct {
	CreationID string  `json:"creation_id"`
	Handle     string  `json:"handle"`
	Numeric    string  `json:"numeric"`
	Pin        *string `json:"pin,omitempty"`
}

type ChangeHandlePayload struct {
	UserID     string `json:"user_id"`
	NewHandle  string `json:"new_handle"`
	NewNumeric string `json:"new_numeric"`
}

type ChangePinPayload struct {
	UserID string `json:"user_id"`
	Pin    string `json:"pin,omitempty"`
}

type ReleaseHandlePayload struct {
	UserID  string `json:"user_id"`
	Handle  string `json:"handle"`
	Numeric string `json:"numeric"`
}

type CreateGamePayload struct {
	CreationID        string            `json:"creation_id"`
	Prompts           []Prompt          `json:"prompts"`
	DifficultyOptions DifficultyOptions `json:"difficulty_options"`
	GameOptions       GameOptions       `json:"game_options"`
}

type AnswerPromptPayload struct {
	PromptID    string `json:"prompt_id"`
	AnswerGiven string `json:"answer_given"`
	Correct     bool   `json:"correct"`
}

type GiveHintPayload struct {
	PromptID  string   `json:"prompt_id"`
	NewPrompt string   `json:"new_prompt"`
	Lines     []string `json:"lines"`
}

type FinishGamePayload struct{}

func PayloadHandler(eventType string, raw []byte, in *interface{}) error {
	switch eventType {
	case "create_user":
		var payload CreateUserPayload
		if err := json.Unmarshal(raw, &payload); err != nil {
			return err
		}

		*in = payload
	case "change_handle":
		var payload ChangeHandlePayload
		if err := json.Unmarshal(raw, &payload); err != nil {
			return err
		}

		*in = payload
	case "change_pin":
		var payload ChangePinPayload
		if err := json.Unmarshal(raw, &payload); err != nil {
			return err
		}

		*in = payload
	case "release_handle":
		var payload ReleaseHandlePayload
		if err := json.Unmarshal(raw, &payload); err != nil {
			return err
		}

		*in = payload
	case "create_game":
		var payload CreateGamePayload
		if err := json.Unmarshal(raw, &payload); err != nil {
			return err
		}

		*in = payload
	case "finish_game":
		*in = FinishGamePayload{}
	case "answer_prompt":
		var payload AnswerPromptPayload
		if err := json.Unmarshal(raw, &payload); err != nil {
			return err
		}

		*in = payload
	case "give_hint":
		var payload GiveHintPayload
		if err := json.Unmarshal(raw, &payload); err != nil {
			return err
		}

		*in = payload
	default:
		return errors.New("unknown event type")
	}

	return nil
}

func SafePublicPayload(in interface{}) (interface{}, error) {
	// TODO(gm): this seems immature compared to other areas of event maanagement
	// let's see if we can clean it up a little by using pointers?
	var out interface{}

	switch v := in.(type) {
	case CreateUserPayload:
		// remove "pin" hash
		v.Pin = nil
		out = v
	case ChangePinPayload:
		// remove "pin" hash
		v.Pin = ""
		out = v
	default:
		// no change
		out = v
	}

	return out, nil
}

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

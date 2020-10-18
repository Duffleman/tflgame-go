package app

import (
	"tflgame"
	"tflgame/server/lib/cher"

	"golang.org/x/crypto/bcrypt"
)

func (a *App) CheckPin(user *tflgame.User, pin *string) error {
	switch true {
	case user.Pin == nil && pin == nil:
		return nil
	case user.Pin != nil && pin == nil, user.Pin == nil && pin != nil:
		return cher.New(cher.Unauthorized, nil)
	}

	ok, err := comparePasswords(*user.Pin, *pin)
	if err != nil {
		return err
	}

	if !ok {
		return cher.New("incorrect_pin", nil)
	}

	return nil
}

func comparePasswords(hashedPwd, plainPwd string) (bool, error) {
	byteHash := []byte(hashedPwd)
	bytePlain := []byte(plainPwd)

	err := bcrypt.CompareHashAndPassword(byteHash, bytePlain)
	if err != nil {
		return false, cher.New(cher.Unauthorized, nil, cher.New("bad_pin", nil))
	}

	return true, nil
}

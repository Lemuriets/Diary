package auth

import (
	"errors"
)

var (
	EmptyLoginOrPassword = errors.New("login or password is underfined")
	WrongPassword        = errors.New("password is wrong")
)

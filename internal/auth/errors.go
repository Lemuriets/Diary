package auth

import (
	"errors"
)

var (
	UnderfinedLoginOrPassword = errors.New("login or password is underfined")
)

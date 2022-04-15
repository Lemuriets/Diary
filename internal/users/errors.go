package users

import (
	"errors"
)

var (
	SchoolIsNotExists = errors.New("School is not exists")
	UndefinedUser     = errors.New("Undefined user")
)

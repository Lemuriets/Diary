package app

import (
	"errors"
)

var (
	MethodsNotSpecified = errors.New("Methods for the handler were not specified")
)

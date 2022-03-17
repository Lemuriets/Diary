package http

import (
	"github.com/Lemuriets/diary/internal/schools"
)

type Handler struct {
	UseCase schools.UseCase
}

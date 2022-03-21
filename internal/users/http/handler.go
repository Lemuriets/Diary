package http

import (
	"github.com/Lemuriets/diary/internal/users"
)

type Handler struct {
	UseCase users.UseCase
}

func NewHandler(uc users.UseCase) *Handler {
	return &Handler{
		UseCase: uc,
	}
}

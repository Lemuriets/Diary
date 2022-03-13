package http

import (
	"github.com/Lemuriets/diary/internal/users"
)

type Handler struct {
	Repository users.UseCase
}

func NewHandler(uc users.UseCase) *Handler {
	return &Handler{
		Repository: uc,
	}
}

package http

import (
	"github.com/Lemuriets/diary/internal/auth"
)

type Handler struct {
	UseCase auth.UseCase
}

func NewHandler(uc auth.UseCase) *Handler {
	return &Handler{
		UseCase: uc,
	}
}

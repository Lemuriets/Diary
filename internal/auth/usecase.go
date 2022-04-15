package auth

import (
	"github.com/Lemuriets/diary/model"
)

type UseCase interface {
	GetCount(user model.User) int64
	Create(user model.User) error
	GenerateAccessToken(login, password string) (string, error)
	GenerateRefreshToken()
}

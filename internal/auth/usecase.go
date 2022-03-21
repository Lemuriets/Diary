package auth

import (
	"github.com/Lemuriets/diary/model"
)

type UseCase interface {
	GetCount(user model.User) int64
	Create(user model.User) error
	GenerateJWT(login, password string) (string, error)
}

package auth

import (
	"github.com/Lemuriets/diary/model"
)

type Repository interface {
	Get(login, password string) (model.User, error)
	GetCount(login, password string) int64
	Create(login, password string) error
}

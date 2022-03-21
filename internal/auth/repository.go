package auth

import (
	"github.com/Lemuriets/diary/model"
)

type Repository interface {
	Get(login, passwordHash string) (model.User, error)
	GetByLogin(login string) (model.User, error)
	GetAll() ([]model.User, error)
	GetById(id uint) (model.User, error)
	GetCountByLogin(login string) int64
	Create(user model.User) error
}

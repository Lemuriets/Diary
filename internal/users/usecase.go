package users

import "github.com/Lemuriets/diary/model"

type UseCase interface {
	GetById(id uint) model.User
}

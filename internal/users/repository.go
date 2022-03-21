package users

import "github.com/Lemuriets/diary/model"

type Repository interface {
	GetById(id uint) (model.User, error)
	Create()
	Update()
	Delete()
}

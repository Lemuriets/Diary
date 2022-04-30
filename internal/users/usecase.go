package users

import "github.com/Lemuriets/diary/model"

type UseCase interface {
	GetById(id uint) model.User
	Update(id uint, updateFields map[string]interface{}) error
	Delete(id uint) error
	MultipleDelete(ids []uint) error
	AddClass(user *model.User, class model.Class) error
}

package users

import "github.com/Lemuriets/diary/model"

type Repository interface {
	GetById(id uint) (model.User, error)
	Update(id uint, updateFields map[string]interface{}) error
	Delete(id uint) error
	MultipleDelete(ids []uint) error
}

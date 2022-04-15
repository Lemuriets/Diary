package users

import "github.com/Lemuriets/diary/model"

type Repository interface {
	GetById(id uint64) (model.User, error)
	Update(id uint64, updateFields map[string]interface{}) error
	Delete(id uint64) error
	MultipleDelete(ids []uint64) error
}

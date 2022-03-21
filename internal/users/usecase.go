package users

import "github.com/Lemuriets/diary/model"

type UseCase interface {
	GetById(id uint64) model.User
	Update(id uint64, updateFields map[string]interface{}) error
	Delete(id uint64) error
}

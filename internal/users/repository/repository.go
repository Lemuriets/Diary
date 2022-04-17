package repository

import (
	"github.com/Lemuriets/diary/model"
	"gorm.io/gorm"
)

type Repository struct {
	DB *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		DB: db,
	}
}

func (repo *Repository) GetById(id uint) (model.User, error) {
	var user model.User

	result := repo.DB.Where("ID = ?", id).First(&user)

	return user, result.Error
}

func (repo *Repository) Update(id uint, updateFields map[string]interface{}) error {
	result := repo.DB.Model(&model.User{}).Where("id = ?", id).Updates(updateFields)

	return result.Error
}

func (repo *Repository) Delete(id uint) error {
	result := repo.DB.Unscoped().Delete(&model.User{}, id)

	return result.Error
}

func (repo *Repository) MultipleDelete(ids []uint) error {
	result := repo.DB.Unscoped().Delete(&model.User{}, ids)

	return result.Error
}

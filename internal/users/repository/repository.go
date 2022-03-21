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

func (repo *Repository) Create() {

}

func (repo *Repository) Update() {

}

func (repo *Repository) Delete() {

}

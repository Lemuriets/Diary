package repository

import (
	"gorm.io/gorm"

	"github.com/Lemuriets/diary/model"
)

type Repository struct {
	DB *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		DB: db,
	}
}

func (repo *Repository) Get(login, password string) (model.User, error) {
	var user model.User

	result := repo.DB.Where("login = ?, password = ?", login, password).First(&user)

	return user, result.Error
}

func (repo *Repository) GetCount(login, password string) int64 {
	var count int64

	repo.DB.Where("login = ?, password = ?", login, password).Count(&count)

	return count
}

func (repo *Repository) Create(user model.User) error {
	res := repo.DB.Create(&user)

	return res.Error
}

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

func (repo *Repository) Get(login, passwordHash string) (model.User, error) {
	var user model.User

	result := repo.DB.Where("login = ? AND password_hash = ?", login, passwordHash).First(&user)
	return user, result.Error
}

func (repo *Repository) GetByLogin(login string) (model.User, error) {
	var user model.User

	result := repo.DB.Where("login = ?", login).First(&user)

	return user, result.Error
}

func (repo *Repository) GetById(id uint) (model.User, error) {
	var user model.User

	result := repo.DB.Where("ID = ?", id).First(&user)

	return user, result.Error
}

func (repo *Repository) GetAll() ([]model.User, error) {
	var users []model.User

	result := repo.DB.Find(&users)

	return users, result.Error
}

func (repo *Repository) GetCountByLogin(login string) int64 {
	var count int64

	repo.DB.Where("login = ?", login).Count(&count)

	return count
}

func (repo *Repository) Create(user model.User) error {
	res := repo.DB.Create(&user)

	return res.Error
}

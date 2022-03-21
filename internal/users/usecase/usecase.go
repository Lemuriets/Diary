package usecase

import (
	"github.com/Lemuriets/diary/internal/users"
	"github.com/Lemuriets/diary/model"
)

type UseCase struct {
	Repository users.Repository
}

func NewUseCase(repo users.Repository) *UseCase {
	return &UseCase{
		Repository: repo,
	}
}

func (uc *UseCase) GetById(id uint) model.User {
	user, err := uc.Repository.GetById(id)

	if err != nil {
		return model.User{}
	}

	return user
}

package auth

import (
	"github.com/Lemuriets/diary/model"
)

type UsecaseMock struct{}

func (uc *UsecaseMock) GetCount(user model.User) int64 {
	return 123
}

func (uc *UsecaseMock) Create(user model.User) error {
	return nil
}

func (uc *UsecaseMock) GenerateAccessToken(login, password string) (string, error) {
	return "access token", nil
}

func (uc *UsecaseMock) GenerateRefreshToken(login string) (string, error) {
	return "refresh token", nil
}

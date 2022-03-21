package usecase

import (
	"os"

	"github.com/Lemuriets/diary/internal/auth"
	"github.com/Lemuriets/diary/pkg/crypto"
	"github.com/golang-jwt/jwt"

	// "github.com/golang-jwt/jwt"
	"github.com/Lemuriets/diary/model"
)

type UseCase struct {
	Repository auth.Repository
}

type tokenClaims struct {
	jwt.StandardClaims
	UserId          uint  `json:"userId"`
	UserPermissions uint8 `json:"userPermissions"`
}

func NewUseCase(repo auth.Repository) *UseCase {
	return &UseCase{
		Repository: repo,
	}
}

func (uc *UseCase) GetCount(user model.User) int64 {
	return uc.Repository.GetCountByLogin(user.Login)
}

func (uc *UseCase) Create(user model.User) error {
	if user.Login == "" || user.PasswordHash == "" {
		return auth.UnderfinedLoginOrPassword
	}

	if err := uc.Repository.Create(user); err != nil {
		return err
	}

	return nil
}

func (uc *UseCase) GenerateJWT(login, password string) (string, error) {
	user, err := uc.Repository.GetByLogin(login)

	if !crypto.CompareHashAndPassword(password, user.PasswordHash) || err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, tokenClaims{
		UserId:          user.ID,
		UserPermissions: user.Permissions,
	})

	return token.SignedString([]byte(os.Getenv("secretkey")))
}

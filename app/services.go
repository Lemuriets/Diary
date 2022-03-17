package app

import (
	"gorm.io/gorm"

	authhttp "github.com/Lemuriets/diary/internal/auth/http"
	authrepo "github.com/Lemuriets/diary/internal/auth/repository"
	authuc "github.com/Lemuriets/diary/internal/auth/usecase"
)

func RegisterAuthService(db *gorm.DB) *authhttp.Handler {
	repo := authrepo.NewRepository(db)
	uc := authuc.NewUseCase(repo)

	return authhttp.NewHandler(uc)
}

func RegisterUsersService() {

}

func RegisterClassesService() {

}

func RegisterSchoolsService() {

}

func RegisterShedulesService() {

}

func RegisterHomeworkService() {

}

func RegisterMarksService() {

}

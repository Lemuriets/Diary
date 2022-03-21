package app

import (
	"gorm.io/gorm"

	authhttp "github.com/Lemuriets/diary/internal/auth/http"
	authrepo "github.com/Lemuriets/diary/internal/auth/repository"
	authuc "github.com/Lemuriets/diary/internal/auth/usecase"

	usershttp "github.com/Lemuriets/diary/internal/users/http"
	usersrepo "github.com/Lemuriets/diary/internal/users/repository"
	usersuc "github.com/Lemuriets/diary/internal/users/usecase"
)

func RegisterAuthService(db *gorm.DB) *authhttp.Handler {
	repo := authrepo.NewRepository(db)
	uc := authuc.NewUseCase(repo)

	return authhttp.NewHandler(uc)
}

func RegisterUsersService(db *gorm.DB) *usershttp.Handler {
	repo := usersrepo.NewRepository(db)
	uc := usersuc.NewUseCase(repo)

	return usershttp.NewHandler(uc)
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

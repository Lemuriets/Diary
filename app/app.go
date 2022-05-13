package app

import (
	"log"

	"github.com/gorilla/mux"

	authhttp "github.com/Lemuriets/diary/internal/auth/http"
	schoolhttp "github.com/Lemuriets/diary/internal/schools/http"
	usershttp "github.com/Lemuriets/diary/internal/users/http"

	"github.com/Lemuriets/diary/model"
	"github.com/Lemuriets/diary/pkg/crypto"
	"github.com/Lemuriets/diary/pkg/db"
)

type App struct {
	Router        *mux.Router
	AuthHandler   *authhttp.Handler
	UsersHandler  *usershttp.Handler
	SchoolHandler *schoolhttp.Handler
}

type Sub struct {
	App    *App
	Prefix string
}

func NewApp() *App {
	database := db.InitDB()

	err := database.AutoMigrate(
		&model.User{},
		&model.Grade{},
		&model.School{},
		&model.Homework{},
		&model.Lesson{},
		&model.Shedule{},
		&model.Mark{},
	)
	superUser := model.User{
		Login:        "Lemuriets",
		PasswordHash: crypto.HashPassword("secret"),
		Name:         "Egor",
		Lastname:     "Madyarov",
		Patronymic:   "Bogdanovich",
		Permissions:  model.SUPERUSER,
	}

	db.CreateSuperUser(database, superUser)

	if err != nil {
		log.Fatal(err)
	}

	return &App{
		Router:       mux.NewRouter(),
		AuthHandler:  RegisterAuthService(database),
		UsersHandler: RegisterUsersService(database),
	}
}

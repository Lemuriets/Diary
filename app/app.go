package app

import (
	"log"

	"github.com/gorilla/mux"

	authhttp "github.com/Lemuriets/diary/internal/auth/http"
	userhttp "github.com/Lemuriets/diary/internal/users/http"

	"github.com/Lemuriets/diary/model"
	"github.com/Lemuriets/diary/pkg/db"
)

type App struct {
	Router      *mux.Router
	AuthHandler *authhttp.Handler
	UserHandler *userhttp.Handler
}

func NewApp() *App {
	database := db.InitDB()

	err := database.AutoMigrate(
		&model.User{},
		&model.Class{},
		&model.School{},
		&model.Homework{},
		&model.Lesson{},
		&model.Shedule{},
	)

	if err != nil {
		log.Fatal(err)
	}

	return &App{
		Router: mux.NewRouter(),
	}
}

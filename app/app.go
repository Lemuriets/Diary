package app

import (
	"github.com/Lemuriets/diary/model"
	"github.com/Lemuriets/diary/pkg/db"
	"github.com/gorilla/mux"
)

type App struct {
	Router *mux.Router
}

func NewApp() *App {
	database := db.InitDB()

	database.AutoMigrate(&model.User{}, &model.Class{}, &model.School{}, &model.Homework{}, &model.Lesson{}, &model.Shedule{})

	return &App{}
}

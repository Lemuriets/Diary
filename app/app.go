package app

import (
	"log"

	"github.com/Lemuriets/diary/model"
	"github.com/Lemuriets/diary/pkg/db"
)

type App struct {
	AuthHandler string
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

	return &App{}
}

package app

import (
	"log"
	"net/http"
	"strings"

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
		&model.Class{},
		&model.School{},
		&model.Homework{},
		&model.Lesson{},
		&model.Shedule{},
		&model.Mark{},
		&model.RefreshSession{},
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

func (app *App) Group(prefix string, handler func(w http.ResponseWriter, r *http.Request), methods ...string) *Sub {
	app.Router.HandleFunc(prefix, handler).Methods(methods...)

	if strings.HasSuffix(prefix, "/") {
		prefix = prefix[:len(prefix)-1]
	}
	return &Sub{
		App:    app,
		Prefix: prefix,
	}
}

func (sub *Sub) RegisterSubHandler(prefix string, handler func(w http.ResponseWriter, r *http.Request), methods ...string) {
	if len(methods) == 0 {
		panic(MethodsNotSpecified)
	}

	sub.App.Router.HandleFunc(sub.Prefix+prefix, handler).Methods(methods...)
}

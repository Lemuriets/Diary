package app

import (
	"log"
	"net/http"
)

// "github.com/Lemuriets/diary/internal/auth/usecase"

func (app *App) Run() {
	app.RouteAuth()
	app.RouteUsers()
	app.RouteMarks()
	app.RouteSchools()
	app.RouteClasses()
	app.RouteHomework()
	app.RouteShedules()

	log.Fatal(http.ListenAndServe("127.0.0.1:8000", app.Router))
}

func (app *App) RouteAuth() {
	sub := app.Group("/auth", app.AuthHandler.SignUp, "GET")

	sub.RegisterSubHandler("/sign-up", app.AuthHandler.SignUp, "POST")
}

func (app *App) RouteUsers() {

}

func (app *App) RouteClasses() {

}

func (app *App) RouteSchools() {

}

func (app *App) RouteHomework() {

}

func (app *App) RouteShedules() {

}

func (app *App) RouteMarks() {

}

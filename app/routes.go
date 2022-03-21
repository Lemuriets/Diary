package app

import (
	"fmt"
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

	fmt.Println("server started")
	log.Fatal(http.ListenAndServe("127.0.0.1:8000", app.Router))
}

func (app *App) RouteAuth() {
	sub := app.Group("/api/auth", app.AuthHandler.HelloMsg, "GET")

	sub.RegisterSubHandler("/sign-up", app.AuthHandler.SignUp, "POST")
	sub.RegisterSubHandler("/sign-in", app.AuthHandler.SignIn, "POST")
}

func (app *App) RouteUsers() {
	sub := app.Group("/api/users", app.UsersHandler.HelloMsg, "GET")

	sub.RegisterSubHandler("/{id:[1-9]+}", app.UsersHandler.Get, "GET")
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

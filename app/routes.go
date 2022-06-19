package app

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Lemuriets/diary/model"
)

func (app *App) Run() {
	app.RouteAuth()
	app.RouteUsers()
	app.RouteMarks()
	app.RouteSchools()
	app.RouteClasses()
	app.RouteHomework()
	app.RouteShedules()

	fmt.Println("The server was started")
	log.Fatal(http.ListenAndServe("127.0.0.1:8000", app.Router))
}

func (app *App) RouteAuth() {
	sub := app.Router.PathPrefix("/auth").Subrouter()

	sub.HandleFunc("/sign-up", app.AuthHandler.SignUp).Methods("POST")
	sub.HandleFunc("/sign-in", app.AuthHandler.SignIn).Methods("POST")
}

func (app *App) RouteUsers() {
	sub := app.Router.PathPrefix("/users").Subrouter()

	sub.HandleFunc("/{id:[1-9]+}", app.UsersHandler.Get).Methods("GET")
	sub.HandleFunc("/update/{id:[1-9]+}", Authorization(
		app.UsersHandler.Update,
		model.ADMINISTRATOR,
	)).Methods("POST")
	sub.HandleFunc("/delete/{id:[1-9]+}", Authorization(
		app.UsersHandler.Delete,
		model.ADMINISTRATOR,
	)).Methods("POST")
	sub.HandleFunc("/multiple-delete", Authorization(
		app.UsersHandler.MultipleDelete,
		model.SUPERUSER,
	)).Methods("POST")
	sub.HandleFunc("/add-grade/{id:[1-9]+}", app.UsersHandler.AddGrade).Methods("GET")
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

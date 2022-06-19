package main

import (
	"log"

	"github.com/Lemuriets/diary/app"
	"github.com/joho/godotenv"
)

func main() {
	// . _ .
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal(err)
	}

	database := app.InitDB()
	server := app.NewApp(database)
	server.Run()
}

package main

import (
	"log"

	"github.com/Lemuriets/diary/app"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal(err)
	}

	server := app.NewApp()
	server.Run()
}

package main

import (
	"github.com/joho/godotenv"
	"github.com/kataras/golog"
	"short-url-service/app/server"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		golog.Errorf("Loading .env file error: %s", err.Error())
	}

	app := server.NewApp()
	app.Start()
}

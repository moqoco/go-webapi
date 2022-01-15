package main

import (
	"api/model"
	"api/server"

	"github.com/joho/godotenv"
)

func main() {
	// read env
	godotenv.Load()

	// DB connection gorm
	model.Conn()

	// server run
	server.Init()
}

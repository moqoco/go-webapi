package main

import (
	"api/model"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	// read env
	godotenv.Load()

	// DB connection gorm
	model.Conn()

	// user model
	user := model.User{}
	err := user.Migrate()
	if err != nil {
		log.Fatal(err.Error())
	}

	log.Print("migration completed.")
}

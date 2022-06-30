package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	env ENV
)

func Init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	env.Port = os.Getenv("PORT")
	env.AuthKey = os.Getenv("AUTH_KEY")
	env.DatabaseConnection = os.Getenv("DATABASE_CONNECTION")
	env.DatabaseName = os.Getenv("DATABASE_NAME")
}

func GetENV() ENV {
	return env
}

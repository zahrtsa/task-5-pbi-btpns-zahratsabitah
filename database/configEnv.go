package database

import (
	"log"

	"github.com/joho/godotenv"
)

func ConfigEnv() {
	err := godotenv.Load() //load env file
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
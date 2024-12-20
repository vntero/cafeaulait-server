package configs

import (
	"log"

	"github.com/joho/godotenv"
)

// load env variables
func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

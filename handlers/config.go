package handlers

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func DotEnvVariable(key string) string {

	// load .env file
	os.Getwd()

	err := godotenv.Load()

	if err != nil {
		log.Printf("Error loading .env file %s", err)
	}

	return os.Getenv(key)
}

package common

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading the .env file")
	}
}

func GetEnv(key string) (env_var string) {
	env_var = os.Getenv(key)
	return
}

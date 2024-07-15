package common

import (
	"log"
	"github.com/joho/godotenv"
	"os"
)

func LoadEnv(){
	err := godotenv.Load()
	if err != nil{
		log.Fatal("Error loading the .env file")
	}
}

func GetEnv(key string) (env_var string) {
	env_var = os.Getenv(key)
	return
}


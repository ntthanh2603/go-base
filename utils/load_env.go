package env

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func Load(key string) string {
	return os.Getenv(key)
}

func Init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

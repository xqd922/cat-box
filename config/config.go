package config

import (
	"os"

	"github.com/joho/godotenv"
)

func Config(key string) string {
	godotenv.Load(".env")
	return os.Getenv(key)
}

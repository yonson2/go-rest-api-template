package config

import (
	"log"
	"os"
	"sync"

	"github.com/joho/godotenv"
)

var once = sync.Once{}

type config struct {
	Port string
}

func getEnv(key, fallback string) string {
	once.Do(func() {
		err := godotenv.Load()
		if err != nil {
			log.Println("Error loading .env file")
		}
	})

	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return fallback
}

func New() config {
	return config{
		Port: getEnv("PORT", "8080"),
	}
}

package config

import (
	"log"
	"os"
	"sync"

	"github.com/joho/godotenv"
)

var Init sync.Once = sync.Once{}

func getEnv(key, fallback string) string {
	Init.Do(func() {
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

var Port = getEnv("PORT", "8080")
var URL = getEnv("URL", "http://localhost:"+Port)

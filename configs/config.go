package configs

import (
	"github.com/joho/godotenv"
	"log"
)

type Config struct {
}

func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file, using default config")
	}
	return &Config{}
}

package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Host          string
	Port          string
	User          string
	Password      string
	DBName        string
	SSLMode       string
	HostRedis     string
	PortRedis     string
	PasswordRedis string
}

func Initialize() Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	config := Config{
		Host:          os.Getenv("DB_HOST"),
		Port:          os.Getenv("DB_PORT"),
		User:          os.Getenv("DB_USER"),
		Password:      os.Getenv("DB_PASSWORD"),
		DBName:        os.Getenv("DB_NAME"),
		SSLMode:       os.Getenv("DB_SSLMODE"),
		HostRedis:     os.Getenv("REDIS_HOST"),
		PortRedis:     os.Getenv("REDIS_PORT"),
		PasswordRedis: os.Getenv("REDIS_PASSWORD"),
	}

	return config
}

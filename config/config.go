package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	FIREBASE_CREDENTIAL string
	CSRF_KEY            string
	PGDATABASE          string
	PGHOST              string
	PGPASSWORD          string
	PGPORT              string
	PGUSER              string
	FRONTEND_URL        string
	LISTEN_PORT         string
	GO_TODO_ENV         string
}

var c *Config

func Init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf(".envが読み込み出来ませんでした: %v", err)
	}
	c = &Config{
		FIREBASE_CREDENTIAL: os.Getenv("FIREBASE_CREDENTIAL"),
		CSRF_KEY:            os.Getenv("CSRF_KEY"),
		PGDATABASE:          os.Getenv("PGDATABASE"),
		PGHOST:              os.Getenv("PGHOST"),
		PGPASSWORD:          os.Getenv("PGPASSWORD"),
		PGPORT:              os.Getenv("PGPORT"),
		PGUSER:              os.Getenv("PGUSER"),
		FRONTEND_URL:        os.Getenv("FRONTEND_URL"),
		LISTEN_PORT:         os.Getenv("LISTEN_PORT"),
		GO_TODO_ENV:         os.Getenv("GO_TODO_ENV"),
	}
}

func GetConfig() *Config {
	return c
}

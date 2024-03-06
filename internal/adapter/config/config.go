package config

import (
	"os"

	"github.com/joho/godotenv"
)

type (
	Container struct {
		HTTP *HTTP
		App  *APP
		DB   *DB
	}
	HTTP struct {
		Env            string
		URL            string
		Port           string
		AllowedOrigins string
	}
	DB struct {
		ConnectionUrl string
	}
	APP struct {
		Env  string
		Name string
	}
)

func New() (*Container, error) {
	if os.Getenv("APP_ENV") != "production" {
		err := godotenv.Load()
		if err != nil {
			return nil, err
		}
	}

	app := &APP{
		Env:  os.Getenv("APP_ENV"),
		Name: os.Getenv("APP_NAME"),
	}

	http := &HTTP{
		Env:            os.Getenv("APP_ENV"),
		URL:            os.Getenv("HTTP_URL"),
		Port:           os.Getenv("HTTP_PORT"),
		AllowedOrigins: os.Getenv("HTTP_ALLOWED_ORIGINS"),
	}

	db := &DB{
		ConnectionUrl: os.Getenv("DB_URL"),
	}

	return &Container{
		http,
		app,
		db,
	}, nil
}

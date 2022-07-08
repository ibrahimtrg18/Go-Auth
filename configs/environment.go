package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Environment struct {
	User     string
	Password string
	Host     string
	Name     string
	Port     string
}

var environment *Environment

func InitEnvironment() *Environment {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASS")
	host := os.Getenv("DB_HOST")
	name := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")

	environment = &Environment{
		User:     user,
		Password: pass,
		Host:     host,
		Name:     name,
		Port:     port,
	}

	return environment
}

func GetEnvironment() *Environment {
	return environment
}

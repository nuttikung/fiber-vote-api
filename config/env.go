package config

import (
	"log"

	"github.com/joho/godotenv"
)

func InitializeEnv() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("LOAD ENV error")
	}
}

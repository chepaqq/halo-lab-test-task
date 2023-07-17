package main

import (
	"log"
	"os"

	"github.com/chepaqq99/halo-lab-test-task/pkg/db"
	"github.com/chepaqq99/halo-lab-test-task/pkg/httpserver"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	cfg := db.Config{
		Username: os.Getenv("POSTGRES_USER"),
		Password: os.Getenv("POSTGRES_PASSWORD"),
		Port:     os.Getenv("POSTGRES_PORT"),
		DBName:   os.Getenv("POSTGRES_DB"),
		Host:     os.Getenv("POSTGRES_HOST"),
		SSLMode:  os.Getenv("POSTGRES_SSLMODE"),
	}

	_, err = db.ConnectPostgres(cfg)
	if err != nil {
		log.Fatalf("failed to initialize db: %s", err.Error())
	}

	srv := new(httpserver.Server)
	if err := srv.Run("8000", nil); err != nil {
		log.Fatalln(err.Error())
	}
}

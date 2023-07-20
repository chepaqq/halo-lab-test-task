package main

import (
	"log"
	"os"

	"github.com/chepaqq99/halo-lab-test-task/internal/api/handler"
	"github.com/chepaqq99/halo-lab-test-task/internal/generation"
	"github.com/chepaqq99/halo-lab-test-task/pkg/httpserver"
	"github.com/joho/godotenv"
)

// @title Halo Lab test task
// @BasePath /
func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	go generation.InitGeneration()
	srv := new(httpserver.Server)
	if err := srv.Run(os.Getenv("PORT"), handler.InitRoutes()); err != nil {
		log.Fatalln(err.Error())
	}
}

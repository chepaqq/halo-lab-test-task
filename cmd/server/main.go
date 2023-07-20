package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	log "github.com/sirupsen/logrus"

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
	go func() {
		if err := srv.Run(os.Getenv("PORT"), handler.InitRoutes()); err != nil {
			log.Fatalln(err.Error())
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	if err := srv.Shutdown(context.Background()); err != nil {
		log.Fatalf("error occured on server shutting down: %s", err.Error())
	}
}

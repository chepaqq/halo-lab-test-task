package main

import (
	"log"
	"os"

	"github.com/chepaqq99/halo-lab-test-task/internal/api/handler"
	"github.com/chepaqq99/halo-lab-test-task/pkg/httpserver"
)

func main() {
	srv := new(httpserver.Server)
	if err := srv.Run(os.Getenv("PORT"), handler.InitRoutes()); err != nil {
		log.Fatalln(err.Error())
	}
}

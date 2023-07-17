package main

import (
	"log"

	"github.com/chepaqq99/halo-lab-test-task/pkg/httpserver"
)

func main() {
	srv := new(httpserver.Server)
	if err := srv.Run("8000", nil); err != nil {
		log.Fatalln(err.Error())
	}
}

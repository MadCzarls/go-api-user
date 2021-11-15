package main

import (
	"log"

	"github.com/mad-czarls/go-api-user/server"
)

func main() {
	if err := server.Run(); err != nil {
		log.Print(err)
	}
}

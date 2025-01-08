package main

import (
	"log"

	"github.com/shaply/File-Transfer/server/cmd/api"
)

func main() {

	server := api.NewServer("localhost:8080", nil)
	if err := server.Start(); err != nil {
		log.Fatal(err)
	}
}

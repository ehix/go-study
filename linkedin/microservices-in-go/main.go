package main

import (
	"log"

	"github.com/ehix/go-microservices/internal/database"
	"github.com/ehix/go-microservices/internal/server"
)

func main() {
	// create instance of DB
	db, err := database.NewDatabaseClient()
	if err != nil {
		log.Fatalf("failed to init Database Client: %s", err)
	}
	// create server and pass in database
	srv := server.NewEchoServer(db)
	if err := srv.Start(); err != nil {
		log.Fatalf(err.Error())
	}
}

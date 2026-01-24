package main

import (
	"log"
	"server/api/server"
)

func main() {
	svr, err := server.NewServer(server.ServerOptions{
		Host: "localhost",
		Port: "8080",
		Cors: nil,
	})
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Launching Server on %s:%s", svr.Options.Host, svr.Options.Port)
	log.Fatal(svr.Launch())
}

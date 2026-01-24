package main

import (
	"log"
	"server/api/server"

	"github.com/joho/godotenv"
)

func main() {
	log.Print("Loading env...")
	if err := godotenv.Load(); err != nil {
		log.Fatal("No evnironment variables found")
	}

	svr, err := server.NewServer(server.ServerOptions{
		Host: "localhost",
		Port: "8080",
		Cors: nil,
	})
	if err != nil {
		log.Fatal(err)
	}

	Main(svr)

	log.Printf("Launching Server on %s:%s", svr.Options.Host, svr.Options.Port)
	log.Fatal(svr.Launch())
}

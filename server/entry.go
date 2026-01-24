package main

import (
	"embed"
	"log"
	"server/api/server"
	"server/setup"

	"github.com/joho/godotenv"
	"github.com/pressly/goose/v3"
)

//go:embed sqlc/migrations/*.sql
var migrations embed.FS

func main() {
	goose.SetBaseFS(migrations)

	log.Print("Loading env...")
	if err := godotenv.Load(); err != nil {
		log.Fatal("No evnironment variables found")
	}

	svr, err := server.NewServer(server.ServerOptions{
		Host: "localhost",
		Port: "8080",
		Cors: nil,
	},
		setup.Handlers,
		setup.Middleware,
	)
	if err != nil {
		log.Fatal(err)
	}

	Main(svr)

	log.Printf("Launching Server on %s:%s", svr.Options.Host, svr.Options.Port)
	log.Fatal(svr.Launch())
}

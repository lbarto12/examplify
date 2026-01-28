package main

import (
	"embed"
	"log"
	"server/api/server"
	"server/environment"
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

	env, err := environment.Get()
	if err != nil {
		log.Fatal("Failed to load environment:", err)
	}

	svr, err := server.NewServer(server.ServerOptions{
		Host: env.ServerHost,
		Port: env.ServerPort,
		Cors: nil,
	},
		setup.Handlers,
		setup.Middleware,
		setup.ConfigureCors,
	)
	if err != nil {
		log.Fatal(err)
	}

	Main(svr)

	log.Printf("Launching Server on %s:%s", svr.Options.Host, svr.Options.Port)
	log.Fatal(svr.Launch())
}

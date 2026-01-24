package server

import (
	"fmt"
	"log"
	"net/http"
	"server/api/serviceaccess"
	"server/api/tools/externaltools/geminiapi"
	"server/api/tools/externaltools/minioapi"
	"server/api/tools/externaltools/postgresapi"
	"server/environment"

	"github.com/go-chi/chi/v5"
)

func NewServer(options ServerOptions) (*Server, error) {
	env, err := environment.Get()
	if err != nil {
		return nil, err
	}

	mux := chi.NewMux()

	// Init Postgres
	log.Println("connecting to postgres...")
	postgresClient, err := postgresapi.Connect(env)
	if err != nil {
		return nil, err
	}

	// Init Minio
	log.Println("connecting to minio...")
	minioClient, err := minioapi.Connect(env)
	if err != nil {
		return nil, err
	}

	// Init Gemini
	log.Println("connecting to gemini...")
	geminiClient, err := geminiapi.Connect(env)
	if err != nil {
		return nil, err
	}

	server := Server{
		Options: options,
		Services: serviceaccess.Access{
			Postgres: postgresClient,
			Minio:    minioClient,
			Gemini:   geminiClient,
		},
		Env: *env,
		Mux: mux,
	}

	var svr server_interface = &server

	return svr.(*Server), nil
}

func (server Server) Launch() error {
	addr := fmt.Sprintf("%s:%s", server.Options.Host, server.Options.Port)
	return http.ListenAndServe(addr, server.Mux)
}

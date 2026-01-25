package server

import (
	"fmt"
	"log"
	"net/http"
	"server/api/serviceaccess"
	"server/api/tools/externaltools/geminiapi"
	"server/api/tools/externaltools/gptapi"
	"server/api/tools/externaltools/minioapi"
	"server/api/tools/externaltools/postgresapi"
	"server/environment"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/pressly/goose/v3"
)

func NewServer(
	options ServerOptions,
	handlers func(mux *chi.Mux, services *serviceaccess.Access),
	middlewares func(mux *chi.Mux),
	corsConfig func(*environment.Vars) *cors.Options,
) (*Server, error) {
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

	if err := goose.SetDialect("postgres"); err != nil {
		return nil, err
	}

	if err := goose.Up(postgresClient, "sqlc/migrations"); err != nil {
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

	// Init OpenAI
	log.Println("connecting to openai...")
	openAIClient, err := gptapi.Connect(env)
	if err != nil {
		return nil, err
	}

	appServices := serviceaccess.Access{
		Postgres: postgresClient,
		Minio:    minioClient,
		Gemini:   geminiClient,
		OpenAI:   openAIClient,
	}

	mux.Use(cors.Handler(*corsConfig(env)))

	middlewares(mux)

	handlers(mux, &appServices)

	server := Server{
		Options:  options,
		Services: appServices,
		Env:      *env,
		Mux:      mux,
	}

	var svr server_interface = &server

	return svr.(*Server), nil
}

func (server Server) Launch() error {
	addr := fmt.Sprintf("%s:%s", server.Options.Host, server.Options.Port)
	return http.ListenAndServe(addr, server.Mux)
}

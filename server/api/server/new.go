package server

import (
	"fmt"
	"net/http"
	"server/api/serviceaccess"
	"server/environment"
)

func NewServer(options ServerOptions) (*Server, error) {
	env, err := environment.Get()
	if err != nil {
		return nil, err
	}

	mux := http.NewServeMux()

	muxHandler := http.Handler(mux)

	server := Server{
		Options:  options,
		Services: serviceaccess.Access{},
		Env:      *env,
		Mux:      &muxHandler,
	}

	return &server, nil
}

func (server Server) Launch() error {
	addr := fmt.Sprintf("%s:%s", server.Options.Host, server.Options.Port)
	return http.ListenAndServe(addr, *server.Mux)
}

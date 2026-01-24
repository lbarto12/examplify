package server

import (
	"fmt"
	"net/http"
	"server/api/serviceaccess"
)

func NewServer(options ServerOptions) (*Server, error) {

	mux := http.NewServeMux()

	muxHandler := http.Handler(mux)

	server := Server{
		Options:  options,
		Services: serviceaccess.Access{},
		Mux:      &muxHandler,
	}

	return &server, nil
}

func (server Server) Launch() error {
	addr := fmt.Sprintf("%s:%s", server.Options.Host, server.Options.Port)
	return http.ListenAndServe(addr, *server.Mux)
}

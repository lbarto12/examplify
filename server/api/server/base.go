package server

import (
	"server/api/serviceaccess"
	"server/environment"

	"github.com/go-chi/chi/v5"
)

type server_interface interface {
	Launch() error
}

type Server struct {
	Options  ServerOptions
	Services serviceaccess.Access
	Env      environment.Vars
	Mux      *chi.Mux
}

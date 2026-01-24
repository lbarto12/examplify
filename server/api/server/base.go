package server

import (
	"net/http"
	"server/api/serviceaccess"
	"server/environment"
)

type server_interface interface {
}

type Server struct {
	Options  ServerOptions
	Services serviceaccess.Access
	Env      environment.Vars
	Mux      *http.Handler
}

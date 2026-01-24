package server

import (
	"net/http"
	"server/api/serviceaccess"
)

type server_interface interface {
}

type Server struct {
	Options  ServerOptions
	Services serviceaccess.Access
	Mux      *http.Handler
}

package server

import "github.com/rs/cors"

type ServerOptions struct {
	Host string
	Port string
	Cors *cors.Options
}

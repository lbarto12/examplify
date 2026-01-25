package setup

import (
	"server/environment"

	"github.com/go-chi/cors"
)

func ConfigureCors(env *environment.Vars) *cors.Options {

	return &cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: true,
	}
}

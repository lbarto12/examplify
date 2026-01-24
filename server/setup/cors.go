package setup

import (
	"net/http"
	"server/environment"

	"github.com/go-chi/cors"
)

func ConfigureCors(env *environment.Vars) *cors.Options {
	origins := []string{
		"*", // lol
	}

	return &cors.Options{
		AllowedOrigins: origins,
		AllowedMethods: []string{
			http.MethodGet,
			http.MethodPost,
			http.MethodHead,
			http.MethodPut,
			http.MethodDelete,
		},
		AllowCredentials: true,
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"*"},
	}
}

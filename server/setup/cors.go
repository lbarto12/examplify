package setup

import (
	"server/environment"
	"strings"

	"github.com/go-chi/cors"
)

func ConfigureCors(env *environment.Vars) *cors.Options {
	// Split comma-separated origins for explicit whitelist
	origins := strings.Split(env.CorsAllowedOrigins, ",")

	return &cors.Options{
		AllowedOrigins:   origins,
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type"},
		AllowCredentials: true,
	}
}

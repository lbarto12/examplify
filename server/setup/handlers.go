package setup

import (
	"log"
	"server/api/serviceaccess"
	"server/core"
	"server/environment"
	"server/handlers/corehandlers"
	"server/handlers/generated/gencore"
	"server/handlers/generated/gensessions"
	"server/handlers/sessionhandlers"
	"server/sqlc/sqlgen"

	"github.com/go-chi/chi/v5"
)

const (
	Public  string = "/v1/public"
	Private string = "/v1/private"
)

func Handlers(mux *chi.Mux, services *serviceaccess.Access) {
	env, err := environment.Get()
	if err != nil {
		log.Fatal("Failed to load environment:", err)
	}

	core, err := core.NewCore(services, env)
	if err != nil {
		log.Fatal(err)
	}

	// Create single shared database query client
	queries := sqlgen.New(services.Postgres)

	gensessions.HandlerWithOptions(sessionhandlers.Handler{
		Services: services,
		Queries:  queries,
	}, gensessions.ChiServerOptions{
		BaseURL:    Public,
		BaseRouter: mux,
	})

	gencore.HandlerWithOptions(corehandlers.Handler{
		Services: services,
		Core:     core,
		Queries:  queries,
	}, gencore.ChiServerOptions{
		BaseURL:    Private,
		BaseRouter: mux,
	})
}

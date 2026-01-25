package setup

import (
	"log"
	"server/api/serviceaccess"
	"server/core"
	"server/handlers/generated/gensessions"
	"server/handlers/sessionhandlers"

	"github.com/go-chi/chi/v5"
)

const (
	Public  string = "/v1/public"
	Private string = "/v1/private"
)

func Handlers(mux *chi.Mux, services *serviceaccess.Access) {

	_, err := core.NewCore(services)
	if err != nil {
		log.Fatal(err)
	}

	gensessions.HandlerWithOptions(sessionhandlers.Handler{
		Services: services,
	}, gensessions.ChiServerOptions{
		BaseURL:    Public,
		BaseRouter: mux,
	})

}

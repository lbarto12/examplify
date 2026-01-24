package setup

import (
	"server/api/serviceaccess"
	"server/handlers/generated/gensessions"
	"server/handlers/sessionhandlers"

	"github.com/go-chi/chi/v5"
)

const (
	Public  string = "/v1/public"
	Private string = "/v1/private"
)

func Handlers(mux *chi.Mux, services *serviceaccess.Access) {

	gensessions.HandlerWithOptions(sessionhandlers.Handler{
		Services: services,
	}, gensessions.ChiServerOptions{
		BaseURL:    Public,
		BaseRouter: mux,
	})

}

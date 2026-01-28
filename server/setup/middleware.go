package setup

import (
	"net/http"
	"server/api/logging"
	"server/api/tools/internaltools/webtokens"

	"github.com/go-chi/chi/v5"
)

var PathExclusions = []string{
	"/v1/public", "v1/public", "/public", "public",
}

func Middleware(mux *chi.Mux) {

	// HTTP request logging middleware (first, so it logs everything)
	mux.Use(logging.HTTPLoggingMiddleware)

	// Authentication middleware
	mux.Use(func(next http.Handler) http.Handler {
		return webtokens.NewWebTokenMiddleware(next, webtokens.WebTokenMiddleWareConfig{
			PathPrefixExclusions: PathExclusions,
		})
	})
}

package webtokens

import (
	"github.com/labstack/gommon/log"

	"net/http"
	"strings"
)

type WebTokenMiddleWare struct {
	next   http.Handler
	config WebTokenMiddleWareConfig
}

type WebTokenMiddleWareConfig struct {
	PathPrefixExclusions []string
}

func NewWebTokenMiddleware(next http.Handler, config WebTokenMiddleWareConfig) *WebTokenMiddleWare {
	return &WebTokenMiddleWare{
		next,
		config,
	}
}

func (mw WebTokenMiddleWare) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := strings.TrimSpace(strings.ToLower(r.URL.Path))

	// Ignore the favicon
	if strings.HasSuffix(path, "favicon.ico") {
		mw.next.ServeHTTP(w, r)
		return
	}

	// Exclude middleware from paths
	if mw.config.PathPrefixExclusions != nil {
		for _, exclusion := range mw.config.PathPrefixExclusions {
			if strings.HasPrefix(path, exclusion) {
				mw.next.ServeHTTP(w, r)
				return
			}
		}
	}

	extractedJWTString, err := ExtractJWT(r)

	if err != nil {
		log.Error(err)
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	jwtParseResponse, err := ParseJWT(extractedJWTString)

	if err != nil || !jwtParseResponse.Valid {
		if jwtParseResponse.HasSeriousErrors {
			log.Error("POTENTIALLY MALICIOUS", err)
		} else {
			log.Error(err)
		}

		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	r = SetTokenInRequest(r, jwtParseResponse.Token)
	mw.next.ServeHTTP(w, r)
}

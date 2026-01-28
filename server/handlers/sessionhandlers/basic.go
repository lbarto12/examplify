package sessionhandlers

import (
	"net/http"
	"server/api/apirequests"
	"server/api/apiresponses"
	"server/handlers/generated/gensessions"
)

// Signs a user in.
// (POST /signin)
func (handler Handler) SignIn(w http.ResponseWriter, r *http.Request) {
	request, err := apirequests.Request[gensessions.SessionItem](r)
	if err != nil {
		apiresponses.BadRequest(w, "Invalid Request", err)
		return
	}

	result, err := handler.SessionManager.SignIn(r.Context(), string(request.Email), string(request.Password))
	if err != nil {
		apiresponses.InternalError(w, "Internal Error", err)
		return
	}

	apiresponses.Success(w, gensessions.SessionResponse{
		Token: result.Token,
	})
}

// Registers a new user.
// (POST /signup)
func (handler Handler) SignUp(w http.ResponseWriter, r *http.Request) {
	request, err := apirequests.Request[gensessions.SessionItem](r)
	if err != nil {
		apiresponses.BadRequest(w, "Invalid Request", err)
		return
	}

	result, err := handler.SessionManager.SignUp(r.Context(), string(request.Email), string(request.Password))
	if err != nil {
		apiresponses.InternalError(w, "Internal Error", err)
		return
	}

	apiresponses.Success(w, gensessions.SessionResponse{
		Token: result.Token,
	})
}

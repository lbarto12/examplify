package sessionhandlers

import (
	"net/http"
	"server/api/apirequests"
	"server/api/apiresponses"
	"server/api/tools/internaltools/passwords"
	"server/api/tools/internaltools/webtokens"
	"server/handlers/generated/gensessions"
	"server/sqlc/sqlgen"

	"github.com/labstack/gommon/log"
)

// Signs a user in.
// (POST /signin)
func (handler Handler) SignIn(w http.ResponseWriter, r *http.Request) {
	request, err := apirequests.Request[gensessions.SessionItem](r)
	if err != nil {
		log.Error(err)
		apiresponses.Error(w, "Invalid Request", http.StatusBadRequest)
		return
	}

	qtx := sqlgen.New(handler.Services.Postgres)

	user, err := qtx.GetUserAccountByEmail(r.Context(), string(request.Email))
	if err != nil {
		log.Error(err)
		apiresponses.Error(w, "Internal Error", http.StatusInternalServerError)
		return
	}

	if err := passwords.CompareHashAndPassword(string(request.Password), user.PasswordHash); err != nil {
		log.Error(err)
		apiresponses.Error(w, "Internal Error", http.StatusInternalServerError)
		return
	}

	token, _, err := webtokens.GenerateJWT(user.ID)
	if err != nil {
		log.Error(err)
		apiresponses.Error(w, "Internal Error", http.StatusInternalServerError)
		return
	}

	apiresponses.Success(w, gensessions.SessionResponse{
		Token: token,
	})

}

// Registers a new user.
// (POST /signup)
func (handler Handler) SignUp(w http.ResponseWriter, r *http.Request) {
	request, err := apirequests.Request[gensessions.SessionItem](r)
	if err != nil {
		log.Error(err)
		apiresponses.Error(w, "Invalid Request", http.StatusBadRequest)
		return
	}

	qtx := sqlgen.New(handler.Services.Postgres)

	pswd, err := passwords.GenerateFromPassword(request.Password, passwords.NewDefaultPasswordGenerationOptions())
	if err != nil {
		log.Error(err)
		apiresponses.Error(w, "Internal Error", http.StatusInternalServerError)
		return
	}

	user, err := qtx.CreateAccount(r.Context(), sqlgen.CreateAccountParams{
		UserEmail:    string(request.Email),
		PasswordHash: pswd,
	})
	if err != nil {
		log.Error(err)
		apiresponses.Error(w, "Internal Error", http.StatusInternalServerError)
		return
	}

	token, _, err := webtokens.GenerateJWT(user.ID)
	if err != nil {
		log.Error(err)
		apiresponses.Error(w, "Internal Error", http.StatusInternalServerError)
		return
	}

	apiresponses.Success(w, gensessions.SessionResponse{
		Token: token,
	})
}

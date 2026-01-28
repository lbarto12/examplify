package sessionhandlers

import (
	"net/http"
	"server/api/apirequests"
	"server/api/apiresponses"
	"server/api/tools/internaltools/passwords"
	"server/api/tools/internaltools/webtokens"
	"server/api/validation"
	"server/handlers/generated/gensessions"
	"server/sqlc/sqlgen"
)

// Signs a user in.
// (POST /signin)
func (handler Handler) SignIn(w http.ResponseWriter, r *http.Request) {
	request, err := apirequests.Request[gensessions.SessionItem](r)
	if err != nil {
		apiresponses.BadRequest(w, "Invalid Request", err)
		return
	}

	// Validate email format
	if err := validation.ValidateEmail(string(request.Email)); err != nil {
		apiresponses.BadRequest(w, err.Error(), err)
		return
	}

	user, err := handler.Queries.GetUserAccountByEmail(r.Context(), string(request.Email))
	if err != nil {
		apiresponses.InternalError(w, "Internal Error", err)
		return
	}

	if err := passwords.CompareHashAndPassword(string(request.Password), user.PasswordHash); err != nil {
		apiresponses.InternalError(w, "Internal Error", err)
		return
	}

	token, _, err := webtokens.GenerateJWT(user.ID)
	if err != nil {
		apiresponses.InternalError(w, "Internal Error", err)
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
		apiresponses.BadRequest(w, "Invalid Request", err)
		return
	}

	// Validate email format
	if err := validation.ValidateEmail(string(request.Email)); err != nil {
		apiresponses.BadRequest(w, err.Error(), err)
		return
	}

	// Validate password is non-empty
	if err := validation.ValidateNonEmpty("password", request.Password); err != nil {
		apiresponses.BadRequest(w, err.Error(), err)
		return
	}

	pswd, err := passwords.GenerateFromPassword(request.Password, passwords.NewDefaultPasswordGenerationOptions())
	if err != nil {
		apiresponses.InternalError(w, "Internal Error", err)
		return
	}

	user, err := handler.Queries.CreateAccount(r.Context(), sqlgen.CreateAccountParams{
		UserEmail:    string(request.Email),
		PasswordHash: pswd,
	})
	if err != nil {
		apiresponses.InternalError(w, "Internal Error", err)
		return
	}

	token, _, err := webtokens.GenerateJWT(user.ID)
	if err != nil {
		apiresponses.InternalError(w, "Internal Error", err)
		return
	}

	apiresponses.Success(w, gensessions.SessionResponse{
		Token: token,
	})
}

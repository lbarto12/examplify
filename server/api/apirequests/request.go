package apirequests

import (
	"encoding/json"
	"net/http"
	"server/api/tools/internaltools/webtokens"

	"github.com/google/uuid"
)

func Request[T any](r *http.Request) (*T, error) {
	var request T
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		return nil, err
	}
	return &request, nil
}

func User(r *http.Request) (*uuid.UUID, error) {
	return webtokens.GetUserIDfromRequest(r)
}

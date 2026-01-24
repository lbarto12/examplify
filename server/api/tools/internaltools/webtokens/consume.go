package webtokens

import (
	"context"
	"errors"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type ContextKey string

const (
	jwtTokenContext ContextKey = "TOKEN_CONTEXT_IDENTIFIER"
)

func SetTokenInRequest(r *http.Request, token *jwt.Token) *http.Request {
	return r.WithContext(context.WithValue(r.Context(), jwtTokenContext, token))
}

func GetTokenFromRequest(r *http.Request) *jwt.Token {
	value := r.Context().Value(jwtTokenContext)
	jwt, ok := value.(*jwt.Token)
	if !ok {
		return nil
	}
	return jwt
}

func GetUserIDFromToken(token *jwt.Token) (*uuid.UUID, error) {
	userIDText, err := token.Claims.GetSubject()
	if err != nil {
		return nil, err
	}

	userID, err := uuid.Parse(userIDText)
	if err != nil {
		return nil, err
	}

	return &userID, nil
}

func GetUserIDfromRequest(r *http.Request) (*uuid.UUID, error) {
	token := GetTokenFromRequest(r)
	if token == nil {
		return nil, errors.New("could not find or parse user token")
	}

	return GetUserIDFromToken(token)
}

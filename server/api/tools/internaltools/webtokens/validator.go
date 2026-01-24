package webtokens

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

type JWTParseResponse struct {
	Valid            bool
	TooEarly         bool
	Expired          bool
	InvalidSignature bool
	Malformed        bool
	HasSeriousErrors bool
	Token            *jwt.Token
	Claims           *JwtClaims
}

func ParseJWT(token string) (*JWTParseResponse, error) {
	claims := &JwtClaims{}
	jwtToken, err := jwt.ParseWithClaims(token, claims, func(t *jwt.Token) (any, error) {
		return jwtSecretKey, nil
	}, jwt.WithExpirationRequired())

	parsedResponse := &JWTParseResponse{
		Token:   jwtToken,
		Claims:  claims,
		Expired: err != nil && errors.Is(err, jwt.ErrTokenExpired),
	}

	// Do this as an if to catch the jwtToken nil error
	if err == nil && jwtToken.Valid {
		parsedResponse.Valid = true
	}

	switch {
	case errors.Is(err, jwt.ErrTokenMalformed):
		parsedResponse.Malformed = true
		parsedResponse.HasSeriousErrors = true
	case errors.Is(err, jwt.ErrTokenSignatureInvalid):
		parsedResponse.InvalidSignature = true
		parsedResponse.HasSeriousErrors = true
	case errors.Is(err, jwt.ErrTokenExpired):
		parsedResponse.Expired = true
	case errors.Is(err, jwt.ErrTokenNotValidYet):
		parsedResponse.TooEarly = true
	default:
		parsedResponse.HasSeriousErrors = true
	}

	return parsedResponse, err
}

func ExtractJWT(r *http.Request) (string, error) {
	reqToken := r.Header.Get("Authorization")
	if strings.TrimSpace(reqToken) == "" {
		return "", errors.New("missing JWT Authorization Header on secure request")
	}

	splitToken := strings.Split(reqToken, "Bearer")
	if len(splitToken) != 2 {
		return "", fmt.Errorf("attempted to authorize with malformed header %s", reqToken)
	}

	return strings.TrimSpace(splitToken[1]), nil
}

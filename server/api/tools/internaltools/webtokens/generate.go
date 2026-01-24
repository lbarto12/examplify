package webtokens

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type JwtClaims struct {
	jwt.RegisteredClaims
}

const JWTExpiry = time.Minute * 30

func GenerateJWT(subject uuid.UUID) (string, time.Time, error) {
	expires := time.Now().UTC().Add(JWTExpiry)

	claims := JwtClaims{
		jwt.RegisteredClaims{
			Subject:   subject.String(),
			IssuedAt:  &jwt.NumericDate{Time: time.Now().UTC()},
			ExpiresAt: &jwt.NumericDate{Time: expires.UTC()},
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedJwt, err := token.SignedString(jwtSecretKey)
	return signedJwt, expires, err
}

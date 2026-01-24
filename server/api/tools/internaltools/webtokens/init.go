package webtokens

import "server/environment"

var jwtSecretKey []byte

func Init(env *environment.Vars) {
	jwtSecretKey = []byte(env.JWTSecretKey)
}

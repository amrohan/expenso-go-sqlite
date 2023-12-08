package auth

import "github.com/go-chi/jwtauth"

func CreateJWTGenerator(secretKey string) *jwtauth.JWTAuth {
   auth := jwtauth.New("HS256", []byte(secretKey), nil)
   return auth
}

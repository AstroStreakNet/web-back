package auth

import "github.com/golang-jwt/jwt/v5"

type Claims struct {
	User string `json:"user"`
	Role string `json:"role"`
	jwt.RegisteredClaims
}

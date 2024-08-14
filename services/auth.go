package services

import (
	"github.com/golang-jwt/jwt/v5"
	"webback/auth"
)

type Auth interface {
	ParseToken(requestHeader string) (*jwt.Token, *auth.Claims, error)
	GenerateToken(user, role string) (string, error)
}

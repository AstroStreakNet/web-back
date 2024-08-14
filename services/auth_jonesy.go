package services

import (
	"github.com/golang-jwt/jwt/v5"
	"strings"
	"time"
	"webback/auth"
	"webback/repositories"
)

type AuthJonesy struct {
	userRepository repositories.User
	jwtExpiration  time.Duration
	jwtKey         string
}

func NewAuthJonesy(userRepository repositories.User, jwtExpiration int, jwtKey string) *AuthJonesy {
	return &AuthJonesy{
		userRepository: userRepository,
		jwtExpiration:  time.Duration(jwtExpiration),
		jwtKey:         jwtKey,
	}
}

func (service *AuthJonesy) ParseToken(requestHeader string) (*jwt.Token, *auth.Claims, error) {
	requestToken := strings.Split(requestHeader, " ")[1]
	claims := &auth.Claims{}
	token, err := jwt.ParseWithClaims(requestToken, claims, func(token *jwt.Token) (interface{}, error) {
		return service.jwtKey, nil
	})
	if err != nil {
		return nil, nil, err
	}
	return token, claims, nil
}

func (service *AuthJonesy) GenerateToken(user, role string) (string, error) {
	expiration := time.Now().Add(service.jwtExpiration * time.Minute)
	claims := &auth.Claims{
		User: user,
		Role: role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expiration),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(service.jwtKey)
}

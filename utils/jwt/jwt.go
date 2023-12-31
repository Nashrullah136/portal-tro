package jwtUtil

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"nashrul-be/crm/entities"
	"nashrul-be/crm/utils/localtime"
	"os"
	"time"
)

type JwtClaims struct {
	ID   uint
	Role string
	jwt.RegisteredClaims
}

func GenerateJWT(actor entities.User) (string, error) {
	now := localtime.Now()
	claims := JwtClaims{
		ID:   actor.ID,
		Role: actor.Role.RoleName,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "localhost",
			Subject:   actor.Username,
			Audience:  []string{"localhost"},
			ExpiresAt: jwt.NewNumericDate(localtime.Now().Add(1 * time.Hour)),
			NotBefore: jwt.NewNumericDate(*now),
			IssuedAt:  jwt.NewNumericDate(*now),
			ID:        "",
		},
	}
	signingKey := []byte(os.Getenv("JWT_SECRET"))
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(signingKey)
	if err != nil {
		return "", err
	}
	return signedToken, nil
}

func AuthenticateJWT(token string) (JwtClaims, error) {
	parsedToken, err := jwt.ParseWithClaims(token, &JwtClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	}, jwt.WithValidMethods([]string{"HS256"}))
	if err != nil {
		return JwtClaims{}, err
	}
	jwtClaims, ok := parsedToken.Claims.(*JwtClaims)
	if !ok || !parsedToken.Valid {
		return JwtClaims{}, errors.New("invalid token")
	}
	return *jwtClaims, nil
}

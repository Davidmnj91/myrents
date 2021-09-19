package jwt

import (
	"fmt"
	"github.com/Davidmnj91/myrents/pkg/domain/auth"
	"github.com/golang-jwt/jwt"
)

type jwtService struct {
	tokenSeed  string
	ttlMinutes int
}

func NewService(tokenSeed string, expirationTime int) auth.JWTService {
	return &jwtService{tokenSeed, expirationTime}
}

func (s *jwtService) SignJWT(claims auth.JWTClaims) (auth.JWTToken, error) {
	claims.Activate(s.ttlMinutes)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, ToClaims(claims))

	tokenString, err := token.SignedString(s.tokenSeed)
	if err != nil {
		return "", err
	}

	return auth.JWTToken(tokenString), nil
}

func (s *jwtService) DecodeJWT(tokenString auth.JWTToken) (auth.JWTClaims, error) {
	token, err := jwt.Parse(string(tokenString), func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return s.tokenSeed, nil
	})

	if err != nil {
		return auth.JWTClaims{}, err
	}

	if claims, ok := token.Claims.(*jwt.StandardClaims); ok && token.Valid {
		return ToDomain(claims), nil
	}

	return auth.JWTClaims{}, err
}

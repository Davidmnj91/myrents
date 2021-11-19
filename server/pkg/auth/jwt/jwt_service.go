package jwt

import (
	"fmt"
	"github.com/Davidmnj91/myrents/pkg/auth/domain"
	"github.com/golang-jwt/jwt"
)

type jwtService struct {
	tokenSeed  string
	ttlMinutes int64
}

func NewService(tokenSeed string, expirationTime int64) domain.JWTService {
	return &jwtService{tokenSeed, expirationTime}
}

func (s *jwtService) SignJWT(claims domain.JWTClaims) (domain.JWTToken, error) {
	claims.Activate(s.ttlMinutes)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, ToClaims(claims))

	tokenString, err := token.SignedString([]byte(s.tokenSeed))
	if err != nil {
		return "", err
	}

	return domain.JWTToken(tokenString), nil
}

func (s *jwtService) DecodeJWT(tokenString domain.JWTToken) (domain.JWTClaims, error) {
	token, err := jwt.ParseWithClaims(string(tokenString), &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(s.tokenSeed), nil
	})

	if err != nil {
		return domain.JWTClaims{}, err
	}

	if claims, ok := token.Claims.(*jwt.StandardClaims); ok && token.Valid {
		return ToDomain(claims), nil
	}

	return domain.JWTClaims{}, err
}

package jwt

import (
	"github.com/Davidmnj91/myrents/pkg/domain/auth"
	"github.com/golang-jwt/jwt"
)

func ToClaims(claims auth.JWTClaims) jwt.Claims {
	return jwt.MapClaims{
		"iss": claims.Issuer,
		"exp": claims.Exp,
		"iat": claims.Iat,
		"sub": claims.Sub,
	}
}

func ToDomain(claims *jwt.StandardClaims) auth.JWTClaims {
	return auth.JWTClaims{
		Issuer: claims.Issuer,
		Exp:    claims.ExpiresAt,
		Iat:    claims.IssuedAt,
		Sub:    claims.Subject,
	}
}

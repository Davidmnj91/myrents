package jwt

import (
	"github.com/Davidmnj91/myrents/pkg/domain/auth"
	"github.com/golang-jwt/jwt"
)

func ToClaims(claims auth.JWTClaims) jwt.Claims {
	return jwt.StandardClaims{
		Audience:  claims.Issuer,
		ExpiresAt: claims.Exp,
		IssuedAt:  claims.Iat,
		NotBefore: claims.Iat,
		Issuer:    claims.Issuer,
		Subject:   claims.Sub,
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

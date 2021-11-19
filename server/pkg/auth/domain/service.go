package domain

type JWTService interface {
	SignJWT(claims JWTClaims) (JWTToken, error)
	DecodeJWT(token JWTToken) (JWTClaims, error)
}

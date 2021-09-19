package auth

import (
	"github.com/Davidmnj91/myrents/pkg/domain/types"
	"time"
)

type JWTToken string

type JWTClaims struct {
	Issuer string
	Exp    int64
	Iat    int64
	Sub    string
}

func NewJWTClaims(uuid domain.UUID) JWTClaims {
	return JWTClaims{Issuer: "MyRents", Sub: uuid.String()}
}

func (j *JWTClaims) Activate(expirationTime int) {
	j.Exp = time.Now().Add(time.Minute * time.Duration(expirationTime)).Unix()
	j.Iat = time.Now().Unix()
}

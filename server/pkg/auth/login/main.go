package login

import (
	"github.com/Davidmnj91/myrents/pkg/auth/domain"
	domain2 "github.com/Davidmnj91/myrents/pkg/user/domain"
	"github.com/Davidmnj91/myrents/pkg/util/validation"
)

func NewLogin(validator validation.Validator, repo domain2.Repository, redis domain.Repository, jwtService domain.JWTService) Handler {
	service := NewService(repo, redis, jwtService)
	return NewHandler(service, validator)
}

package login

import (
	"github.com/Davidmnj91/myrents/pkg/domain/auth"
	"github.com/Davidmnj91/myrents/pkg/domain/user"
	"github.com/Davidmnj91/myrents/pkg/util/validation"
)

func NewLogin(validator validation.Validator, repo user.Repository, redis auth.Repository, jwtService auth.JWTService) Handler {
	service := NewService(repo, redis, jwtService)
	return NewHandler(service, validator)
}

package login

import (
	"github.com/Davidmnj91/myrents/pkg/domain/auth"
	"github.com/Davidmnj91/myrents/pkg/domain/user"
)

func NewLogin(repo user.Repository, redis auth.Repository, jwtService auth.JWTService) Handler {
	service := NewService(repo, redis, jwtService)
	return NewHandler(service)
}

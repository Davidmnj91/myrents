package login

import (
	"context"
	"errors"
	"fmt"
	"github.com/Davidmnj91/myrents/pkg/domain/auth"
	"github.com/Davidmnj91/myrents/pkg/domain/user"
)

type Service interface {
	Login(ctx context.Context, login *auth.Login) (auth.JWTToken, error)
}

type loginService struct {
	repo       user.Repository
	redis      auth.Repository
	jwtService auth.JWTService
}

func NewService(repo user.Repository, redis auth.Repository, jwtService auth.JWTService) Service {
	return &loginService{repo, redis, jwtService}
}

func (s *loginService) Login(ctx context.Context, login *auth.Login) (auth.JWTToken, error) {
	existing, err := s.repo.FindByUsername(ctx, login.Username)

	if err != nil {
		return "", errors.New(fmt.Sprintf("Internal App error: %s", err))
	}

	err = existing.Verify(login.Password)
	if err != nil {
		return "", err
	}

	claims := auth.NewJWTClaims(existing.UserUUID)

	jwtToken, err := s.jwtService.SignJWT(claims)
	if err != nil {
		return "", errors.New(fmt.Sprintf("Internal App error: %s", err))
	}

	err = s.redis.CreateSession(ctx, auth.Session{
		UserUUID: existing.UserUUID,
		Username: existing.Username,
	})

	if err != nil {
		return "", errors.New(fmt.Sprintf("Internal App error: %s", err))
	}

	return jwtToken, nil
}

package login

import (
	"context"
	"errors"
	"fmt"
	"github.com/Davidmnj91/myrents/pkg/auth/domain"
	domain2 "github.com/Davidmnj91/myrents/pkg/user/domain"
)

type Service interface {
	Login(ctx context.Context, login *domain.Login) (domain.JWTToken, error)
}

type loginService struct {
	repo       domain2.Repository
	redis      domain.Repository
	jwtService domain.JWTService
}

func NewService(repo domain2.Repository, redis domain.Repository, jwtService domain.JWTService) Service {
	return &loginService{repo, redis, jwtService}
}

func (s *loginService) Login(ctx context.Context, login *domain.Login) (domain.JWTToken, error) {
	existing, err := s.repo.FindByUsername(ctx, login.Username)

	if err != nil {
		return "", errors.New(fmt.Sprintf("Internal App error: %s", err))
	}

	err = existing.Verify(login.Password)
	if err != nil {
		return "", err
	}

	claims := domain.NewJWTClaims(existing.UserUUID)

	jwtToken, err := s.jwtService.SignJWT(claims)
	if err != nil {
		return "", errors.New(fmt.Sprintf("Internal App error: %s", err))
	}

	err = s.redis.CreateSession(ctx, domain.Session{
		UserUUID: existing.UserUUID,
		Username: existing.Username,
	})

	if err != nil {
		return "", errors.New(fmt.Sprintf("Internal App error: %s", err))
	}

	return jwtToken, nil
}

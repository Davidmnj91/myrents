package middleware

import (
	"context"
	"github.com/Davidmnj91/myrents/pkg/domain/auth"
	"github.com/Davidmnj91/myrents/pkg/domain/types"
	"strings"
)

type Service interface {
	CheckAuth(ctx context.Context, authorizationToken string) (auth.Session, error)
}

type service struct {
	jwtService        auth.JWTService
	sessionRepository auth.Repository
}

func NewService(jwtService auth.JWTService, sessionRepository auth.Repository) Service {
	return &service{jwtService, sessionRepository}
}

func (m *service) CheckAuth(ctx context.Context, authorizationToken string) (auth.Session, error) {
	token := strings.Replace(authorizationToken, "Bearer ", "", 1)

	claims, err := m.jwtService.DecodeJWT(auth.JWTToken(token))
	if err != nil {
		return auth.Session{}, err
	}

	uuid, err := domain.Parse(claims.Sub)
	if err != nil {
		return auth.Session{}, err
	}

	session, err := m.sessionRepository.GetSession(ctx, uuid)

	return session, err
}

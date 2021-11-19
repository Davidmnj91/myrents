package middleware

import (
	"context"
	"github.com/Davidmnj91/myrents/pkg/auth/domain"
	"github.com/Davidmnj91/myrents/pkg/types"
	"strings"
)

type Service interface {
	CheckAuth(ctx context.Context, authorizationToken string) (domain.Session, error)
}

type service struct {
	jwtService        domain.JWTService
	sessionRepository domain.Repository
}

func NewService(jwtService domain.JWTService, sessionRepository domain.Repository) Service {
	return &service{jwtService, sessionRepository}
}

func (m *service) CheckAuth(ctx context.Context, authorizationToken string) (domain.Session, error) {
	token := strings.Replace(authorizationToken, "Bearer ", "", 1)

	claims, err := m.jwtService.DecodeJWT(domain.JWTToken(token))
	if err != nil {
		return domain.Session{}, err
	}

	uuid, err := types.Parse(claims.Sub)
	if err != nil {
		return domain.Session{}, err
	}

	session, err := m.sessionRepository.GetSession(ctx, uuid)

	return session, err
}

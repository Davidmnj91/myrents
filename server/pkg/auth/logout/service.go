package logout

import (
	"context"
	"errors"
	"fmt"
	"github.com/Davidmnj91/myrents/pkg/auth/domain"
	"github.com/Davidmnj91/myrents/pkg/types"
)

type Service interface {
	Logout(ctx context.Context, userUUID types.UUID) error
}

type logoutService struct {
	redis domain.Repository
}

func NewService(redis domain.Repository) Service {
	return &logoutService{redis}
}

func (s *logoutService) Logout(ctx context.Context, userUUID types.UUID) error {
	err := s.redis.RemoveSession(ctx, userUUID)

	if err != nil {
		return errors.New(fmt.Sprintf("Internal App error: %s", err))
	}

	return nil
}

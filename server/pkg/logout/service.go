package logout

import (
	"context"
	"errors"
	"fmt"
	"github.com/Davidmnj91/myrents/pkg/domain/auth"
	"github.com/Davidmnj91/myrents/pkg/domain/types"
)

type Service interface {
	Logout(ctx context.Context, userUUID domain.UUID) error
}

type logoutService struct {
	redis auth.Repository
}

func NewService(redis auth.Repository) Service {
	return &logoutService{redis}
}

func (s *logoutService) Logout(ctx context.Context, userUUID domain.UUID) error {
	err := s.redis.RemoveSession(ctx, userUUID)

	if err != nil {
		return errors.New(fmt.Sprintf("Internal App error: %s", err))
	}

	return nil
}

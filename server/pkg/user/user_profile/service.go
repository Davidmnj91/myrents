package user_profile

import (
	"context"
	"errors"
	"fmt"
	"github.com/Davidmnj91/myrents/pkg/types"
	"github.com/Davidmnj91/myrents/pkg/user/domain"
)

type Service interface {
	Profile(ctx context.Context, userUUID types.UUID) (*domain.User, error)
}

type profileService struct {
	repo domain.Repository
}

func NewService(repo domain.Repository) Service {
	return &profileService{repo}
}

func (s *profileService) Profile(ctx context.Context, userUUID types.UUID) (*domain.User, error) {
	foundUser, err := s.repo.FindById(ctx, userUUID)

	if err != nil {
		return nil, errors.New(fmt.Sprintf("Internal App error: %s", err))
	}

	return foundUser, nil
}

package user_profile

import (
	"context"
	"errors"
	"fmt"
	"github.com/Davidmnj91/myrents/pkg/domain/types"
	"github.com/Davidmnj91/myrents/pkg/domain/user"
)

type Service interface {
	Profile(ctx context.Context, userUUID domain.UUID) (*user.User, error)
}

type profileService struct {
	repo user.Repository
}

func NewService(repo user.Repository) Service {
	return &profileService{repo}
}

func (s *profileService) Profile(ctx context.Context, userUUID domain.UUID) (*user.User, error) {
	foundUser, err := s.repo.FindById(ctx, userUUID)

	if err != nil {
		return nil, errors.New(fmt.Sprintf("Internal App error: %s", err))
	}

	return foundUser, nil
}

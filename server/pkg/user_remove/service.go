package user_remove

import (
	"context"
	"errors"
	"fmt"
	"github.com/Davidmnj91/myrents/pkg/domain/types"
	"github.com/Davidmnj91/myrents/pkg/domain/user"
)

type Service interface {
	RemoveAccount(ctx context.Context, uuid domain.UUID) error
}

type removeService struct {
	repo user.Repository
}

func NewService(repo user.Repository) Service {
	return &removeService{repo: repo}
}

func (s *removeService) RemoveAccount(ctx context.Context, uuid domain.UUID) error {
	foundUser, err := s.repo.FindById(ctx, uuid)

	if err != nil {
		return errors.New(fmt.Sprintf("Internal App error: %s", err))
	}

	foundUser.Delete()

	_, err = s.repo.Update(ctx, foundUser)

	if err != nil {
		return errors.New(fmt.Sprintf("Internal App error: %s", err))
	}

	return nil
}

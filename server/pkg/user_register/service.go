package user_register

import (
	"context"
	"errors"
	"fmt"
	"github.com/Davidmnj91/myrents/pkg/domain/user"
)

type Service interface {
	Register(ctx context.Context, user *user.User) error
}

type registerService struct {
	repo user.Repository
}

func NewService(repo user.Repository) Service {
	return &registerService{repo}
}

func (s *registerService) Register(ctx context.Context, user *user.User) error {
	existing, err := s.repo.Exists(ctx, user)

	if err != nil {
		return errors.New(fmt.Sprintf("Internal App error: %s", err))
	}

	if existing {
		return errors.New(ErrUserAlreadyExists)
	}

	err = user.Create()
	if err != nil {
		return errors.New(ErrPasswordInvalid)
	}

	err = s.repo.Add(ctx, user)

	if err != nil {
		return errors.New(fmt.Sprintf("Internal App error: %s", err))
	}

	return nil
}

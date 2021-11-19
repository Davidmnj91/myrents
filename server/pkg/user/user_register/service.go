package user_register

import (
	"context"
	"errors"
	"fmt"
	"github.com/Davidmnj91/myrents/pkg/user/domain"
)

type Service interface {
	Register(ctx context.Context, user *domain.User) error
}

type registerService struct {
	repo domain.Repository
}

func NewService(repo domain.Repository) Service {
	return &registerService{repo}
}

func (s *registerService) Register(ctx context.Context, user *domain.User) error {
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

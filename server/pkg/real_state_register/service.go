package real_state_register

import (
	"context"
	"errors"
	"fmt"
	"github.com/Davidmnj91/myrents/pkg/domain/real_state"
)

type Service interface {
	Register(ctx context.Context, realState *real_state.RealState) error
}

type registerService struct {
	repo real_state.Repository
}

func NewService(repo real_state.Repository) Service {
	return &registerService{repo}
}

func (s *registerService) Register(ctx context.Context, realState *real_state.RealState) error {
	existing, err := s.repo.Exists(ctx, realState)

	if err != nil {
		return errors.New(fmt.Sprintf("Internal App error: %s", err))
	}

	if existing {
		return errors.New(ErrRealStateAlreadyExists)
	}

	realState.Create()

	err = s.repo.Add(ctx, realState)

	if err != nil {
		return errors.New(fmt.Sprintf("Internal App error: %s", err))
	}

	return nil
}

package real_state_update

import (
	"context"
	"errors"
	"fmt"
	"github.com/Davidmnj91/myrents/pkg/real_state/domain"
)

type Service interface {
	Update(ctx context.Context, realState *domain.RealState) error
}

type registerService struct {
	repo domain.Repository
}

func NewService(repo domain.Repository) Service {
	return &registerService{repo}
}

func (s *registerService) Update(ctx context.Context, realState *domain.RealState) error {
	existing, err := s.repo.Exists(ctx, realState)

	if err != nil {
		return errors.New(fmt.Sprintf("Internal App error: %s", err))
	}

	if !existing {
		return errors.New(ErrRealStatNotExists)
	}

	realState.Update()

	_, err = s.repo.Update(ctx, realState)

	if err != nil {
		return errors.New(fmt.Sprintf("Internal App error: %s", err))
	}

	return nil
}

package real_state_list

import (
	"context"
	"errors"
	"fmt"
	"github.com/Davidmnj91/myrents/pkg/real_state/domain"
)

type Service interface {
	ListById(ctx context.Context, realState *domain.RealState) (*domain.RealState, error)
	ListAll(ctx context.Context, realState *domain.RealState) ([]domain.RealState, error)
}

type listerService struct {
	repo domain.Repository
}

func NewService(repo domain.Repository) Service {
	return &listerService{repo}
}

func (s *listerService) ListById(ctx context.Context, realState *domain.RealState) (*domain.RealState, error) {
	found, err := s.repo.FindByLandReference(ctx, realState.LandReference)

	if err != nil {
		return nil, errors.New(fmt.Sprintf("Internal App error: %s", err))
	}

	if found == nil {
		return nil, errors.New(ErrRealStateNotExists)
	}

	if !found.IsOwnedBy(realState.Landlord) {
		return nil, errors.New(ErrRealStateNotBelongToUser)
	}

	return found, nil
}

func (s *listerService) ListAll(ctx context.Context, realState *domain.RealState) ([]domain.RealState, error) {
	found, err := s.repo.FindByUserId(ctx, realState.Landlord)

	if err != nil {
		return nil, errors.New(fmt.Sprintf("Internal App error: %s", err))
	}

	return found, nil
}

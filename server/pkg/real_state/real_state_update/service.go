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
	found, err := s.repo.FindByLandReference(ctx, realState.LandReference)

	if err != nil {
		return errors.New(fmt.Sprintf("Internal App error: %s", err))
	}

	if found == nil {
		return errors.New(ErrRealStateNotExists)
	}

	if !found.IsOwnedBy(realState.Landlord) {
		return errors.New(ErrRealStateNotBelongToUser)
	}

	found.Update(*realState)

	_, err = s.repo.Update(ctx, found)

	if err != nil {
		return errors.New(fmt.Sprintf("Internal App error: %s", err))
	}

	return nil
}

package real_state_remove

import (
	"context"
	"errors"
	"fmt"
	"github.com/Davidmnj91/myrents/pkg/real_state/domain"
)

type Service interface {
	Remove(ctx context.Context, realState *domain.RealState) error
}

type removerService struct {
	repo domain.Repository
}

func NewService(repo domain.Repository) Service {
	return &removerService{repo}
}

func (s *removerService) Remove(ctx context.Context, realState *domain.RealState) error {
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

	realState.Delete()

	_, err = s.repo.Update(ctx, realState)

	if err != nil {
		return errors.New(fmt.Sprintf("Internal App error: %s", err))
	}

	return nil
}

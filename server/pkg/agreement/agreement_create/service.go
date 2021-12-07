package agreement_create

import (
	"context"
	"errors"
	"fmt"
	"github.com/Davidmnj91/myrents/pkg/agreement/domain"
	realState "github.com/Davidmnj91/myrents/pkg/real_state/domain"
)

type Service interface {
	Create(ctx context.Context, agreement *domain.Agreement) error
}

type createService struct {
	repo          domain.Repository
	realStateRepo realState.Repository
}

func NewService(repo domain.Repository, realStateRepo realState.Repository) Service {
	return &createService{repo, realStateRepo}
}

func (s *createService) Create(ctx context.Context, agreement *domain.Agreement) error {
	found, err := s.realStateRepo.FindByLandReference(ctx, agreement.RealState)

	if err != nil {
		return errors.New(fmt.Sprintf("Internal App error: %s", err))
	}

	if found != nil {
		return errors.New(ErrRealStateNotExists)
	}

	if !found.IsOwnedBy(agreement.Landlord) {
		return errors.New(ErrRealStateNotBelongToLandlord)
	}

	agreement.Create()

	err = s.repo.Add(ctx, agreement)

	if err != nil {
		return errors.New(fmt.Sprintf("Internal App error: %s", err))
	}

	return nil
}

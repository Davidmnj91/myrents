package agreement_list

import (
	"context"
	"errors"
	"fmt"
	"github.com/Davidmnj91/myrents/pkg/agreement/domain"
	"github.com/Davidmnj91/myrents/pkg/types"
)

type Service interface {
	List(ctx context.Context, userUuid types.UUID) ([]domain.Agreement, error)
}

type createService struct {
	repo domain.Repository
}

func NewService(repo domain.Repository) Service {
	return &createService{repo}
}

func (s *createService) List(ctx context.Context, userUuid types.UUID) ([]domain.Agreement, error) {
	agreements, err := s.repo.FindByLandlordOrTenant(ctx, userUuid)

	if err != nil {
		return nil, errors.New(fmt.Sprintf("Internal App error: %s", err))
	}

	return agreements, nil
}

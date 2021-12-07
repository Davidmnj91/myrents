package domain

import (
	"context"
	"github.com/Davidmnj91/myrents/pkg/types"
)

type Repository interface {
	Add(ctx context.Context, realState *Agreement) error
	FindById(ctx context.Context, uuid types.UUID) (*Agreement, error)
	FindByLandReference(ctx context.Context, landReference string) (*Agreement, error)
	FindByLandlordOrTenant(ctx context.Context, userUUID types.UUID) ([]Agreement, error)
}

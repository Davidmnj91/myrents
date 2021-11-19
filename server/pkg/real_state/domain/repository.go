package domain

import (
	"context"
	"github.com/Davidmnj91/myrents/pkg/types"
)

type Repository interface {
	Add(ctx context.Context, realState *RealState) error
	FindById(ctx context.Context, uuid types.UUID) (*RealState, error)
	FindByLandReference(ctx context.Context, landReference string) (*RealState, error)
	FindByUserId(ctx context.Context, userUUID types.UUID) ([]RealState, error)
	Exists(ctx context.Context, realState *RealState) (bool, error)
	Update(ctx context.Context, update *RealState) (*RealState, error)
}

package real_state

import (
	"context"
	domain "github.com/Davidmnj91/myrents/pkg/domain/types"
)

type Repository interface {
	Add(ctx context.Context, realState *RealState) error
	FindById(ctx context.Context, uuid domain.UUID) (*RealState, error)
	FindByUserId(ctx context.Context, userUUID domain.UUID) ([]RealState, error)
	Update(ctx context.Context, update *RealState) (*RealState, error)
}

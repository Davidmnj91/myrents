package real_state

import (
	"context"
	realState "github.com/Davidmnj91/myrents/pkg/domain/real_state"
	domain "github.com/Davidmnj91/myrents/pkg/domain/types"
	"go.mongodb.org/mongo-driver/mongo"
)

type mongoRepository struct {
	db *mongo.Collection
}

func NewRepository(db *mongo.Collection) realState.Repository {
	return &mongoRepository{db}
}

func (r *mongoRepository) Add(ctx context.Context, realState *realState.RealState) error {
	return nil
}

func (r *mongoRepository) FindById(ctx context.Context, uuid domain.UUID) (*realState.RealState, error) {
	return &realState.RealState{}, nil
}

func (r *mongoRepository) FindByUserId(ctx context.Context, userUUID domain.UUID) ([]realState.RealState, error) {
	return []realState.RealState{}, nil
}

func (r *mongoRepository) Update(ctx context.Context, update *realState.RealState) (*realState.RealState, error) {
	return &realState.RealState{}, nil
}

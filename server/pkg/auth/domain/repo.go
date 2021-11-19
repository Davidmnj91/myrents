package domain

import (
	"context"
	"github.com/Davidmnj91/myrents/pkg/types"
)

type Repository interface {
	GetSession(ctx context.Context, uuid types.UUID) (Session, error)
	CreateSession(ctx context.Context, session Session) error
	RefreshSession(ctx context.Context, session Session) error
	RemoveSession(ctx context.Context, userUUID types.UUID) error
}

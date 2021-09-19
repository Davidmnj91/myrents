package auth

import (
	"context"
	"github.com/Davidmnj91/myrents/pkg/domain/types"
)

type Repository interface {
	GetSession(ctx context.Context, uuid domain.UUID) (Session, error)
	CreateSession(ctx context.Context, session Session) error
	RefreshSession(ctx context.Context, session Session) error
	RemoveSession(ctx context.Context, userUUID domain.UUID) error
}

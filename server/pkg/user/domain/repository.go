package domain

import (
	"context"
	"github.com/Davidmnj91/myrents/pkg/types"
)

type Repository interface {
	Add(ctx context.Context, user *User) error
	FindById(ctx context.Context, uuid types.UUID) (*User, error)
	FindByUsername(ctx context.Context, username string) (*User, error)
	Exists(ctx context.Context, user *User) (bool, error)
	Update(ctx context.Context, user *User) (*User, error)
}

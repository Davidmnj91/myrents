package user

import (
	"context"
	"github.com/Davidmnj91/myrents/pkg/domain/types"
)

type Repository interface {
	Add(ctx context.Context, user *User) error
	FindById(ctx context.Context, uuid domain.UUID) (*User, error)
	FindByUsername(ctx context.Context, username string) (*User, error)
	Exists(ctx context.Context, user *User) (bool, error)
	Update(ctx context.Context, user *User) (*User, error)
}

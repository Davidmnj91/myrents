package user_profile

import (
	"github.com/Davidmnj91/myrents/pkg/domain/user"
)

func NewProfile(repo user.Repository) Handler {
	service := NewService(repo)
	return NewHandler(service)
}

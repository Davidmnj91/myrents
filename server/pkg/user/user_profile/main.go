package user_profile

import (
	"github.com/Davidmnj91/myrents/pkg/user/domain"
)

func NewProfile(repo domain.Repository) Handler {
	service := NewService(repo)
	return NewHandler(service)
}

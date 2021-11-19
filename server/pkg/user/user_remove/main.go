package user_remove

import (
	"github.com/Davidmnj91/myrents/pkg/user/domain"
)

func NewReMove(repo domain.Repository) Handler {
	service := NewService(repo)
	return NewHandler(service)
}

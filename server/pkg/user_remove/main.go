package user_remove

import "github.com/Davidmnj91/myrents/pkg/domain/user"

func NewReMove(repo user.Repository) Handler {
	service := NewService(repo)
	return NewHandler(service)
}

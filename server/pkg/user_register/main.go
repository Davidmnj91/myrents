package user_register

import "github.com/Davidmnj91/myrents/pkg/domain/user"

func NewRegister(repo user.Repository) Handler {
	service := NewService(repo)
	return NewHandler(service)
}

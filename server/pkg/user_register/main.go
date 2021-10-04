package user_register

import (
	"github.com/Davidmnj91/myrents/pkg/domain/user"
	"github.com/Davidmnj91/myrents/pkg/util/validation"
)

func NewRegister(repo user.Repository, validator validation.Validator) Handler {
	service := NewService(repo)
	return NewHandler(service, validator)
}

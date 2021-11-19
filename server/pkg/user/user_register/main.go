package user_register

import (
	"github.com/Davidmnj91/myrents/pkg/user/domain"
	"github.com/Davidmnj91/myrents/pkg/util/validation"
)

func NewRegister(repo domain.Repository, validator validation.Validator) Handler {
	service := NewService(repo)
	return NewHandler(service, validator)
}

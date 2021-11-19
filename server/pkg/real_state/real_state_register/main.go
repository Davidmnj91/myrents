package real_state_register

import (
	"github.com/Davidmnj91/myrents/pkg/real_state/domain"
	"github.com/Davidmnj91/myrents/pkg/util/validation"
)

func NewRealStateRegister(repo domain.Repository, validator validation.Validator) Handler {
	service := NewService(repo)
	return NewHandler(service, validator)
}

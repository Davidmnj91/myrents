package real_state_register

import (
	"github.com/Davidmnj91/myrents/pkg/domain/real_state"
	"github.com/Davidmnj91/myrents/pkg/util/validation"
)

func NewRealStateRegister(repo real_state.Repository, validator validation.Validator) Handler {
	service := NewService(repo)
	return NewHandler(service, validator)
}

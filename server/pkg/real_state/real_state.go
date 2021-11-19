package real_state

import (
	"github.com/Davidmnj91/myrents/pkg/real_state/domain"
	"github.com/Davidmnj91/myrents/pkg/real_state/real_state_register"
	"github.com/Davidmnj91/myrents/pkg/util/validation"
)

type RealStateModule struct {
	RegisterHandler real_state_register.Handler
}

func NewRealStateModule(
	realStateRepo domain.Repository,
	validator validation.Validator,
) *RealStateModule {
	realStateRegister := real_state_register.NewRealStateRegister(realStateRepo, validator)

	return &RealStateModule{
		RegisterHandler: realStateRegister,
	}
}

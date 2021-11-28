package real_state

import (
	"github.com/Davidmnj91/myrents/pkg/real_state/domain"
	"github.com/Davidmnj91/myrents/pkg/real_state/real_state_register"
	"github.com/Davidmnj91/myrents/pkg/real_state/real_state_update"
	"github.com/Davidmnj91/myrents/pkg/util/validation"
)

type RealStateModule struct {
	RegisterHandler real_state_register.Handler
	UpdaterHandler  real_state_update.Handler
}

func NewRealStateModule(
	realStateRepo domain.Repository,
	validator validation.Validator,
) *RealStateModule {
	realStateRegister := real_state_register.NewRealStateRegister(realStateRepo, validator)
	realStateUpdater := real_state_update.NewRealStateUpdater(realStateRepo, validator)

	return &RealStateModule{
		RegisterHandler: realStateRegister,
		UpdaterHandler:  realStateUpdater,
	}
}

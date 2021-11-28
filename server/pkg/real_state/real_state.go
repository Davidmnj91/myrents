package real_state

import (
	"github.com/Davidmnj91/myrents/pkg/real_state/domain"
	"github.com/Davidmnj91/myrents/pkg/real_state/real_state_list"
	"github.com/Davidmnj91/myrents/pkg/real_state/real_state_register"
	"github.com/Davidmnj91/myrents/pkg/real_state/real_state_remove"
	"github.com/Davidmnj91/myrents/pkg/real_state/real_state_update"
	"github.com/Davidmnj91/myrents/pkg/util/validation"
)

type RealStateModule struct {
	RegisterHandler real_state_register.Handler
	UpdaterHandler  real_state_update.Handler
	RemoverHandler  real_state_remove.Handler
	ListerHandler   real_state_list.Handler
}

func NewRealStateModule(
	realStateRepo domain.Repository,
	validator validation.Validator,
) *RealStateModule {
	realStateLister := real_state_list.NewRealStateLister(realStateRepo)
	realStateRegister := real_state_register.NewRealStateRegister(realStateRepo, validator)
	realStateUpdater := real_state_update.NewRealStateUpdater(realStateRepo, validator)
	realStateRemover := real_state_remove.NewRealStateRemover(realStateRepo)

	return &RealStateModule{
		ListerHandler:   realStateLister,
		RegisterHandler: realStateRegister,
		UpdaterHandler:  realStateUpdater,
		RemoverHandler:  realStateRemover,
	}
}

package real_state_register

import (
	"github.com/Davidmnj91/myrents/pkg/domain/real_state"
	uuid "github.com/Davidmnj91/myrents/pkg/domain/types"
)

type Mapper interface {
	ToDomain(register Register) *real_state.RealState
}

func ToDomain(register Register) *real_state.RealState {
	landLordUuid, err := uuid.Parse(register.Landlord)
	if err != nil {
	}

	return &real_state.RealState{
		LandReference: register.LandReference,
		Street:        register.Street,
		ZipCode:       register.ZipCode,
		Province:      register.Province,
		Country:       register.Country,
		Gateway:       register.Gateway,
		Door:          register.Door,
		SqMeters:      register.SqMeters,
		Landlord:      landLordUuid,
	}
}

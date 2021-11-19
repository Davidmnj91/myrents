package real_state_register

import (
	"github.com/Davidmnj91/myrents/pkg/real_state/domain"
	uuid "github.com/Davidmnj91/myrents/pkg/types"
)

type Mapper interface {
	ToDomain(register Register) *domain.RealState
}

func ToDomain(register Register) *domain.RealState {
	landLordUuid, err := uuid.Parse(register.Landlord)
	if err != nil {
	}

	return &domain.RealState{
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

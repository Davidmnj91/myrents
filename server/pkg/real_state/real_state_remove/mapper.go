package real_state_remove

import (
	"github.com/Davidmnj91/myrents/pkg/real_state/domain"
	uuid "github.com/Davidmnj91/myrents/pkg/types"
)

type Mapper interface {
	ToDomain(landReference string, landlord string) *domain.RealState
}

func ToDomain(landReference string, landlord string) *domain.RealState {
	landLordUuid, err := uuid.Parse(landlord)
	if err != nil {
	}

	return &domain.RealState{
		LandReference: landReference,
		Landlord:      landLordUuid,
	}
}

package real_state_update

import (
	"github.com/Davidmnj91/myrents/pkg/real_state/domain"
	uuid "github.com/Davidmnj91/myrents/pkg/types"
)

type Mapper interface {
	ToDomain(update Update) *domain.RealState
}

func ToDomain(landReference string, landlord string, update Update) *domain.RealState {
	landLordUuid, err := uuid.Parse(landlord)
	if err != nil {
	}

	return &domain.RealState{
		LandReference: landReference,
		SqMeters:      update.SqMeters,
		Landlord:      landLordUuid,
	}
}

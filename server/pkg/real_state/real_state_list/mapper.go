package real_state_list

import (
	"github.com/Davidmnj91/myrents/pkg/real_state/domain"
	uuid "github.com/Davidmnj91/myrents/pkg/types"
)

type Mapper interface {
	ToDomain(landReference string, landlord string) *domain.RealState
	ToHandler(state *domain.RealState) *ListRealState
}

func ToDomain(landlord string, landReference string) *domain.RealState {
	landLordUuid, err := uuid.Parse(landlord)
	if err != nil {
	}

	return &domain.RealState{
		LandReference: landReference,
		Landlord:      landLordUuid,
	}
}

func ToHandler(state *domain.RealState) *ListRealState {
	return &ListRealState{
		LandReference: state.LandReference,
		Street:        state.Street,
		ZipCode:       state.ZipCode,
		Province:      state.Province,
		Country:       state.Country,
		Gateway:       state.Gateway,
		Door:          state.Door,
		SqMeters:      state.SqMeters,
	}
}

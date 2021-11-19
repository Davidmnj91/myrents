package real_state

import (
	"github.com/Davidmnj91/myrents/pkg/domain/real_state"
	domain "github.com/Davidmnj91/myrents/pkg/domain/types"
)

func ToDomain(realState RealStateStorage) *real_state.RealState {
	uuid, _ := domain.Parse(realState.ID)
	landlordUuid, _ := domain.Parse(realState.ID)

	return &real_state.RealState{
		RealStateUUID: uuid,
		LandReference: realState.LandReference,
		Street:        realState.Street,
		ZipCode:       realState.ZipCode,
		Province:      realState.Province,
		Country:       realState.Country,
		Gateway:       realState.Gateway,
		Door:          realState.Door,
		SqMeters:      realState.SqMeters,
		Landlord:      landlordUuid,
		CreatedAt:     realState.CreatedAt,
		UpdatedAt:     realState.UpdatedAt,
		DeletedAt:     realState.DeletedAt,
	}
}

func ToRepository(realState *real_state.RealState) RealStateStorage {
	return RealStateStorage{
		LandReference: realState.LandReference,
		Street:        realState.Street,
		ZipCode:       realState.ZipCode,
		Province:      realState.Province,
		Country:       realState.Country,
		Gateway:       realState.Gateway,
		Door:          realState.Door,
		SqMeters:      realState.SqMeters,
		Landlord:      realState.Landlord.String(),
		CreatedAt:     realState.CreatedAt,
		UpdatedAt:     realState.UpdatedAt,
		DeletedAt:     realState.DeletedAt,
	}
}

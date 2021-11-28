package real_state

import (
	"github.com/Davidmnj91/myrents/pkg/real_state/domain"
	"github.com/Davidmnj91/myrents/pkg/types"
)

func ToDomain(realState RealStateStorage) *domain.RealState {
	uuid, _ := types.Parse(realState.ID)
	landlordUuid, _ := types.Parse(realState.Landlord)

	return &domain.RealState{
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

func ToRepository(realState *domain.RealState) RealStateStorage {
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

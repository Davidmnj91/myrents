package real_state

import (
	"github.com/Davidmnj91/myrents/pkg/domain/real_state"
	domain "github.com/Davidmnj91/myrents/pkg/domain/types"
)

func ToDomain(realState RealStateStorage) *real_state.RealState {
	uuid, _ := domain.Parse(realState.ID)

	return &real_state.RealState{
		RealStateUUID: uuid,
		Street:        realState.Street,
		ZipCode:       realState.ZipCode,
		Province:      realState.Province,
		Country:       realState.Country,
		Gateway:       realState.Gateway,
		Door:          realState.Door,
		SqMeters:      realState.SqMeters,
		// Landlord:  domain.Parse(string(realState.Landlord)),
		CreatedAt: realState.CreatedAt,
		UpdatedAt: realState.UpdatedAt,
		DeletedAt: realState.DeletedAt,
	}
}

func ToRepository(realState *real_state.RealState) RealStateStorage {
	return RealStateStorage{
		Street:   realState.Street,
		ZipCode:  realState.ZipCode,
		Province: realState.Province,
		Country:  realState.Country,
		Gateway:  realState.Gateway,
		Door:     realState.Door,
		SqMeters: realState.SqMeters,
		// Landlord:  realState.Landlord,
		CreatedAt: realState.CreatedAt,
		UpdatedAt: realState.UpdatedAt,
		DeletedAt: realState.DeletedAt,
	}
}

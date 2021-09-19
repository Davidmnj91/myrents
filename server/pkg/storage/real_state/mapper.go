package real_state

import (
	"github.com/Davidmnj91/myrents/pkg/domain/real_state"
)

func ToDomain(realState RealState) real_state.RealState {
	return real_state.RealState{
		Street:   realState.Street,
		ZipCode:  realState.ZipCode,
		Province: realState.Province,
		Country:  realState.Country,
		Gateway:  realState.Gateway,
		Door:     realState.Door,
		SqMeters: realState.SqMeters,
		// Landlord:  domain.Parse(string(realState.Landlord)),
		CreatedAt: realState.CreatedAt,
		UpdatedAt: realState.UpdatedAt,
		DeletedAt: realState.DeletedAt,
	}
}

func ToRepository(realState real_state.RealState) RealState {
	return RealState{
		// ID:        realState.ID,
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

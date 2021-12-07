package agreement

import (
	"github.com/Davidmnj91/myrents/pkg/agreement/domain"
	"github.com/Davidmnj91/myrents/pkg/types"
)

func ToDomain(storage AgreementStorage) *domain.Agreement {
	uuid, _ := types.Parse(storage.ID)
	landlordUuid, _ := types.Parse(storage.Landlord)
	tenantUuid, _ := types.Parse(storage.Tenant)
	startDate, _ := types.NewDate(storage.StartDate)
	endDate, _ := types.NewDate(storage.EndDate)

	return &domain.Agreement{
		AgreementUUID: uuid,
		RealState:     storage.RealState,
		Landlord:      landlordUuid,
		Tenant:        tenantUuid,
		MonthlyPrice:  storage.RentalCost,
		StartDate:     startDate,
		EndDate:       endDate,
		CreatedAt:     storage.CreatedAt,
		UpdatedAt:     storage.UpdatedAt,
		DeletedAt:     storage.DeletedAt,
	}
}

func ToRepository(agreement *domain.Agreement) AgreementStorage {
	return AgreementStorage{
		ID:         agreement.AgreementUUID.String(),
		RealState:  agreement.RealState,
		Landlord:   agreement.Landlord.String(),
		Tenant:     agreement.Tenant.String(),
		RentalCost: agreement.MonthlyPrice,
		StartDate:  agreement.StartDate.Format(),
		EndDate:    agreement.EndDate.Format(),
		CreatedAt:  agreement.CreatedAt,
		UpdatedAt:  agreement.UpdatedAt,
		DeletedAt:  agreement.DeletedAt,
	}
}

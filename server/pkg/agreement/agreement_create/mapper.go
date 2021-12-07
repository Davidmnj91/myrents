package agreement_create

import (
	"github.com/Davidmnj91/myrents/pkg/agreement/domain"
	"github.com/Davidmnj91/myrents/pkg/types"
)

type Mapper interface {
	ToDomain(agreement CreateAgreement) *domain.Agreement
}

func ToDomain(agreement CreateAgreement) *domain.Agreement {
	landLordUuid, _ := types.Parse(agreement.Landlord)
	tenantUuid, _ := types.Parse(agreement.Tenant)
	startDate, _ := types.NewDate(agreement.StartDate)
	endDate, _ := types.NewDate(agreement.EndDate)

	return &domain.Agreement{
		RealState:    agreement.RealState,
		Landlord:     landLordUuid,
		Tenant:       tenantUuid,
		MonthlyPrice: agreement.MonthlyPrice,
		StartDate:    startDate,
		EndDate:      endDate,
	}
}

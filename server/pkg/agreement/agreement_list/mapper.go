package agreement_list

import (
	"github.com/Davidmnj91/myrents/pkg/agreement/domain"
	"github.com/Davidmnj91/myrents/pkg/types"
)

type Mapper interface {
	ToDomain(user string) *domain.Agreement
	ToHandler(agreement *domain.Agreement) *ListAgreement
}

func ToDomain(user string) *domain.Agreement {
	userUuid, _ := types.Parse(user)

	return &domain.Agreement{
		Landlord: userUuid,
		Tenant:   userUuid,
	}
}

func ToHandler(agreement *domain.Agreement) *ListAgreement {
	return &ListAgreement{
		RealState:    agreement.RealState,
		Landlord:     agreement.Landlord.String(),
		Tenant:       agreement.Tenant.String(),
		MonthlyPrice: agreement.MonthlyPrice,
		StartDate:    agreement.StartDate.Format(),
		EndDate:      agreement.EndDate.Format(),
	}
}

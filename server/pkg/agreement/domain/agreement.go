package domain

import (
	"github.com/Davidmnj91/myrents/pkg/types"
	"time"
)

type Agreement struct {
	AgreementUUID types.UUID
	RealState     string
	Landlord      types.UUID
	Tenant        types.UUID
	MonthlyPrice  float32
	StartDate     types.Date
	EndDate       types.Date
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     time.Time
}

func (a *Agreement) Create() {
	a.CreatedAt = time.Now()
	a.UpdatedAt = time.Now()
}

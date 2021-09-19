package real_state

import (
	domain "github.com/Davidmnj91/myrents/pkg/domain/types"
	"time"
)

type RealState struct {
	RealStateUUID domain.UUID
	Street        string
	ZipCode       string
	Province      string
	Country       string
	Gateway       string
	Floor         string
	Door          string
	SqMeters      float32
	Landlord      domain.UUID
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     time.Time
}

func (u *RealState) Create() {
	u.CreatedAt = time.Now()
	u.UpdatedAt = time.Now()
}

func (u *RealState) Delete() {
	u.DeletedAt = time.Now()
}

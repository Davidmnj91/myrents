package domain

import (
	"github.com/Davidmnj91/myrents/pkg/types"
	"time"
)

type RealState struct {
	RealStateUUID types.UUID
	LandReference string
	Street        string
	ZipCode       string
	Province      string
	Country       string
	Gateway       string
	Floor         string
	Door          string
	SqMeters      string
	Landlord      types.UUID
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     time.Time
}

func (u *RealState) IsOwnedBy(landlord types.UUID) bool {
	return u.Landlord.Equals(landlord)
}

func (u *RealState) Create() {
	u.CreatedAt = time.Now()
	u.UpdatedAt = time.Now()
}

func (u *RealState) Update(state RealState) {
	u.SqMeters = state.SqMeters
	u.UpdatedAt = time.Now()
}

func (u *RealState) Delete() {
	u.DeletedAt = time.Now()
}

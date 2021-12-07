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

func (s *RealState) IsOwnedBy(landlord types.UUID) bool {
	return s.Landlord.Equals(landlord)
}

func (s *RealState) Create() {
	s.CreatedAt = time.Now()
	s.UpdatedAt = time.Now()
}

func (s *RealState) Update(state RealState) {
	s.SqMeters = state.SqMeters
	s.UpdatedAt = time.Now()
}

func (s *RealState) Delete() {
	s.DeletedAt = time.Now()
}

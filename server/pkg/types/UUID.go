package types

import (
	"github.com/google/uuid"
)

// UUID wrapper for Unique identifier
type UUID struct {
	uuid.UUID
}

func Parse(str string) (UUID, error) {
	parse, err := uuid.Parse(str)

	return UUID{parse}, err
}

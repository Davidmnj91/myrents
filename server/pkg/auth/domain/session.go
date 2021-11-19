package domain

import (
	"github.com/Davidmnj91/myrents/pkg/types"
)

type Session struct {
	UserUUID types.UUID
	Username string
}

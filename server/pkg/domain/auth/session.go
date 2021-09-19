package auth

import (
	"github.com/Davidmnj91/myrents/pkg/domain/types"
)

type Session struct {
	UserUUID domain.UUID
	Username string
}

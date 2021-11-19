package logout

import (
	"github.com/Davidmnj91/myrents/pkg/auth/domain"
)

func NewLogout(redis domain.Repository) Handler {
	service := NewService(redis)
	return NewHandler(service)
}

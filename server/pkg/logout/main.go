package logout

import "github.com/Davidmnj91/myrents/pkg/domain/auth"

func NewLogout(redis auth.Repository) Handler {
	service := NewService(redis)
	return NewHandler(service)
}

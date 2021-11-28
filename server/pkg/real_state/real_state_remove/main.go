package real_state_remove

import (
	"github.com/Davidmnj91/myrents/pkg/real_state/domain"
)

func NewRealStateRemover(repo domain.Repository) Handler {
	service := NewService(repo)
	return NewHandler(service)
}

package real_state_list

import (
	"github.com/Davidmnj91/myrents/pkg/real_state/domain"
)

func NewRealStateLister(repo domain.Repository) Handler {
	service := NewService(repo)
	return NewHandler(service)
}

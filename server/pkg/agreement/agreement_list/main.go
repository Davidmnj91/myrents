package agreement_list

import (
	"github.com/Davidmnj91/myrents/pkg/agreement/domain"
)

func NewAgreementList(repo domain.Repository) Handler {
	service := NewService(repo)
	return NewHandler(service)
}

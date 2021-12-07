package agreement_create

import (
	"github.com/Davidmnj91/myrents/pkg/agreement/domain"
	realState "github.com/Davidmnj91/myrents/pkg/real_state/domain"
	"github.com/Davidmnj91/myrents/pkg/util/validation"
)

func NewAgreementCreate(repo domain.Repository, realStateRepo realState.Repository, validator validation.Validator) Handler {
	service := NewService(repo, realStateRepo)
	return NewHandler(service, validator)
}

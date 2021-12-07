package agreement

import (
	"github.com/Davidmnj91/myrents/pkg/agreement/agreement_create"
	"github.com/Davidmnj91/myrents/pkg/agreement/agreement_list"
	"github.com/Davidmnj91/myrents/pkg/agreement/domain"
	realState "github.com/Davidmnj91/myrents/pkg/real_state/domain"
	"github.com/Davidmnj91/myrents/pkg/util/validation"
)

type AgreementModule struct {
	ListHandler   agreement_list.Handler
	CreateHandler agreement_create.Handler
}

func NewAgreementModule(
	repo domain.Repository,
	realStateRepo realState.Repository,
	validator validation.Validator,
) *AgreementModule {
	agreementLister := agreement_list.NewAgreementList(repo)
	agreementCreator := agreement_create.NewAgreementCreate(repo, realStateRepo, validator)

	return &AgreementModule{
		ListHandler:   agreementLister,
		CreateHandler: agreementCreator,
	}
}

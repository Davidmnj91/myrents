package login

import (
	"github.com/Davidmnj91/myrents/pkg/auth/domain"
)

func ToDomain(login Login) domain.Login {
	return domain.Login{
		Username: login.Username,
		Password: login.Password,
	}
}

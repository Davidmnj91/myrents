package login

import (
	"github.com/Davidmnj91/myrents/pkg/domain/auth"
)

func ToDomain(login Login) auth.Login {
	return auth.Login{
		Username: login.Username,
		Password: login.Password,
	}
}

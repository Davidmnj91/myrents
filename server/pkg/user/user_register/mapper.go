package user_register

import (
	"github.com/Davidmnj91/myrents/pkg/user/domain"
)

type Mapper interface {
	ToDomain(register Register) *domain.User
}

func ToDomain(register Register) *domain.User {
	birthDate, err := domain.NewBirthDate(register.BirthDate)
	if err != nil {

	}

	return &domain.User{
		Username:  register.Username,
		Password:  register.Password,
		Name:      register.Name,
		Surname:   register.Surname,
		IDNumber:  register.IDNumber,
		Email:     register.Email,
		Phone:     register.Phone,
		BirthDate: birthDate,
	}
}

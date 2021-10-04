package user_register

import (
	"github.com/Davidmnj91/myrents/pkg/domain/user"
)

type Mapper interface {
	ToDomain(register Register) *user.User
	ToHandler(user user.User) *interface{}
}

func ToDomain(register Register) *user.User {
	birthDate, err := user.NewBirthDate(register.BirthDate)
	if err != nil {

	}

	return &user.User{
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

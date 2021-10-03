package user_register

import (
	"github.com/Davidmnj91/myrents/pkg/domain/user"
	"time"
)

type Mapper interface {
	ToDomain(register Register) *user.User
	ToHandler(user user.User) *interface{}
}

func ToDomain(register Register) *user.User {
	birthDate, err := time.Parse("YYYY-MM-DD", register.BirthDate)
	if err != nil {
		birthDate = time.Now()
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

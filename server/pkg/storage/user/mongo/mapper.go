package mongo

import (
	"github.com/Davidmnj91/myrents/pkg/domain/user"
)

func ToDomain(person Person) user.User {
	return user.User{
		Username:  person.Username,
		Password:  person.Password,
		Name:      person.Name,
		Surname:   person.Surname,
		IDNumber:  person.IDNumber,
		Email:     person.Email,
		Phone:     person.Phone,
		BirthDate: person.BirthDate,
		CreatedAt: person.CreatedAt,
		UpdatedAt: person.UpdatedAt,
		DeletedAt: person.DeletedAt,
	}
}

func ToRepository(user user.User) Person {
	return Person{
		Username:  user.Username,
		Password:  user.Password,
		Name:      user.Name,
		Surname:   user.Surname,
		IDNumber:  user.IDNumber,
		Email:     user.Email,
		Phone:     user.Phone,
		BirthDate: user.BirthDate,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		DeletedAt: user.DeletedAt,
	}
}

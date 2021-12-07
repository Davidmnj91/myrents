package user

import (
	"github.com/Davidmnj91/myrents/pkg/types"
	"github.com/Davidmnj91/myrents/pkg/user/domain"
)

func ToDomain(person Person) *domain.User {
	birthDate, _ := types.NewDate(person.BirthDate)
	uuid, _ := types.Parse(person.ID)

	return &domain.User{
		UserUUID:  uuid,
		Username:  person.Username,
		Password:  person.Password,
		Name:      person.Name,
		Surname:   person.Surname,
		IDNumber:  person.IDNumber,
		Email:     person.Email,
		Phone:     person.Phone,
		BirthDate: birthDate,
		CreatedAt: person.CreatedAt,
		UpdatedAt: person.UpdatedAt,
		DeletedAt: person.DeletedAt,
	}
}

func ToRepository(user *domain.User) Person {
	return Person{
		Username:  user.Username,
		Password:  user.Password,
		Name:      user.Name,
		Surname:   user.Surname,
		IDNumber:  user.IDNumber,
		Email:     user.Email,
		Phone:     user.Phone,
		BirthDate: user.BirthDate.Format(),
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		DeletedAt: user.DeletedAt,
	}
}

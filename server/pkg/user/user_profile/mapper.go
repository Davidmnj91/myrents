package user_profile

import (
	"github.com/Davidmnj91/myrents/pkg/user/domain"
)

type Mapper interface {
	ToHandler(user domain.User) *interface{}
}

func ToHandler(user *domain.User) *Profile {
	return &Profile{
		Username:  user.Username,
		Name:      user.Name,
		Surname:   user.Surname,
		IDNumber:  user.IDNumber,
		Email:     user.Email,
		Phone:     user.Phone,
		BirthDate: user.BirthDate.Format(),
	}
}

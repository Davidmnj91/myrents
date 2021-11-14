package user_profile

import (
	"github.com/Davidmnj91/myrents/pkg/domain/user"
)

type Mapper interface {
	ToHandler(user user.User) *interface{}
}

func ToHandler(user *user.User) *Profile {
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

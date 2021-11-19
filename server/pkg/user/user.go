package user

import (
	"github.com/Davidmnj91/myrents/pkg/user/domain"
	"github.com/Davidmnj91/myrents/pkg/user/user_profile"
	"github.com/Davidmnj91/myrents/pkg/user/user_register"
	"github.com/Davidmnj91/myrents/pkg/user/user_remove"
	"github.com/Davidmnj91/myrents/pkg/util/validation"
)

type UserModule struct {
	RegisterHandler user_register.Handler
	ProfileHandler  user_profile.Handler
	DeleteHandler   user_remove.Handler
}

func NewUserModule(
	userRepo domain.Repository,
	validator validation.Validator,
) *UserModule {
	userRegisterHandler := user_register.NewRegister(userRepo, validator)
	userProfileHandler := user_profile.NewProfile(userRepo)
	userDeleteHandler := user_remove.NewReMove(userRepo)

	return &UserModule{
		RegisterHandler: userRegisterHandler,
		ProfileHandler:  userProfileHandler,
		DeleteHandler:   userDeleteHandler,
	}
}

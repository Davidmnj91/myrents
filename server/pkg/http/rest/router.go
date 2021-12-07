package rest

import (
	"github.com/Davidmnj91/myrents/pkg/agreement/agreement_create"
	"github.com/Davidmnj91/myrents/pkg/auth/login"
	"github.com/Davidmnj91/myrents/pkg/auth/logout"
	"github.com/Davidmnj91/myrents/pkg/auth/middleware"
	"github.com/Davidmnj91/myrents/pkg/real_state/real_state_list"
	"github.com/Davidmnj91/myrents/pkg/real_state/real_state_register"
	"github.com/Davidmnj91/myrents/pkg/real_state/real_state_remove"
	"github.com/Davidmnj91/myrents/pkg/real_state/real_state_update"
	"github.com/Davidmnj91/myrents/pkg/user/user_profile"
	"github.com/Davidmnj91/myrents/pkg/user/user_register"
	"github.com/Davidmnj91/myrents/pkg/user/user_remove"
	"github.com/gofiber/fiber/v2"
)

type Router interface {
	Serve(group fiber.Router)
}

type router struct {
	loginHandler        login.Handler
	logoutHandler       logout.Handler
	userRegisterHandler user_register.Handler
	userDeleteHandler   user_remove.Handler
	userProfileHandler  user_profile.Handler

	realStateListerHandler   real_state_list.Handler
	realStateRegisterHandler real_state_register.Handler
	realStateUpdaterHandler  real_state_update.Handler
	realStateRemoverHandler  real_state_remove.Handler

	agreementCreatorHandler agreement_create.Handler

	authMiddleware middleware.Middleware
}

type Routes struct {
	LoginHandler        login.Handler
	LogoutHandler       logout.Handler
	UserRegisterHandler user_register.Handler
	UserDeleteHandler   user_remove.Handler
	UserProfileHandler  user_profile.Handler

	RealStateListerHandler   real_state_list.Handler
	RealStateRegisterHandler real_state_register.Handler
	RealStateUpdaterHandler  real_state_update.Handler
	RealStateRemoverHandler  real_state_remove.Handler

	AgreementCreatorHandler agreement_create.Handler

	AuthMiddleware middleware.Middleware
}

func NewRouter(routes Routes) Router {
	return &router{
		loginHandler:             routes.LoginHandler,
		logoutHandler:            routes.LogoutHandler,
		userRegisterHandler:      routes.UserRegisterHandler,
		userDeleteHandler:        routes.UserDeleteHandler,
		userProfileHandler:       routes.UserProfileHandler,
		realStateListerHandler:   routes.RealStateListerHandler,
		realStateRegisterHandler: routes.RealStateRegisterHandler,
		realStateUpdaterHandler:  routes.RealStateUpdaterHandler,
		realStateRemoverHandler:  routes.RealStateRemoverHandler,
		agreementCreatorHandler:  routes.AgreementCreatorHandler,
		authMiddleware:           routes.AuthMiddleware,
	}
}

func (r *router) Serve(group fiber.Router) {
	group.Post("/login", r.loginHandler.Login)
	group.Delete("/logout", r.authMiddleware.CheckAuth(), r.logoutHandler.Logout)

	group.Post("/register", r.userRegisterHandler.Register)
	group.Delete("/removeAccount", r.authMiddleware.CheckAuth(), r.userDeleteHandler.RemoveAccount)

	group.Get("/profile", r.authMiddleware.CheckAuth(), r.userProfileHandler.Profile)

	group.Get("/real-state", r.authMiddleware.CheckAuth(), r.realStateListerHandler.List)
	group.Get("/real-state/:landReference", r.authMiddleware.CheckAuth(), r.realStateListerHandler.FindOne)
	group.Post("/real-state/register", r.authMiddleware.CheckAuth(), r.realStateRegisterHandler.Register)
	group.Put("/real-state/:landReference", r.authMiddleware.CheckAuth(), r.realStateUpdaterHandler.Update)
	group.Delete("/real-state/:landReference", r.authMiddleware.CheckAuth(), r.realStateRemoverHandler.Remove)

	group.Post("/agreement", r.authMiddleware.CheckAuth(), r.agreementCreatorHandler.Create)
}

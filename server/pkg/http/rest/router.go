package rest

import (
	"github.com/Davidmnj91/myrents/pkg/auth/login"
	"github.com/Davidmnj91/myrents/pkg/auth/logout"
	"github.com/Davidmnj91/myrents/pkg/auth/middleware"
	"github.com/Davidmnj91/myrents/pkg/real_state/real_state_register"
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

	realStateRegisterHandler real_state_register.Handler
	realStateUpdaterHandler  real_state_update.Handler

	authMiddleware middleware.Middleware
}

type Routes struct {
	LoginHandler        login.Handler
	LogoutHandler       logout.Handler
	UserRegisterHandler user_register.Handler
	UserDeleteHandler   user_remove.Handler
	UserProfileHandler  user_profile.Handler

	RealStateRegisterHandler real_state_register.Handler
	RealStateUpdaterHandler  real_state_update.Handler

	AuthMiddleware middleware.Middleware
}

func NewRouter(routes Routes) Router {
	return &router{
		loginHandler:             routes.LoginHandler,
		logoutHandler:            routes.LogoutHandler,
		userRegisterHandler:      routes.UserRegisterHandler,
		userDeleteHandler:        routes.UserDeleteHandler,
		userProfileHandler:       routes.UserProfileHandler,
		realStateRegisterHandler: routes.RealStateRegisterHandler,
		realStateUpdaterHandler:  routes.RealStateUpdaterHandler,
		authMiddleware:           routes.AuthMiddleware,
	}
}

func (r *router) Serve(group fiber.Router) {
	group.Post("/login", r.loginHandler.Login)
	group.Delete("/logout", r.authMiddleware.CheckAuth(), r.logoutHandler.Logout)

	group.Post("/register", r.userRegisterHandler.Register)
	group.Delete("/removeAccount", r.authMiddleware.CheckAuth(), r.userDeleteHandler.RemoveAccount)

	group.Get("/profile", r.authMiddleware.CheckAuth(), r.userProfileHandler.Profile)

	group.Post("/real-state/register", r.authMiddleware.CheckAuth(), r.realStateRegisterHandler.Register)
	group.Put("/real-state/:landReference", r.authMiddleware.CheckAuth(), r.realStateUpdaterHandler.Update)
}

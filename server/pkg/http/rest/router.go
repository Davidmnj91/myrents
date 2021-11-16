package rest

import (
	"github.com/Davidmnj91/myrents/pkg/login"
	"github.com/Davidmnj91/myrents/pkg/logout"
	"github.com/Davidmnj91/myrents/pkg/middleware"
	"github.com/Davidmnj91/myrents/pkg/user_profile"
	"github.com/Davidmnj91/myrents/pkg/user_register"
	"github.com/Davidmnj91/myrents/pkg/user_remove"
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

	authMiddleware middleware.Middleware
}

type Routes struct {
	LoginHandler        login.Handler
	LogoutHandler       logout.Handler
	UserRegisterHandler user_register.Handler
	UserDeleteHandler   user_remove.Handler
	UserProfileHandler  user_profile.Handler

	AuthMiddleware middleware.Middleware
}

func NewRouter(routes Routes) Router {
	return &router{
		loginHandler:        routes.LoginHandler,
		logoutHandler:       routes.LogoutHandler,
		userRegisterHandler: routes.UserRegisterHandler,
		userDeleteHandler:   routes.UserDeleteHandler,
		userProfileHandler:  routes.UserProfileHandler,
		authMiddleware:      routes.AuthMiddleware,
	}
}

func (r *router) Serve(group fiber.Router) {
	group.Post("/login", r.loginHandler.Login)
	group.Delete("/logout", r.authMiddleware.CheckAuth(), r.logoutHandler.Logout)

	group.Post("/register", r.userRegisterHandler.Register)
	group.Delete("/removeAccount", r.authMiddleware.CheckAuth(), r.userDeleteHandler.RemoveAccount)

	group.Get("/profile", r.authMiddleware.CheckAuth(), r.userProfileHandler.Profile)
}

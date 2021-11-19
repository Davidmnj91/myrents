package auth

import (
	"github.com/Davidmnj91/myrents/pkg/auth/domain"
	"github.com/Davidmnj91/myrents/pkg/auth/jwt"
	"github.com/Davidmnj91/myrents/pkg/auth/login"
	"github.com/Davidmnj91/myrents/pkg/auth/logout"
	"github.com/Davidmnj91/myrents/pkg/auth/middleware"
	user "github.com/Davidmnj91/myrents/pkg/user/domain"
	"github.com/Davidmnj91/myrents/pkg/util/validation"
)

type AuthModule struct {
	AuthMiddleware middleware.Middleware
	LoginHandler   login.Handler
	LogoutHandler  logout.Handler
}

func NewAuthModule(
	tokenSeed string,
	tokenTTL int64,
	redisRepo domain.Repository,
	userRepo user.Repository,
	validator validation.Validator,
) *AuthModule {
	jwtService := jwt.NewService(tokenSeed, tokenTTL)
	authService := middleware.NewService(jwtService, redisRepo)

	authMiddleware := middleware.NewAuthMiddleware(authService)
	loginHandler := login.NewLogin(validator, userRepo, redisRepo, jwtService)
	logoutHandler := logout.NewLogout(redisRepo)

	return &AuthModule{
		AuthMiddleware: authMiddleware,
		LoginHandler:   loginHandler,
		LogoutHandler:  logoutHandler,
	}
}

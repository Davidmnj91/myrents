package middleware

import (
	"github.com/gofiber/fiber/v2"
	"net/http"
)

type Middleware interface {
	CheckAuth() fiber.Handler
}

type middleware struct {
	authService Service
}

func NewAuthMiddleware(authService Service) Middleware {
	return &middleware{authService}
}

func (m *middleware) CheckAuth() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		bearerToken := ctx.Get("authorization")

		session, err := m.authService.CheckAuth(ctx.Context(), bearerToken)

		if err != nil {
			return ctx.Status(http.StatusUnauthorized).JSON(&fiber.Map{
				"status": http.StatusUnauthorized,
				"error":  err.Error(),
			})
		}

		// Always delete the header to avoid security breach
		ctx.Set("user", "")
		ctx.Set("user", session.UserUUID.String())

		return ctx.Next()
	}
}

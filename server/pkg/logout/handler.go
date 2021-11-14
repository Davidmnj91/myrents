package logout

import (
	"github.com/Davidmnj91/myrents/pkg/domain/types"
	"github.com/gofiber/fiber/v2"
)

type Handler interface {
	Logout(ctx *fiber.Ctx) error
}

type logoutHandler struct {
	service Service
}

func NewHandler(service Service) Handler {
	return &logoutHandler{service}
}

// 	Logout swagger:route DELETE /logout auth logout
//
// 	Logs out a user account from the system.
//
// 	Security:
//		loggedIn: []
//
// 	Responses:
// 		200: description:Successful logout
// 		500: description:Internal server error
func (h *logoutHandler) Logout(ctx *fiber.Ctx) error {
	token := ctx.Get("user")
	uuid, err := domain.Parse(token)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(err)
	}

	err = h.service.Logout(ctx.Context(), uuid)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(err)
	}

	return ctx.SendStatus(fiber.StatusAccepted)
}

package user_remove

import (
	"github.com/Davidmnj91/myrents/pkg/types"
	"github.com/gofiber/fiber/v2"
)

type Handler interface {
	RemoveAccount(ctx *fiber.Ctx) error
}

type removeHandler struct {
	service Service
}

func NewHandler(service Service) Handler {
	return &removeHandler{service}
}

// 	RemoveAccount swagger:route DELETE /removeAccount User remove-account
//
// 	Deletes an existing account on the system.
//
// 	Security:
//		loggedIn: []
//
// 	Responses:
// 		200: description:Successful delete
// 		500: description:Internal server error
func (h *removeHandler) RemoveAccount(ctx *fiber.Ctx) error {
	token := ctx.Get("user")

	uuid, err := types.Parse(token)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(err)
	}

	err = h.service.RemoveAccount(ctx.Context(), uuid)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(err)
	}

	return ctx.SendStatus(fiber.StatusAccepted)
}

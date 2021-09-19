package user_remove

import (
	"github.com/Davidmnj91/myrents/pkg/domain/types"
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

// RemoveAccount swagger:route DELETE /account user remove-account
//
// Deletes an existing account on the system.
//
// Responses:
// 		200: description:Successful delete
// 		500: description:Internal server error
func (h *removeHandler) RemoveAccount(ctx *fiber.Ctx) error {
	userUUID := ctx.Params("userUUID")

	if len(userUUID) == 0 {
		return ctx.SendStatus(fiber.StatusForbidden)
	}

	var uuid domain.UUID
	if err := uuid.UUID.Scan(userUUID); err != nil {
		return ctx.SendStatus(fiber.StatusProxyAuthRequired)
	}

	err := h.service.RemoveAccount(ctx.Context(), uuid)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(err)
	}

	return ctx.SendStatus(fiber.StatusAccepted)
}

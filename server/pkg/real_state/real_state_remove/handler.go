package real_state_remove

import (
	"github.com/gofiber/fiber/v2"
)

type Handler interface {
	Remove(ctx *fiber.Ctx) error
}

type removeHandler struct {
	service Service
}

func NewHandler(service Service) Handler {
	return &removeHandler{service}
}

// 	swagger:parameters remove-real-state
type RequestWrapper struct {
	// LandReference of the to be removed real state
	// in:path
	// required: true
	LandReference string
}

// 	Remove swagger:route DELETE /real-state/:landReference RealState remove-real-state
//
// 	Deletes the given real state for the owner in the system.
//
// 	Security:
//		loggedIn: []
//
// 	Responses:
// 		200: description:Successful remove
// 		500: description:Internal server error
func (h *removeHandler) Remove(ctx *fiber.Ctx) error {
	landlord := ctx.Get("user")
	landReference := ctx.Params("landReference")

	realState := ToDomain(landReference, landlord)

	err := h.service.Remove(ctx.Context(), realState)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(err)
	}

	return ctx.SendStatus(fiber.StatusAccepted)
}

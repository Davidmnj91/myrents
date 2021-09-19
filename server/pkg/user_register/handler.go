package user_register

import (
	"github.com/Davidmnj91/myrents/pkg/util/validation"
	"github.com/gofiber/fiber/v2"
)

type Handler interface {
	Register(ctx *fiber.Ctx) error
}

type registerHandler struct {
	service Service
}

func NewHandler(service Service) Handler {
	return &registerHandler{service}
}

// Register swagger:route POST /register user register
//
// Creates a new user in the system.
//
// Responses:
// 		200: description:Successful registration
// 		500: description:Internal server error
func (h *registerHandler) Register(ctx *fiber.Ctx) error {
	reqUser := &Register{}

	if err := ctx.BodyParser(reqUser); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(
			fiber.Map{
				"message": err.Error(),
			})
	}

	if err := validation.ValidateStruct(*reqUser); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(err)
	}

	user := ToDomain(*reqUser)

	err := h.service.Register(ctx.Context(), user)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(err)
	}

	return ctx.SendStatus(fiber.StatusAccepted)
}

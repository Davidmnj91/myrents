package user_register

import (
	"github.com/Davidmnj91/myrents/pkg/util/validation"
	"github.com/gofiber/fiber/v2"
)

type Handler interface {
	Register(ctx *fiber.Ctx) error
}

type registerHandler struct {
	service   Service
	validator validation.Validator
}

func NewHandler(service Service, validator validation.Validator) Handler {
	return &registerHandler{service, validator}
}

// 	swagger:parameters register-user
type RequestWrapper struct {
	// 	Body to register a user
	// 	in:body
	// 	required: true
	Body Register
}

// 	Register swagger:route POST /register user register-user
//
// 	Creates a new user in the system.
//
// 	Responses:
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

	if err := h.validator.ValidateStruct(*reqUser); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(err)
	}

	user := ToDomain(*reqUser)

	err := h.service.Register(ctx.Context(), user)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(err)
	}

	return ctx.SendStatus(fiber.StatusAccepted)
}

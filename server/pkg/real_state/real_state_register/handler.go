package real_state_register

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

// 	swagger:parameters register-real-state
type RequestWrapper struct {
	// 	Body to register a new real state
	// 	in:body
	// 	required: true
	Body Register
}

// 	Register swagger:route POST /real-state/register realState register-real-state
//
// 	Creates a new real state for the owner in the system.
//
// 	Security:
//		loggedIn: []
//
// 	Responses:
// 		200: description:Successful registration
// 		500: description:Internal server error
func (h *registerHandler) Register(ctx *fiber.Ctx) error {
	landlord := ctx.Get("user")
	reqRealState := &Register{}

	if err := ctx.BodyParser(reqRealState); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(
			fiber.Map{
				"message": err.Error(),
			})
	}

	reqRealState.Landlord = landlord

	if err := h.validator.ValidateStruct(*reqRealState); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(err)
	}

	realState := ToDomain(*reqRealState)

	err := h.service.Register(ctx.Context(), realState)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(err)
	}

	return ctx.SendStatus(fiber.StatusAccepted)
}

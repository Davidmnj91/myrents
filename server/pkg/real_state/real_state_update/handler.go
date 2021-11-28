package real_state_update

import (
	"github.com/Davidmnj91/myrents/pkg/util/validation"
	"github.com/gofiber/fiber/v2"
)

type Handler interface {
	Update(ctx *fiber.Ctx) error
}

type updateHandler struct {
	service   Service
	validator validation.Validator
}

func NewHandler(service Service, validator validation.Validator) Handler {
	return &updateHandler{service, validator}
}

// 	swagger:parameters update-real-state
type RequestWrapper struct {
	// LandReference of the to be updated real state
	// in:path
	// required: true
	LandReference string

	// 	Body to update a new real state
	// 	in:body
	// 	required: true
	Body Update
}

// 	Register swagger:route PUT /real-state/:landReference realState update-real-state
//
// 	Updates the given real state for the owner in the system.
//
// 	Security:
//		loggedIn: []
//
// 	Responses:
// 		200: description:Successful update
// 		500: description:Internal server error
func (h *updateHandler) Update(ctx *fiber.Ctx) error {
	landlord := ctx.Get("user")
	landReference := ctx.Params("landReference")
	reqRealState := &Update{}

	if err := ctx.BodyParser(reqRealState); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(
			fiber.Map{
				"message": err.Error(),
			})
	}

	if err := h.validator.ValidateStruct(*reqRealState); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(err)
	}

	realState := ToDomain(landReference, landlord, *reqRealState)

	err := h.service.Update(ctx.Context(), realState)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(err)
	}

	return ctx.SendStatus(fiber.StatusAccepted)
}

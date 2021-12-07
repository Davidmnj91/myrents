package agreement_create

import (
	"github.com/Davidmnj91/myrents/pkg/util/validation"
	"github.com/gofiber/fiber/v2"
)

type Handler interface {
	Create(ctx *fiber.Ctx) error
}

type createHandler struct {
	service   Service
	validator validation.Validator
}

func NewHandler(service Service, validator validation.Validator) Handler {
	return &createHandler{service, validator}
}

// 	swagger:parameters create-agreement
type RequestWrapper struct {
	// 	Body to create a new agreement
	// 	in:body
	// 	required: true
	Body CreateAgreement
}

// 	Create swagger:route POST /agreement Agreement create-agreement
//
// 	Creates a new agreement over a real state in the system
//
// 	Security:
//		loggedIn: []
//
// 	Responses:
// 		200: description:Successful created
// 		500: description:Internal server error
func (h *createHandler) Create(ctx *fiber.Ctx) error {
	landlord := ctx.Get("user")
	reqAgreement := &CreateAgreement{}

	if err := ctx.BodyParser(reqAgreement); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(
			fiber.Map{
				"message": err.Error(),
			})
	}

	reqAgreement.Landlord = landlord

	if err := h.validator.ValidateStruct(*reqAgreement); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(err)
	}

	agreement := ToDomain(*reqAgreement)

	err := h.service.Create(ctx.Context(), agreement)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(err)
	}

	return ctx.SendStatus(fiber.StatusAccepted)
}

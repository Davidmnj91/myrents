package agreement_list

import (
	"github.com/gofiber/fiber/v2"
)

type Handler interface {
	List(ctx *fiber.Ctx) error
}

type listHandler struct {
	service Service
}

func NewHandler(service Service) Handler {
	return &listHandler{service}
}

// 	swagger:response ListAgreements
type ResponseWrapper struct {
	// List of the user's agreements
	// in:body
	Payload []ListAgreement `json:"body"`
}

// 	List swagger:route GET /agreement Agreement list-agreements
//
// 	Returns the list of the registered agreements in the system for a user
//
// 	Security:
//		loggedIn: []
//
// 	Responses:
// 		200: ListAgreements description:Successful created
// 		500: description:Internal server error
func (h *listHandler) List(ctx *fiber.Ctx) error {
	user := ctx.Get("user")
	agreement := ToDomain(user)

	retrieved, err := h.service.List(ctx.Context(), agreement.Landlord)

	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(err)
	}

	var agreements []ListAgreement
	for _, result := range retrieved {
		agreements = append(agreements, *ToHandler(&result))
	}

	return ctx.Status(fiber.StatusOK).JSON(agreements)
}

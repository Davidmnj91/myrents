package real_state_list

import (
	"github.com/gofiber/fiber/v2"
)

type Handler interface {
	FindOne(ctx *fiber.Ctx) error
	List(ctx *fiber.Ctx) error
}

type listHandler struct {
	service Service
}

func NewHandler(service Service) Handler {
	return &listHandler{service}
}

// 	swagger:parameters list-real-state
type RequestWrapper struct {
	// LandReference of the required real state
	// in:path
	// required: true
	LandReference string
}

// 	swagger:response ListRealStates
type ResponseWrapper struct {
	// List of the user's real states
	// in:body
	Payload []ListRealState `json:"body"`
}

// 	List swagger:route GET /real-state/:landReference RealState list-real-state
//
// 	Returns the details of the searched LandReference for the owner in the system.
//
// 	Security:
//		loggedIn: []
//
// 	Responses:
// 		200: ListRealState description:Successful list
// 		500: description:Internal server error
func (h *listHandler) FindOne(ctx *fiber.Ctx) error {
	landlord := ctx.Get("user")
	landReference := ctx.Params("landReference")

	realState := ToDomain(landlord, landReference)

	retrieved, err := h.service.ListById(ctx.Context(), realState)

	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(err)
	}

	response := ToHandler(retrieved)

	return ctx.Status(fiber.StatusOK).JSON(response)
}

// 	List swagger:route GET /real-state RealState list-real-states
//
// 	Returns the existing real states for the owner in the system.
//
// 	Security:
//		loggedIn: []
//
// 	Responses:
// 		200: ListRealStates description:Successful list
// 		500: description:Internal server error
func (h *listHandler) List(ctx *fiber.Ctx) error {
	landlord := ctx.Get("user")

	realState := ToDomain(landlord, "")

	retrieved, err := h.service.ListAll(ctx.Context(), realState)

	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(err)
	}

	var realStates []ListRealState

	for _, result := range retrieved {
		realStates = append(realStates, *ToHandler(&result))
	}

	return ctx.Status(fiber.StatusOK).JSON(realStates)
}

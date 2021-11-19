package user_profile

import (
	"github.com/Davidmnj91/myrents/pkg/types"
	"github.com/gofiber/fiber/v2"
)

type Handler interface {
	Profile(ctx *fiber.Ctx) error
}

type profileHandler struct {
	service Service
}

func NewHandler(service Service) Handler {
	return &profileHandler{service}
}

// 	Profile swagger:route GET /profile user profile-user
//
// 	Returns the authenticated user's profile.
//
// 	Security:
//		loggedIn: []
//
// 	Responses:
// 		200: Profile description:Successful retrieval
// 		500: description:Internal server error
func (h *profileHandler) Profile(ctx *fiber.Ctx) error {
	token := ctx.Get("user")
	uuid, err := types.Parse(token)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(err)
	}

	user, err := h.service.Profile(ctx.Context(), uuid)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(err)
	}

	profile := ToHandler(user)

	return ctx.Status(fiber.StatusOK).JSON(profile)
}

package login

import (
	"github.com/Davidmnj91/myrents/pkg/util/validation"
	"github.com/gofiber/fiber/v2"
)

type Handler interface {
	Login(ctx *fiber.Ctx) error
}

type loginHandler struct {
	service Service
}

func NewHandler(service Service) Handler {
	return &loginHandler{service}
}

// Login swagger:route POST /login auth login-user
//
// Logs in a user account into the system.
//
// Responses:
// 		200: description:Successful login, bearer token in "authorization" header
// 		500: description:Internal server error
func (h *loginHandler) Login(ctx *fiber.Ctx) error {
	loginUser := &Login{}

	if err := ctx.BodyParser(loginUser); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(
			fiber.Map{
				"message": err.Error(),
			})
	}

	if err := validation.ValidateStruct(*loginUser); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(err)
	}

	login := ToDomain(*loginUser)

	token, err := h.service.Login(ctx.Context(), &login)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(err)
	}

	ctx.Set("authorization", "Bearer "+string(token))

	return ctx.SendStatus(fiber.StatusOK)
}

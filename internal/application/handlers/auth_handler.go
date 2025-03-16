package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ssyan-dev/go-fiber-oauth/pkg/response"
)

type AuthHandler struct{}

func NewAuthHandler() *AuthHandler {
	return &AuthHandler{}
}

// GetHelloWorld godoc
// @Summary      Check status
// @Description  Hello world!
// @Tags         Auth
// @Produce      json
// @Success      200
// @Router       /auth [get]
func (h *AuthHandler) GetHelloWorld(ctx *fiber.Ctx) error {
	data := fiber.Map{
		"hello": "world",
	}

	return response.SuccessResponse(ctx, fiber.StatusOK, "hello world!", data)
}

package handlers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/ssyan-dev/go-fiber-oauth/internal/application/config"
	"github.com/ssyan-dev/go-fiber-oauth/pkg/response"
	"github.com/ssyan-dev/go-fiber-oauth/pkg/utils/github_info"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/github"
)

type AuthHandler struct {
	Config       *config.Config
	GitHubConfig oauth2.Config
}

func NewAuthHandler(cfg *config.Config) *AuthHandler {
	return &AuthHandler{
		Config: cfg,
		GitHubConfig: oauth2.Config{
			ClientID:     cfg.OAuth.GitHub.ClientID,
			ClientSecret: cfg.OAuth.GitHub.ClientSecret,
			RedirectURL:  cfg.OAuth.GitHub.RedirectURL,
			Endpoint:     github.Endpoint,
			Scopes:       []string{"read:user", "user:email"},
		},
	}
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

// GetGitHub godoc
// @Summary      GitHub OAuth
// @Description  Redirect to GitHub Auth
// @Tags         Auth
// @Produce      json
// @Success      200
// @Router       /auth/github [get]
func (h *AuthHandler) GetGitHub(ctx *fiber.Ctx) error {
	url := h.GitHubConfig.AuthCodeURL("", oauth2.AccessTypeOffline)
	return ctx.Redirect(url)
}

// GetGitHubCallback godoc
// @Summary      GitHub Auth Callback
// @Description  GitHub Auth Callback
// @Tags         Auth
// @Produce      json
// @Success      200
// @Router       /auth/github/callback [get]
func (h *AuthHandler) GetGitHubCallback(ctx *fiber.Ctx) error {
	code := ctx.Query("code")
	token, err := h.GitHubConfig.Exchange(ctx.Context(), code)
	if err != nil {
		return response.ErrorResponse(ctx, fiber.StatusUnauthorized, "Failed to get token")
	}

	email, err := github_info.GetUserInfo(ctx.Context(), &h.GitHubConfig, token)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).SendString(fmt.Sprintf("Failed to get user info: %v", err))
	}

	data := fiber.Map{
		"token": token,
		"email": email,
	}

	return response.SuccessResponse(ctx, fiber.StatusOK, "Authorized", data)
}

package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
	"github.com/ssyan-dev/go-fiber-oauth/internal/application/config"
	"github.com/ssyan-dev/go-fiber-oauth/internal/application/handlers"

	_ "github.com/ssyan-dev/go-fiber-oauth/docs"
)

type Server struct {
	Config *config.Config
	Http   *fiber.App
}

// @title Go Fiber OAuth
// @version 1.0
// @description API for OAuth using Go Fiber

// @contact.name Developer
// @contact.email me@ssyan.ru

// @license.name MIT
// @license.url https://github.com/ssyan-dev/go-fiber-oauth/blob/main/LICENSE

// @host localhost:8080
// @BasePath /
func NewServer(cfg *config.Config) *Server {
	app := fiber.New(fiber.Config{
		AppName: "go-fiber-oauth",
	})

	server := &Server{
		Config: cfg,
		Http:   app,
	}

	return server
}

func (s *Server) Run() {
	s.registerRoutes()

	s.Http.Listen(s.Config.Server.Address)
}

func (s *Server) registerRoutes() {
	s.Http.Get("/docs/*", swagger.HandlerDefault)

	authGroup := s.Http.Group("/auth")
	authHandler := handlers.NewAuthHandler(s.Config)
	authGroup.Get("/", authHandler.GetHelloWorld)

	// GitHub
	authGroup.Get("/github", authHandler.GetGitHub)
	authGroup.Get("/github/callback", authHandler.GetGitHubCallback)
}

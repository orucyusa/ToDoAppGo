package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

type Server struct {
	app *fiber.App
}

func New() *Server {
	app := fiber.New()
	app.Use(cors.New())
	return &Server{
		app: app,
	}
}

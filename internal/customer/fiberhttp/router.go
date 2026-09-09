package fiberhttp

import "github.com/gofiber/fiber/v3"

func RegisterRoutes(
	app fiber.Router,
	handler *Handler,
) {
	customerApi := app.Group("/customers")
	customerApi.Post("/login", handler.Login)
	customerApi.Post("/register", handler.Register)
}

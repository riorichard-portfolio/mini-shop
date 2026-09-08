package fiberhttp

import "github.com/gofiber/fiber/v3"

func RegisterRoutes(
	app fiber.Router,
	handler *Handler,
) {
	sellerApi := app.Group("/sellers")
	sellerApi.Post("/login", handler.Login)
	sellerApi.Post("/register", handler.Register)
}

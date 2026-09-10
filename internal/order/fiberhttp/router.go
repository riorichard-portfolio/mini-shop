package fiberhttp

import "github.com/gofiber/fiber/v3"

func RegisterRoutes(
	app fiber.Router,
	handler *Handler,
	sellerAuth fiber.Handler,
	customerAuth fiber.Handler,
) {
	orderApi := app.Group("/orders")

	orderSellerApi := orderApi.Group("/sellers", sellerAuth)
	orderSellerApi.Get("", handler.OrderList)
	orderSellerApi.Post("/complete", handler.CompleteOrder)
	orderSellerApi.Post("/cancel", handler.CancelOrder)

	orderCustomerApi := orderApi.Group("/customers", customerAuth)
	orderCustomerApi.Post("", handler.MakeOrder)
}

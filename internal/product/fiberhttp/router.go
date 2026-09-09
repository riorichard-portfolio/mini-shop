package fiberhttp

import "github.com/gofiber/fiber/v3"


func RegisterRoutes(
	app fiber.Router,
	handler *Handler,
	sellerAuth fiber.Handler,
	customerAuth fiber.Handler,
) {
	productApi := app.Group("/products")

	productSellerApi := productApi.Group("/sellers", sellerAuth)
	productSellerApi.Post("", handler.AddNewProduct)

	productCustomerApi := productApi.Group("/customers", customerAuth)
	productCustomerApi.Get("", handler.BrowseProducts)
}

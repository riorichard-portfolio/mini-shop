package product

import (
	"mini-shop/internal/product/fiberhttp"
	"mini-shop/internal/product/usecase"

	"github.com/gofiber/fiber/v3"
)

func RegisterAPI(
	app *fiber.App,
	usc *usecase.Usecase,
	sellerAuth fiber.Handler,
	customerAuth fiber.Handler,
) {
	handler := fiberhttp.NewHandler(usc)
	fiberhttp.RegisterRoutes(
		app,
		handler,
		sellerAuth,
		customerAuth,
	)
}

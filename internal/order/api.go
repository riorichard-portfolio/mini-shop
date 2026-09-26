package order

import (
	"mini-shop/internal/order/fiberhttp"
	"mini-shop/internal/order/usecase"

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

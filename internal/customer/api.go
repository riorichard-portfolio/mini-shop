package customer

import (
	"mini-shop/internal/customer/fiberhttp"
	"mini-shop/internal/customer/usecase"

	"github.com/gofiber/fiber/v3"
)

func RegisterAPI(
	app *fiber.App,
	usc *usecase.Usecase,
) {
	handler := fiberhttp.NewHandler(usc)
	fiberhttp.RegisterRoutes(app, handler)
}

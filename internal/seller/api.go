package seller

import (
	"mini-shop/internal/seller/fiberhttp"
	"mini-shop/internal/seller/usecase"

	"github.com/gofiber/fiber/v3"
)

func RegisterApi(
	app *fiber.App,
	usc *usecase.Usecase,
) {
	handler := fiberhttp.NewHandler(usc)
	fiberhttp.RegisterRoutes(app, handler)
}
package fiberhttp

import (
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/logger"
	"github.com/gofiber/fiber/v3/middleware/recover"
)

func NewApp(
	errHandler func(c fiber.Ctx, err error) error,
) *fiber.App {
	fiberApp := fiber.New(fiber.Config{
		ErrorHandler: errHandler,
	})
	fiberApp.Use(logger.New())
	fiberApp.Use(recover.New())
	return fiberApp
}

package fiberhttp

import "github.com/gofiber/fiber/v3"

var (
	UnauthorizedAuthErr = fiber.NewError(fiber.StatusUnauthorized)
)
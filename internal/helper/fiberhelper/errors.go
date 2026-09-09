package fiberhelper

import "github.com/gofiber/fiber/v3"

var (
	MiddlewareSellerAuthNotSetErr = fiber.NewError(fiber.StatusInternalServerError, "Middleware seller auth not set yet")
	MiddlewareCustomerAuthNotSetErr = fiber.NewError(fiber.StatusInternalServerError, "Middleware customer auth not set yet")
)

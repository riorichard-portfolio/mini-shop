package fiberhelper

import "github.com/gofiber/fiber/v3"

func GetSellerID(c fiber.Ctx) (string, error) {
	sellerID, ok := c.Locals("sellerID").(string)
	if !ok || sellerID == "" {
		return "", MiddlewareSellerAuthNotSetErr
	}
	return sellerID, nil
}

func GetCustomerID(c fiber.Ctx) (string, error) {
	customerID, ok := c.Locals("customerID").(string)
	if !ok || customerID == "" {
		return "", MiddlewareCustomerAuthNotSetErr
	}
	return customerID, nil
}

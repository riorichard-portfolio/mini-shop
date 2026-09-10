package fiberhttp

import (
	"github.com/gofiber/fiber/v3"

	"mini-shop/internal/order"
	"mini-shop/internal/pkg/helper/fiberhelper"
)

type Handler struct {
	usc *order.Usecase
}

func NewHandler(
	usc *order.Usecase,
) *Handler {
	return &Handler{
		usc: usc,
	}
}

func (h *Handler) MakeOrder(c fiber.Ctx) error {
	var req MakeOrderReq
	err := c.Bind().Body(&req)
	if err != nil {
		return err
	}
	customerID, err := fiberhelper.GetCustomerID(c)
	if err != nil {
		return err
	}
	err = h.usc.MakeOrder(c.Context(), order.MakeOrderInput{
		CustomerID: customerID,
		ProductID:  req.ProductID,
		Quantity:   req.Quantity,
	})
	if err != nil {
		return err
	}

	return c.SendStatus(fiber.StatusCreated)
}

func (h *Handler) OrderList(c fiber.Ctx) error {
	var req OrderListReq
	err := c.Bind().Query(&req)
	if err != nil {
		return err
	}
	sellerID, err := fiberhelper.GetSellerID(c)
	if err != nil {
		return err
	}
	res, err := h.usc.OrderList(c.Context(), order.OrderListInput{
		SellerID: sellerID,
		Limit:    req.Limit,
		Offset:   req.Offset,
	})
	if err != nil {
		return err
	}
	data := make([]OrderListItem, 0, len(res))
	for _, item := range res {
		data = append(data, OrderListItem{
			OrderID:     item.OrderID,
			ProductName: item.ProductName,
			Quantity:    item.Quantity,
		})
	}

	return c.Status(fiber.StatusOK).JSON(data)
}

func (h *Handler) CompleteOrder(c fiber.Ctx) error {
	var req CompleteOrderReq
	err := c.Bind().Body(&req)
	if err != nil {
		return err
	}
	sellerID, err := fiberhelper.GetSellerID(c)
	if err != nil {
		return err
	}
	err = h.usc.CompleteOrder(c.Context(), order.CompleteOrderInput{
		OrderID:  req.OrderID,
		SellerID: sellerID,
	})
	if err != nil {
		return err
	}
	return c.SendStatus(fiber.StatusOK)
}

func (h *Handler) CancelOrder(c fiber.Ctx) error {
	var req CancelOrderReq
	err := c.Bind().Body(&req)
	if err != nil {
		return err
	}
	sellerID, err := fiberhelper.GetSellerID(c)
	if err != nil {
		return err
	}
	err = h.usc.CancelOrder(c.Context(), order.CancelOrderInput{
		OrderID:  req.OrderID,
		SellerID: sellerID,
	})
	if err != nil {
		return err
	}
	return c.SendStatus(fiber.StatusOK)
}

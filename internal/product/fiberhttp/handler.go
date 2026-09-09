package fiberhttp

import (
	"github.com/gofiber/fiber/v3"

	"mini-shop/internal/pkg/helper/fiberhelper"
	"mini-shop/internal/product"
)

type Handler struct {
	usc *product.Usecase
}

func NewHandler(
	usc *product.Usecase,
) *Handler {
	return &Handler{
		usc: usc,
	}
}

func (h *Handler) AddNewProduct(c fiber.Ctx) error {
	var req AddNewProductReq
	err := c.Bind().Body(&req)
	if err != nil {
		return err
	}
	sellerID, err := fiberhelper.GetSellerID(c)
	if err != nil {
		return err
	}
	err = h.usc.AddNewProduct(c.Context(), product.AddNewProductInput{
		SellerID: sellerID,
		Name:     req.Name,
		Stock:    req.Stock,
	})
	if err != nil {
		return err
	}

	return c.SendStatus(fiber.StatusCreated)
}

func (h *Handler) BrowseProducts(c fiber.Ctx) error {
	var req BrowseProductsReq
	err := c.Bind().Query(&req)
	if err != nil {
		return err
	}
	res, err := h.usc.BrowseProducts(c.Context(), product.BrowseProductsInput{
		Limit:  req.Limit,
		Offset: req.Offset,
	})
	if err != nil {
		return err
	}
	resp := make(BrowseProductsResp, 0, len(res))
	for _, item := range res {
		resp = append(resp, ProductItem{
			ID:   item.ID,
			Name: item.Name,
		})
	}
	return c.Status(fiber.StatusOK).JSON(resp)
}

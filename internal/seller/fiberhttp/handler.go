package fiberhttp

import (
	"github.com/gofiber/fiber/v3"

	"mini-shop/internal/seller/usecase"
	"mini-shop/internal/seller/dto"
)

type Handler struct {
	usc *usecase.Usecase
}

func NewHandler(
	usc *usecase.Usecase,
) *Handler {
	return &Handler{
		usc: usc,
	}
}

func (h *Handler) Login(c fiber.Ctx) error {
	var req LoginReq
	err := c.Bind().Body(&req)
	if err != nil {
		return err
	}
	res, err := h.usc.Login(c.Context(), dto.LoginInput{
		Email:    req.Email,
		Password: req.Password,
	})
	if err != nil {
		return err
	}
	return c.Status(fiber.StatusOK).JSON(LoginResp{
		AccessToken: res,
	})
}

func (h *Handler) Register(c fiber.Ctx) error {
	var req RegisterReq
	err := c.Bind().Body(&req)
	if err != nil {
		return err
	}
	err = h.usc.Register(c.Context(), dto.RegisterInput{
		Email:    req.Email,
		Password: req.Password,
	})
	if err != nil {
		return err
	}
	return c.SendStatus(fiber.StatusCreated)
}

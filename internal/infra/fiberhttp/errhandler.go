package fiberhttp

import (
	"fmt"
	"log/slog"

	"mini-shop/internal/pkg/err/bizerr"

	"github.com/cockroachdb/errors"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
	"github.com/iancoleman/strcase"
)

func ErrHandler(c fiber.Ctx, err error) error {
	if err == nil {
		return nil
	}
	var ve validator.ValidationErrors
	if errors.As(err, &ve) {
		errorMap := make(map[string]string, len(ve))
		for _, errItem := range ve {
			snakeCaseField := strcase.ToSnake(errItem.Field())
			if errItem.Param() != "" {
				errorMap[snakeCaseField] = errItem.Tag() + ":" + errItem.Param()
			} else {
				errorMap[snakeCaseField] = errItem.Tag()
			}
		}
		return c.Status(fiber.StatusBadRequest).JSON(ErrorResponse{
			Message: "invalid request",
			Errors:  errorMap,
		})
	}
	var fiberErr *fiber.Error
	if errors.As(err, &fiberErr) {
		return c.Status(fiberErr.Code).JSON(ErrorResponse{
			Message: fiberErr.Message,
		})
	}
	var bizErr *bizerr.BizErr
	if errors.As(err, &bizErr) {
		return c.Status(bizErr.Code).JSON(ErrorResponse{
			Message: bizErr.Message,
		})
	}
	slog.Error("unhandled internal server error",
		"path", c.Path(),
		"method", c.Method(),
		"error", err.Error(),
		"stacktrace", fmt.Sprintf("%+v", err),
	)
	return c.Status(fiber.StatusInternalServerError).JSON(InternalServerError)
}

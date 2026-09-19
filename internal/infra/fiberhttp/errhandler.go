package fiberhttp

import (
	"errors"

	errorwrapper "github.com/cockroachdb/errors"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
)

type FieldErrorResponse struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func ErrHandler(c fiber.Ctx, err error) error {
	if err == nil {
		return nil
	}
	var ve validator.ValidationErrors
	if errors.As(err, &ve) {
		errorList := make([]FieldErrorResponse, 0, len(ve))
		for _, errItem := range ve {
			errorList = append(errorList, FieldErrorResponse{
				Field:   errItem.Field(),
				Message: getValidationMsg(errItem),
			})
		}
		return c.Status(fiber.StatusBadRequest).JSON(errorList)
	}
	var fiberErr *fiber.Error
	if errors.As(err, &fiberErr) {
		return c.Status(fiberErr.Code).JSON(fiberErr.Message)
	}
	if errorwrapper.Unwrap(err) != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}
	return c.Status(fiber.StatusPreconditionFailed).JSON(err.Error())
}

func getValidationMsg(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return "field required"
	case "email":
		return "must be an email"
	default:
		return "invalid input"
	}
}

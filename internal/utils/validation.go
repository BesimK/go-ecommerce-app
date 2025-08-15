package utils

import (
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

var validate = validator.New()

func ParseValidated(ctx *fiber.Ctx, dest any) error {
	if err := ctx.BodyParser(dest); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid request body")
	}

	err := validate.Struct(dest)
	if err == nil {
		return nil
	}

	validationErrors := err.(validator.ValidationErrors)

	var messages []string
	for _, fieldErr := range validationErrors {
		var msg string
		fieldName := fieldErr.Field()

		switch fieldErr.Tag() {
		case "required":
			msg = fmt.Sprintf("%s is required", fieldName)
		case "email":
			msg = fmt.Sprintf("%s must be a valid email", fieldName)
		case "min":
			msg = fmt.Sprintf(
				"%s must be at least %s characters",
				fieldName,
				fieldErr.Param(),
			)
		default:
			msg = fmt.Sprintf("%s is invalid", fieldName)
		}

		messages = append(messages, msg)
	}

	return fiber.NewError(fiber.StatusBadRequest, strings.Join(messages, ", "))
}

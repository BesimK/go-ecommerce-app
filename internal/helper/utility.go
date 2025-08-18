// Package helper
package helper

import (
	"crypto/rand"
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func RandomNumbers(length int) (int, error) {
	if length > 9 {
		length = 9
	}
	const numbers = "0123456789"
	bytes := make([]byte, length)

	if _, err := rand.Read(bytes); err != nil {
		return 0, err
	}

	for i := range bytes {
		bytes[i] = numbers[int(bytes[i])%10]
	}

	n, err := strconv.ParseInt(string(bytes), 10, 32)
	if err != nil {
		return 0, err
	}

	return int(n), nil
}

var validate = validator.New()

func ParseValidated(ctx *fiber.Ctx, dest any) error {
	if err := ctx.BodyParser(dest); err != nil {
		return errors.New("invalid request body")
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

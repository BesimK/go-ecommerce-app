// Package rest
package rest

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func ErrorMessage(ctx *fiber.Ctx, status int, err error) error {
	return ctx.Status(status).JSON(err.Error())
}

func SuccessMessage(ctx *fiber.Ctx, msg string, data any) error {
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": msg,
		"data":    data,
	})
}

func InternalError(ctx *fiber.Ctx, msg string, data any) error {
	return nil
}

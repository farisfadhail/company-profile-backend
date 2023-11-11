package middleware

import (
	"plastindo-back-end/utils"

	"github.com/gofiber/fiber/v2"
)

func Authenticated(ctx *fiber.Ctx) error {
	token := ctx.Get("jwt-token")
	if token == "" {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message" : "UNAUTHENTICATED",
		})
	}

	_, err := utils.VerifyTokenJWT(token)
	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message" : "UNAUTHENTICATED",
			"error" : err.Error(),
		})
	}

	return ctx.Next()
}
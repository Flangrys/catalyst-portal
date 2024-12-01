package middlewares

import (
	"github.com/gofiber/fiber/v2"
)

func AuthMiddleware(c *fiber.Ctx) error {
	return c.Next()
}

func LoginMiddleware(c *fiber.Ctx) error {
	return c.Next()
}

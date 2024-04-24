package api

import "github.com/gofiber/fiber/v2"

func Jwt() fiber.Handler {
	return func(c *fiber.Ctx) error {
		c.Context().SetUserValue("user", "user")
		return c.Next()
	}
}

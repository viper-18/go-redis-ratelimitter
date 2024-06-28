package app

import (
	"github.com/gofiber/fiber/v2"
)

func RateLimitterMiddleware(rl *RateLimitter) fiber.Handler {
	return func(c *fiber.Ctx) error {
		clientID := c.IP()
		allowed, err := rl.AllowRequest(clientID)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Internal Server Error"})
		}

		if !allowed {
			return c.Status(fiber.StatusTooManyRequests).JSON(fiber.Map{"error": "Rate limit exceeded"})
		}

		return c.Next()
	}
}

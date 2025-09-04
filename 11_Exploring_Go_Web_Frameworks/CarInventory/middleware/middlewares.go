package middlewares

import (
	"github.com/gofiber/fiber/v2"
)

//we are using the default logger

func SecurityHeaders(c *fiber.Ctx) error {
	
	c.Response().Header.Add("Content-Security-Policy", "default-src 'self'")
	c.Response().Header.Add("X-Frame-Options", "DENY")
	return c.Next()
}


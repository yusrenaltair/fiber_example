package middlware

import "github.com/gofiber/fiber/v2"

func RemoteProtected() fiber.Handler {
	// var err error
	return func(c *fiber.Ctx) error {
		return c.Next()
	}
}

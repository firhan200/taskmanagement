package middlewares

import (
	"net/http"

	"github.com/firhan200/taskmanagement/utils"
	"github.com/gofiber/fiber/v2"
)

func JwtAuthMiddleware(c *fiber.Ctx) error {
	err := utils.TokenValid(c)
	if err != nil {
		c.Status(http.StatusUnauthorized).JSON("Unauthorized")
		return nil
	}
	return c.Next()
}

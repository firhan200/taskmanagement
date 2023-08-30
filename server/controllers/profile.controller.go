package controllers

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func GetProfiles(c *fiber.Ctx) {
	//get body parser
	c.Status(http.StatusOK).JSON(fiber.Map{
		"full_name": "Firhan",
	})
}

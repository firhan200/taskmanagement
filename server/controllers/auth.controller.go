package controllers

import (
	"net/http"

	"github.com/firhan200/taskmanagement/data"
	"github.com/firhan200/taskmanagement/dto"
	"github.com/firhan200/taskmanagement/utils"
	"github.com/gofiber/fiber/v2"
)

func Login(c *fiber.Ctx) error {
	//get body parser
	loginDto := dto.LoginDto{}
	if err := c.BodyParser(&loginDto); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	//validate
	if loginDto.EmailAddress == "" || loginDto.Password == "" {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Email Address or Password cannot be empty",
		})
	}

	db := data.NewConnection()
	userManager := data.NewUserManager(db)
	u, err := userManager.GetByEmailAddressAndPassword(loginDto.EmailAddress, loginDto.Password)

	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	token, err := utils.GenerateToken(u.ID)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"token": token,
	})
}

func Register(c *fiber.Ctx) error {
	//get body parser
	registerDto := dto.RegisterDto{}
	if err := c.BodyParser(&registerDto); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	//validate
	if registerDto.FullName == "" || registerDto.EmailAddress == "" || registerDto.Password == "" {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Full Name, Email Address or Password cannot be empty",
		})
	}

	db := data.NewConnection()
	userManager := data.NewUserManager(db)

	//save and check if error
	u, err := userManager.Register(
		registerDto.FullName,
		registerDto.EmailAddress,
		registerDto.Password,
	)

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	token, err := utils.GenerateToken(u.ID)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	//get body parser
	return c.Status(http.StatusOK).JSON(fiber.Map{
		"token": token,
	})
}

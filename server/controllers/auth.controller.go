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
		c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
		return nil
	}

	//validate
	if loginDto.EmailAddress == "" || loginDto.Password == "" {
		c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Email Address or Password cannot be empty",
		})
		return nil
	}

	//init new user instance
	u := data.User{
		EmailAddress: loginDto.EmailAddress,
		Password:     loginDto.Password,
	}

	err := u.GetByEmailAddressAndPassword()
	if err != nil {
		c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
		return nil
	}

	token, err := utils.GenerateToken(u.ID)
	if err != nil {
		c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
		return nil
	}

	c.Status(http.StatusOK).JSON(fiber.Map{
		"token": token,
	})
	return nil
}

func Register(c *fiber.Ctx) error {
	//get body parser
	registerDto := dto.RegisterDto{}
	if err := c.BodyParser(&registerDto); err != nil {
		c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
		return nil
	}

	//validate
	if registerDto.FullName == "" || registerDto.EmailAddress == "" || registerDto.Password == "" {
		c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Full Name, Email Address or Password cannot be empty",
		})
		return nil
	}

	//init new user instance
	u := data.User{
		FullName:     registerDto.FullName,
		EmailAddress: registerDto.EmailAddress,
		Password:     registerDto.Password,
	}

	//save and check if error
	_, err := u.Save()
	if err != nil {
		c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
		return nil
	}

	token, err := utils.GenerateToken(u.ID)
	if err != nil {
		c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
		return nil
	}

	//get body parser
	c.Status(http.StatusOK).JSON(fiber.Map{
		"token": token,
	})
	return nil
}

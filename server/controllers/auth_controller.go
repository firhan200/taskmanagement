package controllers

import (
	"net/http"

	"github.com/firhan200/taskmanagement/dto"
	"github.com/firhan200/taskmanagement/repositories"
	"github.com/firhan200/taskmanagement/services"
	"github.com/firhan200/taskmanagement/utils"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type LoginHandler struct {
	db *gorm.DB
}

func NewLoginHandler(db *gorm.DB) *LoginHandler {
	return &LoginHandler{
		db: db,
	}
}

func (lh *LoginHandler) Login() fiber.Handler {
	return func(c *fiber.Ctx) error {
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

		repo := repositories.NewUserRepository(lh.db)
		service := services.NewUserService(repo)
		u, err := service.Login(loginDto.EmailAddress, loginDto.Password)

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
}

func (lh *LoginHandler) Register() fiber.Handler {
	return func(c *fiber.Ctx) error {
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

		repo := repositories.NewUserRepository(lh.db)
		service := services.NewUserService(repo)

		//save and check if error
		u, err := service.Register(
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
}

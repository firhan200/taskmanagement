package controllers

import (
	"net/http"

	"github.com/firhan200/taskmanagement/data"
	"github.com/firhan200/taskmanagement/dto"
	"github.com/firhan200/taskmanagement/repositories"
	"github.com/firhan200/taskmanagement/services"
	"github.com/firhan200/taskmanagement/utils"
	"github.com/gofiber/fiber/v2"
)

type IUserService interface {
	Login(email string, pass string) (*data.User, error)
	Register(name string, email string, pass string) (*data.UserSecure, error)
}

type LoginHandler struct {
	userService IUserService
}

func NewLoginHandler() *LoginHandler {
	db := data.GetConnection()
	repo := repositories.NewUserRepository(db)
	service := services.NewUserService(repo)

	return &LoginHandler{
		userService: service,
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

		u, err := lh.userService.Login(loginDto.EmailAddress, loginDto.Password)

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

		//save and check if error
		u, err := lh.userService.Register(
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

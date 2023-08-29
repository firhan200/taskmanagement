package controllers

import (
	"net/http"

	"github.com/firhan200/taskmanagement/data"
	"github.com/firhan200/taskmanagement/dto"
	"github.com/firhan200/taskmanagement/utils"
	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	//get body parser
	loginDto := dto.LoginDto{}
	if err := c.BindJSON(&loginDto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	//validate
	if loginDto.EmailAddress == "" || loginDto.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Email Address or Password cannot be empty",
		})
		return
	}

	//init new user instance
	u := data.User{
		EmailAddress: loginDto.EmailAddress,
		Password:     loginDto.Password,
	}

	err := u.GetByEmailAddressAndPassword()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	token, err := utils.GenerateToken(u.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
	return
}

func Register(c *gin.Context) {
	//get body parser
	registerDto := dto.RegisterDto{}
	if err := c.BindJSON(&registerDto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	//validate
	if registerDto.FullName == "" || registerDto.EmailAddress == "" || registerDto.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Full Name, Email Address or Password cannot be empty",
		})
		return
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
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	token, err := utils.GenerateToken(u.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	//get body parser
	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}

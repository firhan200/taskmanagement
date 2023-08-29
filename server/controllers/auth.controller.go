package controllers

import (
	"net/http"

	"github.com/firhan200/taskmanagement/dto"
	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	//get body parser
	loginDto := dto.Login{}
	if err := c.BindJSON(&loginDto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	//validate
	if loginDto.EmailAddress == "" || loginDto.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "Email Address or Password cannot be empty",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"token":   "1233456",
	})
	return
}

func Register(c *gin.Context) {
	//get body parser
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"token":   "1233456",
	})
}

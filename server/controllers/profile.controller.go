package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetProfiles(c *gin.Context) {
	//get body parser
	c.JSON(http.StatusOK, gin.H{
		"full_name": "Firhan",
	})
}

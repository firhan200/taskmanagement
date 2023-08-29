package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetTasks(c *gin.Context) {
	//get body parser
	c.JSON(http.StatusOK, gin.H{
		"full_name": "Firhan",
	})
}

func GetTaskById(c *gin.Context) {
	id := c.Param("id")
	//get body parser
	c.JSON(http.StatusOK, gin.H{
		"id": id,
	})
}

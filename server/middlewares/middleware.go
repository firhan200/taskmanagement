package middlewares

import (
	"net/http"

	"github.com/firhan200/taskmanagement/utils"
	"github.com/gin-gonic/gin"
)

func JwtAuthMiddleware(c *gin.Context) {
	err := utils.TokenValid(c)
	if err != nil {
		c.String(http.StatusUnauthorized, "Unauthorized")
		c.Abort()
		return
	}
	c.Next()
}

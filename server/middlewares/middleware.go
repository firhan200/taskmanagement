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

func CORSMiddleware(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
	c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
	c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

	if c.Request.Method == "OPTIONS" {
		c.AbortWithStatus(204)
		return
	}

	c.Next()
}

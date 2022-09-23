package middleware


import (
	"net/http"
	"github.com/gin-gonic/gin"
	Token "github.com/rest_api/http/token"
)

func ValidateLogIn() gin.HandlerFunc {
	return func(c *gin.Context) {

		token := c.Request.Header.Get("Token")

		if token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"status": false,
				"message": "Unauthorized, Missing User",
			})
			return
		}

		check, err := Token.ValidateToken(token)

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"status": false,
				"message": "Unauthorized Account!",
			})
			return
		}

		c.JSON(http.StatusAccepted, gin.H{
			"status": check,
			"message": "Authorized !",
		})
		
		c.Next()
	}
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Max-Age", "86400")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "*")
		c.Writer.Header().Set("Access-Control-Expose-Headers", "Content-Length")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(200)
		} else {
			c.Next()
		}
	}
}
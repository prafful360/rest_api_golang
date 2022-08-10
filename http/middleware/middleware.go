package middleware


import (
	"net/http"

	"github.com/shadowshot-x/micro-product-go/authservice/jwt"
	"github.com/gin-gonic/gin"
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
		secret := "S0m3_R4n90m_sss"
		check := jwt.ValidateToken(token, secret)

		if check != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"status": false,
				"message": "Unauthorized Account!",
			})
			return
		}
		c.JSON(http.StatusAccepted, gin.H{
			"status": true,
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
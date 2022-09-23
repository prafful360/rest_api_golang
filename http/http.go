package http

import (
	"github.com/gin-gonic/gin"
	
)

func INIT() {

	router := gin.Default()

	user := router.Group("users")
	user.GET("", GetAll)
	user.POST("", Insert)
	user.PUT("", Update)
	user.DELETE(("/:id"), Remove)

	router.Run("localhost:8081")
}

package http

import (
	"github.com/gin-gonic/gin"
)


func INIT() {
	router := gin.Default()
	// router.Use(CORSMiddleware())
	// router.Use()
	router.GET("/", GetAllUsers)
	router.Run("localhost:8081")
}

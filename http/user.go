package http

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rest_api/data/user"
)

func GetAllUsers(c *gin.Context) {

	users := user.GetAll()
	fmt.Println(users)
	
	t :=  []int{1,2,3,4,5,6}
	c.JSON(http.StatusOK, gin.H{
		"status": true,
		"data":   t,
	})
	
}

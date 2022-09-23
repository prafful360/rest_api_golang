package http

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	Models "github.com/rest_api/Models"
	"github.com/rest_api/data/user"
)

func GetAll(c *gin.Context) {

	users := user.GetAll()

	c.JSON(http.StatusOK, gin.H{
		"status": true,
		"data":   users,
	})

}

func Insert(c *gin.Context) {

	var user_res Models.User

	if err := c.BindJSON(&user_res); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"data":         nil,
			"status":       false,
			"errorMessage": err.Error(),
		})
		return
	}

	res := user.Insert(user_res)

	c.JSON(http.StatusOK, gin.H{
		"status":       res,
		"errorMessage": "",
	})

}


func Remove(c *gin.Context) {

	id_str := c.Param("id")
	id, _ := strconv.Atoi(id_str)
	res, err := user.Remove(id)

	c.JSON(http.StatusOK, gin.H{
		"status":    res,
		"errorMessage": err,
	})
}

func Update(c *gin.Context) {

	var user_res Models.User

	if err := c.BindJSON(&user_res); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"data":         nil,
			"status":       false,
			"errorMessage": err.Error(),
		})
		return
	}

	res, msg := user.Update(user_res)

	c.JSON(http.StatusOK, gin.H{
		"status":    res,
		"message": msg,
	})
}
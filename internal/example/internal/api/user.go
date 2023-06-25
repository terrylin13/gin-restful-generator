package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/terrylin13/gin-restful-generator/internal/example/internal/model"
	"github.com/terrylin13/gin-restful-generator/internal/example/internal/service"
)

func CreateUser(userService *service.UserService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var user model.User
		c.BindJSON(&user)
		err := userService.CreateUser(&user)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Record creation failed!"})
			return
		}
		c.JSON(http.StatusOK, user)
	}
}

func GetUser(userService *service.UserService) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Params.ByName("id")
		var user model.User
		err := userService.GetUser(&user, id)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Record not found!"})
			return
		}
		c.JSON(http.StatusOK, user)
	}
}

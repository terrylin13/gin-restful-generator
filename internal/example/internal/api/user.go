package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/terrylin13/gin-restful-generator/internal/example/internal/model"
	"github.com/terrylin13/gin-restful-generator/internal/example/internal/service"
)

// @BasePath /api
// PingExample godoc
// @Summary 创建用户
// @Description 创建新用户
// @Tags Users
// @Accept json
// @Produce json
// @Param user body User true "用户信息"
// @Success 201 {object} User
// @Router /user [post]
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

// @BasePath /api
// PingExample godoc
// @Summary 根据ID获取用户
// @Description 根据用户ID获取用户信息
// @Tags Users
// @Accept json
// @Produce json
// @Param id path int true "用户ID"
// @Success 200 {object} User
// @Router /users/{id} [get]
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

package restful

import (
	"github.com/gin-gonic/gin"
	"github.com/terrylin13/gin-restful-generator/internal/example/internal/repository"
	"github.com/terrylin13/gin-restful-generator/internal/example/internal/service"
)

func Register(e *gin.Engine) {
	g := e.Group("/api")
	{
		userRepo := &repository.GormUserRepository{}
		userService := service.NewUserService(userRepo)
		g.POST("/user", CreateUser(userService))
		g.GET("/user/:id", GetUser(userService))
	}

}
